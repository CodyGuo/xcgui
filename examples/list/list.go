package main

import (
	"fmt"
)

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWnd_Create(0, 0, 550, 380, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hList := xcgui.XList_Create(20, 40, 500, 300, xcgui.HXCGUI(hWindow))
	xcgui.XList_SetItemTemplateXML(hList, "../xml-template/List_Item.xml")
	xcgui.XSView_SetLineSize(hList, 20, 20)

	for i := 0; i < 3; i++ {
		xcgui.XList_AddColumn(hList, 100)
	}

	hAdapterHeader := xcgui.XAdapterMap_Create()
	xcgui.XList_BindAdapterHeader(hList, hAdapterHeader)

	xcgui.XAdapterMap_AddItemText(hAdapterHeader, "name", "aaa")
	xcgui.XAdapterMap_AddItemText(hAdapterHeader, "name2", "bbb")
	xcgui.XAdapterMap_AddItemText(hAdapterHeader, "name3", "ccc")
	xcgui.XAdapterMap_AddItemText(hAdapterHeader, "name4", "test")

	hAdapter := xcgui.XAdapterTable_Create()
	xcgui.XList_BindAdapter(hList, hAdapter)
	xcgui.XAdapterTable_AddColumn(hAdapter, "name")
	xcgui.XAdapterTable_AddColumn(hAdapter, "name2")
	xcgui.XAdapterTable_AddColumn(hAdapter, "name3")

	xcgui.XList_SetColumnWidth(hList, 0, 150)
	xcgui.XList_SetColumnWidth(hList, 1, 150)
	xcgui.XList_SetColumnWidth(hList, 2, 150)

	for i := 0; i < 20; i++ {
		xcgui.XAdapterTable_AddItemText(hAdapter, "item-"+fmt.Sprint(i))

		xcgui.XAdapterTable_SetItemText(hAdapter, i, 1, "child-"+fmt.Sprint(i)+"-1")
		xcgui.XAdapterTable_SetItemText(hAdapter, i, 2, "child-"+fmt.Sprint(i)+"-2")
	}

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
