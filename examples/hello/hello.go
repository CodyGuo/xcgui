/**
直接使用api
*/
package main

import (
	"fmt"
	xcgui "github.com/CodyGuo/xcgui/xc"
	"syscall"
)

func main() {
	xcgui.XInitXCGUI()
	hwnd := xcgui.XWnd_Create(400, 200, 300, 200, "标题", xcgui.NULL, xcgui.XC_WINDOW_STYLE_DEFAULT)

	parent := xcgui.HXCGUI(hwnd)
	//button
	btn := xcgui.XBtn_Create(10, 5, 80, 22, "关闭", parent)
	xcgui.XBtn_SetType(btn, xcgui.BUTTON_TYPE_CLOSE)
	//监听btn事件
	xcgui.XEle_RegEventC(btn, xcgui.XE_BNCLICK, syscall.NewCallback(OnBtnClick))
	//label
	lb := xcgui.XShapeText_Create(50, 100, 100, 22, "hello world!", parent)
	xcgui.XShapeText_SetText(lb, "hello 世界!")
	xcgui.XShapeText_SetTextColor(lb, 0xff0000, 255)

	//取text及长度
	str := xcgui.XShapeText_GetTextGo(lb)
	fmt.Println(str)
	fmt.Println(xcgui.XShapeText_GetTextLength(lb))

	xcgui.XWnd_ShowWindow(hwnd, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}

func OnBtnClick() int {
	fmt.Println("你点了按钮")
	return 0
}
