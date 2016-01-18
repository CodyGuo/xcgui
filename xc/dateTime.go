package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xDateTime_Create        *syscall.Proc
	xDateTime_SetStyle      *syscall.Proc
	xDateTime_GetStyle      *syscall.Proc
	xDateTime_GetButton     *syscall.Proc
	xDateTime_GetSelBkColor *syscall.Proc
	xDateTime_SetSelBkColor *syscall.Proc
	xDateTime_GetDate       *syscall.Proc
	xDateTime_SetDate       *syscall.Proc
	xDateTime_GetTime       *syscall.Proc
	xDateTime_SetTime       *syscall.Proc
)

func init() {
	// Functions
	xDateTime_Create = xcDLL.MustFindProc("XDateTime_Create")
	xDateTime_SetStyle = xcDLL.MustFindProc("XDateTime_SetStyle")
	xDateTime_GetStyle = xcDLL.MustFindProc("XDateTime_GetStyle")
	xDateTime_GetButton = xcDLL.MustFindProc("XDateTime_GetButton")
	xDateTime_GetSelBkColor = xcDLL.MustFindProc("XDateTime_GetSelBkColor")
	xDateTime_SetSelBkColor = xcDLL.MustFindProc("XDateTime_SetSelBkColor")
	xDateTime_GetDate = xcDLL.MustFindProc("XDateTime_GetDate")
	xDateTime_SetDate = xcDLL.MustFindProc("XDateTime_SetDate")
	xDateTime_GetTime = xcDLL.MustFindProc("XDateTime_GetTime")
	xDateTime_SetTime = xcDLL.MustFindProc("XDateTime_SetTime")
}

/*
创建日期时间元素

参数:
	x x坐标
	y y坐标
	cx 宽度
	cy 高度
	hParent 父为窗口句柄或元素句柄.
返回:
	元素句柄.
*/
func XDateTime_Create(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xDateTime_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
设置样式.

参数:
	hEle 元素句柄.
	nStyle 样式, 0为日期元素,1为时间元素.
*/
func XDateTime_SetStyle(hEle HELE, nStyle int) {
	xDateTime_SetStyle.Call(
		uintptr(hEle),
		uintptr(nStyle))
}

/*
获取样式.

参数:
	hEle 元素句柄.
返回:
	元素样式.
*/
func XDateTime_GetStyle(hEle HELE) int {
	ret, _, _ := xDateTime_GetStyle.Call(uintptr(hEle))

	return int(ret)
}

/*
获取内部按钮元素.

参数:
	hEle 元素句柄.
	nType 按钮类型.
返回:
	元素样式.
*/
func XDateTime_GetButton(hEle HELE, nType int) HELE {
	ret, _, _ := xDateTime_GetButton.Call(
		uintptr(hEle),
		uintptr(nType))

	return HELE(ret)
}

/*
获取被选择文字的背景颜色.

参数:
	hEle 元素句柄.
返回:
	元素样式.
*/
func XDateTime_GetSelBkColor(hEle HELE) COLORREF {
	ret, _, _ := xDateTime_GetSelBkColor.Call(uintptr(hEle))

	return COLORREF(ret)
}

/*
设置被选择文字的背景颜色.

参数:
	hEle 元素句柄.
	crSelectBk 文字被选中背景色.
	alpha 透明度
*/
func XDateTime_SetSelBkColor(hEle HELE, crSelectBk COLORREF, alpha byte) {
	xDateTime_SetSelBkColor.Call(
		uintptr(hEle),
		uintptr(crSelectBk),
		uintptr(alpha))
}

/*
获取当前日期.

参数:
	hEle 元素句柄.
	pnYear 年.[IN,OUT]
	pnMonth 月.[IN,OUT]
	pnDay 日.[IN,OUT]
*/
func XDateTime_GetDate(hEle HELE, pnYear, pnMonth, pnDay *int) {
	xDateTime_GetDate.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pnYear)),
		uintptr(unsafe.Pointer(pnMonth)),
		uintptr(unsafe.Pointer(pnDay)))
}

/*
设置当前日期.

参数:
	hEle 元素句柄.
	nYear 年.
	nMonth 月.
	nDay 日.
*/
func XDateTime_SetDate(hEle HELE, nYear, nMonth, nDay int) {
	xDateTime_SetDate.Call(
		uintptr(hEle),
		uintptr(nYear),
		uintptr(nMonth),
		uintptr(nDay))
}

/*
获取当前时间.

参数:
	hEle 元素句柄.
	pnHour 时.[IN,OUT]
	pnMinute 分.[IN,OUT]
	pnSecond 秒.[IN,OUT]
*/
func XDateTime_GetTime(hEle HELE, pnHour, pnMinute, pnSecond *int) {
	xDateTime_GetTime.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pnHour)),
		uintptr(unsafe.Pointer(pnMinute)),
		uintptr(unsafe.Pointer(pnSecond)))
}

/*
设置当前时分秒.

参数:
	hEle 元素句柄.
	nHour 时.
	nMinute 分.
	nSecond 秒.
*/
func XDateTime_SetTime(hEle HELE, nHour, nMinute, nSecond int) {
	xDateTime_SetTime.Call(
		uintptr(hEle),
		uintptr(nHour),
		uintptr(nMinute),
		uintptr(nSecond))
}
