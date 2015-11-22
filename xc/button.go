package xc

import (
    "syscall"
    "unsafe"
)

var (
    // Functions
    XBtn_Create            *syscall.Proc
    XBtn_SetCheck          *syscall.Proc
    XBtn_IsCheck           *syscall.Proc
    XBtn_SetStyle          *syscall.Proc
    XBtn_GetStyle          *syscall.Proc
    XBtn_SetState          *syscall.Proc
    XBtn_GetState          *syscall.Proc
    XBtn_GetStateEx        *syscall.Proc
    XBtn_SetType           *syscall.Proc
    XBtn_GetType           *syscall.Proc
    XBtn_SetGroupID        *syscall.Proc
    XBtn_GetGroupID        *syscall.Proc
    XBtn_SetBindEle        *syscall.Proc
    XBtn_GetBindEle        *syscall.Proc
    XBtn_SetTextAlign      *syscall.Proc
    XBtn_GetTextAlign      *syscall.Proc
    XBtn_SetIconAlign      *syscall.Proc
    XBtn_SetOffset         *syscall.Proc
    XBtn_SetOffsetIcon     *syscall.Proc
    XBtn_SetIconSpace      *syscall.Proc
    XBtn_SetText           *syscall.Proc
    XBtn_GetText           *syscall.Proc
    XBtn_SetIcon           *syscall.Proc
    XBtn_SetIconDisable    *syscall.Proc
    XBtn_AddAnimationFrame *syscall.Proc
    XBtn_EnableAnimation   *syscall.Proc
)

func init() {
    XBtn_Create = XCDLL.MustFindProc("XBtn_Create")
    XBtn_SetCheck = XCDLL.MustFindProc("XBtn_SetCheck")
    XBtn_IsCheck = XCDLL.MustFindProc("XBtn_IsCheck")
    XBtn_SetStyle = XCDLL.MustFindProc("XBtn_SetStyle")
    XBtn_GetStyle = XCDLL.MustFindProc("XBtn_GetStyle")
    XBtn_SetState = XCDLL.MustFindProc("XBtn_SetState")
    XBtn_GetState = XCDLL.MustFindProc("XBtn_GetState")
    XBtn_GetStateEx = XCDLL.MustFindProc("XBtn_GetStateEx")
    XBtn_SetType = XCDLL.MustFindProc("XBtn_SetType")
    XBtn_GetType = XCDLL.MustFindProc("XBtn_GetType")
    XBtn_SetGroupID = XCDLL.MustFindProc("XBtn_SetGroupID")
    XBtn_GetGroupID = XCDLL.MustFindProc("XBtn_GetGroupID")
    XBtn_SetBindEle = XCDLL.MustFindProc("XBtn_SetBindEle")
    XBtn_GetBindEle = XCDLL.MustFindProc("XBtn_GetBindEle")
    XBtn_SetTextAlign = XCDLL.MustFindProc("XBtn_SetTextAlign")
    XBtn_GetTextAlign = XCDLL.MustFindProc("XBtn_GetTextAlign")
    XBtn_SetIconAlign = XCDLL.MustFindProc("XBtn_SetIconAlign")
    XBtn_SetOffset = XCDLL.MustFindProc("XBtn_SetOffset")
    XBtn_SetOffsetIcon = XCDLL.MustFindProc("XBtn_SetOffsetIcon")
    XBtn_SetIconSpace = XCDLL.MustFindProc("XBtn_SetIconSpace")
    XBtn_SetText = XCDLL.MustFindProc("XBtn_SetText")
    XBtn_GetText = XCDLL.MustFindProc("XBtn_GetText")
    XBtn_SetIcon = XCDLL.MustFindProc("XBtn_SetIcon")
    XBtn_SetIconDisable = XCDLL.MustFindProc("XBtn_SetIconDisable")
    XBtn_AddAnimationFrame = XCDLL.MustFindProc("XBtn_AddAnimationFrame")
    XBtn_EnableAnimation = XCDLL.MustFindProc("XBtn_EnableAnimation")

}

// 关闭
func CloseBtn(hWindow HWINDOW) {
    XBtnSetType(XBtnCreate(10, 5, 35, 20, "关闭", HXCGUI(hWindow)),
        BUTTON_TYPE_CLOSE)
}

// 最小化
func MinBtn(hWindow HWINDOW) {
    XBtnSetType(XBtnCreate(60, 5, 45, 20, "最小化", HXCGUI(hWindow)),
        BUTTON_TYPE_MIN)
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
// @Date: 2015-11-8 20:45:55
// @Function: XBtnSetCheck
// @Description: 设置选中状态.
// @Calls: XBtn_SetCheck
// @Input: hEle 元素句柄. bCheck 是否选中.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XBtnSetCheck(hEle HELE, bCheck BOOL) bool {
    ret, _, _ := XBtn_SetCheck.Call(
        uintptr(hEle),
        uintptr(bCheck))

    if ret != TRUE {
        return false
    }

    return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 20:48:18
// @Function: XBtnIsCheck
// @Description: 是否选中状态.
// @Calls: XBtn_IsCheck
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XBtnIsCheck(hEle HELE) bool {
    ret, _, _ := XBtn_IsCheck.Call(uintptr(hEle))

    if ret != TRUE {
        return false
    }

    return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 20:49:31
// @Function: XBtnSetStyle
// @Description: 设置按钮样式.
// @Calls: XBtn_SetStyle
// @Input: hEle 元素句柄. nStyle 样式.
// @Return:
// *******************************************
func XBtnSetStyle(hEle HELE, nStyle BUTTON_STYLE_) {
    XBtn_SetStyle.Call(
        uintptr(hEle),
        uintptr(nStyle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 20:53:00
// @Function: XBtnGetStyle
// @Description: 获取按钮样式.
// @Calls: XBtn_GetStyle
// @Input: hEle 元素句柄.
// @Return: 按钮样式,参见宏定义.
// *******************************************
func XBtnGetStyle(hEle HELE) BUTTON_STYLE_ {
    ret, _, _ := XBtn_GetStyle.Call(uintptr(hEle))

    return BUTTON_STYLE_(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 20:54:58
// @Function: XBtnSetState
// @Description: 设置按钮状态.
// @Calls: XBtn_SetState
// @Input: hEle 元素句柄. nState 按钮状态见宏定义.
// @Return:
// *******************************************
func XBtnSetState(hEle HELE, nState COMMON_STATE3_) {
    XBtn_SetState.Call(
        uintptr(hEle),
        uintptr(nState))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 20:56:53
// @Function: XBtnGetState
// @Description: 获取按钮状态.
// @Calls: XBtn_GetState
// @Input: hEle 元素句柄.
// @Return: 返回按钮状态,参见宏定义.
// *******************************************
func XBtnGetState(hEle HELE) COMMON_STATE3_ {
    ret, _, _ := XBtn_GetState.Call(uintptr(hEle))

    return COMMON_STATE3_(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-22 22:01:37
// @Function: XBtnGetState
// @Description: 获取按钮状态.
// @Calls: XBtnGetStateEx
// @Input: hEle 元素句柄.
// @Return: 返回按钮状态,参见宏定义 button_state_ .
// *******************************************
func XBtnGetStateEx(hEle HELE) BUTTON_STATE_ {
    ret, _, _ := XBtn_GetStateEx.Call(uintptr(hEle))

    return BUTTON_STATE_(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-6 18:36:20
// @Function: XBtnSetType
// @Description: 设置按钮样式.
// @Calls: XBtn_SetType
// @Input: hEle 元素句柄. nStyle 样式. 参考 button_type_
// @Return:
// *******************************************
func XBtnSetType(hEle HELE, nType BUTTON_TYPE_) {
    XBtn_SetType.Call(
        uintptr(hEle),
        uintptr(nType))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 20:59:12
// @Function: XBtnGetType
// @Description: 获取按钮功能类型.
// @Calls: XBtn_GetType
// @Input: hEle 元素句柄.
// @Return: 按钮类型,参见宏定义 button_type_.
// *******************************************
func XBtnGetType(hEle HELE) BUTTON_TYPE_ {
    ret, _, _ := XBtn_GetType.Call(uintptr(hEle))

    return BUTTON_TYPE_(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:00:25
// @Function: XBtnSetGroupID
// @Description: 设置组ID.
// @Calls: XBtn_SetGroupID
// @Input: hEle 元素句柄. nID 组ID.
// @Return:
// *******************************************
func XBtnSetGroupID(hEle HELE, nID int) {
    XBtn_SetGroupID.Call(
        uintptr(hEle),
        uintptr(nID))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:02:37
// @Function: XBtnGetGroupID
// @Description: 获取组ID.
// @Calls: XBtn_GetGroupID
// @Input: hEle 元素句柄.
// @Return: 组ID.
// *******************************************
func XBtnGetGroupID(hEle HELE) int {
    ret, _, _ := XBtn_GetGroupID.Call(uintptr(hEle))

    return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:04:09
// @Function: XBtnSetBindEle
// @Description: 设置绑定元素.
// @Calls: XBtn_SetBindEle
// @Input: hEle 元素句柄. hBindEle 将要绑定的元素.
// @Return:
// *******************************************
func XBtnSetBindEle(hEle, hBindEle HELE) {
    XBtn_SetBindEle.Call(
        uintptr(hEle),
        uintptr(hBindEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:05:41
// @Function: XBtnGetGroupID
// @Description: 获取绑定的元素.
// @Calls: XBtn_GetBindEle
// @Input: hEle 元素句柄.
// @Return: 绑定的元素句柄.
// *******************************************
func XBtnGetBindEle(hEle HELE) HELE {
    ret, _, _ := XBtn_GetBindEle.Call(uintptr(hEle))

    return HELE(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:06:52
// @Function: XBtnSetTextAlign
// @Description: 设置文本对齐方式.
// @Calls: XBtn_SetTextAlign
// @Input: hEle 元素句柄. nFlags 对齐方式.
// @Return:
// *******************************************
func XBtnSetTextAlign(hEle HELE, nFlags int) {
    XBtn_SetTextAlign.Call(
        uintptr(hEle),
        uintptr(nFlags))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:08:37
// @Function: XBtnGetTextAlign
// @Description: 获取文本对齐方式.
// @Calls: XBtn_GetTextAlign
// @Input: hEle 元素句柄.
// @Return: 文本对齐方式.
// *******************************************
func XBtnGetTextAlign(hEle HELE) int {
    ret, _, _ := XBtn_GetTextAlign.Call(uintptr(hEle))

    return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:10:27
// @Function: XBtnSetIconAlign
// @Description: 设置图标对齐.
// @Calls: XBtn_SetIconAlign
// @Input: hEle 元素句柄. align 对齐方式.
// @Return:
// *******************************************
func XBtnSetIconAlign(hEle HELE, align BUTTON_ICON_ALIGN_) {
    XBtn_SetIconAlign.Call(
        uintptr(hEle),
        uintptr(align))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:11:55
// @Function: XBtnSetOffset
// @Description: 设置按钮文本坐标偏移量.
// @Calls: XBtn_SetOffset
// @Input: hEle 元素句柄. x 偏移量. y 偏移量.
// @Return:
// *******************************************
func XBtnSetOffset(hEle HELE, x, y int) {
    XBtn_SetOffset.Call(
        uintptr(hEle),
        uintptr(x),
        uintptr(y))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:13:30
// @Function: XBtnSetOffsetIcon
// @Description: 设置按钮图标坐标偏移量.
// @Calls: XBtn_SetOffsetIcon
// @Input: hEle 元素句柄. x 偏移量. y 偏移量.
// @Return:
// *******************************************
func XBtnSetOffsetIcon(hEle HELE, x, y int) {
    XBtn_SetOffsetIcon.Call(
        uintptr(hEle),
        uintptr(x),
        uintptr(y))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:13:30
// @Function: XBtnSetIconSpace
// @Description: 设置图标与文本间隔大小.
// @Calls: XBtn_SetIconSpace
// @Input: hEle 元素句柄. size 间隔大小.
// @Return:
// *******************************************
func XBtnSetIconSpace(hEle HELE, size int) {
    XBtn_SetIconSpace.Call(
        uintptr(hEle),
        uintptr(size))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:17:40
// @Function: XBtnSetText
// @Description: 设置文本内容.
// @Calls: XBtn_SetText
// @Input: hEle 元素句柄. pName 文本内容.
// @Return:
// *******************************************
func XBtnSetText(hEle HELE, pName string) {
    XBtn_SetText.Call(
        uintptr(hEle),
        StringToUintPtr(pName))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:08:37
// @Function: XBtnGetText
// @Description: 获取文本内容.
// @Calls: XBtn_GetText
// @Input: hEle 元素句柄. pOut 接收文本内容. nOutLen 接收缓冲区长度,字符为单位.
// @Return:
// *******************************************
func XBtnGetText(hEle HELE, pOut *string, nOutLen int) {
    XBtn_GetText.Call(
        uintptr(hEle),
        uintptr(unsafe.Pointer(pOut)),
        uintptr(nOutLen))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:24:11
// @Function: XBtnSetIcon
// @Description: 设置图标.
// @Calls: XBtn_SetIcon
// @Input: hEle 元素句柄. hImage 图片句柄.
// @Return:
// *******************************************
func XBtnSetIcon(hEle HELE, hImage HIMAGE) {
    XBtn_SetIcon.Call(
        uintptr(hEle),
        uintptr(hImage))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:26:12
// @Function: XBtnSetIconDisable
// @Description: 设置图标禁用状态.
// @Calls: XBtn_SetIconDisable
// @Input: hEle 元素句柄. hImage 图片句柄.
// @Return:
// *******************************************
func XBtnSetIconDisable(hEle HELE, hImage HIMAGE) {
    XBtn_SetIconDisable.Call(
        uintptr(hEle),
        uintptr(hImage))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:27:16
// @Function: XBtnAddAnimationFrame
// @Description: 添加动画帧.
// @Calls: XBtn_AddAnimationFrame
// @Input: hEle 元素句柄. hImage 图片句柄. uElapse 图片帧延时,单位毫秒.
// @Return:
// *******************************************
func XBtnAddAnimationFrame(hEle HELE, hImage HIMAGE, uElapse uint) {
    XBtn_AddAnimationFrame.Call(
        uintptr(hEle),
        uintptr(hImage),
        uintptr(uElapse))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:28:44
// @Function: XBtnEnableAnimation
// @Description: 开始或关闭图片动画的播放.
// @Calls: XBtn_EnableAnimation
// @Input: hEle 元素句柄. bEnable 开始播放动画TRUE,关闭播放动画FALSE. bLoopPlay 是否循环播放.
// @Return:
// *******************************************
func XBtnEnableAnimation(hEle HELE, bEnable, bLoopPlay BOOL) {
    XBtn_EnableAnimation.Call(
        uintptr(hEle),
        uintptr(bEnable),
        uintptr(bLoopPlay))
}
