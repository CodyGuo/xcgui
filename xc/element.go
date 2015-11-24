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

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 19:49:00
// @Function: XEleHitChildEle
// @Description: 检测坐标点所在元素,包含子元素的子元素.
// @Calls: XEle_HitChildEle
// @Input: hEle 元素句柄. pPt 坐标点.
// @Return: 成功返回元素句柄,否则返回NULL.
// *******************************************
func XEleHitChildEle(hEle HELE, pPt *POINT) HELE {
	ret, _, _ := XEle_HitChildEle.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)))

	return HELE(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:15:54
// @Function: XEleIsBkTransparent
// @Description: 是否背景透明.
// @Calls: XEle_IsBkTransparent
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XEleIsBkTransparent(hEle HELE) bool {
	ret, _, _ := XEle_IsBkTransparent.Call(uintptr(hEle))

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
// @Calls: XEle_IsEnableEvent_XE_PAINT_END
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XEleIsEnableEvent_XE_PAINT_END(hEle HELE) bool {
	ret, _, _ := XEle_IsEnableEvent_XE_PAINT_END.Call(uintptr(hEle))

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
// @Calls: XEle_IsKeyTab
// @Input: hEle 元素句柄.
// @Return: 是返回TRUE否则返回FALSE.
// *******************************************
func XEleIsKeyTab(hEle HELE) bool {
	ret, _, _ := XEle_IsKeyTab.Call(uintptr(hEle))

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
// @Calls: XEle_IsSwitchFocus
// @Input: hEle 元素句柄.
// @Return: 是返回TRUE否则返回FALSE.
// *******************************************
func XEleIsSwitchFocus(hEle HELE) bool {
	ret, _, _ := XEle_IsSwitchFocus.Call(uintptr(hEle))

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
// @Calls: XEle_IsEnable_XE_MOUSEWHEEL
// @Input: hEle 元素句柄.
// @Return: 是返回TRUE否则返回FALSE.
// *******************************************
func XEleIsEnable_XE_MOUSEWHEEL(hEle HELE) bool {
	ret, _, _ := XEle_IsEnable_XE_MOUSEWHEEL.Call(uintptr(hEle))

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
// @Calls: XEle_Enable
// @Input: hEle 元素句柄. bEnable 启用或禁用.
// @Return:
// *******************************************
func XEleEnable(hEle HELE, bEnable bool) {
	XEle_Enable.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:32:34
// @Function: XEleEnableFocus
// @Description: 启用焦点.
// @Calls: XEle_EnableFocus
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableFocus(hEle HELE, bEnable bool) {
	XEle_EnableFocus.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:44:01
// @Function: XEleEnableDrawFocus
// @Description: 启用绘制焦点.
// @Calls: XEle_EnableDrawFocus
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableDrawFocus(hEle HELE, bEnable bool) {
	XEle_EnableDrawFocus.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:45:08
// @Function: XEleEnableEvent_XE_PAINT_END
// @Description: 启用XE_PAINT_END事件.
// @Calls: XEle_EnableEvent_XE_PAINT_END
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableEvent_XE_PAINT_END(hEle HELE, bEnable bool) {
	XEle_EnableEvent_XE_PAINT_END.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:46:18
// @Function: XEleEnableBkTransparent
// @Description: 启用背景透明.
// @Calls: XEle_EnableBkTransparent
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableBkTransparent(hEle HELE, bEnable bool) {
	XEle_EnableBkTransparent.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:47:36
// @Function: XEleEnableMouseThrough
// @Description: 启用鼠标穿透.
// @Calls: XEle_EnableMouseThrough
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableMouseThrough(hEle HELE, bEnable bool) {
	XEle_EnableMouseThrough.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 20:48:12
// @Function: XEleEnableKeyTab
// @Description: 启用接收Tab键.
// @Calls: XEle_EnableKeyTab
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableKeyTab(hEle HELE, bEnable bool) {
	XEle_EnableKeyTab.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:04:40
// @Function: XEleEnableSwitchFocus
// @Description: 启用接受通过键盘切换焦点.
// @Calls: XEle_EnableSwitchFocus
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableSwitchFocus(hEle HELE, bEnable bool) {
	XEle_EnableSwitchFocus.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:08:55
// @Function: XEleEnableEvent_XE_MOUSEWHEEL
// @Description: 启用接收鼠标滚动事件,如果禁用那么事件会传递给父元素.
// @Calls: XEle_EnableEvent_XE_MOUSEWHEEL
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableEvent_XE_MOUSEWHEEL(hEle HELE, bEnable bool) {
	XEle_EnableEvent_XE_MOUSEWHEEL.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:10:45
// @Function: XEleGetParentEle
// @Description: 获取父元素.
// @Calls: XEle_GetParentEle
// @Input: hEle 元素句柄.
// @Return: 元素句柄.
// *******************************************
func XEleGetParentEle(hEle HELE) HELE {
	ret, _, _ := XEle_GetParentEle.Call(uintptr(hEle))

	return HELE(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:12:20
// @Function: XEleGetParent
// @Description: 获取父对象,父可能是元素或窗口,通过此函数可以检查是否有父.
// @Calls: XEle_GetParent
// @Input: hEle 元素句柄.
// @Return: 对象句柄.
// *******************************************
func XEleGetParent(hEle HELE) HXCGUI {
	ret, _, _ := XEle_GetParent.Call(uintptr(hEle))

	return HXCGUI(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:14:28
// @Function: XEleRemoveEle
// @Description: 移除元素,但不销毁.
// @Calls: XEle_RemoveEle
// @Input: hEle 元素句柄.
// @Return:
// *******************************************
func XEleRemoveEle(hEle HELE) {
	XEle_RemoveEle.Call(uintptr(hEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:15:26
// @Function: XEleRedrawEle
// @Description: 重绘元素.
// @Calls: XEle_RedrawEle
// @Input: hEle 元素句柄.
// @Return:
// *******************************************
func XEleRedrawEle(hEle HELE) {
	XEle_RedrawEle.Call(uintptr(hEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:16:20
// @Function: XEleRedrawRect
// @Description: 重绘元素指定区域.
// @Calls: XEle_RedrawRect
// @Input: hEle 元素句柄. pRect 相对于元素客户区坐标.
// @Return:
// *******************************************
func XEleRedrawRect(hEle HELE, pRect *RECT) {
	XEle_RedrawRect.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:17:59
// @Function: XEleGetChildCount
// @Description: 获取子元素数量,只检测当前层元素.
// @Calls: XEle_GetChildCount
// @Input: hEle 元素句柄.
// @Return: 子元素数量.
// *******************************************
func XEleGetChildCount(hEle HELE) int {
	ret, _, _ := XEle_GetChildCount.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:20:01
// @Function: XEleGetChildByIndex
// @Description: 获取子元素通过索引,只检测当前层元素.
// @Calls: XEle_GetChildByIndex
// @Input: hEle 元素句柄. index 索引.
// @Return: 元素句柄.
// *******************************************
func XEleGetChildByIndex(hEle HELE, index int) HELE {
	ret, _, _ := XEle_GetChildByIndex.Call(
		uintptr(hEle),
		uintptr(index))

	return HELE(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:22:06
// @Function: XEleGetChildByID
// @Description: 获取子元素通过ID,只检测当前层元素.
// @Calls: XEle_GetChildByID
// @Input: hEle 元素句柄. id 元素ID.
// @Return: 元素句柄.
// *******************************************
func XEleGetChildByID(hEle HELE, id int) HELE {
	ret, _, _ := XEle_GetChildByID.Call(
		uintptr(hEle),
		uintptr(id))

	return HELE(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:23:41
// @Function: XEleGetChildShapeCount
// @Description: 获取子对象(形状对象)数量,只检查当前层.
// @Calls: XEle_GetChildShapeCount
// @Input: hEle 元素句柄.
// @Return: 形状对象数量.
// *******************************************
func XEleGetChildShapeCount(hEle HELE) int {
	ret, _, _ := XEle_GetChildShapeCount.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:25:45
// @Function: XEleGetChildShapeByIndex
// @Description: 通过索引返回形状对象句柄.
// @Calls: XEle_GetChildShapeByIndex
// @Input: hEle 元素句柄. index 索引.
// @Return: 返回形状对象句柄.
// *******************************************
func XEleGetChildShapeByIndex(hEle HELE, index int) HXCGUI {
	ret, _, _ := XEle_GetChildShapeByIndex.Call(
		uintptr(hEle),
		uintptr(index))

	return HXCGUI(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:27:51
// @Function: XEleSetTextColor
// @Description: 设置文本颜色.
// @Calls: XEle_SetTextColor
// @Input: hEle 元素句柄. color RGB颜色值. alpha 透明度.
// @Return:
// *******************************************
func XEleSetTextColor(hEle HELE, color COLORREF, alpha byte) {
	XEle_SetTextColor.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:29:41
// @Function: XEleGetTextColor
// @Description: 获取文本颜色.
// @Calls: XEle_GetTextColor
// @Input: hEle 元素句柄.
// @Return:  文本颜色值.
// *******************************************
func XEleGetTextColor(hEle HELE) COLORREF {
	ret, _, _ := XEle_GetTextColor.Call(uintptr(hEle))

	return COLORREF(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:31:42
// @Function: XEleSetFocusBorderColor
// @Description: 设置焦点边框颜色.
// @Calls: XEle_SetFocusBorderColor
// @Input: hEle 元素句柄. color RGB颜色值. alpha 透明度.
// @Return:
// *******************************************
func XEleSetFocusBorderColor(hEle HELE, color COLORREF, alpha byte) {
	XEle_SetFocusBorderColor.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:33:27
// @Function: XEleGetFocusBorderColor
// @Description: 获取焦点边框颜色.
// @Calls: XEle_GetFocusBorderColor
// @Input: hEle 元素句柄.
// @Return:  返回颜色值,最高位字节是透明度.
// *******************************************
func XEleGetFocusBorderColor(hEle HELE) COLORREF {
	ret, _, _ := XEle_GetFocusBorderColor.Call(uintptr(hEle))

	return COLORREF(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:34:53
// @Function: XEleSetFont
// @Description: 设置元素字体.
// @Calls: XEle_SetFont
// @Input: hEle 元素句柄. hFontx 炫彩字体.
// @Return:
// *******************************************
func XEleSetFont(hEle HELE, hFontx HFONTX) {
	XEle_SetFont.Call(
		uintptr(hEle),
		uintptr(hFontx))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:36:58
// @Function: XEleGetFont
// @Description: 获取元素字体.
// @Calls: XEle_GetFont
// @Input: hEle 元素句柄.
// @Return: 返回炫彩字体句柄.
// *******************************************
func XEleGetFont(hEle HELE) HFONTX {
	ret, _, _ := XEle_GetFont.Call(uintptr(hEle))

	return HFONTX(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:38:15
// @Function: XEleSetAlpha
// @Description: 设置元素透明度.
// @Calls: XEle_SetAlpha
// @Input: hEle 元素句柄. alpha 透明度.
// @Return:
// *******************************************
func XEleSetAlpha(hEle HELE, alpha byte) {
	XEle_SetAlpha.Call(
		uintptr(hEle),
		uintptr(alpha))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:39:30
// @Function: XEleDestroy
// @Description: 销毁元素.
// @Calls: XEle_Destroy
// @Input: hEle 元素句柄.
// @Return:
// *******************************************
func XEleDestroy(hEle HELE) {
	XEle_Destroy.Call(uintptr(hEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:40:47
// @Function: XEleAddBkBorder
// @Description: 添加背景内容边框.
// @Calls: XEle_AddBkBorder
// @Input: hEle 元素句柄. color RGB颜色. alpha 透明度. width 线宽.
// @Return:
// *******************************************
func XEleAddBkBorder(hEle HELE, color COLORREF, alpha byte, width int) {
	XEle_AddBkBorder.Call(
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
// @Calls: XEle_AddBkFill
// @Input: hEle 元素句柄. color RGB颜色. alpha 透明度.
// @Return:
// *******************************************
func XEleAddBkFill(hEle HELE, color COLORREF, alpha byte) {
	XEle_AddBkFill.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:43:45
// @Function: XEleAddBkImage
// @Description: 添加背景内容图片.
// @Calls: XEle_AddBkImage
// @Input: hEle 元素句柄. hImage 图片句柄.
// @Return:
// *******************************************
func XEleAddBkImage(hEle HELE, hImage HIMAGE) {
	XEle_AddBkImage.Call(
		uintptr(hEle),
		uintptr(hImage))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:46:00
// @Function: XEleGetBkInfoCount
// @Description: 获取背景内容数量.
// @Calls: XEle_GetBkInfoCount
// @Input: hEle 元素句柄.
// @Return: 返回背景内容数量.
// *******************************************
func XEleGetBkInfoCount(hEle HELE) int {
	ret, _, _ := XEle_GetBkInfoCount.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:47:25
// @Function: XEleClearBkInfo
// @Description: 清空背景内容; 如果背景没有内容,将使用系统默认内容,以便保证背景正确.
// @Calls: XEle_ClearBkInfo
// @Input: hEle 元素句柄.
// @Return:
// *******************************************
func XEleClearBkInfo(hEle HELE) {
	XEle_ClearBkInfo.Call(uintptr(hEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:48:33
// @Function: XEleGetBkInfoManager
// @Description: 获取元素背景内容管理器.
// @Calls: XEle_GetBkInfoManager
// @Input: hEle 元素句柄.
// @Return: 背景内容管理器.
// *******************************************
func XEleGetBkInfoManager(hEle HELE) HBKINFOM {
	ret, _, _ := XEle_GetBkInfoManager.Call(uintptr(hEle))

	return HBKINFOM(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 21:48:33
// @Function: XEleDrawFocus
// @Description: 绘制元素焦点.
// @Calls: XEle_DrawFocus
// @Input: hEle 元素句柄. hDraw 图形绘制句柄. pRect 区域坐标.
// @Return: 绘制成功返回TRUE,如果不需要绘制焦点返回FALSE.
// *******************************************
func XEleDrawFocus(hEle HELE, hDraw HDRAW, pRect *RECT) bool {
	ret, _, _ := XEle_DrawFocus.Call(
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
// @Calls: XEle_BindLayoutObject
// @Input: hEle 元素句柄. hLayout 布局对象句柄.
// @Return:
// *******************************************
func XEleBindLayoutObject(hEle HELE, hLayout HXCGUI) {
	XEle_BindLayoutObject.Call(
		uintptr(hEle),
		uintptr(hLayout))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:13:10
// @Function: XEleGetLayoutObject
// @Description: 获取绑定的布局对象.
// @Calls: XEle_GetLayoutObject
// @Input: hEle 元素句柄.
// @Return: 布局对象句柄.
// *******************************************
func XEleGetLayoutObject(hEle HELE) HXCGUI {
	ret, _, _ := XEle_GetLayoutObject.Call(uintptr(hEle))

	return HXCGUI(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:15:21
// @Function: XEleGetParentLayoutObject
// @Description: 获取父布局对象.
// @Calls: XEle_GetParentLayoutObject
// @Input: hEle 元素句柄.
// @Return: 布局对象句柄.
// *******************************************
func XEleGetParentLayoutObject(hEle HELE) HXCGUI {
	ret, _, _ := XEle_GetParentLayoutObject.Call(uintptr(hEle))

	return HXCGUI(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:16:37
// @Function: XEleSetUserData
// @Description: 设置用户数据.
// @Calls: XEle_SetUserData
// @Input: hEle 元素句柄. nData 用户数据.
// @Return:
// *******************************************
func XEleSetUserData(hEle HELE, nData int) {
	XEle_SetUserData.Call(
		uintptr(hEle),
		uintptr(nData))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:18:26
// @Function: XEleGetUserData
// @Description: 获取用户数据.
// @Calls: XEle_GetUserData
// @Input: hEle 元素句柄.
// @Return: 用户数据.
// *******************************************
func XEleGetUserData(hEle HELE) int {
	ret, _, _ := XEle_GetUserData.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-23 17:59:36
// @Function: XEleGetContentSize
// @Description: 获取内容大小.
// @Calls: XEle_GetContentSize
// @Input: hEle 元素句柄. pSize 大小.
// @Return: 元素内容大小.
// *******************************************
func XEleGetContentSize(hEle HELE, pSize *uint16) {
	XEle_GetContentSize.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pSize)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:27:57
// @Function: XEleSetCapture
// @Description: 设置鼠标捕获.
// @Calls: XEle_SetCapture
// @Input: hEle 元素句柄. b TRUE设置,FALSE取消.
// @Return:
// *******************************************
func XEleSetCapture(hEle HELE, b bool) {
	XEle_SetCapture.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(b)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:32:55
// @Function: XEleSetLayoutWidth
// @Description: 设置布局宽度.
// @Calls: XEle_SetLayoutWidth
// @Input: hEle 元素句柄. nType 属性标识. nWidth 宽度.
// @Return:
// *******************************************
func XEleSetLayoutWidth(hEle HELE, nType LAYOUT_SIZE_TYPE_, nWidth int) {
	XEle_SetLayoutWidth.Call(
		uintptr(hEle),
		uintptr(nType),
		uintptr(nWidth))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:33:19
// @Function: XEleSetLayoutHeight
// @Description: 设备布局高度.
// @Calls: XEle_SetLayoutHeight
// @Input: hEle 元素句柄. nType 属性标识. nHeight 高度.
// @Return:
// *******************************************
func XEleSetLayoutHeight(hEle HELE, nType LAYOUT_SIZE_TYPE_, nHeight int) {
	XEle_SetLayoutHeight.Call(
		uintptr(hEle),
		uintptr(nType),
		uintptr(nHeight))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:34:23
// @Function: XEleGetLayoutWidth
// @Description: 获取布局宽度.
// @Calls: XEle_GetLayoutWidth
// @Input: hEle 元素句柄. nType 属性标识. pWidth 宽度.
// @Return:
// *******************************************
func XEleGetLayoutWidth(hEle HELE, nType *LAYOUT_SIZE_TYPE_, nWidth *uint16) {
	XEle_GetLayoutWidth.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(nType)),
		uintptr(unsafe.Pointer(nWidth)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:37:16
// @Function: XEleGetLayoutHeight
// @Description: 获取布局宽度.
// @Calls: XEle_GetLayoutHeight
// @Input: hEle 元素句柄. nType 属性标识. nHeight 高度.
// @Return:
// *******************************************
func XEleGetLayoutHeight(hEle HELE, nType *LAYOUT_SIZE_TYPE_, nHeight *uint16) {
	XEle_GetLayoutHeight.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(nType)),
		uintptr(unsafe.Pointer(nHeight)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:38:00
// @Function: XEleEnableTransparentChannel
// @Description: 启用或关闭元素透明通道,如果启用,将强制设置元素背景不透明,默认为启用,此功能是为了兼容GDI不支持透明通道问题.
// @Calls: XEle_EnableTransparentChannel
// @Input: hEle 元素句柄. bEnable 启用或关闭.
// @Return:
// *******************************************
func XEleEnableTransparentChannel(hEle HELE, bEnable bool) {
	XEle_EnableTransparentChannel.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:40:01
// @Function: XEleSetToolTip
// @Description: 设置工具提示内容.
// @Calls: XEle_SetToolTip
// @Input: hEle 元素句柄. pText 工具提示内容.
// @Return:
// *******************************************
func XEleSetToolTip(hEle HELE, pText string) {
	XEle_SetToolTip.Call(
		uintptr(hEle),
		StringToUintPtr(pText))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:41:20
// @Function: XEleGetToolTip
// @Description: 获取工具提示内容.
// @Calls: XEle_GetToolTip
// @Input: hEle 元素句柄. pOut 接收内容缓冲区. nOutLen 缓冲区大小,字符为单位.
// @Return:
// *******************************************
func XEleGetToolTip(hEle HELE, pOut *uint16, nOutLen int) {
	XEle_GetToolTip.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:43:33
// @Function: XEleEnableToolTip
// @Description: 启用工具提示.
// @Calls: XEle_EnableToolTip
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XEleEnableToolTip(hEle HELE, bEnable bool) {
	XEle_EnableToolTip.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:45:01
// @Function: XEleAdjustLayoutObject
// @Description: 调整布局对象.
// @Calls: XEle_AdjustLayoutObject
// @Input: hEle 元素句柄.
// @Return:
// *******************************************
func XEleAdjustLayoutObject(hEle HELE) {
	XEle_AdjustLayoutObject.Call(uintptr(hEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-24 23:46:13
// @Function: XEleAdjustLayout
// @Description: 调整布局.
// @Calls: XEle_AdjustLayout
// @Input: hEle 元素句柄.
// @Return:
// *******************************************
func XEleAdjustLayout(hEle HELE) {
	XEle_AdjustLayout.Call(uintptr(hEle))
}
