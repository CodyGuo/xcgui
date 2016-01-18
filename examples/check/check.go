package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWnd_Create(0, 0, 300, 200, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hCheck1 := xcgui.XBtn_Create(20, 40, 100, 20, "Check1", xcgui.HXCGUI(hWindow))
	hCheck2 := xcgui.XBtn_Create(20, 70, 100, 20, "Check2", xcgui.HXCGUI(hWindow))
	hCheck3 := xcgui.XBtn_Create(20, 100, 100, 20, "Check3", xcgui.HXCGUI(hWindow))
	xcgui.XBtn_SetType(hCheck1, xcgui.BUTTON_TYPE_CHECK)
	xcgui.XBtn_SetType(hCheck2, xcgui.BUTTON_TYPE_CHECK)
	xcgui.XBtn_SetType(hCheck3, xcgui.BUTTON_TYPE_CHECK)

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
