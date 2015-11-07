package xcgui

import (
    "errors"
    // "fmt"
)

import (
    "github.com/codyguo/xcgui/xc"
)

type Window interface {
    AsWindowBase() *WindowBase
    Handle() xc.HWND
    Show() error
}

type WindowBase struct {
    window Window
    hWnd   xc.HWND
    name   string

    hWindow xc.HWINDOW
}

// func InitWindow(window, parent Window, className string, style uint32) error {
//     wb := window.AsWindowBase()
//     wb.window = window

//     var hwndParent xc.HWND
//     if parent != nil {
//         hwndParent = parent.Handle()
//     }

//     if wb.hWnd != 0 {
//         fmt.Println("wb.hWnd :", wb.hWnd)
//     }

//     return nil
// }

func NewMainWindow(width, height int, title string) *WindowBase {
    wb := new(WindowBase)
    wb.hWindow = xc.XWndCreate(0, 0, width, height, title, 0, uint32(xc.XC_WINDOW_STYLE_DEFAULT))
    xc.CloseBtn(wb.hWindow)

    xc.XCDebugToFileInfo("2015")

    return wb
}

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

func (wb *WindowBase) Run() {
    xc.XRunXCGUIFunc()
}

func (wb *WindowBase) Close() {
    xc.XExitXCGUIFunc()
}

func (wb *WindowBase) AsWindowBase() *WindowBase {
    return wb
}

func (wb *WindowBase) Handle() xc.HWND {
    wb.hWnd = xc.XWndGetHWND(wb.hWindow)
    return wb.hWnd
}
