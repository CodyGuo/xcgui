package main

import (
	"log"
)

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

var (
	hWindow xcgui.HWINDOW
)

func main() {
	hWindow = xcgui.XWndCreate(100, 100, 400, 300, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.XBtnSetType(xcgui.XBtnCreate(5, 3, 60, 20, "Close", xcgui.HXCGUI(hWindow)), xcgui.BUTTON_TYPE_CLOSE)

	hButton := xcgui.XBtnCreate(20, 50, 120, 20, "popu-modalWindow", xcgui.HXCGUI(hWindow))
	xcgui.XEleRegEventC(hButton, xcgui.XE_BNCLICK, xcgui.CallBack(OnBtnClick))
	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)

	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func OnBtnClick(pbHandled *bool) int {
	hWindowModal := xcgui.XModalWndCreate(200, 200, "炫彩界面窗口", xcgui.XWndGetHWND(hWindow), xcgui.XC_WINDOW_STYLE_MODAL)
	xcgui.XBtnSetType(xcgui.XBtnCreate(5, 3, 60, 20, "Close", xcgui.HXCGUI(hWindowModal)), xcgui.BUTTON_TYPE_CLOSE)

	nResult := xcgui.XModalWndDoModal(hWindowModal)

	log.Println("exit modal ", nResult)

	*pbHandled = true

	return 0
}
