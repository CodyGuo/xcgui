package main

import (
	"log"
)

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWnd_Create(0, 0, 300, 300, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hScrollView := xcgui.XSView_Create(20, 50, 200, 200, xcgui.HXCGUI(hWindow))
	xcgui.XSView_SetTotalSize(hScrollView, 300, 300)

	xcgui.XBtn_Create(10, 10, 100, 20, "Button1", xcgui.HXCGUI(hScrollView))
	xcgui.XBtn_Create(10, 40, 100, 20, "Button2", xcgui.HXCGUI(hScrollView))
	xcgui.XBtn_Create(10, 70, 100, 20, "Button3", xcgui.HXCGUI(hScrollView))

	xcgui.XEle_RegEventC(hScrollView, xcgui.XE_SCROLLVIEW_SCROLL_H, xcgui.CallBack(OnScrollViewScrollH))
	xcgui.XEle_RegEventC(hScrollView, xcgui.XE_SCROLLVIEW_SCROLL_V, xcgui.CallBack(OnScrollViewScrollV))

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func OnScrollViewScrollH(pos int, pbHandled *bool) int {
	log.Printf("XE_SCROLLVIEW_SCROLL_H %d\n", pos)

	return 0
}

func OnScrollViewScrollV(pos int, pbHandled *bool) int {
	log.Printf("XE_SCROLLVIEW_SCROLL_V %d\n", pos)

	return 0
}
