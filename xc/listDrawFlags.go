package xc

type list_drawItemBk_flags_ uint32

/* List,ListBox,Tree,项背景绘制标志位
list_drawItemBk_flags_nothing 不绘制
list_drawItemBk_flags_leave 绘制鼠标离开状态项背景
list_drawItemBk_flags_stay 绘制鼠标选择状态项背景
list_drawItemBk_flags_select 绘制鼠标停留状态项项背景
*/
const (
	LIST_DRAWITEMBK_FLAGS_NOTHING list_drawItemBk_flags_ = 0x000
	LIST_DRAWITEMBK_FLAGS_LEAVE   list_drawItemBk_flags_ = 0x001
	LIST_DRAWITEMBK_FLAGS_STAY    list_drawItemBk_flags_ = 0x002
	LIST_DRAWITEMBK_FLAGS_SELECT  list_drawItemBk_flags_ = 0x004
)
