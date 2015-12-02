package xc

import (
	"syscall"
	"unsafe"
)

var (
	xMenuBar_Create          *syscall.Proc
	xMenuBar_AddButton       *syscall.Proc
	xMenuBar_SetButtonHeight *syscall.Proc
	xMenuBar_GetMenu         *syscall.Proc
	xMenuBar_DeleteButton    *syscall.Proc
)

func init() {
	xMenuBar_Create = xcDLL.MustFindProc("XMenuBar_Create")
	xMenuBar_AddButton = xcDLL.MustFindProc("XMenuBar_AddButton")
	xMenuBar_SetButtonHeight = xcDLL.MustFindProc("XMenuBar_SetButtonHeight")
	xMenuBar_GetMenu = xcDLL.MustFindProc("XMenuBar_GetMenu")
	xMenuBar_DeleteButton = xcDLL.MustFindProc("XMenuBar_DeleteButton")
}

/*
创建菜单条元素;如果指定了父为窗口,默认调用XWnd_AddMenuBar()函数,将菜单条添加到窗口非客户区.

参数:
	x 元素x坐标.
	y 元素y坐标.
	cx 宽度.
	cy 高度.
	hParent 父是窗口资源句柄或U I元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
返回:
	元素句柄.
*/
func XMenuBarCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xMenuBar_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
添加弹出菜单按钮.

参数:
	hEle 元素句柄.
	pText 文本内容.
返回:
	返回菜单按钮索引.
*/
func XMenuBarAddButton(hEle HELE, pText *uint16) int {
	ret, _, _ := xMenuBar_AddButton.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pText)))

	return int(ret)
}

/*
设置菜单按钮高度.

参数:
	hEle 元素句柄.
	height 高度.
*/
func XMenuBarSetButtonHeight(hEle HELE, height int) {
	xMenuBar_SetButtonHeight.Call(
		uintptr(hEle),
		uintptr(height))
}

/*
获取菜单.

参数:
	hEle 元素句柄.
	nIndex 菜单条上菜单按钮的索引.
返回:
	返回菜单句柄.
*/
func XMenuBarGetMenu(hEle HELE, nIndex int) HMENUX {
	ret, _, _ := xMenuBar_GetMenu.Call(
		uintptr(hEle),
		uintptr(nIndex))
	return HMENUX(ret)
}

/*
删除菜单条上的菜单按钮,同时该按钮下的弹出菜单也被销毁.

参数:
	hEle 元素句柄.
	nIndex 菜单条按钮索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XMenuBarDeleteButton(hEle HELE, nIndex int) bool {
	ret, _, _ := xMenuBar_GetMenu.Call(
		uintptr(hEle),
		uintptr(nIndex))

	return ret == TRUE
}
