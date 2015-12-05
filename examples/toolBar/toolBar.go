package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWndCreate(0, 0, 500, 200, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hToolBar := xcgui.XToolBarCreate(20, 40, 320, 28, xcgui.HXCGUI(hWindow))

	hButton1 := xcgui.XBtnCreate(0, 0, 60, 20, "Button1", xcgui.HXCGUI(hToolBar))
	hButton2 := xcgui.XBtnCreate(0, 0, 60, 20, "Button2", xcgui.HXCGUI(hToolBar))
	hButton3 := xcgui.XBtnCreate(0, 0, 60, 20, "Button3", xcgui.HXCGUI(hToolBar))
	hButton4 := xcgui.XBtnCreate(0, 0, 60, 20, "Button4", xcgui.HXCGUI(hToolBar))
	hButton5 := xcgui.XBtnCreate(0, 0, 60, 20, "Button5", xcgui.HXCGUI(hToolBar))
	xcgui.XToolBarInsertEle(hToolBar, hButton1, -1)
	xcgui.XToolBarInsertEle(hToolBar, hButton2, -1)
	xcgui.XToolBarInsertEle(hToolBar, hButton3, -1)
	xcgui.XToolBarInsertEle(hToolBar, hButton4, -1)
	xcgui.XToolBarInsertEle(hToolBar, hButton5, -1)

	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
