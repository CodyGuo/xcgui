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

func (wb *WindowBase) Run() {
    xc.XRunXCGUIFunc()
}

func (wb *WindowBase) Close() {
    xc.XExitXCGUIFunc()
}
