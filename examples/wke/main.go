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
	hWeb    uintptr
	hWindow xcgui.HWINDOW
	hMenu   xcgui.HMENUX

	hGoBack    xcgui.HXCGUI
	hGoForward xcgui.HXCGUI
	hEditUrl   xcgui.HXCGUI
	hMenuBar   xcgui.HXCGUI
	hState     xcgui.HXCGUI

	rtBody xcgui.RECT
)

func OnWndDrawWindow(hDraw xcgui.HDRAW, pbHandled bool) int {
	xcgui.XEle_RedrawEle(xcgui.HELE(hWeb))
	return 0
}

func main() {
	// xcgui.XC_LoadResource("xml/resource.xml", "")
	hxcgui := xcgui.XC_LoadLayout("xml/layout.xml", 0)
	hWindow = xcgui.HWINDOW(hxcgui)

	xcgui.XWnd_RegEventC(hWindow, xcgui.WM_PAINT, xcgui.CallBack(OnWndDrawWindow))

	// 后退
	hGoBack = xcgui.XC_GetObjectByID(hWindow, 130)
	// xcgui.XEle_RegEventC1(xcgui.HELE(hGoBack), xcgui.XE_BNCLICK, xcgui.CallBack(onBtnClick))
	xcgui.XEle_RegEventC1(xcgui.HELE(hGoBack),
		xcgui.XE_BNCLICK,
		xcgui.CallBackGo(func() {
			fmt.Println("后退中...")
			web.XWeb_GoBack(hWeb)
		}))

	// url地址栏
	hEditUrl = xcgui.XC_GetObjectByID(hWindow, 131)
	xcgui.XEle_EnableSwitchFocus(xcgui.HELE(hEditUrl), true)
	xcgui.XEle_RegEventC(xcgui.HELE(hEditUrl), xcgui.XE_CHAR, xcgui.CallBack(onEventChar))
	// 解决焦点丢失问题
	xcgui.XEle_RegEventC(xcgui.HELE(hEditUrl), xcgui.XE_LBUTTONDOWN, xcgui.CallBack(onLButtonDown))

	// 前进
	hGoForward = xcgui.XC_GetObjectByID(hWindow, 132)
	xcgui.XEle_RegEventC1(xcgui.HELE(hGoForward), xcgui.XE_BNCLICK, xcgui.CallBack(onBtnClick))

	// 浏览器body区
	xcgui.XWnd_GetBodyRect(hWindow, &rtBody)

	// 更多
	hMenuBar = xcgui.XC_GetObjectByID(hWindow, 133)
	xcgui.XMenuBar_AddButton(xcgui.HELE(hMenuBar), "更多")
	hMenu = xcgui.XMenuBar_GetMenu(xcgui.HELE(hMenuBar), 0)

	if hMenu != 0 {
		xcgui.XMenu_AddItem(hMenu, 1330, "主页", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 1331, "重新载入", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 1332, "停止载入", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 1333, "代理加载谷歌", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 1334, "放大", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 1335, "缩小", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 1336, "恢复缩放", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 1337, "获取Cookie", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 1338, "js调用vs函数", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 1340, "设置网页可编辑", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 1341, "设置网页不可编辑", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 1351, "执行js脚本", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 1352, "取浏览器状态", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 1353, "js模拟填写", 0, 0)
		xcgui.XMenu_AddItem(hMenu, 1360, "关于", 0, 0)

	}
	// 菜单选择
	xcgui.XEle_RegEventC(xcgui.HELE(hMenuBar), xcgui.XE_MENU_SELECT, xcgui.CallBack(onWndMenuSelect))

	web.XWeb_Init()
	hWeb = web.XWeb_Create(
		int(rtBody.Left),
		int(rtBody.Top),
		int(rtBody.Right-rtBody.Left),
		int(rtBody.Bottom-rtBody.Top),
		hxcgui)

	// 默认打开百度
	web.XWeb_LoadUrl(hWeb, "http://www.baidu.com")

	// // URL变化
	// web.XWeb_OnURLChanged(hWeb, xcgui.CallBack(onURLChanged), 1)
	//
	// 加载完成
	web.XWeb_OnLoadingFinish(hWeb, xcgui.CallBack(onLoadingFinishCallback), 0)

	// // Title 变化修改状态栏
	// hState = xcgui.XC_GetObjectByID(hWindow, 311)
	// web.XWeb_OnTitleChanged(hWeb, xcgui.CallBack(onTitleChanged), 1)

	// 窗口变化web窗体自适应
	xcgui.XWnd_RegEventC(hWindow, xcgui.WM_SIZE, xcgui.CallBack(OnWndSize))

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	web.XWeb_UnInit()
	xcgui.XExitXCGUI()
}

func onBtnClick(hEventEle xcgui.HELE, pbHandled *bool) int {
	switch hEventEle {
	case xcgui.HELE(hGoBack):
		web.XWeb_GoBack(hWeb)
	case xcgui.HELE(hGoForward):
		web.XWeb_GoForward(hWeb)
	}

	return 0
}

func onEventChar(wParam, lParam uintptr, pbHandled *bool) int {
	// fmt.Println(string(wParam), lParam)
	if wParam == VK_RETURN {
		url, _ := xcgui.XRichEdit_GetTextGo(xcgui.HELE(hEditUrl))
		fmt.Println("----------------------------------------------")
		fmt.Println("URL地址栏: ", url)
		web.XWeb_LoadUrl(hWeb, url)
	}

	return 0
}

func OnWndSize(nFlags uint32, pSize *xcgui.SIZE, pbHandled *bool) int {
	xcgui.XWnd_GetBodyRect(hWindow, &rtBody)
	xcgui.XEle_SetRectEx(xcgui.HELE(hWeb),
		int(rtBody.Left),
		int(rtBody.Top),
		int(rtBody.Right-rtBody.Left),
		int(rtBody.Bottom-rtBody.Top),
		true)

	return 0
}

func onURLChanged(webView, param, url uintptr) int {
	openUrl := web.XWeb_GetStringW(url)
	// xcgui.XRichEdit_SetText(xcgui.HELE(hEditUrl), openUrl)
	fmt.Println("正在访问: ", openUrl)

	return 0
}

func onLoadingFinishCallback(webView, param, url uintptr, result web.WkeLoadingResult, failedReason uintptr) int {
	fmt.Println("onLoadingFinishCallback ----->", web.XWeb_GetStringW(url))
	return 0
}

func onTitleChanged(webView, param, title uintptr) int {
	fmt.Println("浏览器标题: ", web.XWeb_GetStringA(title))
	xcgui.XShapeText_SetText(hState, web.XWeb_GetStringA(title))
	xcgui.XWnd_RedrawWnd(hWindow)

	return 0
}

func onLButtonDown(nFlags uint32, pPt *xcgui.POINT, pbHandled *bool) int {
	win.SetFocus(win.HWND(xcgui.XWnd_GetHWND(hWindow)))

	return 0
}

func onWndMenuSelect(nID int, pbHandled *bool) int {
	switch nID {
	case 1330:
		web.XWeb_LoadUrl(hWeb, "http://www.hupu.cn")
	case 1331:
		web.XWeb_Reload(hWeb)
	case 1332:
		web.XWeb_StopLoading(hWeb)
	case 1333:
		// web.XWeb_SetProxy(hWeb, proxyType, hostName, port, username, pwd)
		xcgui.MessageBox(xcgui.XWnd_GetHWND(hWindow), "代理", "请在源码中设置代理服务器.", xcgui.MB_ICONWARNING)
	case 1334:
		zom := web.XWeb_GetZoom(hWeb)
		zom += 0.3
		web.XWeb_Zoom(hWeb, zom)
	case 1335:
		zom := web.XWeb_GetZoom(hWeb)
		zom -= 0.3
		web.XWeb_Zoom(hWeb, zom)
	case 1336:
		web.XWeb_ZoomReset(hWeb)
	case 1337:
		pCookie := web.XWeb_GetCookie(hWeb)
		xcgui.MessageBox(xcgui.XWnd_GetHWND(hWindow),
			"Cookie", pCookie, xcgui.MB_ICONINFORMATION)
	case 1338:
		web.XWeb_LoadUrl(hWeb, "file://./jsCallFunction.html")
	case 1340:
		web.XWeb_SetEditable(hWeb, true)
	case 1341:
		web.XWeb_SetEditable(hWeb, false)
	case 1351:
		web.XWeb_RunJs(hWeb, "javascript: alert('JS: \n    炫彩界面库-golang.')")
	case 1352:
		if web.XWeb_IsDocumentReady(hWeb) {
			SetStat("网页加载完毕!")
		} else if web.XWeb_IsLoadingCompleted(hWeb) {
			SetStat("网页加载完成!")
		} else if web.XWeb_IsLoadingFailed(hWeb) {
			SetStat("网页加载失败!")
		}
	case 1353:
		fmt.Println("js模拟填写,还未完成.")
	case 1360:
		version := web.XWeb_GetVersionString()
		xcgui.MessageBox(xcgui.XWnd_GetHWND(hWindow), "版本", version, xcgui.MB_ICONINFORMATION)
	}

	return 0
}

func SetStat(str string) {
	xcgui.XShapeText_SetText(hState, str)
	xcgui.XWnd_RedrawWnd(hWindow)
}
