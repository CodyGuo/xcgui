package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xList_Create                       *syscall.Proc
	xList_AddColumn                    *syscall.Proc
	xList_InsertColumn                 *syscall.Proc
	xList_EnableMultiSel               *syscall.Proc
	xList_SetDrawItemBkFlags           *syscall.Proc
	xList_SetColumnWidth               *syscall.Proc
	xList_SetColumnMinWidth            *syscall.Proc
	xList_SetItemData                  *syscall.Proc
	xList_GetItemData                  *syscall.Proc
	xList_SetSelectItem                *syscall.Proc
	xList_GetSelectItem                *syscall.Proc
	xList_GetSelectItemCount           *syscall.Proc
	xList_SelectAll                    *syscall.Proc
	xList_GetHeaderHELE                *syscall.Proc
	xList_DeleteColumn                 *syscall.Proc
	xList_DeleteColumnAll              *syscall.Proc
	xList_BindAdapter                  *syscall.Proc
	xList_BindAdapterHeader            *syscall.Proc
	xList_GetAdapter                   *syscall.Proc
	xList_SetItemTemplateXML           *syscall.Proc
	xList_SetItemTemplateXMLFromString *syscall.Proc
	xList_GetTemplateObject            *syscall.Proc
	xList_GetItemIndexFromHXCGUI       *syscall.Proc
	xList_SetHeaderHeight              *syscall.Proc
	xList_GetHeaderHeight              *syscall.Proc
	xList_AddItemBkBorder              *syscall.Proc
	xList_AddItemBkFill                *syscall.Proc
	xList_AddItemBkImage               *syscall.Proc
	xList_GetItemBkInfoCount           *syscall.Proc
	xList_ClearItemBkInfo              *syscall.Proc
	xList_GetItemBkInfoManager         *syscall.Proc
	xList_SetItemHeightDefault         *syscall.Proc
	xList_GetItemHeightDefault         *syscall.Proc
	xList_HitTest                      *syscall.Proc
	xList_HitTestOffset                *syscall.Proc
	xList_RefreshData                  *syscall.Proc
)

func init() {
	// Functions
	xList_Create = xcDLL.MustFindProc("XList_Create")
	xList_AddColumn = xcDLL.MustFindProc("XList_AddColumn")
	xList_InsertColumn = xcDLL.MustFindProc("XList_InsertColumn")
	xList_EnableMultiSel = xcDLL.MustFindProc("XList_EnableMultiSel")
	xList_SetDrawItemBkFlags = xcDLL.MustFindProc("XList_SetDrawItemBkFlags")
	xList_SetColumnWidth = xcDLL.MustFindProc("XList_SetColumnWidth")
	xList_SetColumnMinWidth = xcDLL.MustFindProc("XList_SetColumnMinWidth")
	xList_SetItemData = xcDLL.MustFindProc("XList_SetItemData")
	xList_GetItemData = xcDLL.MustFindProc("XList_GetItemData")
	xList_SetSelectItem = xcDLL.MustFindProc("XList_SetSelectItem")
	xList_GetSelectItem = xcDLL.MustFindProc("XList_GetSelectItem")
	xList_GetSelectItemCount = xcDLL.MustFindProc("XList_GetSelectItemCount")
	xList_SelectAll = xcDLL.MustFindProc("XList_SelectAll")
	xList_GetHeaderHELE = xcDLL.MustFindProc("XList_GetHeaderHELE")
	xList_DeleteColumn = xcDLL.MustFindProc("XList_DeleteColumn")
	xList_DeleteColumnAll = xcDLL.MustFindProc("XList_DeleteColumnAll")
	xList_BindAdapter = xcDLL.MustFindProc("XList_BindAdapter")
	xList_BindAdapterHeader = xcDLL.MustFindProc("XList_BindAdapterHeader")
	xList_GetAdapter = xcDLL.MustFindProc("XList_GetAdapter")
	xList_SetItemTemplateXML = xcDLL.MustFindProc("XList_SetItemTemplateXML")
	xList_SetItemTemplateXMLFromString = xcDLL.MustFindProc("XList_SetItemTemplateXMLFromString")
	xList_GetTemplateObject = xcDLL.MustFindProc("XList_GetTemplateObject")
	xList_GetItemIndexFromHXCGUI = xcDLL.MustFindProc("XList_GetItemIndexFromHXCGUI")
	xList_SetHeaderHeight = xcDLL.MustFindProc("XList_SetHeaderHeight")
	xList_GetHeaderHeight = xcDLL.MustFindProc("XList_GetHeaderHeight")
	xList_AddItemBkBorder = xcDLL.MustFindProc("XList_AddItemBkBorder")
	xList_AddItemBkFill = xcDLL.MustFindProc("XList_AddItemBkFill")
	xList_AddItemBkImage = xcDLL.MustFindProc("XList_AddItemBkImage")
	xList_GetItemBkInfoCount = xcDLL.MustFindProc("XList_GetItemBkInfoCount")
	xList_ClearItemBkInfo = xcDLL.MustFindProc("XList_ClearItemBkInfo")
	xList_GetItemBkInfoManager = xcDLL.MustFindProc("XList_GetItemBkInfoManager")
	xList_SetItemHeightDefault = xcDLL.MustFindProc("XList_SetItemHeightDefault")
	xList_GetItemHeightDefault = xcDLL.MustFindProc("XList_GetItemHeightDefault")
	xList_HitTest = xcDLL.MustFindProc("XList_HitTest")
	xList_HitTestOffset = xcDLL.MustFindProc("XList_HitTestOffset")
	xList_RefreshData = xcDLL.MustFindProc("XList_RefreshData")
}

func XListCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xList_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

func XListAddColumn(hEle HELE, width int) int {
	ret, _, _ := xList_AddColumn.Call(
		uintptr(hEle),
		uintptr(width))

	return int(ret)
}

func XListInsertColumn(hEle HELE, width, iItem int) int {
	ret, _, _ := xList_InsertColumn.Call(
		uintptr(hEle),
		uintptr(width),
		uintptr(iItem))

	return int(ret)
}

func XListEnableMultiSel(hEle HELE, bEnable bool) {
	xList_EnableMultiSel.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

func XListSetDrawItemBkFlags(hEle HELE, nFlags List_drawitembk_flags_) {
	xList_SetDrawItemBkFlags.Call(
		uintptr(hEle),
		uintptr(nFlags))
}

func XListSetColumnWidth(hEle HELE, iItem, width int) {
	xList_SetColumnWidth.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(width))
}

func XListSetColumnMinWidth(hEle HELE, iItem, width int) {
	xList_SetColumnMinWidth.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(width))
}

func XListSetItemData(hEle HELE, iItem, iSubItem, data int) bool {
	ret, _, _ := xList_SetItemData.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(iSubItem),
		uintptr(data))

	if ret != TRUE {
		return false
	}

	return true
}

func XListGetItemData(hEle HELE, iItem, iSubItem int) int {
	ret, _, _ := xList_GetItemData.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(iSubItem))

	return int(ret)
}

func XListSetSelectItem(hEle HELE, iItem int, bSelect bool) bool {
	ret, _, _ := xList_SetSelectItem.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(BoolToBOOL(bSelect)))

	if ret != TRUE {
		return false
	}

	return true
}

func XListGetSelectItem(hEle HELE) int {
	ret, _, _ := xList_GetSelectItem.Call(uintptr(hEle))

	return int(ret)
}

func XListGetSelectItemCount(hEle HELE) int {
	ret, _, _ := xList_GetSelectItemCount.Call(uintptr(hEle))

	return int(ret)
}

func XListSelectAll(hEle HELE) {
	xList_SelectAll.Call(uintptr(hEle))
}

func XListGetHeaderHELE(hEle HELE) HELE {
	ret, _, _ := xList_GetHeaderHELE.Call(uintptr(hEle))

	return HELE(ret)
}

func XListDeleteColumn(hEle HELE, iItem int) bool {
	ret, _, _ := xList_DeleteColumn.Call(
		uintptr(hEle),
		uintptr(iItem))

	if ret != TRUE {
		return false
	}

	return true
}

func XListDeleteColumnAll(hEle HELE) {
	xList_DeleteColumnAll.Call(uintptr(hEle))
}

func XListBindAdapter(hEle HELE, hAdapter HXCGUI) {
	xList_BindAdapter.Call(
		uintptr(hEle),
		uintptr(hAdapter))
}

func XListBindAdapterHeader(hEle HELE, hAdapter HXCGUI) {
	xList_BindAdapterHeader.Call(
		uintptr(hEle),
		uintptr(hAdapter))
}

func XListGetAdapter(hEle HELE) HELE {
	ret, _, _ := xList_GetAdapter.Call(uintptr(hEle))

	return HELE(ret)
}

func XListSetItemTemplateXML(hEle HELE, pXmlFile string) bool {
	ret, _, _ := xList_SetItemTemplateXML.Call(
		uintptr(hEle),
		StringToUintPtr(pXmlFile))

	if ret != TRUE {
		return false
	}

	return true
}

func XListSetItemTemplateXMLFromString(hEle HELE, pStringXML string) bool {
	ret, _, _ := xList_SetItemTemplateXMLFromString.Call(
		uintptr(hEle),
		StringToUintPtr(pStringXML))

	if ret != TRUE {
		return false
	}

	return true
}

func XListGetTemplateObject(hEle HELE, iItem, nTempItemID int) HXCGUI {
	ret, _, _ := xList_GetTemplateObject.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(nTempItemID))

	return HXCGUI(ret)
}

func XListGetItemIndexFromHXCGUI(hEle HELE, hXCGUI HXCGUI) int {
	ret, _, _ := xList_GetItemIndexFromHXCGUI.Call(
		uintptr(hEle),
		uintptr(hXCGUI))

	return int(ret)
}

func XListSetHeaderHeight(hEle HELE, height int) {
	xList_SetHeaderHeight.Call(
		uintptr(hEle),
		uintptr(height))
}

func XListGetHeaderHeight(hEle HELE) int {
	ret, _, _ := xList_GetHeaderHeight.Call(uintptr(hEle))

	return int(ret)
}

func XListAddItemBkBorder(hEle HELE, nState List_item_state_, color COLORREF, alpha byte, width int) {
	xList_AddItemBkBorder.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha),
		uintptr(width))
}

func XListAddItemBkFill(hEle HELE, nState List_item_state_, color COLORREF, alpha byte) {
	xList_AddItemBkFill.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha))
}

func XListAddItemBkImage(hEle HELE, nState List_item_state_, hImage HIMAGE) {
	xList_AddItemBkImage.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(hImage))
}

func XListGetItemBkInfoCount(hEle HELE, nState List_item_state_) int {
	ret, _, _ := xList_GetItemBkInfoCount.Call(
		uintptr(hEle),
		uintptr(nState))

	return int(ret)

}

func XListClearItemBkInfo(hEle HELE, nState List_item_state_) {
	xList_ClearItemBkInfo.Call(
		uintptr(hEle),
		uintptr(nState))

}

func XListGetItemBkInfoManager(hEle HELE, nState List_item_state_) HBKINFOM {
	ret, _, _ := xList_GetItemBkInfoManager.Call(
		uintptr(hEle),
		uintptr(nState))

	return HBKINFOM(ret)
}

func XListSetItemHeightDefault(hEle HELE, nHeight, nSelHeight int) {
	xList_SetItemHeightDefault.Call(
		uintptr(hEle),
		uintptr(nHeight),
		uintptr(nSelHeight))
}

func XListGetItemHeightDefault(hEle HELE, pHeight, pSelHeight *uint16) {
	xList_GetItemHeightDefault.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pHeight)),
		uintptr(unsafe.Pointer(pSelHeight)))
}

func XListHitTest(hEle HELE, pPt *POINT, piItem, piSubItem *uint16) bool {
	ret, _, _ := xList_HitTest.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)),
		uintptr(unsafe.Pointer(piItem)),
		uintptr(unsafe.Pointer(piSubItem)))

	if ret != TRUE {
		return false
	}

	return true
}

func XListHitTestOffset(hEle HELE, pPt *POINT, piItem, piSubItem *uint16) bool {
	ret, _, _ := xList_HitTestOffset.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)),
		uintptr(unsafe.Pointer(piItem)),
		uintptr(unsafe.Pointer(piSubItem)))

	if ret != TRUE {
		return false
	}

	return true
}

func XListRefreshData(hEle HELE) {
	xList_RefreshData.Call(uintptr(hEle))
}
