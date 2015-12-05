package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWndCreate(0, 0, 300, 200, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hSliderBar := xcgui.XSliderBarCreate(20, 40, 260, 60, xcgui.HXCGUI(hWindow))
	xcgui.XSliderBarSetRange(hSliderBar, 10)
	xcgui.XSliderBarSetButtonHeight(hSliderBar, 27)
	xcgui.XSliderBarSetButtonWidth(hSliderBar, 27)
	xcgui.XSliderBarSetSpaceTwo(hSliderBar, 10, 10)

	xcgui.XEleRegEventC(hSliderBar, xcgui.XE_SLIDERBAR_CHANGE, xcgui.CallBack(OnSliderBarChange))

	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()

}

func OnSliderBarChange(pos int, pbHandled *bool) int {
	*pbHandled = true

	return 0
}
