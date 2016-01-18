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

/*
创建形状对象GIF.

参数:
	x X坐标.
	y Y坐标.
	cx 宽度.
	cy 高度.
	hParent 父对象句柄.
返回:
	成功返回形状对象GIF句柄,否则返回NULL.
*/
func XShapeGif_Create(x, y, cx, cy int, hParent HXCGUI) HXCGUI {
	ret, _, _ := xShapeGif_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HXCGUI(ret)
}

/*
设置GIF图片.

参数:
	hShape 形状对象句柄.
	hImage 图片句柄.
*/
func XShapeGif_SetImage(hShape HXCGUI, hImage HIMAGE) {
	xShapeGif_SetImage.Call(
		uintptr(hShape),
		uintptr(hImage))
}

/*
设置布局宽度.

参数:
	hShape 形状对象句柄.
	nType 属性类型.
	width 宽度.
*/
func XShapeGif_SetLayoutWidth(hShape HXCGUI, nType Layout_size_type_, width int) {
	xShapeGif_SetLayoutWidth.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(width))
}

/*
设置布局高度.

参数:
	hShape 形状对象句柄.
	nType 属性类型.
	height 高度.
*/
func XShapeGif_SetLayoutHeight(hShape HXCGUI, nType Layout_size_type_, height int) {
	xShapeGif_SetLayoutHeight.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(height))
}

/*
获取布局宽度.

参数:
	hShape 形状对象句柄.
	pType 属性类型.
	pWidth 宽度.
*/
func XShapeGif_GetLayoutWidth(hShape HXCGUI, nType Layout_size_type_, pWidth *int) {
	xShapeGif_GetLayoutWidth.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(unsafe.Pointer(pWidth)))
}

/*
获取布局高度.

参数:
	hShape 形状对象句柄.
	pType 属性类型.
	pHeight 高度.
*/
func XShapeGif_GetLayoutHeight(hShape HXCGUI, nType Layout_size_type_, pHeight *int) {
	xShapeGif_GetLayoutHeight.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(unsafe.Pointer(pHeight)))
}
