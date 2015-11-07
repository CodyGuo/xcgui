package xc

type comboBox_state_ uint32

/* ComboBox状态
   comboBox_state_leave 鼠标离开状态
   comboBox_state_stay 鼠标停留状态
   comboBox_state_down 按下状态
*/
const (
    comboBox_state_leave comboBox_state_ = iota
    comboBox_state_stay
    comboBox_state_down
)
