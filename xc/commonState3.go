package xc

type common_state3_ uint32

/* 普通三种状态
   common_state3_leave 离开
   common_state3_stay 停留
   common_state3_down 按下
*/

const (
    COMMON_STATE3_LEAVE common_state3_ = iota
    COMMON_STATE3_STAY
    COMMON_STATE3_DOWN
)
