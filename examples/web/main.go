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
	hWindow := xcgui.XWndCreate(0, 0, 820, 700, "测试", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	button := xcgui.XBtnCreate(300, 35, 50, 30, "执行js", xcgui.HXCGUI(hWindow))

	hWeb = web.XWebCreate(20, 100, 750, 550, xcgui.XWndGetHWND(hWindow))
	web.XWebSetCookieEnabled(hWeb, true)

	version := web.XWebGetVersionString()
	fmt.Println(version)

	web.XWebSetUserAgentA(hWeb, "Mozilla/5.0 (Windows NT 12.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Codyguo/45.0.2454.101 Safari/537.36")

	version2 := web.XWebGetVersion()
	fmt.Println(version2)

	web.XWebLoadUrl(hWeb, "http://www.xcgui.com/bbs/forum.php")

	xcgui.XEleRegEventC1(button, xcgui.XE_BNCLICK, xcgui.CallBack(RunJS))
	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func RunJS() int {
	str := web.XWebGetCookie(hWeb)
	ok := web.XWebIsCookieEnabled(hWeb)
	fmt.Println(str, ok)

	web.XWebGoBack(hWeb)
	js := `alert(navigator.userAgent +'\n golang wke Run JS. \n` + str + `');`
	web.XWebRunJs(hWeb, "javascript: "+js)

	return 0
}
