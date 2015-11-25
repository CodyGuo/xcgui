package main

import (
	"github.com/codyguo/xcgui"

	"github.com/codyguo/xcgui/xc"
)

func main() {
	mw, _ := xcgui.NewMainWindow(xcgui.Size{300, 200}, "富文本")

	hEle := xc.XRichEditCreate(20, 40, 150, 120, xc.HXCGUI(mw.GetHWindow()))
	xc.XRichEditEnableAutoWrap(hEle, true)

	mw.Show()
	mw.Run()
}
