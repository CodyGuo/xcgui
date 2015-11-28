package xc

type LIST_ITEM_STATE_ uint32

/*
List项状态
	list_item_state_leave 项鼠标离开状态
	list_item_state_stay 项鼠标停留状态
	list_item_state_select 项选择状态
*/
const (
	LIST_ITEM_STATE_LEAVE LIST_ITEM_STATE_ = iota
	LIST_ITEM_STATE_STAY
	LIST_ITEM_STATE_SELECT
)
