package xc

import (
	"syscall"
)

var (
	xTextLink_Create                 *syscall.Proc
	xTextLink_EnableUnderlineLeave   *syscall.Proc
	xTextLink_EnableUnderlineStay    *syscall.Proc
	xTextLink_SetTextColorStay       *syscall.Proc
	xTextLink_SetUnderlineColorLeave *syscall.Proc
	xTextLink_SetUnderlineColorStay  *syscall.Proc
)

func init() {
	xTextLink_Create = xcDLL.MustFindProc("XTextLink_Create")
	xTextLink_EnableUnderlineLeave = xcDLL.MustFindProc("XTextLink_EnableUnderlineLeave")
	xTextLink_EnableUnderlineStay = xcDLL.MustFindProc("XTextLink_EnableUnderlineStay")
	xTextLink_SetTextColorStay = xcDLL.MustFindProc("XTextLink_SetTextColorStay")
	xTextLink_SetUnderlineColorLeave = xcDLL.MustFindProc("XTextLink_SetUnderlineColorLeave")
	xTextLink_SetUnderlineColorStay = xcDLL.MustFindProc("XTextLink_SetUnderlineColorStay")
}

func XTextLinkCreate(x int, y int, cx int, cy int, pName string, hParent HXCGUI) HELE {
	ret, _, _ := xTextLink_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		StringToUintPtr(pName),
		uintptr(hParent))
	return HELE(ret)
}

func XTextLinkEnableUnderlineLeave(hEle HELE, bEnable bool) {
	xTextLink_EnableUnderlineLeave.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

func XTextLinkEnableUnderlineStay(hEle HELE, bEnable bool) {
	xTextLink_EnableUnderlineStay.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

func XTextLinkSetTextColorStay(hEle HELE, color COLORREF, alpha byte) {
	xTextLink_SetTextColorStay.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))

}

func XTextLinkSetUnderlineColorLeave(hEle HELE, color COLORREF, alpha byte) {
	xTextLink_SetUnderlineColorLeave.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))
}

func XTextLinkSetUnderlineColorStay(hEle HELE, color COLORREF, alpha byte) {
	xTextLink_SetUnderlineColorStay.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))
}
