package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

func main() {
	xcgui.XC_LoadResource("resource.xml", "")

	hxCGUI := xcgui.XC_LoadLayout("layout.xml", 0)

	xcgui.XWnd_AdjustLayout(xcgui.HWINDOW(hxCGUI))
	xcgui.XWnd_ShowWindow(xcgui.HWINDOW(hxCGUI), xcgui.SW_SHOW)

	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
