package xc

import (
    "syscall"
)

// align_type_
const (
    align_error = -1
    align_left  = iota
    align_top
    align_right
    align_bottom
    align_center
)

var (
    // Functions
    XFrameWnd_Create  *syscall.Proc
    XPane_Create      *syscall.Proc
    XFrameWnd_AddPane *syscall.Proc
)

func init() {
    // Functions
    XFrameWnd_Create = XCDLL.MustFindProc("XFrameWnd_Create")
    XPane_Create = XCDLL.MustFindProc("XPane_Create")
    XFrameWnd_AddPane = XCDLL.MustFindProc("XFrameWnd_AddPane")
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-6 12:59:41
// @Function: XFrameWndCreate
// @Description: 创建框架窗口.
// @Calls: XPane_Create
// @Input: x 窗口左上角x坐标. y 窗口左上角y坐标. cx 窗口宽度. cy 窗口高度. pTitle 窗口标题. hWndParent 父窗口.
//         XCStyle GUI库窗口样式,样式请参见api 常量定义 xc_window_style_.
// @Return: GUI库窗口资源句柄.
// *******************************************
func XFrameWndCreate(x, y, cx, cy int, pTitle string, hWndParent HWND, XCStyle uint32) HWINDOW {
    // API: XFrameWnd_Create
    ret, _, _ := XFrameWnd_Create.Call(
        uintptr(x),
        uintptr(y),
        uintptr(cx),
        uintptr(cy),
        StringToUintPtr(pTitle),
        uintptr(hWndParent),
        uintptr(XCStyle))

    return HWINDOW(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-6 12:53:14
// @Function: XPaneCreate
// @Description: 创建窗格元素.
// @Calls: XPane_Create
// @Input: pName [窗格标题]. nWidth [宽度]. nHeight [高度]. hFrameWnd [框架窗口].
// @Return: 元素句柄.
// *******************************************
func XPaneCreate(pName string, nWidth, nHeight int, hFrameWnd HWINDOW) HELE {
    ret, _, _ := XPane_Create.Call(
        StringToUintPtr(pName),
        uintptr(nWidth),
        uintptr(nHeight),
        uintptr(hFrameWnd))

    return HELE(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-4 23:46:41
// @Function: XFrameWndAddPane
// @Description: 添加窗格到框架窗口.
// @Calls: XFrameWnd_AddPane
// @Input: hWindow [窗口句柄]. hPaneDest [目标窗格]. hPaneNew [当前窗格].  align [对齐方式].
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XFrameWndAddPane(hWindow HWINDOW, hPaneDest HELE, hPaneNew HELE, align int32) bool {
    ret, _, _ := XFrameWnd_AddPane.Call(
        uintptr(hWindow),
        uintptr(hPaneDest),
        uintptr(hPaneNew),
        uintptr(align))

    if ret != TRUE {
        return false
    }

    return true
}
