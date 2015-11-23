package xc

import (
    "syscall"
    "unsafe"
)

var (
    // Functions
    XEle_RegEventC      *syscall.Proc
    XEle_GetContentSize *syscall.Proc
)

func init() {
    // Functions
    XEle_RegEventC = XCDLL.MustFindProc("XEle_RegEventC")
    XEle_GetContentSize = XCDLL.MustFindProc("XEle_GetContentSize")
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-7 08:57:09
// @Function: XEleRegEventC
// @Description: 注册事件C方式,省略2参数.
// @Calls: XEle_RegEventC
// @Input: hEle 元素句柄. nEvent 事件类型. pFun 事件函数指针.
// @Return:
// *******************************************
func XEleRegEventC(hEle HELE, nEvent int, pFun uintptr) {
    XEle_RegEventC.Call(
        uintptr(hEle),
        uintptr(nEvent),
        pFun)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 17:59:36
// @Function: XEleGetContentSize
// @Description: 注册事件C方式,省略2参数.
// @Calls: XEle_GetContentSize
// @Input: hEle 元素句柄. nEvent 事件类型. pFun 事件函数指针.
// @Return:
// *******************************************
func XEleGetContentSize(hEle HELE) uint16 {
    pSize := make([]uint16, 256)
    XEle_GetContentSize.Call(
        uintptr(hEle),
        uintptr(unsafe.Pointer(&pSize[0])))

    return pSize[0]
}
