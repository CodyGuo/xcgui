package xc

type ComboBox_state_ uint32

/* ComboBox状态
comboBox_state_leave 鼠标离开状态
comboBox_state_stay 鼠标停留状态
comboBox_state_down 按下状态
*/
const (
	COMBOBOX_STATE_LEAVE ComboBox_state_ = iota
	COMBOBOX_STATE_STAY
	COMBOBOX_STATE_DOWN
)
