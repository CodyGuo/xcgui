package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xListView_Create                       *syscall.Proc
	xListView_BindAdapter                  *syscall.Proc
	xListView_GetAdapter                   *syscall.Proc
	xListView_SetItemTemplateXML           *syscall.Proc
	xListView_SetItemTemplateXMLFromString *syscall.Proc
	xListView_GetTemplateObject            *syscall.Proc
	xListView_GetTemplateObjectGroup       *syscall.Proc
	xListView_GetItemIDFromHXCGUI          *syscall.Proc
	xListView_HitTest                      *syscall.Proc
	xListView_HitTestOffset                *syscall.Proc
	xListView_EnableMultiSel               *syscall.Proc
	xListView_SetDrawItemBkFlags           *syscall.Proc
	xListView_SetSelectItem                *syscall.Proc
	xListView_GetSelectItem                *syscall.Proc
	xListView_GetSelectItemCount           *syscall.Proc
	xListView_GetSelectItemAll             *syscall.Proc
	xListView_SetSelectItemAll             *syscall.Proc
	xListView_CancelSelectItemAll          *syscall.Proc
	xListView_SetColumnSpace               *syscall.Proc
	xListView_SetRowSpace                  *syscall.Proc
	xListView_SetAlignSizeLeft             *syscall.Proc
	xListView_SetAlignSizeTop              *syscall.Proc
	xListView_SetItemSize                  *syscall.Proc
	xListView_GetItemSize                  *syscall.Proc
	xListView_SetGroupHeight               *syscall.Proc
	xListView_GetGroupHeight               *syscall.Proc
	xListView_AddItemBkBorder              *syscall.Proc
	xListView_AddItemBkFill                *syscall.Proc
	xListView_AddItemBkImage               *syscall.Proc
	xListView_GetItemBkInfoCount           *syscall.Proc
	xListView_ClearItemBkInfo              *syscall.Proc
	// xListView_GetItemBkInfoManager         *syscall.Proc
	xListView_ExpandGroup *syscall.Proc
)

func init() {
	// Functions
	xListView_Create = xcDLL.MustFindProc("XListView_Create")
	xListView_BindAdapter = xcDLL.MustFindProc("XListView_BindAdapter")
	xListView_GetAdapter = xcDLL.MustFindProc("XListView_GetAdapter")
	xListView_SetItemTemplateXML = xcDLL.MustFindProc("XListView_SetItemTemplateXML")
	xListView_SetItemTemplateXMLFromString = xcDLL.MustFindProc("XListView_SetItemTemplateXMLFromString")
	xListView_GetTemplateObject = xcDLL.MustFindProc("XListView_GetTemplateObject")
	xListView_GetTemplateObjectGroup = xcDLL.MustFindProc("XListView_GetTemplateObjectGroup")
	xListView_GetItemIDFromHXCGUI = xcDLL.MustFindProc("XListView_GetItemIDFromHXCGUI")
	xListView_HitTest = xcDLL.MustFindProc("XListView_HitTest")
	xListView_HitTestOffset = xcDLL.MustFindProc("XListView_HitTestOffset")
	xListView_EnableMultiSel = xcDLL.MustFindProc("XListView_EnableMultiSel")
	xListView_SetDrawItemBkFlags = xcDLL.MustFindProc("XListView_SetDrawItemBkFlags")
	xListView_SetSelectItem = xcDLL.MustFindProc("XListView_SetSelectItem")
	xListView_GetSelectItem = xcDLL.MustFindProc("XListView_GetSelectItem")
	xListView_GetSelectItemCount = xcDLL.MustFindProc("XListView_GetSelectItemCount")
	xListView_GetSelectItemAll = xcDLL.MustFindProc("XListView_GetSelectItemAll")
	xListView_SetSelectItemAll = xcDLL.MustFindProc("XListView_SetSelectItemAll")
	xListView_CancelSelectItemAll = xcDLL.MustFindProc("XListView_CancelSelectItemAll")
	xListView_SetColumnSpace = xcDLL.MustFindProc("XListView_SetColumnSpace")
	xListView_SetRowSpace = xcDLL.MustFindProc("XListView_SetRowSpace")
	xListView_SetAlignSizeLeft = xcDLL.MustFindProc("XListView_SetAlignSizeLeft")
	xListView_SetAlignSizeTop = xcDLL.MustFindProc("XListView_SetAlignSizeTop")
	xListView_SetItemSize = xcDLL.MustFindProc("XListView_SetItemSize")
	xListView_GetItemSize = xcDLL.MustFindProc("XListView_GetItemSize")
	xListView_SetGroupHeight = xcDLL.MustFindProc("XListView_SetGroupHeight")
	xListView_GetGroupHeight = xcDLL.MustFindProc("XListView_GetGroupHeight")
	xListView_AddItemBkBorder = xcDLL.MustFindProc("XListView_AddItemBkBorder")
	xListView_AddItemBkFill = xcDLL.MustFindProc("XListView_AddItemBkFill")
	xListView_AddItemBkImage = xcDLL.MustFindProc("XListView_AddItemBkImage")
	xListView_GetItemBkInfoCount = xcDLL.MustFindProc("XListView_GetItemBkInfoCount")
	xListView_ClearItemBkInfo = xcDLL.MustFindProc("XListView_ClearItemBkInfo")
	// xListView_GetItemBkInfoManager = xcDLL.MustFindProc("XListView_GetItemBkInfoManager")
	xListView_ExpandGroup = xcDLL.MustFindProc("XListView_ExpandGroup")
}

/*
创建列表视图元素.

参数:
	x 元素x坐标.
	y 元素y坐标.
	cx 宽度.
	cy 高度.
	hParent 父是窗口资源句柄或U I元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
返回:
	元素句柄.
*/
func XListViewCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xListView_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
绑定数据适配器.

参数:
	hEle 元素句柄.
	hAdapter 数据适配器 XAdapterListView.
*/
func XListViewBindAdapter(hEle HELE, hAdapter HXCGUI) {
	xListView_BindAdapter.Call(
		uintptr(hEle),
		uintptr(hAdapter))
}

/*
获取数据适配器.

参数:
	hEle 元素句柄.
返回:
	返回数据适配器.
*/
func XListViewGetAdapter(hEle HELE) HXCGUI {
	ret, _, _ := xListView_GetAdapter.Call(uintptr(hEle))

	return HXCGUI(ret)
}

/*
设置项模板文件.

参数:
	hEle 元素句柄.
	pXmlFile 文件名.*uint16
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListViewSetItemTemplateXML(hEle HELE, pXmlFile string) bool {
	ret, _, _ := xListView_SetItemTemplateXML.Call(
		uintptr(hEle),
		StringToUintPtr(pXmlFile))
	// uintptr(unsafe.Pointer(pXmlFile)))

	return ret == TRUE
}

/*
设置项布局模板.

参数:
	hEle 元素句柄.
	pStringXML 字符串指针. *uint16
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListViewSetItemTemplateXMLFromString(hEle HELE, pStringXML string) bool {
	ret, _, _ := xListView_SetItemTemplateXMLFromString.Call(
		uintptr(hEle),
		StringToUintPtr(pStringXML))
	// uintptr(unsafe.Pointer(pStringXML)))

	return ret == TRUE
}

/*
通过模板项ID,获取实例化模板项ID对应的对象句柄.

参数:
	hEle 元素句柄.
	iGroup 组索引.
	iItem 项索引.
	nTempItemID 模板项ID.
返回:
	成功返回对象句柄,否则返回NULL.
*/
func XListViewGetTemplateObject(hEle HELE, iGroup, iItem, nTempItemID int) HXCGUI {
	ret, _, _ := xListView_GetTemplateObject.Call(
		uintptr(hEle),
		uintptr(iGroup),
		uintptr(iItem),
		uintptr(nTempItemID))

	return HXCGUI(ret)
}

/*
通过模板项ID,获取实例化模板项ID对应的对象句柄.

参数:
	hEle 元素句柄.
	iGroup 组索引.
	nTempItemID 模板项ID.
返回:
	成功返回对象句柄,否则返回NULL.
*/
func XListViewGetTemplateObjectGroup(hEle HELE, iGroup, nTempItemID int) HXCGUI {
	ret, _, _ := xListView_GetTemplateObjectGroup.Call(
		uintptr(hEle),
		uintptr(iGroup),
		uintptr(nTempItemID))

	return HXCGUI(ret)
}

/*
获取当前对象所在模板实例,属于列表视中哪一个项.

参数:
	hEle 元素句柄.
	hXCGUI 对象句柄, UI元素句柄或形状对象句柄.
	piGroup 接收组索引.
	piItem 接收项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListViewGetItemIDFromHXCGUI(hEle HELE, hXCGUI HXCGUI, piGroup, piItem *int32) bool {
	ret, _, _ := xListView_GetItemIDFromHXCGUI.Call(
		uintptr(hEle),
		uintptr(hXCGUI),
		uintptr(unsafe.Pointer(piGroup)),
		uintptr(unsafe.Pointer(piItem)))

	return ret == TRUE
}

/*
检查坐标点所在项.

参数:
	hEle 元素句柄.
	pPt 坐标点.
	pOutGroup 接收组索引.
	pOutItem 接收项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListViewHitTest(hEle HELE, pPt *POINT, pOutGroup, pOutItem *int32) bool {
	ret, _, _ := xListView_HitTest.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)),
		uintptr(unsafe.Pointer(pOutGroup)),
		uintptr(unsafe.Pointer(pOutItem)))

	return ret == TRUE
}

/*
检查坐标点所在项,自动添加滚动视图偏移量.

参数:
	hEle 元素句柄.
	pPt 坐标点.
	pOutGroup 接收做索引.
	pOutItem 接收项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListViewHitTestOffset(hEle HELE, pPt *POINT, pOutGroup, pOutItem *int32) int {
	ret, _, _ := xListView_HitTestOffset.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)),
		uintptr(unsafe.Pointer(pOutGroup)),
		uintptr(unsafe.Pointer(pOutItem)))

	return int(ret)
}

/*
启用多选.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XListViewEnableMultiSel(hEle HELE, bEnable bool) {
	xListView_EnableMultiSel.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
设置是否绘制指定状态下项的背景.

参数:
	hEle 元素句柄.
	nFlags 标志位 list_drawItemBk_flags_.
*/
func XListViewSetDrawItemBkFlags(hEle HELE, nFlags List_drawItemBk_flags_) {
	xListView_SetDrawItemBkFlags.Call(
		uintptr(hEle),
		uintptr(nFlags))
}

/*
设置选择项.

参数:
	hEle 元素句柄.
	iGroup 组索引.
	iItem 项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListViewSetSelectItem(hEle HELE, iGroup, iItem int) bool {
	ret, _, _ := xListView_SetSelectItem.Call(
		uintptr(hEle),
		uintptr(iGroup),
		uintptr(iItem))

	return ret == TRUE
}

/*
获取选择项.

参数:
	hEle 元素句柄.
	piGroup 接收组索引.
	piItem 接收项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XListViewGetSelectItem(hEle HELE, piGroup, piItem *int32) bool {
	ret, _, _ := xListView_GetSelectItem.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(piGroup)),
		uintptr(unsafe.Pointer(piItem)))

	return ret == TRUE

}

/*
获取选择项数量.

参数:
	hEle 元素句柄.
返回:
	返回选择项数量.
*/
func XListViewGetSelectItemCount(hEle HELE) int {
	ret, _, _ := xListView_GetSelectItemCount.Call(uintptr(hEle))

	return int(ret)
}

/*
获取选择的项ID.

参数:
	hEle 元素句柄.
	pArray 数组,用来接收选择项ID.
	nArraySize 数组大小.
返回:
	返回接收项数量.
*/
func XListViewGetSelectItemAll(hEle HELE, pArray *Listview_item_id_i, nArraySize int) int {
	ret, _, _ := xListView_GetSelectItemAll.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pArray)),
		uintptr(nArraySize))

	return int(ret)
}

/*
选择所有的项.

参数:
	hEle 元素句柄.
*/
func XListViewSetSelectItemAll(hEle HELE) {
	xListView_SetSelectItemAll.Call(uintptr(hEle))

}

/*
取消选择所有项.

参数:
	hEle 元素句柄.
*/
func XListViewCancelSelectItemAll(hEle HELE) {
	xListView_CancelSelectItemAll.Call(uintptr(hEle))
}

/*
设置列间隔大小.

参数:
	hEle 元素句柄.
	space 间隔大小.
*/
func XListViewSetColumnSpace(hEle HELE, space int) {
	xListView_SetColumnSpace.Call(
		uintptr(hEle),
		uintptr(space))
}

/*
设置行间隔大小.

参数:
	hEle 元素句柄.
	space 间隔大小.
*/
func XListViewSetRowSpace(hEle HELE, space int) {
	xListView_SetRowSpace.Call(
		uintptr(hEle),
		uintptr(space))
}

/*
设置左边对齐大小.

参数:
	hEle 元素句柄.
	size 大小.
*/
func XListViewSetAlignSizeLeft(hEle HELE, size int) {
	xListView_SetAlignSizeLeft.Call(
		uintptr(hEle),
		uintptr(size))
}

/*
设置上边对齐大小.

参数:
	hEle 元素句柄.
	size 大小.
*/
func XListViewSetAlignSizeTop(hEle HELE, size int) {
	xListView_SetAlignSizeTop.Call(
		uintptr(hEle),
		uintptr(size))
}

/*
设置项大小.

参数:
	hEle 元素句柄.
	width 宽度.
	height 高度.
*/
func XListViewSetItemSize(hEle HELE, width, height int) {
	xListView_SetItemSize.Call(
		uintptr(hEle),
		uintptr(width),
		uintptr(height))
}

/*
获取项大小.

参数:
	hEle 元素句柄.
	pSize 接收返回大小.
*/
func XListViewGetItemSize(hEle HELE, pSize *SIZE) {
	xListView_GetItemSize.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pSize)))
}

/*
设置组高度.

参数:
	hEle 元素句柄.
	height 高度.
*/
func XListViewSetGroupHeight(hEle HELE, height int) {
	xListView_SetGroupHeight.Call(
		uintptr(hEle),
		uintptr(height))
}

/*
获取组高度.

参数:
	hEle 元素句柄.
返回:
	返回组高度.
*/
func XListViewGetGroupHeight(hEle HELE) int {
	ret, _, _ := xListView_GetGroupHeight.Call(uintptr(hEle))

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
func XListViewAddItemBkBorder(hEle HELE, nState List_item_state_, color COLORREF, alpha byte, width int) {
	xListView_AddItemBkBorder.Call(
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
func XListViewAddItemBkFill(hEle HELE, nState List_item_state_, color COLORREF, alpha byte) {
	xListView_AddItemBkFill.Call(
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
func XListViewAddItemBkImage(hEle HELE, nState List_item_state_, hImage HIMAGE) {
	xListView_AddItemBkImage.Call(
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
func XListViewGetItemBkInfoCount(hEle HELE, nState List_item_state_) int {
	ret, _, _ := xListView_GetItemBkInfoCount.Call(
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
func XListViewClearItemBkInfo(hEle HELE, nState List_item_state_) {
	xListView_ClearItemBkInfo.Call(
		uintptr(hEle),
		uintptr(nState))
}

/*
获取项背景内容管理器. 1.8.9.6

参数:
	hEle 元素句柄.
	nState 项状态.
返回:
	项背景内容管理器.
*/
// func XListViewGetItemBkInfoManager(hEle HELE, nState List_item_state_) HBKINFOM {
// 	ret, _, _ := xListView_GetItemBkInfoManager.Call(
// 		uintptr(hEle),
// 		uintptr(nState))

// 	return HBKINFOM(ret)
// }

/*
展开组.

参数:
	hEle 元素句柄.
	iGroup 组索引.
	bExpand 是否展开.
返回:
	成功返回TRUE否则返回FALSE,如果状态没有改变返回FALSE.
*/
func XListViewExpandGroup(hEle HELE, iGroup int, bExpand bool) bool {
	ret, _, _ := xListView_ExpandGroup.Call(
		uintptr(hEle),
		uintptr(iGroup),
		uintptr(BoolToBOOL(bExpand)))

	return ret == TRUE
}
