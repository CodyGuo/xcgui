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

	hSBar1 := xcgui.XSBar_Create(20, 50, 200, 20, xcgui.HXCGUI(hWindow))
	xcgui.XEle_RegEventC(hSBar1, xcgui.XE_SBAR_SCROLL, xcgui.CallBack(OnSBarScroll))

	hSBar2 := xcgui.XSBar_Create(230, 50, 20, 200, xcgui.HXCGUI(hWindow))
	xcgui.XSBar_SetHorizon(hSBar2, false)

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func OnSBarScroll(pos int, pbHandled *bool) int {
	log.Printf("pos=%d \n", pos)
	*pbHandled = true
	return 0
}
