package main

import (
	xcgui "github.com/codyguo/xcgui/xc"
)

type my_image_info struct {
	hImage1 xcgui.HIMAGE
	hImage2 xcgui.HIMAGE
	hImage3 xcgui.HIMAGE
	hImage4 xcgui.HIMAGE
	hImage5 xcgui.HIMAGE
	hImage6 xcgui.HIMAGE
}

func main() {
	hWindow := xcgui.XWnd_Create(0, 0, 300, 200, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	info := new(my_image_info)

	info.hImage1 = xcgui.XImage_LoadFile("../../img/1.png", false)
	info.hImage2 = xcgui.XImage_LoadFile("../../img/2.png", false)
	info.hImage3 = xcgui.XImage_LoadFile("../../img/3.png", false)
	info.hImage4 = xcgui.XImage_LoadFile("../../img/4.png", false)
	info.hImage5 = xcgui.XImage_LoadFile("../../img/5.png", false)
	info.hImage6 = xcgui.XImage_LoadFile("../../img/6.png", false)

	left, top := 20, 50
	for i := 0; i < 5; i++ {
		CreateButton(left, top, info, hWindow)
		left += 50
	}

	left, top = 20, 100
	for i := 0; i < 5; i++ {
		CreateButtonLoop(left, top, info, hWindow)
		left += 50
	}

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()

}

func CreateButton(left, top int, pInfo *my_image_info, hWindow xcgui.HWINDOW) {
	hButton := xcgui.XBtn_Create(left, top, 37, 42, "A", xcgui.HXCGUI(hWindow))
	xcgui.XBtn_AddAnimationFrame(hButton, pInfo.hImage1, 100)
	xcgui.XBtn_AddAnimationFrame(hButton, pInfo.hImage2, 100)
	xcgui.XBtn_AddAnimationFrame(hButton, pInfo.hImage3, 100)
	xcgui.XBtn_AddAnimationFrame(hButton, pInfo.hImage4, 100)
	xcgui.XBtn_AddAnimationFrame(hButton, pInfo.hImage5, 100)
	xcgui.XBtn_AddAnimationFrame(hButton, pInfo.hImage6, 100)

	xcgui.XBtn_EnableAnimation(hButton, true, false)
}

func CreateButtonLoop(left, top int, pInfo *my_image_info, hWindow xcgui.HWINDOW) {
	hButton := xcgui.XBtn_Create(left, top, 37, 42, "B", xcgui.HXCGUI(hWindow))
	xcgui.XBtn_AddAnimationFrame(hButton, pInfo.hImage1, 100)
	xcgui.XBtn_AddAnimationFrame(hButton, pInfo.hImage2, 100)
	xcgui.XBtn_AddAnimationFrame(hButton, pInfo.hImage3, 100)
	xcgui.XBtn_AddAnimationFrame(hButton, pInfo.hImage4, 100)
	xcgui.XBtn_AddAnimationFrame(hButton, pInfo.hImage5, 100)
	xcgui.XBtn_AddAnimationFrame(hButton, pInfo.hImage6, 100)

	xcgui.XBtn_EnableAnimation(hButton, true, true)
}
