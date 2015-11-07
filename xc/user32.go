package xc

import (
    "syscall"
)

// MessageBox constants
const (
    MB_OK                = 0x00000000
    MB_OKCANCEL          = 0x00000001
    MB_ABORTRETRYIGNORE  = 0x00000002
    MB_YESNOCANCEL       = 0x00000003
    MB_YESNO             = 0x00000004
    MB_RETRYCANCEL       = 0x00000005
    MB_CANCELTRYCONTINUE = 0x00000006
    MB_ICONHAND          = 0x00000010
    MB_ICONQUESTION      = 0x00000020
    MB_ICONEXCLAMATION   = 0x00000030
    MB_ICONASTERISK      = 0x00000040
    MB_USERICON          = 0x00000080
    MB_ICONWARNING       = MB_ICONEXCLAMATION
    MB_ICONERROR         = MB_ICONHAND
    MB_ICONINFORMATION   = MB_ICONASTERISK
    MB_ICONSTOP          = MB_ICONHAND
    MB_DEFBUTTON1        = 0x00000000
    MB_DEFBUTTON2        = 0x00000100
    MB_DEFBUTTON3        = 0x00000200
    MB_DEFBUTTON4        = 0x00000300
)

var (
    User32 *syscall.DLL
)

var (
    // Functions
    MessageBoxW  *syscall.Proc
    SendMessageW *syscall.Proc
)

func init() {
    User32 = syscall.MustLoadDLL("User32.dll")
    MessageBoxW = User32.MustFindProc("MessageBoxW")
    SendMessageW = User32.MustFindProc("SendMessageW")
}

func MessageBox(hWnd HWND, lpText, lpCaption string, uType uint32) int32 {
    ret, _, _ := MessageBoxW.Call(
        uintptr(hWnd),
        StringToUintPtr(lpText),
        StringToUintPtr(lpCaption),
        uintptr(uType))

    return int32(ret)
}

func SendMessage(hWnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
    ret, _, _ := SendMessageW.Call(
        uintptr(hWnd),
        uintptr(msg),
        wParam,
        lParam,
    )

    return ret
}
