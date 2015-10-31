package xcgui

import (
    "github.com/codyguo/sys"
)

var XCDLL *sys.DLLClass

func init() {
    if sys.FileExist("XCGUI.dll") {
        XCDLL = sys.Dll("XCGUI.DLL")
    } else if sys.FileExist("bin/XCGUI.dll") {
        XCDLL = sys.Dll("bin/XCGUI.DLL")
    } else {
        panic("xcgui library not found")
    }

    XCDLL.Call("XInitXCGUI", sys.TEXT("XCGUI Library For Go"))
}
