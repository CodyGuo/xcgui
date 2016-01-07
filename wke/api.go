package wke

import (
	"os"
	"runtime"
	"syscall"
)

const (
	wkeBroDll = "wkeBrowser.dll"
	wkeDll    = "wke.dll"
)

var (
	wkeWebDLL *syscall.DLL
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if FileExist(wkeBroDll) {
		syscall.MustLoadDLL(wkeDll)
		wkeWebDLL = syscall.MustLoadDLL(wkeBroDll)
	} else if FileExist("lib/" + wkeBroDll) {
		syscall.MustLoadDLL("lib/" + wkeDll)
		wkeWebDLL = syscall.MustLoadDLL("lib/" + wkeBroDll)
	} else if FileExist("../lib/" + wkeBroDll) {
		syscall.MustLoadDLL("../lib/" + wkeDll)
		wkeWebDLL = syscall.MustLoadDLL("../lib/" + wkeBroDll)
	} else if FileExist("../../lib/" + wkeBroDll) {
		syscall.MustLoadDLL("../../lib/" + wkeDll)
		wkeWebDLL = syscall.MustLoadDLL("../../lib/" + wkeBroDll)
	} else {
		panic("wkeBrowser library not found,wkeBrowser.dll or ./lib/wkeBrowser.dll or ../lib/wkeBrowser.dll.")
	}
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	} else {
		return true
	}
}
