package xc

import (
	"syscall"
)

var (
	xShapeLine_Create      *syscall.Proc
	xShapeLine_SetPosition *syscall.Proc
	xShapeLine_SetColor    *syscall.Proc
)

func init() {
	xShapeLine_Create = xcDLL.MustFindProc("XShapeLine_Create")
	xShapeLine_SetPosition = xcDLL.MustFindProc("XShapeLine_SetPosition")
	xShapeLine_SetColor = xcDLL.MustFindProc("XShapeLine_SetColor")
}

/*
创建矩形形状对象.

参数:
	x1 坐标.
	y1 坐标.
	x2 坐标.
	y2 坐标.
	hParent 父对象句柄.
返回:
	返回句柄.
*/
func XShapeLine_Create(x1, y1, x2, y2 int, hParent HXCGUI) HXCGUI {
	ret, _, _ := xShapeLine_Create.Call(
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
		uintptr(hParent))

	return HXCGUI(ret)
}

/*
设置位置.

参数:
	hShape 形状对象句柄.
	x1 坐标.
	y1 坐标.
	x2 坐标.
	y2 坐标.
*/
func XShapeLine_SetPosition(hShape HXCGUI, x1, y1, x2, y2 int) {
	xShapeLine_SetPosition.Call(
		uintptr(hShape),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2))
}

/*
设置直线颜色.

参数:
	hShape 形状对象句柄.
	color RGB颜色值.
	alpha 透明度.
*/
func XShapeLine_SetColor(hShape HXCGUI, color COLORREF, alpha byte) {
	xShapeLine_SetColor.Call(
		uintptr(hShape),
		uintptr(color),
		uintptr(alpha))
}
