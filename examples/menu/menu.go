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

type menu_popupWnd_i struct {
	hWindow   xcgui.HWINDOW
	nParentID int
}

type menu_drawBackground_i struct {
	hMenu     xcgui.HMENUX
	hWindow   xcgui.HWINDOW
	nParentID int
}

type menu_drawItem_i struct {
	hMenu   xcgui.HMENUX
	hWindow xcgui.HWINDOW
	nID     int
	nState  int
	rcItem  xcgui.RECT
	hIcon   xcgui.HIMAGE
	pText   string
}

func main() {
	hWindow = xcgui.XWndCreate(0, 0, 300, 200, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hButton := xcgui.XBtnCreate(20, 50, 80, 20, "弹出菜单", xcgui.HXCGUI(hWindow))
	xcgui.XEleRegEventC1(hButton, xcgui.XE_BNCLICK, xcgui.CallBack(OnbtnClick))

	xcgui.XWndRegEventC(hWindow, xcgui.XWM_MENU_POPUP, xcgui.CallBack(OnWndMenuPopup))
	xcgui.XWndRegEventC(hWindow, xcgui.XWM_MENU_POPUP_WND, xcgui.CallBack(OnWndMenuPopupWnd))
	xcgui.XWndRegEventC(hWindow, xcgui.XWM_MENU_SELECT, xcgui.CallBack(OnWndMenuSelect))
	xcgui.XWndRegEventC(hWindow, xcgui.XWM_MENU_EXIT, xcgui.CallBack(OnWndMenuExit))

	xcgui.XWndRegEventC(hWindow, xcgui.XWM_MENU_DRAW_BACKGROUND, xcgui.CallBack(OnWndMenuDrawBackground))

	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()

}

func OnbtnClick(hEventEle xcgui.HELE, pbHandled *bool) int {
	hMenu := xcgui.XMenuCreate()
	xcgui.XMenuSetBkImage(hMenu, xcgui.XImageLoadFile("../img/comma_face_12.png", true))

	xcgui.XMenuAddItemIcon(hMenu, 201, "111", xcgui.XC_ID_ROOT, xcgui.XImageLoadFile("../img/plus.png", false), xcgui.MENU_STATE_FLAGS_CHECK)
	xcgui.XMenuAddItem(hMenu, 202, "222", 0, 0)
	xcgui.XMenuAddItem(hMenu, 203, "333", 0, 0)

	xcgui.XMenuAddItem(hMenu, 204, "444", 203, 0)
	xcgui.XMenuAddItem(hMenu, 205, "555", 203, 0)

	rcButton := new(xcgui.RECT)
	xcgui.XEleGetRect(hEventEle, rcButton)
	pt := xcgui.POINT{int(rcButton.Left), int(rcButton.Bottom)}
	xcgui.ClientToScreen(xcgui.XWndGetHWND(hWindow), &pt)
	xcgui.XMenuPopup(hMenu, xcgui.XWndGetHWND(hWindow), pt.X, pt.Y, 0, 0)

	return 0
}

func OnWndMenuPopup(hMenu xcgui.HMENUX, pbHandled *bool) int {
	log.Println("menu-XWM_MENU_POPU.")

	return 0
}

func OnWndMenuPopupWnd(hMenu xcgui.HMENUX, pInfo *menu_popupWnd_i, pbHandled *bool) int {
	log.Printf("menu-XWM_MENU_PUPUP_WND nParent=%d \n", pInfo.nParentID)

	return 0
}

func OnWndMenuSelect(nID int, pBool *bool) int {
	log.Println("menu-XWM_MENU_SELECT item: ", nID)

	return 0
}

func OnWndMenuExit(pbHandled *bool) int {
	log.Println("menu-XWM_MENU_EXIT exit.")

	return 0
}

func OnWndMenuDrawBackground(hDraw xcgui.HDRAW, pInfo *menu_drawBackground_i, pbHandled *bool) int {
	*pbHandled = true

	return 0
}
