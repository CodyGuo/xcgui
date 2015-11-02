package xc

import (
    "github.com/codyguo/sys"
)

import (
// "fmt"
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

// XWndShowWindow constants
const (
    SW_HIDE            = 0
    SW_NORMAL          = 1
    SW_SHOWNORMAL      = 1
    SW_SHOWMINIMIZED   = 2
    SW_MAXIMIZE        = 3
    SW_SHOWMAXIMIZED   = 3
    SW_SHOWNOACTIVATE  = 4
    SW_SHOW            = 5
    SW_MINIMIZE        = 6
    SW_SHOWMINNOACTIVE = 7
    SW_SHOWNA          = 8
    SW_RESTORE         = 9
    SW_SHOWDEFAULT     = 10
    SW_FORCEMINIMIZE   = 11
)

/*enum   button_type_ {
  button_type_default = 0, button_type_check, button_type_radio, button_type_close,
  button_type_min, button_type_max
  }
*/
const (
    BUTTON_TYPE_DEFAULT = iota
    BUTTON_TYPE_CHECK
    BUTTON_TYPE_RADIO
    BUTTON_TYPE_CLOSE
    BUTTON_TYPE_MIN
    BUTTON_TYPE_MAX
)

type XCMainWindow struct {
    Hwnd uintptr
    // WHwnd   uintptr
    HWindow uintptr
}

// 2.创建窗口
func XWndCreate(x, y, cx, cy int, pTitle string, hWnParent int, XCStyle int) *XCMainWindow {
    xc := new(XCMainWindow)
    // API: XWnd_Create
    xc.HWindow = XCDLL.Call("XWnd_Create",
        uintptr(x),
        uintptr(y),
        uintptr(cx),
        uintptr(cy),
        sys.TEXT(pTitle),
        uintptr(hWnParent),
        uintptr(XCStyle)) // xc_window_style_

    xc.closeBtn()
    xc.minBtn()

    return xc
}

// // 创建窗口,增强功能.
// func XWndCreateEx(x, y, cx, cy int, pTitle string, hWnParent int, XCStyle int) *XCMainWindow {
//     xc := new(XCMainWindow)

//     // API: XWnd_CreateEx
//     xc.Hwnd = XCDLL.Call("XWnd_CreateEx")
// }

func (xc *XCMainWindow) XWndGetHWND() {
    xc.Hwnd = XCDLL.Call(" XWnd_GetHWND",
        xc.HWindow)
}

// 3.显示窗口
func (xc *XCMainWindow) XWndShowWindow(uFlag int) bool {
    ret := XCDLL.Call("XWnd_ShowWindow",
        xc.HWindow,
        uintptr(uFlag))
    // MSDN上返回值：true 为 0，false 为 1
    return ret != 1
}

// 4.运行程序
func (xc *XCMainWindow) XRunXCGUI() {
    XCDLL.Call("XRunXCGUI")
}

// 5.释放UI库
func (xc *XCMainWindow) XExitXCGUI() {
    XCDLL.Call("XExitXCGUI")
}

// 关闭
func (xc *XCMainWindow) closeBtn() {
    xBtnSetType(XBtnCreate(10, 5, 35, 20, "关闭", HWND(xc.HWindow)),
        BUTTON_TYPE_CLOSE)
}

// 最小化
func (xc *XCMainWindow) minBtn() {
    xBtnSetType(XBtnCreate(60, 5, 45, 20, "最小化", HWND(xc.HWindow)),
        BUTTON_TYPE_MIN)
}

// 设置按钮功能类型
func xBtnSetType(hEle uintptr, nType int) {
    XCDLL.Call("XBtn_SetType",
        hEle,
        uintptr(nType))
}

/*

	hWindow := XCGUI.Call("XWnd_Create", 100, 100, 600, 400, TEXT("XCGUI Library For GoLang"), 0, 0x1|0x8)
	hBtn := XCGUI.Call("XBtn_Create", 50, 50, 100, 100, TEXT("Hey Man"), hWindow)
	XCGUI.Call("XEle_RegEventC", hBtn, 34, uintptr(syscall.NewCallback(HeyMan)))
	XCGUI.Call("XWnd_ShowWindow", hWindow, 0x5)
	XCGUI.Call("XRunXCGUI")
*/

/*
// API: XWnd_Create
 HWINDOW WINAPI XWnd_Create  ( int  x,
    int  y,
    int  cx,
    int  cy,
    const wchar_t *  pTitle,
    HWND  hWndParent,
    int  XCStyle
)
创建窗口
参数:
x 窗口左上角x坐标.
y 窗口左上角y坐标.
cx 窗口宽度.
cy 窗口高度.
pTitle 窗口标题.
hWndParent 父窗口.
XCStyle GUI库窗口样式,样式请参见宏定义 xc_window_style_.
返回:GUI库窗口资源句柄.

// API: XWnd_ShowWindow
BOOL WINAPI XWnd_ShowWindow  ( HWINDOW  hWindow,
    int  nCmdShow
 )
显示隐藏窗口.
参数:
hWindow 窗口句柄.
nCmdShow 参见MSDN.
返回:参见MSDN.


// API: XBtn_SetType
void WINAPI XBtn_SetType  ( HELE  hEle,
    button_type_  nType
 )
设置按钮功能类型.
参数:
hEle 元素句柄.
nType 按钮类型,参见宏定义.

*/
