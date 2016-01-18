package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

var (
	mw xcgui.HWINDOW
)

func main() {
	mw = xcgui.XFrameWnd_Create(0, 0, 600, 400, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.XBtn_SetType(xcgui.XBtn_Create(3, 3, 20, 20, "X", xcgui.HXCGUI(mw)), xcgui.BUTTON_TYPE_CLOSE)

	hPane1 := xcgui.XPane_Create("111", 200, 200, mw)
	hPane2 := xcgui.XPane_Create("2222222222", 200, 200, mw)
	hPane3 := xcgui.XPane_Create("333", 200, 200, mw)

	xcgui.XFrameWnd_AddPane(mw, 0, hPane1, xcgui.Align_left)

	xcgui.XFrameWnd_AddPane(mw, 0, hPane2, xcgui.Align_bottom)

	xcgui.XFrameWnd_AddPane(mw, 0, hPane3, xcgui.Align_right)

	xcgui.XWnd_RegEventC(mw, xcgui.WM_RBUTTONUP, xcgui.CallBack(OnWndButtonUp))

	xcgui.XWnd_AdjustLayout(mw)
	xcgui.XWnd_ShowWindow(mw, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func OnWndButtonUp(nFlags uint32, pPt *xcgui.POINT, pbHandled *bool) int {
	pt := *pPt

	hMenu := xcgui.XMenu_Create()
	xcgui.XMenu_AddItem(hMenu, 201, "窗格1", 0, 0)
	xcgui.XMenu_AddItem(hMenu, 202, "窗格2", 0, 0)
	xcgui.XMenu_AddItem(hMenu, 203, "窗格3", 0, 0)

	xcgui.ClientToScreen(xcgui.XWnd_GetHWND(mw), &pt)
	xcgui.XMenu_Popup(hMenu, xcgui.XWnd_GetHWND(mw), pt.X, pt.Y, 0, 0)

	return 0
}
