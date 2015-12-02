package xc

import (
	"syscall"
)

var (
	xSBar_Create             *syscall.Proc
	xSBar_SetRange           *syscall.Proc
	xSBar_GetRange           *syscall.Proc
	xSBar_ShowButton         *syscall.Proc
	xSBar_SetSliderLength    *syscall.Proc
	xSBar_SetSliderMinLength *syscall.Proc
	xSBar_SetHorizon         *syscall.Proc
	xSBar_GetSliderMaxLength *syscall.Proc
	xSBar_ScrollUp           *syscall.Proc
	xSBar_ScrollDown         *syscall.Proc
	xSBar_ScrollTop          *syscall.Proc
	xSBar_ScrollBottom       *syscall.Proc
	xSBar_ScrollPos          *syscall.Proc
	xSBar_GetButtonUp        *syscall.Proc
	xSBar_GetButtonDown      *syscall.Proc
	xSBar_GetButtonSlider    *syscall.Proc
)

func init() {
	xSBar_Create = xcDLL.MustFindProc("XSBar_Create")
	xSBar_SetRange = xcDLL.MustFindProc("XSBar_SetRange")
	xSBar_GetRange = xcDLL.MustFindProc("XSBar_GetRange")
	xSBar_ShowButton = xcDLL.MustFindProc("XSBar_ShowButton")
	xSBar_SetSliderLength = xcDLL.MustFindProc("XSBar_SetSliderLength")
	xSBar_SetSliderMinLength = xcDLL.MustFindProc("XSBar_SetSliderMinLength")
	xSBar_SetHorizon = xcDLL.MustFindProc("XSBar_SetHorizon")
	xSBar_GetSliderMaxLength = xcDLL.MustFindProc("XSBar_GetSliderMaxLength")
	xSBar_ScrollUp = xcDLL.MustFindProc("XSBar_ScrollUp")
	xSBar_ScrollDown = xcDLL.MustFindProc("XSBar_ScrollDown")
	xSBar_ScrollTop = xcDLL.MustFindProc("XSBar_ScrollTop")
	xSBar_ScrollBottom = xcDLL.MustFindProc("XSBar_ScrollBottom")
	xSBar_ScrollPos = xcDLL.MustFindProc("XSBar_ScrollPos")
	xSBar_GetButtonUp = xcDLL.MustFindProc("XSBar_GetButtonUp")
	xSBar_GetButtonDown = xcDLL.MustFindProc("XSBar_GetButtonDown")
	xSBar_GetButtonSlider = xcDLL.MustFindProc("XSBar_GetButtonSlider")
}

/*
创建滚动条元素.

参数:
	x 元素x坐标.
	y 元素y坐标.
	cx 宽度.
	cy 高度.
	hParent 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
返回:
	元素句柄.
*/
func XSBarCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xSBar_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
设置滚动范围.

参数:
	hEle 元素句柄.
	range 范围.
*/
func XSBarSetRange(hEle HELE, irange int) {
	xSBar_SetRange.Call(
		uintptr(hEle),
		uintptr(irange))
}

/*
获取滚动范围.

参数:
	hEle 元素句柄.
	返回:滚动范围.
*/
func XSBarGetRange(hEle HELE) int {
	ret, _, _ := xSBar_GetRange.Call(
		uintptr(hEle))

	return int(ret)
}

/*
显示隐藏滚动条上下按钮.

参数:
	hEle 元素句柄.
	bShow 是否显示.
*/
func XSBarShowButton(hEle HELE, bShow bool) {
	xSBar_ShowButton.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bShow)))
}

/*
设置滑块长度.

参数:
	hEle 元素句柄.
	length 长度.
*/
func XSBarSetSliderLength(hEle HELE, length int) {
	xSBar_SetSliderLength.Call(
		uintptr(hEle),
		uintptr(length))
}

/*
设置滑块最小长度.

参数:
	hEle 元素句柄.
	minLength 长度.
*/
func XSBarSetSliderMinLength(hEle HELE, minLength int) {
	xSBar_SetSliderMinLength.Call(
		uintptr(hEle),
		uintptr(minLength))
}

/*
设置水平或者垂直.

参数:
	hEle 元素句柄.
	bHorizon 水平或垂直.
返回:
	如果改变返回TRUE否则返回FALSE.
*/
func XSBarSetHorizon(hEle HELE, bHorizon bool) bool {
	ret, _, _ := xSBar_SetHorizon.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bHorizon)))

	return ret == TRUE
}

/*
获取滑块最大长度.

参数:
	hEle 元素句柄.
返回:
	长度.
*/
func XSBarGetSliderMaxLength(hEle HELE) int {
	ret, _, _ := xSBar_GetSliderMaxLength.Call(
		uintptr(hEle))

	return int(ret)
}

/*
向上滚动.

参数:
	hEle 元素句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XSBarScrollUp(hEle HELE) bool {
	ret, _, _ := xSBar_ScrollUp.Call(
		uintptr(hEle))

	return ret == TRUE
}

/*
向下滚动.

参数:
	hEle 元素句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XSBarScrollDown(hEle HELE) bool {
	ret, _, _ := xSBar_ScrollDown.Call(
		uintptr(hEle))

	return ret == TRUE
}

/*
滚动到顶部.

参数:
	hEle 元素句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XSBarScrollTop(hEle HELE) bool {
	ret, _, _ := xSBar_ScrollTop.Call(
		uintptr(hEle))

	return ret == TRUE
}

/*
滚动到底部.

参数:
	hEle 元素句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XSBarScrollBottom(hEle HELE) bool {
	ret, _, _ := xSBar_ScrollBottom.Call(
		uintptr(hEle))

	return ret == TRUE
}

/*
滚动到指定位置点.

参数:
	hEle 元素句柄.
	pos 位置点.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XSBarScrollPos(hEle HELE) bool {
	ret, _, _ := xSBar_ScrollPos.Call(
		uintptr(hEle))

	return ret == TRUE
}

/*
获取上按钮.

参数:
	hEle 元素句柄.
返回:
	返回按钮句柄.
*/
func XSBarGetButtonUp(hEle HELE) HELE {
	ret, _, _ := xSBar_GetButtonUp.Call(
		uintptr(hEle))

	return HELE(ret)
}

/*
获取下按钮.

参数:
	hEle 元素句柄.
返回:
	返回按钮句柄.
*/
func XSBarGetButtonDown(hEle HELE) HELE {
	ret, _, _ := xSBar_GetButtonDown.Call(
		uintptr(hEle))

	return HELE(ret)
}

/*
获取滑动按钮.

参数:
	hEle 元素句柄.
返回:
	返回按钮句柄.
*/
func XSBarGetButtonSlider(hEle HELE) HELE {
	ret, _, _ := xSBar_GetButtonSlider.Call(
		uintptr(hEle))

	return HELE(ret)
}
