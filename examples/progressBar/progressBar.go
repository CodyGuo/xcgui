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
	hWindow := xcgui.XWnd_Create(0, 0, 300, 200, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hProgressBar = xcgui.XProgBar_Create(20, 40, 260, 20, xcgui.HXCGUI(hWindow))
	xcgui.XProgBar_SetRange(hProgressBar, 100)
	xcgui.XProgBar_SetPos(hProgressBar, 50)
	xcgui.XProgBar_SetSpaceTwo(hProgressBar, 10, 10)

	hButtonAdd = xcgui.XBtn_Create(20, 70, 50, 18, "+", xcgui.HXCGUI(hWindow))
	hButtonMunus := xcgui.XBtn_Create(80, 70, 50, 18, "-", xcgui.HXCGUI(hWindow))

	xcgui.XEle_RegEventC1(hButtonAdd, xcgui.XE_BNCLICK, xcgui.CallBack(OnBtnClick))
	xcgui.XEle_RegEventC1(hButtonMunus, xcgui.XE_BNCLICK, xcgui.CallBack(OnBtnClick))
	xcgui.XEle_RegEventC(hProgressBar, xcgui.XE_PROGRESSBAR_CHANGE, xcgui.CallBack(OnProgressBarChange))

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()

}

func OnBtnClick(hEventEle xcgui.HELE, pbHandled *bool) int {
	if hButtonAdd == hEventEle {
		xcgui.XProgBar_SetPos(hProgressBar, xcgui.XProgBar_GetPos(hProgressBar)+1)
	} else {
		xcgui.XProgBar_SetPos(hProgressBar, xcgui.XProgBar_GetPos(hProgressBar)-1)
	}
	xcgui.XEle_RedrawEle(hProgressBar)

	return 0
}

func OnProgressBarChange(pos int, pbHandled *bool) int {
	fmt.Println("ProgressBar 进度改变中: ", pos)

	return 0
}
