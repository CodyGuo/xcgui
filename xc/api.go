package xc

import (
    "os"
    "path/filepath"
    "syscall"
    "unsafe"
)

const (
    TRUE  = 1
    FALSE = 0
    NULL  = 0

    xcDll = "XCGUI.dll"
)

var (
    XCDLL *syscall.DLL
)

var (
    // Functions
    XInitXCGUI *syscall.Proc
)

type (
    BOOL    int32
    HRESULT int32

    HWND uintptr
)

// 1.初始化UI库
func init() {
    if FileExist(xcDll) {
        XCDLL = syscall.MustLoadDLL(xcDll)
    } else if FileExist("bin/" + xcDll) {
        XCDLL = syscall.MustLoadDLL("bin/" + xcDll)
    } else {
        panic("xcgui library not found")
    }

    XInitXCGUI = XCDLL.MustFindProc("XInitXCGUI")
    ret, _, _ := XInitXCGUI.Call(StringToUintPtr("XCGUI Library For Go"))
    // XCGUI的返回值: true 为 1 ，false 为 0
    if ret != TRUE {
        panic("XInitXCGUI call failed.")
    }
}

func StringToUintPtr(str string) uintptr {
    return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
}

func FullPath(path string) (p string) {
    p, _ = filepath.Abs(path)
    return
}

func FileExist(path string) bool {
    _, err := os.Stat(path)
    if err != nil {
        return false
    } else {
        return true
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
