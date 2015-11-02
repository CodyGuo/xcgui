package xc

import (
    "github.com/codyguo/sys"
)

var (
    XCDLL *sys.DLLClass
)

const (
    TRUE  = 1
    FALSE = 0
    NULL  = 0
)

type (
    BOOL    int32
    HRESULT int32

    HWND uintptr
)

// 1.初始化UI库
func init() {
    if sys.FileExist("XCGUI.dll") {
        XCDLL = sys.Dll("XCGUI.DLL")
    } else if sys.FileExist("bin/XCGUI.dll") {
        XCDLL = sys.Dll("bin/XCGUI.DLL")
    } else {
        panic("xcgui library not found")
    }

    ret := XCDLL.Call("XInitXCGUI", sys.TEXT("XCGUI Library For Go"))
    // XCGUI的返回值: true 为 1 ，false 为 0
    if ret != TRUE {
        panic("XInitXCGUI call failed.")
    }
}

// func BoolToBOOL(value bool) BOOL {
//     if value {
//         return 1
//     }

//     return 0
// }

// func UTF16PtrToString(s *uint16) string {
//     if s == nil {
//         return ""
//     }
//     return syscall.UTF16ToString((*[1 << 29]uint16)(unsafe.Pointer(s))[0:])
// }

/*
BOOL WINAPI XInitXCGUI  ( wchar_t *  pText )

初始化界面库.
参数:
pText 保留参数.
返回:成功返回TRUE否则返回FALSE.

*/
