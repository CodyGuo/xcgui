package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xLayout_Create          *syscall.Proc
	xLayout_Destroy         *syscall.Proc
	xLayout_AddEle          *syscall.Proc
	xLayout_AddLayoutObject *syscall.Proc
	xLayout_AddShape        *syscall.Proc
	xLayout_RemoveChild     *syscall.Proc
	xLayout_AdjustLayout    *syscall.Proc
	xLayout_GetEle          *syscall.Proc
	xLayout_GetParentLayout *syscall.Proc
	xLayout_GetParent       *syscall.Proc
	xLayout_GetWindow       *syscall.Proc
	xLayout_SetID           *syscall.Proc
	xLayout_GetID           *syscall.Proc
	xLayout_SetHorizon      *syscall.Proc
	xLayout_SetRectFrame    *syscall.Proc
	xLayout_GetRect         *syscall.Proc
	xLayout_GetRectFrame    *syscall.Proc
	xLayout_SetAlignH       *syscall.Proc
	xLayout_SetAlignV       *syscall.Proc
	xLayout_SetPadding      *syscall.Proc
	xLayout_SetLayoutWidth  *syscall.Proc
	xLayout_SetLayoutHeight *syscall.Proc
	xLayout_GetLayoutWidth  *syscall.Proc
	xLayout_GetLayoutHeight *syscall.Proc
	xLayout_SetSpace        *syscall.Proc
	xLayout_GetWidth        *syscall.Proc
	xLayout_GetHeight       *syscall.Proc
	xLayout_GetWidthIn      *syscall.Proc
	xLayout_GetHeightIn     *syscall.Proc
	xLayout_GetContentSize  *syscall.Proc
	xLayout_ShowLayout      *syscall.Proc
	xLayout_GetChildCount   *syscall.Proc
	xLayout_GetChildType    *syscall.Proc
	xLayout_GetChild        *syscall.Proc
	xLayout_Draw            *syscall.Proc
)

func init() {
	// Functions
	xLayout_Create = XCDLL.MustFindProc("XLayout_Create")
	xLayout_Destroy = XCDLL.MustFindProc("XLayout_Destroy")
	xLayout_AddEle = XCDLL.MustFindProc("XLayout_AddEle")
	xLayout_AddLayoutObject = XCDLL.MustFindProc("XLayout_AddLayoutObject")
	xLayout_AddShape = XCDLL.MustFindProc("XLayout_AddShape")
	xLayout_RemoveChild = XCDLL.MustFindProc("XLayout_RemoveChild")
	xLayout_AdjustLayout = XCDLL.MustFindProc("XLayout_AdjustLayout")
	xLayout_GetEle = XCDLL.MustFindProc("XLayout_GetEle")
	xLayout_GetParentLayout = XCDLL.MustFindProc("XLayout_GetParentLayout")
	xLayout_GetParent = XCDLL.MustFindProc("XLayout_GetParent")
	xLayout_GetWindow = XCDLL.MustFindProc("XLayout_GetWindow")
	xLayout_SetID = XCDLL.MustFindProc("XLayout_SetID")
	xLayout_GetID = XCDLL.MustFindProc("XLayout_GetID")
	xLayout_SetHorizon = XCDLL.MustFindProc("XLayout_SetHorizon")
	xLayout_SetRectFrame = XCDLL.MustFindProc("XLayout_SetRectFrame")
	xLayout_GetRect = XCDLL.MustFindProc("XLayout_GetRect")
	xLayout_GetRectFrame = XCDLL.MustFindProc("XLayout_GetRectFrame")
	xLayout_SetAlignH = XCDLL.MustFindProc("XLayout_SetAlignH")
	xLayout_SetAlignV = XCDLL.MustFindProc("XLayout_SetAlignV")
	xLayout_SetPadding = XCDLL.MustFindProc("XLayout_SetPadding")
	xLayout_SetLayoutWidth = XCDLL.MustFindProc("XLayout_SetLayoutWidth")
	xLayout_SetLayoutHeight = XCDLL.MustFindProc("XLayout_SetLayoutHeight")
	xLayout_GetLayoutWidth = XCDLL.MustFindProc("XLayout_GetLayoutWidth")
	xLayout_GetLayoutHeight = XCDLL.MustFindProc("XLayout_GetLayoutHeight")
	xLayout_SetSpace = XCDLL.MustFindProc("XLayout_SetSpace")
	xLayout_GetWidth = XCDLL.MustFindProc("XLayout_GetWidth")
	xLayout_GetHeight = XCDLL.MustFindProc("XLayout_GetHeight")
	xLayout_GetWidthIn = XCDLL.MustFindProc("XLayout_GetWidthIn")
	xLayout_GetHeightIn = XCDLL.MustFindProc("XLayout_GetHeightIn")
	xLayout_GetContentSize = XCDLL.MustFindProc("XLayout_GetContentSize")
	xLayout_ShowLayout = XCDLL.MustFindProc("XLayout_ShowLayout")
	xLayout_GetChildCount = XCDLL.MustFindProc("XLayout_GetChildCount")
	xLayout_GetChildType = XCDLL.MustFindProc("XLayout_GetChildType")
	xLayout_GetChild = XCDLL.MustFindProc("XLayout_GetChild")
	xLayout_Draw = XCDLL.MustFindProc("XLayout_Draw")

}

/*
创建布局对象.

返回:
	返回布局对象句柄.
*/
func XLayoutCreate() HXCGUI {
	ret, _, _ := xLayout_Create.Call()

	return HXCGUI(ret)
}

/*
销毁布局对象.

参数:
	hLayout 布局对象句柄.
*/
func XLayoutDestroy(hLayout HXCGUI) {
	xLayout_Destroy.Call(uintptr(hLayout))
}

/*
添加元素到布局对象,自动将元素添加到父UI中.

参数:
	hLayout 布局对象句柄.
	hEle 元素句柄.
*/
func XLayoutAddEle(hLayout HXCGUI, hEle HELE) {
	xLayout_AddEle.Call(
		uintptr(hLayout),
		uintptr(hEle))
}

/*
添加布局对象到当前布局对象.

参数:
	hLayout 布局对象句柄.
	hLayoutObject 布局对象句柄.
*/
func XLayoutAddLayoutObject(hLayout, hLayoutObject HXCGUI) {
	xLayout_AddLayoutObject.Call(
		uintptr(hLayout),
		uintptr(hLayoutObject))
}

/*
添加形状对象到当前布局对象,自动将形状对象添加到父UI中.

参数:
	hLayout 布局对象句柄.
	hShape 文本块对象.
*/
func XLayoutAddShape(hLayout HXCGUI, hShape HXCGUI) {
	xLayout_AddShape.Call(
		uintptr(hLayout),
		uintptr(hShape))
}

/*
移除子对象.

参数:
	hLayout 布局对象句柄.
	hChild 子对象句柄.
*/
func XLayoutRemoveChild(hLayout, hChild HXCGUI) {
	xLayout_RemoveChild.Call(
		uintptr(hLayout),
		uintptr(hChild))
}

/*
调整布局.

参数:
	hLayout 布局对象句柄.
*/
func XLayoutAdjustLayout(hLayout HXCGUI) {
	xLayout_AdjustLayout.Call(uintptr(hLayout))
}

/*
获取布局对象所在元素.

参数:
	hLayout 布局对象句柄.
返回:
	返回元素句柄.
*/
func XLayoutGetEle(hLayout HXCGUI) HELE {
	ret, _, _ := xLayout_GetEle.Call(uintptr(hLayout))

	return HELE(ret)
}

/*
获取父布局对象.

参数:
	hLayout 布局对象句柄.
返回:
	返回父布局对象,如果没有返回空.
*/
func XLayoutGetParentLayout(hLayout HXCGUI) HXCGUI {
	ret, _, _ := xLayout_GetParentLayout.Call(uintptr(hLayout))

	return HXCGUI(ret)
}

/*
获取父对象,父可能是元素或窗口.

参数:
	hLayout 布局对象句柄.
返回:
	返回父对象.
*/
func XLayoutGetParent(hLayout HXCGUI) HXCGUI {
	ret, _, _ := xLayout_GetParent.Call(uintptr(hLayout))

	return HXCGUI(ret)
}

/*
获取布局对象所在窗口.

参数:
	hLayout 布局对象句柄.
返回:
	返回窗口句柄.
*/
func XLayoutGetWindow(hLayout HXCGUI) HWINDOW {
	ret, _, _ := xLayout_GetWindow.Call(uintptr(hLayout))

	return HWINDOW(ret)
}

/*
设置ID.

参数:
	hLayout 布局对象句柄.
	nID ID值.
*/
func XLayoutSetID(hLayout HXCGUI, nID int) {
	xLayout_SetID.Call(
		uintptr(hLayout),
		uintptr(nID))
}

/*
获取ID.

参数:
	hLayout 布局对象句柄.
返回:
	返回ID值.
*/
func XLayoutGetID(hLayout HXCGUI) int {
	ret, _, _ := xLayout_GetID.Call(uintptr(hLayout))

	return int(ret)
}

/*
设置水平或垂直布局.

参数:
	hLayout 布局对象句柄.
	bHorizon 水平或垂直.
*/
func XLayoutSetHorizon(hLayout HXCGUI, bHorizon bool) {
	xLayout_SetHorizon.Call(
		uintptr(hLayout),
		uintptr(BoolToBOOL(bHorizon)))
}

/*
设置外围框架坐标.

参数:
	hLayout 布局对象句柄.
	pRect 坐标.
*/
func XLayoutSetRectFrame(hLayout HXCGUI, pRect *RECT) {
	xLayout_SetRectFrame.Call(
		uintptr(hLayout),
		uintptr(unsafe.Pointer(pRect)))
}

/*
获取内围坐标.

参数:
	hLayout 布局对象句柄.
	pRect 坐标.
*/
func XLayoutGetRect(hLayout HXCGUI, pRect *RECT) {
	xLayout_GetRect.Call(
		uintptr(hLayout),
		uintptr(unsafe.Pointer(pRect)))
}

/*
获取外围框架坐标.

参数:
	hLayout 布局对象句柄.
	pRect 接收返回坐标值.
*/
func XLayoutGetRectFrame(hLayout HXCGUI, pRect *RECT) {
	xLayout_GetRectFrame.Call(
		uintptr(hLayout),
		uintptr(unsafe.Pointer(pRect)))
}

/*
设置水平对齐方式.

参数:
	hLayout 布局对象句柄.
	nAlign 对齐方式参见宏定义.
*/
func XLayoutSetAlignH(hLayout HXCGUI, nAlign LAYOUT_ALIGN_) {
	xLayout_SetAlignH.Call(
		uintptr(hLayout),
		uintptr(nAlign))
}

/*
设置垂直对齐方式.

参数:
	hLayout 布局对象句柄.
	nAlign 对齐方式参见宏定义.
*/
func XLayoutSetAlignV(hLayout HXCGUI, nAlign LAYOUT_ALIGN_) {
	xLayout_SetAlignV.Call(
		uintptr(hLayout),
		uintptr(nAlign))
}

/*
设置内填充大小.

参数:
	hLayout 布局对象句柄.
	left 左边大小.
	top 上边大小.
	right 右边大小.
	bottom 下边大小.
*/
func XLayoutSetPadding(hLayout HXCGUI, left, top, right, bottom int) {
	xLayout_SetPadding.Call(
		uintptr(hLayout),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom))
}

/*
设置宽度.

参数:
	hLayout 布局对象句柄.
	nType 属性类型.
	nWidth 宽度.
*/
func XLayoutSetLayoutWidth(hLayout HXCGUI, nType LAYOUT_SIZE_TYPE_, nWidth int) {
	xLayout_SetLayoutWidth.Call(
		uintptr(hLayout),
		uintptr(nType),
		uintptr(nWidth))
}

/*
设置高度.

参数:
	hLayout 布局对象句柄.
	nType 属性类型.
	nHeight 高度.
*/
func XLayoutSetLayoutHeight(hLayout HXCGUI, nType LAYOUT_SIZE_TYPE_, nHeight int) {
	xLayout_SetLayoutHeight.Call(
		uintptr(hLayout),
		uintptr(nType),
		uintptr(nHeight))
}

/*
获取布局宽度.

参数:
	hLayout 布局对象句柄.
	pType 属性标识.
	pWidth 宽度.
*/
func XLayoutGetLayoutWidth(hLayout HXCGUI, pType *LAYOUT_SIZE_TYPE_, pWidth *int) {
	xLayout_GetLayoutWidth.Call(
		uintptr(hLayout),
		uintptr(unsafe.Pointer(pType)),
		uintptr(unsafe.Pointer(pWidth)))
}

/*
获取布局高度.

参数:
	hLayout 布局对象句柄.
	pType 属性标识.
	pHeight 高度.
*/
func XLayoutGetLayoutHeight(hLayout HXCGUI, pType *LAYOUT_SIZE_TYPE_, pHeight *int) {
	xLayout_GetLayoutHeight.Call(
		uintptr(hLayout),
		uintptr(unsafe.Pointer(pType)),
		uintptr(unsafe.Pointer(pHeight)))
}

/*
设置子对象之间的间距.

参数:
	hLayout 布局对象句柄.
	nSpace 间距大小.
*/
func XLayoutSetSpace(hLayout HXCGUI, nSpace int) {
	xLayout_SetSpace.Call(
		uintptr(hLayout),
		uintptr(nSpace))
}

/*
获取宽度.

参数:
	hLayout 布局对象句柄.
返回:
	返回宽度.
*/
func XLayoutGetWidth(hLayout HXCGUI) int {
	ret, _, _ := xLayout_GetWidth.Call(uintptr(hLayout))

	return int(ret)
}

/*
获取高度.

参数:
	hLayout 布局对象句柄.
返回:
	返回高度.
*/
func XLayoutGetHeight(hLayout HXCGUI) int {
	ret, _, _ := xLayout_GetHeight.Call(uintptr(hLayout))

	return int(ret)
}

/*
获取宽度,不包含内边距大小.

参数:
	hLayout 布局对象句柄.
返回:
	返回宽度.
*/
func XLayoutGetWidthIn(hLayout HXCGUI) int {
	ret, _, _ := xLayout_GetWidthIn.Call(uintptr(hLayout))

	return int(ret)
}

/*
获取高度,不包含内边距大小.

参数:
	hLayout 布局对象句柄.
返回:
	返回高度.
*/
func XLayoutGetHeightIn(hLayout HXCGUI) int {
	ret, _, _ := xLayout_GetHeightIn.Call(uintptr(hLayout))

	return int(ret)
}

/*
获取内容大小.

参数:
	hLayout 布局对象句柄.
	pSize 接收返回大小值.
*/
func XLayoutGetContentSize(hLayout HXCGUI, pSize *SIZE) {
	xLayout_GetContentSize.Call(
		uintptr(hLayout),
		uintptr(unsafe.Pointer(pSize)))
}

/*
是否显示布局对象.

参数:
	hLayout 布局对象句柄.
	bShow 是否显示.
*/
func XLayoutShowLayout(hLayout HXCGUI, bShow bool) {
	xLayout_ShowLayout.Call(
		uintptr(hLayout),
		uintptr(BoolToBOOL(bShow)))
}

/*
获取子对象数量.

参数:
	hLayout 布局对象句柄.
返回:
	返回子对象数量.
*/
func XLayoutGetChildCount(hLayout HXCGUI) int {
	ret, _, _ := xLayout_GetChildCount.Call(uintptr(hLayout))

	return int(ret)
}

/*
获取子对象类型.

参数:
	hLayout 布局对象句柄.
	index 索引值.
返回:
	返回类型.
*/
func XLayoutGetChildType(hLayout HXCGUI, index int) XC_OBJECT_TYPE {
	ret, _, _ := xLayout_GetChildType.Call(
		uintptr(hLayout),
		uintptr(index))

	return XC_OBJECT_TYPE(ret)
}

/*
获取子对象.

参数:
	hLayout 布局对象句柄.
	index 索引值.
返回:
	返回对象句柄.
*/
func XLayoutGetChild(hLayout HXCGUI, index int) HXCGUI {
	ret, _, _ := xLayout_GetChild.Call(
		uintptr(hLayout),
		uintptr(index))

	return HXCGUI(ret)
}

/*
绘制布局对象.

参数:
	hLayout 布局对象句柄.
	hDraw 图形绘制句柄.
*/
func XLayoutDraw(hLayout HXCGUI, hDraw HDRAW) {
	xLayout_Draw.Call(
		uintptr(hLayout),
		uintptr(hDraw))
}
