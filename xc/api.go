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

// xc_window_style_
const (
    XC_WINDOW_STYLE_NOTHING     = 0x00000000
    XC_WINDOW_STYLE_CAPTION     = 0x00000001
    XC_WINDOW_STYLE_BORDER      = 0x00000002
    XC_WINDOW_STYLE_CENTER      = 0x00000004
    XC_WINDOW_STYLE_DRAG_BORDER = 0x00000008
    XC_WINDOW_STYLE_DRAG_WINDOW = 0x00000010
    XC_WINDOW_STYLE_DEFAULT     = 0x00000001 | 0x00000002 | 0x00000004 | 0x00000008
    XC_WINDOW_STYLE_MODAL       = 0x00000001 | 0x00000004 | 0x00000008
)

var (
    XCDLL *syscall.DLL
)

var (
    // Functions
    XInitXCGUI *syscall.Proc
)

// HXCGUI: 其他句柄
type HXCGUI uintptr

// 句柄列表
type (
    // 系统窗口句柄
    HWND HXCGUI
    // 窗口句柄
    HWINDOW HXCGUI
    // UI元素句柄
    HELE HXCGUI

    HMENUX          HXCGUI
    HSTRING         HXCGUI
    HDRAW           HXCGUI
    HARRAY          HXCGUI
    HIMAGE          HXCGUI
    HXMLRES         HXCGUI
    HARRAY_ITERATOR HXCGUI
    HFONTX          HXCGUI
    HPANE_GROUP     HXCGUI
    HBKINFOM        HXCGUI
)

type (
    BOOL    int32
    HRESULT int32
)

// 1.初始化UI库
func init() {
    if FileExist(xcDll) {
        XCDLL = syscall.MustLoadDLL(xcDll)
    } else if FileExist("bin/" + xcDll) {
        XCDLL = syscall.MustLoadDLL("bin/" + xcDll)
    } else if FileExist("../bin/" + xcDll) {
        XCDLL = syscall.MustLoadDLL("../bin/" + xcDll)
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
