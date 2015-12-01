package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xAdapterListView_Create               *syscall.Proc
	xAdapterListView_Group_AddColumn      *syscall.Proc
	xAdapterListView_Group_AddItemText    *syscall.Proc
	xAdapterListView_Group_AddItemTextEx  *syscall.Proc
	xAdapterListView_Group_AddItemImage   *syscall.Proc
	xAdapterListView_Group_AddItemImageEx *syscall.Proc
	xAdapterListView_Group_SetText        *syscall.Proc
	xAdapterListView_Group_SetTextEx      *syscall.Proc
	xAdapterListView_Group_SetImage       *syscall.Proc
	xAdapterListView_Group_SetImageEx     *syscall.Proc
	xAdapterListView_Item_AddColumn       *syscall.Proc
	xAdapterListView_Item_AddItemText     *syscall.Proc
	xAdapterListView_Item_AddItemTextEx   *syscall.Proc
	xAdapterListView_Item_AddItemImage    *syscall.Proc
	xAdapterListView_Item_AddItemImageEx  *syscall.Proc
	xAdapterListView_Group_DeleteItem     *syscall.Proc
	xAdapterListView_Item_DeleteItem      *syscall.Proc
	xAdapterListView_DeleteAll            *syscall.Proc
	xAdapterListView_Item_GetTextEx       *syscall.Proc
	xAdapterListView_Item_GetImageEx      *syscall.Proc
	xAdapterListView_Item_SetText         *syscall.Proc
	xAdapterListView_Item_SetTextEx       *syscall.Proc
	xAdapterListView_Item_SetImage        *syscall.Proc
	xAdapterListView_Item_SetImageEx      *syscall.Proc
)

func init() {
	// Functions
	xAdapterListView_Create = xcDLL.MustFindProc("XAdapterListView_Create")
	xAdapterListView_Group_AddColumn = xcDLL.MustFindProc("XAdapterListView_Group_AddColumn")
	xAdapterListView_Group_AddItemText = xcDLL.MustFindProc("XAdapterListView_Group_AddItemText")
	xAdapterListView_Group_AddItemTextEx = xcDLL.MustFindProc("XAdapterListView_Group_AddItemTextEx")
	xAdapterListView_Group_AddItemImage = xcDLL.MustFindProc("XAdapterListView_Group_AddItemImage")
	xAdapterListView_Group_AddItemImageEx = xcDLL.MustFindProc("XAdapterListView_Group_AddItemImageEx")
	xAdapterListView_Group_SetText = xcDLL.MustFindProc("XAdapterListView_Group_SetText")
	xAdapterListView_Group_SetTextEx = xcDLL.MustFindProc("XAdapterListView_Group_SetTextEx")
	xAdapterListView_Group_SetImage = xcDLL.MustFindProc("XAdapterListView_Group_SetImage")
	xAdapterListView_Group_SetImageEx = xcDLL.MustFindProc("XAdapterListView_Group_SetImageEx")
	xAdapterListView_Item_AddColumn = xcDLL.MustFindProc("XAdapterListView_Item_AddColumn")
	xAdapterListView_Item_AddItemText = xcDLL.MustFindProc("XAdapterListView_Item_AddItemText")
	xAdapterListView_Item_AddItemTextEx = xcDLL.MustFindProc("XAdapterListView_Item_AddItemTextEx")
	xAdapterListView_Item_AddItemImage = xcDLL.MustFindProc("XAdapterListView_Item_AddItemImage")
	xAdapterListView_Item_AddItemImageEx = xcDLL.MustFindProc("XAdapterListView_Item_AddItemImageEx")
	xAdapterListView_Group_DeleteItem = xcDLL.MustFindProc("XAdapterListView_Group_DeleteItem")
	xAdapterListView_Item_DeleteItem = xcDLL.MustFindProc("XAdapterListView_Item_DeleteItem")
	xAdapterListView_DeleteAll = xcDLL.MustFindProc("XAdapterListView_DeleteAll")
	xAdapterListView_Item_GetTextEx = xcDLL.MustFindProc("XAdapterListView_Item_GetTextEx")
	xAdapterListView_Item_GetImageEx = xcDLL.MustFindProc("XAdapterListView_Item_GetImageEx")
	xAdapterListView_Item_SetText = xcDLL.MustFindProc("XAdapterListView_Item_SetText")
	xAdapterListView_Item_SetTextEx = xcDLL.MustFindProc("XAdapterListView_Item_SetTextEx")
	xAdapterListView_Item_SetImage = xcDLL.MustFindProc("XAdapterListView_Item_SetImage")
	xAdapterListView_Item_SetImageEx = xcDLL.MustFindProc("XAdapterListView_Item_SetImageEx")
}

/*
创建列表视元素数据适配器.

返回:
	返回数据适配器句柄.
*/
func XAdapterListViewCreate() HXCGUI {
	ret, _, _ := xAdapterListView_Create.Call()

	return HXCGUI(ret)
}

/*
组操作,添加数据列.

参数:
	hAdapter 数据适配器句柄.
	pName 名称.
返回:
	返回列索引.
*/
func XAdapterListViewGroupAddColumn(hAdapter HXCGUI, pName *uint16) int {
	ret, _, _ := xAdapterListView_Group_AddColumn.Call(
		uintptr(hAdapter),
		uintptr(unsafe.Pointer(pName)))

	return int(ret)
}

/*
组操作,添加组,数据默认填充到第一列.

参数:
	hAdapter 数据适配器句柄.
	pValue 值.
返回:
	返回组索引.
*/
func XAdapterListViewGroupAddItemText(hAdapter HXCGUI, pValue *uint16) int {
	ret, _, _ := xAdapterListView_Group_AddItemText.Call(
		uintptr(hAdapter),
		uintptr(unsafe.Pointer(pValue)))

	return int(ret)
}

/*
组操作,添加组,数据填充指定列.

参数:
	hAdapter 数据适配器句柄.
	pName 列名.
	pValue 值.
返回:
	返回组索引.
*/
func XAdapterListViewGroupAddItemTextEx(hAdapter HXCGUI, pName, pValue *uint16) int {
	ret, _, _ := xAdapterListView_Group_AddItemTextEx.Call(
		uintptr(hAdapter),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(pValue)))

	return int(ret)
}

/*
组操作,添加组,数据默认填充第一列.

参数:
	hAdapter 数据适配器句柄.
	hImage 图片句柄.
返回:
	返回组索引.
*/
func XAdapterListViewGroupAddItemImage(hAdapter HXCGUI, hImage HIMAGE) int {
	ret, _, _ := xAdapterListView_Group_AddItemImage.Call(
		uintptr(hAdapter),
		uintptr(hImage))

	return int(ret)
}

/*
组操作,添加组,数据填充指定列.

参数:
	hAdapter 数据适配器句柄.
	pName 列名.
	hImage 图片句柄.
返回:
	返回组索引.
*/
func XAdapterListViewGroupAddItemImageEx(hAdapter HXCGUI, pName *uint16, hImage HIMAGE) int {
	ret, _, _ := xAdapterListView_Group_AddItemImageEx.Call(
		uintptr(hAdapter),
		uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return int(ret)
}

/*
组操作,设置指定项数据.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iColumn 列索引.
	pValue 值.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterListViewGroupSetText(hAdapter HXCGUI, iGroup, iColumn int, pValue *uint16) bool {
	ret, _, _ := xAdapterListView_Group_SetText.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iColumn),
		uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
组操作,设置指定项数据.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	pName 字段名.
	pValue 值.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterListViewGroupSetTextEx(hAdapter HXCGUI, iGroup int, pName, pValue *uint16) bool {
	ret, _, _ := xAdapterListView_Group_SetTextEx.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
组操作,设置指定项数据.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iColumn 列索引.
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterListViewGroupSetImage(hAdapter HXCGUI, iGroup, iColumn int, hImage HIMAGE) bool {
	ret, _, _ := xAdapterListView_Group_SetImage.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iColumn),
		uintptr(hImage))

	return ret == TRUE
}

/*
组操作,设置指定项数据.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	pName 字段名.
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterListViewGroupSetImageEx(hAdapter HXCGUI, iGroup int, pName *uint16, hImage HIMAGE) bool {
	ret, _, _ := xAdapterListView_Group_SetImageEx.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return ret == TRUE
}

/*
项操作,添加列.

参数:
	hAdapter 数据适配器句柄.
	pName 名称.
返回:
	返回列索引.
*/
func XAdapterListViewItemAddColumn(hAdapter HXCGUI, pName *uint16) int {
	ret, _, _ := xAdapterListView_Item_AddColumn.Call(
		uintptr(hAdapter),
		uintptr(unsafe.Pointer(pName)))

	return int(ret)
}

/*
项操作,添加项.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	pValue 值.
返回:
	返回项索引.
*/
func XAdapterListViewItemAddItemText(hAdapter HXCGUI, iGroup int, pValue *uint16) int {
	ret, _, _ := xAdapterListView_Item_AddItemText.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(unsafe.Pointer(pValue)))

	return int(ret)
}

/*
项操作,添加项,数据填充指定列.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	pName 列名称.
	pValue 值.
返回:
	返回项索引.
*/
func XAdapterListViewItemAddItemTextEx(hAdapter HXCGUI, iGroup int, pName, pValue *uint16) int {
	ret, _, _ := xAdapterListView_Item_AddItemTextEx.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(pValue)))

	return int(ret)
}

/*
项操作,添加项.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	hImage 图片句柄.
返回:
	返回项索引.
*/
func XAdapterListViewItemAddItemImage(hAdapter HXCGUI, iGroup int, hImage HIMAGE) int {
	ret, _, _ := xAdapterListView_Item_AddItemImage.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(hImage))

	return int(ret)
}

/*
项操作,添加项,填充指定列数据.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	pName 列名称.
	hImage 图片句柄.
返回:
	返回项索引.
*/
func XAdapterListViewItemAddItemImageEx(hAdapter HXCGUI, iGroup int, pName *uint16, hImage HIMAGE) int {
	ret, _, _ := xAdapterListView_Item_AddItemImageEx.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return int(ret)
}

/*
删除组,自动删除子项.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterListViewGroupDeleteItem(hAdapter HXCGUI, iGroup int) bool {
	ret, _, _ := xAdapterListView_Group_DeleteItem.Call(
		uintptr(hAdapter),
		uintptr(iGroup))

	return ret == TRUE
}

/*
删除项.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iItem 项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterListViewItemDeleteItem(hAdapter HXCGUI, iGroup, iItem int) bool {
	ret, _, _ := xAdapterListView_Item_DeleteItem.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iItem))

	return ret == TRUE
}

/*
删除所有.

参数:
	hAdapter 数据适配器句柄.
*/
func XAdapterListViewDeleteAll(hAdapter HXCGUI) {
	xAdapterListView_DeleteAll.Call(uintptr(hAdapter))
}

/*
项操作,获取项文本内容.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iItem 项索引.
	pName 列名.
	pOut 接收内容缓冲区.
	nOutLen 缓冲区大小,字符为单位.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterListViewItemGetTextEx(hAdapter HXCGUI, iGroup, iItem int, pName, pOut *uint16, nOutLen int) bool {
	ret, _, _ := xAdapterListView_Item_GetTextEx.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))

	return ret == TRUE
}

/*
项操作,获取项图片句柄.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iItem 项索引.
	pName 列名.
返回:
	返回图片句柄.
*/
func XAdapterListViewItemGetImageEx(hAdapter HXCGUI, iGroup, iItem int, pName *uint16) HIMAGE {
	ret, _, _ := xAdapterListView_Item_GetImageEx.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pName)))

	return HIMAGE(ret)
}

/*
项操作,数据填充指定列.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iItem 项索引.
	iColumn 列索引.
	pValue 值.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterListViewItemSetText(hAdapter HXCGUI, iGroup, iItem, iColumn int, pValue *uint16) bool {
	ret, _, _ := xAdapterListView_Item_SetText.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iItem),
		uintptr(iColumn),
		uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
项操作,数据填充指定列.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iItem 项索引.
	pName 列名称.
	pValue 值.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterListViewItemSetTextEx(hAdapter HXCGUI, iGroup, iItem int, pName, pValue *uint16) bool {
	ret, _, _ := xAdapterListView_Item_SetTextEx.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
项操作,数据填充指定列.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iItem 项索引.
	iColumn 列索引.
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterListViewItemSetImage(hAdapter HXCGUI, iGroup, iItem, iColumn int, hImage HIMAGE) bool {
	ret, _, _ := xAdapterListView_Item_SetImage.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iItem),
		uintptr(iColumn),
		uintptr(hImage))

	return ret == TRUE
}

/*
项操作,数据填充指定列.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iItem 项索引.
	pName 列名称.
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterListViewItemSetImageEx(hAdapter HXCGUI, iGroup, iItem int, pName *uint16, hImage HIMAGE) bool {
	ret, _, _ := xAdapterListView_Item_SetImageEx.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return ret == TRUE
}
