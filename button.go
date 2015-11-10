package xcgui

import (
    "github.com/codyguo/xcgui/xc"
)

type Button struct {
    WindowBase

    clickedPublisher EventPublisher
}

func NewButton(parent Window, x, y int, size Size) (*Button, error) {
    btn := new(Button)

    btn.hEle = xc.XBtnCreate(
        x,
        y,
        size.Width,
        size.Height,
        "",
        xc.HXCGUI(parent.AsWindowBase().hWindow))

    if btn.hEle == 0 {
        return nil, lastError("XBtnCreate")
    }

    return btn, nil
}

func (b *Button) SetText(value string) {
    xc.XBtnSetText(b.hEle, value)
}

// func (b *Button) Clicked() *Event {
//     return b.clickedPublisher.Event()
// }

func (b *Button) Clicked(function xc.CallBack) {
    xc.XEleRegEventC(
        b.hEle,
        xc.XE_BNCLICK,
        function)
}
