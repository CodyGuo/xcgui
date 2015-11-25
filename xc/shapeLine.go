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
	xShapeLine_Create = XCDLL.MustFindProc("XShapeLine_Create")
	xShapeLine_SetPosition = XCDLL.MustFindProc("XShapeLine_SetPosition")
	xShapeLine_SetColor = XCDLL.MustFindProc("XShapeLine_SetColor")
}

func XShapeLineCreate(x1, int, y1 int, x2 int, y2 int, hParent HXCGUI) HXCGUI {
	ret, _, _ := xShapeLine_Create.Call(
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
		uintptr(hParent))
	return HXCGUI(ret)
}

func XShapeLineSetPosition(hShape HXCGUI, x1 int, y1 int, x2 int, y2 int) {
	xShapeLine_SetPosition.Call(
		uintptr(hShape),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2))
}

func XShapeLineSetColor(hShape HXCGUI, color COLORREF, alpha byte) {
	xShapeLine_SetColor.Call(
		uintptr(hShape),
		uintptr(color),
		uintptr(alpha))
}
