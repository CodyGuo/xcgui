package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xAdapterTree_Create            *syscall.Proc
	xAdapterTree_AddColumn         *syscall.Proc
	xAdapterTree_SetColumn         *syscall.Proc
	xAdapterTree_InsertItemText    *syscall.Proc
	xAdapterTree_InsertItemTextEx  *syscall.Proc
	xAdapterTree_InsertItemImage   *syscall.Proc
	xAdapterTree_InsertItemImageEx *syscall.Proc
	xAdapterTree_GetCount          *syscall.Proc
	xAdapterTree_GetCountColumn    *syscall.Proc
	xAdapterTree_SetItemText       *syscall.Proc
	xAdapterTree_SetItemTextEx     *syscall.Proc
	xAdapterTree_SetItemImage      *syscall.Proc
	xAdapterTree_SetItemImageEx    *syscall.Proc
	xAdapterTree_GetItemTextEx     *syscall.Proc
	xAdapterTree_GetItemImageEx    *syscall.Proc
	xAdapterTree_DeleteItem        *syscall.Proc
	xAdapterTree_DeleteItemAll     *syscall.Proc
	xAdapterTree_DeleteColumnAll   *syscall.Proc
)

func init() {
	// Functions
	xAdapterTree_Create = xcDLL.MustFindProc("XAdapterTree_Create")
	xAdapterTree_AddColumn = xcDLL.MustFindProc("XAdapterTree_AddColumn")
	xAdapterTree_SetColumn = xcDLL.MustFindProc("XAdapterTree_SetColumn")
	xAdapterTree_InsertItemText = xcDLL.MustFindProc("XAdapterTree_InsertItemText")
	xAdapterTree_InsertItemTextEx = xcDLL.MustFindProc("XAdapterTree_InsertItemTextEx")
	xAdapterTree_InsertItemImage = xcDLL.MustFindProc("XAdapterTree_InsertItemImage")
	xAdapterTree_InsertItemImageEx = xcDLL.MustFindProc("XAdapterTree_InsertItemImageEx")
	xAdapterTree_GetCount = xcDLL.MustFindProc("XAdapterTree_GetCount")
	xAdapterTree_GetCountColumn = xcDLL.MustFindProc("XAdapterTree_GetCountColumn")
	xAdapterTree_SetItemText = xcDLL.MustFindProc("XAdapterTree_SetItemText")
	xAdapterTree_SetItemTextEx = xcDLL.MustFindProc("XAdapterTree_SetItemTextEx")
	xAdapterTree_SetItemImage = xcDLL.MustFindProc("XAdapterTree_SetItemImage")
	xAdapterTree_SetItemImageEx = xcDLL.MustFindProc("XAdapterTree_SetItemImageEx")
	xAdapterTree_GetItemTextEx = xcDLL.MustFindProc("XAdapterTree_GetItemTextEx")
	xAdapterTree_GetItemImageEx = xcDLL.MustFindProc("XAdapterTree_GetItemImageEx")
	xAdapterTree_DeleteItem = xcDLL.MustFindProc("XAdapterTree_DeleteItem")
	xAdapterTree_DeleteItemAll = xcDLL.MustFindProc("XAdapterTree_DeleteItemAll")
	xAdapterTree_DeleteColumnAll = xcDLL.MustFindProc("XAdapterTree_DeleteColumnAll")
}

/*
创建树元素数据适配器.

返回:
	返回数据适配器句柄.
*/
func XAdapterTreeCreate() HXCGUI {
	ret, _, _ := xAdapterTree_Create.Call()

	return HXCGUI(ret)
}

/*
添加列.

参数:
	hAdapter 数据适配器句柄.
	pName 名称.
返回:
	返回索引值.
*/
func XAdapterTreeAddColumn(hAdapter HXCGUI, pName *uint16) int {
	ret, _, _ := xAdapterTree_AddColumn.Call(
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
	返回列数量. 注解:例如: XAdapterTree_SetColumn(hAdapter, L"name1,name2,mame3");
*/
func XAdapterTreeSetColumn(hAdapter HXCGUI, pColName *uint16) int {
	ret, _, _ := xAdapterTree_SetColumn.Call(
		uintptr(hAdapter),
		uintptr(unsafe.Pointer(pColName)))

	return int(ret)
}

/*
插入项,数据填充到第一列.

参数:
	hAdapter 数据适配器句柄.
	pValue 值.
	nParentID 父ID.
	insertID 插入位置ID.
返回:
	返回项ID值.
*/
func XAdapterTreeInsertItemText(hAdapter HXCGUI, pValue *uint16, nParentID, insertID int) int {
	ret, _, _ := xAdapterTree_InsertItemText.Call(
		uintptr(hAdapter),
		uintptr(unsafe.Pointer(pValue)),
		uintptr(nParentID),
		uintptr(insertID))

	return int(ret)
}

/*
插入项,数据填充到指定列.

参数:
	hAdapter 数据适配器句柄.
	pName 列名称.
	pValue 值.
	nParentID 父ID.
	insertID 插入位置ID.
返回:
	返回项ID值.
*/
func XAdapterTreeInsertItemTextEx(hAdapter HXCGUI, pName, pValue *uint16, nParentID, insertID int) int {
	ret, _, _ := xAdapterTree_InsertItemTextEx.Call(
		uintptr(hAdapter),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(pValue)),
		uintptr(nParentID),
		uintptr(insertID))

	return int(ret)
}

/*
插入项,数据填充到第一列.

参数:
	hAdapter 数据适配器句柄.
	hImage 图片句柄.
	nParentID 父ID.
	insertID 插入位置ID.
返回:
	返回项ID值.
*/
func XAdapterTreeInsertItemImage(hAdapter HXCGUI, hImage HIMAGE, nParentID, insertID int) int {
	ret, _, _ := xAdapterTree_InsertItemImage.Call(
		uintptr(hAdapter),
		uintptr(hImage),
		uintptr(nParentID),
		uintptr(insertID))

	return int(ret)
}

/*
插入项,数据填充到指定列.

参数:
	hAdapter 数据适配器句柄.
	pName 列名称.
	hImage 图片句柄.
	nParentID 父ID.
	insertID 插入位置ID.
返回:
	返回项ID值.
*/
func XAdapterTreeInsertItemImageEx(hAdapter HXCGUI, pName *uint16, hImage HIMAGE, nParentID, insertID int) int {
	ret, _, _ := xAdapterTree_InsertItemImageEx.Call(
		uintptr(hAdapter),
		uintptr(unsafe.Pointer(pName)),
		uintptr(hImage),
		uintptr(nParentID),
		uintptr(insertID))

	return int(ret)
}

/*
获取项数量.

参数:
	hAdapter 数据适配器句柄.
返回:
	返回数量.
*/
func XAdapterTreeGetCount(hAdapter HXCGUI) int {
	ret, _, _ := xAdapterTree_GetCount.Call(uintptr(hAdapter))

	return int(ret)
}

/*
获取列数量.

参数:
	hAdapter 数据适配器句柄.
返回:
	返回列数量.
*/
func XAdapterTreeGetCountColumn(hAdapter HXCGUI) int {
	ret, _, _ := xAdapterTree_GetCountColumn.Call(uintptr(hAdapter))

	return int(ret)
}

/*
设置项数据.

参数:
	hAdapter 数据适配器句柄.
	nID 项ID.
	iColumn 列索引.
	pValue 值.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTreeSetItemText(hAdapter HXCGUI, nID, iColumn int, pValue *uint16) bool {
	ret, _, _ := xAdapterTree_SetItemText.Call(
		uintptr(hAdapter),
		uintptr(nID),
		uintptr(iColumn),
		uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
设置项文件内容.

参数:
	hAdapter 数据适配器句柄.
	nID 项ID.
	pName 列名.
	pValue 值.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTreeSetItemTextEx(hAdapter HXCGUI, nID int, pName, pValue *uint16) bool {
	ret, _, _ := xAdapterTree_SetItemTextEx.Call(
		uintptr(hAdapter),
		uintptr(nID),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
设置项数据.

参数:
	hAdapter 数据适配器句柄.
	nID 项ID.
	iColumn 列索引.
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTreeSetItemImage(hAdapter HXCGUI, nID, iColumn int, hImage HIMAGE) bool {
	ret, _, _ := xAdapterTree_SetItemImage.Call(
		uintptr(hAdapter),
		uintptr(nID),
		uintptr(iColumn),
		uintptr(hImage))

	return ret == TRUE
}

/*
设置项内容.

参数:
	hAdapter 数据适配器句柄.
	nID 项ID.
	pName 列名.
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTreeSetItemImageEx(hAdapter HXCGUI, nID int, pName *uint16, hImage HIMAGE) bool {
	ret, _, _ := xAdapterTree_SetItemImageEx.Call(
		uintptr(hAdapter),
		uintptr(nID),
		uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return ret == TRUE
}

/*
获取项文本内容.

参数:
	hAdapter 数据适配器句柄.
	nID 项ID.
	pName 列名.
	pOut 接收返回内容缓冲区.
	nOutLen 缓冲区大小,字符为单位.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTreeGetItemTextEx(hAdapter HXCGUI, nID int, pName, pOut *uint16, nOutLen int) bool {
	ret, _, _ := xAdapterTree_GetItemTextEx.Call(
		uintptr(hAdapter),
		uintptr(nID),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))

	return ret == TRUE
}

/*
获取项内容.

参数:
	hAdapter 数据适配器句柄.
	nID 项ID.
	pName 列名.
返回:
	返回图片句柄.
*/
func XAdapterTreeGetItemImageEx(hAdapter HXCGUI, nID int, pName *uint16) HIMAGE {
	ret, _, _ := xAdapterTree_GetItemImageEx.Call(
		uintptr(hAdapter),
		uintptr(nID),
		uintptr(unsafe.Pointer(pName)))

	return HIMAGE(ret)
}

/*
删除项.

参数:
	hAdapter 数据适配器句柄.
	nID 项ID.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterTreeDeleteItem(hAdapter HXCGUI, nID int) bool {
	ret, _, _ := xAdapterTree_DeleteItem.Call(
		uintptr(hAdapter),
		uintptr(nID))

	return ret == TRUE
}

/*
删除所有项.

参数:
	hAdapter 数据适配器句柄.
*/
func XAdapterTreeDeleteItemAll(hAdapter HXCGUI) {
	xAdapterTree_DeleteItemAll.Call(uintptr(hAdapter))
}

/*
删除所有列,并且清空数据.

参数:
	hAdapter 数据适配器句柄.
*/
func XAdapterTreeDeleteColumnAll(hAdapter HXCGUI) {
	xAdapterTree_DeleteColumnAll.Call(uintptr(hAdapter))
}
