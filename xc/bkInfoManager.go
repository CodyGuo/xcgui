package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xBkInfoM_Create    *syscall.Proc
	xBkInfoM_Destroy   *syscall.Proc
	xBkInfoM_SetBkInfo *syscall.Proc
	xBkInfoM_AddBorder *syscall.Proc
	xBkInfoM_AddFill   *syscall.Proc
	xBkInfoM_AddImage  *syscall.Proc
	xBkInfoM_GetCount  *syscall.Proc
	xBkInfoM_Clear     *syscall.Proc
	xBkInfoM_Draw      *syscall.Proc
	xBkInfoM_DrawEx    *syscall.Proc
)

func init() {
	// Functions
	xBkInfoM_Create = xcDLL.MustFindProc("XBkInfoM_Create")
	xBkInfoM_Destroy = xcDLL.MustFindProc("XBkInfoM_Destroy")
	xBkInfoM_SetBkInfo = xcDLL.MustFindProc("XBkInfoM_SetBkInfo")
	xBkInfoM_AddBorder = xcDLL.MustFindProc("XBkInfoM_AddBorder")
	xBkInfoM_AddFill = xcDLL.MustFindProc("XBkInfoM_AddFill")
	xBkInfoM_AddImage = xcDLL.MustFindProc("XBkInfoM_AddImage")
	xBkInfoM_GetCount = xcDLL.MustFindProc("XBkInfoM_GetCount")
	xBkInfoM_Clear = xcDLL.MustFindProc("XBkInfoM_Clear")
	xBkInfoM_Draw = xcDLL.MustFindProc("XBkInfoM_Draw")
}

/*
创建背景内容管理器.

返回:
	背景内容管理器句柄.
*/
func XBkInfoM_Create() HBKINFOM {
	ret, _, _ := xBkInfoM_Create.Call()

	return HBKINFOM(ret)
}

/*
销毁.

参数:
	hBkInfoM 背景内容管理器句柄.
*/
func XBkInfoM_Destroy(hBkInfoM HBKINFOM) {
	xBkInfoM_Destroy.Call(uintptr(hBkInfoM))
}

/*
设置背景内容.

参数:
	hBkInfoM 背景内容管理器句柄.
	pText 背景内容字符串.*uint16
返回:
	返回添加的背景内容数量.
*/
func XBkInfoM_SetBkInfo(hBkInfoM HBKINFOM, pText string) int {
	ret, _, _ := xBkInfoM_SetBkInfo.Call(
		uintptr(hBkInfoM),
		StringToUintPtr(pText))
	// uintptr(unsafe.Pointer(pText)))

	return int(ret)
}

/*
添加背景内容边框.

参数:
	hBkInfoM 背景内容管理器句柄.
	color RGB颜色.
	alpha 透明度.
	width 线宽.
*/
func XBkInfoM_AddBorder(hBkInfoM HBKINFOM, color COLORREF, alpha byte, width int) {
	xBkInfoM_AddBorder.Call(
		uintptr(hBkInfoM),
		uintptr(color),
		uintptr(alpha),
		uintptr(width))
}

/*
添加背景内容填充.

参数:
	hBkInfoM 背景内容管理器句柄.
	color RGB颜色.
	alpha 透明度.
*/
func XBkInfoM_AddFill(hBkInfoM HBKINFOM, color COLORREF, alpha byte) {
	xBkInfoM_AddFill.Call(
		uintptr(hBkInfoM),
		uintptr(color),
		uintptr(alpha))
}

/*
添加背景内容图片.

参数:
	hBkInfoM 背景内容管理器句柄.
	hImage 图片句柄.
*/
func XBkInfoM_AddImage(hBkInfoM HBKINFOM, hImage HIMAGE) {
	xBkInfoM_AddImage.Call(
		uintptr(hBkInfoM),
		uintptr(hImage))
}

/*
获取背景内容数量.

参数:
	hBkInfoM 背景内容管理器句柄.
返回:
	背景内容数量.
*/
func XBkInfoM_GetCount(hBkInfoM HBKINFOM) int {
	ret, _, _ := xBkInfoM_GetCount.Call(uintptr(hBkInfoM))

	return int(ret)
}

/*
清空背景内容.

参数:
	hBkInfoM 背景内容管理器句柄.
*/
func XBkInfoM_Clear(hBkInfoM HBKINFOM) {
	xBkInfoM_Clear.Call(uintptr(hBkInfoM))
}

/*
绘制背景内容.

参数:
	hBkInfoM 背景内容管理器句柄.
	hDraw 图形绘制句柄.
	pRect 区域坐标.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XBkInfoM_Draw(hBkInfoM HBKINFOM, hDraw HDRAW, pRect *RECT) bool {
	ret, _, _ := xBkInfoM_Draw.Call(
		uintptr(hBkInfoM),
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pRect)))

	return ret == TRUE
}

/*
绘制背景内容, 设置条件.

参数:
	hBkInfoM 背景内容管理器句柄.
	nState 组合状态.
	hDraw 图形绘制句柄.
	pRect 区域坐标.
	nStateFilter 当状态组合大于或等于该值时有效.
返回:
	成功返回TRUE否则返回FALSE. 注解:组合状态大于或等于nStateFilter时才有效 ,
	例如用来绘制列表项,过滤掉元素的背景绘制,避免列表项与元素背景叠加.
*/
func XBkInfoM_DrawEx(hBkInfoM HBKINFOM, nState int, hDraw HDRAW, pRect *RECT, nStateFilter int) bool {
	ret, _, _ := xBkInfoM_DrawEx.Call(
		uintptr(hBkInfoM),
		uintptr(nState),
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(nStateFilter))

	return ret == TRUE
}
