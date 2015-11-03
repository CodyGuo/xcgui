package main

import (
    "log"
)
import (
    "github.com/codyguo/xcgui"
)

func main() {
    mw := xcgui.NewMainWindow(600, 500, "测试")

    button := xcgui.NewButton(260, 200, 50, 25, "提交", mw)

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

    defer mw.Close()
}
