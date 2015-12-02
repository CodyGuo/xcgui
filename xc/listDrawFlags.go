package xc

type List_drawItemBk_flags_ uint32

/*
List,ListBox,Tree,项背景绘制标志位
	list_drawItemBk_flags_nothing 不绘制
	list_drawItemBk_flags_leave 绘制鼠标离开状态项背景
	list_drawItemBk_flags_stay 绘制鼠标选择状态项背景
	list_drawItemBk_flags_select 绘制鼠标停留状态项项背景
*/
const (
	LIST_DRAWITEMBK_FLAGS_NOTHING List_drawItemBk_flags_ = 0x000
	LIST_DRAWITEMBK_FLAGS_LEAVE   List_drawItemBk_flags_ = 0x001
	LIST_DRAWITEMBK_FLAGS_STAY    List_drawItemBk_flags_ = 0x002
	LIST_DRAWITEMBK_FLAGS_SELECT  List_drawItemBk_flags_ = 0x004
)

type Listbox_item_info_i struct {
	//用户绑定数据
	nUserData  int
	nHeight    int
	nSelHeight int
}

type Listview_item_id_i struct {
	//<组索引
	iGroup int
	//<项索引
	iItem int
}
