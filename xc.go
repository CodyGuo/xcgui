package xcgui

import (
    "github.com/codyguo/sys"
)

type XCClass struct {
    Text  string
    Hwnd  uintptr
    WHwnd uintptr
}

func (xc *XCClass) Show(uFlag int) {
    XCDLL.Call("XWnd_ShowWindow", xc.Hwnd, uintptr(uFlag))
}

func (xc *XCClass) LoopMessage() {
    XRunXCGUI()
}

func XCGUI(x, y, cx, cy int, text string, hParent int, uFlag int) *XCClass {
    XCCls := &XCClass{}
    XCCls.Hwnd = XCDLL.Call("XWnd_Create", uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), sys.TEXT(text), uintptr(hParent), uintptr(uFlag))
    XCCls.WHwnd = XCDLL.Call("XWnd_GetHWND", XCCls.Hwnd)
    XCCls.Text = text
    return XCCls
}

func XRunXCGUI() {
    XCDLL.Call("XRunXCGUI")
}

/*

	hWindow := XCGUI.Call("XWnd_Create", 100, 100, 600, 400, TEXT("XCGUI Library For GoLang"), 0, 0x1|0x8)
	hBtn := XCGUI.Call("XBtn_Create", 50, 50, 100, 100, TEXT("Hey Man"), hWindow)
	XCGUI.Call("XEle_RegEventC", hBtn, 34, uintptr(syscall.NewCallback(HeyMan)))
	XCGUI.Call("XWnd_ShowWindow", hWindow, 0x5)
	XCGUI.Call("XRunXCGUI")
*/
