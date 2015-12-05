package main

import (
	"log"
)

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWndCreate(0, 0, 300, 200, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	MenuBar := xcgui.XMenuBarCreate(20, 40, 260, 28, xcgui.HXCGUI(hWindow))
	xcgui.XMenuBarAddButton(MenuBar, "文件")
	xcgui.XMenuBarAddButton(MenuBar, "编辑")
	xcgui.XMenuBarAddButton(MenuBar, "视图")
	xcgui.XMenuBarAddButton(MenuBar, "aaa")

	hMenu := xcgui.XMenuBarGetMenu(MenuBar, 0)
	if hMenu != 0 {
		xcgui.XMenuAddItem(hMenu, 101, "101", 0, 0)
		xcgui.XMenuAddItem(hMenu, 102, "102", 0, 0)
		xcgui.XMenuAddItem(hMenu, 103, "103", 0, 0)
		xcgui.XMenuAddItem(hMenu, 1031, "3-1", 103, 0)
	}

	hMenu = xcgui.XMenuBarGetMenu(MenuBar, 1)
	if hMenu != 0 {
		xcgui.XMenuAddItem(hMenu, 201, "201", 0, 0)
		xcgui.XMenuAddItem(hMenu, 202, "202", 0, 0)
		xcgui.XMenuAddItem(hMenu, 203, "203", 0, 0)
	}

	hMenu = xcgui.XMenuBarGetMenu(MenuBar, 2)
	if hMenu != 0 {
		xcgui.XMenuAddItem(hMenu, 301, "301", 0, 0)
	}

	xcgui.XEleRegEventC(MenuBar, xcgui.XE_MENU_SELECT, xcgui.CallBack(OnWndMenuSelet))
	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func OnWndMenuSelet(nID int, pbHandled *bool) int {
	log.Printf("select item %d \n", nID)

	return 0
}
