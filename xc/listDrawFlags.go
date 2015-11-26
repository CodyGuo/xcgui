package xc

type LIST_DRAWITEMBK_FLAGS_ uint32

/* List,ListBox,Tree,项背景绘制标志位
list_drawItemBk_flags_nothing 不绘制
list_drawItemBk_flags_leave 绘制鼠标离开状态项背景
list_drawItemBk_flags_stay 绘制鼠标选择状态项背景
list_drawItemBk_flags_select 绘制鼠标停留状态项项背景
*/
const (
	LIST_DRAWITEMBK_FLAGS_NOTHING LIST_DRAWITEMBK_FLAGS_ = 0x000
	LIST_DRAWITEMBK_FLAGS_LEAVE   LIST_DRAWITEMBK_FLAGS_ = 0x001
	LIST_DRAWITEMBK_FLAGS_STAY    LIST_DRAWITEMBK_FLAGS_ = 0x002
	LIST_DRAWITEMBK_FLAGS_SELECT  LIST_DRAWITEMBK_FLAGS_ = 0x004
)

type LISTBOX_ITEM_INFO_I struct {
	nUserData  int //用户绑定数据
	nHeight    int
	nSelHeight int
}
