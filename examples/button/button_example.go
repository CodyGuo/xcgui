package main

import (
    "log"
)
import (
    "github.com/codyguo/xcgui"
    // "github.com/codyguo/xcgui/xc"
)

func main() {
    mw, _ := xcgui.NewMainWindow(xcgui.Size{800, 600}, "测试")
    mw.SetMinimumSize(xcgui.Size{500, 300})

    button1, _ := xcgui.NewButton(mw, xcgui.Rectangle{300, 180, 100, 25})
    button1.SetText("提示信息")

    button2, _ := xcgui.NewButton(mw, xcgui.Rectangle{300, 235, 100, 30})
    button2.SetText("警告信息")

    button1.OnBtnClick(func() {
        stateEx := button1.GetstateEx()
        log.Println("statex:", stateEx)
        style := button1.GetStyle()
        log.Println("button1-style:", style)

        align := button1.GetTextAlign()
        log.Println("对齐方式:", align)
        xcgui.MsgBox(mw, "提示", "这是一条提示信息...", xcgui.MsgBoxIconInformation)
    })

    button2.OnBtnClick(func() {
        state := button2.GetState()
        log.Println("state:", state)
        // button2.SetType(xc.BUTTON_TYPE_RADIO)
        button1.SetTextAlign(xcgui.BUTTON_FLAGS_LEFT_BOTTOM)

        text := button2.GetText()
        log.Println("button1 - text:", text)
        xcgui.MsgBox(mw, "警告", "这是一条警告信息!", xcgui.MsgBoxIconWarning)
    })

    img, _ := xcgui.NewImageFromFile("../img/open.png")

    button1.SetIcon(img)
    button2.SetIcon(img)

    button1.SetIconAlign(xcgui.BUTTON_ICON_ALIGN_TOP)
    button2.SetIconAlign(xcgui.BUTTON_ICON_ALIGN_RIGHT)

    button1.SetOffset(xcgui.Point{5, 0})

    // button2.SetChecked(true)

    // if button2.Checked() {
    //     log.Println("button2 is checked.")
    //     button2.SetChecked(false)
    // }

    // button2.SetState(xc.COMMON_STATE3_DOWN)

    err := mw.Show()
    if err != nil {
        log.Fatalln(err)
    }

    mw.Run()

    mw.Close()
}
