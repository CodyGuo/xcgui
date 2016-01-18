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

/*
创建形状对象-图片.

参数:
	x x坐标.
	y y坐标.
	cx 宽度.
	cy 高度.
	hParent 父对象句柄.
返回:
	成功返回图片对象句柄,否则返回NULL.
*/
func XShapePic_Create(x, y, cx, cy int, hParent HXCGUI) HXCGUI {
	ret, _, _ := xShapePic_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HXCGUI(ret)
}

/*
设置图片.

参数:
	hShape 形状对象句柄.
	hImage 图片句柄.
*/
func XShapePic_SetImage(hShape HXCGUI, hImage HIMAGE) {
	xShapePic_SetImage.Call(
		uintptr(hShape),
		uintptr(hImage))
}

/*
设置布局宽度.

参数:
	hShape 形状对象句柄.
	nType 标识.
	width 宽度.
*/
func XShapePic_SetLayoutWidth(hShape HXCGUI, nType Layout_size_type_, width int) {
	xShapePic_SetLayoutWidth.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(width))
}

/*
设置布局高度.

参数:
	hShape 形状对象句柄.
	nType 标识.
	height 高度.
*/
func XShapePic_SetLayoutHeight(hShape HXCGUI, nType Layout_size_type_, height int) {
	xShapePic_SetLayoutHeight.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(height))
}

/*
获取布局宽度.

参数:
	hShape 形状对象句柄.
	pType 属性标识.
	pWidth 宽度.
*/
func XShapePic_GetLayoutWidth(hShape HXCGUI, nType Layout_size_type_, pWidth *int) {
	xShapePic_GetLayoutWidth.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(unsafe.Pointer(pWidth)))
}

/*
获取布局高度.

参数:
	hShape 形状对象句柄.
	pType 属性标识.
	pHeight 高度.
*/
func XShapePic_GetLayoutHeight(hShape HXCGUI, nType Layout_size_type_, pHeight *int) {
	xShapePic_GetLayoutHeight.Call(
		uintptr(hShape),
		uintptr(nType),
		uintptr(unsafe.Pointer(pHeight)))
}
