package xc

import (
	"syscall"
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

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-6 19:35:58
// @Function: XModalWndCreate
// @Description: 创建模态窗口;当模态窗口关闭时,会自动销毁模态窗口资源句柄.
// @Calls: xModalWnd_Create
// @Input: nWidth 宽度. nHeight 高度. pTitle 窗口标题内容. hWndParent 父窗口句柄.
//         XCStyle 炫彩窗口样式,样式请参见api 常量定义 xc_window_style_.
// @Return: 模态窗口句柄.
// *******************************************************************
func XModalWndCreate(bWidth, nHeight int, pTitle string, hWndParent HWND, XCStyle uint32) HWINDOW {
	ret, _, _ := xModalWnd_Create.Call(
		uintptr(bWidth),
		uintptr(nHeight),
		StringToUintPtr(pTitle),
		uintptr(hWndParent),
		uintptr(XCStyle))

	return HWINDOW(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-6 19:45:32
// @Function: XModalWndCreateEx
// @Description: 创建模态窗口,增强功能.
// @Calls: xModalWnd_CreateEx
// @Input: dwExStyle 窗口扩展样式. lpClassName 窗口类名. lpWindowName 窗口名. dwStyle 窗口样式.
//         x 窗口左上角x坐标. y 窗口左上角y坐标. cx 窗口宽度. cy 窗口高度. hWndParent 父窗口.
//         XCStyle GUI库窗口样式,样式请参见api 常量定义 xc_window_style_.
// @Return: GUI库窗口资源句柄.
// *******************************************************************
func XModalWndCreateEx(dwExStyle uint32, lpClassName, lpWindowName string, dwStyle uint32, x, y, cx, cy int, hWndParent HWND, XCStyle uint32) HWINDOW {
	ret, _, _ := xModalWnd_CreateEx.Call(
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

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-6 20:00:43
// @Function: XModalWndEnableAutoClose
// @Description: 是否自动关闭窗口,当窗口失去焦点时.
// @Calls: xModalWnd_EnableAutoClose
// @Input: hWindow 模态窗口句柄. bEnable 开启开关.
// @Return:
// *******************************************************************
func XModalWndEnableAutoClose(hWindow HWINDOW, bEnable bool) {
	xModalWnd_EnableAutoClose.Call(
		uintptr(hWindow),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-6 20:04:05
// @Function: XModalWndDoModal
// @Description: 启动显示模态窗口,当窗口关闭时返回.
// @Calls: xModalWnd_DoModal
// @Input: hWindow 模态窗口句柄.
// @Return: XMB_OK:点击确定按钮退出.XMB_CANCEL:点击取消按钮退出.如果返回0,其他方式退出.
// *******************************************************************
func XModalWndDoModal(hWindow HWINDOW) int {
	ret, _, _ := xModalWnd_DoModal.Call(uintptr(hWindow))

	return int(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-6 20:07:33
// @Function: XModalWndEndModal
// @Description: 结束模态窗口.
// @Calls: xModalWnd_EndModal
// @Input: hWindow 窗口句柄. nResult xModalWnd_DoModal() 返回值.
// @Return:
// *******************************************************************
func XModalWndEndModal(hWindow HWINDOW, nResult int) {
	xModalWnd_EndModal.Call(
		uintptr(hWindow),
		uintptr(nResult))
}
