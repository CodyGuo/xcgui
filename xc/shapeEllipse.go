package xc

import (
	"syscall"
)

var (
	xShapeEllipse_Create         *syscall.Proc
	xShapeEllipse_SetBorderColor *syscall.Proc
	xShapeEllipse_SetFillColor   *syscall.Proc
	xShapeEllipse_EnableBorder   *syscall.Proc
	xShapeEllipse_EnableFill     *syscall.Proc
)

func init() {
	xShapeEllipse_Create = xcDLL.MustFindProc("XShapeEllipse_Create")
	xShapeEllipse_SetBorderColor = xcDLL.MustFindProc("XShapeEllipse_SetBorderColor")
	xShapeEllipse_SetFillColor = xcDLL.MustFindProc("XShapeEllipse_SetFillColor")
	xShapeEllipse_EnableBorder = xcDLL.MustFindProc("XShapeEllipse_EnableBorder")
	xShapeEllipse_EnableFill = xcDLL.MustFindProc("XShapeEllipse_EnableFill")
}

/*
创建圆形形状对象.

参数:
	x X坐标.
	y Y坐标.
	cx 宽度.
	cy 高度.
	hParent 父对象句柄.
返回:
	返回句柄.
*/
func XShapeEllipse_Create(x, y, cx, cy int, hParent HXCGUI) HXCGUI {
	ret, _, _ := xShapeEllipse_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HXCGUI(ret)
}

/*
设置边框颜色.

参数:
	hShape 形状对象句柄.
	color RGB颜色值.
	alpha 透明度.
*/
func XShapeEllipse_SetBorderColor(hShape HXCGUI, color COLORREF, alpha byte) {
	xShapeEllipse_SetBorderColor.Call(
		uintptr(hShape),
		uintptr(color),
		uintptr(alpha))
}

/*
设置填充颜色.

参数:
	hShape 形状对象句柄.
	color RGB颜色值.
	alpha 透明度.
*/
func XShapeEllipse_SetFillColor(hShape HXCGUI, color COLORREF, alpha byte) {
	xShapeEllipse_SetFillColor.Call(
		uintptr(hShape),
		uintptr(color),
		uintptr(alpha))
}

/*
启用绘制圆边框.

参数:
	hShape 形状对象句柄.
	bEnable 是否启用.
*/
func XShapeEllipse_EnableBorder(hShape HXCGUI, bEnable bool) {
	xShapeEllipse_EnableBorder.Call(
		uintptr(hShape),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用填充圆.

参数:
	hShape 形状对象句柄.
	bEnable 是否启用.
*/
func XShapeEllipse_EnableFill(hShape HXCGUI, bEnable bool) {
	xShapeEllipse_EnableFill.Call(
		uintptr(hShape),
		uintptr(BoolToBOOL(bEnable)))
}
