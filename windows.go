package xcgui

import (
    "errors"
)

import (
    "github.com/codyguo/xcgui/xc"
)

type MainWindow struct {
    *xc.XCMainWindow
}

func NewMainWindow(width, height int, title string) *MainWindow {
    mw := new(MainWindow)
    mw.XCMainWindow = xc.XWndCreate(0, 0, width, height, title, 0, xc.XC_WINDOW_STYLE_DEFAULT)

    return mw
}

func (mw *MainWindow) Show() (err error) {
    ret := mw.XWndShowWindow(xc.SW_RESTORE)
    if !ret {
        return errors.New("XWndShowWindow call filed.")
    }

    return nil
}

func (mw *MainWindow) GetHWindow() xc.HWND {
    return xc.HWND(mw.HWindow)
}

func (mw *MainWindow) GetHWND() xc.HWND {
    return xc.HWND(mw.Hwnd)
}

func (mw *MainWindow) Run() {
    mw.XRunXCGUI()
}

func (mw *MainWindow) Close() {
    mw.XExitXCGUI()
}
