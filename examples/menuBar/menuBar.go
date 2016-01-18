package main

import (
	"log"
)

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWnd_Create(0, 0, 300, 200, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	MenuBar := xcgui.XMenuBar_Create(20, 40, 260, 28, xcgui.HXCGUI(hWindow))
	xcgui.XMenuBar_AddButton(MenuBar, "文件")
	xcgui.XMenuBar_AddButton(MenuBar, "编辑")
	xcgui.XMenuBar_AddButton(MenuBar, "视图")
	xcgui.XMenuBar_AddButton(MenuBar, "aaa")

	hMenu := xcgui.XMenuBar_GetMenu(MenuBar, 0)
	if hMenu != 0 {
		xcgui.XMenu_AddItem(hMenu, 101, "101", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 102, "102", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 103, "103", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 1031, "3-1", 103, 0)
	}

	hMenu = xcgui.XMenuBar_GetMenu(MenuBar, 1)
	if hMenu != 0 {
		xcgui.XMenu_AddItem(hMenu, 201, "201", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 202, "202", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 203, "203", 0, 0)
	}

	hMenu = xcgui.XMenuBar_GetMenu(MenuBar, 2)
	if hMenu != 0 {
		xcgui.XMenu_AddItem(hMenu, 301, "301", 0, 0)
	}

	xcgui.XEle_RegEventC(MenuBar, xcgui.XE_MENU_SELECT, xcgui.CallBack(OnWndMenuSelet))
	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func OnWndMenuSelet(nID int, pbHandled *bool) int {
	log.Printf("select item %d \n", nID)

	return 0
}
