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

    button1, _ := xcgui.NewButton(mw, xcgui.Rectangle{300, 180, 100, 25})
    button1.SetText("提示信息")

    button2, _ := xcgui.NewButton(mw, xcgui.Rectangle{300, 235, 100, 30})
    button2.SetText("警告信息")

    button1.OnBtnClick(func() {
        xcgui.MsgBox(mw, "提示", "这是一条提示信息...", xcgui.MsgBoxIconInformation)
    })

    button2.OnBtnClick(func() {
        xcgui.MsgBox(mw, "警告", "这是一条警告信息!", xcgui.MsgBoxIconWarning)
    })

    err := mw.Show()
    if err != nil {
        log.Fatalln(err)
    }

    mw.Run()

    mw.Close()
}
