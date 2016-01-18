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
	xEle_SetRectEx                  *syscall.Proc
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
	xEle_AddShape                   *syscall.Proc
	xEle_SetZOrder                  *syscall.Proc
	xEle_SetZOrderEx                *syscall.Proc
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
	xEle_GetStateFlags              *syscall.Proc
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
	xEle_Create = xcDLL.MustFindProc("XEle_Create")
	xEle_RegEventC = xcDLL.MustFindProc("XEle_RegEventC")
	xEle_RegEventC1 = xcDLL.MustFindProc("XEle_RegEventC1")
	xEle_RegEventC2 = xcDLL.MustFindProc("XEle_RegEventC2")
	xEle_RemoveEventC = xcDLL.MustFindProc("XEle_RemoveEventC")
	xEle_SendEvent = xcDLL.MustFindProc("XEle_SendEvent")
	xEle_PostEvent = xcDLL.MustFindProc("XEle_PostEvent")
	xEle_GetRect = xcDLL.MustFindProc("XEle_GetRect")
	xEle_GetRectLogic = xcDLL.MustFindProc("XEle_GetRectLogic")
	xEle_GetClientRect = xcDLL.MustFindProc("XEle_GetClientRect")
	xEle_GetWidth = xcDLL.MustFindProc("XEle_GetWidth")
	xEle_GetHeight = xcDLL.MustFindProc("XEle_GetHeight")
	xEle_RectWndClientToEleClient = xcDLL.MustFindProc("XEle_RectWndClientToEleClient")
	xEle_PointWndClientToEleClient = xcDLL.MustFindProc("XEle_PointWndClientToEleClient")
	xEle_RectClientToWndClient = xcDLL.MustFindProc("XEle_RectClientToWndClient")
	xEle_PointClientToWndClient = xcDLL.MustFindProc("XEle_PointClientToWndClient")
	xEle_GetWndClientRect = xcDLL.MustFindProc("XEle_GetWndClientRect")
	xEle_GetType = xcDLL.MustFindProc("XEle_GetType")
	xEle_GetHWND = xcDLL.MustFindProc("XEle_GetHWND")
	xEle_GetHWINDOW = xcDLL.MustFindProc("XEle_GetHWINDOW")
	xEle_GetCursor = xcDLL.MustFindProc("XEle_GetCursor")
	xEle_SetCursor = xcDLL.MustFindProc("XEle_SetCursor")
	xEle_AddEle = xcDLL.MustFindProc("XEle_AddEle")
	xEle_InsertEle = xcDLL.MustFindProc("XEle_InsertEle")
	xEle_ShowEle = xcDLL.MustFindProc("XEle_ShowEle")
	xEle_SetRect = xcDLL.MustFindProc("XEle_SetRect")
	xEle_SetRectEx = xcDLL.MustFindProc("XEle_SetRectEx")
	xEle_SetRectLogic = xcDLL.MustFindProc("XEle_SetRectLogic")
	xEle_SetID = xcDLL.MustFindProc("XEle_SetID")
	xEle_GetID = xcDLL.MustFindProc("XEle_GetID")
	xEle_IsShow = xcDLL.MustFindProc("XEle_IsShow")
	xEle_IsDrawFocus = xcDLL.MustFindProc("XEle_IsDrawFocus")
	xEle_IsEnable = xcDLL.MustFindProc("XEle_IsEnable")
	xEle_IsEnableFocus = xcDLL.MustFindProc("XEle_IsEnableFocus")
	xEle_IsMouseThrough = xcDLL.MustFindProc("XEle_IsMouseThrough")
	xEle_HitChildEle = xcDLL.MustFindProc("XEle_HitChildEle")
	xEle_IsBkTransparent = xcDLL.MustFindProc("XEle_IsBkTransparent")
	xEle_IsEnableEvent_XE_PAINT_END = xcDLL.MustFindProc("XEle_IsEnableEvent_XE_PAINT_END")
	xEle_IsKeyTab = xcDLL.MustFindProc("XEle_IsKeyTab")
	xEle_IsSwitchFocus = xcDLL.MustFindProc("XEle_IsSwitchFocus")
	xEle_IsEnable_XE_MOUSEWHEEL = xcDLL.MustFindProc("XEle_IsEnable_XE_MOUSEWHEEL")
	xEle_Enable = xcDLL.MustFindProc("XEle_Enable")
	xEle_EnableFocus = xcDLL.MustFindProc("XEle_EnableFocus")
	xEle_EnableDrawFocus = xcDLL.MustFindProc("XEle_EnableDrawFocus")
	xEle_EnableEvent_XE_PAINT_END = xcDLL.MustFindProc("XEle_EnableEvent_XE_PAINT_END")
	xEle_EnableBkTransparent = xcDLL.MustFindProc("XEle_EnableBkTransparent")
	xEle_EnableMouseThrough = xcDLL.MustFindProc("XEle_EnableMouseThrough")
	xEle_EnableKeyTab = xcDLL.MustFindProc("XEle_EnableKeyTab")
	xEle_EnableSwitchFocus = xcDLL.MustFindProc("XEle_EnableSwitchFocus")
	xEle_EnableEvent_XE_MOUSEWHEEL = xcDLL.MustFindProc("XEle_EnableEvent_XE_MOUSEWHEEL")
	xEle_GetParentEle = xcDLL.MustFindProc("XEle_GetParentEle")
	xEle_GetParent = xcDLL.MustFindProc("XEle_GetParent")
	xEle_RemoveEle = xcDLL.MustFindProc("XEle_RemoveEle")
	xEle_AddShape = xcDLL.MustFindProc("XEle_AddShape")
	xEle_SetZOrder = xcDLL.MustFindProc("XEle_SetZOrder")
	xEle_SetZOrderEx = xcDLL.MustFindProc("XEle_SetZOrderEx")
	xEle_RedrawEle = xcDLL.MustFindProc("XEle_RedrawEle")
	xEle_RedrawRect = xcDLL.MustFindProc("XEle_RedrawRect")
	xEle_GetChildCount = xcDLL.MustFindProc("XEle_GetChildCount")
	xEle_GetChildByIndex = xcDLL.MustFindProc("XEle_GetChildByIndex")
	xEle_GetChildByID = xcDLL.MustFindProc("XEle_GetChildByID")
	xEle_GetChildShapeCount = xcDLL.MustFindProc("XEle_GetChildShapeCount")
	xEle_GetChildShapeByIndex = xcDLL.MustFindProc("XEle_GetChildShapeByIndex")
	xEle_SetTextColor = xcDLL.MustFindProc("XEle_SetTextColor")
	xEle_GetTextColor = xcDLL.MustFindProc("XEle_GetTextColor")
	xEle_SetFocusBorderColor = xcDLL.MustFindProc("XEle_SetFocusBorderColor")
	xEle_GetFocusBorderColor = xcDLL.MustFindProc("XEle_GetFocusBorderColor")
	xEle_SetFont = xcDLL.MustFindProc("XEle_SetFont")
	xEle_GetFont = xcDLL.MustFindProc("XEle_GetFont")
	xEle_SetAlpha = xcDLL.MustFindProc("XEle_SetAlpha")
	xEle_Destroy = xcDLL.MustFindProc("XEle_Destroy")
	xEle_AddBkBorder = xcDLL.MustFindProc("XEle_AddBkBorder")
	xEle_AddBkFill = xcDLL.MustFindProc("XEle_AddBkFill")
	xEle_AddBkImage = xcDLL.MustFindProc("XEle_AddBkImage")
	xEle_GetBkInfoCount = xcDLL.MustFindProc("XEle_GetBkInfoCount")
	xEle_ClearBkInfo = xcDLL.MustFindProc("XEle_ClearBkInfo")
	xEle_GetBkInfoManager = xcDLL.MustFindProc("XEle_GetBkInfoManager")
	xEle_GetStateFlags = xcDLL.MustFindProc("XEle_GetStateFlags")
	xEle_DrawFocus = xcDLL.MustFindProc("XEle_DrawFocus")
	xEle_BindLayoutObject = xcDLL.MustFindProc("XEle_BindLayoutObject")
	xEle_GetLayoutObject = xcDLL.MustFindProc("XEle_GetLayoutObject")
	xEle_GetParentLayoutObject = xcDLL.MustFindProc("XEle_GetParentLayoutObject")
	xEle_SetUserData = xcDLL.MustFindProc("XEle_SetUserData")
	xEle_GetUserData = xcDLL.MustFindProc("XEle_GetUserData")
	xEle_GetContentSize = xcDLL.MustFindProc("XEle_GetContentSize")
	xEle_SetCapture = xcDLL.MustFindProc("XEle_SetCapture")
	xEle_SetLayoutWidth = xcDLL.MustFindProc("XEle_SetLayoutWidth")
	xEle_SetLayoutHeight = xcDLL.MustFindProc("XEle_SetLayoutHeight")
	xEle_GetLayoutWidth = xcDLL.MustFindProc("XEle_GetLayoutWidth")
	xEle_GetLayoutHeight = xcDLL.MustFindProc("XEle_GetLayoutHeight")
	xEle_EnableTransparentChannel = xcDLL.MustFindProc("XEle_EnableTransparentChannel")
	xEle_SetToolTip = xcDLL.MustFindProc("XEle_SetToolTip")
	xEle_GetToolTip = xcDLL.MustFindProc("XEle_GetToolTip")
	xEle_EnableToolTip = xcDLL.MustFindProc("XEle_EnableToolTip")
	xEle_AdjustLayoutObject = xcDLL.MustFindProc("XEle_AdjustLayoutObject")
	xEle_AdjustLayout = xcDLL.MustFindProc("XEle_AdjustLayout")
}

/*
创建基础元素.

参数:
	x 元素x坐标.
	y 元素y坐标.
	cx 宽度.
	cy 高度.
	hParent 父是窗口句柄或元素句柄.
返回:
	元素句柄.
*/
func XEle_Create(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xEle_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
注册事件C方式,省略2参数.

参数:
	hEle 元素句柄.
	nEvent 事件类型.
	pFun 事件函数指针.
*/
func XEle_RegEventC(hEle HELE, nEvent int, pFun uintptr) {
	xEle_RegEventC.Call(
		uintptr(hEle),
		uintptr(nEvent),
		pFun)
}

/*
注册事件C方式,省略1参数.

参数:
	hEle 元素句柄.
	nEvent 事件类型.
	pFun 事件函数指针.
*/
func XEle_RegEventC1(hEle HELE, nEvent int, pFun uintptr) {
	xEle_RegEventC1.Call(
		uintptr(hEle),
		uintptr(nEvent),
		pFun)
}

/*
注册事件C方式,不省略参数.

参数:
	hEle 元素句柄.
	nEvent 事件类型.
	pFun 事件函数指针.
*/
func XEle_RegEventC2(hEle HELE, nEvent int, pFun uintptr) {
	xEle_RegEventC2.Call(
		uintptr(hEle),
		uintptr(nEvent),
		pFun)
}

/*
移除事件函数C方式.

参数:
	hEle 元素句柄.
	nEvent 事件类型.
	pFun 事件函数指针.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XEle_RemoveEventC(hEle HELE, nEvent int, pFun uintptr) bool {
	ret, _, _ := xEle_RemoveEventC.Call(
		uintptr(hEle),
		uintptr(nEvent),
		pFun)

	return ret == TRUE
}

/*
发送事件.

参数:
	hEle 目标元素句柄.
	hEventEle 触发事件元素句柄.
	nEvent 事件类型.
	wParam 参数.
	lParam 参数.
返回:
	事件返回值.
*/
func XEle_SendEvent(hEle, hEventEle HELE, nEvent int, wParam, lParam uintptr) int {
	ret, _, _ := xEle_SendEvent.Call(
		uintptr(hEle),
		uintptr(hEventEle),
		uintptr(nEvent),
		wParam,
		lParam)

	return int(ret)
}

/*
POST事件.

参数:
	hEle 目标元素句柄.
	hEventEle 触发事件元素句柄.
	nEvent 事件类型.
	wParam 参数.
	lParam 参数.
返回:
	事件返回值.
*/
func XEle_PostEvent(hEle, hEventEle HELE, nEvent int, wParam, lParam uintptr) int {
	ret, _, _ := xEle_PostEvent.Call(
		uintptr(hEle),
		uintptr(hEventEle),
		uintptr(nEvent),
		wParam,
		lParam)

	return int(ret)
}

/*
获取元素坐标.

参数:
	hEle 元素句柄.
	pRect 坐标.
*/
func XEle_GetRect(hEle HELE, pRect *RECT) {
	xEle_GetRect.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

/*
获取元素坐标,逻辑坐标,包含滚动视图偏移.

参数:
	hEle 元素句柄.
	pRect 坐标.
*/
func XEle_GetRectLogic(hEle HELE, pRect *RECT) {
	xEle_GetRectLogic.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

/*
获取元素客户区坐标.

参数:
	hEle 元素句柄.
	pRect 坐标.
*/
func XEle_GetClientRect(hEle HELE, pRect *RECT) {
	xEle_GetClientRect.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

/*
获取元素宽度.

参数:
	hEle 元素句柄.
返回:
	宽度.
*/
func XEle_GetWidth(hEle HELE) int {
	ret, _, _ := xEle_GetWidth.Call(uintptr(hEle))

	return int(ret)
}

/*
获取元素高度.

参数:
	hEle 元素句柄.
返回:
	高度.
*/
func XEle_GetHeight(hEle HELE) int {
	ret, _, _ := xEle_GetHeight.Call(uintptr(hEle))

	return int(ret)
}

/*
窗口客户区坐标转换到元素客户区坐标.

参数:
	hEle 元素句柄.
	pRect 坐标.
*/
func XEle_RectWndClientToEleClient(hEle HELE, pRect *RECT) {
	xEle_RectWndClientToEleClient.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

/*
窗口客户区坐标转换到元素客户区坐标.

参数:
	hEle 元素句柄.
	pPt 坐标.
*/
func XEle_PointWndClientToEleClient(hEle HELE, pRect *RECT) {
	xEle_PointWndClientToEleClient.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

/*
元素客户区坐标转换到窗口客户区坐标.

参数:
	hEle 元素句柄.
	pRect 坐标.
*/
func XEle_RectClientToWndClient(hEle HELE, pRect *RECT) {
	xEle_RectClientToWndClient.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

/*
元素客户区坐标转换到窗口客户区坐标.

参数:
	hEle 元素句柄.
	pPt 坐标.
*/
func XEle_PointClientToWndClient(hEle HELE, pRect *RECT) {
	xEle_PointClientToWndClient.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

/*
元素基于窗口客户区坐标.

参数:
	hEle 元素句柄.
	pRect 坐标.
*/
func XEle_GetWndClientRect(hEle HELE, pRect *RECT) {
	xEle_GetWndClientRect.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

/*
获取元素类型.

参数:
	hEle 元素句柄.
返回:
	元素类型,参见宏定义.
*/
func XEle_GetType(hEle HELE) XC_OBJECT_TYPE {
	ret, _, _ := xEle_GetType.Call(uintptr(hEle))

	return XC_OBJECT_TYPE(ret)
}

/*
获取HWND句柄.

参数:
	hEle 元素句柄.
返回:
	HWND句柄.
*/
func XEle_GetHWND(hEle HELE) HWND {
	ret, _, _ := xEle_GetHWND.Call(uintptr(hEle))

	return HWND(ret)
}

/*
获取HWINDOW句柄.

参数:
	hEle 元素句柄.
返回:
	HWINDOW句柄.
*/
func XEle_GetHWINDOW(hEle HELE) HWINDOW {
	ret, _, _ := xEle_GetHWINDOW.Call(uintptr(hEle))

	return HWINDOW(ret)
}

/*
获取元素鼠标光标.

参数:
	hEle 元素句柄.
返回:
	返回光标句柄.
*/
func XEle_GetCursor(hEle HELE) HCURSOR {
	ret, _, _ := xEle_GetCursor.Call(uintptr(hEle))

	return HCURSOR(ret)
}

/*
设置元素鼠标光标.

参数:
	hEle 元素句柄.
	hCursor 光标句柄.
*/
func XEle_SetCursor(hEle HELE, hCursor HCURSOR) {
	xEle_SetCursor.Call(
		uintptr(hEle),
		uintptr(hCursor))
}

/*
添加元素.添加一个子元素,到当前元素下.

参数:
	hEle 元素句柄.
	hChildEle 要添加的子元素资源句柄.
返回:
	如果成功返回TRUE,否则相反.
*/
func XEle_AddEle(hEle, hChildEle HELE) bool {
	ret, _, _ := xEle_AddEle.Call(
		uintptr(hEle),
		uintptr(hChildEle))

	return ret == TRUE
}

/*
插入元素到目标元素前面.

参数:
	hEle 元素句柄.
	hChildEle 要插入的子元素句柄.
	hDestEle 目标元素.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XEle_InsertEle(hEle, hChildEle, hDestEle HELE) bool {
	ret, _, _ := xEle_InsertEle.Call(
		uintptr(hEle),
		uintptr(hChildEle),
		uintptr(hDestEle))

	return ret == TRUE
}

/*
显示隐藏元素.

参数:
	hEle 元素句柄.
	bShow TRUE显示元素,否则相反.
*/
func XEle_ShowEle(hEle HELE, bShow bool) {
	xEle_ShowEle.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bShow)))
}

/*
设置元素坐标.

参数:
	hEle 元素句柄.
	pRect 坐标.
	bRedraw 是否重绘.
返回:
	如果坐标未改变返回FALSE,否则相反.
*/
func XEle_SetRect(hEle HELE, pRect *RECT, bRedraw bool) bool {
	ret, _, _ := xEle_SetRect.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(BoolToBOOL(bRedraw)))

	return ret == FALSE
}

/*
设置元素坐标.

参数:
	hEle 元素句柄.
	x X坐标.
	y Y坐标.
	cx 宽度.
	cy 高度.
	bRedraw 是否重绘.
返回:
	如果坐标未改变返回FALSE,否则相反.
*/
func XEle_SetRectEx(hEle HELE, x, y, cx, cy int, bRedraw bool) bool {
	ret, _, _ := xEle_SetRectEx.Call(
		uintptr(hEle),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(BoolToBOOL(bRedraw)))

	return ret == TRUE
}

/*
设置元素坐标,逻辑坐标,包含滚动视图偏移.

参数:
	hEle 元素句柄.
	pRect 坐标.
	bRedraw 是否重绘.
返回:
	如果坐标未改变返回FALSE,否则相反.
*/
func XEle_SetRectLogic(hEle HELE, pRect *RECT, bRedraw bool) bool {
	ret, _, _ := xEle_SetRectLogic.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(BoolToBOOL(bRedraw)))

	return ret == FALSE
}

/*
设置元素ID.

参数:
	hEle 元素句柄.
	id ID值.
*/
func XEle_SetID(hEle HELE, id int) {
	xEle_SetID.Call(
		uintptr(hEle),
		uintptr(id))
}

/*
获取元素ID.

参数:
	hEle 元素句柄.
	返回:返回元素ID.
*/
func XEle_GetID(hEle HELE) int {
	ret, _, _ := xEle_GetID.Call(uintptr(hEle))

	return int(ret)
}

/*
元素是否显示.

参数:
	hEle 元素句柄.
返回:
	如果显示返回TRUE否则返回FALSE.
*/
func XEle_IsShow(hEle HELE) bool {
	ret, _, _ := xEle_IsShow.Call(uintptr(hEle))

	return ret == TRUE
}

/*
元素是否绘制焦点.

参数:
	hEle 元素句柄.
返回:
	如果绘制焦点返回TRUE否则返回FALSE.
*/
func XEle_IsDrawFocus(hEle HELE) bool {
	ret, _, _ := xEle_IsDrawFocus.Call(uintptr(hEle))

	return ret == TRUE
}

/*
元素是否为启用状态.

参数:
	hEle 元素句柄.
返回:
	如果启用状态返回TRUE否则返回FALSE.
*/
func XEle_IsEnable(hEle HELE) bool {
	ret, _, _ := xEle_IsEnable.Call(uintptr(hEle))

	return ret == TRUE
}

/*
元素是否启用焦点.

参数:
	hEle 元素句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XEle_IsEnableFocus(hEle HELE) bool {
	ret, _, _ := xEle_IsEnableFocus.Call(uintptr(hEle))

	return ret == TRUE
}

/*
元素是否启用鼠标穿透.

参数:
	hEle 元素句柄.
	返回:成功返回TRUE否则返回FALSE.
*/
func XEle_IsMouseThrough(hEle HELE) bool {
	ret, _, _ := xEle_IsMouseThrough.Call(uintptr(hEle))

	return ret == TRUE
}

/*
检测坐标点所在元素,包含子元素的子元素.

参数:
	hEle 元素句柄.
	pPt 坐标点.
返回:
	成功返回元素句柄,否则返回NULL.
*/
func XEle_HitChildEle(hEle HELE, pPt *POINT) HELE {
	ret, _, _ := xEle_HitChildEle.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)))

	return HELE(ret)
}

/*
是否背景透明.

参数:
	hEle 元素句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XEle_IsBkTransparent(hEle HELE) bool {
	ret, _, _ := xEle_IsBkTransparent.Call(uintptr(hEle))

	return ret == TRUE
}

/*
是否启XE_PAINT_END用事件.

参数:
	hEle 元素句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XEle_IsEnableEvent_XE_PAINT_END(hEle HELE) bool {
	ret, _, _ := xEle_IsEnableEvent_XE_PAINT_END.Call(uintptr(hEle))

	return ret == TRUE
}

/*
是否接收Tab键.

参数:
	hEle 元素句柄.
返回:
	是返回TRUE否则返回FALSE.
*/
func XEle_IsKeyTab(hEle HELE) bool {
	ret, _, _ := xEle_IsKeyTab.Call(uintptr(hEle))

	return ret == TRUE
}

/*
是否接受通过键盘切换焦点.

参数:
	hEle 元素句柄.
返回:
	是返回TRUE否则返回FALSE.
*/
func XEle_IsSwitchFocus(hEle HELE) bool {
	ret, _, _ := xEle_IsSwitchFocus.Call(uintptr(hEle))

	return ret == TRUE
}

/*
判断是否启用鼠标滚动事件,如果禁用那么事件会发送给他的父元素.

参数:
	hEle 元素句柄.
返回:
	是返回TRUE否则返回FALSE.
*/
func XEle_IsEnable_XE_MOUSEWHEEL(hEle HELE) bool {
	ret, _, _ := xEle_IsEnable_XE_MOUSEWHEEL.Call(uintptr(hEle))

	return ret == TRUE
}

/*
启用或禁用元素.

参数:
	hEle 元素句柄.
	bEnable 启用或禁用.
*/
func XEle_Enable(hEle HELE, bEnable bool) {
	xEle_Enable.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用焦点.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XEle_EnableFocus(hEle HELE, bEnable bool) {
	xEle_EnableFocus.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用绘制焦点.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XEle_EnableDrawFocus(hEle HELE, bEnable bool) {
	xEle_EnableDrawFocus.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用XE_PAINT_END事件.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XEle_EnableEvent_XE_PAINT_END(hEle HELE, bEnable bool) {
	xEle_EnableEvent_XE_PAINT_END.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用背景透明.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XEle_EnableBkTransparent(hEle HELE, bEnable bool) {
	xEle_EnableBkTransparent.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用鼠标穿透.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XEle_EnableMouseThrough(hEle HELE, bEnable bool) {
	xEle_EnableMouseThrough.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用接收Tab键.

参数:
	hEle 元素句柄.
	bEnable XX.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XEle_EnableKeyTab(hEle HELE, bEnable bool) {
	xEle_EnableKeyTab.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用接受通过键盘切换焦点.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XEle_EnableSwitchFocus(hEle HELE, bEnable bool) {
	xEle_EnableSwitchFocus.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用接收鼠标滚动事件,如果禁用那么事件会传递给父元素.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XEle_EnableEvent_XE_MOUSEWHEEL(hEle HELE, bEnable bool) {
	xEle_EnableEvent_XE_MOUSEWHEEL.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
获取父元素.

参数:
	hEle 元素句柄.
返回:
	元素句柄.
*/
func XEle_GetParentEle(hEle HELE) HELE {
	ret, _, _ := xEle_GetParentEle.Call(uintptr(hEle))

	return HELE(ret)
}

/*
获取父对象,父可能是元素或窗口,通过此函数可以检查是否有父.

参数:
	hEle 元素句柄.
返回:
	对象句柄.
*/
func XEle_GetParent(hEle HELE) HXCGUI {
	ret, _, _ := xEle_GetParent.Call(uintptr(hEle))

	return HXCGUI(ret)
}

/*
移除元素,但不销毁.

参数:
	hEle 元素句柄.
*/
func XEle_RemoveEle(hEle HELE) {
	xEle_RemoveEle.Call(uintptr(hEle))
}

/*
添加形状对象.

参数:
	hEle 元素句柄.
	hShape 形状对象句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XEle_AddShape(hEle HELE, hShape HXCGUI) bool {
	ret, _, _ := xEle_AddShape.Call(
		uintptr(hEle),
		uintptr(hShape))

	return ret == TRUE
}

/*
设置元素Z序.

参数:
	hEle 元素句柄.
	index 位置索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XEle_SetZOrder(hEle HELE, index int) bool {
	ret, _, _ := xEle_SetZOrder.Call(
		uintptr(hEle),
		uintptr(index))

	return ret == TRUE
}

/*
设置元素Z序.

参数:
	hEle 元素句柄.
	hDestEle 目标元素.
	nType 类型.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XEle_SetZOrderEx(hEle, hDestEle HELE, nType int) bool {
	ret, _, _ := xEle_SetZOrderEx.Call(
		uintptr(hEle),
		uintptr(hDestEle),
		uintptr(nType))

	return ret == TRUE
}

/*
重绘元素.

参数:
	hEle 元素句柄.
*/
func XEle_RedrawEle(hEle HELE) {
	xEle_RedrawEle.Call(uintptr(hEle))
}

/*
重绘元素指定区域.

参数:
	hEle 元素句柄.
	pRect 相对于元素客户区坐标.
*/
func XEle_RedrawRect(hEle HELE, pRect *RECT) {
	xEle_RedrawRect.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

/*
获取子元素数量,只检测当前层元素.

参数:
	hEle 元素句柄.
返回:
	子元素数量.
*/
func XEle_GetChildCount(hEle HELE) int {
	ret, _, _ := xEle_GetChildCount.Call(uintptr(hEle))

	return int(ret)
}

/*
获取子元素通过索引,只检测当前层元素.

参数:
	hEle 元素句柄.
	index 索引.
返回:
	元素句柄.
*/
func XEle_GetChildByIndex(hEle HELE, index int) HELE {
	ret, _, _ := xEle_GetChildByIndex.Call(
		uintptr(hEle),
		uintptr(index))

	return HELE(ret)
}

/*
获取子元素通过ID,只检测当前层元素.

参数:
	hEle 元素句柄.
	id 元素ID.
返回:
	元素句柄.
*/
func XEle_GetChildByID(hEle HELE, id int) HELE {
	ret, _, _ := xEle_GetChildByID.Call(
		uintptr(hEle),
		uintptr(id))

	return HELE(ret)
}

/*
获取子对象(形状对象)数量,只检查当前层.

参数:
	hEle 元素句柄.
返回:
	形状对象数量.
*/
func XEle_GetChildShapeCount(hEle HELE) int {
	ret, _, _ := xEle_GetChildShapeCount.Call(uintptr(hEle))

	return int(ret)
}

/*
通过索引返回形状对象句柄.

参数:
	hEle 元素句柄.
	index 索引.
返回:
	返回形状对象句柄.
*/
func XEle_GetChildShapeByIndex(hEle HELE, index int) HXCGUI {
	ret, _, _ := xEle_GetChildShapeByIndex.Call(
		uintptr(hEle),
		uintptr(index))

	return HXCGUI(ret)
}

/*
设置文本颜色.

参数:
	hEle 元素句柄.
	color RGB颜色值.
	alpha 透明度.
*/
func XEle_SetTextColor(hEle HELE, color COLORREF, alpha byte) {
	xEle_SetTextColor.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))
}

/*
获取文本颜色.

参数:
	hEle 元素句柄.
返回:
	文本颜色值.
*/
func XEle_GetTextColor(hEle HELE) COLORREF {
	ret, _, _ := xEle_GetTextColor.Call(uintptr(hEle))

	return COLORREF(ret)
}

/*
设置焦点边框颜色.

参数:
	hEle 元素句柄.
	color RGB颜色值.
	alpha 透明度.
*/
func XEle_SetFocusBorderColor(hEle HELE, color COLORREF, alpha byte) {
	xEle_SetFocusBorderColor.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))
}

/*
获取焦点边框颜色.

参数:
	hEle 元素句柄.
返回:
	返回颜色值,最高位字节是透明度.
*/
func XEle_GetFocusBorderColor(hEle HELE) COLORREF {
	ret, _, _ := xEle_GetFocusBorderColor.Call(uintptr(hEle))

	return COLORREF(ret)
}

/*
设置元素字体.

参数:
	hEle 元素句柄.
	hFontx 炫彩字体.
*/
func XEle_SetFont(hEle HELE, hFontx HFONTX) {
	xEle_SetFont.Call(
		uintptr(hEle),
		uintptr(hFontx))
}

/*
获取元素字体.

参数:
	hEle 元素句柄.
返回:
	返回炫彩字体句柄.
*/
func XEle_GetFont(hEle HELE) HFONTX {
	ret, _, _ := xEle_GetFont.Call(uintptr(hEle))

	return HFONTX(ret)
}

/*
设置元素透明度.

参数:
	hEle 元素句柄.
	alpha 透明度.
*/
func XEle_SetAlpha(hEle HELE, alpha byte) {
	xEle_SetAlpha.Call(
		uintptr(hEle),
		uintptr(alpha))
}

/*
销毁元素.

参数:
	hEle 元素句柄.
*/
func XEle_Destroy(hEle HELE) {
	xEle_Destroy.Call(uintptr(hEle))
}

/*
添加背景内容边框.

参数:
	hEle 元素句柄.
	color RGB颜色.
	alpha 透明度.
	width 线宽.
*/
func XEle_AddBkBorder(hEle HELE, color COLORREF, alpha byte, width int) {
	xEle_AddBkBorder.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha),
		uintptr(width))
}

/*
添加背景内容填充.

参数:
	hEle 元素句柄.
	color RGB颜色.
	alpha 透明度.
*/
func XEle_AddBkFill(hEle HELE, color COLORREF, alpha byte) {
	xEle_AddBkFill.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))
}

/*
添加背景内容图片.

参数:
	hEle 元素句柄.
	hImage 图片句柄.
*/
func XEle_AddBkImage(hEle HELE, hImage HIMAGE) {
	xEle_AddBkImage.Call(
		uintptr(hEle),
		uintptr(hImage))
}

/*
获取背景内容数量.

参数:
	hEle 元素句柄.
返回:
	返回背景内容数量.
*/
func XEle_GetBkInfoCount(hEle HELE) int {
	ret, _, _ := xEle_GetBkInfoCount.Call(uintptr(hEle))

	return int(ret)
}

/*
清空背景内容; 如果背景没有内容,将使用系统默认内容,以便保证背景正确.

参数:
	hEle 元素句柄.
*/
func XEle_ClearBkInfo(hEle HELE) {
	xEle_ClearBkInfo.Call(uintptr(hEle))
}

/*
获取元素背景内容管理器.

参数:
	hEle 元素句柄.
返回:
	背景内容管理器.
*/
func XEle_GetBkInfoManager(hEle HELE) HBKINFOM {
	ret, _, _ := xEle_GetBkInfoManager.Call(uintptr(hEle))

	return HBKINFOM(ret)
}

/*
获取组合状态.

参数:
	hEle 元素句柄.
返回:
	返回组合状态.
*/
func XEle_GetStateFlags(hEle HELE) int {
	ret, _, _ := xEle_GetStateFlags.Call(uintptr(hEle))

	return int(ret)
}

/*
绘制元素焦点.

参数:
	hEle 元素句柄.
	hDraw 图形绘制句柄.
	pRect 区域坐标.
返回:
	绘制成功返回TRUE,如果不需要绘制焦点返回FALSE.
*/
func XEle_DrawFocus(hEle HELE, hDraw HDRAW, pRect *RECT) bool {
	ret, _, _ := xEle_DrawFocus.Call(
		uintptr(hEle),
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pRect)))

	return ret == TRUE
}

/*
绑定布局对象,只可以绑定一个.

参数:
	hEle 元素句柄.
	hLayout 布局对象句柄.
*/
func XEle_BindLayoutObject(hEle HELE, hLayout HXCGUI) {
	xEle_BindLayoutObject.Call(
		uintptr(hEle),
		uintptr(hLayout))
}

/*
获取绑定的布局对象.

参数:
	hEle 元素句柄.
返回:
	布局对象句柄.
*/
func XEle_GetLayoutObject(hEle HELE) HXCGUI {
	ret, _, _ := xEle_GetLayoutObject.Call(uintptr(hEle))

	return HXCGUI(ret)
}

/*
获取父布局对象.

参数:
	hEle 元素句柄.
返回:
	布局对象句柄.
*/
func XEle_GetParentLayoutObject(hEle HELE) HXCGUI {
	ret, _, _ := xEle_GetParentLayoutObject.Call(uintptr(hEle))

	return HXCGUI(ret)
}

/*
设置用户数据.

参数:
	hEle 元素句柄.
	nData 用户数据.
*/
func XEle_SetUserData(hEle HELE, nData int) {
	xEle_SetUserData.Call(
		uintptr(hEle),
		uintptr(nData))
}

/*
获取用户数据.

参数:
	hEle 元素句柄.
返回:
	用户数据.
*/
func XEle_GetUserData(hEle HELE) int {
	ret, _, _ := xEle_GetUserData.Call(uintptr(hEle))

	return int(ret)
}

/*
获取内容大小.

参数:
	hEle 元素句柄.
	pSize 大小.
*/
func XEle_GetContentSize(hEle HELE, pSize *SIZE) {
	xEle_GetContentSize.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pSize)))
}

/*
设置鼠标捕获.

参数:
	hEle 元素句柄.
	b TRUE设置,FALSE取消.
*/
func XEle_SetCapture(hEle HELE, b bool) {
	xEle_SetCapture.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(b)))
}

/*
设置布局宽度.

参数:
	hEle 元素句柄.
	nType 属性标识.
	nWidth 宽度.
*/
func XEle_SetLayoutWidth(hEle HELE, nType Layout_size_type_, nWidth int) {
	xEle_SetLayoutWidth.Call(
		uintptr(hEle),
		uintptr(nType),
		uintptr(nWidth))
}

/*
设置布局高度.

参数:
	hEle 元素句柄.
	nType 属性标识.
	nHeight 高度.
*/
func XEle_SetLayoutHeight(hEle HELE, nType Layout_size_type_, nHeight int) {
	xEle_SetLayoutHeight.Call(
		uintptr(hEle),
		uintptr(nType),
		uintptr(nHeight))
}

/*
获取布局宽度.

参数:
	hEle 元素句柄.
	pType 属性标识.
	pWidth 宽度.
*/
func XEle_GetLayoutWidth(hEle HELE, nType *Layout_size_type_, nWidth *int) {
	xEle_GetLayoutWidth.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(nType)),
		uintptr(unsafe.Pointer(nWidth)))
}

/*
获取布局高度.

参数:
	hEle 元素句柄.
	pType 属性标识.
	pHeight 高度.
*/
func XEle_GetLayoutHeight(hEle HELE, nType *Layout_size_type_, nHeight *int) {
	xEle_GetLayoutHeight.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(nType)),
		uintptr(unsafe.Pointer(nHeight)))
}

/*
启用或关闭元素透明通道,如果启用,将强制设置元素背景不透明,默认为启用,此功能是为了兼容GDI不支持透明通道问题.

参数:
	hEle 元素句柄.
	bEnable 启用或关闭.
*/
func XEle_EnableTransparentChannel(hEle HELE, bEnable bool) {
	xEle_EnableTransparentChannel.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
设置工具提示内容.

参数:
	hEle 元素句柄.
	pText 工具提示内容.
*/
func XEle_SetToolTip(hEle HELE, pText string) {
	xEle_SetToolTip.Call(
		uintptr(hEle),
		StringToUintPtr(pText))
}

/*
获取工具提示内容.

参数:
	hEle XX.
	pOut 接收内容缓冲区.
	nOutLen 缓冲区大小,字符为单位.
*/
func XEle_GetToolTip(hEle HELE, pOut *uint16, nOutLen int) {
	xEle_GetToolTip.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))
}

/*
启用工具提示.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XEle_EnableToolTip(hEle HELE, bEnable bool) {
	xEle_EnableToolTip.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
调整布局对象.

参数:
	hEle 元素句柄.
*/
func XEle_AdjustLayoutObject(hEle HELE) {
	xEle_AdjustLayoutObject.Call(uintptr(hEle))
}

/*
调整布局.

参数:
	hEle 元素句柄.
*/
func XEle_AdjustLayout(hEle HELE) {
	xEle_AdjustLayout.Call(uintptr(hEle))
}
