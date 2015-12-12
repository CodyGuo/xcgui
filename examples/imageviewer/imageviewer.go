package main

import (
	"log"
)

import (
	xcgui "github.com/codyguo/xcgui/xc"
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

func main() {
	mw := new(MyMainWindow)

	mw.hWindow = xcgui.XWndCreate(0, 0, 800, 600, "pTitle", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(mw.hWindow)

	hButton := xcgui.XBtnCreate(50, 100, 100, 30, "浏览图片", xcgui.HXCGUI(mw.hWindow))
	hButtonAbout := xcgui.XBtnCreate(200, 100, 100, 30, "关于", xcgui.HXCGUI(mw.hWindow))

	xcgui.XEleRegEventC1(hButton, xcgui.XE_BNCLICK, xcgui.CallBack(mw.openAction_Triggered))
	xcgui.XEleRegEventC1(hButtonAbout, xcgui.XE_BNCLICK, xcgui.CallBack(mw.aboutAction_Triggered))

	xcgui.XWndShowWindow(mw.hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()

}

type MyMainWindow struct {
	hWindow xcgui.HWINDOW
	*walk.MainWindow
	prevFilePath string
}

func (mw *MyMainWindow) openAction_Triggered() int {
	if err := mw.openImage(); err != nil {
		log.Print(err)
	}

	return 0
}

func (mw *MyMainWindow) openImage() error {
	dlg := new(walk.FileDialog)

	dlg.FilePath = mw.prevFilePath
	dlg.Filter = "Image Files (*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff)|*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff"
	dlg.Title = "选择图片"

	if ok, err := dlg.ShowOpen(mw); err != nil {
		return err
	} else if !ok {
		return nil
	}

	mw.prevFilePath = dlg.FilePath

	return nil
}

func (mw *MyMainWindow) aboutAction_Triggered() int {
	walk.MsgBox(mw, "关于", "walk 的msgbox 与 xcgui的混写."+mw.prevFilePath, walk.MsgBoxIconInformation)

	return 0
}

func (mw *MyMainWindow) Handle() win.HWND {
	return win.HWND(xcgui.XWndGetHWND(mw.hWindow))
}
