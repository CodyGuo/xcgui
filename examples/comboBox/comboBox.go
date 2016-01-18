package main

import (
	"fmt"
)

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWnd_Create(0, 0, 300, 200, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hComboBox := xcgui.XComboBox_Create(20, 40, 120, 20, xcgui.HXCGUI(hWindow))
	xcgui.XComboBox_SetItemTemplateXML(hComboBox, "../xml-template/ComboBox_ListBox_Item.xml")
	xcgui.XRichEdit_SetText(hComboBox, "123")

	hAdapter := xcgui.XAdapterTable_Create()
	xcgui.XComboBox_BindApapter(hComboBox, hAdapter)
	xcgui.XAdapterTable_AddColumn(hAdapter, "name")

	for i := 0; i < 20; i++ {
		xcgui.XAdapterTable_AddItemText(hAdapter, "name-"+fmt.Sprint(i)+"-0")
	}

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
