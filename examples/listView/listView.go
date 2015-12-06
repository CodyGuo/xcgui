package main

import (
	"fmt"
)

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

type template_info_i struct {
	info int
}

type listView_item_i struct {
	iGroup    int                    // 项所述组索引 -1没有组
	iItem     int                    // 项在数组中位置索引,如果此致为-1,那么为组
	nUserData int                    // 用户绑定数据
	nState    xcgui.List_item_state_ // 状态
	rcItem    xcgui.RECT             // 整个区域,包含边框
	hLayout   xcgui.HXCGUI           // 布局对象
	pInfo     *template_info_i       // 模板信息
}

func main() {
	hWindow := xcgui.XWndCreate(0, 0, 500, 400, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hListView := xcgui.XListViewCreate(20, 40, 450, 300, xcgui.HXCGUI(hWindow))
	xcgui.XListViewSetItemTemplateXML(hListView, "../xml-template/ListView_Item.xml")

	hAdapter := xcgui.XAdapterListViewCreate()
	xcgui.XAdapterListViewGroupAddColumn(hAdapter, "name3")
	xcgui.XAdapterListViewItemAddColumn(hAdapter, "name")
	xcgui.XAdapterListViewItemAddColumn(hAdapter, "name2")

	xcgui.XListViewBindAdapter(hListView, hAdapter)

	group1 := xcgui.XAdapterListViewGroupAddItemText(hAdapter, "group1")
	group2 := xcgui.XAdapterListViewGroupAddItemText(hAdapter, "group2")

	hImage := xcgui.XImageLoadFile("../img/comma_face_01.png", false)

	for i := 0; i < 20; i++ {
		index1 := xcgui.XAdapterListViewItemAddItemImage(hAdapter, group1, hImage)
		xcgui.XAdapterListViewItemSetText(hAdapter, group1, index1, 1, "group1-item-"+fmt.Sprint(i))

		index2 := xcgui.XAdapterListViewItemAddItemImage(hAdapter, group2, hImage)
		xcgui.XAdapterListViewItemSetText(hAdapter, group2, index2, 1, "group2-item-"+fmt.Sprint(i))
	}

	xcgui.XEleRegEventC(hListView, xcgui.XE_LISTVIEW_SELECT, xcgui.CallBack(OnListViewSelect))
	xcgui.XEleRegEventC(hListView, xcgui.XE_LISTVIEW_EXPAND, xcgui.CallBack(OnListViewExpand))

	xcgui.XEleRegEventC(hListView, xcgui.XE_LISTVIEW_TEMP_CREATE, xcgui.CallBack(OnTemplateCreate))
	xcgui.XEleRegEventC(hListView, xcgui.XE_LISTVIEW_TEMP_DESTROY, xcgui.CallBack(OnTemplateDestroy))
	xcgui.XEleRegEventC(hListView, xcgui.XE_LISTVIEW_TEMP_ADJUST_COORDINATE, xcgui.CallBack(OnTemplateAdjustCoordinate))

	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func OnListViewSelect(iGroup int, iItem int, pbHandled *bool) int {
	*pbHandled = true

	return 0
}

func OnListViewExpand(iGroup int, bExpand bool, pbHandled *bool) int {
	*pbHandled = true

	return 0
}

func OnTemplateCreate(pItem *listView_item_i, pbHandled *bool) int {
	*pbHandled = true

	return 0
}

func OnTemplateDestroy(pItem *listView_item_i, pbHandled *bool) int {
	*pbHandled = true

	return 0
}

func OnTemplateAdjustCoordinate(pItem *listView_item_i, pbHandled *bool) int {
	*pbHandled = true

	return 0
}
