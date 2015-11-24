package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	XEle_Create                     *syscall.Proc
	XEle_RegEventC                  *syscall.Proc
	XEle_RegEventC1                 *syscall.Proc
	XEle_RegEventC2                 *syscall.Proc
	XEle_RemoveEventC               *syscall.Proc
	XEle_SendEvent                  *syscall.Proc
	XEle_PostEvent                  *syscall.Proc
	XEle_GetRect                    *syscall.Proc
	XEle_GetRectLogic               *syscall.Proc
	XEle_GetClientRect              *syscall.Proc
	XEle_GetWidth                   *syscall.Proc
	XEle_GetHeight                  *syscall.Proc
	XEle_RectWndClientToEleClient   *syscall.Proc
	XEle_PointWndClientToEleClient  *syscall.Proc
	XEle_RectClientToWndClient      *syscall.Proc
	XEle_PointClientToWndClient     *syscall.Proc
	XEle_GetWndClientRect           *syscall.Proc
	XEle_GetType                    *syscall.Proc
	XEle_GetHWND                    *syscall.Proc
	XEle_GetHWINDOW                 *syscall.Proc
	XEle_GetCursor                  *syscall.Proc
	XEle_SetCursor                  *syscall.Proc
	XEle_AddEle                     *syscall.Proc
	XEle_InsertEle                  *syscall.Proc
	XEle_ShowEle                    *syscall.Proc
	XEle_SetRect                    *syscall.Proc
	XEle_SetRectLogic               *syscall.Proc
	XEle_SetID                      *syscall.Proc
	XEle_GetID                      *syscall.Proc
	XEle_IsShow                     *syscall.Proc
	XEle_IsDrawFocus                *syscall.Proc
	XEle_IsEnable                   *syscall.Proc
	XEle_IsEnableFocus              *syscall.Proc
	XEle_IsMouseThrough             *syscall.Proc
	XEle_HitChildEle                *syscall.Proc
	XEle_IsBkTransparent            *syscall.Proc
	XEle_IsEnableEvent_XE_PAINT_END *syscall.Proc
	XEle_IsKeyTab                   *syscall.Proc
	XEle_IsSwitchFocus              *syscall.Proc
	XEle_IsEnable_XE_MOUSEWHEEL     *syscall.Proc
	XEle_Enable                     *syscall.Proc
	XEle_EnableFocus                *syscall.Proc
	XEle_EnableDrawFocus            *syscall.Proc
	XEle_EnableEvent_XE_PAINT_END   *syscall.Proc
	XEle_EnableBkTransparent        *syscall.Proc
	XEle_EnableMouseThrough         *syscall.Proc
	XEle_EnableKeyTab               *syscall.Proc
	XEle_EnableSwitchFocus          *syscall.Proc
	XEle_EnableEvent_XE_MOUSEWHEEL  *syscall.Proc
	XEle_GetParentEle               *syscall.Proc
	XEle_GetParent                  *syscall.Proc
	XEle_RemoveEle                  *syscall.Proc
	XEle_RedrawEle                  *syscall.Proc
	XEle_RedrawRect                 *syscall.Proc
	XEle_GetChildCount              *syscall.Proc
	XEle_GetChildByIndex            *syscall.Proc
	XEle_GetChildByID               *syscall.Proc
	XEle_GetChildShapeCount         *syscall.Proc
	XEle_GetChildShapeByIndex       *syscall.Proc
	XEle_SetTextColor               *syscall.Proc
	XEle_GetTextColor               *syscall.Proc
	XEle_SetFocusBorderColor        *syscall.Proc
	XEle_GetFocusBorderColor        *syscall.Proc
	XEle_SetFont                    *syscall.Proc
	XEle_GetFont                    *syscall.Proc
	XEle_SetAlpha                   *syscall.Proc
	XEle_Destroy                    *syscall.Proc
	XEle_AddBkBorder                *syscall.Proc
	XEle_AddBkFill                  *syscall.Proc
	XEle_AddBkImage                 *syscall.Proc
	XEle_GetBkInfoCount             *syscall.Proc
	XEle_ClearBkInfo                *syscall.Proc
	XEle_GetBkInfoManager           *syscall.Proc
	XEle_DrawFocus                  *syscall.Proc
	XEle_BindLayoutObject           *syscall.Proc
	XEle_GetLayoutObject            *syscall.Proc
	XEle_GetParentLayoutObject      *syscall.Proc
	XEle_SetUserData                *syscall.Proc
	XEle_GetUserData                *syscall.Proc
	XEle_GetContentSize             *syscall.Proc
	XEle_SetCapture                 *syscall.Proc
	XEle_SetLayoutWidth             *syscall.Proc
	XEle_SetLayoutHeight            *syscall.Proc
	XEle_GetLayoutWidth             *syscall.Proc
	XEle_GetLayoutHeight            *syscall.Proc
	XEle_EnableTransparentChannel   *syscall.Proc
	XEle_SetToolTip                 *syscall.Proc
	XEle_GetToolTip                 *syscall.Proc
	XEle_EnableToolTip              *syscall.Proc
	XEle_AdjustLayoutObject         *syscall.Proc
	XEle_AdjustLayout               *syscall.Proc
)

func init() {
	// Functions
	XEle_Create = XCDLL.MustFindProc("XEle_Create")
	XEle_RegEventC = XCDLL.MustFindProc("XEle_RegEventC")
	XEle_RegEventC1 = XCDLL.MustFindProc("XEle_RegEventC1")
	XEle_RegEventC2 = XCDLL.MustFindProc("XEle_RegEventC2")
	XEle_RemoveEventC = XCDLL.MustFindProc("XEle_RemoveEventC")
	XEle_SendEvent = XCDLL.MustFindProc("XEle_SendEvent")
	XEle_PostEvent = XCDLL.MustFindProc("XEle_PostEvent")
	XEle_GetRect = XCDLL.MustFindProc("XEle_GetRect")
	XEle_GetRectLogic = XCDLL.MustFindProc("XEle_GetRectLogic")
	XEle_GetClientRect = XCDLL.MustFindProc("XEle_GetClientRect")
	XEle_GetWidth = XCDLL.MustFindProc("XEle_GetWidth")
	XEle_GetHeight = XCDLL.MustFindProc("XEle_GetHeight")
	XEle_RectWndClientToEleClient = XCDLL.MustFindProc("XEle_RectWndClientToEleClient")
	XEle_PointWndClientToEleClient = XCDLL.MustFindProc("XEle_PointWndClientToEleClient")
	XEle_RectClientToWndClient = XCDLL.MustFindProc("XEle_RectClientToWndClient")
	XEle_PointClientToWndClient = XCDLL.MustFindProc("XEle_PointClientToWndClient")
	XEle_GetWndClientRect = XCDLL.MustFindProc("XEle_GetWndClientRect")
	XEle_GetType = XCDLL.MustFindProc("XEle_GetType")
	XEle_GetHWND = XCDLL.MustFindProc("XEle_GetHWND")
	XEle_GetHWINDOW = XCDLL.MustFindProc("XEle_GetHWINDOW")
	XEle_GetCursor = XCDLL.MustFindProc("XEle_GetCursor")
	XEle_SetCursor = XCDLL.MustFindProc("XEle_SetCursor")
	XEle_AddEle = XCDLL.MustFindProc("XEle_AddEle")
	XEle_InsertEle = XCDLL.MustFindProc("XEle_InsertEle")
	XEle_ShowEle = XCDLL.MustFindProc("XEle_ShowEle")
	XEle_SetRect = XCDLL.MustFindProc("XEle_SetRect")
	XEle_SetRectLogic = XCDLL.MustFindProc("XEle_SetRectLogic")
	XEle_SetID = XCDLL.MustFindProc("XEle_SetID")
	XEle_GetID = XCDLL.MustFindProc("XEle_GetID")
	XEle_IsShow = XCDLL.MustFindProc("XEle_IsShow")
	XEle_IsDrawFocus = XCDLL.MustFindProc("XEle_IsDrawFocus")
	XEle_IsEnable = XCDLL.MustFindProc("XEle_IsEnable")
	XEle_IsEnableFocus = XCDLL.MustFindProc("XEle_IsEnableFocus")
	XEle_IsMouseThrough = XCDLL.MustFindProc("XEle_IsMouseThrough")
	XEle_HitChildEle = XCDLL.MustFindProc("XEle_HitChildEle")
	XEle_IsBkTransparent = XCDLL.MustFindProc("XEle_IsBkTransparent")
	XEle_IsEnableEvent_XE_PAINT_END = XCDLL.MustFindProc("XEle_IsEnableEvent_XE_PAINT_END")
	XEle_IsKeyTab = XCDLL.MustFindProc("XEle_IsKeyTab")
	XEle_IsSwitchFocus = XCDLL.MustFindProc("XEle_IsSwitchFocus")
	XEle_IsEnable_XE_MOUSEWHEEL = XCDLL.MustFindProc("XEle_IsEnable_XE_MOUSEWHEEL")
	XEle_Enable = XCDLL.MustFindProc("XEle_Enable")
	XEle_EnableFocus = XCDLL.MustFindProc("XEle_EnableFocus")
	XEle_EnableDrawFocus = XCDLL.MustFindProc("XEle_EnableDrawFocus")
	XEle_EnableEvent_XE_PAINT_END = XCDLL.MustFindProc("XEle_EnableEvent_XE_PAINT_END")
	XEle_EnableBkTransparent = XCDLL.MustFindProc("XEle_EnableBkTransparent")
	XEle_EnableMouseThrough = XCDLL.MustFindProc("XEle_EnableMouseThrough")
	XEle_EnableKeyTab = XCDLL.MustFindProc("XEle_EnableKeyTab")
	XEle_EnableSwitchFocus = XCDLL.MustFindProc("XEle_EnableSwitchFocus")
	XEle_EnableEvent_XE_MOUSEWHEEL = XCDLL.MustFindProc("XEle_EnableEvent_XE_MOUSEWHEEL")
	XEle_GetParentEle = XCDLL.MustFindProc("XEle_GetParentEle")
	XEle_GetParent = XCDLL.MustFindProc("XEle_GetParent")
	XEle_RemoveEle = XCDLL.MustFindProc("XEle_RemoveEle")
	XEle_RedrawEle = XCDLL.MustFindProc("XEle_RedrawEle")
	XEle_RedrawRect = XCDLL.MustFindProc("XEle_RedrawRect")
	XEle_GetChildCount = XCDLL.MustFindProc("XEle_GetChildCount")
	XEle_GetChildByIndex = XCDLL.MustFindProc("XEle_GetChildByIndex")
	XEle_GetChildByID = XCDLL.MustFindProc("XEle_GetChildByID")
	XEle_GetChildShapeCount = XCDLL.MustFindProc("XEle_GetChildShapeCount")
	XEle_GetChildShapeByIndex = XCDLL.MustFindProc("XEle_GetChildShapeByIndex")
	XEle_SetTextColor = XCDLL.MustFindProc("XEle_SetTextColor")
	XEle_GetTextColor = XCDLL.MustFindProc("XEle_GetTextColor")
	XEle_SetFocusBorderColor = XCDLL.MustFindProc("XEle_SetFocusBorderColor")
	XEle_GetFocusBorderColor = XCDLL.MustFindProc("XEle_GetFocusBorderColor")
	XEle_SetFont = XCDLL.MustFindProc("XEle_SetFont")
	XEle_GetFont = XCDLL.MustFindProc("XEle_GetFont")
	XEle_SetAlpha = XCDLL.MustFindProc("XEle_SetAlpha")
	XEle_Destroy = XCDLL.MustFindProc("XEle_Destroy")
	XEle_AddBkBorder = XCDLL.MustFindProc("XEle_AddBkBorder")
	XEle_AddBkFill = XCDLL.MustFindProc("XEle_AddBkFill")
	XEle_AddBkImage = XCDLL.MustFindProc("XEle_AddBkImage")
	XEle_GetBkInfoCount = XCDLL.MustFindProc("XEle_GetBkInfoCount")
	XEle_ClearBkInfo = XCDLL.MustFindProc("XEle_ClearBkInfo")
	XEle_GetBkInfoManager = XCDLL.MustFindProc("XEle_GetBkInfoManager")
	XEle_DrawFocus = XCDLL.MustFindProc("XEle_DrawFocus")
	XEle_BindLayoutObject = XCDLL.MustFindProc("XEle_BindLayoutObject")
	XEle_GetLayoutObject = XCDLL.MustFindProc("XEle_GetLayoutObject")
	XEle_GetParentLayoutObject = XCDLL.MustFindProc("XEle_GetParentLayoutObject")
	XEle_SetUserData = XCDLL.MustFindProc("XEle_SetUserData")
	XEle_GetUserData = XCDLL.MustFindProc("XEle_GetUserData")
	XEle_GetContentSize = XCDLL.MustFindProc("XEle_GetContentSize")
	XEle_SetCapture = XCDLL.MustFindProc("XEle_SetCapture")
	XEle_SetLayoutWidth = XCDLL.MustFindProc("XEle_SetLayoutWidth")
	XEle_SetLayoutHeight = XCDLL.MustFindProc("XEle_SetLayoutHeight")
	XEle_GetLayoutWidth = XCDLL.MustFindProc("XEle_GetLayoutWidth")
	XEle_GetLayoutHeight = XCDLL.MustFindProc("XEle_GetLayoutHeight")
	XEle_EnableTransparentChannel = XCDLL.MustFindProc("XEle_EnableTransparentChannel")
	XEle_SetToolTip = XCDLL.MustFindProc("XEle_SetToolTip")
	XEle_GetToolTip = XCDLL.MustFindProc("XEle_GetToolTip")
	XEle_EnableToolTip = XCDLL.MustFindProc("XEle_EnableToolTip")
	XEle_AdjustLayoutObject = XCDLL.MustFindProc("XEle_AdjustLayoutObject")
	XEle_AdjustLayout = XCDLL.MustFindProc("XEle_AdjustLayout")
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 20:42:40
// @Function: XEleCreate
// @Description: 创建基础元素.
// @Calls: XEle_Create
// @Input: x 元素x坐标. y 元素y坐标. cx 宽度. cy 高度. hParent 父是窗口句柄或元素句柄.
// @Return: 元素句柄.
// *******************************************
func XEleCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := XEle_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
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
// @Date: 2015-11-23 20:47:27
// @Function: XEleRegEventC1
// @Description: 注册事件C方式,省略1参数.
// @Calls: XEle_RegEventC1
// @Input: hEle 元素句柄. nEvent 事件类型. pFun 事件函数指针.
// @Return:
// *******************************************
func XEleRegEventC1(hEle HELE, nEvent int, pFun uintptr) {
	XEle_RegEventC1.Call(
		uintptr(hEle),
		uintptr(nEvent),
		pFun)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 20:49:26
// @Function: XEleRegEventC2
// @Description: 注册事件C方式,不省略参数.
// @Calls: XEle_RegEventC2
// @Input: hEle 元素句柄. nEvent 事件类型. pFun 事件函数指针.
// @Return:
// *******************************************
func XEleRegEventC2(hEle HELE, nEvent int, pFun uintptr) {
	XEle_RegEventC2.Call(
		uintptr(hEle),
		uintptr(nEvent),
		pFun)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 20:50:37
// @Function: XEleRemoveEventC
// @Description: 移除事件函数C方式.
// @Calls: XEle_RemoveEventC
// @Input: hEle 元素句柄. nEvent 事件类型. pFun 事件函数指针.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XEleRemoveEventC(hEle HELE, nEvent int, pFun uintptr) bool {
	ret, _, _ := XEle_RemoveEventC.Call(
		uintptr(hEle),
		uintptr(nEvent),
		pFun)

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 22:49:11
// @Function: XEleSendEvent
// @Description: 发送事件.
// @Calls: XEle_SendEvent
// @Input: hEle 目标元素句柄. hEventEle 触发事件元素句柄. nEvent 事件类型. wParam 参数. lParam 参数.
// @Return: 事件返回值.
// *******************************************
func XEleSendEvent(hEle, hEventEle HELE, nEvent int, wParam, lParam uintptr) int {
	ret, _, _ := XEle_SendEvent.Call(
		uintptr(hEle),
		uintptr(hEventEle),
		uintptr(nEvent),
		wParam,
		lParam)

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 22:54:46
// @Function: XElePostEvent
// @Description: POST事件.
// @Calls: XEle_PostEvent
// @Input: hEle 目标元素句柄. hEventEle 触发事件元素句柄. nEvent 事件类型. wParam 参数. lParam 参数.
// @Return: 事件返回值.
// *******************************************
func XElePostEvent(hEle, hEventEle HELE, nEvent int, wParam, lParam uintptr) int {
	ret, _, _ := XEle_PostEvent.Call(
		uintptr(hEle),
		uintptr(hEventEle),
		uintptr(nEvent),
		wParam,
		lParam)

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 22:56:19
// @Function: XEleGetRect
// @Description: 获取元素坐标.
// @Calls: XEle_GetRect
// @Input: hEle 元素句柄. pRect 坐标.
// @Return:
// *******************************************
func XEleGetRect(hEle HELE, pRect *RECT) {
	XEle_GetRect.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 22:58:51
// @Function: XEleGetRectLogic
// @Description: 获取元素坐标,逻辑坐标,包含滚动视图偏移.
// @Calls: XEle_GetRectLogic
// @Input: hEle 元素句柄. pRect 坐标.
// @Return:
// *******************************************
func XEleGetRectLogic(hEle HELE, pRect *RECT) {
	XEle_GetRectLogic.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 22:59:59
// @Function: XEleGetClientRect
// @Description: 获取元素客户区坐标.
// @Calls: XEle_GetClientRect
// @Input: hEle 元素句柄. pRect 坐标.
// @Return:
// *******************************************
func XEleGetClientRect(hEle HELE, pRect *RECT) {
	XEle_GetClientRect.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:01:30
// @Function: XEleGetWidth
// @Description: 获取元素宽度.
// @Calls: XEle_GetWidth
// @Input: hEle 元素句柄.
// @Return: 宽度.
// *******************************************
func XEleGetWidth(hEle HELE) int {
	ret, _, _ := XEle_GetWidth.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:03:14
// @Function: XEleGetHeight
// @Description: 获取元素高度.
// @Calls: XEle_GetHeight
// @Input: hEle 元素句柄.
// @Return: 高度.
// *******************************************
func XEleGetHeight(hEle HELE) int {
	ret, _, _ := XEle_GetHeight.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:04:35
// @Function: XEleRectWndClientToEleClient
// @Description: 窗口客户区坐标转换到元素客户区坐标.
// @Calls: XEle_RectWndClientToEleClient
// @Input: hEle 元素句柄. pRect 坐标.
// @Return:
// *******************************************
func XEleRectWndClientToEleClient(hEle HELE, pRect *RECT) {
	XEle_RectWndClientToEleClient.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:07:01
// @Function: XElePointWndClientToEleClient
// @Description: 窗口客户区坐标转换到元素客户区坐标.
// @Calls: XEle_PointWndClientToEleClient
// @Input: hEle 元素句柄. pRect 坐标.
// @Return:
// *******************************************
func XElePointWndClientToEleClient(hEle HELE, pRect *RECT) {
	XEle_PointWndClientToEleClient.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:09:08
// @Function: XEleRectClientToWndClient
// @Description: 元素客户区坐标转换到窗口客户区坐标.
// @Calls: XEle_RectClientToWndClient
// @Input: hEle 元素句柄. pRect 坐标.
// @Return:
// *******************************************
func XEleRectClientToWndClient(hEle HELE, pRect *RECT) {
	XEle_RectClientToWndClient.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:10:34
// @Function: XElePointClientToWndClient
// @Description: 元素客户区坐标转换到窗口客户区坐标.
// @Calls: XEle_PointClientToWndClient
// @Input: hEle 元素句柄. pRect 坐标.
// @Return:
// *******************************************
func XElePointClientToWndClient(hEle HELE, pRect *RECT) {
	XEle_PointClientToWndClient.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:12:29
// @Function: XEleGetWndClientRect
// @Description: 元素基于窗口客户区坐标.
// @Calls: XEle_GetWndClientRect
// @Input: hEle 元素句柄. pRect 坐标.
// @Return:
// *******************************************
func XEleGetWndClientRect(hEle HELE, pRect *RECT) {
	XEle_GetWndClientRect.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:13:31
// @Function: XEleGetType
// @Description: 获取元素类型.
// @Calls: XEle_GetType
// @Input: hEle 元素句柄.
// @Return: 元素类型,参见宏定义 XC_OBJECT_TYPE.
// *******************************************
func XEleGetType(hEle HELE) XC_OBJECT_TYPE {
	ret, _, _ := XEle_GetType.Call(uintptr(hEle))

	return XC_OBJECT_TYPE(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:15:25
// @Function: XEleGetHWND
// @Description: 获取HWND句柄.
// @Calls: XEle_GetHWND
// @Input: hEle 元素句柄.
// @Return: HWND句柄.
// *******************************************
func XEleGetHWND(hEle HELE) HWND {
	ret, _, _ := XEle_GetHWND.Call(uintptr(hEle))

	return HWND(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:17:31
// @Function: XEleGetHWINDOW
// @Description: 获取HWINDOW句柄.
// @Calls: XEle_GetHWINDOW
// @Input: hEle 元素句柄.
// @Return: HWINDOW句柄.
// *******************************************
func XEleGetHWINDOW(hEle HELE) HWINDOW {
	ret, _, _ := XEle_GetHWINDOW.Call(uintptr(hEle))

	return HWINDOW(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:18:46
// @Function: XEleGetCursor
// @Description: 获取元素鼠标光标.
// @Calls: XEle_GetCursor
// @Input: hEle 元素句柄.
// @Return: 返回光标句柄.
// *******************************************
func XEleGetCursor(hEle HELE) HCURSOR {
	ret, _, _ := XEle_GetCursor.Call(uintptr(hEle))

	return HCURSOR(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:20:32
// @Function: XEleSetCursor
// @Description: 设置元素鼠标光标.
// @Calls: XEle_SetCursor
// @Input: hEle 元素句柄. hCursor 光标句柄.
// @Return:
// *******************************************
func XEleSetCursor(hEle HELE, hCursor HCURSOR) {
	XEle_SetCursor.Call(
		uintptr(hEle),
		uintptr(hCursor))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:22:06
// @Function: XEleAddEle
// @Description: 添加元素.添加一个子元素,到当前元素下.
// @Calls: XEle_AddEle
// @Input: hEle 元素句柄. hChildEle 要添加的子元素资源句柄.
// @Return: 如果成功返回TRUE,否则相反.
// *******************************************
func XEleAddEle(hEle, hChildEle HELE) bool {
	ret, _, _ := XEle_AddEle.Call(
		uintptr(hEle),
		uintptr(hChildEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:24:25
// @Function: XEleInsertEle
// @Description: 插入元素到目标元素前面.
// @Calls: XEle_InsertEle
// @Input: hEle 元素句柄. hChildEle 要插入的子元素句柄. hDestEle 目标元素.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XEleInsertEle(hEle, hChildEle, hDestEle HELE) bool {
	ret, _, _ := XEle_InsertEle.Call(
		uintptr(hEle),
		uintptr(hChildEle),
		uintptr(hDestEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:27:29
// @Function: XEleShowEle
// @Description: 显示隐藏元素.
// @Calls: XEle_ShowEle
// @Input: hEle 元素句柄. bShow TRUE显示元素,否则相反.
// @Return:
// *******************************************
func XEleShowEle(hEle HELE, bShow bool) {
	XEle_ShowEle.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bShow)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:30:14
// @Function: XEleSetRect
// @Description: 设置元素坐标.
// @Calls: XEle_SetRect
// @Input: hEle 元素句柄. pRect 坐标. bRedraw 是否重绘.
// @Return: 如果坐标未改变返回FALSE,否则相反.
// *******************************************
func XEleSetRect(hEle HELE, pRect *RECT, bRedraw BOOL) bool {
	ret, _, _ := XEle_SetRect.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(bRedraw))

	if ret == FALSE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:34:58
// @Function: XEleSetRectLogic
// @Description: 设置元素坐标,逻辑坐标,包含滚动视图偏移.
// @Calls: XEle_SetRectLogic
// @Input: hEle 元素句柄. pRect 坐标. bRedraw 是否重绘.
// @Return: 如果坐标未改变返回FALSE,否则相反.
// *******************************************
func XEleSetRectLogic(hEle HELE, pRect *RECT, bRedraw BOOL) bool {
	ret, _, _ := XEle_SetRectLogic.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(bRedraw))

	if ret == FALSE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:35:58
// @Function: XEleSetID
// @Description: 设置元素ID.
// @Calls: XEle_SetID
// @Input: hEle 元素句柄. id ID值.
// @Return:
// *******************************************
func XEleSetID(hEle HELE, id int) {
	XEle_SetID.Call(
		uintptr(hEle),
		uintptr(id))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:37:34
// @Function: XEleGetID
// @Description: 获取元素ID.
// @Calls: XEle_GetID
// @Input: hEle 元素句柄.
// @Return: 返回元素ID.
// *******************************************
func XEleGetID(hEle HELE) int {
	ret, _, _ := XEle_GetID.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:38:59
// @Function: XEleIsShow
// @Description: 元素是否显示.
// @Calls: XEle_IsShow
// @Input: hEle 元素句柄.
// @Return: 如果显示返回TRUE否则返回FALSE.
// *******************************************
func XEleIsShow(hEle HELE) bool {
	ret, _, _ := XEle_IsShow.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:40:08
// @Function: XEleIsDrawFocus
// @Description: 元素是否绘制焦点.
// @Calls: XEle_IsDrawFocus
// @Input: hEle 元素句柄.
// @Return: 如果绘制焦点返回TRUE否则返回FALSE.
// *******************************************
func XEleIsDrawFocus(hEle HELE) bool {
	ret, _, _ := XEle_IsDrawFocus.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:41:51
// @Function: XEleIsEnable
// @Description: 元素是否为启用状态.
// @Calls: XEle_IsEnable
// @Input: hEle 元素句柄.
// @Return: 如果启用状态返回TRUE否则返回FALSE.
// *******************************************
func XEleIsEnable(hEle HELE) bool {
	ret, _, _ := XEle_IsEnable.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:43:55
// @Function: XEleIsEnableFocus
// @Description: 元素是否启用焦点.
// @Calls: XEle_IsEnableFocus
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XEleIsEnableFocus(hEle HELE) bool {
	ret, _, _ := XEle_IsEnableFocus.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:45:50
// @Function: XEleIsMouseThrough
// @Description: 元素是否启用鼠标穿透.
// @Calls: XEle_IsMouseThrough
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XEleIsMouseThrough(hEle HELE) bool {
	ret, _, _ := XEle_IsMouseThrough.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

//
//
//
//
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
