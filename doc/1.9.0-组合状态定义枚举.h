
enum  window_state_flag_
{
	window_state_flag_nothing  =0x0000,  ///<无
	window_state_whole_leave   =0x0001,  ///<整个窗口
	window_state_body_leave    =0x0002,  ///<窗口-body
	window_state_top_leave     =0x0004,  ///<窗口-top
	window_state_bottom_leave  =0x0008,  ///<窗口-bottom
	window_state_left_leave    =0x0010,  ///<窗口-left
	window_state_right_leave   =0x0020,  ///<窗口-right
};

enum  element_state_flag_  //组合状态
{
	element_state_flag_nothing   =window_state_flag_nothing,  ///<无
	element_state_flag_enable    =0x0001,  ///<启用
	element_state_flag_disable   =0x0002,  ///<禁用
	element_state_flag_focus     =0x0004,  ///<焦点
	element_state_flag_focus_no  =0x0008,  ///<无焦点

	element_state_flag_leave     =0x0010,  ///<鼠标离开
	element_state_flag_stay      =0x0020,  ///<为扩展模块保留
	element_state_flag_down      =0x0040,  ///<为扩展模块保留
};

enum  button_state_flag_  //组合状态
{
	button_state_flag_leave     =element_state_flag_leave,  ///<鼠标离开
	button_state_flag_stay      =element_state_flag_stay,   ///<鼠标停留
	button_state_flag_down      =element_state_flag_down,   ///<鼠标按下

	button_state_flag_check     =0x0080, ///<选中
	button_state_flag_check_no  =0x0100, ///<未选中
};

enum comboBox_state_flag_
{
	comboBox_state_flag_leave   =element_state_flag_leave, ///<鼠标离开
	comboBox_state_flag_stay    =element_state_flag_stay,  ///<鼠标停留
	comboBox_state_flag_down    =element_state_flag_down,  ///<鼠标按下
};

enum listBox_state_flag_
{
	listBox_state_flag_item_leave  =0x0080, ///<项鼠标离开
	listBox_state_flag_item_stay   =0x0100, ///<项鼠标停留
	listBox_state_flag_item_select =0x0200, ///<项选择
};

enum list_state_flag_
{
	list_state_flag_item_leave  =0x0080, ///<项鼠标离开
	list_state_flag_item_stay   =0x0100, ///<项鼠标停留
	list_state_flag_item_select =0x0200, ///<项选择
};

enum listView_state_flag_
{
	listView_state_flag_item_leave  =0x0080, ///<项鼠标离开
	listView_state_flag_item_stay   =0x0100, ///<项鼠标停留
	listView_state_flag_item_select =0x0200, ///<项选择
};

enum tree_state_flag_
{
	tree_state_flag_item_leave  =0x0080, ///<项鼠标离开
	tree_state_flag_item_select =0x0100, ///<项选择
};
