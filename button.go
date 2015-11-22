package xcgui

import (
    "fmt"
    "github.com/codyguo/xcgui/xc"
    "syscall"
)

type Button struct {
    WindowBase

    clickedPublisher EventPublisher
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

    // var OnEventProc = func(hEventEle xc.HELE, nEvent int, wParam, lParam uintptr, pbHandled *xc.BOOL) uintptr {
    //     fmt.Println(nEvent, wParam, lParam)
    //     fmt.Println("事件正常.")
    //     return uintptr(0)
    // }

    // xc.XEleRegEventC(btn.hEle, xc.XE_BNCLICK, syscall.NewCallback(OnEventProc))

    return btn, nil
}

func (b *Button) SetText(value string) {
    xc.XBtnSetText(b.hEle, value)
}

func (b *Button) OnBtnClick(pFunc func()) {
    var (
        OnBtnClick = func(pbHandled *bool) int {
            // pbHandled = true 取消, false 继续
            // *pbHandled = true
            pFunc()

            return 0
        }
    )

    xc.XEleRegEventC(
        b.hEle,
        xc.XE_BNCLICK,
        syscall.NewCallback(OnBtnClick))
}

// func (b *Button) Clicked() *Event {
//     return b.clickedPublisher.Event(b.hEle)
// }
