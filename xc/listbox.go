package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xListBox_Create                       *syscall.Proc
	xListBox_SetDrawItemBkFlags           *syscall.Proc
	xListBox_SetItemData                  *syscall.Proc
	xListBox_GetItemData                  *syscall.Proc
	xListBox_AddItemBkBorder              *syscall.Proc
	xListBox_AddItemBkFill                *syscall.Proc
	xListBox_AddItemBkImage               *syscall.Proc
	xListBox_GetItemBkInfoCount           *syscall.Proc
	xListBox_ClearItemBkInfo              *syscall.Proc
	xListBox_GetItemBkInfoManager         *syscall.Proc
	xListBox_SetItemInfo                  *syscall.Proc
	xListBox_GetItemInfo                  *syscall.Proc
	xListBox_SetSelectItem                *syscall.Proc
	xListBox_GetSelectItem                *syscall.Proc
	xListBox_CancelSelectItem             *syscall.Proc
	xListBox_CancelSelectAll              *syscall.Proc
	xListBox_GetSelectAll                 *syscall.Proc
	xListBox_GetSelectCount               *syscall.Proc
	xListBox_GetItemMouseStay             *syscall.Proc
	xListBox_SelectAll                    *syscall.Proc
	xListBox_SetItemHeightDefault         *syscall.Proc
	xListBox_GetItemHeightDefault         *syscall.Proc
	xListBox_GetItemIndexFromHXCGUI       *syscall.Proc
	xListBox_HitTest                      *syscall.Proc
	xListBox_HitTestOffset                *syscall.Proc
	xListBox_SetItemTemplateXML           *syscall.Proc
	xListBox_SetItemTemplateXMLFromString *syscall.Proc
	xListBox_GetTemplateObject            *syscall.Proc
	xListBox_EnableMultiSel               *syscall.Proc
	xListBox_BindAdapter                  *syscall.Proc
	xListBox_GetAdapter                   *syscall.Proc
)

func init() {
	// Functions
	xListBox_Create = xcDLL.MustFindProc("XListBox_Create")
	xListBox_SetDrawItemBkFlags = xcDLL.MustFindProc("XListBox_SetDrawItemBkFlags")
	xListBox_SetItemData = xcDLL.MustFindProc("XListBox_SetItemData")
	xListBox_GetItemData = xcDLL.MustFindProc("XListBox_GetItemData")
	xListBox_AddItemBkBorder = xcDLL.MustFindProc("XListBox_AddItemBkBorder")
	xListBox_AddItemBkFill = xcDLL.MustFindProc("XListBox_AddItemBkFill")
	xListBox_AddItemBkImage = xcDLL.MustFindProc("XListBox_AddItemBkImage")
	xListBox_GetItemBkInfoCount = xcDLL.MustFindProc("XListBox_GetItemBkInfoCount")
	xListBox_ClearItemBkInfo = xcDLL.MustFindProc("XListBox_ClearItemBkInfo")
	xListBox_GetItemBkInfoManager = xcDLL.MustFindProc("XListBox_GetItemBkInfoManager")
	xListBox_SetItemInfo = xcDLL.MustFindProc("XListBox_SetItemInfo")
	xListBox_GetItemInfo = xcDLL.MustFindProc("XListBox_GetItemInfo")
	xListBox_SetSelectItem = xcDLL.MustFindProc("XListBox_SetSelectItem")
	xListBox_GetSelectItem = xcDLL.MustFindProc("XListBox_GetSelectItem")
	xListBox_CancelSelectItem = xcDLL.MustFindProc("XListBox_CancelSelectItem")
	xListBox_CancelSelectAll = xcDLL.MustFindProc("XListBox_CancelSelectAll")
	xListBox_GetSelectAll = xcDLL.MustFindProc("XListBox_GetSelectAll")
	xListBox_GetSelectCount = xcDLL.MustFindProc("XListBox_GetSelectCount")
	xListBox_GetItemMouseStay = xcDLL.MustFindProc("XListBox_GetItemMouseStay")
	xListBox_SelectAll = xcDLL.MustFindProc("XListBox_SelectAll")
	xListBox_SetItemHeightDefault = xcDLL.MustFindProc("XListBox_SetItemHeightDefault")
	xListBox_GetItemHeightDefault = xcDLL.MustFindProc("XListBox_GetItemHeightDefault")
	xListBox_GetItemIndexFromHXCGUI = xcDLL.MustFindProc("XListBox_GetItemIndexFromHXCGUI")
	xListBox_HitTest = xcDLL.MustFindProc("XListBox_HitTest")
	xListBox_HitTestOffset = xcDLL.MustFindProc("XListBox_HitTestOffset")
	xListBox_SetItemTemplateXML = xcDLL.MustFindProc("XListBox_SetItemTemplateXML")
	xListBox_SetItemTemplateXMLFromString = xcDLL.MustFindProc("XListBox_SetItemTemplateXMLFromString")
	xListBox_GetTemplateObject = xcDLL.MustFindProc("XListBox_GetTemplateObject")
	xListBox_EnableMultiSel = xcDLL.MustFindProc("XListBox_EnableMultiSel")
	xListBox_BindAdapter = xcDLL.MustFindProc("XListBox_BindAdapter")
	xListBox_GetAdapter = xcDLL.MustFindProc("XListBox_GetAdapter")
}

/*
创建列表框元素.

参数:
	x 元素x坐标.
	y 元素y坐标.
	cx 宽度.
	cy 高度.
	hParent 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
返回:
	元素句柄.
*/
func XListBoxCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xListBox_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
设置是否绘制指定状态下项的背景.

参数:
	hEle 元素句柄.
	nFlags 标志位 list_drawItemBk_flags_.
*/
func XListBoxSetDrawItemBkFlags(hEle HELE, nFlags List_drawitembk_flags_) {
	xListBox_SetDrawItemBkFlags.Call(
		uintptr(hEle),
		uintptr(nFlags))
}

/*
设置项用户数据.

参数:
	hEle 元素句柄.
	iItem 想索引.
	nUserData 用户数据.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListBoxSetItemData(hEle HELE, iItem, nUserData int) bool {
	ret, _, _ := xListBox_SetItemData.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(nUserData))

	return ret == TRUE
}

/*
获取项用户数据.

参数:
	hEle 元素句柄.
	iItem 项索引.
返回:
	用户数据.
*/
func XListBoxGetItemData(hEle HELE, iItem int) int {
	ret, _, _ := xListBox_GetItemData.Call(
		uintptr(hEle),
		uintptr(iItem))

	return int(ret)
}

/*
添加项背景内容边框.

参数:
	hEle 元素句柄.
	nState 项状态.
	color RGB颜色.
	alpha 透明度.
	width 线宽.
*/
func XListBoxAddItemBkBorder(hEle HELE, nState List_item_state_, color COLORREF, alpha byte, width int) {
	xListBox_AddItemBkBorder.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha),
		uintptr(width))
}

/*
添加项背景内容填充.

参数:
	hEle 元素句柄.
	nState 项状态.
	color RGB颜色.
	alpha 透明度.
*/
func XListBoxAddItemBkFill(hEle HELE, nState List_item_state_, color COLORREF, alpha byte) {
	xListBox_AddItemBkFill.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha))
}

/*
添加项背景内容图片.

参数:
	hEle 元素句柄.
	nState 项状态.
	hImage 图片句柄.
*/
func XListBoxAddItemBkImage(hEle HELE, nState List_item_state_, hImage HIMAGE) {
	xListBox_AddItemBkImage.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(hImage))
}

/*
获取背景内容数量.

参数:
	hEle 元素句柄.
	nState 项状态.
返回:
	成功返回背景内容数量,否则返回XC_ID_ERROR.
*/
func XListBoxGetItemBkInfoCount(hEle HELE, nState List_item_state_) int {
	ret, _, _ := xListBox_GetItemBkInfoCount.Call(
		uintptr(hEle),
		uintptr(nState))

	return int(ret)
}

/*
清空项背景内容; 如果背景没有内容,将使用系统默认内容,以便保证背景正确.

参数:
	hEle 元素句柄.
	nState 项状态.
*/
func XListBoxClearItemBkInfo(hEle HELE, nState List_item_state_) {
	xListBox_ClearItemBkInfo.Call(
		uintptr(hEle),
		uintptr(nState))
}

/*
获取项背景内容管理器.

参数:
	hEle 元素句柄.
	nState 项状态.
返回:
	项背景内容管理器.
*/
func XListBoxGetItemBkInfoManager(hEle HELE, nState List_item_state_) HBKINFOM {
	ret, _, _ := xListBox_GetItemBkInfoManager.Call(
		uintptr(hEle),
		uintptr(nState))

	return HBKINFOM(ret)
}

/*
设置项信息.

参数:
	hEle 元素句柄.
	iItem 项索引.
	pItem 项信息.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListBoxSetItemInfo(hEle HELE, iItemint int, pItem *Listbox_item_info_i) bool {
	ret, _, _ := xListBox_SetItemInfo.Call(
		uintptr(hEle),
		uintptr(iItemint),
		uintptr(unsafe.Pointer(pItem)))

	return ret == TRUE
}

/*
获取项信息.

参数:
	hEle 元素句柄.
	iItem 项索引.
	pItem 项信息.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListBoxGetItemInfo(hEle HELE, iItem int, pItem *Listbox_item_info_i) bool {
	ret, _, _ := xListBox_GetItemInfo.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pItem)))

	return ret == TRUE
}

/*
设置选择选.

参数:
	hEle 元素句柄.
	iItem 项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListBoxSetSelectItem(hEle HELE, iItem int) bool {
	ret, _, _ := xListBox_SetSelectItem.Call(
		uintptr(hEle),
		uintptr(iItem))

	return ret == TRUE
}

/*
获取选择项.

参数:
	hEle 元素句柄.
返回:
	项索引.
*/
func XListBoxGetSelectItem(hEle HELE) int {
	ret, _, _ := xListBox_GetSelectItem.Call(uintptr(hEle))

	return int(ret)
}

/*
取消选择指定项.

参数:
	hEle 元素句柄.
	iItem 项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListBoxCancelSelectItem(hEle HELE, iItem int) bool {
	ret, _, _ := xListBox_CancelSelectItem.Call(
		uintptr(hEle),
		uintptr(iItem))

	return ret == TRUE
}

/*
取消所有选中的项.

参数:
	hEle 元素句柄.
返回:
	如果之前有选择状态的项返回TRUE,此时可以更新UI,否则返回FALSE.
*/
func XListBoxCancelSelectAll(hEle HELE) bool {
	ret, _, _ := xListBox_CancelSelectAll.Call(uintptr(hEle))

	return ret == TRUE
}

/*
获取所有选择项.

参数:
	hEle 元素句柄.
	pArray 数组缓冲区.
	nArraySize 数组大小.
返回:
	返回接收数量.
*/
func XListBoxGetSelectAll(hEle HELE, pArray *uint16, nArraySize int) int {
	ret, _, _ := xListBox_GetSelectAll.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pArray)),
		uintptr(nArraySize))

	return int(ret)
}

/*
获取选择项数量.

参数:
	hEle 元素句柄.
返回:
	返回数量.
*/
func XListBoxGetSelectCount(hEle HELE) int {
	ret, _, _ := xListBox_GetSelectCount.Call(uintptr(hEle))

	return int(ret)
}

/*
获取鼠标停留项.

参数:
	hEle 元素句柄.
返回:
	返回鼠标所在项.
*/
func XListBoxGetItemMouseStay(hEle HELE) int {
	ret, _, _ := xListBox_GetItemMouseStay.Call(uintptr(hEle))

	return int(ret)
}

/*
选择所有项.

参数:
	hEle 元素句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListBoxSelectAll(hEle HELE) bool {
	ret, _, _ := xListBox_SelectAll.Call(uintptr(hEle))

	return ret == TRUE
}

/*
设置项默认高度.

参数:
	hEle 元素句柄.
	nHeight 项高度.
	nSelHeight 选中项高度.
*/
func XListBoxSetItemHeightDefault(hEle HELE, nHeight, nSelHeight int) {
	xListBox_SetItemHeightDefault.Call(
		uintptr(hEle),
		uintptr(nHeight),
		uintptr(nSelHeight))
}

/*
获取项默认高度.

参数:
	hEle 元素句柄.
	pHeight 高度.
	pSelHeight 选中时高度.
*/
func XListBoxGetItemHeightDefault(hEle HELE, pHeight, pSelHeight *uint16) {
	xListBox_GetItemHeightDefault.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pHeight)),
		uintptr(unsafe.Pointer(pSelHeight)))
}

/*
获取当前对象所在模板实例,属于列表中哪一个项.

参数:
	hEle 元素句柄.
	hXCGUI 对象句柄, UI元素句柄或形状对象句柄.
返回:
	成功返回项索引, 否则返回XC_ID_ERROR.
*/
func XListBoxGetItemIndexFromHXCGUI(hEle HELE, hXCGUI HXCGUI) int {
	ret, _, _ := xListBox_GetItemIndexFromHXCGUI.Call(
		uintptr(hEle),
		uintptr(hXCGUI))

	return int(ret)
}

/*
检测坐标点所在项.

参数:
	hEle 元素句柄.
	pPt 坐标点.
返回:
	返回项索引.
*/
func XListBoxHitTest(hEle HELE, pPt *POINT) int {
	ret, _, _ := xListBox_HitTest.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)))

	return int(ret)
}

/*
检测坐标点所在项,自动添加滚动视图偏移量.

参数:
	hEle 元素句柄.
	pPt 坐标点.
返回:
	项索引.
*/
func XListBoxHitTestOffset(hEle HELE, pPt *POINT) int {
	ret, _, _ := xListBox_HitTestOffset.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)))

	return int(ret)
}

/*
设置项布局模板文件.

参数:
	hEle 元素句柄.
	pXmlFile 文件名.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListBoxSetItemTemplateXML(hEle HELE, pXmlFile *uint16) bool {
	ret, _, _ := xListBox_SetItemTemplateXML.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pXmlFile)))

	return ret == TRUE
}

/*
设置项布局模板文件.

参数:
	hEle 元素句柄.
	pStringXML 字符串指针.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListBoxSetItemTemplateXMLFromString(hEle HELE, pStringXML *uint16) bool {
	ret, _, _ := xListBox_SetItemTemplateXMLFromString.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pStringXML)))

	return ret == TRUE
}

/*
通过模板项ID,获取实例化模板项ID对应的对象句柄.

参数:
	hEle 元素句柄.
	iItem 项索引.
	nTempItemID 模板项ID.
返回:
	成功返回对象句柄,否则返回NULL.
*/
func XListBoxGetTemplateObject(hEle HELE, iItem, nTempItemID int) HXCGUI {
	ret, _, _ := xListBox_GetTemplateObject.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(nTempItemID))

	return HXCGUI(ret)
}

/*
是否启用多行选择功能.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XListBox_EnableMultiSel(hEle HELE, bEnable bool) {
	xListBox_EnableMultiSel.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
绑定数据适配器.

参数:
	hEle 元素句柄.
	hAdapter 数据适配器句柄 XAdapterTable.
*/
func XListBox_BindAdapter(hEle HELE, hAdapter HXCGUI) {
	xListBox_BindAdapter.Call(
		uintptr(hEle),
		uintptr(hAdapter))
}

/*
获取绑定的数据适配器.

参数:
	hEle 元素句柄.
返回:
	返回数据适配器句柄.
*/
func XListBox_GetAdapter(hEle HELE) HXCGUI {
	ret, _, _ := xListBox_GetAdapter.Call(uintptr(hEle))

	return HXCGUI(ret)
}
