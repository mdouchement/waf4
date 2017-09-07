#!/bin/sh

WORKDIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
E_VERSION="v1.7.5" # Edit also util/provisionner_<GOARCH>.go file
A_VERSION="v0.6.0" # Edit also util/provisionner_<GOARCH>.go file

mkdir ${WORKDIR}/vendor

echo "Darwin build"
cd ${WORKDIR}/vendor
rm *
curl -LO https://github.com/electron/electron/releases/download/${E_VERSION}/electron-${E_VERSION}-darwin-x64.zip
curl -LO https://github.com/asticode/astilectron/archive/${A_VERSION}.zip
cd $WORKDIR
go-bindata -pkg util -o util/vendor_darwin.go vendor/

echo "Windows build"
cd ${WORKDIR}/vendor
rm *
curl -LO https://github.com/electron/electron/releases/download/${E_VERSION}/electron-${E_VERSION}-win32-ia32.zip
curl -LO https://github.com/asticode/astilectron/archive/${A_VERSION}.zip
cd $WORKDIR
go-bindata -pkg util -o util/vendor_windows.go vendor/

echo "Linux build"
cd ${WORKDIR}/vendor
rm *
curl -LO https://github.com/electron/electron/releases/download/${E_VERSION}/electron-${E_VERSION}-linux-x64.zip
curl -LO https://github.com/asticode/astilectron/archive/${A_VERSION}.zip
cd $WORKDIR
go-bindata -pkg util -o util/vendor_linux.go vendor/

rm -rf ${WORKDIR}/vendor
