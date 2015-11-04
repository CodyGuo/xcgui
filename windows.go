package xcgui

import (
    "errors"
    // "fmt"
)

import (
    "github.com/codyguo/xcgui/xc"
)

type XCWindowStyle uint

const (
    // 什么也没有
    XCWindowsStyleNothing XCWindowStyle = xc.XC_WINDOW_STYLE_NOTHING
    // top布局,如果指定该属性,默认为绑定标题栏元素
    XCWindowsStyleCaption XCWindowStyle = xc.XC_WINDOW_STYLE_CAPTION
    //边框,指定默认上下左右布局大小,如果没有指定,那么边框布局大小为0
    XCWindowsStyleBorder XCWindowStyle = xc.XC_WINDOW_STYLE_BORDER
    // 窗口居中
    XCWindowsStyleCenter XCWindowStyle = xc.XC_WINDOW_STYLE_CENTER
    // 拖动窗口边框
    XCWindowsStyleDragBorder XCWindowStyle = xc.XC_WINDOW_STYLE_DRAG_BORDER
    // 拖动窗口
    XCWindowsStyleDragWindow XCWindowStyle = xc.XC_WINDOW_STYLE_DRAG_WINDOW
    // 窗口默认样式
    XCWindowsStyleDefault XCWindowStyle = xc.XC_WINDOW_STYLE_DEFAULT
    // 模态窗口样式
    XCWindowsStyleModal XCWindowStyle = xc.XC_WINDOW_STYLE_MODAL
)

type Window interface {
    AsWindowBase() *WindowBase
    Handle() xc.HWND
    Show() error
}

type WindowBase struct {
    window  Window
    hWnd    xc.HWND
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
    wb.hWindow = xc.XWndCreate(0, 0, width, height, title, 0, uint32(XCWindowsStyleDefault))
    xc.CloseBtn(wb.hWindow)

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
    return wb.hWnd
}
