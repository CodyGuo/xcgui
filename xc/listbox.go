package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xListBox_Create                       *syscall.Proc
	xListBox_SetDrawItemBkFlags           *syscall.Proc
	xListBox_SetItemData                  *syscall.Proc
	xListBox_GetItemData                  *syscall.Proc
	xListBox_AddItemBkBorder              *syscall.Proc
	xListBox_AddItemBkFill                *syscall.Proc
	xListBox_AddItemBkImage               *syscall.Proc
	xListBox_GetItemBkInfoCount           *syscall.Proc
	xListBox_ClearItemBkInfo              *syscall.Proc
	xListBox_GetItemBkInfoManager         *syscall.Proc
	xListBox_SetItemInfo                  *syscall.Proc
	xListBox_GetItemInfo                  *syscall.Proc
	xListBox_SetSelectItem                *syscall.Proc
	xListBox_GetSelectItem                *syscall.Proc
	xListBox_CancelSelectItem             *syscall.Proc
	xListBox_CancelSelectAll              *syscall.Proc
	xListBox_GetSelectAll                 *syscall.Proc
	xListBox_GetSelectCount               *syscall.Proc
	xListBox_GetItemMouseStay             *syscall.Proc
	xListBox_SelectAll                    *syscall.Proc
	xListBox_SetItemHeightDefault         *syscall.Proc
	xListBox_GetItemHeightDefault         *syscall.Proc
	xListBox_GetItemIndexFromHXCGUI       *syscall.Proc
	xListBox_HitTest                      *syscall.Proc
	xListBox_HitTestOffset                *syscall.Proc
	xListBox_SetItemTemplateXML           *syscall.Proc
	xListBox_SetItemTemplateXMLFromString *syscall.Proc
	xListBox_GetTemplateObject            *syscall.Proc
	xListBox_EnableMultiSel               *syscall.Proc
	xListBox_BindAdapter                  *syscall.Proc
	xListBox_GetAdapter                   *syscall.Proc
)

func init() {
	// Functions
	xListBox_Create = xcDLL.MustFindProc("XListBox_Create")
	xListBox_SetDrawItemBkFlags = xcDLL.MustFindProc("XListBox_SetDrawItemBkFlags")
	xListBox_SetItemData = xcDLL.MustFindProc("XListBox_SetItemData")
	xListBox_GetItemData = xcDLL.MustFindProc("XListBox_GetItemData")
	xListBox_AddItemBkBorder = xcDLL.MustFindProc("XListBox_AddItemBkBorder")
	xListBox_AddItemBkFill = xcDLL.MustFindProc("XListBox_AddItemBkFill")
	xListBox_AddItemBkImage = xcDLL.MustFindProc("XListBox_AddItemBkImage")
	xListBox_GetItemBkInfoCount = xcDLL.MustFindProc("XListBox_GetItemBkInfoCount")
	xListBox_ClearItemBkInfo = xcDLL.MustFindProc("XListBox_ClearItemBkInfo")
	xListBox_GetItemBkInfoManager = xcDLL.MustFindProc("XListBox_GetItemBkInfoManager")
	xListBox_SetItemInfo = xcDLL.MustFindProc("XListBox_SetItemInfo")
	xListBox_GetItemInfo = xcDLL.MustFindProc("XListBox_GetItemInfo")
	xListBox_SetSelectItem = xcDLL.MustFindProc("XListBox_SetSelectItem")
	xListBox_GetSelectItem = xcDLL.MustFindProc("XListBox_GetSelectItem")
	xListBox_CancelSelectItem = xcDLL.MustFindProc("XListBox_CancelSelectItem")
	xListBox_CancelSelectAll = xcDLL.MustFindProc("XListBox_CancelSelectAll")
	xListBox_GetSelectAll = xcDLL.MustFindProc("XListBox_GetSelectAll")
	xListBox_GetSelectCount = xcDLL.MustFindProc("XListBox_GetSelectCount")
	xListBox_GetItemMouseStay = xcDLL.MustFindProc("XListBox_GetItemMouseStay")
	xListBox_SelectAll = xcDLL.MustFindProc("XListBox_SelectAll")
	xListBox_SetItemHeightDefault = xcDLL.MustFindProc("XListBox_SetItemHeightDefault")
	xListBox_GetItemHeightDefault = xcDLL.MustFindProc("XListBox_GetItemHeightDefault")
	xListBox_GetItemIndexFromHXCGUI = xcDLL.MustFindProc("XListBox_GetItemIndexFromHXCGUI")
	xListBox_HitTest = xcDLL.MustFindProc("XListBox_HitTest")
	xListBox_HitTestOffset = xcDLL.MustFindProc("XListBox_HitTestOffset")
	xListBox_SetItemTemplateXML = xcDLL.MustFindProc("XListBox_SetItemTemplateXML")
	xListBox_SetItemTemplateXMLFromString = xcDLL.MustFindProc("XListBox_SetItemTemplateXMLFromString")
	xListBox_GetTemplateObject = xcDLL.MustFindProc("XListBox_GetTemplateObject")
	xListBox_EnableMultiSel = xcDLL.MustFindProc("XListBox_EnableMultiSel")
	xListBox_BindAdapter = xcDLL.MustFindProc("XListBox_BindAdapter")
	xListBox_GetAdapter = xcDLL.MustFindProc("XListBox_GetAdapter")
}

func XListBoxCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xListBox_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

func XListBoxSetDrawItemBkFlags(hEle HELE, nFlags LIST_DRAWITEMBK_FLAGS_) {
	xListBox_SetDrawItemBkFlags.Call(
		uintptr(hEle),
		uintptr(nFlags))
}

func XListBoxSetItemData(hEle HELE, iItem, nUserData int) bool {
	ret, _, _ := xListBox_SetItemData.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(nUserData))

	if ret != TRUE {
		return false
	}

	return true
}

func XListBoxGetItemData(hEle HELE, iItem int) int {
	ret, _, _ := xListBox_GetItemData.Call(
		uintptr(hEle),
		uintptr(iItem))

	return int(ret)
}

func XListBoxAddItemBkBorder(hEle HELE, nState LIST_ITEM_STATE_, color COLORREF, alpha byte, width int) {
	xListBox_AddItemBkBorder.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha),
		uintptr(width))
}

func XListBoxAddItemBkFill(hEle HELE, nState LIST_ITEM_STATE_, color COLORREF, alpha byte) {
	xListBox_AddItemBkFill.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha))
}

func XListBoxAddItemBkImage(hEle HELE, nState LIST_ITEM_STATE_, hImage HIMAGE) {
	xListBox_AddItemBkImage.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(hImage))
}

func XListBoxGetItemBkInfoCount(hEle HELE, nState LIST_ITEM_STATE_) int {
	ret, _, _ := xListBox_GetItemBkInfoCount.Call(
		uintptr(hEle),
		uintptr(nState))

	return int(ret)
}

func XListBoxClearItemBkInfo(hEle HELE, nState LIST_ITEM_STATE_) {
	xListBox_ClearItemBkInfo.Call(
		uintptr(hEle),
		uintptr(nState))
}

func XListBoxGetItemBkInfoManager(hEle HELE, nState LIST_ITEM_STATE_) HBKINFOM {
	ret, _, _ := xListBox_GetItemBkInfoManager.Call(
		uintptr(hEle),
		uintptr(nState))

	return HBKINFOM(ret)
}

func XListBoxSetItemInfo(hEle HELE, iItemint int, pItem *LISTBOX_ITEM_INFO_I) bool {
	ret, _, _ := xListBox_SetItemInfo.Call(
		uintptr(hEle),
		uintptr(iItemint),
		uintptr(unsafe.Pointer(pItem)))

	if ret != TRUE {
		return false
	}

	return true
}

func XListBoxGetItemInfo(hEle HELE, iItem int, pItem *LISTBOX_ITEM_INFO_I) bool {
	ret, _, _ := xListBox_GetItemInfo.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pItem)))

	if ret != TRUE {
		return false
	}

	return true
}

func XListBoxSetSelectItem(hEle HELE, iItem int) bool {
	ret, _, _ := xListBox_SetSelectItem.Call(
		uintptr(hEle),
		uintptr(iItem))

	if ret != TRUE {
		return false
	}

	return true
}

func XListBoxGetSelectItem(hEle HELE) int {
	ret, _, _ := xListBox_GetSelectItem.Call(uintptr(hEle))

	return int(ret)
}

func XListBoxCancelSelectItem(hEle HELE, iItem int) bool {
	ret, _, _ := xListBox_CancelSelectItem.Call(
		uintptr(hEle),
		uintptr(iItem))

	if ret != TRUE {
		return false
	}

	return true
}

func XListBoxCancelSelectAll(hEle HELE) bool {
	ret, _, _ := xListBox_CancelSelectAll.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

func XListBoxGetSelectAll(hEle HELE, pArray *uint16, nArraySize int) int {
	ret, _, _ := xListBox_GetSelectAll.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pArray)),
		uintptr(nArraySize))

	return int(ret)
}

func XListBoxGetSelectCount(hEle HELE) int {
	ret, _, _ := xListBox_GetSelectCount.Call(uintptr(hEle))

	return int(ret)
}

func XListBoxGetItemMouseStay(hEle HELE) int {
	ret, _, _ := xListBox_GetItemMouseStay.Call(uintptr(hEle))

	return int(ret)
}

func XListBoxSelectAll(hEle HELE) bool {
	ret, _, _ := xListBox_SelectAll.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

func XListBoxSetItemHeightDefault(hEle HELE, nHeight, nSelHeight int) {
	xListBox_SetItemHeightDefault.Call(
		uintptr(hEle),
		uintptr(nHeight),
		uintptr(nSelHeight))
}

func XListBoxGetItemHeightDefault(hEle HELE, pHeight, pSelHeight *uint16) {
	xListBox_GetItemHeightDefault.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pHeight)),
		uintptr(unsafe.Pointer(pSelHeight)))
}

func XListBoxGetItemIndexFromHXCGUI(hEle HELE, hXCGUI HXCGUI) int {
	ret, _, _ := xListBox_GetItemIndexFromHXCGUI.Call(
		uintptr(hEle),
		uintptr(hXCGUI))

	return int(ret)
}

func XListBoxHitTest(hEle HELE, pPt *POINT) int {
	ret, _, _ := xListBox_HitTest.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)))

	return int(ret)
}

func XListBoxHitTestOffset(hEle HELE, pPt *POINT) int {
	ret, _, _ := xListBox_HitTestOffset.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)))

	return int(ret)
}

func XListBoxSetItemTemplateXML(hEle HELE, pXmlFile string) bool {
	ret, _, _ := xListBox_SetItemTemplateXML.Call(
		uintptr(hEle),
		StringToUintPtr(pXmlFile))

	if ret != TRUE {
		return false
	}

	return true
}

func XListBoxSetItemTemplateXMLFromString(hEle HELE, pStringXML string) bool {
	ret, _, _ := xListBox_SetItemTemplateXMLFromString.Call(
		uintptr(hEle),
		StringToUintPtr(pStringXML))

	if ret != TRUE {
		return false
	}

	return true
}

func XListBoxGetTemplateObject(hEle HELE, iItem, nTempItemID int) HXCGUI {
	ret, _, _ := xListBox_GetTemplateObject.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(nTempItemID))

	return HXCGUI(ret)
}

func XListBox_EnableMultiSel(hEle HELE, bEnable bool) {
	xListBox_EnableMultiSel.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

func XListBox_BindAdapter(hEle HELE, hAdapter HXCGUI) {
	xListBox_BindAdapter.Call(
		uintptr(hEle),
		uintptr(hAdapter))
}

func XListBox_GetAdapter(hEle HELE) HXCGUI {
	ret, _, _ := xListBox_GetAdapter.Call(uintptr(hEle))

	return HXCGUI(ret)
}
