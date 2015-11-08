package xc

import (
    "syscall"
)

var (
    // Functions
    XBtn_Create *syscall.Proc
)

func init() {
    XBtn_Create = XCDLL.MustFindProc("XBtn_Create")
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-6 18:33:01
// @Function: XBtnCreate
// @Description: 创建按钮元素.
// @Calls: XBtn_Create
// @Input: x 按钮x坐标. y 按钮y坐标. cx 宽度. cy 高度. pName 标题.
//         hParent 父是窗口句柄或元素句柄.
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
