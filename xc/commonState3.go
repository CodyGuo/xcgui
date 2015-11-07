package xc

type common_state3_ uint32

/* 普通三种状态
   common_state3_leave 离开
   common_state3_stay 停留
   common_state3_down 按下
*/

const (
    common_state3_leave common_state3_ = iota
    common_state3_stay
    common_state3_down
)
