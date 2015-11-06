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

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-6 18:33:01
// @Function: XBtnCreate
// @Description: 创建按钮元素.
// @Calls: XBtn_Create
// @Input: x [按钮x坐标]. y [按钮y坐标]. cx [宽度]. cy [高度]. pName [标题].
//         hParent [父是窗口句柄或元素句柄].
// @Return: 按钮元素句柄.
// *******************************************
func XBtnCreate(x, y, cx, cy int, pName string, hParent HXCGUI) HELE {
    ret, _, _ := XBtn_Create.Call(
        uintptr(x),
        uintptr(y),
        uintptr(cx),
        uintptr(cy),
        StringToUintPtr(pName),
        uintptr(hParent))

    return HELE(ret)
}

// @Author: cody.guo
// @Date: 2015-11-6 18:45:49
// @Function: XEleRegEventC
// @Description: 注册元素事件,将类成员函数作为事件回调函数.回调函数省略元素自身句柄hEle,省略触发事件元素句柄hEventEle.
// @Calls: XEle_RegEventC
// @Input: hEle [元素句柄]. nEvent [事件类型]. memberFunction [类成员函数].
func XEleRegEventC(hEle HELE, nEvent int, memberFunction CallBack) {
    XEle_RegEventC.Call(
        uintptr(hEle),
        uintptr(nEvent),
        syscall.NewCallback(memberFunction))
}
