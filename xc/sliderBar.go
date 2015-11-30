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

func XSliderBarCreate(x int, y int, cx int, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xSliderBar_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

func XSliderBarSetRange(hEle HELE, irange int) {
	xSliderBar_SetRange.Call(
		uintptr(hEle),
		uintptr(irange))
}

func XSliderBarGetRange(hEle HELE) int {
	ret, _, _ := xSliderBar_GetRange.Call(
		uintptr(hEle))
	return int(ret)
}

func XSliderBarSetImageLoad(hEle HELE, hImage HIMAGE) {
	xSliderBar_SetImageLoad.Call(
		uintptr(hEle),
		uintptr(hImage))
}

func XSliderBarSetButtonWidth(hEle HELE, width int) {
	xSliderBar_SetButtonWidth.Call(
		uintptr(hEle),
		uintptr(width))
}

func XSliderBarSetButtonHeight(hEle HELE, height int) {
	xSliderBar_SetButtonHeight.Call(
		uintptr(hEle),
		uintptr(height))
}

func XSliderBarSetSpaceTwo(hEle HELE, leftSize int, rightSize int) {
	xSliderBar_SetSpaceTwo.Call(
		uintptr(hEle),
		uintptr(leftSize),
		uintptr(rightSize))

}

func XSliderBarSetPos(hEle HELE, pos int) {
	xSliderBar_SetPos.Call(
		uintptr(hEle),
		uintptr(pos))
}

func XSliderBarGetPos(hEle HELE) int {
	ret, _, _ := xSliderBar_GetPos.Call(uintptr(hEle))
	return int(ret)
}

func XSliderBarGetButton(hEle HELE) HELE {
	ret, _, _ := xSliderBar_GetButton.Call(uintptr(hEle))
	return HELE(ret)
}

func XSliderBarSetHorizon(hEle HELE, bHorizon bool) {
	xSliderBar_SetHorizon.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bHorizon)))
}
