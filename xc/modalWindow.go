package xc

import (
	"syscall"
	// "unsafe"
)

var (
	// Functions
	xModalWnd_Create          *syscall.Proc
	xModalWnd_CreateEx        *syscall.Proc
	xModalWnd_EnableAutoClose *syscall.Proc
	xModalWnd_DoModal         *syscall.Proc
	xModalWnd_EndModal        *syscall.Proc
)

func init() {
	// Funtions
	xModalWnd_Create = xcDLL.MustFindProc("XModalWnd_Create")
	xModalWnd_CreateEx = xcDLL.MustFindProc("XModalWnd_CreateEx")
	xModalWnd_EnableAutoClose = xcDLL.MustFindProc("XModalWnd_EnableAutoClose")
	xModalWnd_DoModal = xcDLL.MustFindProc("XModalWnd_DoModal")
	xModalWnd_EndModal = xcDLL.MustFindProc("XModalWnd_EndModal")
}

/*
创建模态窗口;当模态窗口关闭时,会自动销毁模态窗口资源句柄.

参数:
	nWidth 宽度.
	nHeight 高度.
	pTitle 窗口标题内容. *uint16
	hWndParent 父窗口句柄.
	XCStyle 炫彩窗口样式,样式请参见宏定义.
返回:
	模态窗口句柄.
*/
func XModalWndCreate(bWidth, nHeight int, pTitle string, hWndParent HWND, XCStyle int) HWINDOW {
	ret, _, _ := xModalWnd_Create.Call(
		uintptr(bWidth),
		uintptr(nHeight),
		StringToUintPtr(pTitle),
		// uintptr(unsafe.Pointer(pTitle)),
		uintptr(hWndParent),
		uintptr(XCStyle))

	return HWINDOW(ret)
}

/*
创建模态窗口,增强功能.

参数:
	dwExStyle 窗口扩展样式.
	lpClassName 窗口类名.*uint16
	lpWindowName 窗口名.*uint16
	dwStyle 窗口样式.
	x 窗口左上角x坐标.
	y 窗口左上角y坐标.
	cx 窗口宽度.
	cy 窗口高度.
	hWndParent 父窗口.
	XCStyle GUI库窗口样式,样式请参见宏定义 xc_window_style_.
返回:
	GUI库窗口资源句柄.
*/
func XModalWndCreateEx(dwExStyle uint32, lpClassName, lpWindowName string, dwStyle uint32, x, y, cx, cy int, hWndParent HWND, XCStyle int) HWINDOW {
	ret, _, _ := xModalWnd_CreateEx.Call(
		uintptr(dwExStyle),
		StringToUintPtr(lpClassName),
		StringToUintPtr(lpWindowName),
		// uintptr(unsafe.Pointer(lpClassName)),
		// uintptr(unsafe.Pointer(lpWindowName)),
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
是否自动关闭窗口,当窗口失去焦点时.

参数:
	hWindow 模态窗口句柄.
	bEnable 开启开关.
*/
func XModalWndEnableAutoClose(hWindow HWINDOW, bEnable bool) {
	xModalWnd_EnableAutoClose.Call(
		uintptr(hWindow),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启动显示模态窗口,当窗口关闭时返回.

参数:
	hWindow 模态窗口句柄.
返回:
	XMB_OK:点击确定按钮退出.XMB_CANCEL:点击取消按钮退出.如果返回0,其他方式退出.
*/
func XModalWndDoModal(hWindow HWINDOW) int {
	ret, _, _ := xModalWnd_DoModal.Call(uintptr(hWindow))

	return int(ret)
}

/*
结束模态窗口.

参数:
	hWindow 窗口句柄.
	nResult XModalWnd_DoModal() 返回值.
*/
func XModalWndEndModal(hWindow HWINDOW, nResult int) {
	xModalWnd_EndModal.Call(
		uintptr(hWindow),
		uintptr(nResult))
}
