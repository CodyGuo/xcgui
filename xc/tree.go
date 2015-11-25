package xc

import (
	"syscall"
	"unsafe"
)

var (
	xTree_Create                          *syscall.Proc
	xTree_SetItemTemplateXML              *syscall.Proc
	xTree_SetItemTemplateXMLSel           *syscall.Proc
	xTree_SetItemTemplateXMLFromString    *syscall.Proc
	xTree_SetItemTemplateXMLSelFromString *syscall.Proc
	xTree_SetDrawItemBkFlags              *syscall.Proc
	xTree_SetItemData                     *syscall.Proc
	xTree_GetItemData                     *syscall.Proc
	xTree_SetSelectItem                   *syscall.Proc
	xTree_GetSelectItem                   *syscall.Proc
	xTree_IsExpand                        *syscall.Proc
	xTree_ExpandItem                      *syscall.Proc
	xTree_HitTest                         *syscall.Proc
	xTree_HitTestOffet                    *syscall.Proc
	xTree_GetFirstChildItem               *syscall.Proc
	xTree_GetNextSiblingItem              *syscall.Proc
	xTree_GetParentItem                   *syscall.Proc
	xTree_BindAdapter                     *syscall.Proc
	xTree_GetAdapter                      *syscall.Proc
	xTree_SetIndentation                  *syscall.Proc
	xTree_GetIndentation                  *syscall.Proc
	xTree_SetItemHeightDefault            *syscall.Proc
	xTree_GetItemHeightDefault            *syscall.Proc
	xTree_SetItemHeight                   *syscall.Proc
	xTree_GetItemHeight                   *syscall.Proc
	xTree_AddItemBkBorder                 *syscall.Proc
	xTree_AddItemBkFill                   *syscall.Proc
	xTree_AddItemBkImage                  *syscall.Proc
	xTree_GetItemBkInfoCount              *syscall.Proc
	xTree_ClearItemBkInfo                 *syscall.Proc
	xTree_GetItemBkInfoManager            *syscall.Proc
	xTree_GetTemplateObject               *syscall.Proc
	xTree_GetItemIDFromHXCGUI             *syscall.Proc
)

func init() {
	xTree_Create = XCDLL.MustFindProc("XTree_Create")
	xTree_SetItemTemplateXML = XCDLL.MustFindProc("XTree_SetItemTemplateXML")
	xTree_SetItemTemplateXMLSel = XCDLL.MustFindProc("XTree_SetItemTemplateXMLSel")
	xTree_SetItemTemplateXMLFromString = XCDLL.MustFindProc("XTree_SetItemTemplateXMLFromString")
	xTree_SetItemTemplateXMLSelFromString = XCDLL.MustFindProc("XTree_SetItemTemplateXMLSelFromString")
	xTree_SetDrawItemBkFlags = XCDLL.MustFindProc("XTree_SetDrawItemBkFlags")
	xTree_SetItemData = XCDLL.MustFindProc("XTree_SetItemData")

	xTree_GetItemData = XCDLL.MustFindProc("XTree_GetItemData")
	xTree_SetSelectItem = XCDLL.MustFindProc("XTree_SetSelectItem")
	xTree_GetSelectItem = XCDLL.MustFindProc("XTree_GetSelectItem")
	xTree_IsExpand = XCDLL.MustFindProc("XTree_IsExpand")
	xTree_ExpandItem = XCDLL.MustFindProc("XTree_ExpandItem")
	xTree_HitTest = XCDLL.MustFindProc("XTree_HitTest")
	xTree_HitTestOffet = XCDLL.MustFindProc("XTree_HitTestOffet")
	xTree_GetFirstChildItem = XCDLL.MustFindProc("XTree_GetFirstChildItem")
	xTree_GetNextSiblingItem = XCDLL.MustFindProc("XTree_GetNextSiblingItem")
	xTree_GetParentItem = XCDLL.MustFindProc("XTree_GetParentItem")
	xTree_BindAdapter = XCDLL.MustFindProc("XTree_BindAdapter")
	xTree_GetAdapter = XCDLL.MustFindProc("XTree_GetAdapter")
	xTree_SetIndentation = XCDLL.MustFindProc("XTree_SetIndentation")
	xTree_GetIndentation = XCDLL.MustFindProc("XTree_GetIndentation")
	xTree_SetItemHeightDefault = XCDLL.MustFindProc("XTree_SetItemHeightDefault")
	xTree_GetItemHeightDefault = XCDLL.MustFindProc("XTree_GetItemHeightDefault")
	xTree_SetItemHeight = XCDLL.MustFindProc("XTree_SetItemHeight")
	xTree_GetItemHeight = XCDLL.MustFindProc("XTree_GetItemHeight")
	xTree_AddItemBkBorder = XCDLL.MustFindProc("XTree_AddItemBkBorder")
	xTree_AddItemBkFill = XCDLL.MustFindProc("XTree_AddItemBkFill")
	xTree_AddItemBkImage = XCDLL.MustFindProc("XTree_AddItemBkImage")
	xTree_GetItemBkInfoCount = XCDLL.MustFindProc("XTree_GetItemBkInfoCount")
	xTree_ClearItemBkInfo = XCDLL.MustFindProc("XTree_ClearItemBkInfo")
	xTree_GetItemBkInfoManager = XCDLL.MustFindProc("XTree_GetItemBkInfoManager")
	xTree_GetTemplateObject = XCDLL.MustFindProc("XTree_GetTemplateObject")
	xTree_GetItemIDFromHXCGUI = XCDLL.MustFindProc("XTree_GetItemIDFromHXCGUI")

}

func XTreeCreate(x int, y int, cx int, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xTree_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

func XTreeSetItemTemplateXML(hEle HELE, pXmlFile string) bool {
	ret, _, _ := xTree_SetItemTemplateXML.Call(
		uintptr(hEle),
		StringToUintPtr(pXmlFile))
	return ret == TRUE
}

func XTreeSetItemTemplateXMLSel(hEle HELE, pXmlFile string) bool {
	ret, _, _ := xTree_SetItemTemplateXMLSel.Call(
		uintptr(hEle),
		StringToUintPtr(pXmlFile))
	return ret == TRUE
}

func XTreeSetItemTemplateXMLFromString(hEle HELE, pXmlFile string) bool {
	ret, _, _ := xTree_SetItemTemplateXMLFromString.Call(
		uintptr(hEle),
		StringToUintPtr(pXmlFile))
	return ret == TRUE
}

func XTreeSetItemTemplateXMLSelFromString(hEle HELE, pXmlFile string) bool {
	ret, _, _ := xTree_SetItemTemplateXMLSelFromString.Call(
		uintptr(hEle),
		StringToUintPtr(pXmlFile))
	return ret == TRUE
}

func XTreeSetDrawItemBkFlags(hEle HELE, nFlags int) {
	xTree_SetDrawItemBkFlags.Call(uintptr(hEle), uintptr(nFlags))
}

func XTreeSetItemData(hEle HELE, nID int, nUserData int) bool {
	ret, _, _ := xTree_SetItemData.Call(
		uintptr(hEle),
		uintptr(nID),
		uintptr(nUserData))
	return ret == TRUE

}

func XTreeGetItemData(hEle HELE, nID int) int {
	ret, _, _ := xTree_GetItemData.Call(uintptr(hEle),
		uintptr(nID))
	return int(ret)
}

func XTreeSetSelectItem(hEle HELE, nID int) bool {
	ret, _, _ := xTree_SetSelectItem.Call(uintptr(hEle),
		uintptr(nID))
	return ret == TRUE
}

func XTreeGetSelectItem(hEle HELE) int {
	ret, _, _ := xTree_GetSelectItem.Call(uintptr(hEle))
	return int(ret)
}

func XTreeIsExpand(hEle HELE, nID int) bool {
	ret, _, _ := xTree_IsExpand.Call(uintptr(hEle), uintptr(nID))
	return ret == TRUE
}

func XTreeExpandItem(hEle HELE, nID int, bExpand bool) bool {
	ret, _, _ := xTree_ExpandItem.Call(uintptr(hEle), uintptr(nID), uintptr(BoolToBOOL(bExpand)))
	return ret == TRUE
}

func XTreeHitTest(hEle HELE, pPt *POINT) int {
	ret, _, _ := xTree_HitTest.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)))
	return int(ret)
}

func XTreeHitTestOffet(hEle HELE, pPt *POINT) int {
	ret, _, _ := xTree_HitTestOffet.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)))
	return int(ret)
}

func XTreeGetFirstChildItem(hEle HELE, nID int) int {
	ret, _, _ := xTree_GetFirstChildItem.Call(uintptr(hEle), uintptr(nID))
	return int(ret)
}

func XTreeGetNextSiblingItem(hEle HELE, nID int) int {
	ret, _, _ := xTree_GetNextSiblingItem.Call(uintptr(hEle), uintptr(nID))
	return int(ret)
}

func XTreeGetParentItem(hEle HELE, nID int) int {
	ret, _, _ := xTree_GetParentItem.Call(uintptr(hEle), uintptr(nID))
	return int(ret)
}

func XTreeBindAdapter(hEle HELE, hAdapter HXCGUI) {
	xTree_BindAdapter.Call(
		uintptr(hEle), uintptr(hAdapter))
}

func XTreeGetAdapter(hEle HELE) HXCGUI {
	ret, _, _ := xTree_GetAdapter.Call(uintptr(hEle))
	return HXCGUI(ret)
}

func XTreeSetIndentation(hEle HELE, nWidth int) {
	xTree_SetIndentation.Call(uintptr(hEle), uintptr(nWidth))
}

func XTreeGetIndentation(hEle HELE) int {
	ret, _, _ := xTree_GetIndentation.Call(uintptr(hEle))
	return int(ret)
}

func XTreeSetItemHeightDefault(hEle HELE, nHeight int, nSelHeight int) {
	xTree_SetItemHeightDefault.Call(uintptr(hEle), uintptr(nHeight), uintptr(nSelHeight))
}

func XTreeGetItemHeightDefault(hEle HELE, pHeight *int, pSelHeight *int) {
	xTree_GetItemHeightDefault.Call(uintptr(hEle), uintptr(unsafe.Pointer(pHeight)), uintptr(unsafe.Pointer(pSelHeight)))
}

func XTreeSetItemHeight(hEle HELE, nID int, nHeight int, nSelHeight int) {
	xTree_SetItemHeight.Call(uintptr(hEle), uintptr(nHeight), uintptr(nSelHeight))
}

func XTreeGetItemHeight(hEle HELE, nID int, pHeight *int, pSelHeight *int) {
	xTree_GetItemHeight.Call(uintptr(hEle), uintptr(nID), uintptr(unsafe.Pointer(pHeight)), uintptr(unsafe.Pointer(pSelHeight)))

}

func XTreeAddItemBkBorder(hEle HELE, nState tree_item_state_, color COLORREF, alpha byte, width int) {
	xTree_AddItemBkBorder.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha),
		uintptr(width))
}

func XTreeAddItemBkFill(hEle HELE, nState tree_item_state_, color COLORREF, alpha byte) {
	xTree_AddItemBkFill.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha))
}

func XTreeAddItemBkImage(hEle HELE, nState tree_item_state_, hImage HIMAGE) {
	xTree_AddItemBkImage.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(hImage))

}

func XTreeGetItemBkInfoCount(hEle HELE, nState tree_item_state_) int {
	ret, _, _ := xTree_GetItemBkInfoCount.Call(
		uintptr(hEle),
		uintptr(nState))
	return int(ret)
}

func XTreeClearItemBkInfo(hEle HELE, nState tree_item_state_) {
	xTree_ClearItemBkInfo.Call(
		uintptr(hEle),
		uintptr(nState))
}

func XTreeGetItemBkInfoManager(hEle HELE, nState tree_item_state_) HBKINFOM {
	ret, _, _ := xTree_GetItemBkInfoManager.Call(
		uintptr(hEle),
		uintptr(nState))
	return HBKINFOM(ret)
}

func XTreeGetTemplateObject(hEle HELE, nID int, nTempItemID int) HXCGUI {
	ret, _, _ := xTree_GetTemplateObject.Call(
		uintptr(hEle),
		uintptr(nID),
		uintptr(nTempItemID))
	return HXCGUI(ret)
}

func XTreeGetItemIDFromHXCGUI(hEle HELE, hXCGUI HXCGUI) int {
	ret, _, _ := xTree_GetItemIDFromHXCGUI.Call(
		uintptr(hEle),
		uintptr(hXCGUI))
	return int(ret)
}
