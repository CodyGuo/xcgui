package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xAdapterMap_Create       *syscall.Proc
	xAdapterMap_AddItemText  *syscall.Proc
	xAdapterMap_AddItemImage *syscall.Proc
	xAdapterMap_DeleteItem   *syscall.Proc
	xAdapterMap_GetCount     *syscall.Proc
	xAdapterMap_GetItemText  *syscall.Proc
	xAdapterMap_GetItemImage *syscall.Proc
)

func init() {
	// Funtions
	xAdapterMap_Create = xcDLL.MustFindProc("XAdapterMap_Create")
	xAdapterMap_AddItemText = xcDLL.MustFindProc("XAdapterMap_AddItemText")
	xAdapterMap_AddItemImage = xcDLL.MustFindProc("XAdapterMap_AddItemImage")
	xAdapterMap_DeleteItem = xcDLL.MustFindProc("XAdapterMap_DeleteItem")
	xAdapterMap_GetCount = xcDLL.MustFindProc("XAdapterMap_GetCount")
	xAdapterMap_GetItemText = xcDLL.MustFindProc("XAdapterMap_GetItemText")
	xAdapterMap_GetItemImage = xcDLL.MustFindProc("XAdapterMap_GetItemImage")
}

/*
创建数据适配器,单列数据.

返回:
	返回数据适配器句柄.
*/
func XAdapterMapCreate() HXCGUI {
	ret, _, _ := xAdapterMap_Create.Call()

	return HXCGUI(ret)
}

/*
增加数据项.

参数:
	hAdapter 数据适配器句柄.
	pName 项名称.
	pValue 值.*uint16
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterMapAddItemText(hAdapter HXCGUI, pName, pValue string) bool {
	ret, _, _ := xAdapterMap_AddItemText.Call(
		uintptr(hAdapter),
		StringToUintPtr(pName),
		StringToUintPtr(pValue))
	// uintptr(unsafe.Pointer(pName)),
	// uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
增加数据项.

参数:
	hAdapter 数据适配器句柄.
	pName 项名称.*uint16
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterMapAddItemImage(hAdapter HXCGUI, pName string, hImage HIMAGE) bool {
	ret, _, _ := xAdapterMap_AddItemImage.Call(
		uintptr(hAdapter),
		StringToUintPtr(pName),
		// uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return ret == TRUE
}

/*
删除数据项.

参数:
	hAdapter 数据适配器句柄.
	pName 项名称.*uint16
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterMapDeleteItem(hAdapter HXCGUI, pName string) bool {
	ret, _, _ := xAdapterMap_DeleteItem.Call(
		uintptr(hAdapter),
		StringToUintPtr(pName))
	// uintptr(unsafe.Pointer(pName)))

	return ret == TRUE
}

/*
获取项数量.

参数:
	hAdapter 数据适配器句柄.
返回:
	返回项数量.
*/
func XAdapterMapGetCount(hAdapter HXCGUI) int {
	ret, _, _ := xAdapterMap_GetCount.Call(uintptr(hAdapter))

	return int(ret)
}

/*
获取项内容,如果内容为文本.

参数:
	hAdapter 数据适配器句柄.
	pName 项名称.
	pOut 接收返回文本缓冲区.
	nOutLen 缓冲区长度,字符为单位.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdapterMapGetItemText(hAdapter HXCGUI, pName, pOut *uint16, nOutLen int) bool {
	ret, _, _ := xAdapterMap_GetItemText.Call(
		uintptr(hAdapter),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))

	return ret == TRUE
}

/*
获取项内容,如果内容为图片句柄.

参数:
	hAdapter 数据适配器句柄.
	pName 项名称.*uint16
返回:
	返回图片句柄.
*/
func XAdapterMapGetItemImage(hAdapter HXCGUI, pName string) HIMAGE {
	ret, _, _ := xAdapterMap_GetItemImage.Call(
		uintptr(hAdapter),
		StringToUintPtr(pName))
	// uintptr(unsafe.Pointer(pName)))

	return HIMAGE(ret)
}
