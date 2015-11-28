package xc

type COMMON_STATE3_ uint32

/*
common_state3_ 普通三种状态
	common_state3_leave 离开
	common_state3_stay 停留
	common_state3_down 按下
*/

const (
	COMMON_STATE3_LEAVE COMMON_STATE3_ = iota
	COMMON_STATE3_STAY
	COMMON_STATE3_DOWN
)
