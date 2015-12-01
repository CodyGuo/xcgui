package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xAdapterTable_Create            *syscall.Proc
	xAdapterTable_AddColumn         *syscall.Proc
	xAdapterTable_SetColumn         *syscall.Proc
	xAdapterTable_AddItemText       *syscall.Proc
	xAdapterTable_AddItemTextEx     *syscall.Proc
	xAdapterTable_AddItemImage      *syscall.Proc
	xAdapterTable_AddItemImageEx    *syscall.Proc
	xAdapterTable_InsertItemText    *syscall.Proc
	xAdapterTable_InsertItemTextEx  *syscall.Proc
	xAdapterTable_InsertItemImage   *syscall.Proc
	xAdapterTable_InsertItemImageEx *syscall.Proc
	xAdapterTable_SetItemText       *syscall.Proc
	xAdapterTable_SetItemTextEx     *syscall.Proc
	xAdapterTable_SetItemImage      *syscall.Proc
	xAdapterTable_SetItemImageEx    *syscall.Proc
	xAdapterTable_DeleteItem        *syscall.Proc
	xAdapterTable_DeleteItemEx      *syscall.Proc
	xAdapterTable_DeleteItemAll     *syscall.Proc
	xAdapterTable_DeleteColumnAll   *syscall.Proc
	xAdapterTable_GetCount          *syscall.Proc
	xAdapterTable_GetCountColumn    *syscall.Proc
	xAdapterTable_GetItemText       *syscall.Proc
	xAdapterTable_GetItemImage      *syscall.Proc
	xAdapterTable_GetItemTextEx     *syscall.Proc
	xAdapterTable_GetItemImageEx    *syscall.Proc
)

func init() {
	// Functions
	xAdapterTable_Create = xcDLL.MustFindProc("XAdapterTable_Create")
	xAdapterTable_AddColumn = xcDLL.MustFindProc("XAdapterTable_AddColumn")
	xAdapterTable_SetColumn = xcDLL.MustFindProc("XAdapterTable_SetColumn")
	xAdapterTable_AddItemText = xcDLL.MustFindProc("XAdapterTable_AddItemText")
	xAdapterTable_AddItemTextEx = xcDLL.MustFindProc("XAdapterTable_AddItemTextEx")
	xAdapterTable_AddItemImage = xcDLL.MustFindProc("XAdapterTable_AddItemImage")
	xAdapterTable_AddItemImageEx = xcDLL.MustFindProc("XAdapterTable_AddItemImageEx")
	xAdapterTable_InsertItemText = xcDLL.MustFindProc("XAdapterTable_InsertItemText")
	xAdapterTable_InsertItemTextEx = xcDLL.MustFindProc("XAdapterTable_InsertItemTextEx")
	xAdapterTable_InsertItemImage = xcDLL.MustFindProc("XAdapterTable_InsertItemImage")
	xAdapterTable_InsertItemImageEx = xcDLL.MustFindProc("XAdapterTable_InsertItemImageEx")
	xAdapterTable_SetItemText = xcDLL.MustFindProc("XAdapterTable_SetItemText")
	xAdapterTable_SetItemTextEx = xcDLL.MustFindProc("XAdapterTable_SetItemTextEx")
	xAdapterTable_SetItemImage = xcDLL.MustFindProc("XAdapterTable_SetItemImage")
	xAdapterTable_SetItemImageEx = xcDLL.MustFindProc("XAdapterTable_SetItemImageEx")
	xAdapterTable_DeleteItem = xcDLL.MustFindProc("XAdapterTable_DeleteItem")
	xAdapterTable_DeleteItemEx = xcDLL.MustFindProc("XAdapterTable_DeleteItemEx")
	xAdapterTable_DeleteItemAll = xcDLL.MustFindProc("XAdapterTable_DeleteItemAll")
	xAdapterTable_DeleteColumnAll = xcDLL.MustFindProc("XAdapterTable_DeleteColumnAll")
	xAdapterTable_GetCount = xcDLL.MustFindProc("XAdapterTable_GetCount")
	xAdapterTable_GetCountColumn = xcDLL.MustFindProc("XAdapterTable_GetCountColumn")
	xAdapterTable_GetItemText = xcDLL.MustFindProc("XAdapterTable_GetItemText")
	xAdapterTable_GetItemImage = xcDLL.MustFindProc("XAdapterTable_GetItemImage")
	xAdapterTable_GetItemTextEx = xcDLL.MustFindProc("XAdapterTable_GetItemTextEx")
	xAdapterTable_GetItemImageEx = xcDLL.MustFindProc("XAdapterTable_GetItemImageEx")
}

/*
创建列表框元素数据适配器.

返回:
	返回数据适配器句柄.
*/
func XAdapterTableCreate() HXCGUI {
	ret, _, _ := xAdapterTable_Create.Call()

	return HXCGUI(ret)
}

/*
添加数据列.

参数:
	hAdapter 数据适配器句柄.
	pName 名称.
返回:
	返回列索引.
*/
func XAdapterTableAddColumn(hAdapter HXCGUI, pName *uint16) int {
	ret, _, _ := xAdapterTable_AddColumn.Call(
		uintptr(hAdapter),
		uintptr(unsafe.Pointer(pName)))

	return int(ret)
}

/*
设置列.

参数:
	hAdapter 数据适配器句柄.
	pColName 列名,多个列名用逗号分开.
返回:
	返回列数量. 注解:例如: XAdapterTable_SetColumn(hAdapter, L"name1,name2,mame3");
*/
func XAdapterTableSetColumn(hAdapter HXCGUI, pColName *uint16) int {
	ret, _, _ := xAdapterTable_SetColumn.Call(
		uintptr(hAdapter),
		uintptr(unsafe.Pointer(pColName)))

	return int(ret)
}

/*
添加数据项,默认值放到第一列.

参数:
	hAdapter 数据适配器句柄.
	pValue 值.
返回:
	返回项索引值.
*/
func XAdapterTableAddItemText(hAdapter HXCGUI, pValue *uint16) int {
	ret, _, _ := xAdapterTable_AddItemText.Call(
		uintptr(hAdapter),
		uintptr(unsafe.Pointer(pValue)))

	return int(ret)
}

/*
添加数据项,填充指定列数据.

参数:
	hAdapter 数据适配器句柄.
	pName 列名.
	pValue 值.
返回:
	返回项索引.
*/
func XAdapterTableAddItemTextEx(hAdapter HXCGUI, pName, pValue *uint16) int {
	ret, _, _ := xAdapterTable_AddItemTextEx.Call(
		uintptr(hAdapter),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(pValue)))

	return int(ret)
}

/*
添加数据项,默认值放到第一列.

参数:
	hAdapter 数据适配器句柄.
	hImage 图片句柄.
返回:
	返回项索引值.
*/
func XAdapterTableAddItemImage(hAdapter HXCGUI, hImage HIMAGE) int {
	ret, _, _ := xAdapterTable_AddItemImage.Call(
		uintptr(hAdapter),
		uintptr(hImage))

	return int(ret)
}

/*
添加数据项,并填充指定列数据.

参数:
	hAdapter 数据适配器句柄.
	pName 列名.
	hImage 图片句柄.
返回:

*/
func XAdapterTableAddItemImageEx(hAdapter HXCGUI, pName *uint16, hImage HIMAGE) int {
	ret, _, _ := xAdapterTable_AddItemImageEx.Call(
		uintptr(hAdapter),
		uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return int(ret)
}

/*
插入数据项,填充第一列数据.

参数:
	hAdapter 数据适配器句柄.
	iItem 插入位置索引.
	pValue 值.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTableInsertItemText(hAdapter HXCGUI, iItem int, pValue *uint16) bool {
	ret, _, _ := xAdapterTable_InsertItemText.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
插入数据项,并填充指定列数据.

参数:
	hAdapter 数据适配器句柄.
	iItem 插入位置索引.
	pName 列名.
	pValue 值.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTableInsertItemTextEx(hAdapter HXCGUI, iItem int, pName, pValue *uint16) bool {
	ret, _, _ := xAdapterTable_InsertItemTextEx.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
插入数据项,填充第一列数据.

参数:
	hAdapter 数据适配器句柄.
	iItem 插入位置索引.
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTableInsertItemImage(hAdapter HXCGUI, iItem int, hImage HIMAGE) bool {
	ret, _, _ := xAdapterTable_InsertItemImage.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(hImage))

	return ret == TRUE
}

/*
插入数据项,并填充指定列数据.

参数:
	hAdapter 数据适配器句柄.
	iItem 插入位置索引.
	pName 列名.
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTableInsertItemImageEx(hAdapter HXCGUI, iItem int, pName *uint16, hImage HIMAGE) bool {
	ret, _, _ := xAdapterTable_InsertItemImageEx.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return ret == TRUE
}

/*
设置项数据.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
	iColumn 列索引.
	pValue 值.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTableSetItemText(hAdapter HXCGUI, iItem, iColumn int, pValue *uint16) bool {
	ret, _, _ := xAdapterTable_SetItemText.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(iColumn),
		uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
设置项数据.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
	pName 列名.
	pValue 值.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTableSetItemTextEx(hAdapter HXCGUI, iItem int, pName, pValue *uint16) bool {
	ret, _, _ := xAdapterTable_SetItemTextEx.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
设置项数据.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
	iColumn 列索引.
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTableSetItemImage(hAdapter HXCGUI, iItem, iColumn int, hImage HIMAGE) bool {
	ret, _, _ := xAdapterTable_SetItemImage.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(iColumn),
		uintptr(hImage))

	return ret == TRUE
}

/*
设置项数据.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
	pName 列名.
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTableSetItemImageEx(hAdapter HXCGUI, iItem int, pName *uint16, hImage HIMAGE) bool {
	ret, _, _ := xAdapterTable_SetItemImageEx.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return ret == TRUE
}

/*
删除项.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTableDeleteItem(hAdapter HXCGUI, iItem int) bool {
	ret, _, _ := xAdapterTable_DeleteItem.Call(
		uintptr(hAdapter),
		uintptr(iItem))

	return ret == TRUE
}

/*
删除多个项.

参数:
	hAdapter 数据适配器句柄.
	iItem 项起始索引.
	nCount 删除项数量.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTableDeleteItemEx(hAdapter HXCGUI, iItem, nCount int) bool {
	ret, _, _ := xAdapterTable_DeleteItemEx.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(nCount))

	return ret == TRUE
}

/*
删除所有项.

参数:
	hAdapter 数据适配器句柄.
*/
func XAdapterTableDeleteItemAll(hAdapter HXCGUI) {
	xAdapterTable_DeleteItemAll.Call(uintptr(hAdapter))
}

/*
删除所有列,并且清空所有数据.

参数:
	hAdapter 数据适配器句柄.
*/
func XAdapterTableDeleteColumnAll(hAdapter HXCGUI) {
	xAdapterTable_DeleteColumnAll.Call(uintptr(hAdapter))
}

/*
获取项数量.

参数:
	hAdapter 数据适配器句柄.
返回:
	返回数量.
*/
func XAdapterTableGetCount(hAdapter HXCGUI) int {
	ret, _, _ := xAdapterTable_GetCount.Call(uintptr(hAdapter))

	return int(ret)
}

/*
获取列数量.

参数:
	hAdapter 数据适配器句柄.
返回:
	返回列数量.
*/
func XAdapterTableGetCountColumn(hAdapter HXCGUI) int {
	ret, _, _ := xAdapterTable_GetCountColumn.Call(uintptr(hAdapter))

	return int(ret)
}

/*
获取项文本内容.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
	iColumn 列索引.
	pOut 接收内容缓冲区.
	nOutLen 缓冲区大小,字符为单位.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTableGetItemText(hAdapter HXCGUI, iItem, iColumn int, pOut *uint16, nOutLen int) bool {
	ret, _, _ := xAdapterTable_GetItemText.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(iColumn),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))

	return ret == TRUE
}

/*
获取项图片句柄.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
	iColumn 列索引.
返回:
	返回图片句柄.
*/
func XAdapterTableGetItemImage(hAdapter HXCGUI, iItem, iColumn int) HIMAGE {
	ret, _, _ := xAdapterTable_GetItemImage.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(iColumn))

	return HIMAGE(ret)
}

/*
获取项文本内容.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
	pName 列名.
	pOut 接收内容缓冲区.
	nOutLen 缓冲区大小,字符为单位.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTableGetItemTextEx(hAdapter HXCGUI, iItem int, pName, pOut *uint16, nOutLen int) bool {
	ret, _, _ := xAdapterTable_GetItemTextEx.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))

	return ret == TRUE
}

/*
获取项图片句柄.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
	pName 列名.
返回:
	返回图片句柄.
*/
func XAdapterTableGetItemImageEx(hAdapter HXCGUI, iItem int, pName *uint16) HIMAGE {
	ret, _, _ := xAdapterTable_GetItemImageEx.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pName)))

	return HIMAGE(ret)
}
