package xc

import (
	"syscall"
	"unsafe"
)

var (
	xShapeGif_Create          *syscall.Proc
	xShapeGif_SetImage        *syscall.Proc
	xShapeGif_SetLayoutWidth  *syscall.Proc
	xShapeGif_SetLayoutHeight *syscall.Proc
	xShapeGif_GetLayoutWidth  *syscall.Proc
	xShapeGif_GetLayoutHeight *syscall.Proc
)

func init() {
	xShapeGif_Create = xcDLL.MustFindProc("XShapeGif_Create")
	xShapeGif_SetImage = xcDLL.MustFindProc("XShapeGif_SetImage")
	xShapeGif_SetLayoutWidth = xcDLL.MustFindProc("XShapeGif_SetLayoutWidth")
	xShapeGif_SetLayoutHeight = xcDLL.MustFindProc("XShapeGif_SetLayoutHeight")
	xShapeGif_GetLayoutWidth = xcDLL.MustFindProc("XShapeGif_GetLayoutWidth")
	xShapeGif_GetLayoutHeight = xcDLL.MustFindProc("XShapeGif_GetLayoutHeight")
}

func xShapeGifCreate(x int, y int, cx int, cy int, hParent HXCGUI) HXCGUI {
	ret, _, _ := xShapeGif_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))
	return HXCGUI(ret)
}

func xShapeGifSetImage(hShape HXCGUI, hImage HIMAGE) {
	xShapeGif_SetImage.Call(
		uintptr(hShape),
		uintptr(hImage))
}

func xShapeGifSetLayoutWidth(hShape HXCGUI, nType LAYOUT_SIZE_TYPE_, width int) {
	xShapeGif_SetLayoutWidth.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(width))
}

func xShapeGifSetLayoutHeight(hShape HXCGUI, nType LAYOUT_SIZE_TYPE_, height int) {
	xShapeGif_SetLayoutHeight.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(height))
}

func xShapeGifGetLayoutWidth(hShape HXCGUI, nType LAYOUT_SIZE_TYPE_, pWidth *int) {
	xShapeGif_GetLayoutWidth.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(unsafe.Pointer(pWidth)))
}

func xShapeGifGetLayoutHeight(hShape HXCGUI, nType LAYOUT_SIZE_TYPE_, pHeight *int) {
	xShapeGif_GetLayoutHeight.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(unsafe.Pointer(pHeight)))
}
