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

func XToolBarCreate(x int, y int, cx int, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xTextLink_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))
	return HELE(ret)
}

func XToolBarInsertEle(hEle HELE, hNewEle HELE, index int) int {
	ret, _, _ := xToolBar_InsertEle.Call(
		uintptr(hEle),
		uintptr(hNewEle),
		uintptr(index))
	return int(ret)
}
func XToolBarInsertSeparator(hEle HELE, index int) int {
	ret, _, _ := xToolBar_InsertSeparator.Call(
		uintptr(hEle),
		uintptr(index))
	return int(ret)
}

func XToolBarEnableButtonMenu(hEle HELE, bEnable bool) {
	xToolBar_EnableButtonMenu.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))

}

func XToolBarGetHEle(hEle HELE, index int) HELE {
	ret, _, _ := xToolBar_GetHEle.Call(
		uintptr(hEle),
		uintptr(index))
	return HELE(ret)
}

func XToolBarGetButtonLeft(hEle HELE) HELE {
	ret, _, _ := xToolBar_GetButtonLeft.Call(
		uintptr(hEle))
	return HELE(ret)
}

func XToolBarGetButtonRight(hEle HELE) HELE {
	ret, _, _ := xToolBar_GetButtonRight.Call(
		uintptr(hEle))
	return HELE(ret)
}

func XToolBarGetButtonMenu(hEle HELE) HELE {
	ret, _, _ := xToolBar_GetButtonMenu.Call(
		uintptr(hEle))
	return HELE(ret)
}

func XToolBarSetSpace(hEle HELE, nSize int) {
	xToolBar_SetSpace.Call(uintptr(hEle), uintptr(nSize))
}
func XToolBarDeleteEle(hEle HELE, index int) {
	xToolBar_DeleteEle.Call(uintptr(hEle), uintptr(index))
}
func XToolBarDeleteAllEle(hEle HELE) {
	xToolBar_DeleteAllEle.Call(uintptr(hEle))
}
