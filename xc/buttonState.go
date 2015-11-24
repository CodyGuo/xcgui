package xc

type BUTTON_STATE_ uint32

/* button_state_ 按钮状态
button_state_leave 离开状态
button_state_stay 停留状态
button_state_down 按下状态
button_state_check 选中状态
button_state_disable 禁用状态
*/

const (
	BUTTON_STATE_LEAVE BUTTON_STATE_ = iota
	BUTTON_STATE_STAY
	BUTTON_STATE_DOWN
	BUTTON_STATE_CHECK
	BUTTON_STATE_DISABLE
)
