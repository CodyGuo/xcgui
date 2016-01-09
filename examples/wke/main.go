package main

import (
	"fmt"
)

import (
	web "github.com/codyguo/xcgui/wke"
	xcgui "github.com/codyguo/xcgui/xc"
)

const (
	VK_RETURN = 13
)

var (
	hWeb    xcgui.HELE
	hWindow xcgui.HWINDOW
	hMenu   xcgui.HMENUX

	hEditUrl xcgui.HXCGUI
	hMenuBar xcgui.HXCGUI

	rtBody xcgui.RECT
)

func main() {
	xcgui.XCLoadResource("xml/resource.xml", "")
	hxcgui := xcgui.XCLoadLayout("xml/layout.xml", 0)
	hWindow = xcgui.HWINDOW(hxcgui)

	xcgui.XWndAdjustLayout(hWindow)

	// 获取url编辑文本
	hEditUrl = xcgui.XCGetObjectByID(131)
	xcgui.XEleRegEventC(xcgui.HELE(hEditUrl), xcgui.XE_CHAR, xcgui.CallBack(onEventChar))

	// 浏览器body区
	xcgui.XWndGetBodyRect(hWindow, &rtBody)

	hWeb = web.XWebCreate(
		int(rtBody.Left),
		int(rtBody.Top),
		int(rtBody.Right-rtBody.Left),
		int(rtBody.Bottom-rtBody.Top),
		xcgui.XWndGetHWND(hWindow))

	// 默认打开百度
	web.XWebLoadUrl(hWeb, "http://www.baidu.com")

	// URL变化
	web.XWebOnURLChanged(hWeb, xcgui.CallBack(onURLChanged), 1)

	// 窗口变化web窗体自适应
	xcgui.XWndRegEventC(hWindow, xcgui.WM_SIZE, xcgui.CallBack(OnWndSize))

	// 更多
	hMenuBar = xcgui.XCGetObjectByID(133)
	xcgui.XMenuBarAddButton(xcgui.HELE(hMenuBar), "更多")
	hMenu = xcgui.XMenuBarGetMenu(xcgui.HELE(hMenuBar), 0)

	if hMenu != 0 {
		xcgui.XMenuAddItem(hMenu, 1330, "主页", 0, 0)
		xcgui.XMenuAddItem(hMenu, 1331, "重新载入", 0, 0)

	}

	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func onEventChar(wParam, lParam uintptr, pbHandled *bool) int {
	// fmt.Println(string(wParam), lParam)
	if wParam == VK_RETURN {
		url, _ := xcgui.XRichEditGetTextGo(xcgui.HELE(hEditUrl))
		fmt.Println(url)
		web.XWebLoadUrl(hWeb, url)
	}

	return 0
}

func OnWndSize(nFlags uint32, pSize *xcgui.SIZE, pbHandled *bool) int {
	xcgui.XWndGetBodyRect(hWindow, &rtBody)
	web.XWebMoveWindow(hWeb,
		int(rtBody.Left),
		int(rtBody.Top),
		int(rtBody.Right-rtBody.Left),
		int(rtBody.Bottom-rtBody.Top))

	return 0
}

func onURLChanged(webView, param, url uintptr) int {
	openUrl := web.XWebGetStringW(url)
	xcgui.XRichEditSetText(xcgui.HELE(hEditUrl), openUrl)
	fmt.Println(openUrl)

	return 0
}
