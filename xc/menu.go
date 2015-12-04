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

/*
创建菜单.默认弹出菜单窗口关闭后自动销毁.

返回:
	菜单句柄.
*/
func XMenuCreate() HMENUX {
	ret, _, _ := xMenu_Create.Call()

	return HMENUX(ret)
}

/*
添加菜单项.

参数:
	hMenu 菜单句柄.
	nID 项ID.
	pText 文本内容.*uint16
	nParentID 父项ID.
	nFlags 标识参见宏定义 menu_state_flags_.
*/
func XMenuAddItem(hMenu HMENUX, nID int, pText string, nParentID int, nFlags Menu_state_flags_) {
	xMenu_AddItem.Call(
		uintptr(hMenu),
		uintptr(nID),
		StringToUintPtr(pText),
		// uintptr(unsafe.Pointer(pText)),
		uintptr(nParentID),
		uintptr(nFlags))
}

/*
添加菜单项.

参数:
	hMenu 菜单句柄.
	nID 项ID.
	pText 文本内容.*uint16
	nParentID 父项ID.
	hIcon 菜单项图标句柄.
	nFlags 标识参见宏定义 menu_state_flags_.
*/
func XMenuAddItemIcon(hMenu HMENUX, nID int, pText string, nParentID int, hIcon HIMAGE, nFlags Menu_state_flags_) {
	xMenu_AddItemIcon.Call(
		uintptr(hMenu),
		uintptr(nID),
		StringToUintPtr(pText),
		// uintptr(unsafe.Pointer(pText)),
		uintptr(nParentID),
		uintptr(hIcon),
		uintptr(nFlags))
}

/*
插入菜单项.

参数:
	hMenu 菜单句柄.
	nID 项ID.
	pText 文本内容.*uint16
	nFlags 标识参见宏定义 menu_state_flags_.
	insertID 插入位置ID.
*/
func XMenuInsertItem(hMenu HMENUX, nID int, pText string, nFlags Menu_state_flags_, insertID int) {
	xMenu_InsertItem.Call(
		uintptr(hMenu),
		uintptr(nID),
		StringToUintPtr(pText),
		// uintptr(unsafe.Pointer(pText)),
		uintptr(nFlags),
		uintptr(insertID))
}

/*
插入菜单项.

参数:
	hMenu 菜单句柄.
	nID 项ID.
	pText 文本内容. *uint16
	hIcon 菜单项图标句柄.
	nFlags 标识参见宏定义 menu_state_flags_.
	insertID 插入位置ID.
*/
func XMenuInsertItemIcon(hMenu HMENUX, nID int, pText string, hIcon HIMAGE, nFlags Menu_state_flags_, insertID int) {
	xMenu_InsertItemIcon.Call(
		uintptr(hMenu),
		uintptr(nID),
		StringToUintPtr(pText),
		// uintptr(unsafe.Pointer(pText)),
		uintptr(hIcon),
		uintptr(nFlags),
		uintptr(insertID))
}

/*
设置是否自动销毁菜单.

参数:
	hMenu 菜单句柄.
	bAuto 是否自动销毁.
*/
func XMenuSetAutoDestroy(hMenu HMENUX, bAuto bool) {
	xMenu_SetAutoDestroy.Call(
		uintptr(hMenu),
		uintptr(BoolToBOOL(bAuto)))
}

/*
是否有用户绘制菜单背景,如果启用XWM_MENU_DRAW_BACKGROUND和XE_MENU_DRAW_BACKGROUND事件有效.

参数:
	hMenu 菜单句柄.
	bEnable 是否启用.
*/
func XMenuEnableDrawBackground(hMenu HMENUX, bEnable bool) {
	xMenu_EnableDrawBackground.Call(
		uintptr(hMenu),
		uintptr(BoolToBOOL(bEnable)))
}

/*
是否有用户绘制菜单项,如果启用XWM_MENU_DRAWITEM和XE_MENU_DRAWITEM事件有效.

参数:
	hMenu 菜单句柄.
	bEnable 是否启用.
*/
func XMenuEnableDrawItem(hMenu HMENUX, bEnable bool) {
	xMenu_EnableDrawItem.Call(
		uintptr(hMenu),
		uintptr(BoolToBOOL(bEnable)))
}

/*
弹出菜单.

参数:
	hMenu 菜单句柄.
	hParentWnd 父窗口句柄.
	x x坐标.
	y y坐标.
	hParentEle 父元素句柄,如果该值不为NULL,hParentEle元素将接收菜单消息事件, 否则将由hParentWnd窗口接收菜单的消息事件
	nPosition 弹出位置,参见宏定义.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XMenuPopup(hMenu HMENUX, hParentWnd HWND, x, y int, hParentEle HELE, nPosition Menu_popup_position_) bool {
	ret, _, _ := xMenu_Popup.Call(
		uintptr(hMenu),
		uintptr(hParentWnd),
		uintptr(x),
		uintptr(y),
		uintptr(hParentEle),
		uintptr(nPosition))

	return ret == TRUE

}

/*
销毁菜单.

参数:
	hMenu 菜单句柄.
*/
func XMenuDestroyMenu(hMenu HMENUX) {
	xMenu_DestroyMenu.Call(uintptr(hMenu))
}

/*
关闭菜单.

参数:
	hMenu 菜单句柄.
*/
func XMenuCloseMenu(hMenu HMENUX) {
	xMenu_CloseMenu.Call(uintptr(hMenu))
}

/*
设置菜单背景图片.

参数:
	hMenu 菜单句柄.
	hImage 图片句柄.
*/
func XMenuSetBkImage(hMenu HMENUX, hImage HIMAGE) {
	xMenu_SetBkImage.Call(
		uintptr(hMenu),
		uintptr(hImage))
}

/*
设置项文本.

参数:
	hMenu 菜单句柄.
	nID 项ID.
	pText 文本内容.*uint16
返回:
	成功返回TRUE否则返回FALSE.
*/
func XMenuSetItemText(hMenu HMENUX, nID int, pText string) bool {
	ret, _, _ := xMenu_SetItemText.Call(
		uintptr(hMenu),
		uintptr(nID),
		StringToUintPtr(pText))
	// uintptr(unsafe.Pointer(pText)))

	return ret == TRUE
}

/*
获取项文件.

参数:
	hMenu 菜单句柄.
	nID 项ID.
	pOut 接收内容缓冲区.
	nOutLen 缓冲区长度,字符为单位.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XMenuGetItemText(hMenu HMENUX, nID int, pOut *uint16, nOutLen int) bool {
	ret, _, _ := xMenu_GetItemText.Call(
		uintptr(hMenu),
		uintptr(nID),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))

	return ret == TRUE
}

/*
获取项文本长度,不包含字符串空终止符.

参数:
	hMenu 菜单句柄.
	nID 项ID.
返回:
	长度,字符为单位.
*/
func XMenuGetItemTextLength(hMenu HMENUX, nID int) int {
	ret, _, _ := xMenu_GetItemTextLength.Call(
		uintptr(hMenu),
		uintptr(nID))

	return int(ret)
}

/*
设置菜单项图标.

参数:
	hMenu 菜单句柄.
	nID 项ID.
	hIcon 菜单项图标句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XMenuSetItemIcon(hMenu HMENUX, nID int, hIcon HIMAGE) bool {
	ret, _, _ := xMenu_SetItemIcon.Call(
		uintptr(hMenu),
		uintptr(nID),
		uintptr(hIcon))

	return ret == TRUE
}

/*
设置项标识.

参数:
	hMenu 菜单句柄.
	nID 项ID.
	uFlags 标识参见宏定义 menu_state_flags_.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XMenuSetItemFlags(hMenu HMENUX, nID int, uFlags Menu_state_flags_) bool {
	ret, _, _ := xMenu_SetItemFlags.Call(
		uintptr(hMenu),
		uintptr(nID),
		uintptr(uFlags))

	return ret == TRUE
}

/*
设置项高度.

参数:
	hMenu 菜单句柄.
	height 高度.
*/
func XMenuSetItemHeight(hMenu HMENUX, height int) {
	xMenu_SetItemHeight.Call(
		uintptr(hMenu),
		uintptr(height))
}

/*
获取项高度.

参数:
	hMenu 菜单句柄.
返回:
	返回项高度.
*/
func XMenuGetItemHeight(hMenu HMENUX) int {
	ret, _, _ := xMenu_GetItemHeight.Call(uintptr(hMenu))

	return int(ret)
}

/*
设置菜单边框颜色.

参数:
	hMenu 菜单句柄.
	crColor 颜色.
	alpha 透明度0-255.
*/
func XMenuSetBorderColor(hMenu HMENUX, crColor COLORREF, alpha byte) {
	xMenu_SetBorderColor.Call(
		uintptr(hMenu),
		uintptr(crColor),
		uintptr(alpha))
}

/*
获取左侧区域宽度.

参数:
	hMenu 菜单句柄.
返回:
	返回左侧区域宽度.
*/
func XMenuGetLeftWidth(hMenu HMENUX) int {
	ret, _, _ := xMenu_GetLeftWidth.Call(uintptr(hMenu))

	return int(ret)
}

/*
获取菜单项文本左间隔.

参数:
	hMenu 菜单句柄.
返回:
	返回菜单项文件左间隔大小.
*/
func XMenuGetLeftSpaceText(hMenu HMENUX) int {
	ret, _, _ := xMenu_GetLeftSpaceText.Call(uintptr(hMenu))

	return int(ret)
}

/*
获取菜单项数量,包含子菜单项.

参数:
	hMenu 菜单句柄.
返回:
	菜单项数量.
*/
func XMenuGetItemCount(hMenu HMENUX) int {
	ret, _, _ := xMenu_GetItemCount.Call(uintptr(hMenu))

	return int(ret)
}

/*
设置菜单项勾选状态.

参数:
	hMenu 菜单句柄.
	nID 菜单项ID
	bCheck 勾选TRUE,非勾选FALSE
返回:
	如果勾选返回TRUE,否则返回FALSE.
*/
func XMenuSetItemCheck(hMenu HMENUX, nID int, bCheck bool) {
	xMenu_SetItemCheck.Call(
		uintptr(hMenu),
		uintptr(nID),
		uintptr(BoolToBOOL(bCheck)))
}

/*
判断菜单项是否勾选.

参数:
	hMenu 菜单句柄.
	nID 菜单项ID
返回:
	如果勾选返回TRUE,否则返回FALSE.
*/
func XMenuIsItemCheck(hMenu HMENUX, nID int) bool {
	ret, _, _ := xMenu_IsItemCheck.Call(
		uintptr(hMenu),
		uintptr(nID))

	return ret == TRUE
}
