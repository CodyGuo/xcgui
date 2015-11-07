package xc

type tree_item_state_ uint32

/* Tree项状态
   tree_item_state_leave 项鼠标离开状态
   tree_item_state_select 项选择状态
*/
const (
    tree_item_state_leave tree_item_state_ = iota
    tree_item_state_select
)
