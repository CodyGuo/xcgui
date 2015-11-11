package xcgui

import (
    "github.com/codyguo/xcgui/xc"
)

type Button struct {
    WindowBase

    clickedPublisher EventPublisher
}

func (b *Button) WndProc(hwnd xc.HWND, msg uint32, wParam, lParam uintptr) uintptr {
    return 0
}

func NewButton(parent Window, rect Rectangle) (*Button, error) {
    btn := new(Button)

    btn.hEle = xc.XBtnCreate(
        rect.X,
        rect.Y,
        rect.Width,
        rect.Height,
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
//     return b.clickedPublisher.Event(b.hEle)
// }

func (b *Button) Clicked(function xc.CallBack) {
    xc.XEleRegEventC(
        b.hEle,
        xc.XE_BNCLICK,
        function)
}
