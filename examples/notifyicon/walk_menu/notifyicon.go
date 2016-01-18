package main

import (
	"log"
)

import (
	xcgui "github.com/codyguo/xcgui/xc"
	"github.com/lxn/walk"
)

func main() {
	xcgui.XWnd_Create(0, 0, 0, 0, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_NOTHING)

	// We load our icon from a file.
	icon, err := walk.NewIconFromFile("../../img/x.ico")
	if err != nil {
		log.Fatal(err)
	}

	// Create the notify icon and make sure we clean it up on exit.
	ni, err := walk.NewNotifyIcon()
	if err != nil {
		log.Fatal(err)
	}
	defer ni.Dispose()

	// Set the icon and a tool tip text.
	if err := ni.SetIcon(icon); err != nil {
		log.Fatal(err)
	}
	if err := ni.SetToolTip("托盘"); err != nil {
		log.Fatal(err)
	}

	// When the left mouse button is pressed, bring up our balloon.
	ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}

		if err := ni.ShowCustom(
			"自定义消息",
			"这是一个带图标的自定义消息."); err != nil {
			log.Fatal(err)
		}
	})

	// 菜单使用walk的，主程序为xcgui.
	exitAction := walk.NewAction()
	if err := exitAction.SetText("退出"); err != nil {
		log.Fatal(err)
	}
	exitAction.Triggered().Attach(func() {
		ni.Dispose()
		walk.App().Exit(0)
		xcgui.XExitXCGUI()
	})

	if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
		log.Fatal(err)
	}

	// 托盘图标默认为隐藏状态，需设置为显示。
	if err := ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}

	// Now that the icon is visible, we can bring up an info balloon.
	if err := ni.ShowInfo("托盘", "正在运行中."); err != nil {
		log.Fatal(err)
	}

	// Run the message loop.
	xcgui.XRunXCGUI()
}
