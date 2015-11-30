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

func XPaneCreate(pName string, nWidth int, nHeight int, hFrameWnd HWINDOW) HELE {
	ret, _, _ := xPane_Create.Call(
		StringToUintPtr(pName),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hFrameWnd))
	return HELE(ret)
}

func XPaneSetView(hEle HELE, hView HELE) {
	xPane_SetView.Call(
		uintptr(hEle),
		uintptr(hView))
}

func XPaneGetTitle(hEle HELE, pOut *uint16, nOutLen int) {
	xPane_GetTitle.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))
}

func XPaneGetTitleGo(hEle HELE) string {
	buf_szize := 256
	buf := make([]uint16, buf_szize)
	XPaneGetTitle(hEle, &buf[0], buf_szize)
	return syscall.UTF16ToString(buf)
}

func XPaneIsShowPane(hEle HELE) bool {
	ret, _, _ := xPane_IsShowPane.Call(uintptr(hEle))
	return ret == TRUE
}

func XPaneHidePane(hEle HELE) {
	xPane_HidePane.Call(uintptr(hEle))
}
func XPaneShowPane(hEle HELE) {
	xPane_ShowPane.Call(uintptr(hEle))
}
func XPaneDockPane(hEle HELE) {
	xPane_DockPane.Call(uintptr(hEle))
}
func XPaneLockPane(hEle HELE) {
	xPane_LockPane.Call(uintptr(hEle))
}
func XPaneFloatPane(hEle HELE) {
	xPane_FloatPane.Call(uintptr(hEle))
}
