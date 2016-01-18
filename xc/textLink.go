package xc

import (
	"syscall"
	// "unsafe"
)

var (
	xTextLink_Create                 *syscall.Proc
	xTextLink_EnableUnderlineLeave   *syscall.Proc
	xTextLink_EnableUnderlineStay    *syscall.Proc
	xTextLink_SetTextColorStay       *syscall.Proc
	xTextLink_SetUnderlineColorLeave *syscall.Proc
	xTextLink_SetUnderlineColorStay  *syscall.Proc
)

func init() {
	xTextLink_Create = xcDLL.MustFindProc("XTextLink_Create")
	xTextLink_EnableUnderlineLeave = xcDLL.MustFindProc("XTextLink_EnableUnderlineLeave")
	xTextLink_EnableUnderlineStay = xcDLL.MustFindProc("XTextLink_EnableUnderlineStay")
	xTextLink_SetTextColorStay = xcDLL.MustFindProc("XTextLink_SetTextColorStay")
	xTextLink_SetUnderlineColorLeave = xcDLL.MustFindProc("XTextLink_SetUnderlineColorLeave")
	xTextLink_SetUnderlineColorStay = xcDLL.MustFindProc("XTextLink_SetUnderlineColorStay")
}

/*
创建静态文本连接元素.

参数:
	x 元素x坐标.
	y 元素y坐标.
	cx 宽度.
	cy 高度.
	pName 文本内容.*uint16
	hParent 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
返回:
	元素句柄.
*/
func XTextLink_Create(x, y, cx, cy int, pName string, hParent HXCGUI) HELE {
	ret, _, _ := xTextLink_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		StringToUintPtr(pName),
		// uintptr(unsafe.Pointer(pName)),
		uintptr(hParent))

	return HELE(ret)
}

/*
启用下划线,鼠标离开状态.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XTextLink_EnableUnderlineLeave(hEle HELE, bEnable bool) {
	xTextLink_EnableUnderlineLeave.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用下划线,鼠标停留状态.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XTextLink_EnableUnderlineStay(hEle HELE, bEnable bool) {
	xTextLink_EnableUnderlineStay.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
设置文本颜色,鼠标停留状态.

参数:
	hEle 元素句柄.
	color RGB颜色值.
	alpha 透明度.
*/
func XTextLink_SetTextColorStay(hEle HELE, color COLORREF, alpha byte) {
	xTextLink_SetTextColorStay.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))

}

/*
设置下划线颜色,鼠标离开状态.

参数:
	hEle 元素句柄.
	color RGB颜色值.
	alpha 透明度.
*/
func XTextLink_SetUnderlineColorLeave(hEle HELE, color COLORREF, alpha byte) {
	xTextLink_SetUnderlineColorLeave.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))
}

/*
设置下划线颜色,鼠标停留状态.

参数:
	hEle 元素句柄.
	color RGB颜色值.
	alpha 透明度.
*/
func XTextLink_SetUnderlineColorStay(hEle HELE, color COLORREF, alpha byte) {
	xTextLink_SetUnderlineColorStay.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))
}
