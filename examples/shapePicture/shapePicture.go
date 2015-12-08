package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWndCreate(0, 0, 400, 400, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hPicture := xcgui.XShapePicCreate(20, 50, 100, 100, xcgui.HXCGUI(hWindow))
	xcgui.XShapePicSetImage(hPicture, xcgui.XImageLoadFile("../img/comma_face_02.png", false))

	xcgui.XWndAdjustLayout(hWindow)
	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
