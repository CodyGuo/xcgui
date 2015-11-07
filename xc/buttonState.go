package xc

type button_state_ uint32

/* 按钮状态
   button_state_leave 离开状态
   button_state_stay 停留状态
   button_state_down 按下状态
   button_state_check 选中状态
   button_state_disable 禁用状态
*/

const (
    button_state_leave button_state_ = iota
    button_state_stay
    button_state_down
    button_state_check
    button_state_disable
)
