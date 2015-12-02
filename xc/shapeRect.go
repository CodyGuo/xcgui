package xc

import (
	"syscall"
	"unsafe"
)

var (
	xShapeRect_Create           *syscall.Proc
	xShapeRect_SetBorderColor   *syscall.Proc
	xShapeRect_SetFillColor     *syscall.Proc
	xShapeRect_SetRoundAngle    *syscall.Proc
	xShapeRect_GetRoundAngle    *syscall.Proc
	xShapeRect_EnableBorder     *syscall.Proc
	xShapeRect_EnableFill       *syscall.Proc
	xShapeRect_EnableRoundAngle *syscall.Proc
)

func init() {
	xShapeRect_Create = xcDLL.MustFindProc("XShapeRect_Create")
	xShapeRect_SetBorderColor = xcDLL.MustFindProc("XShapeRect_SetBorderColor")
	xShapeRect_SetFillColor = xcDLL.MustFindProc("XShapeRect_SetFillColor")
	xShapeRect_SetRoundAngle = xcDLL.MustFindProc("XShapeRect_SetRoundAngle")
	xShapeRect_GetRoundAngle = xcDLL.MustFindProc("XShapeRect_GetRoundAngle")
	xShapeRect_EnableBorder = xcDLL.MustFindProc("XShapeRect_EnableBorder")
	xShapeRect_Create = xcDLL.MustFindProc("XShapeRect_Create")
	xShapeRect_EnableFill = xcDLL.MustFindProc("XShapeRect_EnableFill")
	xShapeRect_EnableRoundAngle = xcDLL.MustFindProc("XShapeRect_EnableRoundAngle")
}

/*
创建矩形形状对象.

参数:
	x X坐标.
	y Y坐标.
	cx 宽度.
	cy 高度.
	hParent 父对象句柄.
返回:
	返回句柄.
*/
func XShapeRectCreate(x, y, cx, cy int, hParent HXCGUI) HXCGUI {
	ret, _, _ := xShapeRect_Create.Call(
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
func XShapeRectSetBorderColor(hShape HXCGUI, color COLORREF, alpha byte) {
	xShapeRect_SetBorderColor.Call(
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
func XShapeRectSetFillColor(hShape HXCGUI, color COLORREF, alpha byte) {
	xShapeRect_SetFillColor.Call(
		uintptr(hShape),
		uintptr(color),
		uintptr(alpha))
}

/*
设置圆角大小.

参数:
	hShape 形状对象句柄.
	nWidth 圆角宽度.
	nHeight 圆角高度.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XShapeRectSetRoundAngle(hShape HXCGUI, nWidth, nHeight int) {
	xShapeRect_SetRoundAngle.Call(
		uintptr(hShape),
		uintptr(nWidth),
		uintptr(nHeight))
}

/*
获取圆角大小.

参数:
	hShape 形状对象句柄.
	pWidth 圆角宽度.
	pHeight 圆角高度.
*/
func XShapeRectGetRoundAngle(hShape HXCGUI, pWidth, pHeight *int) {
	xShapeRect_GetRoundAngle.Call(
		uintptr(hShape),
		uintptr(unsafe.Pointer(pWidth)),
		uintptr(unsafe.Pointer(pHeight)))
}

/*
启用绘制矩形边框.

参数:
	hShape 形状对象句柄.
	bEnable 是否启用.
*/
func XShapeRectEnableBorder(hShape HXCGUI, bEnable bool) {
	xShapeRect_EnableBorder.Call(
		uintptr(hShape),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用填充矩形.

参数:
	hShape 形状对象句柄.
	bEnable 是否启用.
*/
func XShapeRectEnableFill(hShape HXCGUI, bEnable bool) {
	xShapeRect_EnableFill.Call(
		uintptr(hShape),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用圆角.

参数:
	hShape 形状对象句柄.
	bEnable 是否启用.
*/
func XShapeRectEnableRoundAngle(hShape HXCGUI, bEnable bool) {
	xShapeRect_EnableRoundAngle.Call(
		uintptr(hShape),
		uintptr(BoolToBOOL(bEnable)))
}
