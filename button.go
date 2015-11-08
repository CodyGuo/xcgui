package xcgui

import (
    "github.com/codyguo/xcgui/xc"
)

type Button struct {
    WidgetBase
    X      int
    Y      int
    Width  int
    Height int
    Text   string
    hEle   xc.HELE

    clickedPublisher EventPublisher
}

func NewButton(x, y, w, h int, text string, hParent xc.HWINDOW) *Button {
    btn := new(Button)
    btn.SetBounds(x, y, w, h)

    btn.hEle = xc.XBtnCreate(
        btn.X,
        btn.Y,
        btn.Width,
        btn.Height,
        btn.Text,
        xc.HXCGUI(hParent))

    return btn
}

func (b *Button) SetText(value string) {
    xc.XBtnSetText(b.hEle, value)
}

func (b *Button) SetBounds(x, y, w, h int) {
    b.X = x
    b.Y = y
    b.Width = w
    b.Height = h
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
