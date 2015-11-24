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
	xShapeEllipse_Create = XCDLL.MustFindProc("XShapeEllipse_Create")
	xShapeEllipse_SetBorderColor = XCDLL.MustFindProc("XShapeEllipse_SetBorderColor")
	xShapeEllipse_SetFillColor = XCDLL.MustFindProc("XShapeEllipse_SetFillColor")
	xShapeEllipse_EnableBorder = XCDLL.MustFindProc("XShapeEllipse_EnableBorder")
	xShapeEllipse_EnableFill = XCDLL.MustFindProc("XShapeEllipse_EnableFill")
}

func XShapeEllipseCreate(x int, y int, cx int, cy int, hParent HXCGUI) HXCGUI {
	ret, _, _ := xShapeEllipse_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))
	return HXCGUI(ret)
}

func XShapeEllipseSetBorderColor(hShape HXCGUI, color COLORREF, alpha byte) {
	xShapeEllipse_SetBorderColor.Call(
		uintptr(hShape),
		uintptr(color),
		uintptr(alpha))
}

func XShapeEllipseSetFillColor(hShape HXCGUI, color COLORREF, alpha byte) {
	xShapeEllipse_SetFillColor.Call(
		uintptr(hShape),
		uintptr(color),
		uintptr(alpha))
}

func XShapeEllipseEnableBorder(hShape HXCGUI, bEnable bool) {
	xShapeEllipse_EnableBorder.Call(
		uintptr(hShape),
		uintptr(BoolToBOOL(bEnable)))
}

func XShapeEllipseEnableFill(hShape HXCGUI, bEnable bool) {
	xShapeEllipse_EnableFill.Call(
		uintptr(hShape),
		uintptr(BoolToBOOL(bEnable)))
}
