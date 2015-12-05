package main

import (
	"fmt"
)

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWndCreate(0, 0, 300, 200, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hComboBox := xcgui.XComboBoxCreate(20, 40, 120, 20, xcgui.HXCGUI(hWindow))
	xcgui.XComboBoxSetItemTemplateXML(hComboBox, "../xml-template/ComboBox_ListBox_Item.xml")
	xcgui.XRichEditSetText(hComboBox, "123")

	hAdapter := xcgui.XAdapterTableCreate()
	xcgui.XComboBoxBindApapter(hComboBox, hAdapter)
	xcgui.XAdapterTableAddColumn(hAdapter, "name")

	for i := 0; i < 20; i++ {
		xcgui.XAdapterTableAddItemText(hAdapter, "name-"+fmt.Sprint(i)+"-0")
	}

	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
