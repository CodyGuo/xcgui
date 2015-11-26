package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xBtn_Create            *syscall.Proc
	xBtn_SetCheck          *syscall.Proc
	xBtn_IsCheck           *syscall.Proc
	xBtn_SetStyle          *syscall.Proc
	xBtn_GetStyle          *syscall.Proc
	xBtn_SetState          *syscall.Proc
	xBtn_GetState          *syscall.Proc
	xBtn_GetStateEx        *syscall.Proc
	xBtn_SetType           *syscall.Proc
	xBtn_GetType           *syscall.Proc
	xBtn_SetGroupID        *syscall.Proc
	xBtn_GetGroupID        *syscall.Proc
	xBtn_SetBindEle        *syscall.Proc
	xBtn_GetBindEle        *syscall.Proc
	xBtn_SetTextAlign      *syscall.Proc
	xBtn_GetTextAlign      *syscall.Proc
	xBtn_SetIconAlign      *syscall.Proc
	xBtn_SetOffset         *syscall.Proc
	xBtn_SetOffsetIcon     *syscall.Proc
	xBtn_SetIconSpace      *syscall.Proc
	xBtn_SetText           *syscall.Proc
	xBtn_GetText           *syscall.Proc
	xBtn_SetIcon           *syscall.Proc
	xBtn_SetIconDisable    *syscall.Proc
	xBtn_AddAnimationFrame *syscall.Proc
	xBtn_EnableAnimation   *syscall.Proc
	xBtn_AddBkBorder       *syscall.Proc
	xBtn_AddBkFill         *syscall.Proc
	xBtn_AddBkImage        *syscall.Proc
	xBtn_GetBkInfoCount    *syscall.Proc
	xBtn_ClearBkInfo       *syscall.Proc
	xBtn_GetBkInfoManager  *syscall.Proc
)

func init() {
	xBtn_Create = XCDLL.MustFindProc("XBtn_Create")
	xBtn_SetCheck = XCDLL.MustFindProc("XBtn_SetCheck")
	xBtn_IsCheck = XCDLL.MustFindProc("XBtn_IsCheck")
	xBtn_SetStyle = XCDLL.MustFindProc("XBtn_SetStyle")
	xBtn_GetStyle = XCDLL.MustFindProc("XBtn_GetStyle")
	xBtn_SetState = XCDLL.MustFindProc("XBtn_SetState")
	xBtn_GetState = XCDLL.MustFindProc("XBtn_GetState")
	xBtn_GetStateEx = XCDLL.MustFindProc("XBtn_GetStateEx")
	xBtn_SetType = XCDLL.MustFindProc("XBtn_SetType")
	xBtn_GetType = XCDLL.MustFindProc("XBtn_GetType")
	xBtn_SetGroupID = XCDLL.MustFindProc("XBtn_SetGroupID")
	xBtn_GetGroupID = XCDLL.MustFindProc("XBtn_GetGroupID")
	xBtn_SetBindEle = XCDLL.MustFindProc("XBtn_SetBindEle")
	xBtn_GetBindEle = XCDLL.MustFindProc("XBtn_GetBindEle")
	xBtn_SetTextAlign = XCDLL.MustFindProc("XBtn_SetTextAlign")
	xBtn_GetTextAlign = XCDLL.MustFindProc("XBtn_GetTextAlign")
	xBtn_SetIconAlign = XCDLL.MustFindProc("XBtn_SetIconAlign")
	xBtn_SetOffset = XCDLL.MustFindProc("XBtn_SetOffset")
	xBtn_SetOffsetIcon = XCDLL.MustFindProc("XBtn_SetOffsetIcon")
	xBtn_SetIconSpace = XCDLL.MustFindProc("XBtn_SetIconSpace")
	xBtn_SetText = XCDLL.MustFindProc("XBtn_SetText")
	xBtn_GetText = XCDLL.MustFindProc("XBtn_GetText")
	xBtn_SetIcon = XCDLL.MustFindProc("XBtn_SetIcon")
	xBtn_SetIconDisable = XCDLL.MustFindProc("XBtn_SetIconDisable")
	xBtn_AddAnimationFrame = XCDLL.MustFindProc("XBtn_AddAnimationFrame")
	xBtn_EnableAnimation = XCDLL.MustFindProc("XBtn_EnableAnimation")
	xBtn_AddBkBorder = XCDLL.MustFindProc("XBtn_AddBkBorder")
	xBtn_AddBkFill = XCDLL.MustFindProc("XBtn_AddBkFill")
	xBtn_AddBkImage = XCDLL.MustFindProc("XBtn_AddBkImage")
	xBtn_GetBkInfoCount = XCDLL.MustFindProc("XBtn_GetBkInfoCount")
	xBtn_ClearBkInfo = XCDLL.MustFindProc("XBtn_ClearBkInfo")
	xBtn_GetBkInfoManager = XCDLL.MustFindProc("XBtn_GetBkInfoManager")

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
// @Calls: xBtn_Create
// @Input: x 按钮x坐标. y 按钮y坐标. cx 宽度. cy 高度. pName 标题.
//         hParent 父是窗口句柄或元素句柄.
// @Return: 按钮元素句柄.
// *******************************************
func XBtnCreate(x, y, cx, cy int, pName string, hParent HXCGUI) HELE {
	ret, _, _ := xBtn_Create.Call(
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
// @Calls: xBtn_SetCheck
// @Input: hEle 元素句柄. bCheck 是否选中.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XBtnSetCheck(hEle HELE, bCheck BOOL) bool {
	ret, _, _ := xBtn_SetCheck.Call(
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
// @Calls: xBtn_IsCheck
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XBtnIsCheck(hEle HELE) bool {
	ret, _, _ := xBtn_IsCheck.Call(uintptr(hEle))

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
// @Calls: xBtn_SetStyle
// @Input: hEle 元素句柄. nStyle 样式.
// @Return:
// *******************************************
func XBtnSetStyle(hEle HELE, nStyle BUTTON_STYLE_) {
	xBtn_SetStyle.Call(
		uintptr(hEle),
		uintptr(nStyle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 20:53:00
// @Function: XBtnGetStyle
// @Description: 获取按钮样式.
// @Calls: xBtn_GetStyle
// @Input: hEle 元素句柄.
// @Return: 按钮样式,参见宏定义.
// *******************************************
func XBtnGetStyle(hEle HELE) BUTTON_STYLE_ {
	ret, _, _ := xBtn_GetStyle.Call(uintptr(hEle))

	return BUTTON_STYLE_(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 20:54:58
// @Function: XBtnSetState
// @Description: 设置按钮状态.
// @Calls: xBtn_SetState
// @Input: hEle 元素句柄. nState 按钮状态见宏定义.
// @Return:
// *******************************************
func XBtnSetState(hEle HELE, nState COMMON_STATE3_) {
	xBtn_SetState.Call(
		uintptr(hEle),
		uintptr(nState))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 20:56:53
// @Function: XBtnGetState
// @Description: 获取按钮状态.
// @Calls: xBtn_GetState
// @Input: hEle 元素句柄.
// @Return: 返回按钮状态,参见宏定义.
// *******************************************
func XBtnGetState(hEle HELE) COMMON_STATE3_ {
	ret, _, _ := xBtn_GetState.Call(uintptr(hEle))

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
	ret, _, _ := xBtn_GetStateEx.Call(uintptr(hEle))

	return BUTTON_STATE_(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-6 18:36:20
// @Function: XBtnSetType
// @Description: 设置按钮样式.
// @Calls: xBtn_SetType
// @Input: hEle 元素句柄. nStyle 样式. 参考 button_type_
// @Return:
// *******************************************
func XBtnSetType(hEle HELE, nType BUTTON_TYPE_) {
	xBtn_SetType.Call(
		uintptr(hEle),
		uintptr(nType))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 20:59:12
// @Function: XBtnGetType
// @Description: 获取按钮功能类型.
// @Calls: xBtn_GetType
// @Input: hEle 元素句柄.
// @Return: 按钮类型,参见宏定义 button_type_.
// *******************************************
func XBtnGetType(hEle HELE) BUTTON_TYPE_ {
	ret, _, _ := xBtn_GetType.Call(uintptr(hEle))

	return BUTTON_TYPE_(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:00:25
// @Function: XBtnSetGroupID
// @Description: 设置组ID.
// @Calls: xBtn_SetGroupID
// @Input: hEle 元素句柄. nID 组ID.
// @Return:
// *******************************************
func XBtnSetGroupID(hEle HELE, nID int) {
	xBtn_SetGroupID.Call(
		uintptr(hEle),
		uintptr(nID))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:02:37
// @Function: XBtnGetGroupID
// @Description: 获取组ID.
// @Calls: xBtn_GetGroupID
// @Input: hEle 元素句柄.
// @Return: 组ID.
// *******************************************
func XBtnGetGroupID(hEle HELE) int {
	ret, _, _ := xBtn_GetGroupID.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:04:09
// @Function: XBtnSetBindEle
// @Description: 设置绑定元素.
// @Calls: xBtn_SetBindEle
// @Input: hEle 元素句柄. hBindEle 将要绑定的元素.
// @Return:
// *******************************************
func XBtnSetBindEle(hEle, hBindEle HELE) {
	xBtn_SetBindEle.Call(
		uintptr(hEle),
		uintptr(hBindEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:05:41
// @Function: XBtnGetGroupID
// @Description: 获取绑定的元素.
// @Calls: xBtn_GetBindEle
// @Input: hEle 元素句柄.
// @Return: 绑定的元素句柄.
// *******************************************
func XBtnGetBindEle(hEle HELE) HELE {
	ret, _, _ := xBtn_GetBindEle.Call(uintptr(hEle))

	return HELE(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:06:52
// @Function: XBtnSetTextAlign
// @Description: 设置文本对齐方式.
// @Calls: xBtn_SetTextAlign
// @Input: hEle 元素句柄. nFlags 对齐方式.
// @Return:
// *******************************************
func XBtnSetTextAlign(hEle HELE, nFlags int) {
	xBtn_SetTextAlign.Call(
		uintptr(hEle),
		uintptr(nFlags))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:08:37
// @Function: XBtnGetTextAlign
// @Description: 获取文本对齐方式.
// @Calls: xBtn_GetTextAlign
// @Input: hEle 元素句柄.
// @Return: 文本对齐方式.
// *******************************************
func XBtnGetTextAlign(hEle HELE) int {
	ret, _, _ := xBtn_GetTextAlign.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:10:27
// @Function: XBtnSetIconAlign
// @Description: 设置图标对齐.
// @Calls: xBtn_SetIconAlign
// @Input: hEle 元素句柄. align 对齐方式.
// @Return:
// *******************************************
func XBtnSetIconAlign(hEle HELE, align BUTTON_ICON_ALIGN_) {
	xBtn_SetIconAlign.Call(
		uintptr(hEle),
		uintptr(align))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:11:55
// @Function: XBtnSetOffset
// @Description: 设置按钮文本坐标偏移量.
// @Calls: xBtn_SetOffset
// @Input: hEle 元素句柄. x 偏移量. y 偏移量.
// @Return:
// *******************************************
func XBtnSetOffset(hEle HELE, x, y int) {
	xBtn_SetOffset.Call(
		uintptr(hEle),
		uintptr(x),
		uintptr(y))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:13:30
// @Function: XBtnSetOffsetIcon
// @Description: 设置按钮图标坐标偏移量.
// @Calls: xBtn_SetOffsetIcon
// @Input: hEle 元素句柄. x 偏移量. y 偏移量.
// @Return:
// *******************************************
func XBtnSetOffsetIcon(hEle HELE, x, y int) {
	xBtn_SetOffsetIcon.Call(
		uintptr(hEle),
		uintptr(x),
		uintptr(y))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:13:30
// @Function: XBtnSetIconSpace
// @Description: 设置图标与文本间隔大小.
// @Calls: xBtn_SetIconSpace
// @Input: hEle 元素句柄. size 间隔大小.
// @Return:
// *******************************************
func XBtnSetIconSpace(hEle HELE, size int) {
	xBtn_SetIconSpace.Call(
		uintptr(hEle),
		uintptr(size))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:17:40
// @Function: XBtnSetText
// @Description: 设置文本内容.
// @Calls: xBtn_SetText
// @Input: hEle 元素句柄. pName 文本内容.
// @Return:
// *******************************************
func XBtnSetText(hEle HELE, pName string) {
	xBtn_SetText.Call(
		uintptr(hEle),
		StringToUintPtr(pName))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:08:37
// @Function: XBtnGetText
// @Description: 获取文本内容.
// @Calls: xBtn_GetText
// @Input: hEle 元素句柄. pOut 接收文本内容. nOutLen 接收缓冲区长度,字符为单位.
// @Return:
// *******************************************
func XBtnGetText(hEle HELE, pOut *uint16, nOutLen int) {
	xBtn_GetText.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:24:11
// @Function: XBtnSetIcon
// @Description: 设置图标.
// @Calls: xBtn_SetIcon
// @Input: hEle 元素句柄. hImage 图片句柄.
// @Return:
// *******************************************
func XBtnSetIcon(hEle HELE, hImage HIMAGE) {
	xBtn_SetIcon.Call(
		uintptr(hEle),
		uintptr(hImage))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:26:12
// @Function: XBtnSetIconDisable
// @Description: 设置图标禁用状态.
// @Calls: xBtn_SetIconDisable
// @Input: hEle 元素句柄. hImage 图片句柄.
// @Return:
// *******************************************
func XBtnSetIconDisable(hEle HELE, hImage HIMAGE) {
	xBtn_SetIconDisable.Call(
		uintptr(hEle),
		uintptr(hImage))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:27:16
// @Function: XBtnAddAnimationFrame
// @Description: 添加动画帧.
// @Calls: xBtn_AddAnimationFrame
// @Input: hEle 元素句柄. hImage 图片句柄. uElapse 图片帧延时,单位毫秒.
// @Return:
// *******************************************
func XBtnAddAnimationFrame(hEle HELE, hImage HIMAGE, uElapse uint) {
	xBtn_AddAnimationFrame.Call(
		uintptr(hEle),
		uintptr(hImage),
		uintptr(uElapse))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-8 21:28:44
// @Function: XBtnEnableAnimation
// @Description: 开始或关闭图片动画的播放.
// @Calls: xBtn_EnableAnimation
// @Input: hEle 元素句柄. bEnable 开始播放动画TRUE,关闭播放动画FALSE. bLoopPlay 是否循环播放.
// @Return:
// *******************************************
func XBtnEnableAnimation(hEle HELE, bEnable, bLoopPlay BOOL) {
	xBtn_EnableAnimation.Call(
		uintptr(hEle),
		uintptr(bEnable),
		uintptr(bLoopPlay))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 18:34:05
// @Function: XBtnAddBkBorder
// @Description: 添加背景内容边框.
// @Calls: xBtn_AddBkBorder
// @Input: hEle 元素句柄. nState 按钮状态. color RGB颜色. alpha 透明度. width 线宽.
// @Return:
// *******************************************
func XBtnAddBkBorder(hEle HELE, nState BUTTON_STATE_, color COLORREF, alpha byte, width int) {
	xBtn_AddBkBorder.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha),
		uintptr(width))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 18:36:46
// @Function: XBtnAddBkFill
// @Description: 添加背景内容填充.
// @Calls: xBtn_AddBkFill
// @Input: hEle 元素句柄. nState 按钮状态. color RGB颜色. alpha 透明度.
// @Return:
// *******************************************
func XBtnAddBkFill(hEle HELE, nState BUTTON_STATE_, color COLORREF, alpha byte) {
	xBtn_AddBkFill.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 18:38:53
// @Function: XBtnAddBkImage
// @Description: 添加背景内容图片.
// @Calls: xBtn_AddBkImage
// @Input: hEle 元素句柄. nState 按钮状态. hImage 图片句柄.
// @Return:
// *******************************************
func XBtnAddBkImage(hEle HELE, nState BUTTON_STATE_, hImage HIMAGE) {
	xBtn_AddBkImage.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(hImage))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 18:40:22
// @Function: XBtnGetBkInfoCount
// @Description: 获取背景内容数量.
// @Calls: xBtn_GetBkInfoCount
// @Input: hEle 元素句柄. nState 按钮状态.
// @Return: 成功返回背景内容数量,否则返回 XC_ID_ERROR.
// *******************************************
func XBtnGetBkInfoCount(hEle HELE, nState BUTTON_STATE_) int {
	ret, _, _ := xBtn_GetBkInfoCount.Call(
		uintptr(hEle),
		uintptr(nState))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 18:43:04
// @Function: XBtnClearBkInfo
// @Description: 清空背景内容; 如果背景没有内容,将使用系统默认内容,以便保证背景正确.
// @Calls: xBtn_ClearBkInfo
// @Input: hEle 元素句柄. nState 按钮状态.
// @Return:
// *******************************************
func XBtnClearBkInfo(hEle HELE, nState BUTTON_STATE_) {
	xBtn_ClearBkInfo.Call(
		uintptr(hEle),
		uintptr(nState))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 18:44:10
// @Function: XBtnGetBkInfoManager
// @Description: 获取背景内容管理器.
// @Calls: xBtn_GetBkInfoManager
// @Input: hEle 元素句柄. nState 按钮状态.
// @Return: 背景内容管理器.
// *******************************************
func XBtnGetBkInfoManager(hEle HELE, nState BUTTON_STATE_) HBKINFOM {
	ret, _, _ := xBtn_GetBkInfoManager.Call(
		uintptr(hEle),
		uintptr(nState))

	return HBKINFOM(ret)
}
