package main

import (
	"fmt"
	"log"
)

import (
	xcgui "github.com/codyguo/xcgui/xc"
	"github.com/lxn/walk"
)

var (
	hWindow xcgui.HWINDOW
	ni      *walk.NotifyIcon
)

func main() {
	hWindow = xcgui.XWndCreate(0, 0, 0, 0, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_NOTHING)

	// We load our icon from a file.
	icon, err := walk.NewIconFromFile("../../img/x.ico")
	if err != nil {
		log.Fatal(err)
	}

	// Create the notify icon and make sure we clean it up on exit.
	ni, err = walk.NewNotifyIcon()
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
		if button == walk.LeftButton {
			if err := ni.ShowCustom(
				"自定义消息",
				"这是一个带图标的自定义消息."); err != nil {
				log.Fatal(err)
			}
		} else {
			RunXCMenu()
		}

	})

	// 托盘图标默认为隐藏状态，需设置为显示。
	if err := ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}

	// Now that the icon is visible, we can bring up an info balloon.
	if err := ni.ShowInfo("托盘", "正在运行中."); err != nil {
		log.Fatal(err)
	}

	// 注册菜单选择事件
	xcgui.XWndRegEventC(hWindow, xcgui.XWM_MENU_SELECT, xcgui.CallBack(OnWndMenuSelect))

	// Run the message loop.
	xcgui.XRunXCGUI()
}

func RunXCMenu() {
	var p xcgui.POINT
	xcgui.GetCursorPos(&p)
	hMenu := xcgui.XMenuCreate()
	xcgui.XMenuAddItem(hMenu, 101, "炫彩菜单101", 0, xcgui.MENU_STATE_FLAGS_POPUP)
	xcgui.XMenuAddItem(hMenu, 102, "炫彩菜单102", 0, 0)
	xcgui.XMenuAddItem(hMenu, 103, "炫彩菜单103", 0, 0)
	xcgui.XMenuAddItem(hMenu, 104, "炫彩菜单104", 0, 0)
	xcgui.XMenuAddItem(hMenu, 105, "炫彩菜单105", 0, 0)
	xcgui.XMenuAddItem(hMenu, 106, "退出程序", 0, 0)

	xcgui.XMenuAddItem(hMenu, 1011, "炫彩菜单101-1", 101, 0)
	xcgui.XMenuAddItem(hMenu, 1012, "炫彩菜单101-2", 101, 0)

	xcgui.SetForegroundWindow(xcgui.XWndGetHWND(hWindow))
	xcgui.XMenuPopup(hMenu, xcgui.XWndGetHWND(hWindow), p.X, p.Y, 0, 0)
}

func OnWndMenuSelect(nID int, pBool bool) int {
	switch nID {
	case 1011:
		ni.ShowInfo("提示信息", "炫彩菜单101-1")
		xcgui.MessageBox(xcgui.XWndGetHWND(hWindow), "提示信息", "炫彩菜单101-1", xcgui.MB_ICONINFORMATION)
	case 1012:
		ni.ShowWarning("警告信息", "炫彩菜单101-2")
		xcgui.MessageBox(xcgui.XWndGetHWND(hWindow), "警告信息", "炫彩菜单101-2", xcgui.MB_ICONWARNING)
	case 102:
		ni.ShowError("错误信息", "炫彩菜单102")
		xcgui.MessageBox(xcgui.XWndGetHWND(hWindow), "错误信息", "炫彩菜单102", xcgui.MB_ICONERROR)
	case 106:
		ni.ShowMessage("退出程序", "正在退出程序...")
		ni.Dispose()
		walk.App().Exit(0)
		xcgui.XExitXCGUI()
	default:
		ni.ShowMessage("其他信息", "您选择的菜单："+fmt.Sprint(nID))
		xcgui.MessageBox(xcgui.XWndGetHWND(hWindow), "其他信息", "您选择了其他菜单.", xcgui.MB_USERICON)
	}

	return 0
}
