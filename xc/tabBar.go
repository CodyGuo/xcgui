package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xTabBar_Create            *syscall.Proc
	xTabBar_AddLabel          *syscall.Proc
	xTabBar_InsertLabel       *syscall.Proc
	xTabBar_DeleteLabel       *syscall.Proc
	xTabBar_DeleteLabelAll    *syscall.Proc
	xTabBar_GetLabel          *syscall.Proc
	xTabBar_GetLabelClose     *syscall.Proc
	xTabBar_GetButtonLeft     *syscall.Proc
	xTabBar_GetButtonRight    *syscall.Proc
	xTabBar_GetSelect         *syscall.Proc
	xTabBar_GetLabelSpacing   *syscall.Proc
	xTabBar_GetLabelCount     *syscall.Proc
	xTabBar_SetLabelSpacing   *syscall.Proc
	xTabBar_SetSelect         *syscall.Proc
	xTabBar_SetUp             *syscall.Proc
	xTabBar_SetDown           *syscall.Proc
	xTabBar_EnableTile        *syscall.Proc
	xTabBar_EnableClose       *syscall.Proc
	xTabBar_SetCloseSize      *syscall.Proc
	xTabBar_SetTurnButtonSize *syscall.Proc
	xTabBar_ShowLabel         *syscall.Proc
)

func init() {
	// Functions
	xTabBar_Create = XCDLL.MustFindProc("XTabBar_Create")
	xTabBar_AddLabel = XCDLL.MustFindProc("XTabBar_AddLabel")
	xTabBar_InsertLabel = XCDLL.MustFindProc("XTabBar_InsertLabel")
	xTabBar_DeleteLabel = XCDLL.MustFindProc("XTabBar_DeleteLabel")
	xTabBar_DeleteLabelAll = XCDLL.MustFindProc("XTabBar_DeleteLabelAll")
	xTabBar_GetLabel = XCDLL.MustFindProc("XTabBar_GetLabel")
	xTabBar_GetLabelClose = XCDLL.MustFindProc("XTabBar_GetLabelClose")
	xTabBar_GetButtonLeft = XCDLL.MustFindProc("XTabBar_GetButtonLeft")
	xTabBar_GetButtonRight = XCDLL.MustFindProc("XTabBar_GetButtonRight")
	xTabBar_GetSelect = XCDLL.MustFindProc("XTabBar_GetSelect")
	xTabBar_GetLabelSpacing = XCDLL.MustFindProc("XTabBar_GetLabelSpacing")
	xTabBar_GetLabelCount = XCDLL.MustFindProc("XTabBar_GetLabelCount")
	xTabBar_SetLabelSpacing = XCDLL.MustFindProc("XTabBar_SetLabelSpacing")
	xTabBar_SetSelect = XCDLL.MustFindProc("XTabBar_SetSelect")
	xTabBar_SetUp = XCDLL.MustFindProc("XTabBar_SetUp")
	xTabBar_SetDown = XCDLL.MustFindProc("XTabBar_SetDown")
	xTabBar_EnableTile = XCDLL.MustFindProc("XTabBar_EnableTile")
	xTabBar_EnableClose = XCDLL.MustFindProc("XTabBar_EnableClose")
	xTabBar_SetCloseSize = XCDLL.MustFindProc("XTabBar_SetCloseSize")
	xTabBar_SetTurnButtonSize = XCDLL.MustFindProc("XTabBar_SetTurnButtonSize")
	xTabBar_ShowLabel = XCDLL.MustFindProc("XTabBar_ShowLabel")
}

func XTabBarCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xTabBar_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

func XTabBarAddLabel(hEle HELE, pName string) int {
	ret, _, _ := xTabBar_AddLabel.Call(
		uintptr(hEle),
		StringToUintPtr(pName))

	return int(ret)
}

func XTabBarInsertLabel(hEle HELE, index int, pName string) int {
	ret, _, _ := xTabBar_InsertLabel.Call(
		uintptr(hEle),
		uintptr(index),
		StringToUintPtr(pName))

	return int(ret)
}

func XTabBarDeleteLabel(hEle HELE, index int) bool {
	ret, _, _ := xTabBar_DeleteLabel.Call(
		uintptr(hEle),
		uintptr(index))

	if ret != TRUE {
		return false
	}

	return true
}

func XTabBarDeleteLabelAll(hEle HELE) {
	xTabBar_DeleteLabelAll.Call(uintptr(hEle))
}

func XTabBarGetLabel(hEle HELE, index int) HELE {
	ret, _, _ := xTabBar_GetLabel.Call(
		uintptr(hEle),
		uintptr(index))

	return HELE(ret)
}

func XTabBarGetLabelClose(hEle HELE, index int) HELE {
	ret, _, _ := xTabBar_GetLabelClose.Call(
		uintptr(hEle),
		uintptr(index))

	return HELE(ret)
}

func XTabBarGetButtonLeft(hEle HELE) HELE {
	ret, _, _ := xTabBar_GetButtonLeft.Call(uintptr(hEle))

	return HELE(ret)
}

func XTabBarGetButtonRight(hEle HELE) HELE {
	ret, _, _ := xTabBar_GetButtonRight.Call(uintptr(hEle))

	return HELE(ret)
}

func XTabBarGetSelect(hEle HELE) int {
	ret, _, _ := xTabBar_GetSelect.Call(uintptr(hEle))

	return int(ret)
}

func XTabBarGetLabelSpacing(hEle HELE) int {
	ret, _, _ := xTabBar_GetLabelSpacing.Call(uintptr(hEle))

	return int(ret)
}

func XTabBarGetLabelCount(hEle HELE) int {
	ret, _, _ := xTabBar_GetLabelCount.Call(uintptr(hEle))

	return int(ret)
}

func XTabBarSetLabelSpacing(hEle HELE, spacing int) {
	xTabBar_SetLabelSpacing.Call(
		uintptr(hEle),
		uintptr(spacing))
}

func XTabBarSetSelect(hEle HELE, index int) {
	xTabBar_SetSelect.Call(
		uintptr(hEle),
		uintptr(index))
}

func XTabBarSetUp(hEle HELE) {
	xTabBar_SetUp.Call(uintptr(hEle))
}

func XTabBarSetDown(hEle HELE) {
	xTabBar_SetDown.Call(uintptr(hEle))
}

func XTabBarEnableTile(hEle HELE, bTile bool) {
	xTabBar_EnableTile.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bTile)))
}

func XTabBarEnableClose(hEle HELE, bEnable bool) {
	xTabBar_EnableClose.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

func XTabBarSetCloseSize(hEle HELE, pSize *SIZE) {
	xTabBar_SetCloseSize.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pSize)))
}

func XTabBarSetTurnButtonSize(hEle HELE, pSize *SIZE) {
	xTabBar_SetTurnButtonSize.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pSize)))
}

func XTabBarShowLabel(hEle HELE, index int, bShow bool) bool {
	ret, _, _ := xTabBar_ShowLabel.Call(
		uintptr(hEle),
		uintptr(index),
		uintptr(BoolToBOOL(bShow)))

	if ret != TRUE {
		return false
	}

	return true
}
