package main

import (
	"fmt"
)

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

//模板项
type template_info_i struct {
	info int
}

type tree_item_i struct {
	nID        int
	nDepth     int
	nState     xcgui.Tree_item_state_
	nHeight    int
	nSelHeight int
	nUserData  int
	bExpand    bool
	rcItem     xcgui.RECT
	hLayout    xcgui.HXCGUI     ///<布局对象
	pInfo      *template_info_i ///<模板信息
}

func main() {
	hWindow := xcgui.XWndCreate(0, 0, 350, 400, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hTree := xcgui.XTreeCreate(20, 40, 300, 300, xcgui.HXCGUI(hWindow))
	xcgui.XTreeSetItemTemplateXML(hTree, "../xml-template/Tree_Item.xml")

	hAdapter := xcgui.XAdapterTreeCreate()
	xcgui.XTreeBindAdapter(hTree, hAdapter)
	xcgui.XAdapterTreeAddColumn(hAdapter, "name")

	for i := 0; i < 20; i++ {
		xcgui.XAdapterTreeInsertItemText(hAdapter, "name-"+fmt.Sprint(i)+"-0", xcgui.XC_ID_ROOT, xcgui.XC_ID_LAST)
	}

	xcgui.XAdapterTreeInsertItemText(hAdapter, "Item1-----1", 1, xcgui.XC_ID_LAST)         // idc1
	idc2 := xcgui.XAdapterTreeInsertItemText(hAdapter, "Item1-----2", 1, xcgui.XC_ID_LAST) // idc2
	xcgui.XAdapterTreeInsertItemText(hAdapter, "Item1-----2--1", idc2, xcgui.XC_ID_LAST)   // idc2-1

	xcgui.XEleRegEventC(hTree, xcgui.XE_TREE_SELECT, xcgui.CallBack(OnTreeSelect))
	xcgui.XEleRegEventC(hTree, xcgui.XE_TREE_EXPAND, xcgui.CallBack(OnTreeExpand))

	xcgui.XEleRegEventC(hTree, xcgui.XE_TREE_TEMP_CREATE, xcgui.CallBack(OnTemplateCreate))
	xcgui.XEleRegEventC(hTree, xcgui.XE_TREE_TEMP_DESTROY, xcgui.CallBack(OnTemplateDestroy))
	xcgui.XEleRegEventC(hTree, xcgui.XE_TREE_TEMP_ADJUST_COORDINATE, xcgui.CallBack(OnTemplateAdjustCoordinate))

	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func OnTreeSelect(nItemID int, pbHandled *bool) int {
	*pbHandled = true
	fmt.Println("选择的 ItemID 为: ", nItemID)

	return 0
}

func OnTreeExpand(id int, bExpand bool, pbHandled *bool) int {
	*pbHandled = true
	fmt.Printf("ID 为 %d, 是否展开: %v \n", id, bExpand)

	return 0
}

func OnTemplateCreate(pItem *tree_item_i, pbHandled *bool) int {
	*pbHandled = true
	fmt.Println("OnTemplateCreate -> tree_item_i: ", pItem.nID, pItem.nDepth, pItem.nHeight,
		pItem.nSelHeight, pItem.nUserData, pItem.bExpand, pItem.rcItem,
		pItem.nState, pItem.hLayout, pItem.pInfo.info)

	return 0
}

func OnTemplateDestroy(pItem *tree_item_i, pbHandled *bool) int {
	*pbHandled = true
	fmt.Println("OnTemplateDestroy -> tree_item_i: ", pItem.nID, pItem.nDepth, pItem.nHeight,
		pItem.nSelHeight, pItem.nUserData, pItem.bExpand, pItem.rcItem,
		pItem.nState, pItem.hLayout)

	return 0
}

func OnTemplateAdjustCoordinate(pItem *tree_item_i, pbHandled *bool) int {
	*pbHandled = true
	fmt.Println("OnTemplateAdjustCoordinate -> tree_item_i: ", pItem.nID, pItem.nDepth, pItem.nHeight,
		pItem.nSelHeight, pItem.nUserData, pItem.bExpand, pItem.rcItem,
		pItem.nState, pItem.hLayout)

	return 0
}
