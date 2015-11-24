package xc

import (
	"syscall"
	"unsafe"
)

var (
	xShapePic_Create          *syscall.Proc
	xShapePic_SetImage        *syscall.Proc
	xShapePic_SetLayoutWidth  *syscall.Proc
	xShapePic_SetLayoutHeight *syscall.Proc
	xShapePic_GetLayoutWidth  *syscall.Proc
	xShapePic_GetLayoutHeight *syscall.Proc
)

func init() {
	xShapePic_Create = XCDLL.MustFindProc("XShapePic_Create")
	xShapePic_SetImage = XCDLL.MustFindProc("XShapePic_SetImage")
	xShapePic_SetLayoutWidth = XCDLL.MustFindProc("XShapePic_SetLayoutWidth")
	xShapePic_SetLayoutHeight = XCDLL.MustFindProc("XShapePic_SetLayoutHeight")
	xShapePic_GetLayoutWidth = XCDLL.MustFindProc("XShapePic_GetLayoutWidth")
	xShapePic_GetLayoutHeight = XCDLL.MustFindProc("XShapePic_GetLayoutHeight")
}

func XShapePicCreate(x int, y int, cx int, cy int, hParent HXCGUI) HXCGUI {
	ret, _, _ := xShapePic_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))
	return HXCGUI(ret)
}

func XShapePicSetImage(hShape HXCGUI, hImage HIMAGE) {
	xShapePic_SetImage.Call(
		uintptr(hShape),
		uintptr(hImage))
}

func XShapePicSetLayoutWidth(hShape HXCGUI, nType layout_size_type_, width int) {
	xShapePic_SetLayoutWidth.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(width))
}

func XShapePicSetLayoutHeight(hShape HXCGUI, nType layout_size_type_, height int) {
	xShapePic_SetLayoutHeight.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(height))
}

func XShapePicGetLayoutWidth(hShape HXCGUI, nType layout_size_type_, pWidth *int) {
	xShapePic_GetLayoutWidth.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(unsafe.Pointer(pWidth)))
}

func XShapePicGetLayoutHeight(hShape HXCGUI, nType layout_size_type_, pHeight *int) {
	xShapePic_GetLayoutHeight.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(unsafe.Pointer(pHeight)))
}
