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
	xShapeGif_Create = XCDLL.MustFindProc("XShapeGif_Create")
	xShapeGif_SetImage = XCDLL.MustFindProc("XShapeGif_SetImage")
	xShapeGif_SetLayoutWidth = XCDLL.MustFindProc("XShapeGif_SetLayoutWidth")
	xShapeGif_SetLayoutHeight = XCDLL.MustFindProc("XShapeGif_SetLayoutHeight")
	xShapeGif_GetLayoutWidth = XCDLL.MustFindProc("XShapeGif_GetLayoutWidth")
	xShapeGif_GetLayoutHeight = XCDLL.MustFindProc("XShapeGif_GetLayoutHeight")
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

func xShapeGifSetLayoutWidth(hShape HXCGUI, nType layout_size_type_, width int) {
	xShapeGif_SetLayoutWidth.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(width))
}

func xShapeGifSetLayoutHeight(hShape HXCGUI, nType layout_size_type_, height int) {
	xShapeGif_SetLayoutHeight.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(height))
}

func xShapeGifGetLayoutWidth(hShape HXCGUI, nType layout_size_type_, pWidth *int) {
	xShapeGif_GetLayoutWidth.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(unsafe.Pointer(pWidth)))
}

func xShapeGifGetLayoutHeight(hShape HXCGUI, nType layout_size_type_, pHeight *int) {
	xShapeGif_GetLayoutHeight.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(unsafe.Pointer(pHeight)))
}
