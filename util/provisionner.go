package util

import (
	"fmt"

	astilectron "github.com/asticode/go-astilectron"
)

const (
	eVersion = "v1.7.5"
	aVersion = "v0.6.0"
)

var (
	ePath string
	aPath string
)

func init() {
	aPath = fmt.Sprintf("vendor/%s.zip", aVersion)
}

func NewProvisioner() astilectron.Provisioner {
	return astilectron.NewDisembedderProvisioner(Asset, aPath, ePath)
}
