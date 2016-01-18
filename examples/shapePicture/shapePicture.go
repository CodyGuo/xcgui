package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWnd_Create(0, 0, 400, 400, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hPicture := xcgui.XShapePic_Create(20, 50, 100, 100, xcgui.HXCGUI(hWindow))
	xcgui.XShapePic_SetImage(hPicture, xcgui.XImage_LoadFile("../img/comma_face_02.png", false))

	xcgui.XWnd_AdjustLayout(hWindow)
	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
