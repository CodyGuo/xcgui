package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWnd_Create(0, 0, 500, 200, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hToolBar := xcgui.XToolBar_Create(20, 40, 320, 28, xcgui.HXCGUI(hWindow))

	hButton1 := xcgui.XBtn_Create(0, 0, 60, 20, "Button1", xcgui.HXCGUI(hToolBar))
	hButton2 := xcgui.XBtn_Create(0, 0, 60, 20, "Button2", xcgui.HXCGUI(hToolBar))
	hButton3 := xcgui.XBtn_Create(0, 0, 60, 20, "Button3", xcgui.HXCGUI(hToolBar))
	hButton4 := xcgui.XBtn_Create(0, 0, 60, 20, "Button4", xcgui.HXCGUI(hToolBar))
	hButton5 := xcgui.XBtn_Create(0, 0, 60, 20, "Button5", xcgui.HXCGUI(hToolBar))
	xcgui.XToolBar_InsertEle(hToolBar, hButton1, -1)
	xcgui.XToolBar_InsertEle(hToolBar, hButton2, -1)
	xcgui.XToolBar_InsertEle(hToolBar, hButton3, -1)
	xcgui.XToolBar_InsertEle(hToolBar, hButton4, -1)
	xcgui.XToolBar_InsertEle(hToolBar, hButton5, -1)

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
