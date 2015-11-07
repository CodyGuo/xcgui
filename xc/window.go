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

var (
    // Functions
    XWnd_Create     *syscall.Proc
    XWnd_GetHWND    *syscall.Proc
    XWnd_ShowWindow *syscall.Proc
    XBtn_SetType    *syscall.Proc
    XWnd_RegEventC  *syscall.Proc
)

func init() {
    // Functions
    XWnd_Create = XCDLL.MustFindProc("XWnd_Create")
    XWnd_ShowWindow = XCDLL.MustFindProc("XWnd_ShowWindow")

    XWnd_GetHWND = XCDLL.MustFindProc("XWnd_GetHWND")
    XBtn_SetType = XCDLL.MustFindProc("XBtn_SetType")

    XWnd_RegEventC = XCDLL.MustFindProc("XWnd_RegEventC")
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

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-4 23:46:41
// @Function: XWndShowWindow
// @Description: 显示窗口.
// @Calls: XWnd_ShowWindow
// @Input: hWindow 窗口句柄. nCmdShow XWndShowWindow constants.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************************************
func XWndShowWindow(hWindow HWINDOW, nCmdShow int) bool {
    ret, _, _ := XWnd_ShowWindow.Call(
        uintptr(hWindow),
        uintptr(nCmdShow))
    // MSDN上返回值：true 为 0，false 为 1
    return ret != 1
}

// 关闭
func CloseBtn(hWindow HWINDOW) {
    xBtnSetType(XBtnCreate(10, 5, 35, 20, "关闭", HXCGUI(hWindow)),
        uint32(button_type_close))
}

// 最小化
func MinBtn(hWindow HWINDOW) {
    xBtnSetType(XBtnCreate(60, 5, 45, 20, "最小化", HXCGUI(hWindow)),
        uint32(button_type_min))
}

// 获取系统窗口句柄
func XWndGetHWND(hWindow HWINDOW) HWND {
    ret, _, _ := XWnd_GetHWND.Call(uintptr(hWindow))

    return HWND(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-6 18:36:20
// @Function: xBtnSetType
// @Description: 设置按钮样式.
// @Calls: XBtn_SetType
// @Input: hEle 元素句柄. nStyle 样式. 参考 button_type_
// *******************************************
func xBtnSetType(hEle HELE, nType uint32) {
    XBtn_SetType.Call(
        uintptr(hEle),
        uintptr(nType))
}

// @Author: cody.guo
// @Date: 2015-11-6 18:45:49
// @Function: XWndRegEventC
// @Description: 注册事件函数C方式,事件函数不省略自身HWINDOW句柄.
// @Calls: XEle_RegEventC
// @Input: hWindow 窗口句柄. nEvent 事件类型. pFun 事件函数.
func XWndRegEventC(hWindow HWINDOW, nEvent int, pFun CallBack) {
    XWnd_RegEventC.Call(
        uintptr(hWindow),
        uintptr(nEvent),
        syscall.NewCallback(pFun))
}
