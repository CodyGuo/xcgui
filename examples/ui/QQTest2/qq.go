package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	xcgui.XCLoadResource("resource.xml", "")

	hxCGUI := xcgui.XCLoadLayout("layout.xml", 0)

	xcgui.XWndAdjustLayout(xcgui.HWINDOW(hxCGUI))
	xcgui.XWndShowWindow(xcgui.HWINDOW(hxCGUI), xcgui.SW_SHOW)

	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
