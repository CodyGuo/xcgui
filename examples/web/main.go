package main

import (
	"fmt"
)
import (
	web "github.com/codyguo/xcgui/wke"
	xcgui "github.com/codyguo/xcgui/xc"
)

var (
	hWeb xcgui.HWINDOW
)

func main() {
	hWindow := xcgui.XWnd_Create(0, 0, 820, 700, "测试", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	button := xcgui.XBtn_Create(300, 35, 50, 30, "执行js", xcgui.HXCGUI(hWindow))

	hWeb = web.XWeb_Create(20, 100, 750, 550, xcgui.XWnd_GetHWND(hWindow))
	web.XWeb_SetCookieEnabled(hWeb, true)

	version := web.XWeb_GetVersionString()
	fmt.Println(version)

	web.XWeb_SetUserAgentA(hWeb, "Mozilla/5.0 (Windows NT 12.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Codyguo/45.0.2454.101 Safari/537.36")

	version2 := web.XWeb_GetVersion()
	fmt.Println(version2)

	web.XWeb_LoadUrl(hWeb, "http://www.xcgui.com/bbs/forum.php")

	xcgui.XEle_RegEventC1(button, xcgui.XE_BNCLICK, xcgui.CallBack(RunJS))
	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func RunJS() int {
	str := web.XWeb_GetCookie(hWeb)
	ok := web.XWeb_IsCookieEnabled(hWeb)
	fmt.Println(str, ok)

	web.XWeb_GoBack(hWeb)
	js := `alert(navigator.userAgent +'\n golang wke Run JS. \n` + str + `');`
	web.XWeb_RunJs(hWeb, "javascript: "+js)

	return 0
}
