package xc

import (
	"syscall"
	"unsafe"
)

var (
	xComboBox_Create                       *syscall.Proc
	xComboBox_SetSelItem                   *syscall.Proc
	xComboBox_BindApapter                  *syscall.Proc
	xComboBox_GetButtonRect                *syscall.Proc
	xComboBox_SetButtonSize                *syscall.Proc
	xComboBox_SetDropHeight                *syscall.Proc
	xComboBox_GetDropHeight                *syscall.Proc
	xComboBox_SetItemTemplateXML           *syscall.Proc
	xComboBox_SetItemTemplateXMLFromString *syscall.Proc
	xComboBox_EnableDrawButton             *syscall.Proc
	xComboBox_EnableEdit                   *syscall.Proc
	xComboBox_AddBkBorder                  *syscall.Proc
	xComboBox_AddBkFill                    *syscall.Proc
	xComboBox_AddBkImage                   *syscall.Proc
	XComboboX_GetBkInfoCount               *syscall.Proc
	xComboBox_ClearBkInfo                  *syscall.Proc
	xComboBox_GetBkInfoManager             *syscall.Proc
	xComboBox_GetSelItem                   *syscall.Proc
	xComboBox_GetState                     *syscall.Proc
)

func init() {
	xComboBox_Create = XCDLL.MustFindProc("XComboBox_Create")
	xComboBox_SetSelItem = XCDLL.MustFindProc("XComboBox_SetSelItem")
	xComboBox_BindApapter = XCDLL.MustFindProc("XComboBox_BindApapter")
	xComboBox_GetButtonRect = XCDLL.MustFindProc("XComboBox_GetButtonRect")
	xComboBox_SetButtonSize = XCDLL.MustFindProc("XComboBox_SetButtonSize")
	xComboBox_SetDropHeight = XCDLL.MustFindProc("XComboBox_SetDropHeight")
	xComboBox_GetDropHeight = XCDLL.MustFindProc("XComboBox_GetDropHeight")
	xComboBox_SetItemTemplateXML = XCDLL.MustFindProc("XComboBox_SetItemTemplateXML")
	xComboBox_SetItemTemplateXMLFromString = XCDLL.MustFindProc("XComboBox_SetItemTemplateXMLFromString")
	xComboBox_EnableDrawButton = XCDLL.MustFindProc("XComboBox_EnableDrawButton")
	xComboBox_EnableEdit = XCDLL.MustFindProc("XComboBox_EnableEdit")
	xComboBox_AddBkBorder = XCDLL.MustFindProc("XComboBox_AddBkBorder")
	xComboBox_AddBkFill = XCDLL.MustFindProc("XComboBox_AddBkFill")
	xComboBox_AddBkImage = XCDLL.MustFindProc("XComboBox_AddBkImage")
	xComboBox_GetBkInfoManager = XCDLL.MustFindProc("XComboBox_GetBkInfoManager")
	XComboboX_GetBkInfoCount = XCDLL.MustFindProc("XComboboX_GetBkInfoCount")
	xComboBox_Create = XCDLL.MustFindProc("XComboBox_Create")
	xComboBox_GetSelItem = XCDLL.MustFindProc("XComboBox_GetSelItem")
	xComboBox_ClearBkInfo = XCDLL.MustFindProc("XComboBox_ClearBkInfo")
}

func XComboBoxCreate(x int, y int, cx int, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xComboBox_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))
	return HELE(ret)
}

func XComboBoxSetSelItem(hEle HELE, iIndex int) bool {
	ret, _, _ := xComboBox_SetSelItem.Call(
		uintptr(hEle),
		uintptr(iIndex))
	return ret == TRUE

}

func XComboBoxBindApapter(hEle HELE, hAdapter HXCGUI) {
	xComboBox_BindApapter.Call(
		uintptr(hEle),
		uintptr(hAdapter))
}

func XComboBoxGetButtonRect(hEle HELE, pRect *RECT) {
	xComboBox_GetButtonRect.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

func XComboBoxSetButtonSize(hEle HELE, size int) {
	xComboBox_SetButtonSize.Call(
		uintptr(hEle),
		uintptr(size))
}

func XComboBoxSetDropHeight(hEle HELE, height int) {
	xComboBox_SetDropHeight.Call(
		uintptr(hEle),
		uintptr(height))
}

func XComboBoxGetDropHeight(hEle HELE) int {
	ret, _, _ := xComboBox_GetDropHeight.Call(uintptr(hEle))
	return int(ret)
}

func XComboBoxSetItemTemplateXML(hEle HELE, pXmlFile string) {
	xComboBox_SetItemTemplateXML.Call(
		uintptr(hEle),
		StringToUintPtr(pXmlFile))
}

func XComboBoxSetItemTemplateXMLFromString(hEle HELE, pStringXML string) {
	xComboBox_SetItemTemplateXMLFromString.Call(
		uintptr(hEle),
		StringToUintPtr(pStringXML))
}

func XComboBoxEnableDrawButton(hEle HELE, bEnable bool) {
	xComboBox_EnableDrawButton.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

func XComboBoxEnableEdit(hEle HELE, bEdit bool) {
	xComboBox_EnableEdit.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEdit)))
}

func XComboBoxAddBkBorder(hEle HELE, nState comboBox_state_, color COLORREF, alpha byte, width int) {
	xComboBox_AddBkBorder.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha),
		uintptr(width))
}

func XComboBoxAddBkFill(hEle HELE, nState comboBox_state_, color COLORREF, alpha byte) {
	xComboBox_AddBkFill.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha))
}

func XComboBoxAddBkImage(hEle HELE, nState comboBox_state_, hImage HIMAGE) {
	xComboBox_AddBkImage.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(hImage))
}

func XComboboXGetBkInfoCount(hEle HELE, nState comboBox_state_) int {
	ret, _, _ := XComboboX_GetBkInfoCount.Call(
		uintptr(hEle),
		uintptr(nState))
	return int(ret)
}

func XComboBoxClearBkInfo(hEle HELE, nState comboBox_state_) {
	xComboBox_ClearBkInfo.Call(
		uintptr(hEle),
		uintptr(nState))
}

func XComboBoxGetBkInfoManager(hEle HELE, nState comboBox_state_) HBKINFOM {
	ret, _, _ := xComboBox_GetBkInfoManager.Call(
		uintptr(hEle),
		uintptr(nState))
	return HBKINFOM(ret)
}

func XComboBoxGetSelItem(hEle HELE) int {
	ret, _, _ := xComboBox_GetSelItem.Call(
		uintptr(hEle))
	return int(ret)
}

func XComboBoxGetState(hEle HELE) comboBox_state_ {
	ret, _, _ := xComboBox_GetState.Call(
		uintptr(hEle))
	return comboBox_state_(ret)
}
