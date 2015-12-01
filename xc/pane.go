package xc

import (
	"syscall"
	"unsafe"
)

var (
	xPane_Create     *syscall.Proc
	xPane_SetView    *syscall.Proc
	xPane_GetTitle   *syscall.Proc
	xPane_IsShowPane *syscall.Proc
	xPane_HidePane   *syscall.Proc
	xPane_ShowPane   *syscall.Proc
	xPane_DockPane   *syscall.Proc
	xPane_LockPane   *syscall.Proc
	xPane_FloatPane  *syscall.Proc
)

func init() {
	xPane_Create = xcDLL.MustFindProc("XPane_Create")
	xPane_SetView = xcDLL.MustFindProc("XPane_SetView")
	xPane_GetTitle = xcDLL.MustFindProc("XPane_GetTitle")
	xPane_IsShowPane = xcDLL.MustFindProc("XPane_IsShowPane")
	xPane_HidePane = xcDLL.MustFindProc("XPane_HidePane")
	xPane_ShowPane = xcDLL.MustFindProc("XPane_ShowPane")
	xPane_DockPane = xcDLL.MustFindProc("XPane_DockPane")
	xPane_LockPane = xcDLL.MustFindProc("XPane_LockPane")
	xPane_FloatPane = xcDLL.MustFindProc("XPane_FloatPane")

}

/*
创建窗格元素.

参数:
	pName 窗格标题.
	nWidth 宽度.
	nHeight 高度.
	hFrameWnd 框架窗口.
返回:
	元素句柄.
*/
func XPaneCreate(pName string, nWidth int, nHeight int, hFrameWnd HWINDOW) HELE {
	ret, _, _ := xPane_Create.Call(
		StringToUintPtr(pName),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hFrameWnd))

	return HELE(ret)
}

/*
设置窗格视图元素.

参数:
	hEle 元素句柄.
	hView 视图颜色.
*/
func XPaneSetView(hEle HELE, hView HELE) {
	xPane_SetView.Call(
		uintptr(hEle),
		uintptr(hView))
}

/*
获取标题.

参数:
	hEle 元素句柄.
	pOut 接收内容缓冲区.
	nOutLen 缓冲区长度,字符为单位.
*/
func XPaneGetTitle(hEle HELE, pOut *uint16, nOutLen int) {
	xPane_GetTitle.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))
}

/*
判断窗格是否显示.

参数:
	hEle 元素句柄.
返回:
	如果显示返回TRUE否则返回FALSE.
*/
func XPaneIsShowPane(hEle HELE) bool {
	ret, _, _ := xPane_IsShowPane.Call(uintptr(hEle))

	return ret == TRUE
}

/*
隐藏窗格.

参数:
	hEle 元素句柄.
*/
func XPaneHidePane(hEle HELE) {
	xPane_HidePane.Call(uintptr(hEle))
}

/*
显示窗格.

参数:
	hEle 元素句柄.
*/
func XPaneShowPane(hEle HELE) {
	xPane_ShowPane.Call(uintptr(hEle))
}

/*
窗格停靠到码头.

参数:
	hEle 元素句柄.
*/
func XPaneDockPane(hEle HELE) {
	xPane_DockPane.Call(uintptr(hEle))
}

/*
锁定窗格.

参数:
	hEle 元素句柄.
*/
func XPaneLockPane(hEle HELE) {
	xPane_LockPane.Call(uintptr(hEle))
}

/*
浮动窗格.

参数:
	hEle 元素句柄.
*/
func XPaneFloatPane(hEle HELE) {
	xPane_FloatPane.Call(uintptr(hEle))
}

func XPaneGetTitleGo(hEle HELE) string {
	buf_szize := 256
	buf := make([]uint16, buf_szize)
	XPaneGetTitle(hEle, &buf[0], buf_szize)

	return syscall.UTF16ToString(buf)
}
