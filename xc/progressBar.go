//ProgressBar-进度条
package xc

import (
	"syscall"
)

var (
	xProgBar_Create       *syscall.Proc
	xProgBar_SetRange     *syscall.Proc
	xProgBar_GetRange     *syscall.Proc
	xProgBar_SetImageLoad *syscall.Proc
	xProgBar_SetSpaceTwo  *syscall.Proc
	xProgBar_SetPos       *syscall.Proc
	xProgBar_GetPos       *syscall.Proc
	xProgBar_SetHorizon   *syscall.Proc
)

func init() {
	xProgBar_Create = XCDLL.MustFindProc("XProgBar_Create")
	xProgBar_SetRange = XCDLL.MustFindProc("XProgBar_SetRange")
	xProgBar_GetRange = XCDLL.MustFindProc("XProgBar_GetRange")
	xProgBar_SetImageLoad = XCDLL.MustFindProc("XProgBar_SetImageLoad")
	xProgBar_SetSpaceTwo = XCDLL.MustFindProc("XProgBar_SetSpaceTwo")
	xProgBar_SetPos = XCDLL.MustFindProc("XProgBar_SetPos")
	xProgBar_GetPos = XCDLL.MustFindProc("XProgBar_GetPos")
	xProgBar_SetHorizon = XCDLL.MustFindProc("XProgBar_SetHorizon")
}

func XProgBarCreate(x int, y int, cx int, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xProgBar_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))
	return HELE(ret)
}

func XProgBarSetRange(hEle HELE, irange int) {
	xProgBar_SetRange.Call(uintptr(hEle), uintptr(irange))
}

func XProgBarGetRange(hEle HELE) int {
	ret, _, _ := xProgBar_GetRange.Call(uintptr(hEle))
	return int(ret)
}

func XProgBarSetImageLoad(hEle HELE, hImage HIMAGE) {
	xProgBar_SetImageLoad.Call(uintptr(hEle), uintptr(hImage))
}

func XProgBarSetSpaceTwo(hEle HELE, leftSize int, rightSize int) {
	xProgBar_SetSpaceTwo.Call(uintptr(hEle), uintptr(leftSize), uintptr(rightSize))
}

func XProgBarSetPos(hEle HELE, pos int) {
	xProgBar_SetPos.Call(uintptr(hEle), uintptr(pos))
}

func XProgBarGetPos(hEle HELE) int {
	ret, _, _ := xProgBar_GetPos.Call(uintptr(hEle))
	return int(ret)
}

func XProgBarSetHorizon(hEle HELE, bHorizon bool) {
	xProgBar_SetPos.Call(uintptr(hEle), uintptr(BoolToBOOL(bHorizon)))
}
