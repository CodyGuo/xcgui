package xc

import (
    . "github.com/codyguo/sys"
)

var XCDLL *DLLClass

func init() {
    if FileExist("XCGUI.dll") {
        XCDLL = Dll("XCGUI.DLL")
    } else if FileExist("bin/XCGUI.dll") {
        XCDLL = Dll("bin/XCGUI.DLL")
    } else {
        panic("xcgui library not found")
    }

    XCDLL.Call("XInitXCGUI", TEXT("XCGUI Library For Go"))
}
