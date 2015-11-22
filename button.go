package xcgui

import (
    "fmt"
    "syscall"
)

import (
    "github.com/codyguo/xcgui/xc"
)

/* nFlags 文本对齐方式
   BUTTON_FLAGS_LEFT_TOP 水平靠左,垂直靠上
   BUTTON_FLAGS_CENTER_TOP 水平居中,垂直靠上
   BUTTON_FLAGS_RIGHT_TOP 水平靠右,垂直靠上
   BUTTON_FLAGS_LEFT_CENTER 水平靠左,垂直居中
   BUTTON_FLAGS_CENTER_CENTER 水平居中,垂直居中
   BUTTON_FLAGS_RIGHT_CENTER 水平靠右,垂直居中
   BUTTON_FLAGS_LEFT_BOTTOM 水平靠左,垂直靠下
   BUTTON_FLAGS_CENTER_BOTTOM 水平居中,垂直靠下
   BUTTON_FLAGS_RIGHT_BOTTOM 水平靠右,垂直靠下

*/
const (
    BUTTON_FLAGS_LEFT_TOP      = 0
    BUTTON_FLAGS_CENTER_TOP    = 1
    BUTTON_FLAGS_RIGHT_TOP     = 2
    BUTTON_FLAGS_LEFT_CENTER   = 4
    BUTTON_FLAGS_CENTER_CENTER = 5
    BUTTON_FLAGS_RIGHT_CENTER  = 6
    BUTTON_FLAGS_LEFT_BOTTOM   = 8
    BUTTON_FLAGS_CENTER_BOTTOM = 9
    BUTTON_FLAGS_RIGHT_BOTTOM  = 10
)

const (
    BUTTON_ICON_ALIGN_LEFT   = xc.BUTTON_ICON_ALIGN_LEFT
    BUTTON_ICON_ALIGN_TOP    = xc.BUTTON_ICON_ALIGN_TOP
    BUTTON_ICON_ALIGN_RIGHT  = xc.BUTTON_ICON_ALIGN_RIGHT
    BUTTON_ICON_ALIGN_BOTTOM = xc.BUTTON_ICON_ALIGN_BOTTOM
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

func (b *Button) Checked() bool {
    return xc.XBtnIsCheck(b.hEle)
}

func (b *Button) SetChecked(checked bool) {
    if checked == b.Checked() {
        return
    }

    xc.XBtnSetCheck(b.hEle, xc.BoolToBOOL(checked))
}

func (b *Button) SetStyle(nStyle xc.BUTTON_STYLE_) {
    xc.XBtnSetStyle(b.hEle, nStyle)
}

func (b *Button) GetStyle() xc.BUTTON_STYLE_ {
    return xc.XBtnGetStyle(b.hEle)
}

func (b *Button) SetState(nState xc.COMMON_STATE3_) {
    xc.XBtnSetState(b.hEle, nState)
}

func (b *Button) GetState() xc.COMMON_STATE3_ {
    return xc.XBtnGetState(b.hEle)
}

func (b *Button) GetstateEx() xc.BUTTON_STATE_ {
    return xc.XBtnGetStateEx(b.hEle)
}

func (b *Button) SetType(nType xc.BUTTON_TYPE_) {
    xc.XBtnSetType(b.hEle, nType)
}

func (b *Button) GetType() xc.BUTTON_TYPE_ {
    return xc.XBtnGetType(b.hEle)
}

func (b *Button) SetGroupID(nID int) {
    xc.XBtnSetGroupID(b.hEle, nID)
}

func (b *Button) GetGroupID() int {
    return xc.XBtnGetGroupID(b.hEle)
}

func (b *Button) SetBindEle(hBindEle xc.HELE) {
    xc.XBtnSetBindEle(b.hEle, hBindEle)
}

func (b *Button) GetBindEle() xc.HELE {
    return xc.XBtnGetBindEle(b.hEle)
}

func (b *Button) SetTextAlign(nFlags int) {
    xc.XBtnSetTextAlign(b.hEle, nFlags)
}

func (b *Button) GetTextAlign() int {
    return xc.XBtnGetTextAlign(b.hEle)
}

func (b *Button) SetIconAlign(align xc.BUTTON_ICON_ALIGN_) {
    xc.XBtnSetIconAlign(b.hEle, align)
}

func (b *Button) SetOffset(point Point) {
    xc.XBtnSetOffset(b.hEle, point.X, point.Y)
}

func (b *Button) SetOffsetIcon(point Point) {
    xc.XBtnSetOffsetIcon(b.hEle, point.X, point.Y)
}

func (b *Button) SetIconSpace(sizeNum int) {
    xc.XBtnSetIconSpace(b.hEle, sizeNum)
}

func (b *Button) SetText(value string) {
    xc.XBtnSetText(b.hEle, value)
}

func (b *Button) GetText() string {
    textLength := xc.SendMessage(xc.HWND(b.hEle), xc.WM_GETTEXTLENGTH, 0, 0)
    fmt.Println(textLength)
    var buf []string
    buf = make([]string, 4)
    xc.XBtnGetText(b.hEle, &buf[0], 3)
    fmt.Println(buf)

    return buf[0]
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

func (b *Button) SetIcon(hImage xc.HIMAGE) {
    xc.XBtnSetIcon(b.hEle, hImage)
}
