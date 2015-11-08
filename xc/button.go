package xc

import (
    "syscall"
)

var (
    // Functions
    XBtn_Create  *syscall.Proc
    XBtn_SetType *syscall.Proc
)

func init() {
    XBtn_Create = XCDLL.MustFindProc("XBtn_Create")
    XBtn_SetType = XCDLL.MustFindProc("XBtn_SetType")

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

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-6 18:36:20
// @Function: xBtnSetType
// @Description: 设置按钮样式.
// @Calls: XBtn_SetType
// @Input: hEle 元素句柄. nStyle 样式. 参考 button_type_
// *******************************************
func xBtnSetType(hEle HELE, nType uint32) {
    XBtn_SetType.Call(
        uintptr(hEle),
        uintptr(nType))
}

// 关闭
func CloseBtn(hWindow HWINDOW) {
    xBtnSetType(XBtnCreate(10, 5, 35, 20, "关闭", HXCGUI(hWindow)),
        uint32(BUTTON_TYPE_CLOSE))
}

// 最小化
func MinBtn(hWindow HWINDOW) {
    xBtnSetType(XBtnCreate(60, 5, 45, 20, "最小化", HXCGUI(hWindow)),
        uint32(BUTTON_TYPE_MIN))
}
