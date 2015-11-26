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
	xList_Create = XCDLL.MustFindProc("XList_Create")
	xList_AddColumn = XCDLL.MustFindProc("XList_AddColumn")
	xList_InsertColumn = XCDLL.MustFindProc("XList_InsertColumn")
	xList_EnableMultiSel = XCDLL.MustFindProc("XList_EnableMultiSel")
	xList_SetDrawItemBkFlags = XCDLL.MustFindProc("XList_SetDrawItemBkFlags")
	xList_SetColumnWidth = XCDLL.MustFindProc("XList_SetColumnWidth")
	xList_SetColumnMinWidth = XCDLL.MustFindProc("XList_SetColumnMinWidth")
	xList_SetItemData = XCDLL.MustFindProc("XList_SetItemData")
	xList_GetItemData = XCDLL.MustFindProc("XList_GetItemData")
	xList_SetSelectItem = XCDLL.MustFindProc("XList_SetSelectItem")
	xList_GetSelectItem = XCDLL.MustFindProc("XList_GetSelectItem")
	xList_GetSelectItemCount = XCDLL.MustFindProc("XList_GetSelectItemCount")
	xList_SelectAll = XCDLL.MustFindProc("XList_SelectAll")
	xList_GetHeaderHELE = XCDLL.MustFindProc("XList_GetHeaderHELE")
	xList_DeleteColumn = XCDLL.MustFindProc("XList_DeleteColumn")
	xList_DeleteColumnAll = XCDLL.MustFindProc("XList_DeleteColumnAll")
	xList_BindAdapter = XCDLL.MustFindProc("XList_BindAdapter")
	xList_BindAdapterHeader = XCDLL.MustFindProc("XList_BindAdapterHeader")
	xList_GetAdapter = XCDLL.MustFindProc("XList_GetAdapter")
	xList_SetItemTemplateXML = XCDLL.MustFindProc("XList_SetItemTemplateXML")
	xList_SetItemTemplateXMLFromString = XCDLL.MustFindProc("XList_SetItemTemplateXMLFromString")
	xList_GetTemplateObject = XCDLL.MustFindProc("XList_GetTemplateObject")
	xList_GetItemIndexFromHXCGUI = XCDLL.MustFindProc("XList_GetItemIndexFromHXCGUI")
	xList_SetHeaderHeight = XCDLL.MustFindProc("XList_SetHeaderHeight")
	xList_GetHeaderHeight = XCDLL.MustFindProc("XList_GetHeaderHeight")
	xList_AddItemBkBorder = XCDLL.MustFindProc("XList_AddItemBkBorder")
	xList_AddItemBkFill = XCDLL.MustFindProc("XList_AddItemBkFill")
	xList_AddItemBkImage = XCDLL.MustFindProc("XList_AddItemBkImage")
	xList_GetItemBkInfoCount = XCDLL.MustFindProc("XList_GetItemBkInfoCount")
	xList_ClearItemBkInfo = XCDLL.MustFindProc("XList_ClearItemBkInfo")
	xList_GetItemBkInfoManager = XCDLL.MustFindProc("XList_GetItemBkInfoManager")
	xList_SetItemHeightDefault = XCDLL.MustFindProc("XList_SetItemHeightDefault")
	xList_GetItemHeightDefault = XCDLL.MustFindProc("XList_GetItemHeightDefault")
	xList_HitTest = XCDLL.MustFindProc("XList_HitTest")
	xList_HitTestOffset = XCDLL.MustFindProc("XList_HitTestOffset")
	xList_RefreshData = XCDLL.MustFindProc("XList_RefreshData")
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

func XListSetDrawItemBkFlags(hEle HELE, nFlags LIST_DRAWITEMBK_FLAGS_) {
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

func XListAddItemBkBorder(hEle HELE, nState LIST_ITEM_STATE_, color COLORREF, alpha byte, width int) {
	xList_AddItemBkBorder.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha),
		uintptr(width))
}

func XListAddItemBkFill(hEle HELE, nState LIST_ITEM_STATE_, color COLORREF, alpha byte) {
	xList_AddItemBkFill.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha))
}

func XListAddItemBkImage(hEle HELE, nState LIST_ITEM_STATE_, hImage HIMAGE) {
	xList_AddItemBkImage.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(hImage))
}

func XListGetItemBkInfoCount(hEle HELE, nState LIST_ITEM_STATE_) int {
	ret, _, _ := xList_GetItemBkInfoCount.Call(
		uintptr(hEle),
		uintptr(nState))

	return int(ret)

}

func XListClearItemBkInfo(hEle HELE, nState LIST_ITEM_STATE_) {
	xList_ClearItemBkInfo.Call(
		uintptr(hEle),
		uintptr(nState))

}

func XListGetItemBkInfoManager(hEle HELE, nState LIST_ITEM_STATE_) HBKINFOM {
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
