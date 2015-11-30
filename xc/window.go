package xc

import (
	"syscall"
	"unsafe"
	// "fmt"
)

var (
	// Functions
	xWnd_Create               *syscall.Proc
	xWnd_CreateEx             *syscall.Proc
	xWnd_RegEventC            *syscall.Proc
	xWnd_RegEventC1           *syscall.Proc
	xWnd_RemovegEvent         *syscall.Proc
	xWnd_AddEle               *syscall.Proc
	xWnd_InsertEle            *syscall.Proc
	xWnd_AddShape             *syscall.Proc
	xWnd_RedrawWnd            *syscall.Proc
	xWnd_RedrawWndRect        *syscall.Proc
	xWnd_SetFocusEle          *syscall.Proc
	xWnd_GetFocusEle          *syscall.Proc
	xWnd_SetCursor            *syscall.Proc
	xWnd_GetCursor            *syscall.Proc
	xWnd_GetHWND              *syscall.Proc
	xWnd_EnableDragBorder     *syscall.Proc
	xWnd_EnableDragWindow     *syscall.Proc
	xWnd_EnableDrawBk         *syscall.Proc
	xWnd_EnableAutoFocus      *syscall.Proc
	xWnd_EnableMaxWindow      *syscall.Proc
	xWnd_SetCaptureEle        *syscall.Proc
	xWnd_GetCaptureEle        *syscall.Proc
	xWnd_GetDrawRect          *syscall.Proc
	xWnd_ShowWindow           *syscall.Proc
	xWnd_BindLayoutEle        *syscall.Proc
	xWnd_GetLayoutEle         *syscall.Proc
	xWnd_SetCursorSys         *syscall.Proc
	xWnd_SetFont              *syscall.Proc
	xWnd_SetID                *syscall.Proc
	xWnd_GetID                *syscall.Proc
	xWnd_SetLayoutSize        *syscall.Proc
	xWnd_GetLayoutSize        *syscall.Proc
	xWnd_SetDragBorderSize    *syscall.Proc
	xWnd_GetDragBorderSize    *syscall.Proc
	xWnd_SetMinimumSize       *syscall.Proc
	xWnd_HitChildEle          *syscall.Proc
	xWnd_GetChildCount        *syscall.Proc
	xWnd_GetChildByIndex      *syscall.Proc
	xWnd_GetChildByID         *syscall.Proc
	xWnd_GetEle               *syscall.Proc
	xWnd_GetChildShapeCount   *syscall.Proc
	xWnd_GetChildShapeByIndex *syscall.Proc
	xWnd_CloseWindow          *syscall.Proc
	xWnd_BindLayoutObject     *syscall.Proc
	xWnd_GetLayoutObject      *syscall.Proc
	xWnd_AdjustLayout         *syscall.Proc
	xWnd_AdjustLayoutObject   *syscall.Proc
	xWnd_CreateCaret          *syscall.Proc
	xWnd_SetCaretSize         *syscall.Proc
	xWnd_SetCaretPos          *syscall.Proc
	xWnd_SetCaretPosEx        *syscall.Proc
	xWnd_SetCaretColor        *syscall.Proc
	xWnd_ShowCaret            *syscall.Proc
	xWnd_DestroyCaret         *syscall.Proc
	xWnd_GetCaretHELE         *syscall.Proc
	xWnd_GetClientRect        *syscall.Proc
	xWnd_GetBodyRect          *syscall.Proc
	xWnd_SetTimer             *syscall.Proc
	xWnd_KillTimer            *syscall.Proc
	xWnd_GetBkInfoManager     *syscall.Proc
	xWnd_SetTransparentType   *syscall.Proc
	xWnd_SetTransparentAlpha  *syscall.Proc
	xWnd_SetTransparentColor  *syscall.Proc
)

func init() {
	// Functions
	xWnd_Create = xcDLL.MustFindProc("XWnd_Create")
	xWnd_CreateEx = xcDLL.MustFindProc("XWnd_CreateEx")
	xWnd_RegEventC = xcDLL.MustFindProc("XWnd_RegEventC")
	xWnd_RegEventC1 = xcDLL.MustFindProc("XWnd_RegEventC1")
	xWnd_RemovegEvent = xcDLL.MustFindProc("XWnd_RemovegEvent")
	xWnd_AddEle = xcDLL.MustFindProc("XWnd_AddEle")
	xWnd_InsertEle = xcDLL.MustFindProc("XWnd_InsertEle")
	xWnd_AddShape = xcDLL.MustFindProc("XWnd_AddShape")
	xWnd_RedrawWnd = xcDLL.MustFindProc("XWnd_RedrawWnd")
	xWnd_RedrawWndRect = xcDLL.MustFindProc("XWnd_RedrawWndRect")
	xWnd_SetFocusEle = xcDLL.MustFindProc("XWnd_SetFocusEle")
	xWnd_GetFocusEle = xcDLL.MustFindProc("XWnd_GetFocusEle")
	xWnd_SetCursor = xcDLL.MustFindProc("XWnd_SetCursor")
	xWnd_GetCursor = xcDLL.MustFindProc("XWnd_GetCursor")
	xWnd_GetHWND = xcDLL.MustFindProc("XWnd_GetHWND")
	xWnd_EnableDragBorder = xcDLL.MustFindProc("XWnd_EnableDragBorder")
	xWnd_EnableDragWindow = xcDLL.MustFindProc("XWnd_EnableDragWindow")
	xWnd_EnableDrawBk = xcDLL.MustFindProc("XWnd_EnableDrawBk")
	xWnd_EnableAutoFocus = xcDLL.MustFindProc("XWnd_EnableAutoFocus")
	xWnd_EnableMaxWindow = xcDLL.MustFindProc("XWnd_EnableMaxWindow")
	xWnd_SetCaptureEle = xcDLL.MustFindProc("XWnd_SetCaptureEle")
	xWnd_GetCaptureEle = xcDLL.MustFindProc("XWnd_GetCaptureEle")
	xWnd_GetDrawRect = xcDLL.MustFindProc("XWnd_GetDrawRect")
	xWnd_ShowWindow = xcDLL.MustFindProc("XWnd_ShowWindow")
	xWnd_BindLayoutEle = xcDLL.MustFindProc("XWnd_BindLayoutEle")
	xWnd_GetLayoutEle = xcDLL.MustFindProc("XWnd_GetLayoutEle")
	xWnd_SetCursorSys = xcDLL.MustFindProc("XWnd_SetCursorSys")
	xWnd_SetFont = xcDLL.MustFindProc("XWnd_SetFont")
	xWnd_SetID = xcDLL.MustFindProc("XWnd_SetID")
	xWnd_GetID = xcDLL.MustFindProc("XWnd_GetID")
	xWnd_SetLayoutSize = xcDLL.MustFindProc("XWnd_SetLayoutSize")
	xWnd_GetLayoutSize = xcDLL.MustFindProc("XWnd_GetLayoutSize")
	xWnd_SetDragBorderSize = xcDLL.MustFindProc("XWnd_SetDragBorderSize")
	xWnd_GetDragBorderSize = xcDLL.MustFindProc("XWnd_GetDragBorderSize")
	xWnd_SetMinimumSize = xcDLL.MustFindProc("XWnd_SetMinimumSize")
	xWnd_HitChildEle = xcDLL.MustFindProc("XWnd_HitChildEle")
	xWnd_GetChildCount = xcDLL.MustFindProc("XWnd_GetChildCount")
	xWnd_GetChildByIndex = xcDLL.MustFindProc("XWnd_GetChildByIndex")
	xWnd_GetChildByID = xcDLL.MustFindProc("XWnd_GetChildByID")
	xWnd_GetEle = xcDLL.MustFindProc("XWnd_GetEle")
	xWnd_GetChildShapeCount = xcDLL.MustFindProc("XWnd_GetChildShapeCount")
	xWnd_GetChildShapeByIndex = xcDLL.MustFindProc("XWnd_GetChildShapeByIndex")
	xWnd_CloseWindow = xcDLL.MustFindProc("XWnd_CloseWindow")
	xWnd_BindLayoutObject = xcDLL.MustFindProc("XWnd_BindLayoutObject")
	xWnd_GetLayoutObject = xcDLL.MustFindProc("XWnd_GetLayoutObject")
	xWnd_AdjustLayout = xcDLL.MustFindProc("XWnd_AdjustLayout")
	xWnd_AdjustLayoutObject = xcDLL.MustFindProc("XWnd_AdjustLayoutObject")
	xWnd_CreateCaret = xcDLL.MustFindProc("XWnd_CreateCaret")
	xWnd_SetCaretSize = xcDLL.MustFindProc("XWnd_SetCaretSize")
	xWnd_SetCaretPos = xcDLL.MustFindProc("XWnd_SetCaretPos")
	xWnd_SetCaretPosEx = xcDLL.MustFindProc("XWnd_SetCaretPosEx")
	xWnd_SetCaretColor = xcDLL.MustFindProc("XWnd_SetCaretColor")
	xWnd_ShowCaret = xcDLL.MustFindProc("XWnd_ShowCaret")
	xWnd_DestroyCaret = xcDLL.MustFindProc("XWnd_DestroyCaret")
	xWnd_GetCaretHELE = xcDLL.MustFindProc("XWnd_GetCaretHELE")
	xWnd_GetClientRect = xcDLL.MustFindProc("XWnd_GetClientRect")
	xWnd_GetBodyRect = xcDLL.MustFindProc("XWnd_GetBodyRect")
	xWnd_SetTimer = xcDLL.MustFindProc("XWnd_SetTimer")
	xWnd_KillTimer = xcDLL.MustFindProc("XWnd_KillTimer")
	xWnd_GetBkInfoManager = xcDLL.MustFindProc("XWnd_GetBkInfoManager")
	xWnd_SetTransparentType = xcDLL.MustFindProc("XWnd_SetTransparentType")
	xWnd_SetTransparentAlpha = xcDLL.MustFindProc("XWnd_SetTransparentAlpha")
	xWnd_SetTransparentColor = xcDLL.MustFindProc("XWnd_SetTransparentColor")
}

/*
创建窗口

参数:
	x 窗口左上角x坐标.
	y 窗口左上角y坐标.
	cx 窗口宽度.
	cy 窗口高度.
	pTitle 窗口标题.
	hWndParent 父窗口.
	XCStyle GUI库窗口样式,样式请参见宏定义 xc_window_style_.
返回:
	GUI库窗口资源句柄.
*/
func XWndCreate(x, y, cx, cy int, pTitle *uint16, hWndParent HWND, XCStyle int) HWINDOW {
	ret, _, _ := xWnd_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(unsafe.Pointer(pTitle)),
		uintptr(hWndParent),
		uintptr(XCStyle))

	return HWINDOW(ret)
}

/*
创建窗口,增强功能.

参数:
	dwExStyle 窗口扩展样式.
	lpClassName 窗口类名.
	lpWindowName 窗口名.
	dwStyle 窗口样式
	x 窗口左上角x坐标.
	y 窗口左上角y坐标.
	cx 窗口宽度.
	cy 窗口高度.
	hWndParent 父窗口.
	XCStyle GUI库窗口样式,样式请参见宏定义 xc_window_style_.
返回:
	GUI库窗口资源句柄.
*/
func XWndCreateEx(dwExStyle uint32, lpClassName, lpWindowName *uint16, dwStyle uint32, x, y, cx, cy int, hWndParent HWND, XCStyle int) HWINDOW {
	ret, _, _ := xWnd_CreateEx.Call(
		uintptr(dwExStyle),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hWndParent),
		uintptr(XCStyle))

	return HWINDOW(ret)
}

/*
注册事件函数C方式,事件函数省略自身HWINDOW句柄.

参数:
	hWindow 窗口句柄.
	nEvent 事件类型.
	pFun 事件函数.
*/
func XWndRegEventC(hWindow HWINDOW, nEvent int, pFun uintptr) {
	xWnd_RegEventC.Call(
		uintptr(hWindow),
		uintptr(nEvent),
		pFun)
}

/*
注册事件函数C方式,事件函数不省略自身HWINDOW句柄.

参数:
	hWindow 窗口句柄.
	nEvent 事件类型.
	pFun 事件函数.
*/
func XWndRegEventC1(hWindow HWINDOW, nEvent int, pFun uintptr) {
	xWnd_RegEventC1.Call(
		uintptr(hWindow),
		uintptr(nEvent),
		pFun)
}

/*
移除事件函数.

参数:
	hWindow 窗口句柄.
	nEvent 事件类型.
	pFun 事件函数.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XWndRemovegEvent(hWindow HWINDOW, nEvent int, pFun uintptr) bool {
	ret, _, _ := xWnd_RemovegEvent.Call(
		uintptr(hWindow),
		uintptr(nEvent),
		pFun)

	return ret == TRUE
}

/*
添加元素到窗口

参数:
	hWindow 窗口句柄
	hEle 要添加的元素句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XWndAddEle(hWindow HWINDOW, hEle HELE) bool {
	ret, _, _ := xWnd_AddEle.Call(
		uintptr(hWindow),
		uintptr(hEle))

	return ret == TRUE
}

/*
插入元素到目标元素前面.

参数:
	hWindow 窗口句柄.
	hChildEle 要添加的元素句柄.
	hDestEle 目标元素.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XWndInsertEle(hWindow HWINDOW, hChildEle, hDestEle HELE) bool {
	ret, _, _ := xWnd_InsertEle.Call(
		uintptr(hWindow),
		uintptr(hChildEle),
		uintptr(hDestEle))

	return ret == TRUE
}

/*
添加形状对象到窗口.

参数:
	hWindow 窗口句柄.
	hShape 形状对象句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XWndAddShape(hWindow HWINDOW, hShape HXCGUI) bool {
	ret, _, _ := xWnd_AddShape.Call(
		uintptr(hWindow),
		uintptr(hShape))

	return ret == TRUE
}

/*
重绘窗口.

参数:
	hWindow 窗口资源句柄.
*/
func XWndRedrawWnd(hWindow HWINDOW) {
	xWnd_RedrawWnd.Call(uintptr(hWindow))
}

/*
重绘窗口指定区域.

参数:
	hWindow 窗口资源句柄.
	pRect 需要重绘的区域坐标.
	bImmediately TRUE立即重绘,FALSE放入消息队列延迟重绘.
*/
func XWndRedrawWndRect(hWindow HWINDOW, pRect *RECT, bImmediately bool) {
	xWnd_RedrawWndRect.Call(
		uintptr(hWindow),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(BoolToBOOL(bImmediately)))
}

/*
设置焦点元素.

参数:
	hWindow 窗口资源句柄.
	hFocusEle 将获得焦点的元素.
*/
func XWndSetFocusEle(hWindow HWINDOW, hFocusEle HELE) {
	xWnd_SetFocusEle.Call(
		uintptr(hWindow),
		uintptr(hFocusEle))
}

/*
获得拥有输入焦点的元素.

参数:
	hWindow 窗口资源句柄.
返回:
	元素句柄.
*/
func XWndGetFocusEle(hWindow HWINDOW) HELE {
	ret, _, _ := xWnd_GetFocusEle.Call(uintptr(hWindow))

	return HELE(ret)
}

/*
设置窗口鼠标光标.

参数:
	hWindow 窗口句柄.
	hCursor 鼠标光标句柄.
*/
func XWndSetCursor(hWindow HWINDOW, hCursor HCURSOR) {
	xWnd_SetCursor.Call(
		uintptr(hWindow),
		uintptr(hCursor))
}

/*
获取窗口的鼠标光标.

参数:
	hWindow 窗口句柄.
返回:
	鼠标光标句柄.
*/
func XWndGetCursor(hWindow HWINDOW) HSTRING {
	ret, _, _ := xWnd_GetCursor.Call(uintptr(hWindow))

	return HSTRING(ret)
}

/*
获取HWND句柄.

参数:
	hWindow 窗口句柄.
返回:
	HWND句柄.
*/
func XWndGetHWND(hWindow HWINDOW) HWND {
	ret, _, _ := xWnd_GetHWND.Call(uintptr(hWindow))

	return HWND(ret)
}

/*
启用拖动窗口边框.

参数:
	hWindow 窗口句柄.
	bEnable 是否启用.
*/
func XWndEnableDragBorder(hWindow HWINDOW, bEnable bool) {
	xWnd_EnableDragBorder.Call(
		uintptr(hWindow),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用拖动窗口.

参数:
	hWindow 窗口句柄.
	bEnable 是否启用.
*/
func XWndEnableDragWindow(hWindow HWINDOW, bEnable bool) {
	xWnd_EnableDragWindow.Call(
		uintptr(hWindow),
		uintptr(BoolToBOOL(bEnable)))
}

/*
是否绘制窗口背景.

参数:
	hWindow 窗口句柄.
	bEnable 是否启用.
*/
func XWndEnableDrawBk(hWindow HWINDOW, bEnable bool) {
	xWnd_EnableDrawBk.Call(
		uintptr(hWindow),
		uintptr(BoolToBOOL(bEnable)))
}

/*
当鼠标左键按下是否获得焦点.

参数:
	hWindow 窗口句柄.
	bEnable 是否启用.
*/
func XWndEnableAutoFocus(hWindow HWINDOW, bEnable bool) {
	xWnd_EnableAutoFocus.Call(
		uintptr(hWindow),
		uintptr(BoolToBOOL(bEnable)))
}

/*
允许窗口最大化.

参数:
	hWindow 窗口句柄.
	bEnable 是否启用.
*/
func XWndEnableMaxWindow(hWindow HWINDOW, bEnable bool) {
	xWnd_EnableMaxWindow.Call(
		uintptr(hWindow),
		uintptr(BoolToBOOL(bEnable)))
}

/*
设置鼠标捕获元素.

参数:
	hWindow 窗口句柄.
	hEle 元素句柄.
*/
func XWndSetCaptureEle(hWindow HWINDOW, hEle HELE) {
	xWnd_SetCaptureEle.Call(
		uintptr(hWindow),
		uintptr(hEle))
}

/*
获取当前鼠标捕获元素.

参数:
	hWindow 窗口句柄.
返回:
	元素句柄.
*/
func XWndGetCaptureEle(hWindow HWINDOW) HELE {
	ret, _, _ := xWnd_GetCaptureEle.Call(uintptr(hWindow))

	return HELE(ret)
}

/*
获取重绘区域.

参数:
	hWindow 窗口句柄.
	pRcPaint 重绘区域坐标.
*/
func XWndGetDrawRect(hWindow HWINDOW, pRcPaint *RECT) {
	xWnd_GetDrawRect.Call(
		uintptr(hWindow),
		uintptr(unsafe.Pointer(pRcPaint)))
}

/*
显示隐藏窗口.

参数:
	hWindow 窗口句柄.
	nCmdShow 参见MSDN.
返回:
	参见MSDN.
*/
func XWndShowWindow(hWindow HWINDOW, nCmdShow int) bool {
	ret, _, _ := xWnd_ShowWindow.Call(
		uintptr(hWindow),
		uintptr(nCmdShow))
	// MSDN上返回值：true 为 0，false 为 1
	return ret != TRUE
}

/*
设置布局元素

参数:
	hWindow 窗口句柄.
	nPosition 参见宏定义.
	hEle 元素句柄.
返回:
	成功返回TRUE否则返回FALSE
*/
func XWndBindLayoutEle(hWindow HWINDOW, nPosition WINDOW_POSITION_, hEle HELE) bool {
	ret, _, _ := xWnd_BindLayoutEle.Call(
		uintptr(hWindow),
		uintptr(nPosition),
		uintptr(hEle))

	return ret == TRUE
}

/*
获取布局元素.

参数:
	hWindow 窗口句柄.
	nPosition 参见宏定义.
返回:
	元素句柄.
*/
func XWndGetLayoutEle(hWindow HWINDOW, nPosition WINDOW_POSITION_) HELE {
	ret, _, _ := xWnd_GetLayoutEle.Call(
		uintptr(hWindow),
		uintptr(nPosition))

	return HELE(ret)
}

/*
系统函数,设置窗口类光标句柄.

参数:
	hWindow 窗口句柄.
	hCursor 光标句柄.
返回:
	返回先前的光标句柄.
*/
func XWndSetCursorSys(hWindow HWINDOW, hCursor HCURSOR) HCURSOR {
	ret, _, _ := xWnd_SetCursorSys.Call(
		uintptr(hWindow),
		uintptr(hCursor))

	return HCURSOR(ret)
}

/*
设置窗口字体.

参数:
	hWindow 窗口句柄.
	hFontx 炫彩字体句柄.
*/
func XWndSetFont(hWindow HWINDOW, hFontx HFONTX) {
	xWnd_SetFont.Call(
		uintptr(hWindow),
		uintptr(hFontx))
}

/*
设置窗口ID.

参数:
	hWindow 窗口句柄.
	nID ID值.
*/
func XWndSetID(hWindow HWINDOW, nID int) {
	xWnd_SetID.Call(
		uintptr(hWindow),
		uintptr(nID))
}

/*
获取窗口ID.

参数:
	hWindow 窗口句柄.
返回:
	返回窗口ID值.
*/
func XWndGetID(hWindow HWINDOW) int {
	ret, _, _ := xWnd_GetID.Call(uintptr(hWindow))

	return int(ret)
}

/*
设置布局大小.

参数:
	hWindow 窗口句柄.
	left 窗口左边大小.
	top 窗口上边大小.
	right 窗口右边大小.
	bottom 窗口底部大小.
*/
func XWndSetLayoutSize(hWindow HWINDOW, left, top, right, bottom int) {
	xWnd_SetLayoutSize.Call(
		uintptr(hWindow),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom))
}

/*
获取布局大小.

参数:
	hWindow 窗口句柄.
	pBorderSize 布局大小.
*/
func XWndGetLayoutSize(hWindow HWINDOW, pBorderSize *BorderSize_) {
	xWnd_GetLayoutSize.Call(
		uintptr(hWindow),
		uintptr(unsafe.Pointer(pBorderSize)))
}

/*
设置窗口拖动边框大小.

参数:
	hWindow 窗口句柄.
	left 窗口左边大小.
	top 窗口上边大小.
	right 窗口右边大小.
	bottom 窗口底边大小.
*/
func XWndSetDragBorderSize(hWindow HWINDOW, left, top, right, bottom int) {
	xWnd_SetDragBorderSize.Call(
		uintptr(hWindow),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom))
}

/*
获取窗口拖动边框大小.

参数:
	hWindow 窗口句柄.
	pSize 拖动边框大小.
*/
func XWndGetDragBorderSize(hWindow HWINDOW, pSize *BorderSize_) {
	xWnd_GetDragBorderSize.Call(
		uintptr(hWindow),
		uintptr(unsafe.Pointer(pSize)))
}

/*
设置窗口的最小宽度和高度.

参数:
	hWindow 窗口句柄.
	width 最小宽度.
	height 最小高度.
*/
func XWndSetMinimumSize(hWindow HWINDOW, width, height int) {
	xWnd_SetMinimumSize.Call(
		uintptr(hWindow),
		uintptr(width),
		uintptr(height))
}

/*
检测所在元素.

参数:
	hWindow 窗口句柄.
	pPt 左边点.
返回:
	元素句柄.
*/
func XWndHitChildEle(hWindow HWINDOW, pPt *POINT) HELE {
	ret, _, _ := xWnd_HitChildEle.Call(
		uintptr(hWindow),
		uintptr(unsafe.Pointer(pPt)))

	return HELE(ret)
}

/*
获取当前层子元素数量,不包含子元素的子元素.

参数:
	hWindow 窗口句柄.
返回:
	子元素数量.
*/
func XWndGetChildCount(hWindow HWINDOW) int {
	ret, _, _ := xWnd_GetChildCount.Call(uintptr(hWindow))

	return int(ret)
}

/*
获取当前层子元素通过索引.

参数:
	hWindow 窗口句柄.
	index 元素索引.
返回:
	元素句柄.
*/
func XWndGetChildByIndex(hWindow HWINDOW, index int) HELE {
	ret, _, _ := xWnd_GetChildByIndex.Call(
		uintptr(hWindow),
		uintptr(index))

	return HELE(ret)
}

/*
获取当前层子元素通过元素ID.

参数:
	hWindow 窗口句柄.
	nID 元素ID,ID必须大于0.
返回:
	元素句柄.
*/
func XWndGetChildByID(hWindow HWINDOW, nID int) HELE {
	ret, _, _ := xWnd_GetChildByID.Call(
		uintptr(hWindow),
		uintptr(nID))

	return HELE(ret)
}

/*
获取元素通过元素ID,如果元素不在该窗口上无效.

参数:
	hWindow 窗口句柄.
	nID 元素ID,ID必须大于0.
返回:
	元素句柄.
*/
func XWndGetEle(hWindow HWINDOW, nID int) HELE {
	ret, _, _ := xWnd_GetEle.Call(
		uintptr(hWindow),
		uintptr(nID))

	return HELE(ret)
}

/*
获取子对象(形状对象)数量,只检查当前层.

参数:
	hWindow 窗口句柄.
	返回:返回形状对象数量.
*/
func XWndGetChildShapeCount(hWindow HWINDOW) int {
	ret, _, _ := xWnd_GetChildShapeCount.Call(uintptr(hWindow))

	return int(ret)
}

/*
获取子对象(形状对象)通过索引.

参数:
	hWindow 窗口句柄.
	index 索引.
返回:
	返回形状对象句柄.
*/
func XWndGetChildShapeByIndex(hWindow HWINDOW, index int) HXCGUI {
	ret, _, _ := xWnd_GetChildShapeByIndex.Call(
		uintptr(hWindow),
		uintptr(index))

	return HXCGUI(ret)
}

/*
关闭窗口.

参数:
	hWindow 窗口句柄.
*/
func XWndCloseWindow(hWindow HWINDOW) {
	xWnd_CloseWindow.Call(uintptr(hWindow))
}

/*
绑定布局对象.

参数:
	hWindow 窗口句柄.
	nPosition 参见宏定义.
	hLayout 布局对象.
*/
func XWndBindLayoutObject(hWindow HWINDOW, nPosition WINDOW_POSITION_, hLayout HXCGUI) {
	xWnd_BindLayoutObject.Call(
		uintptr(hWindow),
		uintptr(nPosition),
		uintptr(hLayout))
}

/*
获取布局对象.

参数:
	hWindow 窗口句柄.
	nPosition 参见宏定义.
返回:
	布局对象句柄.
*/
func XWndGetLayoutObject(hWindow HWINDOW, nPosition WINDOW_POSITION_) HXCGUI {
	ret, _, _ := xWnd_GetLayoutObject.Call(
		uintptr(hWindow),
		uintptr(nPosition))

	return HXCGUI(ret)
}

/*
调整窗口布局,该函数只能调整受窗口布局管理的对象,不受窗口布局管理的[布局对象]请使用函数 XWnd_AdjustLayoutObject().

参数:
	hWindow 窗口句柄.
*/
func XWndAdjustLayout(hWindow HWINDOW) {
	xWnd_AdjustLayout.Call(uintptr(hWindow))
}

/*
调整窗口上所有的[布局对象],如果你的[布局对象]不受窗口布局管理,那么你需要调用该函数调整布局,以便显示正确的坐标.

参数:
	hWindow 窗口句柄.
	注解:XWnd_AdjustLayoutObject() 与 XWnd_AdjustLayout() 不分先后顺序 .
*/
func XWndAdjustLayoutObject(hWindow HWINDOW) {
	xWnd_AdjustLayoutObject.Call(uintptr(hWindow))
}

/*
创建插入符.

参数:
	hWindow 窗口句柄.
	hEle 元素句柄.
	width 宽度.
	height 高度.
*/
func XWndCreateCaret(hWindow HWINDOW, hEle HELE, width, height int) {
	xWnd_CreateCaret.Call(
		uintptr(hWindow),
		uintptr(hEle),
		uintptr(width),
		uintptr(height))
}

/*
设置插入符大小.

参数:
	hWindow 窗口句柄.
	width 宽度.
	height 高度.
*/
func XWndSetCaretSize(hWindow HWINDOW, width, height int) {
	xWnd_SetCaretSize.Call(
		uintptr(hWindow),
		uintptr(width),
		uintptr(height))
}

/*
设置插入符位置.

参数:
	hWindow 窗口句柄.
	x x坐标.
	y y坐标.
*/
func XWndSetCaretPos(hWindow HWINDOW, x, y int) {
	xWnd_SetCaretPos.Call(
		uintptr(hWindow),
		uintptr(x),
		uintptr(y))
}

/*
设置插入符位置.

参数:
	hWindow 窗口句柄.
	x x坐标.
	y y坐标.
	width 宽度.
	height 高度.
*/
func XWndSetCaretPosEx(hWindow HWINDOW, x, y, width, height int) {
	xWnd_SetCaretPosEx.Call(
		uintptr(hWindow),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
}

/*
设置插入符颜色.

参数:
	hWindow 窗口句柄.
	color 颜色值.
*/
func XWndSetCaretColor(hWindow HWINDOW, color COLORREF) {
	xWnd_SetCaretColor.Call(
		uintptr(hWindow),
		uintptr(color))
}

/*
显示插入符.

参数:
	hWindow 窗口句柄.
	bShow 是否显示.
*/
func XWndShowCaret(hWindow HWINDOW, bShow BOOL) {
	xWnd_ShowCaret.Call(
		uintptr(hWindow),
		uintptr(bShow))
}

/*
销毁插入符.

参数:
	hWindow 窗口句柄.
*/
func XWndDestroyCaret(hWindow HWINDOW) {
	xWnd_DestroyCaret.Call(uintptr(hWindow))
}

/*
获取插入符所属元素.

参数:
	hWindow 窗口句柄.
返回:
	元素句柄.
*/
func XWndGetCaretHELE(hWindow HWINDOW) HELE {
	ret, _, _ := xWnd_GetCaretHELE.Call(uintptr(hWindow))

	return HELE(ret)
}

/*
获取窗口客户区坐标.

参数:
	hWindow 窗口句柄.
	pRect 坐标.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XWndGetClientRect(hWindow HWINDOW, pRect *RECT) bool {
	ret, _, _ := xWnd_GetClientRect.Call(
		uintptr(hWindow),
		uintptr(unsafe.Pointer(pRect)))

	return ret == TRUE
}

/*
获取body坐标.

参数:
	hWindow 窗口句柄.
	pRect 坐标.
*/
func XWndGetBodyRect(hWindow HWINDOW, pRect *RECT) {
	xWnd_GetBodyRect.Call(
		uintptr(hWindow),
		uintptr(unsafe.Pointer(pRect)))
}

/*
设置窗口定时器.

参数:
	hWindow 窗口句柄.
	nIDEvent 定时器ID.
	uElapse 间隔值,单位毫秒.
返回:
	参见MSDN.
*/
func XWndSetTimer(hWindow HWINDOW, nIDEvent, uElapse uint32) uint32 {
	ret, _, _ := xWnd_SetTimer.Call(
		uintptr(hWindow),
		uintptr(nIDEvent),
		uintptr(uElapse))

	return uint32(ret)
}

/*
关闭定时器.

参数:
	hWindow 窗口句柄.
	nIDEvent 定时器ID.
返回:
	参见MSDN.
*/
func XWndKillTimer(hWindow HWINDOW, nIDEvent uint32) bool {
	ret, _, _ := xWnd_KillTimer.Call(
		uintptr(hWindow),
		uintptr(nIDEvent))

	return ret == TRUE
}

/*
获取背景内容管理器.

参数:
	hWindow 窗口句柄.
	nPosition 窗口布局位置.
返回:
	背景内容管理器.
*/
func XWndGetBkInfoManager(hWindow HWINDOW, nPosition WINDOW_POSITION_) HBKINFOM {
	ret, _, _ := xWnd_GetBkInfoManager.Call(
		uintptr(hWindow),
		uintptr(nPosition))

	return HBKINFOM(ret)
}

/*
设置透明窗口,同时可以通过该函数关闭透明窗口.

参数:
	hWindow 窗口句柄.
	nType 窗口透明类型.
*/
func XWndSetTransparentType(hWindow HWINDOW, nType WINDOW_TRANSPARENT_) {
	xWnd_SetTransparentType.Call(
		uintptr(hWindow),
		uintptr(nType))
}

/*
设置透明窗口的透明度,设置后调用重绘窗口API来更新.

参数:
	hWindow 窗口句柄.
	alpha 窗口透明度,范围0-255之间,0透明,255不透明.
*/
func XWndSetTransparentAlpha(hWindow HWINDOW, alpha byte) {
	xWnd_SetTransparentAlpha.Call(
		uintptr(hWindow),
		uintptr(alpha))
}

/*
设置透明窗口的透明色.

参数:
	hWindow 窗口句柄.
	color 窗口透明色.
*/
func XWndSetTransparentColor(hWindow HWINDOW, color COLORREF) {
	xWnd_SetTransparentColor.Call(
		uintptr(hWindow),
		uintptr(color))
}
