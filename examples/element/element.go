package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWnd_Create(0, 0, 300, 200, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.XBtn_SetType(xcgui.XBtn_Create(10, 5, 60, 20, "Close", xcgui.HXCGUI(hWindow)), xcgui.BUTTON_TYPE_CLOSE)

	xcgui.XEle_Create(20, 50, 100, 100, xcgui.HXCGUI(hWindow))

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)

	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
