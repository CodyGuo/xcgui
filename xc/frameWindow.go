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
	xFrameWnd_Create               *syscall.Proc
	xFrameWnd_CreateEx             *syscall.Proc
	xFrameWnd_GetLayoutAreaRect    *syscall.Proc
	xFrameWnd_SetView              *syscall.Proc
	xFrameWnd_SetPaneSplitBarColor *syscall.Proc
	xFrameWnd_AddPane              *syscall.Proc
	xFrameWnd_MergePane            *syscall.Proc
)

func init() {
	// Functions
	xFrameWnd_Create = XCDLL.MustFindProc("XFrameWnd_Create")
	xFrameWnd_CreateEx = XCDLL.MustFindProc("XFrameWnd_CreateEx")
	xFrameWnd_GetLayoutAreaRect = XCDLL.MustFindProc("XFrameWnd_GetLayoutAreaRect")
	xFrameWnd_SetView = XCDLL.MustFindProc("XFrameWnd_SetView")
	xFrameWnd_SetPaneSplitBarColor = XCDLL.MustFindProc("XFrameWnd_SetPaneSplitBarColor")
	xFrameWnd_AddPane = XCDLL.MustFindProc("XFrameWnd_AddPane")
	xFrameWnd_MergePane = XCDLL.MustFindProc("XFrameWnd_MergePane")
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-6 12:59:41
// @Function: XFrameWndCreate
// @Description: 创建框架窗口.
// @Calls: xFrameWnd_Create
// @Input: x 窗口左上角x坐标. y 窗口左上角y坐标. cx 窗口宽度. cy 窗口高度. pTitle 窗口标题. hWndParent 父窗口.
//         XCStyle GUI库窗口样式,样式请参见api 常量定义 xc_window_style_.
// @Return: GUI库窗口资源句柄.
// *******************************************
func XFrameWndCreate(x, y, cx, cy int, pTitle string, hWndParent HWND, XCStyle uint32) HWINDOW {
	ret, _, _ := xFrameWnd_Create.Call(
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
// @Date: 2015-11-6 20:20:46
// @Function: XFrameWndCreateEx
// @Description: 创建框架窗口,增强功能.
// @Calls: xFrameWnd_CreateEx
// @Input: dwExStyle 窗口扩展样式. lpClassName 窗口类名. lpWindowName 窗口名. dwStyle 窗口样式
//         x 窗口左上角x坐标. y 窗口左上角y坐标. cx 窗口宽度. cy 窗口高度. hWndParent 父窗口.
//         XCStyle GUI库窗口样式,样式请参见api 常量定义 xc_window_style_.
// @Return: GUI库窗口资源句柄.
// *******************************************
func XFrameWndCreateEx(dwExStyle uint32, lpClassName, lpWindowName string, dwStyle uint32, x, y, cx, cy int, hWndParent HWND, XCStyle uint32) HWINDOW {
	ret, _, _ := xFrameWnd_CreateEx.Call(
		uintptr(dwExStyle),
		StringToUintPtr(lpClassName),
		StringToUintPtr(lpWindowName),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hWndParent),
		uintptr(XCStyle))

	return HWINDOW(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-6 20:28:09
// @Function: XFrameWndGetLayoutAreaRect
// @Description: 用来布局窗格的区域坐标,不包含码头.
// @Calls: xFrameWnd_GetLayoutAreaRect
// @Input: hWindow 窗口句柄. pRect 返回坐标.
// @Return:
// *******************************************
func XFrameWndGetLayoutAreaRect(hWindow HWINDOW, pRect uint32) {
	xFrameWnd_GetLayoutAreaRect.Call(
		uintptr(hWindow),
		uintptr(pRect))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-6 20:35:53
// @Function: XFrameWndSetView
// @Description: 设置主视图元素.
// @Calls: xFrameWnd_SetView
// @Input: hWindow 窗口句柄. hEle 元素句柄.
// @Return:
// *******************************************
func XFrameWndSetView(hWindow HWINDOW, hEle HELE) {
	xFrameWnd_SetView.Call(
		uintptr(hWindow),
		uintptr(hEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-6 20:38:37
// @Function: XFrameWndSetPaneSplitBarColor
// @Description: 设置窗格分隔条颜色.
// @Calls: xFrameWnd_SetPaneSplitBarColor
// @Input: hWindow 窗口句柄. color RGB颜色值. alpha 透明度.
// @Return:
// *******************************************
func XFrameWndSetPaneSplitBarColor(hWindow HWINDOW, color COLORREF, alpha int) {
	xFrameWnd_SetPaneSplitBarColor.Call(
		uintptr(hWindow),
		uintptr(color),
		uintptr(alpha))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-4 23:46:41
// @Function: XFrameWndAddPane
// @Description: 添加窗格到框架窗口.
// @Calls: xFrameWnd_AddPane
// @Input: hWindow 窗口句柄. hPaneDest 目标窗格. hPaneNew 当前窗格.  align 对齐方式.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XFrameWndAddPane(hWindow HWINDOW, hPaneDest HELE, hPaneNew HELE, align int32) bool {
	ret, _, _ := xFrameWnd_AddPane.Call(
		uintptr(hWindow),
		uintptr(hPaneDest),
		uintptr(hPaneNew),
		uintptr(align))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-6 19:24:15
// @Function: XFrameWndMergePane
// @Description: 合并窗格.
// @Calls:  xFrameWnd_MergePane
// @Input: hWindow 窗口句柄. hPaneDest 目标窗格. hPaneNew 当前窗格.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XFrameWndMergePane(hWindow HWINDOW, hPaneDest, hPaneNew HELE) bool {
	ret, _, _ := xFrameWnd_MergePane.Call(
		uintptr(hWindow),
		uintptr(hPaneDest),
		uintptr(hPaneNew))

	if ret != TRUE {
		return false
	}

	return true
}
