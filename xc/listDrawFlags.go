package xc

type list_drawItemBk_flags_ uint32

/* List,ListBox,Tree,项背景绘制标志位
   list_drawItemBk_flags_nothing 不绘制
   list_drawItemBk_flags_leave 绘制鼠标离开状态项背景
   list_drawItemBk_flags_stay 绘制鼠标选择状态项背景
   list_drawItemBk_flags_select 绘制鼠标停留状态项项背景
*/
const (
    list_drawItemBk_flags_nothing list_drawItemBk_flags_ = 0x000
    list_drawItemBk_flags_leave   list_drawItemBk_flags_ = 0x001
    list_drawItemBk_flags_stay    list_drawItemBk_flags_ = 0x002
    list_drawItemBk_flags_select  list_drawItemBk_flags_ = 0x004
)
