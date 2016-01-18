package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xShapeGroupBox_Create           *syscall.Proc
	xShapeGroupBox_SetBorderColor   *syscall.Proc
	xShapeGroupBox_SetTextColor     *syscall.Proc
	xShapeGroupBox_SetFontX         *syscall.Proc
	xShapeGroupBox_SetTextOffset    *syscall.Proc
	xShapeGroupBox_SetRoundAngle    *syscall.Proc
	xShapeGroupBox_SetText          *syscall.Proc
	xShapeGroupBox_GetTextOffset    *syscall.Proc
	xShapeGroupBox_GetRoundAngle    *syscall.Proc
	xShapeGroupBox_EnableRoundAngle *syscall.Proc
)

func init() {
	// Functions
	xShapeGroupBox_Create = xcDLL.MustFindProc("XShapeGroupBox_Create")
	xShapeGroupBox_SetBorderColor = xcDLL.MustFindProc("XShapeGroupBox_SetBorderColor")
	xShapeGroupBox_SetTextColor = xcDLL.MustFindProc("XShapeGroupBox_SetTextColor")
	xShapeGroupBox_SetFontX = xcDLL.MustFindProc("XShapeGroupBox_SetFontX")
	xShapeGroupBox_SetTextOffset = xcDLL.MustFindProc("XShapeGroupBox_SetTextOffset")
	xShapeGroupBox_SetRoundAngle = xcDLL.MustFindProc("XShapeGroupBox_SetRoundAngle")
	xShapeGroupBox_SetText = xcDLL.MustFindProc("XShapeGroupBox_SetText")
	xShapeGroupBox_GetTextOffset = xcDLL.MustFindProc("XShapeGroupBox_GetTextOffset")
	xShapeGroupBox_GetRoundAngle = xcDLL.MustFindProc("XShapeGroupBox_GetRoundAngle")
	xShapeGroupBox_EnableRoundAngle = xcDLL.MustFindProc("XShapeGroupBox_EnableRoundAngle")
}

/*
创建组框形状对象.

参数:
	x X坐标.
	y Y坐标.
	cx 宽度.
	cy 高度.
	pName 名称.*uint16
	hParent 父对象句柄.

返回:
	返回句柄.
*/
func XShapeGroupBox_Create(x, y, cx, cy int, pName string, hParent HXCGUI) HXCGUI {
	ret, _, _ := xShapeGroupBox_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		StringToUintPtr(pName),
		// uintptr(unsafe.Pointer(pName)),
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
func XShapeGroupBox_SetBorderColor(hShape HXCGUI, color COLORREF, alpha byte) {
	xShapeGroupBox_SetBorderColor.Call(
		uintptr(hShape),
		uintptr(color),
		uintptr(alpha))
}

/*
设置文本内容.

参数:
	hShape 形状对象句柄.
	pText 文本内容.*uint16
*/
func XShapeGroupBox_SetTextColor(hShape HXCGUI, pText string) {
	xShapeGroupBox_SetTextColor.Call(
		uintptr(hShape),
		StringToUintPtr(pText))
	// uintptr(unsafe.Pointer(pText)))
}

/*
设置字体.

参数:
	hShape 形状对象句柄.
	hFontX 炫彩字体.
*/
func XShapeGroupBox_SetFontX(hShape HXCGUI, hFontX HFONTX) {
	xShapeGroupBox_SetFontX.Call(
		uintptr(hShape),
		uintptr(hFontX))
}

/*
设置文本偏移量.

参数:
	hShape 形状对象句柄.
	offsetX 水平偏移.
	offsetY 垂直偏移.
*/
func XShapeGroupBox_SetTextOffset(hShape HXCGUI, offsetX, offsetY int) {
	xShapeGroupBox_SetTextOffset.Call(
		uintptr(hShape),
		uintptr(offsetX),
		uintptr(offsetY))
}

/*
设置圆角大小.

参数:
	hShape 形状对象句柄.
	nWidth 圆角宽度.
	nHeight 圆角高度.
*/
func XShapeGroupBox_SetRoundAngle(hShape HXCGUI, nWidth, nHeight int) {
	xShapeGroupBox_SetRoundAngle.Call(
		uintptr(hShape),
		uintptr(nWidth),
		uintptr(nHeight))
}

/*
设置文本内容.

参数:
	hShape 形状对象句柄.
	pText 文本内容.*uint16
*/
func XShapeGroupBox_SetText(hShape HXCGUI, pText string) {
	xShapeGroupBox_SetText.Call(
		uintptr(hShape),
		StringToUintPtr(pText))
	// uintptr(unsafe.Pointer(pText)))
}

/*
获取文本偏移量.

参数:
	hShape 形状对象句柄.
	pOffsetX X坐标偏移量.
	pOffsetY Y坐标偏移量.
*/
func XShapeGroupBox_GetTextOffset(hShape HXCGUI, pOffsetX, pOffsetY *int) {
	xShapeGroupBox_GetTextOffset.Call(
		uintptr(hShape),
		uintptr(unsafe.Pointer(pOffsetX)),
		uintptr(unsafe.Pointer(pOffsetY)))
}

/*
获取圆角大小.

参数:
	hShape 形状对象句柄.
	pWidth 返回圆角宽度.
	pHeight 返回圆角高度.
*/
func XShapeGroupBox_GetRoundAngle(hShape HXCGUI, pWidth, pHeight *int) {
	xShapeGroupBox_GetRoundAngle.Call(
		uintptr(hShape),
		uintptr(unsafe.Pointer(pWidth)),
		uintptr(unsafe.Pointer(pHeight)))
}

/*
启用圆角.

参数:
	hShape 形状对象句柄.
	bEnable 是否启用.
*/
func XShapeGroupBox_EnableRoundAngle(hShape HXCGUI, bEnable bool) {
	xShapeGroupBox_EnableRoundAngle.Call(
		uintptr(hShape),
		uintptr(BoolToBOOL(bEnable)))
}
