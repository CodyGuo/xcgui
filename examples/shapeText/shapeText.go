package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWnd_Create(0, 0, 400, 300, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hEle := xcgui.XEle_Create(60, 60, 200, 200, xcgui.HXCGUI(hWindow))
	hTextBlock := xcgui.XShapeText_Create(0, 0, 100, 20, "123456", xcgui.HXCGUI(hEle))

	xcgui.XShapeText_SetLayoutWidth(hTextBlock, xcgui.LAYOUT_SIZE_TYPE_AUTO, 0)
	xcgui.XShapeText_SetLayoutHeight(hTextBlock, xcgui.LAYOUT_SIZE_TYPE_AUTO, 0)

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
