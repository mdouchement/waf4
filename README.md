# Web Application Frontend

Simple frontend for webapps.

This project is based on the [Astilectron](https://github.com/asticode/go-astilectron) project.

## Build application

```sh
# Go dependencies
$ go get ./...

# Get Electron stuff
$ go get -u github.com/jteeuwen/go-bindata/...
$ ./get-electron.sh

# Build app
$ go build -o waf *.go
```

## Run application

```sh
$ WAF_URL=https://github.com/mdouchement ./waf
```

## License

**MIT**
