package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWnd_Create(0, 0, 300, 200, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hSliderBar := xcgui.XSliderBar_Create(20, 40, 260, 60, xcgui.HXCGUI(hWindow))
	xcgui.XSliderBar_SetRange(hSliderBar, 10)
	xcgui.XSliderBar_SetButtonHeight(hSliderBar, 27)
	xcgui.XSliderBar_SetButtonWidth(hSliderBar, 27)
	xcgui.XSliderBar_SetSpaceTwo(hSliderBar, 10, 10)

	xcgui.XEle_RegEventC(hSliderBar, xcgui.XE_SLIDERBAR_CHANGE, xcgui.CallBack(OnSliderBarChange))

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()

}

func OnSliderBarChange(pos int, pbHandled *bool) int {
	*pbHandled = true

	return 0
}
