package main

import (
	"github.com/codyguo/web"
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWndCreate(0, 0, 800, 600, "测试", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hWeb := web.XWebCreate(20, 30, 750, 550, xcgui.XWndGetHWND(hWindow))
	web.XWebLoadUrl(hWeb, "http://wwww.baidu.com")

	xcgui.XWndShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
