package xc

import (
	"syscall"
	"unsafe"
	// "fmt"
)

var (
	// Functions
	XWnd_Create              *syscall.Proc
	XWnd_CreateEx            *syscall.Proc
	XWnd_RegEventC           *syscall.Proc
	XWnd_RegEventC1          *syscall.Proc
	XWnd_RemovegEvent        *syscall.Proc
	XWnd_AddEle              *syscall.Proc
	XWnd_InsertEle           *syscall.Proc
	XWnd_AddShape            *syscall.Proc
	XWnd_RedrawWnd           *syscall.Proc
	XWnd_RedrawWndRect       *syscall.Proc
	XWnd_SetFocusEle         *syscall.Proc
	XWnd_GetFocusEle         *syscall.Proc
	XWnd_SetCursor           *syscall.Proc
	XWnd_GetCursor           *syscall.Proc
	XWnd_GetHWND             *syscall.Proc
	XWnd_EnableDragBorder    *syscall.Proc
	XWnd_EnableDragWindow    *syscall.Proc
	XWnd_EnableDrawBk        *syscall.Proc
	XWnd_EnableAutoFocus     *syscall.Proc
	XWnd_SetCaptureEle       *syscall.Proc
	XWnd_GetCaptureEle       *syscall.Proc
	XWnd_GetDrawRect         *syscall.Proc
	XWnd_ShowWindow          *syscall.Proc
	XWnd_BindLayoutEle       *syscall.Proc
	XWnd_GetLayoutEle        *syscall.Proc
	XWnd_SetCursorSys        *syscall.Proc
	XWnd_SetFont             *syscall.Proc
	XWnd_SetID               *syscall.Proc
	XWnd_GetID               *syscall.Proc
	XWnd_SetLayoutSize       *syscall.Proc
	XWnd_GetLayoutSize       *syscall.Proc
	XWnd_SetDragBorderSize   *syscall.Proc
	XWnd_GetDragBorderSize   *syscall.Proc
	XWnd_SetMinimumSize      *syscall.Proc
	XWnd_HitChildEle         *syscall.Proc
	XWnd_GetChildCount       *syscall.Proc
	XWnd_GetChildByIndex     *syscall.Proc
	XWnd_GetChildByID        *syscall.Proc
	XWnd_GetEle              *syscall.Proc
	XWnd_CloseWindow         *syscall.Proc
	XWnd_BindLayoutObject    *syscall.Proc
	XWnd_GetLayoutObject     *syscall.Proc
	XWnd_AdjustLayout        *syscall.Proc
	XWnd_CreateCaret         *syscall.Proc
	XWnd_SetCaretSize        *syscall.Proc
	XWnd_SetCaretPos         *syscall.Proc
	XWnd_SetCaretPosEx       *syscall.Proc
	XWnd_SetCaretColor       *syscall.Proc
	XWnd_ShowCaret           *syscall.Proc
	XWnd_DestroyCaret        *syscall.Proc
	XWnd_GetCaretHELE        *syscall.Proc
	XWnd_GetClientRect       *syscall.Proc
	XWnd_GetBodyRect         *syscall.Proc
	XWnd_SetTimer            *syscall.Proc
	XWnd_KillTimer           *syscall.Proc
	XWnd_GetBkInfoManager    *syscall.Proc
	XWnd_SetTransparentType  *syscall.Proc
	XWnd_SetTransparentAlpha *syscall.Proc
	XWnd_SetTransparentColor *syscall.Proc
)

func init() {
	// Functions
	XWnd_Create = XCDLL.MustFindProc("XWnd_Create")
	XWnd_CreateEx = XCDLL.MustFindProc("XWnd_CreateEx")
	XWnd_RegEventC = XCDLL.MustFindProc("XWnd_RegEventC")
	XWnd_RegEventC1 = XCDLL.MustFindProc("XWnd_RegEventC1")
	XWnd_RemovegEvent = XCDLL.MustFindProc("XWnd_RemovegEvent")
	XWnd_AddEle = XCDLL.MustFindProc("XWnd_AddEle")
	XWnd_InsertEle = XCDLL.MustFindProc("XWnd_InsertEle")
	XWnd_AddShape = XCDLL.MustFindProc("XWnd_AddShape")
	XWnd_RedrawWnd = XCDLL.MustFindProc("XWnd_RedrawWnd")
	XWnd_RedrawWndRect = XCDLL.MustFindProc("XWnd_RedrawWndRect")
	XWnd_SetFocusEle = XCDLL.MustFindProc("XWnd_SetFocusEle")
	XWnd_GetFocusEle = XCDLL.MustFindProc("XWnd_GetFocusEle")
	XWnd_SetCursor = XCDLL.MustFindProc("XWnd_SetCursor")
	XWnd_GetCursor = XCDLL.MustFindProc("XWnd_GetCursor")
	XWnd_GetHWND = XCDLL.MustFindProc("XWnd_GetHWND")
	XWnd_EnableDragBorder = XCDLL.MustFindProc("XWnd_EnableDragBorder")
	XWnd_EnableDragWindow = XCDLL.MustFindProc("XWnd_EnableDragWindow")
	XWnd_EnableDrawBk = XCDLL.MustFindProc("XWnd_EnableDrawBk")
	XWnd_EnableAutoFocus = XCDLL.MustFindProc("XWnd_EnableAutoFocus")
	XWnd_SetCaptureEle = XCDLL.MustFindProc("XWnd_SetCaptureEle")
	XWnd_GetCaptureEle = XCDLL.MustFindProc("XWnd_GetCaptureEle")
	XWnd_GetDrawRect = XCDLL.MustFindProc("XWnd_GetDrawRect")
	XWnd_ShowWindow = XCDLL.MustFindProc("XWnd_ShowWindow")
	XWnd_BindLayoutEle = XCDLL.MustFindProc("XWnd_BindLayoutEle")
	XWnd_GetLayoutEle = XCDLL.MustFindProc("XWnd_GetLayoutEle")
	XWnd_SetCursorSys = XCDLL.MustFindProc("XWnd_SetCursorSys")
	XWnd_SetFont = XCDLL.MustFindProc("XWnd_SetFont")
	XWnd_SetID = XCDLL.MustFindProc("XWnd_SetID")
	XWnd_GetID = XCDLL.MustFindProc("XWnd_GetID")
	XWnd_SetLayoutSize = XCDLL.MustFindProc("XWnd_SetLayoutSize")
	XWnd_GetLayoutSize = XCDLL.MustFindProc("XWnd_GetLayoutSize")
	XWnd_SetDragBorderSize = XCDLL.MustFindProc("XWnd_SetDragBorderSize")
	XWnd_GetDragBorderSize = XCDLL.MustFindProc("XWnd_GetDragBorderSize")
	XWnd_SetMinimumSize = XCDLL.MustFindProc("XWnd_SetMinimumSize")
	XWnd_HitChildEle = XCDLL.MustFindProc("XWnd_HitChildEle")
	XWnd_GetChildCount = XCDLL.MustFindProc("XWnd_GetChildCount")
	XWnd_GetChildByIndex = XCDLL.MustFindProc("XWnd_GetChildByIndex")
	XWnd_GetChildByID = XCDLL.MustFindProc("XWnd_GetChildByID")
	XWnd_GetEle = XCDLL.MustFindProc("XWnd_GetEle")
	XWnd_CloseWindow = XCDLL.MustFindProc("XWnd_CloseWindow")
	XWnd_BindLayoutObject = XCDLL.MustFindProc("XWnd_BindLayoutObject")
	XWnd_GetLayoutObject = XCDLL.MustFindProc("XWnd_GetLayoutObject")
	XWnd_AdjustLayout = XCDLL.MustFindProc("XWnd_AdjustLayout")
	XWnd_CreateCaret = XCDLL.MustFindProc("XWnd_CreateCaret")
	XWnd_SetCaretSize = XCDLL.MustFindProc("XWnd_SetCaretSize")
	XWnd_SetCaretPos = XCDLL.MustFindProc("XWnd_SetCaretPos")
	XWnd_SetCaretPosEx = XCDLL.MustFindProc("XWnd_SetCaretPosEx")
	XWnd_SetCaretColor = XCDLL.MustFindProc("XWnd_SetCaretColor")
	XWnd_ShowCaret = XCDLL.MustFindProc("XWnd_ShowCaret")
	XWnd_DestroyCaret = XCDLL.MustFindProc("XWnd_DestroyCaret")
	XWnd_GetCaretHELE = XCDLL.MustFindProc("XWnd_GetCaretHELE")
	XWnd_GetClientRect = XCDLL.MustFindProc("XWnd_GetClientRect")
	XWnd_GetBodyRect = XCDLL.MustFindProc("XWnd_GetBodyRect")
	XWnd_SetTimer = XCDLL.MustFindProc("XWnd_SetTimer")
	XWnd_KillTimer = XCDLL.MustFindProc("XWnd_KillTimer")
	XWnd_GetBkInfoManager = XCDLL.MustFindProc("XWnd_GetBkInfoManager")
	XWnd_SetTransparentType = XCDLL.MustFindProc("XWnd_SetTransparentType")
	XWnd_SetTransparentAlpha = XCDLL.MustFindProc("XWnd_SetTransparentAlpha")
	XWnd_SetTransparentColor = XCDLL.MustFindProc("XWnd_SetTransparentColor")
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 14:47:35
// @Function: XWndCreate
// @Description: 显示窗口.
// @Calls: XWnd_Create
// @Input: x 窗口左上角x坐标. y 窗口左上角y坐标. cx 窗口宽度. cy 窗口高度. pTitle 窗口标题.
//         hWndParent 父窗口. XCStyle GUI库窗口样式,样式请参见windowStyle.go宏定义 xc_window_style_.

// @Return: GUI库窗口资源句柄.
// *******************************************************************
func XWndCreate(x, y, cx, cy int, pTitle string, hWndParent HWND, XCStyle int) HWINDOW {
	ret, _, _ := XWnd_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		StringToUintPtr(pTitle),
		uintptr(hWndParent),
		uintptr(XCStyle))

	return HWINDOW(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 14:51:11
// @Function: XWndCreateEx
// @Description: 显示窗口.
// @Calls: XWnd_Create
// @Input: dwExStyle 窗口扩展样式. lpClassName 窗口类名. lpWindowName 窗口名. dwStyle 窗口样式
//         x 窗口左上角x坐标. y 窗口左上角y坐标. cx 窗口宽度. cy 窗口高度.
//         hWndParent 父窗口. XCStyle GUI库窗口样式,样式请参见windowStyle.go宏定义 xc_window_style_.
// @Return: GUI库窗口资源句柄.
// *******************************************************************
func XWndCreateEx(dwExStyle uint32, lpClassName, lpWindowName string, dwStyle uint32, x, y, cx, cy int, hWndParent HWND, XCStyle int) HWINDOW {
	ret, _, _ := XWnd_CreateEx.Call(
		uintptr(dwExStyle),
		StringToUintPtr(lpClassName),
		StringToUintPtr(lpWindowName),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hWndParent),
		uintptr(XCStyle))

	return HWINDOW(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-6 18:45:49
// @Function: XWndRegEventC
// @Description: 注册事件函数C方式,事件函数省略自身HWINDOW句柄.
// @Calls: XEle_RegEventC
// @Input: hWindow 窗口句柄. nEvent 事件类型. pFun 事件函数.
// *******************************************************************
func XWndRegEventC(hWindow HWINDOW, nEvent int, pFun uintptr) {
	XWnd_RegEventC.Call(
		uintptr(hWindow),
		uintptr(nEvent),
		pFun)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 15:19:00
// @Function: XWndRegEventC1
// @Description: 注册事件函数C方式,事件函数不省略自身HWINDOW句柄.
// @Calls: XWnd_RegEventC1
// @Input: hWindow 窗口句柄. nEvent 事件类型. pFun 事件函数.
// @Return:
// *******************************************************************
func XWndRegEventC1(hWindow HWINDOW, nEvent int, pFun uintptr) {
	XWnd_RegEventC1.Call(
		uintptr(hWindow),
		uintptr(nEvent),
		pFun)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 15:21:17
// @Function: XWndRemovegEvent
// @Description: 移除事件函数.
// @Calls: XWnd_RemovegEvent
// @Input: hWindow 窗口句柄. nEvent 事件类型. pFun 事件函数.
// @Return:
// *******************************************************************
func XWndRemovegEvent(hWindow HWINDOW, nEvent int, pFun uintptr) {
	XWnd_RemovegEvent.Call(
		uintptr(hWindow),
		uintptr(nEvent),
		pFun)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 15:23:50
// @Function: XWndAddEle
// @Description: 添加元素到窗口.
// @Calls: XWnd_AddEle
// @Input: hWindow 窗口句柄. hEle 要添加的元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************************************
func XWndAddEle(hWindow HWINDOW, hEle HELE) bool {
	ret, _, _ := XWnd_AddEle.Call(
		uintptr(hWindow),
		uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 15:27:45
// @Function: XWndInsertEle
// @Description: 插入元素到目标元素前面.
// @Calls: XWnd_InsertEle
// @Input: hWindow 窗口句柄. hChildEle 要添加的元素句柄. hDestEle 目标元素.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************************************
func XWndInsertEle(hWindow HWINDOW, hChildEle HELE, hDestEle HELE) bool {
	ret, _, _ := XWnd_InsertEle.Call(
		uintptr(hWindow),
		uintptr(hChildEle),
		uintptr(hDestEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 15:30:49
// @Function: XWndAddShape
// @Description: 添加形状对象到窗口.
// @Calls: XWnd_AddShape
// @Input: hWindow 窗口句柄. hShape 形状对象句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************************************
func XWndAddShape(hWindow HWINDOW, hShape HXCGUI) bool {
	ret, _, _ := XWnd_AddShape.Call(
		uintptr(hWindow),
		uintptr(hShape))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 15:33:57
// @Function: XWndRedrawWnd
// @Description: 重绘窗口.
// @Calls: XWnd_RedrawWnd
// @Input: hWindow 窗口资源句柄.
// @Return:
// *******************************************************************
func XWndRedrawWnd(hWindow HWINDOW) {
	XWnd_RedrawWnd.Call(uintptr(hWindow))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 15:36:11
// @Function: XWndRedrawWndRect
// @Description: 重绘窗口指定区域.
// @Calls: XWnd_RedrawWndRect
// @Input: hWindow 窗口资源句柄. pRect 需要重绘的区域坐标. bImmediately TRUE立即重绘,FALSE放入消息队列延迟重绘.
// @Return:
// *******************************************************************
func XWndRedrawWndRect(hWindow HWINDOW, pRect *RECT, bImmediately BOOL) {
	XWnd_RedrawWndRect.Call(
		uintptr(hWindow),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(bImmediately))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 15:42:48
// @Function: XWndSetFocusEle
// @Description: 设置焦点元素.
// @Calls: XWnd_SetFocusEle
// @Input: hWindow 窗口资源句柄. hFocusEle 将获得焦点的元素.
// @Return:
// *******************************************************************
func XWndSetFocusEle(hWindow HWINDOW, hFocusEle HELE) {
	XWnd_SetFocusEle.Call(
		uintptr(hWindow),
		uintptr(hFocusEle))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 15:46:22
// @Function: XWndGetFocusEle
// @Description: 获得拥有输入焦点的元素.
// @Calls: XWnd_GetFocusEle
// @Input: hWindow 窗口资源句柄.
// @Return: 元素句柄.
// *******************************************************************
func XWndGetFocusEle(hWindow HWINDOW) HELE {
	ret, _, _ := XWnd_GetFocusEle.Call(uintptr(hWindow))

	return HELE(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 15:48:38
// @Function: XWndSetCursor
// @Description: 设置窗口鼠标光标.
// @Calls: XWnd_SetCursor
// @Input: hWindow 窗口句柄. hCursor 鼠标光标句柄.
// @Return:
// *******************************************************************
func XWndSetCursor(hWindow HWINDOW, hCursor HCURSOR) {
	XWnd_SetCursor.Call(
		uintptr(hWindow),
		uintptr(hCursor))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 16:08:13
// @Function: XWndGetCursor
// @Description: 获取窗口的鼠标光标.
// @Calls: XWnd_GetCursor
// @Input: hWindow 窗口句柄.
// @Return: 鼠标光标句柄.
// *******************************************************************
func XWndGetCursor(hWindow HWINDOW) HSTRING {
	ret, _, _ := XWnd_GetCursor.Call(uintptr(hWindow))

	return HSTRING(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 16:09:58
// @Function: XWndGetHWND
// @Description: 获取HWND句柄.
// @Calls: XWnd_GetHWND
// @Input: hWindow 窗口句柄.
// @Return: HWND句柄.
// *******************************************************************
func XWndGetHWND(hWindow HWINDOW) HWND {
	ret, _, _ := XWnd_GetHWND.Call(uintptr(hWindow))

	return HWND(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 16:12:11
// @Function: XWndEnableDragBorder
// @Description: 启用拖动窗口边框.
// @Calls: XWnd_EnableDragBorder
// @Input: hWindow 窗口句柄. bEnable 是否启用.
// @Return:
// *******************************************************************
func XWndEnableDragBorder(hWindow HWINDOW, bEnable BOOL) {
	XWnd_EnableDragBorder.Call(
		uintptr(hWindow),
		uintptr(bEnable))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 16:17:49
// @Function: XWndEnableDragWindow
// @Description: 启用拖动窗口.
// @Calls: XWnd_EnableDragWindow
// @Input: hWindow 窗口句柄. bEnable 是否启用.
// @Return:
// *******************************************************************
func XWndEnableDragWindow(hWindow HWINDOW, bEnable BOOL) {
	XWnd_EnableDragWindow.Call(
		uintptr(hWindow),
		uintptr(bEnable))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 16:24:37
// @Function: XWndEnableDrawBk
// @Description: 是否绘制窗口背景.
// @Calls: XWnd_EnableDrawBk
// @Input: hWindow 窗口句柄. bEnable 是否启用.
// @Return:
// *******************************************************************
func XWndEnableDrawBk(hWindow HWINDOW, bEnable BOOL) {
	XWnd_EnableDrawBk.Call(
		uintptr(hWindow),
		uintptr(bEnable))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 16:26:34
// @Function: XWndEnableAutoFocus
// @Description: 当鼠标左键按下是否获得焦点.
// @Calls: XWnd_EnableAutoFocus
// @Input: hWindow 窗口句柄. bEnable 是否启用.
// @Return:
// *******************************************************************
func XWndEnableAutoFocus(hWindow HWINDOW, bEnable BOOL) {
	XWnd_EnableAutoFocus.Call(
		uintptr(hWindow),
		uintptr(bEnable))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 16:30:02
// @Function: XWndSetCaptureEle
// @Description: 设置鼠标捕获元素.
// @Calls: XWnd_SetCaptureEle
// @Input: hWindow 窗口句柄. hEle 元素句柄.
// @Return:
// *******************************************************************
func XWndSetCaptureEle(hWindow HWINDOW, hEle HELE) {
	XWnd_SetCaptureEle.Call(
		uintptr(hWindow),
		uintptr(hEle))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 16:32:19
// @Function: XWndGetCaptureEle
// @Description: 获取当前鼠标捕获元素.
// @Calls: XWnd_GetCaptureEle
// @Input: hWindow 窗口句柄.
// @Return: 元素句柄.
// *******************************************************************
func XWndGetCaptureEle(hWindow HWINDOW) HELE {
	ret, _, _ := XWnd_GetCaptureEle.Call(uintptr(hWindow))

	return HELE(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 16:34:21
// @Function: XWndGetDrawRect
// @Description: 获取重绘区域.
// @Calls: XWnd_GetDrawRect
// @Input: hWindow 窗口句柄. pRcPaint 重绘区域坐标.
// @Return:
// *******************************************************************
func XWndGetDrawRect(hWindow HWINDOW, pRcPaint *RECT) {
	XWnd_GetDrawRect.Call(
		uintptr(hWindow),
		uintptr(unsafe.Pointer(pRcPaint)))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-4 23:46:41
// @Function: XWndShowWindow
// @Description: 显示窗口.
// @Calls: XWnd_ShowWindow
// @Input: hWindow 窗口句柄. nCmdShow XWndShowWindow constants.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************************************
func XWndShowWindow(hWindow HWINDOW, nCmdShow int) bool {
	ret, _, _ := XWnd_ShowWindow.Call(
		uintptr(hWindow),
		uintptr(nCmdShow))
	// MSDN上返回值：true 为 0，false 为 1
	return ret != 1
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 16:47:05
// @Function: XWndBindLayoutEle
// @Description: 设置布局元素
// @Calls: XWnd_BindLayoutEle
// @Input: hWindow 窗口句柄. nPosition 参见宏定义. hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************************************
func XWndBindLayoutEle(hWindow HWINDOW, nPosition window_position_, hEle HELE) bool {
	ret, _, _ := XWnd_BindLayoutEle.Call(
		uintptr(hWindow),
		uintptr(nPosition),
		uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 16:52:45
// @Function: XWndGetLayoutEle
// @Description: 获取布局元素.
// @Calls: XWnd_GetLayoutEle
// @Input: hWindow 窗口句柄. nPosition 参见宏定义.
// @Return: 元素句柄.
// *******************************************************************
func XWndGetLayoutEle(hWindow HWINDOW, nPosition window_position_) HELE {
	ret, _, _ := XWnd_GetLayoutEle.Call(
		uintptr(hWindow),
		uintptr(nPosition))

	return HELE(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 16:55:08
// @Function: XWndSetCursorSys
// @Description: 系统函数,设置窗口类光标句柄.
// @Calls: XWnd_SetCursorSys
// @Input: hWindow 窗口句柄. hCursor 光标句柄.
// @Return: 返回先前的光标句柄.
// *******************************************************************
func XWndSetCursorSys(hWindow HWINDOW, hCursor HCURSOR) HCURSOR {
	ret, _, _ := XWnd_SetCursorSys.Call(
		uintptr(hWindow),
		uintptr(hCursor))

	return HCURSOR(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 16:58:42
// @Function: XWndSetFont
// @Description: 设置窗口字体.
// @Calls: XWnd_SetFont
// @Input: hWindow 窗口句柄. hFontx 炫彩字体句柄.
// @Return:
// *******************************************************************
func XWndSetFont(hWindow HWINDOW, hFontx HFONTX) {
	XWnd_SetFont.Call(
		uintptr(hWindow),
		uintptr(hFontx))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:07:26
// @Function: XWndSetID
// @Description: 设置窗口ID.
// @Calls: XWnd_SetID
// @Input: hWindow 窗口句柄. nID ID值.
// @Return:
// *******************************************************************
func XWndSetID(hWindow HWINDOW, nID int) {
	XWnd_SetID.Call(
		uintptr(hWindow),
		uintptr(nID))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:24:30
// @Function: XWndGetID
// @Description: 获取窗口ID.
// @Calls: XWnd_GetID
// @Input: hWindow 窗口句柄.
// @Return: 返回窗口ID值.
// *******************************************************************
func XWndGetID(hWindow HWINDOW) int {
	ret, _, _ := XWnd_GetID.Call(uintptr(hWindow))

	return int(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:26:50
// @Function: XWndSetLayoutSize
// @Description: 设置布局大小.
// @Calls: XWnd_SetLayoutSize
// @Input: hWindow 窗口句柄. left 窗口左边大小. top 窗口上边大小. right 窗口右边大小. bottom 窗口底部大小.
// @Return:
// *******************************************************************
func XWndSetLayoutSize(hWindow HWINDOW, left, top, right, bottom int) {
	XWnd_SetLayoutSize.Call(
		uintptr(hWindow),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:29:10
// @Function: XWndGetLayoutSize
// @Description: 获取布局大小.
// @Calls: XWnd_GetLayoutSize
// @Input: hWindow 窗口句柄. pBorderSize 布局大小.
// @Return:
// *******************************************************************
func XWndGetLayoutSize(hWindow HWINDOW, pBorderSize *BorderSize_) {
	XWnd_GetLayoutSize.Call(
		uintptr(hWindow),
		uintptr(unsafe.Pointer(pBorderSize)))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:33:53
// @Function: XWndSetDragBorderSize
// @Description: 设置窗口拖动边框大小.
// @Calls: XWnd_SetDragBorderSize
// @Input: hWindow 窗口句柄. left 窗口左边大小. top 窗口上边大小. right 窗口右边大小. bottom 窗口底边大小.
// @Return:
// *******************************************************************
func XWndSetDragBorderSize(hWindow HWINDOW, left, top, right, bottom int) {
	XWnd_SetDragBorderSize.Call(
		uintptr(hWindow),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:36:06
// @Function: XWndGetDragBorderSize
// @Description: 获取窗口拖动边框大小.
// @Calls: XWnd_GetDragBorderSize
// @Input: hWindow 窗口句柄. pSize 拖动边框大小.
// @Return:
// *******************************************************************
func XWndGetDragBorderSize(hWindow HWINDOW, pSize *BorderSize_) {
	XWnd_GetDragBorderSize.Call(
		uintptr(hWindow),
		uintptr(unsafe.Pointer(pSize)))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:38:07
// @Function: XWndSetMinimumSize
// @Description: 设置窗口的最小宽度和高度.
// @Calls: XWnd_SetMinimumSize
// @Input: hWindow 窗口句柄. width 最小宽度. height 最小高度.
// @Return:
// *******************************************************************
func XWndSetMinimumSize(hWindow HWINDOW, width, height int) {
	XWnd_SetMinimumSize.Call(
		uintptr(hWindow),
		uintptr(width),
		uintptr(height))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:43:32
// @Function: XWndHitChildEle
// @Description: 检测所在元素.
// @Calls: XWnd_HitChildEle
// @Input: hWindow 窗口句柄. pPt 左边点.
// @Return: 元素句柄.
// *******************************************************************
func XWndHitChildEle(hWindow HWINDOW, pPt *POINT) HELE {
	ret, _, _ := XWnd_HitChildEle.Call(
		uintptr(hWindow),
		uintptr(unsafe.Pointer(pPt)))

	return HELE(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:45:30
// @Function: XWndGetChildCount
// @Description: 获取当前层子元素数量,不包含子元素的子元素.
// @Calls: XWnd_GetChildCount
// @Input: hWindow 窗口句柄.
// @Return: 子元素数量.
// *******************************************************************
func XWndGetChildCount(hWindow HWINDOW) int {
	ret, _, _ := XWnd_GetChildCount.Call(uintptr(hWindow))

	return int(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:47:01
// @Function: XWndGetChildByIndex
// @Description: 获取当前层子元素通过索引.
// @Calls: XWnd_GetChildByIndex
// @Input: hWindow 窗口句柄. index 元素索引.
// @Return: 元素句柄.
// *******************************************************************
func XWndGetChildByIndex(hWindow HWINDOW, index int) HELE {
	ret, _, _ := XWnd_GetChildByIndex.Call(
		uintptr(hWindow),
		uintptr(index))

	return HELE(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 18:43:06
// @Function: XWndGetChildByID
// @Description: 获取当前层子元素通过元素ID.
// @Calls: XWnd_GetChildByID
// @Input: hWindow 窗口句柄. nID 元素ID,ID必须大于0.
// @Return: 元素句柄.
// *******************************************************************
func XWndGetChildByID(hWindow HWINDOW, nID int) HELE {
	ret, _, _ := XWnd_GetChildByID.Call(
		uintptr(hWindow),
		uintptr(nID))

	return HELE(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:49:18
// @Function: XWndGetEle
// @Description: 获取元素通过元素ID,如果元素不在该窗口上无效.
// @Calls: XWnd_GetEle
// @Input: hWindow 窗口句柄. nID 元素ID,ID必须大于0.
// @Return: 元素句柄.
// *******************************************************************
func XWndGetEle(hWindow HWINDOW, nID int) HELE {
	ret, _, _ := XWnd_GetEle.Call(
		uintptr(hWindow),
		uintptr(nID))

	return HELE(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:50:47
// @Function: XWndCloseWindow
// @Description: 关闭窗口.
// @Calls: XWnd_CloseWindow
// @Input: hWindow 窗口句柄.
// @Return:
// *******************************************************************
func XWndCloseWindow(hWindow HWINDOW) {
	XWnd_CloseWindow.Call(uintptr(hWindow))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:52:06
// @Function: XWndBindLayoutObject
// @Description: 绑定布局对象.
// @Calls: XWnd_BindLayoutObject
// @Input: hWindow 窗口句柄. nPosition 参见宏定义. hLayout 布局对象.
// @Return:
// *******************************************************************
func XWndBindLayoutObject(hWindow HWINDOW, nPosition window_position_, hLayout HXCGUI) {
	XWnd_BindLayoutObject.Call(
		uintptr(hWindow),
		uintptr(nPosition),
		uintptr(hLayout))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:54:09
// @Function: XWndGetLayoutObject
// @Description: 获取布局对象.
// @Calls: XWnd_GetLayoutObject
// @Input: hWindow 窗口句柄. nPosition 参见宏定义.
// @Return: 布局对象句柄.
// *******************************************************************
func XWndGetLayoutObject(hWindow HWINDOW, nPosition window_position_) HXCGUI {
	ret, _, _ := XWnd_GetLayoutObject.Call(
		uintptr(hWindow),
		uintptr(nPosition))

	return HXCGUI(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:55:48
// @Function: XWndAdjustLayout
// @Description: 调整窗口布局.
// @Calls: XWnd_AdjustLayout
// @Input: hWindow 窗口句柄.
// @Return:
// *******************************************************************
func XWndAdjustLayout(hWindow HWINDOW) {
	XWnd_AdjustLayout.Call(uintptr(hWindow))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:57:12
// @Function: XWndCreateCaret
// @Description: 创建插入符.
// @Calls: XWnd_CreateCaret
// @Input: hWindow 窗口句柄. hEle 元素句柄. width 宽度. height 高度.
// @Return:
// *******************************************************************
func XWndCreateCaret(hWindow HWINDOW, hEle HELE, width, height int) {
	XWnd_CreateCaret.Call(
		uintptr(hWindow),
		uintptr(hEle),
		uintptr(width),
		uintptr(height))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 17:59:12
// @Function: XWndSetCaretSize
// @Description: 设置插入符大小.
// @Calls: XWnd_SetCaretSize
// @Input: hWindow 窗口句柄. width 宽度. height 高度.
// @Return:
// *******************************************************************
func XWndSetCaretSize(hWindow HWINDOW, width, height int) {
	XWnd_SetCaretSize.Call(
		uintptr(hWindow),
		uintptr(width),
		uintptr(height))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 18:00:17
// @Function: XWndSetCaretPos
// @Description: 设置插入符位置.
// @Calls: XWnd_SetCaretPos
// @Input: hWindow 窗口句柄. x x坐标. y y坐标.
// @Return:
// *******************************************************************
func XWndSetCaretPos(hWindow HWINDOW, x, y int) {
	XWnd_SetCaretPos.Call(
		uintptr(hWindow),
		uintptr(x),
		uintptr(y))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 18:01:50
// @Function: XWndSetCaretPosEx
// @Description: 设置插入符位置.
// @Calls: XWnd_SetCaretPosEx
// @Input: hWindow 窗口句柄. x x坐标. y y坐标. width 宽度. height 高度.
// @Return:
// *******************************************************************
func XWndSetCaretPosEx(hWindow HWINDOW, x, y, width, height int) {
	XWnd_SetCaretPosEx.Call(
		uintptr(hWindow),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 18:03:15
// @Function: XWndSetCaretColor
// @Description: 设置插入符颜色.
// @Calls:  XWnd_SetCaretColor
// @Input: hWindow 窗口句柄. color 颜色值.
// @Return:
// *******************************************************************
func XWndSetCaretColor(hWindow HWINDOW, color COLORREF) {
	XWnd_SetCaretColor.Call(
		uintptr(hWindow),
		uintptr(color))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 18:04:35
// @Function: XWndShowCaret
// @Description: 显示插入符.
// @Calls:  XWnd_ShowCaret
// @Input: hWindow 窗口句柄. bShow 是否显示.
// @Return:
// *******************************************************************
func XWndShowCaret(hWindow HWINDOW, bShow BOOL) {
	XWnd_ShowCaret.Call(
		uintptr(hWindow),
		uintptr(bShow))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 18:05:38
// @Function: XWndDestroyCaret
// @Description: 销毁插入符.
// @Calls:  XWnd_DestroyCaret
// @Input: hWindow 窗口句柄.
// @Return:
// *******************************************************************
func XWndDestroyCaret(hWindow HWINDOW) {
	XWnd_DestroyCaret.Call(uintptr(hWindow))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 18:06:46
// @Function: XWndGetCaretHELE
// @Description: 获取插入符所属元素.
// @Calls:  XWnd_GetCaretHELE
// @Input: hWindow 窗口句柄.
// @Return: 元素句柄.
// *******************************************************************
func XWndGetCaretHELE(hWindow HWINDOW) HELE {
	ret, _, _ := XWnd_GetCaretHELE.Call(uintptr(hWindow))

	return HELE(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 18:07:53
// @Function: XWndGetClientRect
// @Description: 获取窗口客户区坐标.
// @Calls:  XWnd_GetClientRect
// @Input: hWindow 窗口句柄. pRect 坐标.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************************************
func XWndGetClientRect(hWindow HWINDOW, pRect *RECT) bool {
	ret, _, _ := XWnd_GetClientRect.Call(
		uintptr(hWindow),
		uintptr(unsafe.Pointer(pRect)))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 18:09:55
// @Function: XWndGetBodyRect
// @Description: 获取body坐标.
// @Calls:  XWnd_GetBodyRect
// @Input: hWindow 窗口句柄. pRect 坐标.
// @Return:
// *******************************************************************
func XWndGetBodyRect(hWindow HWINDOW, pRect *RECT) {
	XWnd_GetBodyRect.Call(
		uintptr(hWindow),
		uintptr(unsafe.Pointer(pRect)))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 18:11:17
// @Function: XWndSetTimer
// @Description: 设置窗口定时器.
// @Calls:  XWnd_SetTimer
// @Input: hWindow 窗口句柄. nIDEvent 定时器ID. uElapse 间隔值,单位毫秒.
// @Return: 参加MSDN.
// *******************************************************************
func XWndSetTimer(hWindow HWINDOW, nIDEvent, uElapse uint) uint {
	ret, _, _ := XWnd_SetTimer.Call(
		uintptr(hWindow),
		uintptr(nIDEvent),
		uintptr(uElapse))

	return uint(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 18:15:00
// @Function: XWndKillTimer
// @Description: 关闭定时器.
// @Calls:  XWnd_KillTimer
// @Input: hWindow 窗口句柄. nIDEvent 定时器ID.
// @Return: 参见MSDN.
// *******************************************************************
func XWndKillTimer(hWindow HWINDOW, nIDEvent uint) bool {
	ret, _, _ := XWnd_KillTimer.Call(
		uintptr(hWindow),
		uintptr(nIDEvent))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 18:16:31
// @Function: XWndGetBkInfoManager
// @Description: 获取背景内容管理器.
// @Calls:  XWnd_GetBkInfoManager
// @Input: hWindow 窗口句柄. nPosition 窗口布局位置.
// @Return: 背景内容管理器.
// *******************************************************************
func XWndGetBkInfoManager(hWindow HWINDOW, nPosition window_position_) HBKINFOM {
	ret, _, _ := XWnd_GetBkInfoManager.Call(
		uintptr(hWindow),
		uintptr(nPosition))

	return HBKINFOM(ret)
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 18:20:39
// @Function: XWndSetTransparentType
// @Description: 设置透明窗口,同时可以通过该函数关闭透明窗口.
// @Calls:  XWnd_SetTransparentType
// @Input: hWindow 窗口句柄. nType 窗口透明类型.
// @Return:
// *******************************************************************
func XWndSetTransparentType(hWindow HWINDOW, nType window_transparent_) {
	XWnd_SetTransparentType.Call(
		uintptr(hWindow),
		uintptr(nType))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 18:22:20
// @Function: XWndSetTransparentAlpha
// @Description: 设置透明窗口的透明度,设置后调用重绘窗口API来更新.
// @Calls:  XWnd_SetTransparentAlpha
// @Input: hWindow 窗口句柄. alpha 窗口透明度,范围0-255之间,0透明,255不透明.
// @Return:
// *******************************************************************
func XWndSetTransparentAlpha(hWindow HWINDOW, alpha byte) {
	XWnd_SetTransparentAlpha.Call(
		uintptr(hWindow),
		uintptr(alpha))
}

// *******************************************************************
// @Author: cody.guo
// @Date: 2015-11-8 18:34:05
// @Function: XWndSetTransparentColor
// @Description: 设置透明窗口的透明色.
// @Calls:  XWnd_SetTransparentColor
// @Input: hWindow 窗口句柄. color 窗口透明色.
// @Return:
// *******************************************************************
func XWndSetTransparentColor(hWindow HWINDOW, color COLORREF) {
	XWnd_SetTransparentColor.Call(
		uintptr(hWindow),
		uintptr(color))
}
