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
	hWindow = xcgui.XWnd_Create(100, 100, 400, 300, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.XBtn_SetType(xcgui.XBtn_Create(5, 3, 60, 20, "Close", xcgui.HXCGUI(hWindow)), xcgui.BUTTON_TYPE_CLOSE)

	hButton := xcgui.XBtn_Create(20, 50, 120, 20, "popu-modalWindow", xcgui.HXCGUI(hWindow))
	xcgui.XEle_RegEventC(hButton, xcgui.XE_BNCLICK, xcgui.CallBack(OnBtnClick))
	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)

	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func OnBtnClick(pbHandled *bool) int {
	hWindowModal := xcgui.XModalWnd_Create(200, 200, "炫彩界面窗口", xcgui.XWnd_GetHWND(hWindow), xcgui.XC_WINDOW_STYLE_MODAL)
	xcgui.XBtn_SetType(xcgui.XBtn_Create(5, 3, 60, 20, "Close", xcgui.HXCGUI(hWindowModal)), xcgui.BUTTON_TYPE_CLOSE)

	nResult := xcgui.XModalWnd_DoModal(hWindowModal)

	log.Println("exit modal ", nResult)

	*pbHandled = true

	return 0
}
