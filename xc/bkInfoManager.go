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
)

func init() {
	// Functions
	xBkInfoM_Create = XCDLL.MustFindProc("XBkInfoM_Create")
	xBkInfoM_Destroy = XCDLL.MustFindProc("XBkInfoM_Destroy")
	xBkInfoM_SetBkInfo = XCDLL.MustFindProc("XBkInfoM_SetBkInfo")
	xBkInfoM_AddBorder = XCDLL.MustFindProc("XBkInfoM_AddBorder")
	xBkInfoM_AddFill = XCDLL.MustFindProc("XBkInfoM_AddFill")
	xBkInfoM_AddImage = XCDLL.MustFindProc("XBkInfoM_AddImage")
	xBkInfoM_GetCount = XCDLL.MustFindProc("XBkInfoM_GetCount")
	xBkInfoM_Clear = XCDLL.MustFindProc("XBkInfoM_Clear")
	xBkInfoM_Draw = XCDLL.MustFindProc("XBkInfoM_Draw")
}

/*
创建背景内容管理器.

返回:
	背景内容管理器句柄.
*/
func XBkInfoMCreate() HBKINFOM {
	ret, _, _ := xBkInfoM_Create.Call()

	return HBKINFOM(ret)
}

/*
销毁.

参数:
	hBkInfoM 背景内容管理器句柄.
*/
func XBkInfoMDestroy(hBkInfoM HBKINFOM) {
	xBkInfoM_Destroy.Call(uintptr(hBkInfoM))
}

/*
设置背景内容.

参数:
	hBkInfoM 背景内容管理器句柄.
	pText 背景内容字符串.
返回:
	返回添加的背景内容数量.
*/
func XBkInfoMSetBkInfo(hBkInfoM HBKINFOM, pText string) int {
	ret, _, _ := xBkInfoM_SetBkInfo.Call(
		uintptr(hBkInfoM),
		StringToUintPtr(pText))

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
func XBkInfoMAddBorder(hBkInfoM HBKINFOM, color COLORREF, alpha byte, width int) {
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
func XBkInfoMAddFill(hBkInfoM HBKINFOM, color COLORREF, alpha byte) {
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
func XBkInfoMAddImage(hBkInfoM HBKINFOM, hImage HIMAGE) {
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
func XBkInfoMGetCount(hBkInfoM HBKINFOM) int {
	ret, _, _ := xBkInfoM_GetCount.Call(uintptr(hBkInfoM))

	return int(ret)
}

/*
清空背景内容.

参数:
	hBkInfoM 背景内容管理器句柄.
*/
func XBkInfoMClear(hBkInfoM HBKINFOM) {
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
func XBkInfoMDraw(hBkInfoM HBKINFOM, hDraw HDRAW, pRect *RECT) bool {
	ret, _, _ := xBkInfoM_Draw.Call(
		uintptr(hBkInfoM),
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pRect)))

	return ret == TRUE
}
