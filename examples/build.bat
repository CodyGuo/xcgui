@if exist "rsrc.syso" (
    @del "rsrc.syso"
)

set GOARCH=386
set GOGCCFLAGS=-m32 -mthreads -fmessage-length=0
rsrc -manifest main.manifest -o rsrc.syso
go build
pause
exit
-ldflags="-H windowsgui"
 -ldflags "-s -w"


set GOARCH=386
set GOBIN=D:\go\bin
set GOEXE=.exe
set GOHOSTARCH=386
set GOHOSTOS=windows
set GOOS=windows
set GOPATH=D:\go\gopath
set GORACE=
set GOROOT=D:\go
set GOTOOLDIR=D:\go\pkg\tool\windows_386
set GO15VENDOREXPERIMENT=
set CC=gcc
set GOGCCFLAGS=-m32 -mthreads -fmessage-length=0
set CXX=g++
set CGO_ENABLED=1