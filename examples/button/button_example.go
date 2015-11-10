package main

import (
    "log"
)
import (
    "github.com/codyguo/xcgui"
)

func main() {
    mw, _ := xcgui.NewMainWindow(xcgui.Size{800, 600}, "测试")
    mw.SetMinimumSize(xcgui.Size{500, 300})

    button, _ := xcgui.NewButton(mw, 300, 235, xcgui.Size{50, 30})
    button.SetText("改名字了..")

    button2, _ := xcgui.NewButton(mw, 300, 180, xcgui.Size{40, 25})
    button2.SetText("取消")

    var msgFunc = func() uintptr {
        ret := xcgui.MsgBox(mw, "警告", "您的点击不正确.", xcgui.MsgBoxIconWarning)
        return uintptr(ret)
    }

    button.Clicked(msgFunc)

    err := mw.Show()
    if err != nil {
        log.Fatalln(err)
    }

    mw.Run()

    // defer mw.Close()
}
