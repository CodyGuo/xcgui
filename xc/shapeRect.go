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

func XShapeRectCreate(x1, int, y1 int, x2 int, y2 int, hParent HXCGUI) HXCGUI {
	ret, _, _ := xShapeRect_Create.Call(
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
		uintptr(hParent))
	return HXCGUI(ret)
}

func XShapeRectSetBorderColor(hShape HXCGUI, color COLORREF, alpha byte) {
	xShapeRect_SetBorderColor.Call(
		uintptr(hShape),
		uintptr(color),
		uintptr(alpha))
}

func XShapeRectSetFillColor(hShape HXCGUI, color COLORREF, alpha byte) {
	xShapeRect_SetFillColor.Call(
		uintptr(hShape),
		uintptr(color),
		uintptr(alpha))
}

func XShapeRectSetRoundAngle(hShape HXCGUI, nWidth int, nHeight int) {
	xShapeRect_SetRoundAngle.Call(
		uintptr(hShape),
		uintptr(nWidth),
		uintptr(nHeight))
}

func XShapeRectGetRoundAngle(hShape HXCGUI, pWidth *int, pHeight *int) {
	xShapeRect_GetRoundAngle.Call(
		uintptr(hShape),
		uintptr(unsafe.Pointer(pWidth)),
		uintptr(unsafe.Pointer(pHeight)))
}

func XShapeRectEnableBorder(hShape HXCGUI, bEnable bool) {
	xShapeRect_EnableBorder.Call(
		uintptr(hShape),
		uintptr(BoolToBOOL(bEnable)))
}

func XShapeRectEnableFill(hShape HXCGUI, bEnable bool) {
	xShapeRect_EnableFill.Call(
		uintptr(hShape),
		uintptr(BoolToBOOL(bEnable)))
}

func XShapeRectEnableRoundAngle(hShape HXCGUI, bEnable bool) {
	xShapeRect_EnableRoundAngle.Call(
		uintptr(hShape),
		uintptr(BoolToBOOL(bEnable)))
}
