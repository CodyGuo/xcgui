package main

import (
	"fmt"
)

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

var (
	hListBox  xcgui.HELE
	hRichEidt xcgui.HELE
)

type template_info_i struct {
	info int
}

type listBox_item_i struct {
	nUserData  int
	nHeight    int
	nSelHeight int
	//	BOOL    bSelect;
	nState  xcgui.List_item_state_
	index   int
	rcItem  xcgui.RECT
	hLayout xcgui.HXCGUI
	pInfo   template_info_i
}

func main() {
	hWindow := xcgui.XWndCreate(0, 0, 500, 400, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hListBox = xcgui.XListBoxCreate(20, 40, 300, 310, xcgui.HXCGUI(hWindow))
	xcgui.XListBoxSetItemTemplateXML(hListBox, "../xml-template/ListBox_Item.xml")
	xcgui.XSViewSetLineSize(hListBox, 20, 20)

	hAdapter := xcgui.XAdapterTableCreate()
	xcgui.XListBoxBindAdapter(hListBox, hAdapter)
	xcgui.XAdapterTableAddColumn(hAdapter, "name")
	xcgui.XAdapterTableAddColumn(hAdapter, "name2")
	xcgui.XAdapterTableAddColumn(hAdapter, "name3")

	for i := 0; i < 20; i++ {
		index := xcgui.XAdapterTableAddItemTextEx(hAdapter, "name", "item-"+fmt.Sprint(i))
		if index != xcgui.XC_ID_ERROR {
			xcgui.XAdapterTableSetItemTextEx(hAdapter, index, "name2", "item-1-"+fmt.Sprint(i))
			xcgui.XAdapterTableSetItemTextEx(hAdapter, index, "name3", "item-2-"+fmt.Sprint(i))
		}
	}

	hButton1 := xcgui.XBtnCreate(330, 40, 150, 18, "插入", xcgui.HXCGUI(hWindow))
	xcgui.XEleRegEventC(hButton1, xcgui.XE_BNCLICK, xcgui.CallBack(OnBtnClickAdd))

	hButton2 := xcgui.XBtnCreate(330, 60, 150, 18, "删除", xcgui.HXCGUI(hWindow))
	xcgui.XEleRegEventC(hButton2, xcgui.XE_BNCLICK, xcgui.CallBack(OnBtnClickDel))

	hButton3 := xcgui.XBtnCreate(330, 84, 150, 18, "删除 index=(1-3)", xcgui.HXCGUI(hWindow))
	xcgui.XEleRegEventC(hButton3, xcgui.XE_BNCLICK, xcgui.CallBack(OnBtnClickDelEx))

	hRichEidt = xcgui.XRichEditCreate(330, 106, 120, 200, xcgui.HXCGUI(hWindow))
	xcgui.XSViewShowSBarV(hRichEidt, true)
	xcgui.XRichEditEnableMultiLine(hRichEidt, true)

	xcgui.XEleRegEventC(hListBox, xcgui.XE_LISTBOX_TEMP_CREATE_END, xcgui.CallBack(OnTemplateCreate))
	xcgui.XEleRegEventC(hListBox, xcgui.XE_LISTBOX_TEMP_DESTROY, xcgui.CallBack(OnTemplateDestroy))
	xcgui.XEleRegEventC(hListBox, xcgui.XE_LISTBOX_TEMP_ADJUST_COORDINATE, xcgui.CallBack(OnTemplateAdjustCoordinate))
	xcgui.XEleRegEventC(hListBox, xcgui.XE_LISTBOX_SELECT, xcgui.CallBack(OnListBoxSelect))

	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func OnBtnClickAdd(pbHandled *bool) int {
	hAdapter := xcgui.XListBoxGetAdapter(hListBox)
	xcgui.XAdapterTableInsertItemText(hAdapter, xcgui.XAdapterTableGetCount(hAdapter), "test - insert"+fmt.Sprint(xcgui.XAdapterTableGetCount(hAdapter)))
	xcgui.XEleRedrawEle(hListBox)

	return 0
}
func OnBtnClickDel(pbHandled *bool) int {
	hAdapter := xcgui.XListBoxGetAdapter(hListBox)
	array := make([]uint16, 256)
	count := xcgui.XListBoxGetSelectAll(hListBox, &array[0], 256)
	for i := 0; i < count*2; i += 2 {
		ok := xcgui.XAdapterTableDeleteItem(hAdapter, int(array[i]))
		if ok {
			xcgui.XEleRedrawEle(hListBox)
		}
	}

	return 0
}

func OnBtnClickDelEx(pbHandled *bool) int {
	hAdapter := xcgui.XListBoxGetAdapter(hListBox)
	ok := xcgui.XAdapterTableDeleteItemEx(hAdapter, 1, 3)
	if ok {
		xcgui.XEleRedrawEle(hListBox)
	}

	return 0
}

func OnListBoxSelect(iItem int, pbHandled *bool) int {
	xcgui.XRichEditDeleteAll(hRichEidt)
	array := make([]uint16, 256)
	count := xcgui.XListBoxGetSelectAll(hListBox, &array[0], 256)

	hAdapter := xcgui.XListBoxGetAdapter(hListBox)
	var szItemList string
	name := make([]uint16, 256)
	for i := 0; i < count*2; i += 2 {
		ok := xcgui.XAdapterTableGetItemText(hAdapter, int(array[i]), 0, &name[0], 256)
		if ok {
			szItemList += xcgui.UTF16PtrToString(&name[0])
			szItemList += "\n"
		}
	}

	xcgui.XRichEditSetText(hRichEidt, szItemList)

	return 0
}

func OnTemplateCreate(pItem *listBox_item_i, pbHandled *bool) int {
	*pbHandled = true

	return 0
}

func OnTemplateDestroy(pItem *listBox_item_i, pbHandled *bool) int {
	*pbHandled = true

	return 0
}

func OnTemplateAdjustCoordinate(pItem *listBox_item_i, pbHandled *bool) int {
	*pbHandled = true

	return 0
}
