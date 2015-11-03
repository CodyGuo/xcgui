package xc

import (
    "syscall"
)

const (
    XE_BNCLICK      = 34
    XE_BUTTON_CHECK = 35
)

var (
    // Functions
    XBtn_Create    *syscall.Proc
    XEle_RegEventC *syscall.Proc
)

func init() {
    XBtn_Create = XCDLL.MustFindProc("XBtn_Create")
    XEle_RegEventC = XCDLL.MustFindProc("XEle_RegEventC")
}

type CallBack func() uintptr

// 创建按钮
func XBtnCreate(x, y, cx, cy int, pName string, hParent HWND) (hEle uintptr) {
    hEle, _, _ = XBtn_Create.Call(
        uintptr(x),
        uintptr(y),
        uintptr(cx),
        uintptr(cy),
        StringToUintPtr(pName),
        uintptr(hParent))

    return hEle
}

// 注册事件
func XEleRegEventCPP(hEle uintptr, nEvent int, memberFunction CallBack) {
    XEle_RegEventC.Call(
        hEle,
        uintptr(nEvent),
        syscall.NewCallback(memberFunction))
}
