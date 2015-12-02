package xc

import (
	"syscall"
	"unsafe"
)

var (
	xTree_Create                          *syscall.Proc
	xTree_SetItemTemplateXML              *syscall.Proc
	xTree_SetItemTemplateXMLSel           *syscall.Proc
	xTree_SetItemTemplateXMLFromString    *syscall.Proc
	xTree_SetItemTemplateXMLSelFromString *syscall.Proc
	xTree_SetDrawItemBkFlags              *syscall.Proc
	xTree_SetItemData                     *syscall.Proc
	xTree_GetItemData                     *syscall.Proc
	xTree_SetSelectItem                   *syscall.Proc
	xTree_GetSelectItem                   *syscall.Proc
	xTree_IsExpand                        *syscall.Proc
	xTree_ExpandItem                      *syscall.Proc
	xTree_HitTest                         *syscall.Proc
	xTree_HitTestOffet                    *syscall.Proc
	xTree_GetFirstChildItem               *syscall.Proc
	xTree_GetNextSiblingItem              *syscall.Proc
	xTree_GetParentItem                   *syscall.Proc
	xTree_BindAdapter                     *syscall.Proc
	xTree_GetAdapter                      *syscall.Proc
	xTree_SetIndentation                  *syscall.Proc
	xTree_GetIndentation                  *syscall.Proc
	xTree_SetItemHeightDefault            *syscall.Proc
	xTree_GetItemHeightDefault            *syscall.Proc
	xTree_SetItemHeight                   *syscall.Proc
	xTree_GetItemHeight                   *syscall.Proc
	xTree_AddItemBkBorder                 *syscall.Proc
	xTree_AddItemBkFill                   *syscall.Proc
	xTree_AddItemBkImage                  *syscall.Proc
	xTree_GetItemBkInfoCount              *syscall.Proc
	xTree_ClearItemBkInfo                 *syscall.Proc
	xTree_GetItemBkInfoManager            *syscall.Proc
	xTree_GetTemplateObject               *syscall.Proc
	xTree_GetItemIDFromHXCGUI             *syscall.Proc
)

func init() {
	xTree_Create = xcDLL.MustFindProc("XTree_Create")
	xTree_SetItemTemplateXML = xcDLL.MustFindProc("XTree_SetItemTemplateXML")
	xTree_SetItemTemplateXMLSel = xcDLL.MustFindProc("XTree_SetItemTemplateXMLSel")
	xTree_SetItemTemplateXMLFromString = xcDLL.MustFindProc("XTree_SetItemTemplateXMLFromString")
	xTree_SetItemTemplateXMLSelFromString = xcDLL.MustFindProc("XTree_SetItemTemplateXMLSelFromString")
	xTree_SetDrawItemBkFlags = xcDLL.MustFindProc("XTree_SetDrawItemBkFlags")
	xTree_SetItemData = xcDLL.MustFindProc("XTree_SetItemData")
	xTree_GetItemData = xcDLL.MustFindProc("XTree_GetItemData")
	xTree_SetSelectItem = xcDLL.MustFindProc("XTree_SetSelectItem")
	xTree_GetSelectItem = xcDLL.MustFindProc("XTree_GetSelectItem")
	xTree_IsExpand = xcDLL.MustFindProc("XTree_IsExpand")
	xTree_ExpandItem = xcDLL.MustFindProc("XTree_ExpandItem")
	xTree_HitTest = xcDLL.MustFindProc("XTree_HitTest")
	xTree_HitTestOffet = xcDLL.MustFindProc("XTree_HitTestOffet")
	xTree_GetFirstChildItem = xcDLL.MustFindProc("XTree_GetFirstChildItem")
	xTree_GetNextSiblingItem = xcDLL.MustFindProc("XTree_GetNextSiblingItem")
	xTree_GetParentItem = xcDLL.MustFindProc("XTree_GetParentItem")
	xTree_BindAdapter = xcDLL.MustFindProc("XTree_BindAdapter")
	xTree_GetAdapter = xcDLL.MustFindProc("XTree_GetAdapter")
	xTree_SetIndentation = xcDLL.MustFindProc("XTree_SetIndentation")
	xTree_GetIndentation = xcDLL.MustFindProc("XTree_GetIndentation")
	xTree_SetItemHeightDefault = xcDLL.MustFindProc("XTree_SetItemHeightDefault")
	xTree_GetItemHeightDefault = xcDLL.MustFindProc("XTree_GetItemHeightDefault")
	xTree_SetItemHeight = xcDLL.MustFindProc("XTree_SetItemHeight")
	xTree_GetItemHeight = xcDLL.MustFindProc("XTree_GetItemHeight")
	xTree_AddItemBkBorder = xcDLL.MustFindProc("XTree_AddItemBkBorder")
	xTree_AddItemBkFill = xcDLL.MustFindProc("XTree_AddItemBkFill")
	xTree_AddItemBkImage = xcDLL.MustFindProc("XTree_AddItemBkImage")
	xTree_GetItemBkInfoCount = xcDLL.MustFindProc("XTree_GetItemBkInfoCount")
	xTree_ClearItemBkInfo = xcDLL.MustFindProc("XTree_ClearItemBkInfo")
	xTree_GetItemBkInfoManager = xcDLL.MustFindProc("XTree_GetItemBkInfoManager")
	xTree_GetTemplateObject = xcDLL.MustFindProc("XTree_GetTemplateObject")
	xTree_GetItemIDFromHXCGUI = xcDLL.MustFindProc("XTree_GetItemIDFromHXCGUI")

}

/*
创建树元素.

参数:
	x 元素x坐标.
	y 元素y坐标.
	cx 宽度.
	cy 高度.
	hParent 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
返回:
	元素句柄.
*/
func XTreeCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xTree_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
设置项模板文件.

参数:
	hEle 元素句柄.
	pXmlFile 文件名.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XTreeSetItemTemplateXML(hEle HELE, pXmlFile *uint16) bool {
	ret, _, _ := xTree_SetItemTemplateXML.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pXmlFile)))

	return ret == TRUE
}

/*
设置项模板文件,项选中状态.

参数:
	hEle 元素句柄.
	pXmlFile 文件名.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XTreeSetItemTemplateXMLSel(hEle HELE, pXmlFile *uint16) bool {
	ret, _, _ := xTree_SetItemTemplateXMLSel.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pXmlFile)))

	return ret == TRUE
}

/*
设置项模板文件.

参数:
	hEle 元素句柄.
	pStringXML 字符串指针.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XTreeSetItemTemplateXMLFromString(hEle HELE, pXmlFile *uint16) bool {
	ret, _, _ := xTree_SetItemTemplateXMLFromString.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pXmlFile)))

	return ret == TRUE
}

/*
设置项模板文件,项选中状态.

参数:
	hEle 元素句柄.
	pStringXML 字符串指针.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XTreeSetItemTemplateXMLSelFromString(hEle HELE, pXmlFile *uint16) bool {
	ret, _, _ := xTree_SetItemTemplateXMLSelFromString.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pXmlFile)))

	return ret == TRUE
}

/*
设置是否绘制指定状态下项的背景.

参数:
	hEle 元素句柄.
	nFlags 标志位 List_drawItemBk_flags_.
*/
func XTreeSetDrawItemBkFlags(hEle HELE, nFlags int) {
	xTree_SetDrawItemBkFlags.Call(
		uintptr(hEle),
		uintptr(nFlags))
}

/*
设置项用户数据.

参数:
	hEle 元素句柄.
	nID 项ID.
	nUserData 用户数据.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XTreeSetItemData(hEle HELE, nID int, nUserData int) bool {
	ret, _, _ := xTree_SetItemData.Call(
		uintptr(hEle),
		uintptr(nID),
		uintptr(nUserData))

	return ret == TRUE

}

/*
获取项用户数据.

参数:
	hEle 元素句柄.
	nID 想ID.
返回:
	项用户数据.
*/
func XTreeGetItemData(hEle HELE, nID int) int {
	ret, _, _ := xTree_GetItemData.Call(
		uintptr(hEle),
		uintptr(nID))

	return int(ret)
}

/*
设置选择项.

参数:
	hEle 元素句柄.
	nID 项ID.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XTreeSetSelectItem(hEle HELE, nID int) bool {
	ret, _, _ := xTree_SetSelectItem.Call(
		uintptr(hEle),
		uintptr(nID))

	return ret == TRUE
}

/*
获取选择项.

参数:
	hEle 元素句柄.
返回:
	项ID.
*/
func XTreeGetSelectItem(hEle HELE) int {
	ret, _, _ := xTree_GetSelectItem.Call(uintptr(hEle))

	return int(ret)
}

/*
判断项是否展开.

参数:
	hEle 元素句柄.
	nID 项ID.
返回:
	如果展开返回TRUE否则返回FALSE.
*/
func XTreeIsExpand(hEle HELE, nID int) bool {
	ret, _, _ := xTree_IsExpand.Call(
		uintptr(hEle),
		uintptr(nID))

	return ret == TRUE
}

/*
展开项.

参数:
	hEle 元素句柄.
	nID 项ID.
	bExpand 是否展开.
返回:
	成功返回TRUE,如果项已经展开或失败返回FALSE.
*/
func XTreeExpandItem(hEle HELE, nID int, bExpand bool) bool {
	ret, _, _ := xTree_ExpandItem.Call(
		uintptr(hEle),
		uintptr(nID),
		uintptr(BoolToBOOL(bExpand)))

	return ret == TRUE
}

/*
检测坐标点所在项.

参数:
	hEle 元素句柄.
	pPt 坐标点.
返回:
	项ID.
*/
func XTreeHitTest(hEle HELE, pPt *POINT) int {
	ret, _, _ := xTree_HitTest.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)))

	return int(ret)
}

/*
检测坐标点所在项,自动添加滚动视图偏移坐标.

参数:
	hEle 元素句柄.
	pPt 坐标点.
返回:
	项ID.
*/
func XTreeHitTestOffet(hEle HELE, pPt *POINT) int {
	ret, _, _ := xTree_HitTestOffet.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)))

	return int(ret)
}

/*
获取第一个子项.

参数:
	hEle 元素句柄.
	nID 项ID.
返回:
	返回项ID,失败返回XC_ID_ERROR.
*/
func XTreeGetFirstChildItem(hEle HELE, nID int) int {
	ret, _, _ := xTree_GetFirstChildItem.Call(
		uintptr(hEle),
		uintptr(nID))

	return int(ret)
}

/*
获取下一个兄弟项.

参数:
	hEle 元素句柄.
	nID 项ID.
返回:
	返回下一个兄弟项ID.
*/
func XTreeGetNextSiblingItem(hEle HELE, nID int) int {
	ret, _, _ := xTree_GetNextSiblingItem.Call(
		uintptr(hEle),
		uintptr(nID))

	return int(ret)
}

/*
获取父项.

参数:
	hEle 元素句柄.
	nID 项ID.
返回:
	返回父项ID,错误返回-1.
*/
func XTreeGetParentItem(hEle HELE, nID int) int {
	ret, _, _ := xTree_GetParentItem.Call(
		uintptr(hEle),
		uintptr(nID))

	return int(ret)
}

/*
绑定数据适配器.

参数:
	hEle 元素句柄.
	hAdapter 数据适配器句柄, XAdapterTree.
*/
func XTreeBindAdapter(hEle HELE, hAdapter HXCGUI) {
	xTree_BindAdapter.Call(
		uintptr(hEle),
		uintptr(hAdapter))
}

/*
获取数据适配器.

参数:
	hEle 元素句柄.
返回:
	返回数据适配器句柄.
*/
func XTreeGetAdapter(hEle HELE) HXCGUI {
	ret, _, _ := xTree_GetAdapter.Call(uintptr(hEle))

	return HXCGUI(ret)
}

/*
设置缩进大小.

参数:
	hEle 元素句柄.
	nWidth 缩进宽度.
*/
func XTreeSetIndentation(hEle HELE, nWidth int) {
	xTree_SetIndentation.Call(
		uintptr(hEle),
		uintptr(nWidth))
}

/*
获取缩进值.

参数:
	hEle 元素句柄.
返回:
	返回缩进值大小.
*/
func XTreeGetIndentation(hEle HELE) int {
	ret, _, _ := xTree_GetIndentation.Call(uintptr(hEle))

	return int(ret)
}

/*
设置项默认高度.

参数:
	hEle 元素句柄.
	nHeight 高度.
	nSelHeight 选中时高度.
*/
func XTreeSetItemHeightDefault(hEle HELE, nHeight int, nSelHeight int) {
	xTree_SetItemHeightDefault.Call(
		uintptr(hEle),
		uintptr(nHeight),
		uintptr(nSelHeight))
}

/*
获取项默认高度.

参数:
	hEle 元素句柄.
	pHeight 接收返回高度.
	pSelHeight 接收返回值,当项选中时的高度.
*/
func XTreeGetItemHeightDefault(hEle HELE, pHeight *int, pSelHeight *int) {
	xTree_GetItemHeightDefault.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pHeight)),
		uintptr(unsafe.Pointer(pSelHeight)))
}

/*
设置指定项高度.

参数:
	hEle 元素句柄.
	nID 项ID.
	nHeight 高度.
	nSelHeight 选中时高度.
*/
func XTreeSetItemHeight(hEle HELE, nID int, nHeight int, nSelHeight int) {
	xTree_SetItemHeight.Call(
		uintptr(hEle),
		uintptr(nHeight),
		uintptr(nSelHeight))
}

/*
获取指定项高度.

参数:
	hEle 元素句柄.
	nID 项ID.
	pHeight 接收返回高度.
	pSelHeight 接收返回值,当项选中时的高度.
*/
func XTreeGetItemHeight(hEle HELE, nID int, pHeight *int, pSelHeight *int) {
	xTree_GetItemHeight.Call(
		uintptr(hEle),
		uintptr(nID),
		uintptr(unsafe.Pointer(pHeight)),
		uintptr(unsafe.Pointer(pSelHeight)))

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
func XTreeAddItemBkBorder(hEle HELE, nState Tree_item_state_, color COLORREF, alpha byte, width int) {
	xTree_AddItemBkBorder.Call(
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
func XTreeAddItemBkFill(hEle HELE, nState Tree_item_state_, color COLORREF, alpha byte) {
	xTree_AddItemBkFill.Call(
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
func XTreeAddItemBkImage(hEle HELE, nState Tree_item_state_, hImage HIMAGE) {
	xTree_AddItemBkImage.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(hImage))

}

/*
获取项背景内容数量.

参数:
	hEle 元素句柄.
	nState 项状态.
返回:
	项背景内容数量.
*/
func XTreeGetItemBkInfoCount(hEle HELE, nState Tree_item_state_) int {
	ret, _, _ := xTree_GetItemBkInfoCount.Call(
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
func XTreeClearItemBkInfo(hEle HELE, nState Tree_item_state_) {
	xTree_ClearItemBkInfo.Call(
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
func XTreeGetItemBkInfoManager(hEle HELE, nState Tree_item_state_) HBKINFOM {
	ret, _, _ := xTree_GetItemBkInfoManager.Call(
		uintptr(hEle),
		uintptr(nState))

	return HBKINFOM(ret)
}

/*
通过模板项ID,获取实例化模板项ID对应的对象句柄.

参数:
	hEle 元素句柄.
	nID 树项ID.
	nTempItemID 模板项ID.
返回:
	成功返回对象句柄,否则返回NULL.
*/
func XTreeGetTemplateObject(hEle HELE, nID int, nTempItemID int) HXCGUI {
	ret, _, _ := xTree_GetTemplateObject.Call(
		uintptr(hEle),
		uintptr(nID),
		uintptr(nTempItemID))

	return HXCGUI(ret)
}

/*
获取当前对象所在模板实例,属于列表树中哪一个项.

参数:
	hEle 元素句柄.
	hXCGUI 对象句柄, UI元素句柄或形状对象句柄..
返回:
	成功返回项ID, 否则返回XC_ID_ERROR.
*/
func XTreeGetItemIDFromHXCGUI(hEle HELE, hXCGUI HXCGUI) int {
	ret, _, _ := xTree_GetItemIDFromHXCGUI.Call(
		uintptr(hEle),
		uintptr(hXCGUI))

	return int(ret)
}
