package xc

import (
    "syscall"
)

var (
    // Functions
    XEle_RegEventC *syscall.Proc
)

func init() {
    // Functions
    XEle_RegEventC = XCDLL.MustFindProc("XEle_RegEventC")
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
