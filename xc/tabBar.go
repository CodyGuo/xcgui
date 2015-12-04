package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xTabBar_Create            *syscall.Proc
	xTabBar_AddLabel          *syscall.Proc
	xTabBar_InsertLabel       *syscall.Proc
	xTabBar_DeleteLabel       *syscall.Proc
	xTabBar_DeleteLabelAll    *syscall.Proc
	xTabBar_GetLabel          *syscall.Proc
	xTabBar_GetLabelClose     *syscall.Proc
	xTabBar_GetButtonLeft     *syscall.Proc
	xTabBar_GetButtonRight    *syscall.Proc
	xTabBar_GetSelect         *syscall.Proc
	xTabBar_GetLabelSpacing   *syscall.Proc
	xTabBar_GetLabelCount     *syscall.Proc
	xTabBar_SetLabelSpacing   *syscall.Proc
	xTabBar_SetSelect         *syscall.Proc
	xTabBar_SetUp             *syscall.Proc
	xTabBar_SetDown           *syscall.Proc
	xTabBar_EnableTile        *syscall.Proc
	xTabBar_EnableClose       *syscall.Proc
	xTabBar_SetCloseSize      *syscall.Proc
	xTabBar_SetTurnButtonSize *syscall.Proc
	xTabBar_ShowLabel         *syscall.Proc
)

func init() {
	// Functions
	xTabBar_Create = xcDLL.MustFindProc("XTabBar_Create")
	xTabBar_AddLabel = xcDLL.MustFindProc("XTabBar_AddLabel")
	xTabBar_InsertLabel = xcDLL.MustFindProc("XTabBar_InsertLabel")
	xTabBar_DeleteLabel = xcDLL.MustFindProc("XTabBar_DeleteLabel")
	xTabBar_DeleteLabelAll = xcDLL.MustFindProc("XTabBar_DeleteLabelAll")
	xTabBar_GetLabel = xcDLL.MustFindProc("XTabBar_GetLabel")
	xTabBar_GetLabelClose = xcDLL.MustFindProc("XTabBar_GetLabelClose")
	xTabBar_GetButtonLeft = xcDLL.MustFindProc("XTabBar_GetButtonLeft")
	xTabBar_GetButtonRight = xcDLL.MustFindProc("XTabBar_GetButtonRight")
	xTabBar_GetSelect = xcDLL.MustFindProc("XTabBar_GetSelect")
	xTabBar_GetLabelSpacing = xcDLL.MustFindProc("XTabBar_GetLabelSpacing")
	xTabBar_GetLabelCount = xcDLL.MustFindProc("XTabBar_GetLabelCount")
	xTabBar_SetLabelSpacing = xcDLL.MustFindProc("XTabBar_SetLabelSpacing")
	xTabBar_SetSelect = xcDLL.MustFindProc("XTabBar_SetSelect")
	xTabBar_SetUp = xcDLL.MustFindProc("XTabBar_SetUp")
	xTabBar_SetDown = xcDLL.MustFindProc("XTabBar_SetDown")
	xTabBar_EnableTile = xcDLL.MustFindProc("XTabBar_EnableTile")
	xTabBar_EnableClose = xcDLL.MustFindProc("XTabBar_EnableClose")
	xTabBar_SetCloseSize = xcDLL.MustFindProc("XTabBar_SetCloseSize")
	xTabBar_SetTurnButtonSize = xcDLL.MustFindProc("XTabBar_SetTurnButtonSize")
	xTabBar_ShowLabel = xcDLL.MustFindProc("XTabBar_ShowLabel")
}

/*
创建tabBar元素.

参数:
	x 元素x坐标.
	y 元素y坐标.
	cx 宽度.
	cy 高度.
	hParent 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
返回:
	元素句柄.
*/
func XTabBarCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xTabBar_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
添加一个标签.

参数:
	hEle 元素句柄
	*pName 标签文本内容.*uint16
返回:
	标签索引.
*/
func XTabBarAddLabel(hEle HELE, pName string) int {
	ret, _, _ := xTabBar_AddLabel.Call(
		uintptr(hEle),
		StringToUintPtr(pName))
	// uintptr(unsafe.Pointer(pName)))

	return int(ret)
}

/*
插入一个标签.

参数:
	hEle 元素句柄.
	index 插入位置.
	*pName 标签文本内容.*uint16
返回:
	标签索引.
*/
func XTabBarInsertLabel(hEle HELE, index int, pName string) int {
	ret, _, _ := xTabBar_InsertLabel.Call(
		uintptr(hEle),
		uintptr(index),
		StringToUintPtr(pName))
	// uintptr(unsafe.Pointer(pName)))

	return int(ret)
}

/*
删除一个标签.

参数:
	hEle 元素句柄.
	index 位置索引.
返回:
	成功返回TRUE否则FALSE.
*/
func XTabBarDeleteLabel(hEle HELE, index int) bool {
	ret, _, _ := xTabBar_DeleteLabel.Call(
		uintptr(hEle),
		uintptr(index))

	return ret == TRUE
}

/*
删除所有标签.

参数:
	hEle 元素句柄.
*/
func XTabBarDeleteLabelAll(hEle HELE) {
	xTabBar_DeleteLabelAll.Call(uintptr(hEle))
}

/*
获取标签按钮Button.

参数:
	hEle 元素句柄.
	index 位置索引.
返回:
	按钮句柄.
*/
func XTabBarGetLabel(hEle HELE, index int) HELE {
	ret, _, _ := xTabBar_GetLabel.Call(
		uintptr(hEle),
		uintptr(index))

	return HELE(ret)
}

/*
获取标签上关闭按钮.

参数:
	hEle 元素句柄.
	index 位置索引.
返回:
	按钮句柄.
*/
func XTabBarGetLabelClose(hEle HELE, index int) HELE {
	ret, _, _ := xTabBar_GetLabelClose.Call(
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
func XTabBarGetButtonLeft(hEle HELE) HELE {
	ret, _, _ := xTabBar_GetButtonLeft.Call(uintptr(hEle))

	return HELE(ret)
}

/*
获取右滚动按钮.

参数:
	hEle 元素句柄.
返回:
	返回按钮句柄.
*/
func XTabBarGetButtonRight(hEle HELE) HELE {
	ret, _, _ := xTabBar_GetButtonRight.Call(uintptr(hEle))

	return HELE(ret)
}

/*
获取选择的标签索引.

参数:
	hEle 元素句柄.
返回:
	标签位置索引.
*/
func XTabBarGetSelect(hEle HELE) int {
	ret, _, _ := xTabBar_GetSelect.Call(uintptr(hEle))

	return int(ret)
}

/*
获取标签间距, 0没有间距.

参数:
	hEle 元素句柄.
返回:
	标签间隔大小
*/
func XTabBarGetLabelSpacing(hEle HELE) int {
	ret, _, _ := xTabBar_GetLabelSpacing.Call(uintptr(hEle))

	return int(ret)
}

/*
获取标签项数量.

参数:
	hEle 元素句柄.
返回:
	标签项数量.
*/
func XTabBarGetLabelCount(hEle HELE) int {
	ret, _, _ := xTabBar_GetLabelCount.Call(uintptr(hEle))

	return int(ret)
}

/*
设置标签间距, 0没有间距.

参数:
	hEle 元素句柄.
	spacing 标签间隔大小.
*/
func XTabBarSetLabelSpacing(hEle HELE, spacing int) {
	xTabBar_SetLabelSpacing.Call(
		uintptr(hEle),
		uintptr(spacing))
}

/*
设置选择标签.

参数:
	hEle 元素句柄.
	index 标签位置索引.
*/
func XTabBarSetSelect(hEle HELE, index int) {
	xTabBar_SetSelect.Call(
		uintptr(hEle),
		uintptr(index))
}

/*
左按钮滚动.

参数:
	hEle 元素句柄.
*/
func XTabBarSetUp(hEle HELE) {
	xTabBar_SetUp.Call(uintptr(hEle))
}

/*
右按钮滚动.

参数:
	hEle 元素句柄.
*/
func XTabBarSetDown(hEle HELE) {
	xTabBar_SetDown.Call(uintptr(hEle))
}

/*
平铺标签,每个标签显示相同大小.

参数:
	hEle 元素句柄.
	bTile 是否启用.
*/
func XTabBarEnableTile(hEle HELE, bTile bool) {
	xTabBar_EnableTile.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bTile)))
}

/*
启用关闭标签功能.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XTabBarEnableClose(hEle HELE, bEnable bool) {
	xTabBar_EnableClose.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
设置关闭按钮大小.

参数:
	hEle 元素句柄.
	pSize 大小值, 宽度和高度可以为-1,-1代表默认值.
*/
func XTabBarSetCloseSize(hEle HELE, pSize *SIZE) {
	xTabBar_SetCloseSize.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pSize)))
}

/*
设置翻滚按钮大小.

参数:
	hEle 元素句柄.
	pSize 大小值, 宽度和高度可以为-1,-1代表默认值.
*/
func XTabBarSetTurnButtonSize(hEle HELE, pSize *SIZE) {
	xTabBar_SetTurnButtonSize.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pSize)))
}

/*
显示或隐藏指定标签.

参数:
	hEle 元素句柄.
	index 标签索引.
	bShow 是否显示.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XTabBarShowLabel(hEle HELE, index int, bShow bool) bool {
	ret, _, _ := xTabBar_ShowLabel.Call(
		uintptr(hEle),
		uintptr(index),
		uintptr(BoolToBOOL(bShow)))

	return ret == TRUE
}
