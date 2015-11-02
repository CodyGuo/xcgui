package xcgui

import (
    "github.com/codyguo/xcgui/xc"
)

type Button struct {
    X      int
    Y      int
    Width  int
    Height int
    Text   string
    hEle   uintptr

    clickedPublisher EventPublisher
}

func NewButton(x, y, w, h int, text string, owner *MainWindow) *Button {
    btn := new(Button)
    btn.SetBounds(x, y, w, h)
    btn.SetText(text)

    var pHwnd xc.HWND

    if owner != nil {
        pHwnd = owner.GetHWindow()
    }

    btn.hEle = xc.XBtnCreate(btn.X,
        btn.Y,
        btn.Width,
        btn.Height,
        btn.Text,
        pHwnd)

    return btn
}

func (b *Button) SetText(value string) {
    b.Text = value
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
    xc.XEleRegEventCPP(b.hEle,
        xc.XE_BNCLICK,
        function)
}
