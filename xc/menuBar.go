package xc

import (
	"syscall"
)

var (
	xMenuBar_Create          *syscall.Proc
	xMenuBar_AddButton       *syscall.Proc
	xMenuBar_SetButtonHeight *syscall.Proc
	xMenuBar_GetMenu         *syscall.Proc
	xMenuBar_DeleteButton    *syscall.Proc
)

func init() {
	xMenuBar_Create = xcDLL.MustFindProc("XMenuBar_Create")
	xMenuBar_AddButton = xcDLL.MustFindProc("XMenuBar_AddButton")
	xMenuBar_SetButtonHeight = xcDLL.MustFindProc("XMenuBar_SetButtonHeight")
	xMenuBar_GetMenu = xcDLL.MustFindProc("XMenuBar_GetMenu")
	xMenuBar_DeleteButton = xcDLL.MustFindProc("XMenuBar_DeleteButton")
}

func XMenuBarCreate(x int, y int, cx int, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xMenuBar_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))
	return HELE(ret)
}

func XMenuBarAddButton(hEle HELE, pText string) int {
	ret, _, _ := xMenuBar_AddButton.Call(
		uintptr(hEle),
		StringToUintPtr(pText))
	return int(ret)
}

func XMenuBarSetButtonHeight(hEle HELE, height int) {
	xMenuBar_SetButtonHeight.Call(
		uintptr(hEle),
		uintptr(height))
}

func XMenuBarGetMenu(hEle HELE, nIndex int) HMENUX {
	ret, _, _ := xMenuBar_GetMenu.Call(
		uintptr(hEle),
		uintptr(nIndex))
	return HMENUX(ret)
}

func XMenuBarDeleteButton(hEle HELE, nIndex int) bool {
	ret, _, _ := xMenuBar_GetMenu.Call(
		uintptr(hEle),
		uintptr(nIndex))
	return ret == TRUE
}
