package main

import (
	"fmt"
)

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWndCreate(0, 0, 550, 380, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hList := xcgui.XListCreate(20, 40, 500, 300, xcgui.HXCGUI(hWindow))
	xcgui.XListSetItemTemplateXML(hList, "../xml-template/List_Item.xml")
	xcgui.XSViewSetLineSize(hList, 20, 20)

	for i := 0; i < 3; i++ {
		xcgui.XListAddColumn(hList, 100)
	}

	hAdapterHeader := xcgui.XAdapterMapCreate()
	xcgui.XListBindAdapterHeader(hList, hAdapterHeader)

	xcgui.XAdapterMapAddItemText(hAdapterHeader, "name", "aaa")
	xcgui.XAdapterMapAddItemText(hAdapterHeader, "name2", "bbb")
	xcgui.XAdapterMapAddItemText(hAdapterHeader, "name3", "ccc")
	xcgui.XAdapterMapAddItemText(hAdapterHeader, "name4", "test")

	hAdapter := xcgui.XAdapterTableCreate()
	xcgui.XListBindAdapter(hList, hAdapter)
	xcgui.XAdapterTableAddColumn(hAdapter, "name")
	xcgui.XAdapterTableAddColumn(hAdapter, "name2")
	xcgui.XAdapterTableAddColumn(hAdapter, "name3")

	xcgui.XListSetColumnWidth(hList, 0, 150)
	xcgui.XListSetColumnWidth(hList, 1, 150)
	xcgui.XListSetColumnWidth(hList, 2, 150)

	for i := 0; i < 20; i++ {
		xcgui.XAdapterTableAddItemText(hAdapter, "item-"+fmt.Sprint(i))

		xcgui.XAdapterTableSetItemText(hAdapter, i, 1, "child-"+fmt.Sprint(i)+"-1")
		xcgui.XAdapterTableSetItemText(hAdapter, i, 2, "child-"+fmt.Sprint(i)+"-2")
	}

	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
