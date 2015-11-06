package xc

import (
    "syscall"
)

var (
    XPane_Create *syscall.Proc
)

func init() {
    XPane_Create = XCDLL.MustFindProc("XPane_Create")
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-6 12:53:14
// @Function: XPaneCreate
// @Description: 创建窗格元素.
// @Calls: XPane_Create
// @Input: pName 窗格标题. nWidth 宽度. nHeight 高度. hFrameWnd 框架窗口.
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
