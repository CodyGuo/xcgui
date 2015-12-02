package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

var (
	mw xcgui.HWINDOW
)

func main() {
	mw = xcgui.XFrameWndCreate(0, 0, 600, 400, xcgui.StringToUTF16Ptr("炫彩界面库窗口"), xcgui.NULL, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.XBtnSetType(xcgui.XBtnCreate(3, 3, 20, 20, xcgui.StringToUTF16Ptr("X"), xcgui.HXCGUI(mw)), xcgui.BUTTON_TYPE_CLOSE)

	hPane1 := xcgui.XPaneCreate(xcgui.StringToUTF16Ptr("111"), 200, 200, mw)
	hPane2 := xcgui.XPaneCreate(xcgui.StringToUTF16Ptr("2222222222"), 200, 200, mw)
	hPane3 := xcgui.XPaneCreate(xcgui.StringToUTF16Ptr("333"), 200, 200, mw)

	xcgui.XFrameWndAddPane(mw, xcgui.NULL, hPane1, xcgui.Align_left)

	xcgui.XFrameWndAddPane(mw, xcgui.NULL, hPane2, xcgui.Align_bottom)

	xcgui.XFrameWndAddPane(mw, xcgui.NULL, hPane3, xcgui.Align_right)

	xcgui.XWndRegEventC(mw, xcgui.WM_RBUTTONUP, xcgui.CallBack(OnWndButtonUp))

	xcgui.XWndAdjustLayout(mw)
	xcgui.XWndShowWindow(mw, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func OnWndButtonUp(nFlags uint32, pPt *xcgui.POINT, pbHandled *bool) int {
	pt := *pPt

	hMenu := xcgui.XMenuCreate()
	xcgui.XMenuAddItem(hMenu, 201, xcgui.StringToUTF16Ptr("窗格1"), xcgui.NULL, xcgui.NULL)
	xcgui.XMenuAddItem(hMenu, 202, xcgui.StringToUTF16Ptr("窗格2"), xcgui.NULL, xcgui.NULL)
	xcgui.XMenuAddItem(hMenu, 203, xcgui.StringToUTF16Ptr("窗格3"), xcgui.NULL, xcgui.NULL)

	xcgui.ClientToScreen(xcgui.XWndGetHWND(mw), &pt)
	xcgui.XMenuPopup(hMenu, xcgui.XWndGetHWND(mw), pt.X, pt.Y, xcgui.NULL, xcgui.NULL)

	return 0
}
