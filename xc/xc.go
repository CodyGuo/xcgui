package xc

import (
    "syscall"
    // "fmt"
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

var (
    // Functions
    XWnd_Create     *syscall.Proc
    XWnd_GetHWND    *syscall.Proc
    XWnd_ShowWindow *syscall.Proc
    XRunXCGUI       *syscall.Proc
    XExitXCGUI      *syscall.Proc
    XBtn_SetType    *syscall.Proc
)

func init() {
    // Functions
    XWnd_Create = XCDLL.MustFindProc("XWnd_Create")
    XWnd_ShowWindow = XCDLL.MustFindProc("XWnd_ShowWindow")

    XWnd_GetHWND = XCDLL.MustFindProc("XWnd_GetHWND")
    XBtn_SetType = XCDLL.MustFindProc("XBtn_SetType")

    XRunXCGUI = XCDLL.MustFindProc("XRunXCGUI")
    XExitXCGUI = XCDLL.MustFindProc("XExitXCGUI")
}

// 创建窗口
func XWndCreate(x, y, cx, cy int, pTitle string, hWndParent HWND, XCStyle uint32) HWINDOW {
    // API: XWnd_Create
    ret, _, _ := XWnd_Create.Call(
        uintptr(x),
        uintptr(y),
        uintptr(cx),
        uintptr(cy),
        StringToUintPtr(pTitle),
        uintptr(hWndParent),
        uintptr(XCStyle))

    return HWINDOW(ret)
}

// // 创建窗口,增强功能.
// func XWndCreateEx(x, y, cx, cy int, pTitle string, hWnParent int, XCStyle int) *XCMainWindow {
//     xc := new(XCMainWindow)

//     // API: XWnd_CreateEx
//     xc.Hwnd = XCDLL.Call("XWnd_CreateEx")
// }

// @Author: cody.guo
// @Date: 2015-11-4 23:46:41
// @Function: XWndShowWindow
// @Description: 显示窗口.
// @Calls: XWnd_ShowWindow
// @Input: hWindow: [窗口句柄], nCmdShow: [XWndShowWindow constants]
// @Return: 成功返回TRUE否则返回FALSE.
func XWndShowWindow(hWindow HWINDOW, nCmdShow int) bool {
    ret, _, _ := XWnd_ShowWindow.Call(
        uintptr(hWindow),
        uintptr(nCmdShow))
    // MSDN上返回值：true 为 0，false 为 1
    return ret != 1
}

// 运行程序
func XRunXCGUIFunc() {
    XRunXCGUI.Call()
}

// 释放UI库
func XExitXCGUIFunc() {
    XExitXCGUI.Call()
}

// 关闭
func CloseBtn(hWindow HWINDOW) {
    xBtnSetType(XBtnCreate(10, 5, 35, 20, "关闭", HXCGUI(hWindow)),
        BUTTON_TYPE_CLOSE)
}

// 最小化
func MinBtn(hWindow HWINDOW) {
    xBtnSetType(XBtnCreate(60, 5, 45, 20, "最小化", HXCGUI(hWindow)),
        BUTTON_TYPE_MIN)
}

// 获取系统窗口句柄
func XWndGetHWND(hWindow HWINDOW) HWND {
    ret, _, _ := XWnd_GetHWND.Call(uintptr(hWindow))

    return HWND(ret)
}

// 设置按钮功能类型
func xBtnSetType(hEle HELE, nType uint32) {
    XBtn_SetType.Call(
        uintptr(hEle),
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
