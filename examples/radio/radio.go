package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWndCreate(0, 0, 300, 200, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.XBtnSetType(xcgui.XBtnCreate(5, 3, 60, 20, "Close", xcgui.HXCGUI(hWindow)), xcgui.BUTTON_TYPE_CLOSE)

	hCheck1 := xcgui.XBtnCreate(20, 40, 100, 20, "Radio1", xcgui.HXCGUI(hWindow))
	hCheck2 := xcgui.XBtnCreate(20, 70, 100, 20, "Radio2", xcgui.HXCGUI(hWindow))
	hCheck3 := xcgui.XBtnCreate(20, 100, 100, 20, "Radio3", xcgui.HXCGUI(hWindow))

	xcgui.XBtnSetGroupID(hCheck1, 1)
	xcgui.XBtnSetGroupID(hCheck2, 1)
	xcgui.XBtnSetGroupID(hCheck3, 1)

	xcgui.XBtnSetType(hCheck1, xcgui.BUTTON_TYPE_RADIO)
	xcgui.XBtnSetType(hCheck2, xcgui.BUTTON_TYPE_RADIO)
	xcgui.XBtnSetType(hCheck3, xcgui.BUTTON_TYPE_RADIO)

	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
