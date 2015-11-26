package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xEle_Create                     *syscall.Proc
	xEle_RegEventC                  *syscall.Proc
	xEle_RegEventC1                 *syscall.Proc
	xEle_RegEventC2                 *syscall.Proc
	xEle_RemoveEventC               *syscall.Proc
	xEle_SendEvent                  *syscall.Proc
	xEle_PostEvent                  *syscall.Proc
	xEle_GetRect                    *syscall.Proc
	xEle_GetRectLogic               *syscall.Proc
	xEle_GetClientRect              *syscall.Proc
	xEle_GetWidth                   *syscall.Proc
	xEle_GetHeight                  *syscall.Proc
	xEle_RectWndClientToEleClient   *syscall.Proc
	xEle_PointWndClientToEleClient  *syscall.Proc
	xEle_RectClientToWndClient      *syscall.Proc
	xEle_PointClientToWndClient     *syscall.Proc
	xEle_GetWndClientRect           *syscall.Proc
	xEle_GetType                    *syscall.Proc
	xEle_GetHWND                    *syscall.Proc
	xEle_GetHWINDOW                 *syscall.Proc
	xEle_GetCursor                  *syscall.Proc
	xEle_SetCursor                  *syscall.Proc
	xEle_AddEle                     *syscall.Proc
	xEle_InsertEle                  *syscall.Proc
	xEle_ShowEle                    *syscall.Proc
	xEle_SetRect                    *syscall.Proc
	xEle_SetRectLogic               *syscall.Proc
	xEle_SetID                      *syscall.Proc
	xEle_GetID                      *syscall.Proc
	xEle_IsShow                     *syscall.Proc
	xEle_IsDrawFocus                *syscall.Proc
	xEle_IsEnable                   *syscall.Proc
	xEle_IsEnableFocus              *syscall.Proc
	xEle_IsMouseThrough             *syscall.Proc
	xEle_HitChildEle                *syscall.Proc
	xEle_IsBkTransparent            *syscall.Proc
	xEle_IsEnableEvent_XE_PAINT_END *syscall.Proc
	xEle_IsKeyTab                   *syscall.Proc
	xEle_IsSwitchFocus              *syscall.Proc
	xEle_IsEnable_XE_MOUSEWHEEL     *syscall.Proc
	xEle_Enable                     *syscall.Proc
	xEle_EnableFocus                *syscall.Proc
	xEle_EnableDrawFocus            *syscall.Proc
	xEle_EnableEvent_XE_PAINT_END   *syscall.Proc
	xEle_EnableBkTransparent        *syscall.Proc
	xEle_EnableMouseThrough         *syscall.Proc
	xEle_EnableKeyTab               *syscall.Proc
	xEle_EnableSwitchFocus          *syscall.Proc
	xEle_EnableEvent_XE_MOUSEWHEEL  *syscall.Proc
	xEle_GetParentEle               *syscall.Proc
	xEle_GetParent                  *syscall.Proc
	xEle_RemoveEle                  *syscall.Proc
	xEle_RedrawEle                  *syscall.Proc
	xEle_RedrawRect                 *syscall.Proc
	xEle_GetChildCount              *syscall.Proc
	xEle_GetChildByIndex            *syscall.Proc
	xEle_GetChildByID               *syscall.Proc
	xEle_GetChildShapeCount         *syscall.Proc
	xEle_GetChildShapeByIndex       *syscall.Proc
	xEle_SetTextColor               *syscall.Proc
	xEle_GetTextColor               *syscall.Proc
	xEle_SetFocusBorderColor        *syscall.Proc
	xEle_GetFocusBorderColor        *syscall.Proc
	xEle_SetFont                    *syscall.Proc
	xEle_GetFont                    *syscall.Proc
	xEle_SetAlpha                   *syscall.Proc
	xEle_Destroy                    *syscall.Proc
	xEle_AddBkBorder                *syscall.Proc
	xEle_AddBkFill                  *syscall.Proc
	xEle_AddBkImage                 *syscall.Proc
	xEle_GetBkInfoCount             *syscall.Proc
	xEle_ClearBkInfo                *syscall.Proc
	xEle_GetBkInfoManager           *syscall.Proc
	xEle_DrawFocus                  *syscall.Proc
	xEle_BindLayoutObject           *syscall.Proc
	xEle_GetLayoutObject            *syscall.Proc
	xEle_GetParentLayoutObject      *syscall.Proc
	xEle_SetUserData                *syscall.Proc
	xEle_GetUserData                *syscall.Proc
	xEle_GetContentSize             *syscall.Proc
	xEle_SetCapture                 *syscall.Proc
	xEle_SetLayoutWidth             *syscall.Proc
	xEle_SetLayoutHeight            *syscall.Proc
	xEle_GetLayoutWidth             *syscall.Proc
	xEle_GetLayoutHeight            *syscall.Proc
	xEle_EnableTransparentChannel   *syscall.Proc
	xEle_SetToolTip                 *syscall.Proc
	xEle_GetToolTip                 *syscall.Proc
	xEle_EnableToolTip              *syscall.Proc
	xEle_AdjustLayoutObject         *syscall.Proc
	xEle_AdjustLayout               *syscall.Proc
)

func init() {
	// Functions
	xEle_Create = XCDLL.MustFindProc("XEle_Create")
	xEle_RegEventC = XCDLL.MustFindProc("XEle_RegEventC")
	xEle_RegEventC1 = XCDLL.MustFindProc("XEle_RegEventC1")
	xEle_RegEventC2 = XCDLL.MustFindProc("XEle_RegEventC2")
	xEle_RemoveEventC = XCDLL.MustFindProc("XEle_RemoveEventC")
	xEle_SendEvent = XCDLL.MustFindProc("XEle_SendEvent")
	xEle_PostEvent = XCDLL.MustFindProc("XEle_PostEvent")
	xEle_GetRect = XCDLL.MustFindProc("XEle_GetRect")
	xEle_GetRectLogic = XCDLL.MustFindProc("XEle_GetRectLogic")
	xEle_GetClientRect = XCDLL.MustFindProc("XEle_GetClientRect")
	xEle_GetWidth = XCDLL.MustFindProc("XEle_GetWidth")
	xEle_GetHeight = XCDLL.MustFindProc("XEle_GetHeight")
	xEle_RectWndClientToEleClient = XCDLL.MustFindProc("XEle_RectWndClientToEleClient")
	xEle_PointWndClientToEleClient = XCDLL.MustFindProc("XEle_PointWndClientToEleClient")
	xEle_RectClientToWndClient = XCDLL.MustFindProc("XEle_RectClientToWndClient")
	xEle_PointClientToWndClient = XCDLL.MustFindProc("XEle_PointClientToWndClient")
	xEle_GetWndClientRect = XCDLL.MustFindProc("XEle_GetWndClientRect")
	xEle_GetType = XCDLL.MustFindProc("XEle_GetType")
	xEle_GetHWND = XCDLL.MustFindProc("XEle_GetHWND")
	xEle_GetHWINDOW = XCDLL.MustFindProc("XEle_GetHWINDOW")
	xEle_GetCursor = XCDLL.MustFindProc("XEle_GetCursor")
	xEle_SetCursor = XCDLL.MustFindProc("XEle_SetCursor")
	xEle_AddEle = XCDLL.MustFindProc("XEle_AddEle")
	xEle_InsertEle = XCDLL.MustFindProc("XEle_InsertEle")
	xEle_ShowEle = XCDLL.MustFindProc("XEle_ShowEle")
	xEle_SetRect = XCDLL.MustFindProc("XEle_SetRect")
	xEle_SetRectLogic = XCDLL.MustFindProc("XEle_SetRectLogic")
	xEle_SetID = XCDLL.MustFindProc("XEle_SetID")
	xEle_GetID = XCDLL.MustFindProc("XEle_GetID")
	xEle_IsShow = XCDLL.MustFindProc("XEle_IsShow")
	xEle_IsDrawFocus = XCDLL.MustFindProc("XEle_IsDrawFocus")
	xEle_IsEnable = XCDLL.MustFindProc("XEle_IsEnable")
	xEle_IsEnableFocus = XCDLL.MustFindProc("XEle_IsEnableFocus")
	xEle_IsMouseThrough = XCDLL.MustFindProc("XEle_IsMouseThrough")
	xEle_HitChildEle = XCDLL.MustFindProc("XEle_HitChildEle")
	xEle_IsBkTransparent = XCDLL.MustFindProc("XEle_IsBkTransparent")
	xEle_IsEnableEvent_XE_PAINT_END = XCDLL.MustFindProc("XEle_IsEnableEvent_XE_PAINT_END")
	xEle_IsKeyTab = XCDLL.MustFindProc("XEle_IsKeyTab")
	xEle_IsSwitchFocus = XCDLL.MustFindProc("XEle_IsSwitchFocus")
	xEle_IsEnable_XE_MOUSEWHEEL = XCDLL.MustFindProc("XEle_IsEnable_XE_MOUSEWHEEL")
	xEle_Enable = XCDLL.MustFindProc("XEle_Enable")
	xEle_EnableFocus = XCDLL.MustFindProc("XEle_EnableFocus")
	xEle_EnableDrawFocus = XCDLL.MustFindProc("XEle_EnableDrawFocus")
	xEle_EnableEvent_XE_PAINT_END = XCDLL.MustFindProc("XEle_EnableEvent_XE_PAINT_END")
	xEle_EnableBkTransparent = XCDLL.MustFindProc("XEle_EnableBkTransparent")
	xEle_EnableMouseThrough = XCDLL.MustFindProc("XEle_EnableMouseThrough")
	xEle_EnableKeyTab = XCDLL.MustFindProc("XEle_EnableKeyTab")
	xEle_EnableSwitchFocus = XCDLL.MustFindProc("XEle_EnableSwitchFocus")
	xEle_EnableEvent_XE_MOUSEWHEEL = XCDLL.MustFindProc("XEle_EnableEvent_XE_MOUSEWHEEL")
	xEle_GetParentEle = XCDLL.MustFindProc("XEle_GetParentEle")
	xEle_GetParent = XCDLL.MustFindProc("XEle_GetParent")
	xEle_RemoveEle = XCDLL.MustFindProc("XEle_RemoveEle")
	xEle_RedrawEle = XCDLL.MustFindProc("XEle_RedrawEle")
	xEle_RedrawRect = XCDLL.MustFindProc("XEle_RedrawRect")
	xEle_GetChildCount = XCDLL.MustFindProc("XEle_GetChildCount")
	xEle_GetChildByIndex = XCDLL.MustFindProc("XEle_GetChildByIndex")
	xEle_GetChildByID = XCDLL.MustFindProc("XEle_GetChildByID")
	xEle_GetChildShapeCount = XCDLL.MustFindProc("XEle_GetChildShapeCount")
	xEle_GetChildShapeByIndex = XCDLL.MustFindProc("XEle_GetChildShapeByIndex")
	xEle_SetTextColor = XCDLL.MustFindProc("XEle_SetTextColor")
	xEle_GetTextColor = XCDLL.MustFindProc("XEle_GetTextColor")
	xEle_SetFocusBorderColor = XCDLL.MustFindProc("XEle_SetFocusBorderColor")
	xEle_GetFocusBorderColor = XCDLL.MustFindProc("XEle_GetFocusBorderColor")
	xEle_SetFont = XCDLL.MustFindProc("XEle_SetFont")
	xEle_GetFont = XCDLL.MustFindProc("XEle_GetFont")
	xEle_SetAlpha = XCDLL.MustFindProc("XEle_SetAlpha")
	xEle_Destroy = XCDLL.MustFindProc("XEle_Destroy")
	xEle_AddBkBorder = XCDLL.MustFindProc("XEle_AddBkBorder")
	xEle_AddBkFill = XCDLL.MustFindProc("XEle_AddBkFill")
	xEle_AddBkImage = XCDLL.MustFindProc("XEle_AddBkImage")
	xEle_GetBkInfoCount = XCDLL.MustFindProc("XEle_GetBkInfoCount")
	xEle_ClearBkInfo = XCDLL.MustFindProc("XEle_ClearBkInfo")
	xEle_GetBkInfoManager = XCDLL.MustFindProc("XEle_GetBkInfoManager")
	xEle_DrawFocus = XCDLL.MustFindProc("XEle_DrawFocus")
	xEle_BindLayoutObject = XCDLL.MustFindProc("XEle_BindLayoutObject")
	xEle_GetLayoutObject = XCDLL.MustFindProc("XEle_GetLayoutObject")
	xEle_GetParentLayoutObject = XCDLL.MustFindProc("XEle_GetParentLayoutObject")
	xEle_SetUserData = XCDLL.MustFindProc("XEle_SetUserData")
	xEle_GetUserData = XCDLL.MustFindProc("XEle_GetUserData")
	xEle_GetContentSize = XCDLL.MustFindProc("XEle_GetContentSize")
	xEle_SetCapture = XCDLL.MustFindProc("XEle_SetCapture")
	xEle_SetLayoutWidth = XCDLL.MustFindProc("XEle_SetLayoutWidth")
	xEle_SetLayoutHeight = XCDLL.MustFindProc("XEle_SetLayoutHeight")
	xEle_GetLayoutWidth = XCDLL.MustFindProc("XEle_GetLayoutWidth")
	xEle_GetLayoutHeight = XCDLL.MustFindProc("XEle_GetLayoutHeight")
	xEle_EnableTransparentChannel = XCDLL.MustFindProc("XEle_EnableTransparentChannel")
	xEle_SetToolTip = XCDLL.MustFindProc("XEle_SetToolTip")
	xEle_GetToolTip = XCDLL.MustFindProc("XEle_GetToolTip")
	xEle_EnableToolTip = XCDLL.MustFindProc("XEle_EnableToolTip")
	xEle_AdjustLayoutObject = XCDLL.MustFindProc("XEle_AdjustLayoutObject")
	xEle_AdjustLayout = XCDLL.MustFindProc("XEle_AdjustLayout")
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 20:42:40
// @Function: XEleCreate
// @Description: 创建基础元素.
// @Calls: xEle_Create
// @Input: x 元素x坐标. y 元素y坐标. cx 宽度. cy 高度. hParent 父是窗口句柄或元素句柄.
// @Return: 元素句柄.
// *******************************************
func XEleCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xEle_Create.Call(
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
// @Calls: xEle_RegEventC
// @Input: hEle 元素句柄. nEvent 事件类型. pFun 事件函数指针.
// @Return:
// *******************************************
func XEleRegEventC(hEle HELE, nEvent int, pFun uintptr) {
	xEle_RegEventC.Call(
		uintptr(hEle),
		uintptr(nEvent),
		pFun)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 20:47:27
// @Function: XEleRegEventC1
// @Description: 注册事件C方式,省略1参数.
// @Calls: xEle_RegEventC1
// @Input: hEle 元素句柄. nEvent 事件类型. pFun 事件函数指针.
// @Return:
// *******************************************
func XEleRegEventC1(hEle HELE, nEvent int, pFun uintptr) {
	xEle_RegEventC1.Call(
		uintptr(hEle),
		uintptr(nEvent),
		pFun)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 20:49:26
// @Function: XEleRegEventC2
// @Description: 注册事件C方式,不省略参数.
// @Calls: xEle_RegEventC2
// @Input: hEle 元素句柄. nEvent 事件类型. pFun 事件函数指针.
// @Return:
// *******************************************
func XEleRegEventC2(hEle HELE, nEvent int, pFun uintptr) {
	xEle_RegEventC2.Call(
		uintptr(hEle),
		uintptr(nEvent),
		pFun)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 20:50:37
// @Function: XEleRemoveEventC
// @Description: 移除事件函数C方式.
// @Calls: xEle_RemoveEventC
// @Input: hEle 元素句柄. nEvent 事件类型. pFun 事件函数指针.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XEleRemoveEventC(hEle HELE, nEvent int, pFun uintptr) bool {
	ret, _, _ := xEle_RemoveEventC.Call(
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
// @Calls: xEle_SendEvent
// @Input: hEle 目标元素句柄. hEventEle 触发事件元素句柄. nEvent 事件类型. wParam 参数. lParam 参数.
// @Return: 事件返回值.
// *******************************************
func XEleSendEvent(hEle, hEventEle HELE, nEvent int, wParam, lParam uintptr) int {
	ret, _, _ := xEle_SendEvent.Call(
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
// @Calls: xEle_PostEvent
// @Input: hEle 目标元素句柄. hEventEle 触发事件元素句柄. nEvent 事件类型. wParam 参数. lParam 参数.
// @Return: 事件返回值.
// *******************************************
func XElePostEvent(hEle, hEventEle HELE, nEvent int, wParam, lParam uintptr) int {
	ret, _, _ := xEle_PostEvent.Call(
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
// @Calls: xEle_GetRect
// @Input: hEle 元素句柄. pRect 坐标.
// @Return:
// *******************************************
func XEleGetRect(hEle HELE, pRect *RECT) {
	xEle_GetRect.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 22:58:51
// @Function: XEleGetRectLogic
// @Description: 获取元素坐标,逻辑坐标,包含滚动视图偏移.
// @Calls: xEle_GetRectLogic
// @Input: hEle 元素句柄. pRect 坐标.
// @Return:
// *******************************************
func XEleGetRectLogic(hEle HELE, pRect *RECT) {
	xEle_GetRectLogic.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 22:59:59
// @Function: XEleGetClientRect
// @Description: 获取元素客户区坐标.
// @Calls: xEle_GetClientRect
// @Input: hEle 元素句柄. pRect 坐标.
// @Return:
// *******************************************
func XEleGetClientRect(hEle HELE, pRect *RECT) {
	xEle_GetClientRect.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:01:30
// @Function: XEleGetWidth
// @Description: 获取元素宽度.
// @Calls: xEle_GetWidth
// @Input: hEle 元素句柄.
// @Return: 宽度.
// *******************************************
func XEleGetWidth(hEle HELE) int {
	ret, _, _ := xEle_GetWidth.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:03:14
// @Function: XEleGetHeight
// @Description: 获取元素高度.
// @Calls: xEle_GetHeight
// @Input: hEle 元素句柄.
// @Return: 高度.
// *******************************************
func XEleGetHeight(hEle HELE) int {
	ret, _, _ := xEle_GetHeight.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:04:35
// @Function: XEleRectWndClientToEleClient
// @Description: 窗口客户区坐标转换到元素客户区坐标.
// @Calls: xEle_RectWndClientToEleClient
// @Input: hEle 元素句柄. pRect 坐标.
// @Return:
// *******************************************
func XEleRectWndClientToEleClient(hEle HELE, pRect *RECT) {
	xEle_RectWndClientToEleClient.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:07:01
// @Function: XElePointWndClientToEleClient
// @Description: 窗口客户区坐标转换到元素客户区坐标.
// @Calls: xEle_PointWndClientToEleClient
// @Input: hEle 元素句柄. pRect 坐标.
// @Return:
// *******************************************
func XElePointWndClientToEleClient(hEle HELE, pRect *RECT) {
	xEle_PointWndClientToEleClient.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:09:08
// @Function: XEleRectClientToWndClient
// @Description: 元素客户区坐标转换到窗口客户区坐标.
// @Calls: xEle_RectClientToWndClient
// @Input: hEle 元素句柄. pRect 坐标.
// @Return:
// *******************************************
func XEleRectClientToWndClient(hEle HELE, pRect *RECT) {
	xEle_RectClientToWndClient.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:10:34
// @Function: XElePointClientToWndClient
// @Description: 元素客户区坐标转换到窗口客户区坐标.
// @Calls: xEle_PointClientToWndClient
// @Input: hEle 元素句柄. pRect 坐标.
// @Return:
// *******************************************
func XElePointClientToWndClient(hEle HELE, pRect *RECT) {
	xEle_PointClientToWndClient.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:12:29
// @Function: XEleGetWndClientRect
// @Description: 元素基于窗口客户区坐标.
// @Calls: xEle_GetWndClientRect
// @Input: hEle 元素句柄. pRect 坐标.
// @Return:
// *******************************************
func XEleGetWndClientRect(hEle HELE, pRect *RECT) {
	xEle_GetWndClientRect.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:13:31
// @Function: XEleGetType
// @Description: 获取元素类型.
// @Calls: xEle_GetType
// @Input: hEle 元素句柄.
// @Return: 元素类型,参见宏定义 XC_OBJECT_TYPE.
// *******************************************
func XEleGetType(hEle HELE) XC_OBJECT_TYPE {
	ret, _, _ := xEle_GetType.Call(uintptr(hEle))

	return XC_OBJECT_TYPE(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:15:25
// @Function: XEleGetHWND
// @Description: 获取HWND句柄.
// @Calls: xEle_GetHWND
// @Input: hEle 元素句柄.
// @Return: HWND句柄.
// *******************************************
func XEleGetHWND(hEle HELE) HWND {
	ret, _, _ := xEle_GetHWND.Call(uintptr(hEle))

	return HWND(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:17:31
// @Function: XEleGetHWINDOW
// @Description: 获取HWINDOW句柄.
// @Calls: xEle_GetHWINDOW
// @Input: hEle 元素句柄.
// @Return: HWINDOW句柄.
// *******************************************
func XEleGetHWINDOW(hEle HELE) HWINDOW {
	ret, _, _ := xEle_GetHWINDOW.Call(uintptr(hEle))

	return HWINDOW(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:18:46
// @Function: XEleGetCursor
// @Description: 获取元素鼠标光标.
// @Calls: xEle_GetCursor
// @Input: hEle 元素句柄.
// @Return: 返回光标句柄.
// *******************************************
func XEleGetCursor(hEle HELE) HCURSOR {
	ret, _, _ := xEle_GetCursor.Call(uintptr(hEle))

	return HCURSOR(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:20:32
// @Function: XEleSetCursor
// @Description: 设置元素鼠标光标.
// @Calls: xEle_SetCursor
// @Input: hEle 元素句柄. hCursor 光标句柄.
// @Return:
// *******************************************
func XEleSetCursor(hEle HELE, hCursor HCURSOR) {
	xEle_SetCursor.Call(
		uintptr(hEle),
		uintptr(hCursor))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:22:06
// @Function: XEleAddEle
// @Description: 添加元素.添加一个子元素,到当前元素下.
// @Calls: xEle_AddEle
// @Input: hEle 元素句柄. hChildEle 要添加的子元素资源句柄.
// @Return: 如果成功返回TRUE,否则相反.
// *******************************************
func XEleAddEle(hEle, hChildEle HELE) bool {
	ret, _, _ := xEle_AddEle.Call(
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
// @Calls: xEle_InsertEle
// @Input: hEle 元素句柄. hChildEle 要插入的子元素句柄. hDestEle 目标元素.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XEleInsertEle(hEle, hChildEle, hDestEle HELE) bool {
	ret, _, _ := xEle_InsertEle.Call(
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
// @Calls: xEle_ShowEle
// @Input: hEle 元素句柄. bShow TRUE显示元素,否则相反.
// @Return:
// *******************************************
func XEleShowEle(hEle HELE, bShow bool) {
	xEle_ShowEle.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bShow)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:30:14
// @Function: XEleSetRect
// @Description: 设置元素坐标.
// @Calls: xEle_SetRect
// @Input: hEle 元素句柄. pRect 坐标. bRedraw 是否重绘.
// @Return: 如果坐标未改变返回FALSE,否则相反.
// *******************************************
func XEleSetRect(hEle HELE, pRect *RECT, bRedraw BOOL) bool {
	ret, _, _ := xEle_SetRect.Call(
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
// @Calls: xEle_SetRectLogic
// @Input: hEle 元素句柄. pRect 坐标. bRedraw 是否重绘.
// @Return: 如果坐标未改变返回FALSE,否则相反.
// *******************************************
func XEleSetRectLogic(hEle HELE, pRect *RECT, bRedraw BOOL) bool {
	ret, _, _ := xEle_SetRectLogic.Call(
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
// @Calls: xEle_SetID
// @Input: hEle 元素句柄. id ID值.
// @Return:
// *******************************************
func XEleSetID(hEle HELE, id int) {
	xEle_SetID.Call(
		uintptr(hEle),
		uintptr(id))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:37:34
// @Function: XEleGetID
// @Description: 获取元素ID.
// @Calls: xEle_GetID
// @Input: hEle 元素句柄.
// @Return: 返回元素ID.
// *******************************************
func XEleGetID(hEle HELE) int {
	ret, _, _ := xEle_GetID.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 23:38:59
// @Function: XEleIsShow
// @Description: 元素是否显示.
// @Calls: xEle_IsShow
// @Input: hEle 元素句柄.
// @Return: 如果显示返回TRUE否则返回FALSE.
// *******************************************
func XEleIsShow(hEle HELE) bool {
	ret, _, _ := xEle_IsShow.Call(uintptr(hEle))

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
// @Calls: xEle_IsDrawFocus
// @Input: hEle 元素句柄.
// @Return: 如果绘制焦点返回TRUE否则返回FALSE.
// *******************************************
func XEleIsDrawFocus(hEle HELE) bool {
	ret, _, _ := xEle_IsDrawFocus.Call(uintptr(hEle))

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
// @Calls: xEle_IsEnable
// @Input: hEle 元素句柄.
// @Return: 如果启用状态返回TRUE否则返回FALSE.
// *******************************************
func XEleIsEnable(hEle HELE) bool {
	ret, _, _ := xEle_IsEnable.Call(uintptr(hEle))

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
// @Calls: xEle_IsEnableFocus
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XEleIsEnableFocus(hEle HELE) bool {
	ret, _, _ := xEle_IsEnableFocus.Call(uintptr(hEle))

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
// @Calls: xEle_IsMouseThrough
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XEleIsMouseThrough(hEle HELE) bool {
	ret, _, _ := xEle_IsMouseThrough.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 19:49:00
// @Function: XEleHitChildEle
// @Description: 检测坐标点所在元素,包含子元素的子元素.
// @Calls: xEle_HitChildEle
// @Input: hEle 元素句柄. pPt 坐标点.
// @Return: 成功返回元素句柄,否则返回NULL.
// *******************************************
func XEleHitChildEle(hEle HELE, pPt *POINT) HELE {
	ret, _, _ := xEle_HitChildEle.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)))

	return HELE(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:15:54
// @Function: XEleIsBkTransparent
// @Description: 是否背景透明.
// @Calls: xEle_IsBkTransparent
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XEleIsBkTransparent(hEle HELE) bool {
	ret, _, _ := xEle_IsBkTransparent.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:18:23
// @Function: XEleIsEnableEvent_XE_PAINT_END
// @Description: 是否启XE_PAINT_END用事件.
// @Calls: xEle_IsEnableEvent_XE_PAINT_END
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XEleIsEnableEvent_XE_PAINT_END(hEle HELE) bool {
	ret, _, _ := xEle_IsEnableEvent_XE_PAINT_END.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:26:14
// @Function: XEleIsKeyTab
// @Description: 是否接收Tab键.
// @Calls: xEle_IsKeyTab
// @Input: hEle 元素句柄.
// @Return: 是返回TRUE否则返回FALSE.
// *******************************************
func XEleIsKeyTab(hEle HELE) bool {
	ret, _, _ := xEle_IsKeyTab.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:27:41
// @Function: XEleIsSwitchFocus
// @Description: 是否接受通过键盘切换焦点.
// @Calls: xEle_IsSwitchFocus
// @Input: hEle 元素句柄.
// @Return: 是返回TRUE否则返回FALSE.
// *******************************************
func XEleIsSwitchFocus(hEle HELE) bool {
	ret, _, _ := xEle_IsSwitchFocus.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:29:31
// @Function: XEleIsEnable_XE_MOUSEWHEEL
// @Description: 判断是否启用鼠标滚动事件,如果禁用那么事件会发送给他的父元素.
// @Calls: xEle_IsEnable_XE_MOUSEWHEEL
// @Input: hEle 元素句柄.
// @Return: 是返回TRUE否则返回FALSE.
// *******************************************
func XEleIsEnable_XE_MOUSEWHEEL(hEle HELE) bool {
	ret, _, _ := xEle_IsEnable_XE_MOUSEWHEEL.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:30:55
// @Function: XEleEnable
// @Description: 启用或禁用元素.
// @Calls: xEle_Enable
// @Input: hEle 元素句柄. bEnable 启用或禁用.
// @Return:
// *******************************************
func XEleEnable(hEle HELE, bEnable bool) {
	xEle_Enable.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:32:34
// @Function: XEleEnableFocus
// @Description: 启用焦点.
// @Calls: xEle_EnableFocus
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableFocus(hEle HELE, bEnable bool) {
	xEle_EnableFocus.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:44:01
// @Function: XEleEnableDrawFocus
// @Description: 启用绘制焦点.
// @Calls: xEle_EnableDrawFocus
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableDrawFocus(hEle HELE, bEnable bool) {
	xEle_EnableDrawFocus.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:45:08
// @Function: XEleEnableEvent_XE_PAINT_END
// @Description: 启用XE_PAINT_END事件.
// @Calls: xEle_EnableEvent_XE_PAINT_END
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableEvent_XE_PAINT_END(hEle HELE, bEnable bool) {
	xEle_EnableEvent_XE_PAINT_END.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:46:18
// @Function: XEleEnableBkTransparent
// @Description: 启用背景透明.
// @Calls: xEle_EnableBkTransparent
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableBkTransparent(hEle HELE, bEnable bool) {
	xEle_EnableBkTransparent.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:47:36
// @Function: XEleEnableMouseThrough
// @Description: 启用鼠标穿透.
// @Calls: xEle_EnableMouseThrough
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableMouseThrough(hEle HELE, bEnable bool) {
	xEle_EnableMouseThrough.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:48:12
// @Function: XEleEnableKeyTab
// @Description: 启用接收Tab键.
// @Calls: xEle_EnableKeyTab
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableKeyTab(hEle HELE, bEnable bool) {
	xEle_EnableKeyTab.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:04:40
// @Function: XEleEnableSwitchFocus
// @Description: 启用接受通过键盘切换焦点.
// @Calls: xEle_EnableSwitchFocus
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableSwitchFocus(hEle HELE, bEnable bool) {
	xEle_EnableSwitchFocus.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:08:55
// @Function: XEleEnableEvent_XE_MOUSEWHEEL
// @Description: 启用接收鼠标滚动事件,如果禁用那么事件会传递给父元素.
// @Calls: xEle_EnableEvent_XE_MOUSEWHEEL
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableEvent_XE_MOUSEWHEEL(hEle HELE, bEnable bool) {
	xEle_EnableEvent_XE_MOUSEWHEEL.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:10:45
// @Function: XEleGetParentEle
// @Description: 获取父元素.
// @Calls: xEle_GetParentEle
// @Input: hEle 元素句柄.
// @Return: 元素句柄.
// *******************************************
func XEleGetParentEle(hEle HELE) HELE {
	ret, _, _ := xEle_GetParentEle.Call(uintptr(hEle))

	return HELE(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:12:20
// @Function: XEleGetParent
// @Description: 获取父对象,父可能是元素或窗口,通过此函数可以检查是否有父.
// @Calls: xEle_GetParent
// @Input: hEle 元素句柄.
// @Return: 对象句柄.
// *******************************************
func XEleGetParent(hEle HELE) HXCGUI {
	ret, _, _ := xEle_GetParent.Call(uintptr(hEle))

	return HXCGUI(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:14:28
// @Function: XEleRemoveEle
// @Description: 移除元素,但不销毁.
// @Calls: xEle_RemoveEle
// @Input: hEle 元素句柄.
// @Return:
// *******************************************
func XEleRemoveEle(hEle HELE) {
	xEle_RemoveEle.Call(uintptr(hEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:15:26
// @Function: XEleRedrawEle
// @Description: 重绘元素.
// @Calls: xEle_RedrawEle
// @Input: hEle 元素句柄.
// @Return:
// *******************************************
func XEleRedrawEle(hEle HELE) {
	xEle_RedrawEle.Call(uintptr(hEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:16:20
// @Function: XEleRedrawRect
// @Description: 重绘元素指定区域.
// @Calls: xEle_RedrawRect
// @Input: hEle 元素句柄. pRect 相对于元素客户区坐标.
// @Return:
// *******************************************
func XEleRedrawRect(hEle HELE, pRect *RECT) {
	xEle_RedrawRect.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:17:59
// @Function: XEleGetChildCount
// @Description: 获取子元素数量,只检测当前层元素.
// @Calls: xEle_GetChildCount
// @Input: hEle 元素句柄.
// @Return: 子元素数量.
// *******************************************
func XEleGetChildCount(hEle HELE) int {
	ret, _, _ := xEle_GetChildCount.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:20:01
// @Function: XEleGetChildByIndex
// @Description: 获取子元素通过索引,只检测当前层元素.
// @Calls: xEle_GetChildByIndex
// @Input: hEle 元素句柄. index 索引.
// @Return: 元素句柄.
// *******************************************
func XEleGetChildByIndex(hEle HELE, index int) HELE {
	ret, _, _ := xEle_GetChildByIndex.Call(
		uintptr(hEle),
		uintptr(index))

	return HELE(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:22:06
// @Function: XEleGetChildByID
// @Description: 获取子元素通过ID,只检测当前层元素.
// @Calls: xEle_GetChildByID
// @Input: hEle 元素句柄. id 元素ID.
// @Return: 元素句柄.
// *******************************************
func XEleGetChildByID(hEle HELE, id int) HELE {
	ret, _, _ := xEle_GetChildByID.Call(
		uintptr(hEle),
		uintptr(id))

	return HELE(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:23:41
// @Function: XEleGetChildShapeCount
// @Description: 获取子对象(形状对象)数量,只检查当前层.
// @Calls: xEle_GetChildShapeCount
// @Input: hEle 元素句柄.
// @Return: 形状对象数量.
// *******************************************
func XEleGetChildShapeCount(hEle HELE) int {
	ret, _, _ := xEle_GetChildShapeCount.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:25:45
// @Function: XEleGetChildShapeByIndex
// @Description: 通过索引返回形状对象句柄.
// @Calls: xEle_GetChildShapeByIndex
// @Input: hEle 元素句柄. index 索引.
// @Return: 返回形状对象句柄.
// *******************************************
func XEleGetChildShapeByIndex(hEle HELE, index int) HXCGUI {
	ret, _, _ := xEle_GetChildShapeByIndex.Call(
		uintptr(hEle),
		uintptr(index))

	return HXCGUI(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:27:51
// @Function: XEleSetTextColor
// @Description: 设置文本颜色.
// @Calls: xEle_SetTextColor
// @Input: hEle 元素句柄. color RGB颜色值. alpha 透明度.
// @Return:
// *******************************************
func XEleSetTextColor(hEle HELE, color COLORREF, alpha byte) {
	xEle_SetTextColor.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:29:41
// @Function: XEleGetTextColor
// @Description: 获取文本颜色.
// @Calls: xEle_GetTextColor
// @Input: hEle 元素句柄.
// @Return:  文本颜色值.
// *******************************************
func XEleGetTextColor(hEle HELE) COLORREF {
	ret, _, _ := xEle_GetTextColor.Call(uintptr(hEle))

	return COLORREF(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:31:42
// @Function: XEleSetFocusBorderColor
// @Description: 设置焦点边框颜色.
// @Calls: xEle_SetFocusBorderColor
// @Input: hEle 元素句柄. color RGB颜色值. alpha 透明度.
// @Return:
// *******************************************
func XEleSetFocusBorderColor(hEle HELE, color COLORREF, alpha byte) {
	xEle_SetFocusBorderColor.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:33:27
// @Function: XEleGetFocusBorderColor
// @Description: 获取焦点边框颜色.
// @Calls: xEle_GetFocusBorderColor
// @Input: hEle 元素句柄.
// @Return:  返回颜色值,最高位字节是透明度.
// *******************************************
func XEleGetFocusBorderColor(hEle HELE) COLORREF {
	ret, _, _ := xEle_GetFocusBorderColor.Call(uintptr(hEle))

	return COLORREF(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:34:53
// @Function: XEleSetFont
// @Description: 设置元素字体.
// @Calls: xEle_SetFont
// @Input: hEle 元素句柄. hFontx 炫彩字体.
// @Return:
// *******************************************
func XEleSetFont(hEle HELE, hFontx HFONTX) {
	xEle_SetFont.Call(
		uintptr(hEle),
		uintptr(hFontx))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:36:58
// @Function: XEleGetFont
// @Description: 获取元素字体.
// @Calls: xEle_GetFont
// @Input: hEle 元素句柄.
// @Return: 返回炫彩字体句柄.
// *******************************************
func XEleGetFont(hEle HELE) HFONTX {
	ret, _, _ := xEle_GetFont.Call(uintptr(hEle))

	return HFONTX(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:38:15
// @Function: XEleSetAlpha
// @Description: 设置元素透明度.
// @Calls: xEle_SetAlpha
// @Input: hEle 元素句柄. alpha 透明度.
// @Return:
// *******************************************
func XEleSetAlpha(hEle HELE, alpha byte) {
	xEle_SetAlpha.Call(
		uintptr(hEle),
		uintptr(alpha))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:39:30
// @Function: XEleDestroy
// @Description: 销毁元素.
// @Calls: xEle_Destroy
// @Input: hEle 元素句柄.
// @Return:
// *******************************************
func XEleDestroy(hEle HELE) {
	xEle_Destroy.Call(uintptr(hEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:40:47
// @Function: XEleAddBkBorder
// @Description: 添加背景内容边框.
// @Calls: xEle_AddBkBorder
// @Input: hEle 元素句柄. color RGB颜色. alpha 透明度. width 线宽.
// @Return:
// *******************************************
func XEleAddBkBorder(hEle HELE, color COLORREF, alpha byte, width int) {
	xEle_AddBkBorder.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha),
		uintptr(width))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:42:36
// @Function: XEleAddBkFill
// @Description: 添加背景内容填充.
// @Calls: xEle_AddBkFill
// @Input: hEle 元素句柄. color RGB颜色. alpha 透明度.
// @Return:
// *******************************************
func XEleAddBkFill(hEle HELE, color COLORREF, alpha byte) {
	xEle_AddBkFill.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:43:45
// @Function: XEleAddBkImage
// @Description: 添加背景内容图片.
// @Calls: xEle_AddBkImage
// @Input: hEle 元素句柄. hImage 图片句柄.
// @Return:
// *******************************************
func XEleAddBkImage(hEle HELE, hImage HIMAGE) {
	xEle_AddBkImage.Call(
		uintptr(hEle),
		uintptr(hImage))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:46:00
// @Function: XEleGetBkInfoCount
// @Description: 获取背景内容数量.
// @Calls: xEle_GetBkInfoCount
// @Input: hEle 元素句柄.
// @Return: 返回背景内容数量.
// *******************************************
func XEleGetBkInfoCount(hEle HELE) int {
	ret, _, _ := xEle_GetBkInfoCount.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:47:25
// @Function: XEleClearBkInfo
// @Description: 清空背景内容; 如果背景没有内容,将使用系统默认内容,以便保证背景正确.
// @Calls: xEle_ClearBkInfo
// @Input: hEle 元素句柄.
// @Return:
// *******************************************
func XEleClearBkInfo(hEle HELE) {
	xEle_ClearBkInfo.Call(uintptr(hEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:48:33
// @Function: XEleGetBkInfoManager
// @Description: 获取元素背景内容管理器.
// @Calls: xEle_GetBkInfoManager
// @Input: hEle 元素句柄.
// @Return: 背景内容管理器.
// *******************************************
func XEleGetBkInfoManager(hEle HELE) HBKINFOM {
	ret, _, _ := xEle_GetBkInfoManager.Call(uintptr(hEle))

	return HBKINFOM(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:48:33
// @Function: XEleDrawFocus
// @Description: 绘制元素焦点.
// @Calls: xEle_DrawFocus
// @Input: hEle 元素句柄. hDraw 图形绘制句柄. pRect 区域坐标.
// @Return: 绘制成功返回TRUE,如果不需要绘制焦点返回FALSE.
// *******************************************
func XEleDrawFocus(hEle HELE, hDraw HDRAW, pRect *RECT) bool {
	ret, _, _ := xEle_DrawFocus.Call(
		uintptr(hEle),
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pRect)))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 22:25:57
// @Function: XEleBindLayoutObject
// @Description: 绑定布局对象,只可以绑定一个.
// @Calls: xEle_BindLayoutObject
// @Input: hEle 元素句柄. hLayout 布局对象句柄.
// @Return:
// *******************************************
func XEleBindLayoutObject(hEle HELE, hLayout HXCGUI) {
	xEle_BindLayoutObject.Call(
		uintptr(hEle),
		uintptr(hLayout))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:13:10
// @Function: XEleGetLayoutObject
// @Description: 获取绑定的布局对象.
// @Calls: xEle_GetLayoutObject
// @Input: hEle 元素句柄.
// @Return: 布局对象句柄.
// *******************************************
func XEleGetLayoutObject(hEle HELE) HXCGUI {
	ret, _, _ := xEle_GetLayoutObject.Call(uintptr(hEle))

	return HXCGUI(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:15:21
// @Function: XEleGetParentLayoutObject
// @Description: 获取父布局对象.
// @Calls: xEle_GetParentLayoutObject
// @Input: hEle 元素句柄.
// @Return: 布局对象句柄.
// *******************************************
func XEleGetParentLayoutObject(hEle HELE) HXCGUI {
	ret, _, _ := xEle_GetParentLayoutObject.Call(uintptr(hEle))

	return HXCGUI(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:16:37
// @Function: XEleSetUserData
// @Description: 设置用户数据.
// @Calls: xEle_SetUserData
// @Input: hEle 元素句柄. nData 用户数据.
// @Return:
// *******************************************
func XEleSetUserData(hEle HELE, nData int) {
	xEle_SetUserData.Call(
		uintptr(hEle),
		uintptr(nData))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:18:26
// @Function: XEleGetUserData
// @Description: 获取用户数据.
// @Calls: xEle_GetUserData
// @Input: hEle 元素句柄.
// @Return: 用户数据.
// *******************************************
func XEleGetUserData(hEle HELE) int {
	ret, _, _ := xEle_GetUserData.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 17:59:36
// @Function: XEleGetContentSize
// @Description: 获取内容大小.
// @Calls: xEle_GetContentSize
// @Input: hEle 元素句柄. pSize 大小.
// @Return: 元素内容大小.
// *******************************************
func XEleGetContentSize(hEle HELE, pSize *uint16) {
	xEle_GetContentSize.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pSize)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:27:57
// @Function: XEleSetCapture
// @Description: 设置鼠标捕获.
// @Calls: xEle_SetCapture
// @Input: hEle 元素句柄. b TRUE设置,FALSE取消.
// @Return:
// *******************************************
func XEleSetCapture(hEle HELE, b bool) {
	xEle_SetCapture.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(b)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:32:55
// @Function: XEleSetLayoutWidth
// @Description: 设置布局宽度.
// @Calls: xEle_SetLayoutWidth
// @Input: hEle 元素句柄. nType 属性标识. nWidth 宽度.
// @Return:
// *******************************************
func XEleSetLayoutWidth(hEle HELE, nType LAYOUT_SIZE_TYPE_, nWidth int) {
	xEle_SetLayoutWidth.Call(
		uintptr(hEle),
		uintptr(nType),
		uintptr(nWidth))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:33:19
// @Function: XEleSetLayoutHeight
// @Description: 设备布局高度.
// @Calls: xEle_SetLayoutHeight
// @Input: hEle 元素句柄. nType 属性标识. nHeight 高度.
// @Return:
// *******************************************
func XEleSetLayoutHeight(hEle HELE, nType LAYOUT_SIZE_TYPE_, nHeight int) {
	xEle_SetLayoutHeight.Call(
		uintptr(hEle),
		uintptr(nType),
		uintptr(nHeight))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:34:23
// @Function: XEleGetLayoutWidth
// @Description: 获取布局宽度.
// @Calls: xEle_GetLayoutWidth
// @Input: hEle 元素句柄. nType 属性标识. pWidth 宽度.
// @Return:
// *******************************************
func XEleGetLayoutWidth(hEle HELE, nType *LAYOUT_SIZE_TYPE_, nWidth *uint16) {
	xEle_GetLayoutWidth.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(nType)),
		uintptr(unsafe.Pointer(nWidth)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:37:16
// @Function: XEleGetLayoutHeight
// @Description: 获取布局宽度.
// @Calls: xEle_GetLayoutHeight
// @Input: hEle 元素句柄. nType 属性标识. nHeight 高度.
// @Return:
// *******************************************
func XEleGetLayoutHeight(hEle HELE, nType *LAYOUT_SIZE_TYPE_, nHeight *uint16) {
	xEle_GetLayoutHeight.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(nType)),
		uintptr(unsafe.Pointer(nHeight)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:38:00
// @Function: XEleEnableTransparentChannel
// @Description: 启用或关闭元素透明通道,如果启用,将强制设置元素背景不透明,默认为启用,此功能是为了兼容GDI不支持透明通道问题.
// @Calls: xEle_EnableTransparentChannel
// @Input: hEle 元素句柄. bEnable 启用或关闭.
// @Return:
// *******************************************
func XEleEnableTransparentChannel(hEle HELE, bEnable bool) {
	xEle_EnableTransparentChannel.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:40:01
// @Function: XEleSetToolTip
// @Description: 设置工具提示内容.
// @Calls: xEle_SetToolTip
// @Input: hEle 元素句柄. pText 工具提示内容.
// @Return:
// *******************************************
func XEleSetToolTip(hEle HELE, pText string) {
	xEle_SetToolTip.Call(
		uintptr(hEle),
		StringToUintPtr(pText))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:41:20
// @Function: XEleGetToolTip
// @Description: 获取工具提示内容.
// @Calls: xEle_GetToolTip
// @Input: hEle 元素句柄. pOut 接收内容缓冲区. nOutLen 缓冲区大小,字符为单位.
// @Return:
// *******************************************
func XEleGetToolTip(hEle HELE, pOut *uint16, nOutLen int) {
	xEle_GetToolTip.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:43:33
// @Function: XEleEnableToolTip
// @Description: 启用工具提示.
// @Calls: xEle_EnableToolTip
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableToolTip(hEle HELE, bEnable bool) {
	xEle_EnableToolTip.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:45:01
// @Function: XEleAdjustLayoutObject
// @Description: 调整布局对象.
// @Calls: xEle_AdjustLayoutObject
// @Input: hEle 元素句柄.
// @Return:
// *******************************************
func XEleAdjustLayoutObject(hEle HELE) {
	xEle_AdjustLayoutObject.Call(uintptr(hEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:46:13
// @Function: XEleAdjustLayout
// @Description: 调整布局.
// @Calls: xEle_AdjustLayout
// @Input: hEle 元素句柄.
// @Return:
// *******************************************
func XEleAdjustLayout(hEle HELE) {
	xEle_AdjustLayout.Call(uintptr(hEle))
}
