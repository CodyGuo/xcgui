package xc

import (
    "syscall"
)

import (
    "github.com/codyguo/sys"
)

const (
    XE_BNCLICK      = 34
    XE_BUTTON_CHECK = 35
)

type CallBack func() uintptr

// 创建按钮
func XBtnCreate(x, y, cx, cy int, pName string, hParent HWND) (hEle uintptr) {
    hEle = XCDLL.Call("XBtn_Create",
        uintptr(x),
        uintptr(y),
        uintptr(cx),
        uintptr(cy),
        sys.TEXT(pName),
        uintptr(hParent))

    return hEle
}

// 注册事件
func XEleRegEventCPP(hEle uintptr, nEvent int, memberFunction CallBack) {
    XCDLL.Call("XEle_RegEventC",
        hEle,
        uintptr(nEvent),
        syscall.NewCallback(memberFunction))
}
