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

// 创建框架窗口.
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

// 创建窗格元素.
func XPaneCreate(pName string, nWidth, nHeight int, hFrameWnd HWINDOW) HELE {
    // API: XPane_Create
    ret, _, _ := XPane_Create.Call(
        StringToUintPtr(pName),
        uintptr(nWidth),
        uintptr(nHeight),
        uintptr(hFrameWnd))

    return HELE(ret)
}

// @Author: cody.guo
// @Date: 2015-11-4 23:46:41
// @Function: XFrameWndAddPane
// @Description: 添加窗格到框架窗口.
// @Calls: XFrameWnd_AddPane
// @Input: align -> align_type_
// @Return: 成功返回TRUE否则返回FALSE.
func XFrameWndAddPane(hWindow HWINDOW, hPaneDest HELE, hPaneNew HELE, align int32) bool {
    // API: XFrameWnd_AddPane
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
