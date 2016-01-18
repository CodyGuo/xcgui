package xcgui

import (
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

const (
	BUTTON_STATE_LEAVE   = xc.BUTTON_STATE_LEAVE
	BUTTON_STATE_STAY    = xc.BUTTON_STATE_STAY
	BUTTON_STATE_DOWN    = xc.BUTTON_STATE_DOWN
	BUTTON_STATE_CHECK   = xc.BUTTON_STATE_CHECK
	BUTTON_STATE_DISABLE = xc.BUTTON_STATE_DISABLE
)

type Button struct {
	WindowBase

	clickedPublisher EventPublisher
}

func NewButton(parent Window, rect Rectangle) (*Button, error) {
	btn := new(Button)

	btn.hEle = xc.XBtn_Create(
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
	return xc.XBtn_IsCheck(b.hEle)
}

func (b *Button) SetChecked(checked bool) {
	if checked == b.Checked() {
		return
	}

	xc.XBtn_SetCheck(b.hEle, checked)
}

func (b *Button) SetStyle(nStyle xc.Button_style_) {
	xc.XBtn_SetStyle(b.hEle, nStyle)
}

func (b *Button) GetStyle() xc.Button_style_ {
	return xc.XBtn_GetStyle(b.hEle)
}

func (b *Button) SetState(nState xc.Common_state3_) {
	xc.XBtn_SetState(b.hEle, nState)
}

func (b *Button) GetState() xc.Common_state3_ {
	return xc.XBtn_GetState(b.hEle)
}

func (b *Button) GetstateEx() xc.Button_state_ {
	return xc.XBtn_GetStateEx(b.hEle)
}

func (b *Button) SetType(nType xc.Button_type_) {
	xc.XBtn_SetType(b.hEle, nType)
}

func (b *Button) GetType() xc.Button_type_ {
	return xc.XBtn_GetType(b.hEle)
}

func (b *Button) SetGroupID(nID int) {
	xc.XBtn_SetGroupID(b.hEle, nID)
}

func (b *Button) GetGroupID() int {
	return xc.XBtn_GetGroupID(b.hEle)
}

func (b *Button) SetBindEle(hBindEle xc.HELE) {
	xc.XBtn_SetBindEle(b.hEle, hBindEle)
}

func (b *Button) GetBindEle() xc.HELE {
	return xc.XBtn_GetBindEle(b.hEle)
}

func (b *Button) SetTextAlign(nFlags int) {
	xc.XBtn_SetTextAlign(b.hEle, nFlags)
}

func (b *Button) GetTextAlign() int {
	return xc.XBtn_GetTextAlign(b.hEle)
}

func (b *Button) SetIconAlign(align xc.Button_icon_align_) {
	xc.XBtn_SetIconAlign(b.hEle, align)
}

func (b *Button) SetOffset(point Point) {
	xc.XBtn_SetOffset(b.hEle, point.X, point.Y)
}

func (b *Button) SetOffsetIcon(point Point) {
	xc.XBtn_SetOffsetIcon(b.hEle, point.X, point.Y)
}

func (b *Button) SetIconSpace(sizeNum int) {
	xc.XBtn_SetIconSpace(b.hEle, sizeNum)
}

func (b *Button) SetText(value string) {
	xc.XBtn_SetText(b.hEle, value)
}

func (b *Button) GetText() string {
	// textLength := xc.SendMessage(xc.HWND(b.hEle), xc.WM_GETTEXTLENGTH, 0, 0)
	// pSize := make([]uint16, 256)
	// xc.XEleGetContentSize(b.hEle, &pSize[0])
	// textLength := pSize[0]
	// fmt.Println("textLength:", textLength)

	pSize := new(xc.SIZE)
	xc.XEle_GetContentSize(b.hEle, pSize)
	textLength := pSize.CX

	buf := make([]uint16, textLength+1)
	xc.XBtn_GetText(b.hEle, &buf[0], int(textLength)+1)

	return syscall.UTF16ToString(buf)
}

func (b *Button) SetIcon(hImage xc.HIMAGE) {
	xc.XBtn_SetIcon(b.hEle, hImage)
}

func (b *Button) SetIconDisable(hImage xc.HIMAGE) {
	xc.XBtn_SetIconDisable(b.hEle, hImage)
}

func (b *Button) AddAnimationFrame(hImage xc.HIMAGE, uElapse uint32) {
	xc.XBtn_AddAnimationFrame(b.hEle, hImage, uElapse)
}

func (b *Button) EnableAnimation(bEnable, bLoopPlay bool) {
	xc.XBtn_EnableAnimation(b.hEle, bEnable, bLoopPlay)
}

func (b *Button) AddBkBorder(nState xc.Button_state_, color Color, alpha byte, width int) {
	xc.XBtn_AddBkBorder(b.hEle, nState, xc.COLORREF(color), alpha, width)
}

func (b *Button) AddBkFill(nState xc.Button_state_, color Color, alpha byte) {
	xc.XBtn_AddBkFill(b.hEle, nState, xc.COLORREF(color), alpha)
}

func (b *Button) AddBkImage(nState xc.Button_state_, hImage xc.HIMAGE) {
	xc.XBtn_AddBkImage(b.hEle, nState, hImage)
}

func (b *Button) GetBkInfoCount(nState xc.Button_state_) int {
	return xc.XBtn_GetBkInfoCount(b.hEle, nState)
}

func (b *Button) ClearBkInfo(nState xc.Button_state_) {
	xc.XBtn_ClearBkInfo(b.hEle, nState)
}

// func (b *Button) GetBkInfoManager(nState xc.Button_state_) int {
// 	return int(xc.XBtn_GetBkInfoManager(b.hEle, nState))
// }

func (b *Button) OnBtnClick(pFunc func()) {
	var OnBtnClick = func(pbHandled *bool) int {
		// pbHandled = true 取消, false 继续
		// *pbHandled = true
		pFunc()

		return 0
	}

	xc.XEle_RegEventC(
		b.hEle,
		xc.XE_BNCLICK,
		syscall.NewCallback(OnBtnClick))
}

// func (b *Button) Clicked() *Event {
//     return b.clickedPublisher.Event(b.hEle)
// }
