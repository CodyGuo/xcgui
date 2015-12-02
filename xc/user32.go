package xc

import (
	"syscall"
	"unsafe"
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

type MSG struct {
	HWnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

var (
	user32 *syscall.DLL
)

var (
	// Functions
	clientToScreen *syscall.Proc

	messageBox       *syscall.Proc
	sendMessage      *syscall.Proc
	getMessage       *syscall.Proc
	translateMessage *syscall.Proc
	dispatchMessage  *syscall.Proc
)

func init() {
	user32 = syscall.MustLoadDLL("User32.dll")

	clientToScreen = user32.MustFindProc("ClientToScreen")
	messageBox = user32.MustFindProc("MessageBoxW")
	sendMessage = user32.MustFindProc("SendMessageW")
	getMessage = user32.MustFindProc("GetMessageW")
	translateMessage = user32.MustFindProc("TranslateMessage")
	dispatchMessage = user32.MustFindProc("DispatchMessageW")
}

func ClientToScreen(hwnd HWND, lpPoint *POINT) bool {
	ret, _, _ := clientToScreen.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(lpPoint)))

	return ret != 0
}

func MessageBox(hWnd HWND, lpText, lpCaption string, uType uint32) int32 {
	ret, _, _ := messageBox.Call(
		uintptr(hWnd),
		StringToUintPtr(lpText),
		StringToUintPtr(lpCaption),
		uintptr(uType))

	return int32(ret)
}

func SendMessage(hWnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := sendMessage.Call(
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam,
	)

	return ret
}

func GetMessage(msg *MSG, hWnd HWND, msgFilterMin, msgFilterMax uint32) BOOL {
	ret, _, _ := getMessage.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hWnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax))

	return BOOL(ret)
}

func TranslateMessage(msg *MSG) bool {
	ret, _, _ := translateMessage.Call(uintptr(unsafe.Pointer(msg)))

	return ret != 0
}

func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := dispatchMessage.Call(
		uintptr(unsafe.Pointer(msg)))

	return ret
}
