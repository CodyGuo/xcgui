package xc

import (
	"syscall"
)

var (
	xProgBar_Create       *syscall.Proc
	xProgBar_SetRange     *syscall.Proc
	xProgBar_GetRange     *syscall.Proc
	xProgBar_SetImageLoad *syscall.Proc
	xProgBar_SetSpaceTwo  *syscall.Proc
	xProgBar_SetPos       *syscall.Proc
	xProgBar_GetPos       *syscall.Proc
	xProgBar_SetHorizon   *syscall.Proc
)

func init() {
	xProgBar_Create = xcDLL.MustFindProc("XProgBar_Create")
	xProgBar_SetRange = xcDLL.MustFindProc("XProgBar_SetRange")
	xProgBar_GetRange = xcDLL.MustFindProc("XProgBar_GetRange")
	xProgBar_SetImageLoad = xcDLL.MustFindProc("XProgBar_SetImageLoad")
	xProgBar_SetSpaceTwo = xcDLL.MustFindProc("XProgBar_SetSpaceTwo")
	xProgBar_SetPos = xcDLL.MustFindProc("XProgBar_SetPos")
	xProgBar_GetPos = xcDLL.MustFindProc("XProgBar_GetPos")
	xProgBar_SetHorizon = xcDLL.MustFindProc("XProgBar_SetHorizon")
}

/*
创建进度条元素.

参数:
	x 元素x坐标.
	y 元素y坐标.
	cx 宽度.
	cy 高度.
	hParent 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
返回:
	元素句柄.
*/
func XProgBar_Create(x int, y int, cx int, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xProgBar_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
设置范围.

参数:
	hEle 元句柄.
	range 范围.
*/
func XProgBar_SetRange(hEle HELE, barRange int) {
	xProgBar_SetRange.Call(
		uintptr(hEle),
		uintptr(barRange))
}

/*
获取范围.

参数:
	hEle 元句柄.
返回:
	返回范围.
*/
func XProgBar_GetRange(hEle HELE) int {
	ret, _, _ := xProgBar_GetRange.Call(uintptr(hEle))

	return int(ret)
}

/*
设置进度贴图.

参数:
	hEle 元句柄.
	hImage 图片句柄.
*/
func XProgBar_SetImageLoad(hEle HELE, hImage HIMAGE) {
	xProgBar_SetImageLoad.Call(
		uintptr(hEle),
		uintptr(hImage))
}

/*
设置两端间隔大小.

参数:
	hEle 元句柄.
	leftSize 左边间隔大小.
	rightSize 右边间隔大小.
*/
func XProgBar_SetSpaceTwo(hEle HELE, leftSize int, rightSize int) {
	xProgBar_SetSpaceTwo.Call(
		uintptr(hEle),
		uintptr(leftSize),
		uintptr(rightSize))
}

/*
设置位置点.

参数:
	hEle 元句柄.
	pos 位置点.
*/
func XProgBar_SetPos(hEle HELE, pos int) {
	xProgBar_SetPos.Call(
		uintptr(hEle),
		uintptr(pos))
}

/*
获取当前位置点.

参数:
	hEle 元句柄.
返回:
	返回当前位置点.
*/
func XProgBar_GetPos(hEle HELE) int {
	ret, _, _ := xProgBar_GetPos.Call(uintptr(hEle))

	return int(ret)
}

/*
设置水平或垂直.

参数:
	hEle 元句柄.
	bHorizon 水平或垂直.
*/
func XProgBar_SetHorizon(hEle HELE, bHorizon bool) {
	xProgBar_SetPos.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bHorizon)))
}
