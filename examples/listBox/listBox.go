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
	hWindow := xcgui.XWnd_Create(0, 0, 500, 400, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hListBox = xcgui.XListBox_Create(20, 40, 300, 310, xcgui.HXCGUI(hWindow))
	xcgui.XListBox_SetItemTemplateXML(hListBox, "../xml-template/ListBox_Item.xml")
	xcgui.XSView_SetLineSize(hListBox, 20, 20)

	hAdapter := xcgui.XAdapterTable_Create()
	xcgui.XListBox_BindAdapter(hListBox, hAdapter)
	xcgui.XAdapterTable_AddColumn(hAdapter, "name")
	xcgui.XAdapterTable_AddColumn(hAdapter, "name2")
	xcgui.XAdapterTable_AddColumn(hAdapter, "name3")

	for i := 0; i < 20; i++ {
		index := xcgui.XAdapterTable_AddItemTextEx(hAdapter, "name", "item-"+fmt.Sprint(i))
		if index != xcgui.XC_ID_ERROR {
			xcgui.XAdapterTable_SetItemTextEx(hAdapter, index, "name2", "item-1-"+fmt.Sprint(i))
			xcgui.XAdapterTable_SetItemTextEx(hAdapter, index, "name3", "item-2-"+fmt.Sprint(i))
		}
	}

	hButton1 := xcgui.XBtn_Create(330, 40, 150, 18, "插入", xcgui.HXCGUI(hWindow))
	xcgui.XEle_RegEventC(hButton1, xcgui.XE_BNCLICK, xcgui.CallBack(OnBtnClickAdd))

	hButton2 := xcgui.XBtn_Create(330, 60, 150, 18, "删除", xcgui.HXCGUI(hWindow))
	xcgui.XEle_RegEventC(hButton2, xcgui.XE_BNCLICK, xcgui.CallBack(OnBtnClickDel))

	hButton3 := xcgui.XBtn_Create(330, 84, 150, 18, "删除 index=(1-3)", xcgui.HXCGUI(hWindow))
	xcgui.XEle_RegEventC(hButton3, xcgui.XE_BNCLICK, xcgui.CallBack(OnBtnClickDelEx))

	hRichEidt = xcgui.XRichEdit_Create(330, 106, 120, 200, xcgui.HXCGUI(hWindow))
	xcgui.XSView_ShowSBarV(hRichEidt, true)
	xcgui.XRichEdit_EnableMultiLine(hRichEidt, true)

	xcgui.XEle_RegEventC(hListBox, xcgui.XE_LISTBOX_TEMP_CREATE_END, xcgui.CallBack(OnTemplateCreate))
	xcgui.XEle_RegEventC(hListBox, xcgui.XE_LISTBOX_TEMP_DESTROY, xcgui.CallBack(OnTemplateDestroy))
	xcgui.XEle_RegEventC(hListBox, xcgui.XE_LISTBOX_TEMP_ADJUST_COORDINATE, xcgui.CallBack(OnTemplateAdjustCoordinate))
	xcgui.XEle_RegEventC(hListBox, xcgui.XE_LISTBOX_SELECT, xcgui.CallBack(OnListBoxSelect))

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func OnBtnClickAdd(pbHandled *bool) int {
	hAdapter := xcgui.XListBox_GetAdapter(hListBox)
	xcgui.XAdapterTable_InsertItemText(hAdapter, xcgui.XAdapterTable_GetCount(hAdapter), "test - insert"+fmt.Sprint(xcgui.XAdapterTable_GetCount(hAdapter)))
	xcgui.XEle_RedrawEle(hListBox)

	return 0
}
func OnBtnClickDel(pbHandled *bool) int {
	hAdapter := xcgui.XListBox_GetAdapter(hListBox)
	array := make([]uint16, 256)
	count := xcgui.XListBox_GetSelectAll(hListBox, &array[0], 256)
	for i := 0; i < count*2; i += 2 {
		ok := xcgui.XAdapterTable_DeleteItem(hAdapter, int(array[i]))
		if ok {
			xcgui.XEle_RedrawEle(hListBox)
		}
	}

	return 0
}

func OnBtnClickDelEx(pbHandled *bool) int {
	hAdapter := xcgui.XListBox_GetAdapter(hListBox)
	ok := xcgui.XAdapterTable_DeleteItemEx(hAdapter, 1, 3)
	if ok {
		xcgui.XEle_RedrawEle(hListBox)
	}

	return 0
}

func OnListBoxSelect(iItem int, pbHandled *bool) int {
	xcgui.XRichEdit_DeleteAll(hRichEidt)
	array := make([]uint16, 256)
	count := xcgui.XListBox_GetSelectAll(hListBox, &array[0], 256)

	hAdapter := xcgui.XListBox_GetAdapter(hListBox)
	var szItemList string
	name := make([]uint16, 256)
	for i := 0; i < count*2; i += 2 {
		ok := xcgui.XAdapterTable_GetItemText(hAdapter, int(array[i]), 0, &name[0], 256)
		if ok {
			szItemList += xcgui.UTF16PtrToString(&name[0])
			szItemList += "\n"
		}
	}

	xcgui.XRichEdit_SetText(hRichEidt, szItemList)

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
