package main

import (
	"fmt"
)

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

var (
	hButtonAdd   xcgui.HELE
	hProgressBar xcgui.HELE
)

func main() {
	hWindow := xcgui.XWndCreate(0, 0, 300, 200, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hProgressBar = xcgui.XProgBarCreate(20, 40, 260, 20, xcgui.HXCGUI(hWindow))
	xcgui.XProgBarSetRange(hProgressBar, 100)
	xcgui.XProgBarSetPos(hProgressBar, 50)
	xcgui.XProgBarSetSpaceTwo(hProgressBar, 10, 10)

	hButtonAdd = xcgui.XBtnCreate(20, 70, 50, 18, "+", xcgui.HXCGUI(hWindow))
	hButtonMunus := xcgui.XBtnCreate(80, 70, 50, 18, "-", xcgui.HXCGUI(hWindow))

	xcgui.XEleRegEventC1(hButtonAdd, xcgui.XE_BNCLICK, xcgui.CallBack(OnBtnClick))
	xcgui.XEleRegEventC1(hButtonMunus, xcgui.XE_BNCLICK, xcgui.CallBack(OnBtnClick))
	xcgui.XEleRegEventC(hProgressBar, xcgui.XE_PROGRESSBAR_CHANGE, xcgui.CallBack(OnProgressBarChange))

	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()

}

func OnBtnClick(hEventEle xcgui.HELE, pbHandled *bool) int {
	if hButtonAdd == hEventEle {
		xcgui.XProgBarSetPos(hProgressBar, xcgui.XProgBarGetPos(hProgressBar)+1)
	} else {
		xcgui.XProgBarSetPos(hProgressBar, xcgui.XProgBarGetPos(hProgressBar)-1)
	}
	xcgui.XEleRedrawEle(hProgressBar)

	return 0
}

func OnProgressBarChange(pos int, pbHandled *bool) int {
	fmt.Println("ProgressBar 进度改变中: ", pos)

	return 0
}
