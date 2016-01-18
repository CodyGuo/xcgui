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
	hWindow := xcgui.XWnd_Create(0, 0, 500, 400, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hListView := xcgui.XListView_Create(20, 40, 450, 300, xcgui.HXCGUI(hWindow))
	xcgui.XListView_SetItemTemplateXML(hListView, "../xml-template/ListView_Item.xml")

	hAdapter := xcgui.XAdapterListView_Create()
	xcgui.XAdapterListView_GroupAddColumn(hAdapter, "name3")
	xcgui.XAdapterListView_ItemAddColumn(hAdapter, "name")
	xcgui.XAdapterListView_ItemAddColumn(hAdapter, "name2")

	xcgui.XListView_BindAdapter(hListView, hAdapter)

	group1 := xcgui.XAdapterListView_GroupAddItemText(hAdapter, "group1")
	group2 := xcgui.XAdapterListView_GroupAddItemText(hAdapter, "group2")

	hImage := xcgui.XImage_LoadFile("../img/comma_face_01.png", false)

	for i := 0; i < 20; i++ {
		index1 := xcgui.XAdapterListView_ItemAddItemImage(hAdapter, group1, hImage)
		xcgui.XAdapterListView_ItemSetText(hAdapter, group1, index1, 1, "group1-item-"+fmt.Sprint(i))

		index2 := xcgui.XAdapterListView_ItemAddItemImage(hAdapter, group2, hImage)
		xcgui.XAdapterListView_ItemSetText(hAdapter, group2, index2, 1, "group2-item-"+fmt.Sprint(i))
	}

	xcgui.XEle_RegEventC(hListView, xcgui.XE_LISTVIEW_SELECT, xcgui.CallBack(OnListViewSelect))
	xcgui.XEle_RegEventC(hListView, xcgui.XE_LISTVIEW_EXPAND, xcgui.CallBack(OnListViewExpand))

	xcgui.XEle_RegEventC(hListView, xcgui.XE_LISTVIEW_TEMP_CREATE, xcgui.CallBack(OnTemplateCreate))
	xcgui.XEle_RegEventC(hListView, xcgui.XE_LISTVIEW_TEMP_DESTROY, xcgui.CallBack(OnTemplateDestroy))
	xcgui.XEle_RegEventC(hListView, xcgui.XE_LISTVIEW_TEMP_ADJUST_COORDINATE, xcgui.CallBack(OnTemplateAdjustCoordinate))

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
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
