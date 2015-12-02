package xc

import (
	"syscall"
	"unsafe"
)

// Align_type_
type Align_type_ int32

const (
	Align_error Align_type_ = -1
	Align_left  Align_type_ = iota - 1
	Align_top
	Align_right
	Align_bottom
	// Align_center
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
	xFrameWnd_Create = xcDLL.MustFindProc("XFrameWnd_Create")
	xFrameWnd_CreateEx = xcDLL.MustFindProc("XFrameWnd_CreateEx")
	xFrameWnd_GetLayoutAreaRect = xcDLL.MustFindProc("XFrameWnd_GetLayoutAreaRect")
	xFrameWnd_SetView = xcDLL.MustFindProc("XFrameWnd_SetView")
	xFrameWnd_SetPaneSplitBarColor = xcDLL.MustFindProc("XFrameWnd_SetPaneSplitBarColor")
	xFrameWnd_AddPane = xcDLL.MustFindProc("XFrameWnd_AddPane")
	xFrameWnd_MergePane = xcDLL.MustFindProc("XFrameWnd_MergePane")
}

/*
创建框架窗口

参数:
	x 窗口左上角x坐标.
	y 窗口左上角y坐标.
	cx 窗口宽度.
	cy 窗口高度.
	pTitle 窗口标题.
	hWndParent 父窗口.
	XCStyle GUI库窗口样式,样式请参见宏定义 xc_window_style_.
返回:
	GUI库窗口资源句柄.
*/
func XFrameWndCreate(x, y, cx, cy int, pTitle *uint16, hWndParent HWND, XCStyle int) HWINDOW {
	ret, _, _ := xFrameWnd_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(unsafe.Pointer(pTitle)),
		uintptr(hWndParent),
		uintptr(XCStyle))

	return HWINDOW(ret)
}

/*
创建框架窗口,增强功能.

参数:
	dwExStyle 窗口扩展样式.
	lpClassName 窗口类名.
	lpWindowName 窗口名.
	dwStyle 窗口样式
	x 窗口左上角x坐标.
	y 窗口左上角y坐标.
	cx 窗口宽度.
	cy 窗口高度.
	hWndParent 父窗口.
	XCStyle GUI库窗口样式,样式请参见宏定义 xc_window_style_.
返回:
	GUI库窗口资源句柄.
*/
func XFrameWndCreateEx(dwExStyle uint32, lpClassName, lpWindowName *uint16, dwStyle uint32, x, y, cx, cy int, hWndParent HWND, XCStyle uint32) HWINDOW {
	ret, _, _ := xFrameWnd_CreateEx.Call(
		uintptr(dwExStyle),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hWndParent),
		uintptr(XCStyle))

	return HWINDOW(ret)
}

/*
用来布局窗格的区域坐标,不包含码头.

参数:
	hWindow 窗口句柄.
	pRect 返回坐标.
*/
func XFrameWndGetLayoutAreaRect(hWindow HWINDOW, pRect *RECT) {
	xFrameWnd_GetLayoutAreaRect.Call(
		uintptr(hWindow),
		uintptr(unsafe.Pointer(pRect)))
}

/*
设置主视图元素.

参数:
	hWindow 窗口句柄.
	hEle 元素句柄.
*/
func XFrameWndSetView(hWindow HWINDOW, hEle HELE) {
	xFrameWnd_SetView.Call(
		uintptr(hWindow),
		uintptr(hEle))
}

/*
设置窗格分隔条颜色.

参数:
	hWindow 窗口句柄.
	color RGB颜色值.
	alpha 透明度.
*/
func XFrameWndSetPaneSplitBarColor(hWindow HWINDOW, color COLORREF, alpha byte) {
	xFrameWnd_SetPaneSplitBarColor.Call(
		uintptr(hWindow),
		uintptr(color),
		uintptr(alpha))
}

/*
添加窗格到框架窗口.

参数:
	hWindow 窗口句柄.
	hPaneDest 目标窗格.
	hPaneNew 当前窗格.
	align 对齐方式.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XFrameWndAddPane(hWindow HWINDOW, hPaneDest HELE, hPaneNew HELE, align Align_type_) bool {
	ret, _, _ := xFrameWnd_AddPane.Call(
		uintptr(hWindow),
		uintptr(hPaneDest),
		uintptr(hPaneNew),
		uintptr(align))

	return ret == TRUE
}

/*
合并窗格.

参数:
	hWindow 窗口句柄.
	hPaneDest 目标窗格.
	hPaneNew 当前窗格.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XFrameWndMergePane(hWindow HWINDOW, hPaneDest, hPaneNew HELE) bool {
	ret, _, _ := xFrameWnd_MergePane.Call(
		uintptr(hWindow),
		uintptr(hPaneDest),
		uintptr(hPaneNew))

	return ret == TRUE
}
