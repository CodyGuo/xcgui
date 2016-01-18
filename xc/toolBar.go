package xc

import (
	"syscall"
)

var (
	xToolBar_Create           *syscall.Proc
	xToolBar_InsertEle        *syscall.Proc
	xToolBar_InsertSeparator  *syscall.Proc
	xToolBar_EnableButtonMenu *syscall.Proc
	xToolBar_GetHEle          *syscall.Proc
	xToolBar_GetButtonLeft    *syscall.Proc
	xToolBar_GetButtonRight   *syscall.Proc
	xToolBar_GetButtonMenu    *syscall.Proc
	xToolBar_SetSpace         *syscall.Proc
	xToolBar_DeleteEle        *syscall.Proc
	xToolBar_DeleteAllEle     *syscall.Proc
)

func init() {
	xToolBar_Create = xcDLL.MustFindProc("XToolBar_Create")
	xToolBar_InsertEle = xcDLL.MustFindProc("XToolBar_InsertEle")
	xToolBar_InsertSeparator = xcDLL.MustFindProc("XToolBar_InsertSeparator")
	xToolBar_EnableButtonMenu = xcDLL.MustFindProc("XToolBar_EnableButtonMenu")
	xToolBar_GetHEle = xcDLL.MustFindProc("XToolBar_GetHEle")
	xToolBar_GetButtonLeft = xcDLL.MustFindProc("XToolBar_GetButtonLeft")
	xToolBar_GetButtonRight = xcDLL.MustFindProc("XToolBar_GetButtonRight")
	xToolBar_GetButtonMenu = xcDLL.MustFindProc("XToolBar_GetButtonMenu")
	xToolBar_SetSpace = xcDLL.MustFindProc("XToolBar_SetSpace")
	xToolBar_DeleteEle = xcDLL.MustFindProc("XToolBar_DeleteEle")
	xToolBar_DeleteAllEle = xcDLL.MustFindProc("XToolBar_DeleteAllEle")
}

/*
创建工具条元素;如果指定了父为窗口,默认调用XWnd_AddToolBar()函数,将工具条添加到窗口非客户区.

参数:
	x 元素x坐标.
	y 元素y坐标.
	cx 宽度.
	cy 高度.
	hParent 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
返回:
	元素句柄.
*/
func XToolBar_Create(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xToolBar_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
插入元素到工具条.

参数:
	hEle 元素句柄.
	hNewEle 将要插入的元素.
	index 插入位置索引.
返回:
	返回插入位置索引.
*/
func XToolBar_InsertEle(hEle, hNewEle HELE, index int) int {
	ret, _, _ := xToolBar_InsertEle.Call(
		uintptr(hEle),
		uintptr(hNewEle),
		uintptr(index))

	return int(ret)
}

/*
插入分隔符到工具条.

参数:
	hEle 元素句柄.
	index 插入位置索引.
返回:
	返回插入位置索引.
*/
func XToolBar_InsertSeparator(hEle HELE, index int) int {
	ret, _, _ := xToolBar_InsertSeparator.Call(
		uintptr(hEle),
		uintptr(index))

	return int(ret)
}

/*
启用下拉菜单,显示隐藏的项.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XToolBar_EnableButtonMenu(hEle HELE, bEnable bool) {
	xToolBar_EnableButtonMenu.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))

}

/*
获取工具条上指定元素.

参数:
	hEle 元素句柄.
	index 索引值.
返回:
	返回元素句柄.
*/
func XToolBar_GetHEle(hEle HELE, index int) HELE {
	ret, _, _ := xToolBar_GetHEle.Call(
		uintptr(hEle),
		uintptr(index))

	return HELE(ret)
}

/*
获取左滚动按钮.

参数:
	hEle 元素句柄.
返回:
	返回按钮句柄.
*/
func XToolBar_GetButtonLeft(hEle HELE) HELE {
	ret, _, _ := xToolBar_GetButtonLeft.Call(uintptr(hEle))

	return HELE(ret)
}

/*
获取右滚动按钮.

参数:
	hEle 元素句柄.
返回:
	返回按钮句柄.
*/
func XToolBar_GetButtonRight(hEle HELE) HELE {
	ret, _, _ := xToolBar_GetButtonRight.Call(uintptr(hEle))

	return HELE(ret)
}

/*
获取菜单按钮.

参数:
	hEle 元素句柄.
返回:
	返回菜单按钮句柄.
*/
func XToolBar_GetButtonMenu(hEle HELE) HELE {
	ret, _, _ := xToolBar_GetButtonMenu.Call(uintptr(hEle))

	return HELE(ret)
}

/*
设置对象之间的间距.

参数:
	hEle 元素句柄.
	nSize 间距大小.
*/
func XToolBar_SetSpace(hEle HELE, nSize int) {
	xToolBar_SetSpace.Call(
		uintptr(hEle),
		uintptr(nSize))
}

/*
删除元素,并且销毁.

参数:
	hEle 元素句柄.
	index 索引值.
*/
func XToolBar_DeleteEle(hEle HELE, index int) {
	xToolBar_DeleteEle.Call(
		uintptr(hEle),
		uintptr(index))
}

/*
删除所有元素,并且销毁.

参数:
	hEle 元素句柄.
*/
func XToolBar_DeleteAllEle(hEle HELE) {
	xToolBar_DeleteAllEle.Call(uintptr(hEle))
}
