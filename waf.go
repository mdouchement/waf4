package main

import (
	"os"
	"path/filepath"

	astilectron "github.com/asticode/go-astilectron"
	astilog "github.com/asticode/go-astilog"
	"github.com/mdouchement/waf4/util"
	"github.com/sirupsen/logrus"
)

var (
	url string
	l   *logrus.Logger
)

func init() {
	url = os.Getenv("WAF_URL")
	if url == "" {
		url = "https://google.fr"
	}

	// Framework logging
	l = logrus.New()
	l.Level = logrus.WarnLevel
	astilog.SetLogger(l)
}

func main() {
	base := filepath.Join(os.TempDir(), "electron-testouille")
	os.MkdirAll(base, 0755)

	app, err := astilectron.New(astilectron.Options{
		AppName:           "Astilectron",
		BaseDirectoryPath: base,
	})
	check(err)

	app.SetProvisioner(util.NewProvisioner())
	defer app.Close()

	app.HandleSignals()

	// Start
	err = app.Start()
	check(err)

	// Create window
	w, err := app.NewWindow(url, &astilectron.WindowOptions{
		Center: astilectron.PtrBool(true),
		Width:  astilectron.PtrInt(1024),
		Height: astilectron.PtrInt(768),
		Frame:  astilectron.PtrBool(true),
	})
	check(err)

	err = w.Create()
	check(err)

	// Blocking pattern
	app.Wait()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
