package xc

import (
	"syscall"
	"unsafe"
)

type PaddingSize_ int

var (
	xSView_Create                  *syscall.Proc
	xSView_SetTotalSize            *syscall.Proc
	xSView_GetTotalSize            *syscall.Proc
	xSView_SetLineSize             *syscall.Proc
	xSView_GetLineSize             *syscall.Proc
	xSView_GetViewPosH             *syscall.Proc
	xSView_GetViewPosV             *syscall.Proc
	xSView_GetViewWidth            *syscall.Proc
	xSView_GetViewHeight           *syscall.Proc
	xSView_GetViewRect             *syscall.Proc
	xSView_GetScrollBarH           *syscall.Proc
	xSView_GetScrollBarV           *syscall.Proc
	xSView_SetPadding              *syscall.Proc
	xSView_GetPadding              *syscall.Proc
	xSView_ScrollPosH              *syscall.Proc
	xSView_ScrollPosV              *syscall.Proc
	xSView_ScrollPosXH             *syscall.Proc
	xSView_ScrollPosYV             *syscall.Proc
	xSView_ShowSBarH               *syscall.Proc
	xSView_ShowSBarV               *syscall.Proc
	xSView_EnableAutoShowScrollBar *syscall.Proc
	xSView_ScrollLeftLine          *syscall.Proc
	xSView_ScrollRightLine         *syscall.Proc
	xSView_ScrollTopLine           *syscall.Proc
	xSView_ScrollBottomLine        *syscall.Proc
	xSView_ScrollLeft              *syscall.Proc
	xSView_ScrollRight             *syscall.Proc
	xSView_ScrollTop               *syscall.Proc
	xSView_ScrollBottom            *syscall.Proc
)

func init() {
	xSView_Create = xcDLL.MustFindProc("XSView_Create")
	xSView_SetTotalSize = xcDLL.MustFindProc("XSView_SetTotalSize")
	xSView_GetTotalSize = xcDLL.MustFindProc("XSView_GetTotalSize")
	xSView_SetLineSize = xcDLL.MustFindProc("XSView_SetLineSize")
	xSView_GetLineSize = xcDLL.MustFindProc("XSView_GetLineSize")
	xSView_GetViewPosH = xcDLL.MustFindProc("XSView_GetViewPosH")
	xSView_GetViewPosV = xcDLL.MustFindProc("XSView_GetViewPosV")
	xSView_GetViewWidth = xcDLL.MustFindProc("XSView_GetViewWidth")
	xSView_GetViewHeight = xcDLL.MustFindProc("XSView_GetViewHeight")
	xSView_GetViewRect = xcDLL.MustFindProc("XSView_GetViewRect")
	xSView_GetScrollBarH = xcDLL.MustFindProc("XSView_GetScrollBarH")
	xSView_GetScrollBarV = xcDLL.MustFindProc("XSView_GetScrollBarV")
	xSView_SetPadding = xcDLL.MustFindProc("XSView_SetPadding")
	xSView_GetPadding = xcDLL.MustFindProc("XSView_GetPadding")
	xSView_ScrollPosH = xcDLL.MustFindProc("XSView_ScrollPosH")
	xSView_ScrollPosV = xcDLL.MustFindProc("XSView_ScrollPosV")
	xSView_ScrollPosXH = xcDLL.MustFindProc("XSView_ScrollPosXH")
	xSView_ScrollPosYV = xcDLL.MustFindProc("XSView_ScrollPosYV")
	xSView_ShowSBarH = xcDLL.MustFindProc("XSView_ShowSBarH")
	xSView_ShowSBarV = xcDLL.MustFindProc("XSView_ShowSBarV")
	xSView_EnableAutoShowScrollBar = xcDLL.MustFindProc("XSView_EnableAutoShowScrollBar")
	xSView_ScrollLeftLine = xcDLL.MustFindProc("XSView_ScrollLeftLine")
	xSView_ScrollRightLine = xcDLL.MustFindProc("XSView_ScrollRightLine")
	xSView_ScrollTopLine = xcDLL.MustFindProc("XSView_ScrollTopLine")
	xSView_ScrollBottomLine = xcDLL.MustFindProc("XSView_ScrollBottomLine")
	xSView_ScrollLeft = xcDLL.MustFindProc("XSView_ScrollLeft")
	xSView_ScrollRight = xcDLL.MustFindProc("XSView_ScrollRight")
	xSView_ScrollTop = xcDLL.MustFindProc("XSView_ScrollTop")
	xSView_ScrollBottom = xcDLL.MustFindProc("XSView_ScrollBottom")
}

/*
创建滚动视图元素.

参数:
	x 元素x坐标.
	y 元素y坐标.
	cx 宽度.
	cy 高度.
	hParent 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
返回:
	元素句柄.
*/
func XSView_Create(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xSView_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
设置内容大小.

参数:
	hEle 元素句柄.
	cx 宽度.
	cy 高度.
返回:
	如果内容改变返回TRUE否则返回FALSE.
*/
func XSView_SetTotalSize(hEle HELE, cx, cy int) bool {
	ret, _, _ := xSView_SetTotalSize.Call(
		uintptr(hEle),
		uintptr(cx),
		uintptr(cy))

	return ret == TRUE

}

/*
获取内容总大小.

参数:
	hEle 元素句柄.
	pSize 大小.
*/
func XSView_GetTotalSize(hEle HELE, pSize *SIZE) {
	xSView_GetTotalSize.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pSize)))
}

/*
设置滚动单位大小.

参数:
	hEle 元素句柄.
	nWidth 宽度.
	nHeight 高度.
返回:
	如果内容改变返回TRUE否则返回FALSE.
*/
func XSView_SetLineSize(hEle HELE, nWidth, nHeight int) bool {
	ret, _, _ := xSView_SetLineSize.Call(
		uintptr(hEle),
		uintptr(nWidth),
		uintptr(nHeight))

	return ret == TRUE
}

/*
获取滚动单位大小.

参数:
	hEle 元素句柄.
	pSize 返回大小.
*/
func XSView_GetLineSize(hEle HELE, pSize *SIZE) {
	xSView_GetLineSize.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pSize)))
}

/*
获取视口原点X坐标.

参数:
	hEle 元素句柄.
返回:
	视口原点X坐标.
*/
func XSView_GetViewPosH(hEle HELE) int {
	ret, _, _ := xSView_GetViewPosH.Call(
		uintptr(hEle))

	return int(ret)
}

/*
获取视口原点Y坐标.

参数:
	hEle 元素句柄.
返回:
	视口原点Y坐标.
*/
func XSView_GetViewPosV(hEle HELE) int {
	ret, _, _ := xSView_GetViewPosV.Call(
		uintptr(hEle))

	return int(ret)
}

/*
获取视口宽度.

参数:
	hEle 元素句柄.
返回:
	返回视口宽度.
*/
func XSView_GetViewWidth(hEle HELE) int {
	ret, _, _ := xSView_GetViewWidth.Call(
		uintptr(hEle))

	return int(ret)
}

/*
获取视口高度.

参数:
	hEle 元素句柄.
返回:
	返回视口高度.
*/
func XSView_GetViewHeight(hEle HELE) int {
	ret, _, _ := xSView_GetViewHeight.Call(
		uintptr(hEle))

	return int(ret)
}

/*
获取视口坐标.

参数:
	hEle 元素句柄.
	pRect 坐标.
*/
func XSView_GetViewRect(hEle HELE, pRect *RECT) {
	xSView_GetLineSize.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

/*
获取水平滚动条.

参数:
	hEle 元素句柄.
返回:
	滚动条句柄.
*/
func XSView_GetScrollBarH(hEle HELE) HELE {
	ret, _, _ := xSView_GetScrollBarH.Call(
		uintptr(hEle))

	return HELE(ret)
}

/*
获取垂直滚动条.

参数:
	hEle 元素句柄.
返回:
	垂直滚动条句柄.
*/
func XSView_GetScrollBarV(hEle HELE) HELE {
	ret, _, _ := xSView_GetScrollBarV.Call(
		uintptr(hEle))

	return HELE(ret)
}

/*
设置边框内填充大小.

参数:
	hEle 元素句柄.
	left 左边大小.
	top 上边大小.
	right 右边大小.
	bottom 下边大小.
*/
func XSView_SetPadding(hEle HELE, left, top, right, bottom int) {
	xSView_SetPadding.Call(
		uintptr(hEle),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom))
}

/*
获取边框内填充大小.

参数:
	hEle 元素句柄.
	pPadding 大小.
*/
func XSView_GetPadding(hEle HELE, pPadding *PaddingSize_) {
	xSView_GetPadding.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPadding)))
}

/*
水平滚动条,滚动到指定位置点.

参数:
	hEle 元素句柄.
	pos 位置点.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XSView_ScrollPosH(hEle HELE, pos int) bool {
	ret, _, _ := xSView_ScrollPosH.Call(
		uintptr(hEle),
		uintptr(pos))

	return ret == TRUE
}

/*
垂直滚动条,滚动到指定位置点.

参数:
	hEle 元素句柄.
	pos 位置点.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XSView_ScrollPosV(hEle HELE, pos int) bool {
	ret, _, _ := xSView_ScrollPosV.Call(
		uintptr(hEle),
		uintptr(pos))

	return ret == TRUE
}

/*
水平滚动条,滚动到指定坐标.

参数:
	hEle 元素句柄.
	posX X坐标.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XSView_ScrollPosXH(hEle HELE, posX int) bool {
	ret, _, _ := xSView_ScrollPosXH.Call(
		uintptr(hEle),
		uintptr(posX))

	return ret == TRUE
}

/*
垂直滚动条,滚动到指定坐标.

参数:
	hEle 元素句柄.
	posY Y坐标.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XSView_ScrollPosYV(hEle HELE, posY int) bool {
	ret, _, _ := xSView_ScrollPosYV.Call(
		uintptr(hEle),
		uintptr(posY))

	return ret == TRUE
}

/*
显示水平滚动条.

参数:
	hEle 元素句柄.
	bShow 是否显示.
*/
func XSView_ShowSBarH(hEle HELE, bShow bool) {
	xSView_ShowSBarH.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bShow)))
}

/*
显示垂直滚动条.

参数:
	hEle 元素句柄.
	bShow 是否显示.
*/
func XSView_ShowSBarV(hEle HELE, bShow bool) {
	xSView_ShowSBarV.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bShow)))
}

/*
启用自动显示滚动条.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XSView_EnableAutoShowScrollBar(hEle HELE, bEnable bool) {
	xSView_EnableAutoShowScrollBar.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
向左滚动.

参数:
	hEle 元素句柄.
返回:
	如果成功返回TRUE,否则相反.
*/
func XSView_ScrollLeftLine(hEle HELE) bool {
	ret, _, _ := xSView_ScrollLeftLine.Call(
		uintptr(hEle))

	return ret == TRUE
}

/*
向右滚动.

参数:
	hEle 元素句柄.
返回:
	如果成功返回TRUE,否则相反.
*/
func XSView_ScrollRightLine(hEle HELE) bool {
	ret, _, _ := xSView_ScrollRightLine.Call(
		uintptr(hEle))

	return ret == TRUE
}

/*
向上滚动.

参数:
	hEle 元素句柄.
返回:
	如果成功返回TRUE,否则相反.
*/
func XSView_ScrollTopLine(hEle HELE) bool {
	ret, _, _ := xSView_ScrollTopLine.Call(
		uintptr(hEle))

	return ret == TRUE
}

/*
向下滚动.

参数:
	hEle 元素句柄.
返回:
	如果成功返回TRUE,否则相反.
*/
func XSView_ScrollBottomLine(hEle HELE) bool {
	ret, _, _ := xSView_ScrollBottomLine.Call(
		uintptr(hEle))

	return ret == TRUE
}

/*
水平滚动到左侧.

参数:
	hEle 元素句柄.
返回:
	如果成功返回TRUE,否则相反.
*/
func XSView_ScrollLeft(hEle HELE) bool {
	ret, _, _ := xSView_ScrollLeft.Call(
		uintptr(hEle))

	return ret == TRUE
}

/*
水平滚动到右侧.

参数:
	hEle 元素句柄.
返回:
	如果成功返回TRUE,否则相反.
*/
func XSView_ScrollRight(hEle HELE) bool {
	ret, _, _ := xSView_ScrollRight.Call(
		uintptr(hEle))

	return ret == TRUE
}

/*
垂直滚动到顶部.

参数:
	hEle 元素句柄.
返回:
	如果成功返回TRUE,否则相反.
*/
func XSView_ScrollTop(hEle HELE) bool {
	ret, _, _ := xSView_ScrollTop.Call(
		uintptr(hEle))

	return ret == TRUE
}

/*
垂直滚动到底部.

参数:
	hEle 元素句柄.
返回:
	如果成功返回TRUE,否则相反.
*/
func XSView_ScrollBottom(hEle HELE) bool {
	ret, _, _ := xSView_ScrollBottom.Call(
		uintptr(hEle))

	return ret == TRUE
}
