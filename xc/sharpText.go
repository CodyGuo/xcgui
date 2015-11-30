package xc

import (
	"syscall"
	"unsafe"
)

type out_int int

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

func XShapeTextCreate(x int, y int, cx int, cy int, pName string, hParent HXCGUI) HXCGUI {
	ret, _, _ := xShapeText_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		StringToUintPtr(pName),
		uintptr(hParent))

	return HXCGUI(ret)
}

func XShapeTextSetText(hTextBlock HXCGUI, pName string) {
	xShapeText_SetText.Call(
		uintptr(hTextBlock),
		StringToUintPtr(pName))
}

/*
buf_szize := 255 + 1
buf := make([]uint16, buf_szize)
xcgui.ShapeText_GetText(lb, &buf[0], buf_szize)
str := syscall.UTF16ToString(buf)
*/
func XShapeTextGetText(hTextBlock HXCGUI, pOut *uint16, nOutLen int) {
	xShapeText_GetText.Call(
		uintptr(hTextBlock),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))
}

//非官方api,功能同上
//更方便golang
func XShapeTextGetTextGo(hTextBlock HXCGUI) string {
	buf_szize := 256
	buf := make([]uint16, buf_szize)
	XShapeTextGetText(hTextBlock, &buf[0], buf_szize)
	return syscall.UTF16ToString(buf)
}

func XShapeTextGetTextLength(hTextBlock HXCGUI) int {
	ret, _, _ := xShapeText_GetTextLength.Call(
		uintptr(hTextBlock))
	return int(ret)
}

func XShapeTextSetFont(hTextBlock HXCGUI, hFontx HFONTX) {
	xShapeText_SetFont.Call(
		uintptr(hTextBlock),
		uintptr(hFontx))
}

func XShapeTextSetTextColor(hTextBlock HXCGUI, color COLORREF, alpha byte) {
	xShapeText_SetTextColor.Call(
		uintptr(hTextBlock),
		uintptr(color),
		uintptr(alpha))
}

func XShapeTextGetTextColor(hTextBlock HXCGUI) COLORREF {
	ret, _, _ := xShapeText_GetTextColor.Call(
		uintptr(hTextBlock))
	return COLORREF(ret)
}

func XShapeTextSetTextAlign(hTextBlock HXCGUI, align int) {
	xShapeText_SetTextAlign.Call(
		uintptr(hTextBlock),
		uintptr(align))
}

func XShapeTextSetOffset(hTextBlock HXCGUI, x int, y int) {
	xShapeText_SetOffset.Call(
		uintptr(hTextBlock),
		uintptr(x),
		uintptr(y))
}

func XShapeTextSetLayoutWidth(hTextBlock HXCGUI, nType LAYOUT_SIZE_TYPE_, width int) {
	xShapeText_SetLayoutWidth.Call(
		uintptr(hTextBlock),
		uintptr(nType),
		uintptr(width))
}

func XShapeTextSetLayoutHeight(hTextBlock HXCGUI, nType LAYOUT_SIZE_TYPE_, height int) {
	xShapeText_SetLayoutHeight.Call(
		uintptr(hTextBlock),
		uintptr(nType),
		uintptr(height))
}

func XShapeTextGetLayoutWidth(hTextBlock HXCGUI, pType *LAYOUT_SIZE_TYPE_, pWidth *out_int) {
	xShapeText_GetLayoutWidth.Call(
		uintptr(hTextBlock),
		uintptr(unsafe.Pointer(pType)),
		uintptr(unsafe.Pointer(pWidth)))
}

func XShapeTextGetLayoutHeight(hTextBlock HXCGUI, pType *LAYOUT_SIZE_TYPE_, pHeight *out_int) {
	xShapeText_GetLayoutHeight.Call(
		uintptr(hTextBlock),
		uintptr(unsafe.Pointer(pType)),
		uintptr(unsafe.Pointer(pHeight)))
}
