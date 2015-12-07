package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWndCreate(0, 0, 400, 300, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hEle := xcgui.XEleCreate(60, 60, 200, 200, xcgui.HXCGUI(hWindow))
	hTextBlock := xcgui.XShapeTextCreate(0, 0, 100, 20, "123456", xcgui.HXCGUI(hEle))

	xcgui.XShapeTextSetLayoutWidth(hTextBlock, xcgui.LAYOUT_SIZE_TYPE_AUTO, 0)
	xcgui.XShapeTextSetLayoutHeight(hTextBlock, xcgui.LAYOUT_SIZE_TYPE_AUTO, 0)

	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
