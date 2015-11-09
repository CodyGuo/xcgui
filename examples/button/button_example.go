package main

import (
    "log"
)
import (
    "github.com/codyguo/xcgui"
)

func main() {
    mw, _ := xcgui.NewMainWindow(800, 600, "测试")
    mw.SetMinimumSize(xcgui.Size{500, 300})
    hWindow := mw.GetHWindow()

    button := xcgui.NewButton(260, 200, 50, 25, "提交", hWindow)

    hWnd := mw.Handle()

    var msgFunc = func() uintptr {
        ret := xcgui.MsgBox(hWnd, "警告", "您的点击不正确.", xcgui.MsgBoxIconWarning)
        return uintptr(ret)
    }

    button.Clicked(msgFunc)
    button.SetText("改名字了..")

    err := mw.Show()
    if err != nil {
        log.Fatalln(err)
    }

    mw.Run()

    // defer mw.Close()
}
