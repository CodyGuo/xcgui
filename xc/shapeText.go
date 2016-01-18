package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xShapeText_Create          *syscall.Proc
	xShapeText_SetText         *syscall.Proc
	xShapeText_GetText         *syscall.Proc
	xShapeText_GetTextLength   *syscall.Proc
	xShapeText_SetFont         *syscall.Proc
	xShapeText_SetTextColor    *syscall.Proc
	xShapeText_GetTextColor    *syscall.Proc
	xShapeText_SetTextAlign    *syscall.Proc
	xShapeText_SetOffset       *syscall.Proc
	xShapeText_SetLayoutWidth  *syscall.Proc
	xShapeText_SetLayoutHeight *syscall.Proc
	xShapeText_GetLayoutWidth  *syscall.Proc
	xShapeText_GetLayoutHeight *syscall.Proc
)

func init() {
	xShapeText_Create = xcDLL.MustFindProc("XShapeText_Create")
	xShapeText_SetText = xcDLL.MustFindProc("XShapeText_SetText")
	xShapeText_GetText = xcDLL.MustFindProc("XShapeText_GetText")
	xShapeText_GetTextLength = xcDLL.MustFindProc("XShapeText_GetTextLength")
	xShapeText_SetFont = xcDLL.MustFindProc("XShapeText_SetFont")
	xShapeText_SetTextColor = xcDLL.MustFindProc("XShapeText_SetTextColor")
	xShapeText_GetTextColor = xcDLL.MustFindProc("XShapeText_GetTextColor")
	xShapeText_SetTextAlign = xcDLL.MustFindProc("XShapeText_SetTextAlign")
	xShapeText_SetOffset = xcDLL.MustFindProc("XShapeText_SetOffset")
	xShapeText_SetLayoutWidth = xcDLL.MustFindProc("XShapeText_SetLayoutWidth")
	xShapeText_SetLayoutHeight = xcDLL.MustFindProc("XShapeText_SetLayoutHeight")
	xShapeText_GetLayoutWidth = xcDLL.MustFindProc("XShapeText_GetLayoutWidth")
	xShapeText_GetLayoutHeight = xcDLL.MustFindProc("XShapeText_GetLayoutHeight")
}

/*
创建形状对象文本.

参数:
	x X坐标.
	y Y坐标.
	cx 宽度.
	cy 高度.
	pName 文本内容.*uint16
	hParent 父对象句柄.
返回:
	返回文本块句柄.
*/
func XShapeText_Create(x, y, cx, cy int, pName string, hParent HXCGUI) HXCGUI {
	ret, _, _ := xShapeText_Create.Call(
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
设置文本内容.

参数:
	hTextBlock 形状对象文本句柄.
	pName 文本内容.*uint16
*/
func XShapeText_SetText(hTextBlock HXCGUI, pName string) {
	xShapeText_SetText.Call(
		uintptr(hTextBlock),
		StringToUintPtr(pName))
	// uintptr(unsafe.Pointer(pName)))
}

/*
获取文本内容.

参数:
	hTextBlock 形状对象文本句柄.
	pOut 接收内容缓冲区.
	nOutLen 缓冲区大小,字符数量为单位.
*/
func XShapeText_GetText(hTextBlock HXCGUI, pOut *uint16, nOutLen int) {
	xShapeText_GetText.Call(
		uintptr(hTextBlock),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))
}

//非官方api,功能同上, 更方便golang
func XShapeText_GetTextGo(hTextBlock HXCGUI) string {
	buf_szize := 256
	buf := make([]uint16, buf_szize)
	XShapeText_GetText(hTextBlock, &buf[0], buf_szize)

	return syscall.UTF16ToString(buf)
}

/*
获取文本长度.

参数:
	hTextBlock 形状对象文本句柄.
返回:
	文本长度.
*/
func XShapeText_GetTextLength(hTextBlock HXCGUI) int {
	ret, _, _ := xShapeText_GetTextLength.Call(
		uintptr(hTextBlock))

	return int(ret)
}

/*
设置字体.

参数:
	hTextBlock 形状对象文本句柄.
	hFontx 字体句柄.
*/
func XShapeText_SetFont(hTextBlock HXCGUI, hFontx HFONTX) {
	xShapeText_SetFont.Call(
		uintptr(hTextBlock),
		uintptr(hFontx))
}

/*
设置文本颜色.

参数:
	hTextBlock 形状对象文本句柄.
	color RGB颜色值.
	alpha 透明度
*/
func XShapeText_SetTextColor(hTextBlock HXCGUI, color COLORREF, alpha byte) {
	xShapeText_SetTextColor.Call(
		uintptr(hTextBlock),
		uintptr(color),
		uintptr(alpha))
}

/*
获取文本颜色.

参数:
	hTextBlock 形状对象文本句柄.
返回:
	颜色值.
*/
func XShapeText_GetTextColor(hTextBlock HXCGUI) COLORREF {
	ret, _, _ := xShapeText_GetTextColor.Call(
		uintptr(hTextBlock))

	return COLORREF(ret)
}

/*
设置文本对齐方式.

参数:
	hTextBlock 形状对象文本句柄.
	align 文本对齐方式.
*/
func XShapeText_SetTextAlign(hTextBlock HXCGUI, align int) {
	xShapeText_SetTextAlign.Call(
		uintptr(hTextBlock),
		uintptr(align))
}

/*
设置内容偏移.

参数:
	hTextBlock 形状对象文本句柄.
	x X坐标.
	y Y坐标.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XShapeText_SetOffset(hTextBlock HXCGUI, x, y int) {
	xShapeText_SetOffset.Call(
		uintptr(hTextBlock),
		uintptr(x),
		uintptr(y))
}

/*
设置布局宽度.

参数:
	hTextBlock 形状对象文本句柄.
	nType 指定属性类型
	width 布局宽度.
*/
func XShapeText_SetLayoutWidth(hTextBlock HXCGUI, nType Layout_size_type_, width int) {
	xShapeText_SetLayoutWidth.Call(
		uintptr(hTextBlock),
		uintptr(nType),
		uintptr(width))
}

/*
设置布局高度.

参数:
	hTextBlock 形状对象文本句柄.
	nType 指定属性类型
	height 布局高度.
*/
func XShapeText_SetLayoutHeight(hTextBlock HXCGUI, nType Layout_size_type_, height int) {
	xShapeText_SetLayoutHeight.Call(
		uintptr(hTextBlock),
		uintptr(nType),
		uintptr(height))
}

/*
获取布局宽度.

参数:
	hTextBlock 形状对象文本句柄.
	pType 属性类型
	pWidth 宽度
*/
func XShapeText_GetLayoutWidth(hTextBlock HXCGUI, pType *Layout_size_type_, pWidth *int) {
	xShapeText_GetLayoutWidth.Call(
		uintptr(hTextBlock),
		uintptr(unsafe.Pointer(pType)),
		uintptr(unsafe.Pointer(pWidth)))
}

/*
获取布局高度.

参数:
	hTextBlock 形状对象文本句柄.
	pType 属性类型
	pHeight 高度
*/
func XShapeText_GetLayoutHeight(hTextBlock HXCGUI, pType *Layout_size_type_, pHeight *int) {
	xShapeText_GetLayoutHeight.Call(
		uintptr(hTextBlock),
		uintptr(unsafe.Pointer(pType)),
		uintptr(unsafe.Pointer(pHeight)))
}
