package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xList_Create                       *syscall.Proc
	xList_AddColumn                    *syscall.Proc
	xList_InsertColumn                 *syscall.Proc
	xList_EnableMultiSel               *syscall.Proc
	xList_SetDrawItemBkFlags           *syscall.Proc
	xList_SetColumnWidth               *syscall.Proc
	xList_SetColumnMinWidth            *syscall.Proc
	xList_SetItemData                  *syscall.Proc
	xList_GetItemData                  *syscall.Proc
	xList_SetSelectItem                *syscall.Proc
	xList_GetSelectItem                *syscall.Proc
	xList_GetSelectItemCount           *syscall.Proc
	xList_SelectAll                    *syscall.Proc
	xList_GetHeaderHELE                *syscall.Proc
	xList_DeleteColumn                 *syscall.Proc
	xList_DeleteColumnAll              *syscall.Proc
	xList_BindAdapter                  *syscall.Proc
	xList_BindAdapterHeader            *syscall.Proc
	xList_GetAdapter                   *syscall.Proc
	xList_SetItemTemplateXML           *syscall.Proc
	xList_SetItemTemplateXMLFromString *syscall.Proc
	xList_GetTemplateObject            *syscall.Proc
	xList_GetItemIndexFromHXCGUI       *syscall.Proc
	xList_SetHeaderHeight              *syscall.Proc
	xList_GetHeaderHeight              *syscall.Proc
	xList_AddItemBkBorder              *syscall.Proc
	xList_AddItemBkFill                *syscall.Proc
	xList_AddItemBkImage               *syscall.Proc
	xList_GetItemBkInfoCount           *syscall.Proc
	xList_ClearItemBkInfo              *syscall.Proc
	xList_GetItemBkInfoManager         *syscall.Proc
	xList_SetItemHeightDefault         *syscall.Proc
	xList_GetItemHeightDefault         *syscall.Proc
	xList_HitTest                      *syscall.Proc
	xList_HitTestOffset                *syscall.Proc
	xList_RefreshData                  *syscall.Proc
)

func init() {
	// Functions
	xList_Create = xcDLL.MustFindProc("XList_Create")
	xList_AddColumn = xcDLL.MustFindProc("XList_AddColumn")
	xList_InsertColumn = xcDLL.MustFindProc("XList_InsertColumn")
	xList_EnableMultiSel = xcDLL.MustFindProc("XList_EnableMultiSel")
	xList_SetDrawItemBkFlags = xcDLL.MustFindProc("XList_SetDrawItemBkFlags")
	xList_SetColumnWidth = xcDLL.MustFindProc("XList_SetColumnWidth")
	xList_SetColumnMinWidth = xcDLL.MustFindProc("XList_SetColumnMinWidth")
	xList_SetItemData = xcDLL.MustFindProc("XList_SetItemData")
	xList_GetItemData = xcDLL.MustFindProc("XList_GetItemData")
	xList_SetSelectItem = xcDLL.MustFindProc("XList_SetSelectItem")
	xList_GetSelectItem = xcDLL.MustFindProc("XList_GetSelectItem")
	xList_GetSelectItemCount = xcDLL.MustFindProc("XList_GetSelectItemCount")
	xList_SelectAll = xcDLL.MustFindProc("XList_SelectAll")
	xList_GetHeaderHELE = xcDLL.MustFindProc("XList_GetHeaderHELE")
	xList_DeleteColumn = xcDLL.MustFindProc("XList_DeleteColumn")
	xList_DeleteColumnAll = xcDLL.MustFindProc("XList_DeleteColumnAll")
	xList_BindAdapter = xcDLL.MustFindProc("XList_BindAdapter")
	xList_BindAdapterHeader = xcDLL.MustFindProc("XList_BindAdapterHeader")
	xList_GetAdapter = xcDLL.MustFindProc("XList_GetAdapter")
	xList_SetItemTemplateXML = xcDLL.MustFindProc("XList_SetItemTemplateXML")
	xList_SetItemTemplateXMLFromString = xcDLL.MustFindProc("XList_SetItemTemplateXMLFromString")
	xList_GetTemplateObject = xcDLL.MustFindProc("XList_GetTemplateObject")
	xList_GetItemIndexFromHXCGUI = xcDLL.MustFindProc("XList_GetItemIndexFromHXCGUI")
	xList_SetHeaderHeight = xcDLL.MustFindProc("XList_SetHeaderHeight")
	xList_GetHeaderHeight = xcDLL.MustFindProc("XList_GetHeaderHeight")
	xList_AddItemBkBorder = xcDLL.MustFindProc("XList_AddItemBkBorder")
	xList_AddItemBkFill = xcDLL.MustFindProc("XList_AddItemBkFill")
	xList_AddItemBkImage = xcDLL.MustFindProc("XList_AddItemBkImage")
	xList_GetItemBkInfoCount = xcDLL.MustFindProc("XList_GetItemBkInfoCount")
	xList_ClearItemBkInfo = xcDLL.MustFindProc("XList_ClearItemBkInfo")
	xList_GetItemBkInfoManager = xcDLL.MustFindProc("XList_GetItemBkInfoManager")
	xList_SetItemHeightDefault = xcDLL.MustFindProc("XList_SetItemHeightDefault")
	xList_GetItemHeightDefault = xcDLL.MustFindProc("XList_GetItemHeightDefault")
	xList_HitTest = xcDLL.MustFindProc("XList_HitTest")
	xList_HitTestOffset = xcDLL.MustFindProc("XList_HitTestOffset")
	xList_RefreshData = xcDLL.MustFindProc("XList_RefreshData")
}

/*
创建列表元素.

参数:
	x 元素x坐标.
	y 元素y坐标.
	cx 宽度.
	cy 高度.
	hParent 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
返回:
	元素句柄.
*/
func XListCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xList_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
增加列.

参数:
	hEle 元素句柄.
	width 列宽度.
返回:
	返回位置索引.
*/
func XListAddColumn(hEle HELE, width int) int {
	ret, _, _ := xList_AddColumn.Call(
		uintptr(hEle),
		uintptr(width))

	return int(ret)
}

/*
插入列.

参数:
	hEle 元素句柄.
	width 列宽度.
	iItem 插入位置索引.
返回:
	返回插入位置索引.
*/
func XListInsertColumn(hEle HELE, width, iItem int) int {
	ret, _, _ := xList_InsertColumn.Call(
		uintptr(hEle),
		uintptr(width),
		uintptr(iItem))

	return int(ret)
}

/*
启用或关闭多选功能.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XListEnableMultiSel(hEle HELE, bEnable bool) {
	xList_EnableMultiSel.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
设置是否绘制指定状态下项的背景.

参数:
	hEle 元素句柄.
	nFlags 标志位 list_drawItemBk_flags_.
*/
func XListSetDrawItemBkFlags(hEle HELE, nFlags List_drawItemBk_flags_) {
	xList_SetDrawItemBkFlags.Call(
		uintptr(hEle),
		uintptr(nFlags))
}

/*
设置列宽度.

参数:
	hEle 元素句柄.
	iItem 列索引.
	width 宽度.
*/
func XListSetColumnWidth(hEle HELE, iItem, width int) {
	xList_SetColumnWidth.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(width))
}

/*
设置列最小宽度.

参数:
	hEle 元素句柄.
	iItem 列索引.
	width 宽度.
*/
func XListSetColumnMinWidth(hEle HELE, iItem, width int) {
	xList_SetColumnMinWidth.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(width))
}

/*
设置项用户数据.

参数:
	hEle 元素句柄.
	iItem 项索引.
	iSubItem 子项索引.
	data 用户数据.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListSetItemData(hEle HELE, iItem, iSubItem, data int) bool {
	ret, _, _ := xList_SetItemData.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(iSubItem),
		uintptr(data))

	return ret == TRUE
}

/*
获取项用户数据.

参数:
	hEle 元素句柄.
	iItem 项索引.
	iSubItem 子项索引.
返回:
	返回用户数据.
*/
func XListGetItemData(hEle HELE, iItem, iSubItem int) int {
	ret, _, _ := xList_GetItemData.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(iSubItem))

	return int(ret)
}

/*
设置选择项.

参数:
	hEle 元素句柄.
	iItem 项索引.
	bSelect 是否选择.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListSetSelectItem(hEle HELE, iItem int, bSelect bool) bool {
	ret, _, _ := xList_SetSelectItem.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(BoolToBOOL(bSelect)))

	return ret == TRUE
}

/*
获取选择项.

参数:
	hEle 元素句柄.
返回:
	项索引.
*/
func XListGetSelectItem(hEle HELE) int {
	ret, _, _ := xList_GetSelectItem.Call(uintptr(hEle))

	return int(ret)
}

/*
获取选择项数量.

参数:
	hEle 元素句柄.
返回:
	返回选择项数量.
*/
func XListGetSelectItemCount(hEle HELE) int {
	ret, _, _ := xList_GetSelectItemCount.Call(uintptr(hEle))

	return int(ret)
}

/*
选择全部行.

参数:
	hEle 元素句柄.
*/
func XListSelectAll(hEle HELE) {
	xList_SelectAll.Call(uintptr(hEle))
}

/*
获取列表头元素.

参数:
	hEle 元素句柄.
返回:
	返回列表头元素句柄.
*/
func XListGetHeaderHELE(hEle HELE) HELE {
	ret, _, _ := xList_GetHeaderHELE.Call(uintptr(hEle))

	return HELE(ret)
}

/*
删除列.

参数:
	hEle 元素句柄.
	iItem 项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListDeleteColumn(hEle HELE, iItem int) bool {
	ret, _, _ := xList_DeleteColumn.Call(
		uintptr(hEle),
		uintptr(iItem))

	return ret == TRUE
}

/*
删除所有的列.

参数:
	hEle 元素句柄.
*/
func XListDeleteColumnAll(hEle HELE) {
	xList_DeleteColumnAll.Call(uintptr(hEle))
}

/*
绑定数据适配器.

参数:
	hEle 元素句柄.
	hAdapter 数据适配器句柄 XAdapterTable.
*/
func XListBindAdapter(hEle HELE, hAdapter HXCGUI) {
	xList_BindAdapter.Call(
		uintptr(hEle),
		uintptr(hAdapter))
}

/*
列表头绑定数据适配器.

参数:
	hEle 元素句柄.
	hAdapter 数据适配器句柄 XAdapterMap.
*/
func XListBindAdapterHeader(hEle HELE, hAdapter HXCGUI) {
	xList_BindAdapterHeader.Call(
		uintptr(hEle),
		uintptr(hAdapter))
}

/*
获取数据适配器.

参数:
	hEle 元素句柄.
返回:
	数据适配器句柄.
*/
func XListGetAdapter(hEle HELE) HELE {
	ret, _, _ := xList_GetAdapter.Call(uintptr(hEle))

	return HELE(ret)
}

/*
设置项布局模板文件.

参数:
	hEle 元素句柄.
	pXmlFile 文件名.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListSetItemTemplateXML(hEle HELE, pXmlFile *uint16) bool {
	ret, _, _ := xList_SetItemTemplateXML.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pXmlFile)))

	return ret == TRUE
}

/*
设置项布局模板文件.
参数:
hEle 元素句柄.
pStringXML 字符串指针.
返回:成功返回TRUE否则返回FALSE.
*/
func XListSetItemTemplateXMLFromString(hEle HELE, pStringXML *uint16) bool {
	ret, _, _ := xList_SetItemTemplateXMLFromString.Call(
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
func XListGetTemplateObject(hEle HELE, iItem, nTempItemID int) HXCGUI {
	ret, _, _ := xList_GetTemplateObject.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(nTempItemID))

	return HXCGUI(ret)
}

/*
获取当前对象所在模板实例,属于列表中哪一个项.

参数:
	hEle 元素句柄.
	hXCGUI 对象句柄, UI元素句柄或形状对象句柄.
返回:
	成功返回项索引, 否则返回XC_ID_ERROR.
*/
func XListGetItemIndexFromHXCGUI(hEle HELE, hXCGUI HXCGUI) int {
	ret, _, _ := xList_GetItemIndexFromHXCGUI.Call(
		uintptr(hEle),
		uintptr(hXCGUI))

	return int(ret)
}

/*
设置列表头高度.

参数:
	hEle 元素句柄.
	height 高度.
*/
func XListSetHeaderHeight(hEle HELE, height int) {
	xList_SetHeaderHeight.Call(
		uintptr(hEle),
		uintptr(height))
}

/*
获取列表头高度.

参数:
	hEle 元素句柄.
返回:
	返回列表头高度.
*/
func XListGetHeaderHeight(hEle HELE) int {
	ret, _, _ := xList_GetHeaderHeight.Call(uintptr(hEle))

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
func XListAddItemBkBorder(hEle HELE, nState List_item_state_, color COLORREF, alpha byte, width int) {
	xList_AddItemBkBorder.Call(
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
func XListAddItemBkFill(hEle HELE, nState List_item_state_, color COLORREF, alpha byte) {
	xList_AddItemBkFill.Call(
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
func XListAddItemBkImage(hEle HELE, nState List_item_state_, hImage HIMAGE) {
	xList_AddItemBkImage.Call(
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
func XListGetItemBkInfoCount(hEle HELE, nState List_item_state_) int {
	ret, _, _ := xList_GetItemBkInfoCount.Call(
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
func XListClearItemBkInfo(hEle HELE, nState List_item_state_) {
	xList_ClearItemBkInfo.Call(
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
func XListGetItemBkInfoManager(hEle HELE, nState List_item_state_) HBKINFOM {
	ret, _, _ := xList_GetItemBkInfoManager.Call(
		uintptr(hEle),
		uintptr(nState))

	return HBKINFOM(ret)
}

/*
设置项默认高度.

参数:
	hEle 元素句柄.
	nHeight 高度.
	nSelHeight 选中时高度.
*/
func XListSetItemHeightDefault(hEle HELE, nHeight, nSelHeight int) {
	xList_SetItemHeightDefault.Call(
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
func XListGetItemHeightDefault(hEle HELE, pHeight, pSelHeight *uint16) {
	xList_GetItemHeightDefault.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pHeight)),
		uintptr(unsafe.Pointer(pSelHeight)))
}

/*
检测坐标点所在项.

参数:
	hEle 元素句柄.
	pPt 坐标点.
	piItem 项索引.
	piSubItem 子项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListHitTest(hEle HELE, pPt *POINT, piItem, piSubItem *uint16) bool {
	ret, _, _ := xList_HitTest.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)),
		uintptr(unsafe.Pointer(piItem)),
		uintptr(unsafe.Pointer(piSubItem)))

	return ret == TRUE
}

/*
检查坐标点所在项,自动添加滚动视图偏移量.

参数:
	hEle 元素句柄.
	pPt 坐标点.
	piItem 项索引.
	piSubItem 子项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListHitTestOffset(hEle HELE, pPt *POINT, piItem, piSubItem *uint16) bool {
	ret, _, _ := xList_HitTestOffset.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)),
		uintptr(unsafe.Pointer(piItem)),
		uintptr(unsafe.Pointer(piSubItem)))

	return ret == TRUE
}

/*
刷新数据.

参数:
	hEle 元素句柄.
*/
func XListRefreshData(hEle HELE) {
	xList_RefreshData.Call(uintptr(hEle))
}
