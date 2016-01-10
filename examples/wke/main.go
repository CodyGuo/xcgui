package main

import (
	"fmt"
)

import (
	web "github.com/codyguo/xcgui/wke"
	xcgui "github.com/codyguo/xcgui/xc"
	"github.com/lxn/win"
)

const (
	VK_RETURN = 13
)

var (
	hWeb    xcgui.HWINDOW
	hWindow xcgui.HWINDOW
	hMenu   xcgui.HMENUX

	hGoBack    xcgui.HXCGUI
	hGoForward xcgui.HXCGUI
	hEditUrl   xcgui.HXCGUI
	hMenuBar   xcgui.HXCGUI
	hState     xcgui.HXCGUI

	rtBody xcgui.RECT
)

func main() {
	xcgui.XCLoadResource("xml/resource.xml", "")
	hxcgui := xcgui.XCLoadLayout("xml/layout.xml", 0)
	hWindow = xcgui.HWINDOW(hxcgui)
	xcgui.XWndAdjustLayout(hWindow)

	// 后退
	hGoBack = xcgui.XCGetObjectByID(130)
	// xcgui.XEleRegEventC1(xcgui.HELE(hGoBack), xcgui.XE_BNCLICK, xcgui.CallBack(onBtnClick))
	xcgui.XEleRegEventC1(xcgui.HELE(hGoBack),
		xcgui.XE_BNCLICK,
		xcgui.CallBackGo(func() {
			fmt.Println("后退中...")
			web.XWebGoBack(hWeb)
		}))

	// url地址栏
	hEditUrl = xcgui.XCGetObjectByID(131)
	xcgui.XEleEnableSwitchFocus(xcgui.HELE(hEditUrl), true)
	xcgui.XEleRegEventC(xcgui.HELE(hEditUrl), xcgui.XE_CHAR, xcgui.CallBack(onEventChar))
	// 解决焦点丢失问题
	xcgui.XEleRegEventC(xcgui.HELE(hEditUrl), xcgui.XE_LBUTTONDOWN, xcgui.CallBack(onLButtonDown))

	// 前进
	hGoForward = xcgui.XCGetObjectByID(132)
	xcgui.XEleRegEventC1(xcgui.HELE(hGoForward), xcgui.XE_BNCLICK, xcgui.CallBack(onBtnClick))

	// 浏览器body区
	xcgui.XWndGetBodyRect(hWindow, &rtBody)

	hWeb = web.XWebCreate(
		int(rtBody.Left),
		int(rtBody.Top),
		int(rtBody.Right-rtBody.Left),
		int(rtBody.Bottom-rtBody.Top),
		xcgui.XWndGetHWND(hWindow))

	// 更多
	hMenuBar = xcgui.XCGetObjectByID(133)
	xcgui.XMenuBarAddButton(xcgui.HELE(hMenuBar), "更多")
	hMenu = xcgui.XMenuBarGetMenu(xcgui.HELE(hMenuBar), 0)

	if hMenu != 0 {
		xcgui.XMenuAddItem(hMenu, 1330, "主页", 0, 0)
		xcgui.XMenuAddItem(hMenu, 1331, "重新载入", 0, 0)

	}

	// 默认打开百度
	web.XWebLoadUrl(hWeb, "http://www.baidu.com")

	// URL变化
	web.XWebOnURLChanged(hWeb, xcgui.CallBack(onURLChanged), 1)

	// Title 变化修改状态栏
	hState = xcgui.XCGetObjectByID(311)
	web.XWebOnTitleChanged(hWeb, xcgui.CallBack(onTitleChanged), 1)

	// 窗口变化web窗体自适应
	xcgui.XWndRegEventC(hWindow, xcgui.WM_SIZE, xcgui.CallBack(OnWndSize))

	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func onBtnClick(hEventEle xcgui.HELE, pbHandled *bool) int {
	switch hEventEle {
	case xcgui.HELE(hGoBack):
		web.XWebGoBack(hWeb)
	case xcgui.HELE(hGoForward):
		web.XWebGoForward(hWeb)
	}

	return 0
}

func onEventChar(wParam, lParam uintptr, pbHandled *bool) int {
	// fmt.Println(string(wParam), lParam)
	if wParam == VK_RETURN {
		url, _ := xcgui.XRichEditGetTextGo(xcgui.HELE(hEditUrl))
		fmt.Println("----------------------------------------------")
		fmt.Println("URL地址栏: ", url)
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
	fmt.Println("正在访问: ", openUrl)

	return 0
}

func onTitleChanged(webView, param, title uintptr) int {
	fmt.Println("浏览器标题: ", web.XWebGetStringA(title))
	xcgui.XShapeTextSetText(hState, web.XWebGetStringA(title))
	xcgui.XWndRedrawWnd(hWindow)

	return 0
}

func onLButtonDown(nFlags uint32, pPt *xcgui.POINT, pbHandled *bool) int {
	win.SetFocus(win.HWND(xcgui.XWndGetHWND(hWindow)))

	return 0
}
