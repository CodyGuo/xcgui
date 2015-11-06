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

/* enum button_type_
   button_type_default 默认类型
   button_type_check 复选按钮
   button_type_radio 单选按钮
   button_type_close 窗口关闭按钮
   button_type_min 窗口最小化按钮
   button_type_max 窗口最大化还原按钮
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

// 运行消息循环,当炫彩窗口数量为0时退出.
func XRunXCGUIFunc() {
    XRunXCGUI.Call()
}

// 退出界面库释放资源.
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
// @Description: 注册窗口事件,将类成员函数作为事件回调函数.回调函数省略参数窗口自身句柄hWindow.
// @Calls: XEle_RegEventC
// @Input: hWindow 窗口句柄. nEvent 事件类型. memberFunction 类成员函数.
func XWndRegEventC(hWindow HWINDOW, nEvent int, memberFunction CallBack) {
    XWnd_RegEventC.Call(
        uintptr(hWindow),
        uintptr(nEvent),
        syscall.NewCallback(memberFunction))
}
