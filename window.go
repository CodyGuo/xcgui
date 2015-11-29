package xcgui

import (
	"errors"
	"fmt"
	// "syscall"
)

import (
	"github.com/codyguo/xcgui/xc"
)

var (
	hwnd2WindowBase = make(map[xc.HWND]*WindowBase)
)

type Window interface {
	AsWindowBase() *WindowBase
	Handle() xc.HWND
	Show() error

	WndProc(hwnd xc.HWND, msg uint32, wParam, lParam uintptr) uintptr
}

type WindowBase struct {
	window Window
	hWnd   xc.HWND
	hEle   xc.HELE
	name   string

	hWindow xc.HWINDOW
}

// WindowBase simply returns the receiver.
func (wb *WindowBase) AsWindowBase() *WindowBase {
	return wb
}

func InitWindow(window, parent Window, width, height int, title string, style uint32) error {
	wb := window.AsWindowBase()
	wb.window = window

	var hwndParent xc.HWND
	if parent != nil {
		hwndParent = parent.Handle()
	}

	if wb.hWnd != 0 {
		return lastError("XWndCreate")
	}

	wb.hWindow = xc.XWndCreate(
		xc.CW_USEDEFAULT,
		xc.CW_USEDEFAULT,
		width,
		height,
		title,
		hwndParent,
		int(style))

	xc.CloseBtn(wb.hWindow)

	if wb.hWindow == 0 {
		return lastError("XWndCreate")
	}

	succeeded := false
	go func() {
		if !succeeded {
			newErrorNoPanic("XWndCreate")
		}
	}()

	succeeded = true

	return nil
}

func (wb *WindowBase) Handle() xc.HWND {
	wb.hWnd = xc.XWndGetHWND(wb.hWindow)
	return wb.hWnd
}

func (wb *WindowBase) SetMinimumSize(min Size) error {
	if min.Width < 0 || min.Height < 0 {
		return newError("min must be positive")
	}

	xc.XWndSetMinimumSize(wb.hWindow, min.Width, min.Height)

	return nil
}

// func NewMainWindow(width, height int, title string) *WindowBase {
//     wb := new(WindowBase)
//     wb.hWindow = xc.XWndCreate(0, 0, width, height, title, 0, int(xc.XC_WINDOW_STYLE_DEFAULT))
//     xc.CloseBtn(wb.hWindow)
//     // 透明
//     xc.XWndSetTransparentType(wb.hWindow, xc.WINDOW_TRANSPARENT_SIMPLE)
//     // xc.XWndEnableDragBorder(wb.hWindow, xc.BoolToBOOL(false))

//     // xc.XCDebugToFileInfo("2015")

//     return wb
// }

func (wb *WindowBase) Show() (err error) {
	ret := xc.XWndShowWindow(wb.hWindow, xc.SW_RESTORE)
	if !ret {
		return errors.New("XWndShowWindow call filed.")
	}

	return nil
}

func (wb *WindowBase) GetHWindow() xc.HWINDOW {
	return wb.hWindow
}

// func (wb *WindowBase) GetHWND() xc.HWND {
//     return xc.HWND(mw.Hwnd)
// }

func (wb *WindowBase) Run() int {

	// var msg xc.MSG

	// for wb.hWindow != 0 {
	//     switch xc.GetMessage(&msg, 0, 0, 0) {
	//     case 0:
	//         return int(msg.WParam)
	//     case -1:
	//         return -1
	//     }

	//     xc.TranslateMessage(&msg)
	//     xc.DispatchMessage(&msg)

	// }
	xc.XRunXCGUI()

	return 0
}

func (wb *WindowBase) Close() {
	wb.SendMessage(xc.WM_CLOSE, 0, 0)
	xc.XExitXCGUI()
}

// SendMessage sends a message to the window and returns the result.
func (wb *WindowBase) SendMessage(msg uint32, wParam, lParam uintptr) uintptr {
	return xc.SendMessage(wb.hWnd, msg, wParam, lParam)
}

func windowFromHandle(hwnd xc.HWND) Window {
	if wb := hwnd2WindowBase[hwnd]; wb != nil {
		return wb.window
	}

	return nil
}

func defaultWndProc(hwnd xc.HWND, msg uint32, wParam, lParam uintptr) (result uintptr) {

	wi := windowFromHandle(hwnd)

	result = wi.WndProc(hwnd, msg, wParam, lParam)

	return
}

func (wb *WindowBase) WndProc(hwnd xc.HWND, msg uint32, wparam, lparam uintptr) uintptr {
	fmt.Println("WidgetBase.WndProc")
	return uintptr(0)
}

func DebugToFileInfo(str string) {
	// 中文转换为Ansi
	buffer := make([]byte, len(str)+1)
	xc.XCUnicodeToAnsi(xc.StringToUTF16Ptr(str), len(str), &buffer[0], len(str)*2)

	xc.XCDebugToFileInfo(&buffer[0])
}
