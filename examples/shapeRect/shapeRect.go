package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWndCreate(0, 0, 400, 300, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	//    hShapeRect
	xcgui.XShapeRectCreate(20, 50, 100, 100, xcgui.HXCGUI(hWindow))
	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
