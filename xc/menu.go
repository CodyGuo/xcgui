package xc

import (
	"syscall"
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
	xMenu_Create = XCDLL.MustFindProc("XMenu_Create")
	xMenu_AddItem = XCDLL.MustFindProc("XMenu_AddItem")
	xMenu_AddItemIcon = XCDLL.MustFindProc("XMenu_AddItemIcon")
	xMenu_InsertItem = XCDLL.MustFindProc("XMenu_InsertItem")
	xMenu_InsertItemIcon = XCDLL.MustFindProc("XMenu_InsertItemIcon")
	xMenu_SetAutoDestroy = XCDLL.MustFindProc("XMenu_SetAutoDestroy")
	xMenu_EnableDrawBackground = XCDLL.MustFindProc("XMenu_EnableDrawBackground")
	xMenu_EnableDrawItem = XCDLL.MustFindProc("XMenu_EnableDrawItem")
	xMenu_Popup = XCDLL.MustFindProc("XMenu_Popup")
	xMenu_DestroyMenu = XCDLL.MustFindProc("XMenu_DestroyMenu")
	xMenu_CloseMenu = XCDLL.MustFindProc("XMenu_CloseMenu")
	xMenu_SetBkImage = XCDLL.MustFindProc("XMenu_SetBkImage")
	xMenu_SetItemText = XCDLL.MustFindProc("XMenu_SetItemText")
	xMenu_GetItemText = XCDLL.MustFindProc("XMenu_GetItemText")
	xMenu_GetItemTextLength = XCDLL.MustFindProc("XMenu_GetItemTextLength")
	xMenu_SetItemIcon = XCDLL.MustFindProc("XMenu_SetItemIcon")
	xMenu_SetItemFlags = XCDLL.MustFindProc("XMenu_SetItemFlags")
	xMenu_SetItemHeight = XCDLL.MustFindProc("XMenu_SetItemHeight")
	xMenu_GetItemHeight = XCDLL.MustFindProc("XMenu_GetItemHeight")
	xMenu_SetBorderColor = XCDLL.MustFindProc("XMenu_SetBorderColor")
	xMenu_GetLeftWidth = XCDLL.MustFindProc("XMenu_GetLeftWidth")
	xMenu_GetLeftSpaceText = XCDLL.MustFindProc("XMenu_GetLeftSpaceText")
	xMenu_GetItemCount = XCDLL.MustFindProc("XMenu_GetItemCount")
	xMenu_SetItemCheck = XCDLL.MustFindProc("XMenu_SetItemCheck")
	xMenu_IsItemCheck = XCDLL.MustFindProc("XMenu_IsItemCheck")

}

// xMenu_Create
// xMenu_AddItem
// xMenu_AddItemIcon
// xMenu_InsertItem
// xMenu_InsertItemIcon
// xMenu_SetAutoDestroy
// xMenu_EnableDrawBackground
// xMenu_EnableDrawItem
// xMenu_Popup
// xMenu_DestroyMenu
// xMenu_CloseMenu
// xMenu_SetBkImage
// xMenu_SetItemText
// xMenu_GetItemText
// xMenu_GetItemTextLength
// xMenu_SetItemIcon
// xMenu_SetItemFlags
// xMenu_SetItemHeight
// xMenu_GetItemHeight
// xMenu_SetBorderColor
// xMenu_GetLeftWidth
// xMenu_GetLeftSpaceText
// xMenu_GetItemCount
// xMenu_SetItemCheck
// xMenu_IsItemCheck
