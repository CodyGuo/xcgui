package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xMenu_Create               *syscall.Proc
	xMenu_AddItem              *syscall.Proc
	xMenu_AddItemIcon          *syscall.Proc
	xMenu_InsertItem           *syscall.Proc
	xMenu_InsertItemIcon       *syscall.Proc
	xMenu_SetAutoDestroy       *syscall.Proc
	xMenu_EnableDrawBackground *syscall.Proc
	xMenu_EnableDrawItem       *syscall.Proc
	xMenu_Popup                *syscall.Proc
	xMenu_DestroyMenu          *syscall.Proc
	xMenu_CloseMenu            *syscall.Proc
	xMenu_SetBkImage           *syscall.Proc
	xMenu_SetItemText          *syscall.Proc
	xMenu_GetItemText          *syscall.Proc
	xMenu_GetItemTextLength    *syscall.Proc
	xMenu_SetItemIcon          *syscall.Proc
	xMenu_SetItemFlags         *syscall.Proc
	xMenu_SetItemHeight        *syscall.Proc
	xMenu_GetItemHeight        *syscall.Proc
	xMenu_SetBorderColor       *syscall.Proc
	xMenu_GetLeftWidth         *syscall.Proc
	xMenu_GetLeftSpaceText     *syscall.Proc
	xMenu_GetItemCount         *syscall.Proc
	xMenu_SetItemCheck         *syscall.Proc
	xMenu_IsItemCheck          *syscall.Proc
)

func init() {
	// Functions
	xMenu_Create = xcDLL.MustFindProc("XMenu_Create")
	xMenu_AddItem = xcDLL.MustFindProc("XMenu_AddItem")
	xMenu_AddItemIcon = xcDLL.MustFindProc("XMenu_AddItemIcon")
	xMenu_InsertItem = xcDLL.MustFindProc("XMenu_InsertItem")
	xMenu_InsertItemIcon = xcDLL.MustFindProc("XMenu_InsertItemIcon")
	xMenu_SetAutoDestroy = xcDLL.MustFindProc("XMenu_SetAutoDestroy")
	xMenu_EnableDrawBackground = xcDLL.MustFindProc("XMenu_EnableDrawBackground")
	xMenu_EnableDrawItem = xcDLL.MustFindProc("XMenu_EnableDrawItem")
	xMenu_Popup = xcDLL.MustFindProc("XMenu_Popup")
	xMenu_DestroyMenu = xcDLL.MustFindProc("XMenu_DestroyMenu")
	xMenu_CloseMenu = xcDLL.MustFindProc("XMenu_CloseMenu")
	xMenu_SetBkImage = xcDLL.MustFindProc("XMenu_SetBkImage")
	xMenu_SetItemText = xcDLL.MustFindProc("XMenu_SetItemText")
	xMenu_GetItemText = xcDLL.MustFindProc("XMenu_GetItemText")
	xMenu_GetItemTextLength = xcDLL.MustFindProc("XMenu_GetItemTextLength")
	xMenu_SetItemIcon = xcDLL.MustFindProc("XMenu_SetItemIcon")
	xMenu_SetItemFlags = xcDLL.MustFindProc("XMenu_SetItemFlags")
	xMenu_SetItemHeight = xcDLL.MustFindProc("XMenu_SetItemHeight")
	xMenu_GetItemHeight = xcDLL.MustFindProc("XMenu_GetItemHeight")
	xMenu_SetBorderColor = xcDLL.MustFindProc("XMenu_SetBorderColor")
	xMenu_GetLeftWidth = xcDLL.MustFindProc("XMenu_GetLeftWidth")
	xMenu_GetLeftSpaceText = xcDLL.MustFindProc("XMenu_GetLeftSpaceText")
	xMenu_GetItemCount = xcDLL.MustFindProc("XMenu_GetItemCount")
	xMenu_SetItemCheck = xcDLL.MustFindProc("XMenu_SetItemCheck")
	xMenu_IsItemCheck = xcDLL.MustFindProc("XMenu_IsItemCheck")

}

func XMenuCreate() HMENUX {
	ret, _, _ := xMenu_Create.Call()

	return HMENUX(ret)
}

func XMenuAddItem(hMenu HMENUX, nID int, pText string, nParentID int, nFlags MENU_STATE_FLAGS_) {
	xMenu_AddItem.Call(
		uintptr(hMenu),
		uintptr(nID),
		StringToUintPtr(pText),
		uintptr(nParentID),
		uintptr(nFlags))
}

func XMenuAddItemIcon(hMenu HMENUX, nID int, pText string, nParentID int, hIcon HIMAGE, nFlags MENU_STATE_FLAGS_) {
	xMenu_AddItemIcon.Call(
		uintptr(hMenu),
		uintptr(nID),
		StringToUintPtr(pText),
		uintptr(nParentID),
		uintptr(hIcon),
		uintptr(nFlags))
}

func XMenuInsertItem(hMenu HMENUX, nID int, pText string, nFlags MENU_STATE_FLAGS_, insertID int) {
	xMenu_InsertItem.Call(
		uintptr(hMenu),
		uintptr(nID),
		StringToUintPtr(pText),
		uintptr(nFlags),
		uintptr(insertID))
}

func XMenuInsertItemIcon(hMenu HMENUX, nID int, pText string, hIcon HIMAGE, nFlags MENU_STATE_FLAGS_, insertID int) {
	xMenu_InsertItemIcon.Call(
		uintptr(hMenu),
		uintptr(nID),
		StringToUintPtr(pText),
		uintptr(hIcon),
		uintptr(nFlags),
		uintptr(insertID))
}

func XMenuSetAutoDestroy(hMenu HMENUX, bAuto bool) {
	xMenu_SetAutoDestroy.Call(
		uintptr(hMenu),
		uintptr(BoolToBOOL(bAuto)))
}

func XMenuEnableDrawBackground(hMenu HMENUX, bEnable bool) {
	xMenu_EnableDrawBackground.Call(
		uintptr(hMenu),
		uintptr(BoolToBOOL(bEnable)))
}

func XMenuEnableDrawItem(hMenu HMENUX, bEnable bool) {
	xMenu_EnableDrawItem.Call(
		uintptr(hMenu),
		uintptr(BoolToBOOL(bEnable)))
}

func XMenuPopup(hMenu HMENUX, hParentWnd HWND, x, y int, hParentEle HELE, nPosition MENU_POPUP_POSITION_) bool {
	ret, _, _ := xMenu_Popup.Call(
		uintptr(hMenu),
		uintptr(hParentWnd),
		uintptr(x),
		uintptr(y),
		uintptr(hParentEle),
		uintptr(nPosition))

	if ret != TRUE {
		return false
	}

	return true

}

func XMenuDestroyMenu(hMenu HMENUX) {
	xMenu_DestroyMenu.Call(uintptr(hMenu))
}

func XMenuCloseMenu(hMenu HMENUX) {
	xMenu_CloseMenu.Call(uintptr(hMenu))
}

func XMenuSetBkImage(hMenu HMENUX, hImage HIMAGE) {
	xMenu_SetBkImage.Call(
		uintptr(hMenu),
		uintptr(hImage))
}

func XMenuSetItemText(hMenu HMENUX, nID int, pText string) bool {
	ret, _, _ := xMenu_SetItemText.Call(
		uintptr(hMenu),
		uintptr(nID),
		StringToUintPtr(pText))

	if ret != TRUE {
		return false
	}

	return true
}

func XMenuGetItemText(hMenu HMENUX, nID int, pOut *uint16, nOutLen int) bool {
	ret, _, _ := xMenu_GetItemText.Call(
		uintptr(hMenu),
		uintptr(nID),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))

	if ret != TRUE {
		return false
	}

	return true
}

func XMenuGetItemTextLength(hMenu HMENUX, nID int) int {
	ret, _, _ := xMenu_GetItemTextLength.Call(
		uintptr(hMenu),
		uintptr(nID))

	return int(ret)
}

func XMenuSetItemIcon(hMenu HMENUX, nID int, hIcon HIMAGE) bool {
	ret, _, _ := xMenu_SetItemIcon.Call(
		uintptr(hMenu),
		uintptr(nID),
		uintptr(hIcon))

	if ret != TRUE {
		return false
	}

	return true
}

func XMenuSetItemFlags(hMenu HMENUX, nID int, uFlags MENU_STATE_FLAGS_) bool {
	ret, _, _ := xMenu_SetItemFlags.Call(
		uintptr(hMenu),
		uintptr(nID),
		uintptr(uFlags))

	if ret != TRUE {
		return false
	}

	return true
}

func XMenuSetItemHeight(hMenu HMENUX, height int) {
	xMenu_SetItemHeight.Call(
		uintptr(hMenu),
		uintptr(height))
}

func XMenuGetItemHeight(hMenu HMENUX) int {
	ret, _, _ := xMenu_GetItemHeight.Call(uintptr(hMenu))

	return int(ret)
}

func XMenuSetBorderColor(hMenu HMENUX, crColor COLORREF, alpha byte) {
	xMenu_SetBorderColor.Call(
		uintptr(hMenu),
		uintptr(crColor),
		uintptr(alpha))
}

func XMenuGetLeftWidth(hMenu HMENUX) int {
	ret, _, _ := xMenu_GetLeftWidth.Call(uintptr(hMenu))

	return int(ret)
}

func XMenuGetLeftSpaceText(hMenu HMENUX) int {
	ret, _, _ := xMenu_GetLeftSpaceText.Call(uintptr(hMenu))

	return int(ret)
}

func XMenuGetItemCount(hMenu HMENUX) int {
	ret, _, _ := xMenu_GetItemCount.Call(uintptr(hMenu))

	return int(ret)
}

func XMenuSetItemCheck(hMenu HMENUX, nID int, bCheck bool) {
	xMenu_SetItemCheck.Call(
		uintptr(hMenu),
		uintptr(nID),
		uintptr(BoolToBOOL(bCheck)))
}

func XMenuIsItemCheck(hMenu HMENUX, nID int) bool {
	ret, _, _ := xMenu_IsItemCheck.Call(
		uintptr(hMenu),
		uintptr(nID))

	if ret != TRUE {
		return false
	}

	return true
}
