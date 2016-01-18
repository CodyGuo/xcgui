package xc

import (
	"syscall"
)

var (
	xSliderBar_Create          *syscall.Proc
	xSliderBar_SetRange        *syscall.Proc
	xSliderBar_GetRange        *syscall.Proc
	xSliderBar_SetImageLoad    *syscall.Proc
	xSliderBar_SetButtonWidth  *syscall.Proc
	xSliderBar_SetButtonHeight *syscall.Proc
	xSliderBar_SetSpaceTwo     *syscall.Proc
	xSliderBar_SetPos          *syscall.Proc
	xSliderBar_GetPos          *syscall.Proc
	xSliderBar_GetButton       *syscall.Proc
	xSliderBar_SetHorizon      *syscall.Proc
)

func init() {
	xSliderBar_Create = xcDLL.MustFindProc("XSliderBar_Create")
	xSliderBar_SetRange = xcDLL.MustFindProc("XSliderBar_SetRange")
	xSliderBar_GetRange = xcDLL.MustFindProc("XSliderBar_GetRange")
	xSliderBar_SetImageLoad = xcDLL.MustFindProc("XSliderBar_SetImageLoad")
	xSliderBar_SetButtonWidth = xcDLL.MustFindProc("XSliderBar_SetButtonWidth")
	xSliderBar_SetButtonHeight = xcDLL.MustFindProc("XSliderBar_SetButtonHeight")
	xSliderBar_SetSpaceTwo = xcDLL.MustFindProc("XSliderBar_SetSpaceTwo")
	xSliderBar_SetPos = xcDLL.MustFindProc("XSliderBar_SetPos")
	xSliderBar_GetPos = xcDLL.MustFindProc("XSliderBar_GetPos")
	xSliderBar_GetButton = xcDLL.MustFindProc("XSliderBar_GetButton")
	xSliderBar_SetHorizon = xcDLL.MustFindProc("XSliderBar_SetHorizon")
}

/*
创建滑动条元素.

参数:
	x 元素x坐标.
	y 元素y坐标.
	cx 宽度.
	cy 高度.
	hParent 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
返回:
	元素句柄.
*/
func XSliderBar_Create(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xSliderBar_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
设置滑动范围.

参数:
	hEle 元素句柄.
	range 范围.
*/
func XSliderBar_SetRange(hEle HELE, irange int) {
	xSliderBar_SetRange.Call(
		uintptr(hEle),
		uintptr(irange))
}

/*
获取滚动范围.

参数:
	hEle 元素句柄.
返回:
	返回滚动范围.
*/
func XSliderBar_GetRange(hEle HELE) int {
	ret, _, _ := xSliderBar_GetRange.Call(uintptr(hEle))

	return int(ret)
}

/*
设置进度贴图.

参数:
	hEle 元素句柄.
	hImage 图片句柄.
*/
func XSliderBar_SetImageLoad(hEle HELE, hImage HIMAGE) {
	xSliderBar_SetImageLoad.Call(
		uintptr(hEle),
		uintptr(hImage))
}

/*
设置滑块按钮宽度.

参数:
	hEle 元素句柄.
	width 宽度.
*/
func XSliderBar_SetButtonWidth(hEle HELE, width int) {
	xSliderBar_SetButtonWidth.Call(
		uintptr(hEle),
		uintptr(width))
}

/*
设置滑块按钮高度.

参数:
	hEle 元素句柄.
	height 高度.
*/
func XSliderBar_SetButtonHeight(hEle HELE, height int) {
	xSliderBar_SetButtonHeight.Call(
		uintptr(hEle),
		uintptr(height))
}

/*
设置两端间隔大小.

参数:
	hEle 元素句柄.
	leftSize 左边间隔大小.
	rightSize 右边间隔大小.
*/
func XSliderBar_SetSpaceTwo(hEle HELE, leftSize, rightSize int) {
	xSliderBar_SetSpaceTwo.Call(
		uintptr(hEle),
		uintptr(leftSize),
		uintptr(rightSize))

}

/*
设置当前进度点.

参数:
	hEle 元素句柄.
	pos 进度点.
*/
func XSliderBar_SetPos(hEle HELE, pos int) {
	xSliderBar_SetPos.Call(
		uintptr(hEle),
		uintptr(pos))
}

/*
获取当前进度点.

参数:
	hEle 元素句柄.
返回:
	返回当前进度点.
*/
func XSliderBar_GetPos(hEle HELE) int {
	ret, _, _ := xSliderBar_GetPos.Call(uintptr(hEle))

	return int(ret)
}

/*
获取滑块按钮.

参数:
	hEle 元素句柄.
返回:
	按钮句柄.
*/
func XSliderBar_GetButton(hEle HELE) HELE {
	ret, _, _ := xSliderBar_GetButton.Call(uintptr(hEle))

	return HELE(ret)
}

/*
设置水平或垂直.

参数:
hEle 元素句柄.
bHorizon 水平或垂直.
*/
func XSliderBar_SetHorizon(hEle HELE, bHorizon bool) {
	xSliderBar_SetHorizon.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bHorizon)))
}
