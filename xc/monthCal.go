package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xMonthCal_Create     *syscall.Proc
	xMonthCal_GetButton  *syscall.Proc
	xMonthCal_SetToday   *syscall.Proc
	xMonthCal_GetToday   *syscall.Proc
	xMonthCal_SeSelDate  *syscall.Proc
	xMonthCal_GetSelDate *syscall.Proc
)

func init() {
	// Functions
	xMonthCal_Create = xcDLL.MustFindProc("XMonthCal_Create")
	xMonthCal_GetButton = xcDLL.MustFindProc("XMonthCal_GetButton")
	xMonthCal_SetToday = xcDLL.MustFindProc("XMonthCal_SetToday")
	xMonthCal_GetToday = xcDLL.MustFindProc("XMonthCal_GetToday")
	xMonthCal_SeSelDate = xcDLL.MustFindProc("XMonthCal_SeSelDate")
	xMonthCal_GetSelDate = xcDLL.MustFindProc("XMonthCal_GetSelDate")
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
func XMonthCalCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xMonthCal_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
获取内部按钮元素.

参数:
	hEle 元素句柄.
	nType 按钮类型.
返回:
	元素样式.
*/
func XMonthCalGetButton(hEle HELE, nType MonthCal_button_type_) HELE {
	ret, _, _ := xMonthCal_GetButton.Call(
		uintptr(hEle),
		uintptr(nType))

	return HELE(ret)
}

/*
设置月历当前年月日.

参数:
	hEle 元素句柄.
	nYear 年.
	nMonth 月.
	nDay 日.
*/
func XMonthCalSetToday(hEle HELE, nYear, nMonth, nDay int) {
	xMonthCal_SetToday.Call(
		uintptr(hEle),
		uintptr(nYear),
		uintptr(nMonth),
		uintptr(nDay))
}

/*
获取月历当前年月日.

参数:
	hEle 元素句柄.
	pnYear 年.[INT,OUT]
	pnMonth 月.[INT,OUT]
	pnDay 日.[INT,OUT]
*/
func XMonthCalGetToday(hEle HELE, pnYear, pnMonth, pnDay *int) {
	xMonthCal_GetToday.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pnYear)),
		uintptr(unsafe.Pointer(pnMonth)),
		uintptr(unsafe.Pointer(pnDay)))
}

/*
设置月历选中的年月日.

参数:
	hEle 元素句柄.
	nYear 年.
	nMonth 月.
	nDay 日.
*/
func XMonthCalSeSelDate(hEle HELE, nYear, nMonth, nDay int) {
	xMonthCal_SeSelDate.Call(
		uintptr(hEle),
		uintptr(nYear),
		uintptr(nMonth),
		uintptr(nDay))
}

/*
获取月历选中的年月日.

参数:
	hEle 元素句柄.
	pnYear 年.[INT,OUT]
	pnMonth 月.[INT,OUT]
	pnDay 日.[INT,OUT]
*/
func XMonthCalGetSelDate(hEle HELE, pnYear, pnMonth, pnDay *int) {
	xMonthCal_GetSelDate.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pnYear)),
		uintptr(unsafe.Pointer(pnMonth)),
		uintptr(unsafe.Pointer(pnDay)))
}
