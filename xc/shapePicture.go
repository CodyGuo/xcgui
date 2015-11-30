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
	xShapePic_Create = xcDLL.MustFindProc("XShapePic_Create")
	xShapePic_SetImage = xcDLL.MustFindProc("XShapePic_SetImage")
	xShapePic_SetLayoutWidth = xcDLL.MustFindProc("XShapePic_SetLayoutWidth")
	xShapePic_SetLayoutHeight = xcDLL.MustFindProc("XShapePic_SetLayoutHeight")
	xShapePic_GetLayoutWidth = xcDLL.MustFindProc("XShapePic_GetLayoutWidth")
	xShapePic_GetLayoutHeight = xcDLL.MustFindProc("XShapePic_GetLayoutHeight")
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

func XShapePicSetLayoutWidth(hShape HXCGUI, nType LAYOUT_SIZE_TYPE_, width int) {
	xShapePic_SetLayoutWidth.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(width))
}

func XShapePicSetLayoutHeight(hShape HXCGUI, nType LAYOUT_SIZE_TYPE_, height int) {
	xShapePic_SetLayoutHeight.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(height))
}

func XShapePicGetLayoutWidth(hShape HXCGUI, nType LAYOUT_SIZE_TYPE_, pWidth *int) {
	xShapePic_GetLayoutWidth.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(unsafe.Pointer(pWidth)))
}

func XShapePicGetLayoutHeight(hShape HXCGUI, nType LAYOUT_SIZE_TYPE_, pHeight *int) {
	xShapePic_GetLayoutHeight.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(unsafe.Pointer(pHeight)))
}
