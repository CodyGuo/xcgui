package main

import (
    "log"
)
import (
    "github.com/codyguo/xcgui"
)

func main() {
    mw := xcgui.NewMainWindow(600, 500, "测试")
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
