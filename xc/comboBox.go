package xc

import (
	"syscall"
	"unsafe"
)

var (
	xComboBox_Create                       *syscall.Proc
	xComboBox_SetSelItem                   *syscall.Proc
	xComboBox_BindApapter                  *syscall.Proc
	xComboBox_GetApapter                   *syscall.Proc
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
	xCombobox_GetBkInfoCount               *syscall.Proc
	xComboBox_ClearBkInfo                  *syscall.Proc
	// xComboBox_GetBkInfoManager             *syscall.Proc
	xComboBox_GetSelItem *syscall.Proc
	xComboBox_GetState   *syscall.Proc
)

func init() {
	xComboBox_Create = xcDLL.MustFindProc("XComboBox_Create")
	xComboBox_SetSelItem = xcDLL.MustFindProc("XComboBox_SetSelItem")
	xComboBox_BindApapter = xcDLL.MustFindProc("XComboBox_BindApapter")
	xComboBox_GetApapter = xcDLL.MustFindProc("XComboBox_GetApapter")
	xComboBox_GetButtonRect = xcDLL.MustFindProc("XComboBox_GetButtonRect")
	xComboBox_SetButtonSize = xcDLL.MustFindProc("XComboBox_SetButtonSize")
	xComboBox_SetDropHeight = xcDLL.MustFindProc("XComboBox_SetDropHeight")
	xComboBox_GetDropHeight = xcDLL.MustFindProc("XComboBox_GetDropHeight")
	xComboBox_SetItemTemplateXML = xcDLL.MustFindProc("XComboBox_SetItemTemplateXML")
	xComboBox_SetItemTemplateXMLFromString = xcDLL.MustFindProc("XComboBox_SetItemTemplateXMLFromString")
	xComboBox_EnableDrawButton = xcDLL.MustFindProc("XComboBox_EnableDrawButton")
	xComboBox_EnableEdit = xcDLL.MustFindProc("XComboBox_EnableEdit")
	xComboBox_AddBkBorder = xcDLL.MustFindProc("XComboBox_AddBkBorder")
	xComboBox_AddBkFill = xcDLL.MustFindProc("XComboBox_AddBkFill")
	xComboBox_AddBkImage = xcDLL.MustFindProc("XComboBox_AddBkImage")
	// xComboBox_GetBkInfoManager = xcDLL.MustFindProc("XComboBox_GetBkInfoManager")
	xCombobox_GetBkInfoCount = xcDLL.MustFindProc("XComboboX_GetBkInfoCount")
	xComboBox_Create = xcDLL.MustFindProc("XComboBox_Create")
	xComboBox_GetSelItem = xcDLL.MustFindProc("XComboBox_GetSelItem")
	xComboBox_ClearBkInfo = xcDLL.MustFindProc("XComboBox_ClearBkInfo")
}

/*
创建下拉组合框元素.

参数:
	x 元素x坐标.
	y 元素y坐标.
	cx 宽度.
	cy 高度.
	hParent 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
返回:
	元素句柄.
*/
func XComboBox_Create(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xComboBox_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
设置选择项.

参数:
	hEle 元素句柄.
	iIndex 想索引.
返回:
	成功返回否则返回FALSE.
*/
func XComboBox_SetSelItem(hEle HELE, iIndex int) bool {
	ret, _, _ := xComboBox_SetSelItem.Call(
		uintptr(hEle),
		uintptr(iIndex))

	return ret == TRUE

}

/*
绑定数据适配器.

参数:
	hEle 元素句柄.
	hAdapter 适配器句柄.
*/
func XComboBox_BindApapter(hEle HELE, hAdapter HXCGUI) {
	xComboBox_BindApapter.Call(
		uintptr(hEle),
		uintptr(hAdapter))
}

/*
获取绑定的数据适配器.

参数:
	hEle 元素句柄.
返回:
	返回数据适配器.
*/
func XComboBox_GetApapter(hEle HELE) HXCGUI {
	ret, _, _ := xComboBox_GetApapter.Call(uintptr(hEle))

	return HXCGUI(ret)
}

/*
获取下拉按钮坐标.

参数:
	hEle 元素句柄.
	pRect 坐标.
*/
func XComboBox_GetButtonRect(hEle HELE, pRect *RECT) {
	xComboBox_GetButtonRect.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

/*
设置下拉按钮大小.

参数:
	hEle 元素句柄.
	size 大小.
*/
func XComboBox_SetButtonSize(hEle HELE, size int) {
	xComboBox_SetButtonSize.Call(
		uintptr(hEle),
		uintptr(size))
}

/*
设置下拉列表高度.

参数:
	hEle 元素句柄.
	height 高度.
*/
func XComboBox_SetDropHeight(hEle HELE, height int) {
	xComboBox_SetDropHeight.Call(
		uintptr(hEle),
		uintptr(height))
}

/*
获取下拉列表高度.

参数:
	hEle 元素句柄.
返回:
	下拉列表高度.
*/
func XComboBox_GetDropHeight(hEle HELE) int {
	ret, _, _ := xComboBox_GetDropHeight.Call(uintptr(hEle))

	return int(ret)
}

/*
设置下拉列表项模板文件.

参数:
	hEle 元素句柄.
	pXmlFile 项模板文件.*uint16
*/
func XComboBox_SetItemTemplateXML(hEle HELE, pXmlFile string) {
	xComboBox_SetItemTemplateXML.Call(
		uintptr(hEle),
		StringToUintPtr(pXmlFile))
	// uintptr(unsafe.Pointer(pXmlFile)))
}

/*
设置下拉列表项模板.

参数:
	hEle 元素句柄.
	pStringXML 字符串指针. *uint16
*/
func XComboBox_SetItemTemplateXMLFromString(hEle HELE, pStringXML string) {
	xComboBox_SetItemTemplateXMLFromString.Call(
		uintptr(hEle),
		StringToUintPtr(pStringXML))
	// uintptr(unsafe.Pointer(pStringXML)))
}

/*
是否绘制下拉按钮.

参数:
	hEle 元素句柄.
	bEnable 是否绘制.
*/
func XComboBox_EnableDrawButton(hEle HELE, bEnable bool) {
	xComboBox_EnableDrawButton.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用可编辑显示的文本内容.

参数:
	hEle 元素句柄.
	bEdit TRUE可编辑,否则相反.
*/
func XComboBox_EnableEdit(hEle HELE, bEdit bool) {
	xComboBox_EnableEdit.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEdit)))
}

/*
添加背景内容边框.

参数:
	hEle 元素句柄.
	nState 按钮状态.
	color RGB颜色.
	alpha 透明度.
	width 线宽.
*/
func XComboBox_AddBkBorder(hEle HELE, nState ComboBox_state_, color COLORREF, alpha byte, width int) {
	xComboBox_AddBkBorder.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha),
		uintptr(width))
}

/*
添加背景内容填充.
参数:
	hEle 元素句柄.
	nState 按钮状态.
	color RGB颜色.
	alpha 透明度.
*/
func XComboBox_AddBkFill(hEle HELE, nState ComboBox_state_, color COLORREF, alpha byte) {
	xComboBox_AddBkFill.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha))
}

/*
添加背景内容图片.

参数:
	hEle 元素句柄.
	nState 按钮状态.
	hImage 图片句柄.
*/
func XComboBox_AddBkImage(hEle HELE, nState ComboBox_state_, hImage HIMAGE) {
	xComboBox_AddBkImage.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(hImage))
}

/*
获取背景内容数量.

参数:
	hEle 元素句柄.
	nState 按钮状态.
返回:
	成功返回背景内容数量,否则返回XC_ID_ERROR.
*/
func XComboboxGetBkInfoCount(hEle HELE, nState ComboBox_state_) int {
	ret, _, _ := xCombobox_GetBkInfoCount.Call(
		uintptr(hEle),
		uintptr(nState))

	return int(ret)
}

/*
清空背景内容; 如果背景没有内容,将使用系统默认内容,以便保证背景正确.

参数:
	hEle 元素句柄.
	nState 按钮状态.
*/
func XComboBox_ClearBkInfo(hEle HELE, nState ComboBox_state_) {
	xComboBox_ClearBkInfo.Call(
		uintptr(hEle),
		uintptr(nState))
}

/*
获取背景内容管理器. 1.8.9.6

参数:
	hEle 元素句柄.
	nState 状态.
返回:
	背景内容管理器.
*/
// func XComboBox_GetBkInfoManager(hEle HELE, nState ComboBox_state_) HBKINFOM {
// 	ret, _, _ := xComboBox_GetBkInfoManager.Call(
// 		uintptr(hEle),
// 		uintptr(nState))

// 	return HBKINFOM(ret)
// }

/*
获取组合框下拉列表中选择项索引.

参数:
	hEle 元素句柄.
返回:
	返回项索引.
*/
func XComboBox_GetSelItem(hEle HELE) int {
	ret, _, _ := xComboBox_GetSelItem.Call(
		uintptr(hEle))

	return int(ret)
}

/*
获取状态.

参数:
	hEle 元素句柄.
返回:
	状态.
*/
func XComboBox_GetState(hEle HELE) ComboBox_state_ {
	ret, _, _ := xComboBox_GetState.Call(
		uintptr(hEle))

	return ComboBox_state_(ret)
}
