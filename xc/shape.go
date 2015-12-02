package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Funcations
	xShape_GetParentEle    *syscall.Proc
	xShape_GetParentLayout *syscall.Proc
	xShape_GetHWINDOW      *syscall.Proc
	xShape_GetParent       *syscall.Proc
	xShape_SetID           *syscall.Proc
	xShape_GetID           *syscall.Proc
	xShape_Redraw          *syscall.Proc
	xShape_GetWidth        *syscall.Proc
	xShape_GetHeight       *syscall.Proc
	xShape_GetRect         *syscall.Proc
	xShape_SetRect         *syscall.Proc
	xShape_GetContentSize  *syscall.Proc
	xShape_ShowLayout      *syscall.Proc
	xShape_AdjustLayout    *syscall.Proc
	xShape_Destroy         *syscall.Proc
)

func init() {
	// Functions
	xShape_GetParentEle = xcDLL.MustFindProc("XShape_GetParentEle")
	xShape_GetParentLayout = xcDLL.MustFindProc("XShape_GetParentLayout")
	xShape_GetHWINDOW = xcDLL.MustFindProc("XShape_GetHWINDOW")
	xShape_GetParent = xcDLL.MustFindProc("XShape_GetParent")
	xShape_SetID = xcDLL.MustFindProc("XShape_SetID")
	xShape_GetID = xcDLL.MustFindProc("XShape_GetID")
	xShape_Redraw = xcDLL.MustFindProc("XShape_Redraw")
	xShape_GetWidth = xcDLL.MustFindProc("XShape_GetWidth")
	xShape_GetHeight = xcDLL.MustFindProc("XShape_GetHeight")
	xShape_GetRect = xcDLL.MustFindProc("XShape_GetRect")
	xShape_SetRect = xcDLL.MustFindProc("XShape_SetRect")
	xShape_GetContentSize = xcDLL.MustFindProc("XShape_GetContentSize")
	xShape_ShowLayout = xcDLL.MustFindProc("XShape_ShowLayout")
	xShape_AdjustLayout = xcDLL.MustFindProc("XShape_AdjustLayout")
	xShape_Destroy = xcDLL.MustFindProc("XShape_Destroy")
}

/*
获取所在元素句柄.

参数:
	hShape 形状对象句柄.
返回:
	返回元素句柄.
*/
func XShapeGetParentEle(hShape HXCGUI) HELE {
	ret, _, _ := xShape_GetParentEle.Call(uintptr(hShape))

	return HELE(ret)
}

/*
获取所在布局对象.

参数:
	hShape 形状对象句柄.
返回:
	布局对象句柄.
*/
func XShapeGetParentLayout(hShape HXCGUI) HXCGUI {
	ret, _, _ := xShape_GetParentLayout.Call(uintptr(hShape))

	return HXCGUI(ret)
}

/*
获取窗口句柄.

参数:
	hShape 形状对象句柄.
返回:
	窗口句柄.
*/
func XShapeGetHWINDOW(hShape HXCGUI) HWINDOW {
	ret, _, _ := xShape_GetHWINDOW.Call(uintptr(hShape))

	return HWINDOW(ret)
}

/*
获取父对象,父可能是元素或窗口.

参数:
	hShape 形状对象句柄.
返回:
	返回父对象.
*/
func XShapeGetParent(hShape HXCGUI) HXCGUI {
	ret, _, _ := xShape_GetParent.Call(uintptr(hShape))

	return HXCGUI(ret)
}

/*
设置ID.

参数:
	hShape 形状对象句柄.
	nID ID值.
*/
func XShapeSetID(hShape HXCGUI, nID int) {
	xShape_SetID.Call(
		uintptr(hShape),
		uintptr(nID))
}

/*
获取ID.

参数:
	hShape 形状对象句柄.
返回:
	返回ID.
*/
func XShapeGetID(hShape HXCGUI) int {
	ret, _, _ := xShape_GetID.Call(uintptr(hShape))

	return int(ret)
}

/*
重绘形状对象.

参数:
	hShape 形状对象句柄.
*/
func XShapeRedraw(hShape HXCGUI) {
	xShape_Redraw.Call(uintptr(hShape))
}

/*
获取内容宽度.

参数:
	hShape 形状对象句柄.
返回:
	返回内容宽度.
*/
func XShapeGetWidth(hShape HXCGUI) int {
	ret, _, _ := xShape_GetWidth.Call(uintptr(hShape))

	return int(ret)
}

/*
获取内容高度.

参数:
	hShape 形状对象句柄.
返回:
	返回内容高度.
*/
func XShapeGetHeight(hShape HXCGUI) int {
	ret, _, _ := xShape_GetHeight.Call(uintptr(hShape))

	return int(ret)
}

/*
获取坐标.

参数:
	hShape 形状对象句柄.
	pRect 接收返回坐标.
*/
func XShapeGetRect(hShape HXCGUI, pRect *RECT) {
	xShape_GetRect.Call(
		uintptr(hShape),
		uintptr(unsafe.Pointer(pRect)))
}

/*
设置坐标.

参数:
	hShape 形状对象句柄.
	pRect 坐标.
*/
func XShapeSetRect(hShape HXCGUI, pRect *RECT) {
	xShape_SetRect.Call(
		uintptr(hShape),
		uintptr(unsafe.Pointer(pRect)))
}

/*
获取内容大小.

参数:
	hShape 形状对象句柄.
	pSize 接收返回内容大小值.
*/
func XShapeGetContentSize(hShape HXCGUI, pSize *SIZE) {
	xShape_GetContentSize.Call(
		uintptr(hShape),
		uintptr(unsafe.Pointer(pSize)))
}

/*
是否显示布局边界.

参数:
	hShape 形状对象句柄.
	bShow 是否显示.
*/
func XShapeShowLayout(hShape HXCGUI, bShow bool) {
	xShape_ShowLayout.Call(
		uintptr(hShape),
		uintptr(BoolToBOOL(bShow)))
}

/*
调整布局.

参数:
	hShape 形状对象句柄.
*/
func XShapeAdjustLayout(hShape HXCGUI) {
	xShape_AdjustLayout.Call(uintptr(hShape))
}

/*
销毁形状对象.

参数:
	hShape 形状对象句柄.
*/
func XShapeDestroy(hShape HXCGUI) {
	xShape_Destroy.Call(uintptr(hShape))
}
