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

func XSBarCreate(x int, y int, cx int, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xSBar_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))
	return HELE(ret)
}

func XSBarSetRange(hEle HELE, irange int) {
	xSBar_SetRange.Call(
		uintptr(hEle),

		uintptr(irange))
}

func XSBarGetRange(hEle HELE) int {
	ret, _, _ := xSBar_GetRange.Call(
		uintptr(hEle))
	return int(ret)
}

func XSBarShowButton(hEle HELE, bShow bool) {
	xSBar_ShowButton.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bShow)))
}

func XSBarSetSliderLength(hEle HELE, length int) {
	xSBar_SetSliderLength.Call(
		uintptr(hEle),
		uintptr(length))
}

func XSBarSetSliderMinLength(hEle HELE, minLength int) {
	xSBar_SetSliderMinLength.Call(
		uintptr(hEle),
		uintptr(minLength))
}

func XSBarSetHorizon(hEle HELE, bHorizon bool) bool {
	ret, _, _ := xSBar_SetHorizon.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bHorizon)))
	return ret == TRUE
}

func XSBarGetSliderMaxLength(hEle HELE) int {
	ret, _, _ := xSBar_GetSliderMaxLength.Call(
		uintptr(hEle))
	return int(ret)
}

func XSBarScrollUp(hEle HELE) bool {
	ret, _, _ := xSBar_ScrollUp.Call(
		uintptr(hEle))
	return ret == TRUE
}

func XSBarScrollDown(hEle HELE) bool {
	ret, _, _ := xSBar_ScrollDown.Call(
		uintptr(hEle))
	return ret == TRUE
}

func XSBarScrollTop(hEle HELE) bool {
	ret, _, _ := xSBar_ScrollTop.Call(
		uintptr(hEle))
	return ret == TRUE
}

func XSBarScrollBottom(hEle HELE) bool {
	ret, _, _ := xSBar_ScrollBottom.Call(
		uintptr(hEle))
	return ret == TRUE
}

func XSBarScrollPos(hEle HELE) bool {
	ret, _, _ := xSBar_ScrollPos.Call(
		uintptr(hEle))
	return ret == TRUE
}

func XSBarGetButtonUp(hEle HELE) HELE {
	ret, _, _ := xSBar_GetButtonUp.Call(
		uintptr(hEle))
	return HELE(ret)
}

func XSBarGetButtonDown(hEle HELE) HELE {
	ret, _, _ := xSBar_GetButtonDown.Call(
		uintptr(hEle))
	return HELE(ret)
}

func XSBarGetButtonSlider(hEle HELE) HELE {
	ret, _, _ := xSBar_GetButtonSlider.Call(
		uintptr(hEle))
	return HELE(ret)
}
