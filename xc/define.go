package xc

// --- 接口句柄类型 ---
type XC_OBJECT_TYPE int32

/*
接口句柄类型
	XC_ERROR 错误类型
	XC_WINDOW 窗口
	XC_MODALWINDOW 模态窗口
	XC_FRAMEWND 框架窗口
	XC_FLOATWND 浮动窗口
*/
const (
	XC_ERROR  XC_OBJECT_TYPE = -1
	XC_WINDOW XC_OBJECT_TYPE = iota
	XC_MODALWINDOW
	XC_FRAMEWND
	XC_FLOATWND
)

/*
接口句柄类型
	XC_ELE 基础元素
	XC_BUTTON 按钮
	XC_RICHEDIT 富文本编辑框
	XC_COMBOBOX 下拉组合框元素
	XC_SCROLLBAR 滚动条元素
	XC_SCROLLVIEW 滚动视图元素
	XC_LIST 列表元素
	XC_LISTBOX 列表框元素
	XC_LISTVIEW 列表视图,大图标
	XC_TREE 树元素
	XC_MENUBAR 菜单条
	XC_SLIDERBAR 滑动条
	XC_PROGRESSBAR 进度条
	XC_TOOLBAR 工具条
	XC_MONTHCAL 月历元素
	XC_DATETIME 时间元素
	XC_PROPERTYGRID 属性网格
	XC_RICHEDIT_COLOR 颜色选择元素
	XC_TABBAR tab条
	XC_TEXTLINK 文本链接按钮
	XC_PANE 布局窗格
	XC_PANE_SPLIT 布局窗格拖动分割条
	XC_MENUBAR_BUTTON 菜单条上的按钮
	XC_TOOLBAR_BUTTON 工具条上按钮
	XC_PROPERTYPAGE_LABEL 属性页标签按钮
	XC_PIER 窗格停靠码头
	XC_BUTTON_MENU 弹出菜单按钮
	XC_VIRTUAL_ELE 虚拟元素
	XC_RICHEDIT_FILE RichEditFile 文件选择编辑框
	XC_RICHEDIT_FOLDER RichEditFolder 文件夹选择编辑框
	XC_LIST_HEADER 列表头元素
*/
const (
	XC_ELE    XC_OBJECT_TYPE = 21
	XC_BUTTON XC_OBJECT_TYPE = iota + XC_ELE
	XC_RICHEDIT
	XC_COMBOBOX
	XC_SCROLLBAR
	XC_SCROLLVIEW
	XC_LIST
	XC_LISTBOX
	XC_LISTVIEW
	XC_TREE
	XC_MENUBAR
	XC_SLIDERBAR
	XC_PROGRESSBAR
	XC_TOOLBAR
	XC_MONTHCAL
	XC_DATETIME
	XC_PROPERTYGRID
	XC_RICHEDIT_COLOR
	XC_TABBAR
	XC_TEXTLINK
	XC_PANE
	XC_PANE_SPLIT
	XC_MENUBAR_BUTTON
	XC_TOOLBAR_BUTTON
	XC_PROPERTYPAGE_LABEL
	XC_PIER
	XC_BUTTON_MENU
	XC_VIRTUAL_ELE
	XC_RICHEDIT_FILE
	XC_RICHEDIT_FOLDER
	XC_LIST_HEADER
)

/*
接口句柄类型
	XC_SHAPE 形状对象
	XC_SHAPE_TEXT 形状对象-文本
	XC_SHAPE_PICTURE 形状对象-图片
	XC_SHAPE_RECT 形状对象-矩形
	XC_SHAPE_ELLIPSE 形状对象-圆
	XC_SHAPE_LINE 形状对象-直线
	XC_SHAPE_GROUPBOX 形状对象-组框
	XC_SHAPE_GIF 形状对象-GIF
*/
const (
	XC_SHAPE      XC_OBJECT_TYPE = 61
	XC_SHAPE_TEXT XC_OBJECT_TYPE = iota + XC_SHAPE
	XC_SHAPE_PICTURE
	XC_SHAPE_RECT
	XC_SHAPE_ELLIPSE
	XC_SHAPE_LINE
	XC_SHAPE_GROUPBOX
	XC_SHAPE_GIF
)

/*
接口句柄类型
	XC_MENU 弹出菜单
	XC_IMAGE 图片操作
	XC_HDRAW 绘图操作
	XC_FONT 炫彩字体
	XC_FLASH flash
*/
const (
	XC_MENU  XC_OBJECT_TYPE = 81
	XC_IMAGE XC_OBJECT_TYPE = iota + XC_MENU
	XC_HDRAW
	XC_FONT
	XC_FLASH
)

/*
接口句柄类型
	XC_LAYOUT_OBJECT 布局对象LayoutObject
	XC_ADAPTER_TABLE 数据适配器AdapterTable
	XC_ADAPTER_TREE 数据适配器AdapterTree
	XC_ADAPTER_LISTVIEW 数据适配器AdapterListView
	XC_ADAPTER_MAP 数据适配器AdapterMap
*/
const (
	XC_LAYOUT_OBJECT XC_OBJECT_TYPE = 101
	XC_ADAPTER_TABLE XC_OBJECT_TYPE = iota + XC_LAYOUT_OBJECT
	XC_ADAPTER_TREE
	XC_ADAPTER_LISTVIEW
	XC_ADAPTER_MAP
)

// HXCGUI: 资源句柄
type HXCGUI uintptr

// 句柄列表
type (
	// 窗口资源句柄
	HWINDOW HXCGUI

	// 系统窗口句柄
	HWND HXCGUI

	// 元素资源句柄
	HELE HXCGUI

	// 菜单资源句柄
	HMENUX HXCGUI

	// 字符串资源句柄
	HSTRING HXCGUI

	// 图形绘制资源句柄
	HDRAW HXCGUI

	// 数组资源句柄
	HARRAY HXCGUI

	// 图片资源句柄
	HIMAGE HXCGUI

	// ICO资源句柄
	HICON HXCGUI

	// XML资源句柄
	HXMLRES HXCGUI

	// 数组迭代器
	HARRAY_ITERATOR HXCGUI

	// 炫彩字体句柄
	HFONTX HXCGUI

	// 炫彩HFONT句柄
	HFONT HXCGUI

	// 窗格组句柄
	HPANE_GROUP HXCGUI

	// 背景内容管理器句柄
	HBKINFOM HXCGUI

	// 鼠标光标句柄.
	HCURSOR HXCGUI

	// 设备上下文句柄HDC
	HDC HXCGUI

	// 纯色逻辑刷句柄
	HBRUSH HXCGUI

	// 逻辑笔句柄
	HPEN HXCGUI

	// 区域句柄
	HRGN HXCGUI

	// 位图句柄
	HBITMAP HXCGUI
)

// 窗口默认大小
const CW_USEDEFAULT = ^0x7fffffff

type Xc_window_style_ uint32

/*
炫彩窗口样式
	xc_window_style_nothing 什么也没有
	xc_window_style_caption top布局,如果指定该属性,默认为绑定标题栏元素
	xc_window_style_border 边框,指定默认上下左右布局大小,如果没有指定,那么边框布局大小为0
	xc_window_style_center 窗口居中
	xc_window_style_drag_border 拖动窗口边框
	xc_window_style_drag_window 拖动窗口
	xc_window_style_default 窗口默认样式
	xc_window_style_modal 模态窗口样式
*/
const (
	XC_WINDOW_STYLE_NOTHING     Xc_window_style_ = 0x00000000
	XC_WINDOW_STYLE_CAPTION     Xc_window_style_ = 0x00000001
	XC_WINDOW_STYLE_BORDER      Xc_window_style_ = 0x00000002
	XC_WINDOW_STYLE_CENTER      Xc_window_style_ = 0x00000004
	XC_WINDOW_STYLE_DRAG_BORDER Xc_window_style_ = 0x00000008
	XC_WINDOW_STYLE_DRAG_WINDOW Xc_window_style_ = 0x00000010
	XC_WINDOW_STYLE_DEFAULT     Xc_window_style_ = 0x00000001 | 0x00000002 | 0x00000004 | 0x00000008
	XC_WINDOW_STYLE_MODAL       Xc_window_style_ = 0x00000001 | 0x00000004 | 0x00000008
)

// Window style constants
const (
	WS_OVERLAPPED       = 0X00000000
	WS_POPUP            = 0X80000000
	WS_CHILD            = 0X40000000
	WS_MINIMIZE         = 0X20000000
	WS_VISIBLE          = 0X10000000
	WS_DISABLED         = 0X08000000
	WS_CLIPSIBLINGS     = 0X04000000
	WS_CLIPCHILDREN     = 0X02000000
	WS_MAXIMIZE         = 0X01000000
	WS_CAPTION          = 0X00C00000
	WS_BORDER           = 0X00800000
	WS_DLGFRAME         = 0X00400000
	WS_VSCROLL          = 0X00200000
	WS_HSCROLL          = 0X00100000
	WS_SYSMENU          = 0X00080000
	WS_THICKFRAME       = 0X00040000
	WS_GROUP            = 0X00020000
	WS_TABSTOP          = 0X00010000
	WS_MINIMIZEBOX      = 0X00020000
	WS_MAXIMIZEBOX      = 0X00010000
	WS_TILED            = 0X00000000
	WS_ICONIC           = 0X20000000
	WS_SIZEBOX          = 0X00040000
	WS_OVERLAPPEDWINDOW = 0X00000000 | 0X00C00000 | 0X00080000 | 0X00040000 | 0X00020000 | 0X00010000
	WS_POPUPWINDOW      = 0X80000000 | 0X00800000 | 0X00080000
	WS_CHILDWINDOW      = 0X40000000
)

// Extended window style constants
const (
	WS_EX_DLGMODALFRAME    = 0X00000001
	WS_EX_NOPARENTNOTIFY   = 0X00000004
	WS_EX_TOPMOST          = 0X00000008
	WS_EX_ACCEPTFILES      = 0X00000010
	WS_EX_TRANSPARENT      = 0X00000020
	WS_EX_MDICHILD         = 0X00000040
	WS_EX_TOOLWINDOW       = 0X00000080
	WS_EX_WINDOWEDGE       = 0X00000100
	WS_EX_CLIENTEDGE       = 0X00000200
	WS_EX_CONTEXTHELP      = 0X00000400
	WS_EX_RIGHT            = 0X00001000
	WS_EX_LEFT             = 0X00000000
	WS_EX_RTLREADING       = 0X00002000
	WS_EX_LTRREADING       = 0X00000000
	WS_EX_LEFTSCROLLBAR    = 0X00004000
	WS_EX_RIGHTSCROLLBAR   = 0X00000000
	WS_EX_CONTROLPARENT    = 0X00010000
	WS_EX_STATICEDGE       = 0X00020000
	WS_EX_APPWINDOW        = 0X00040000
	WS_EX_OVERLAPPEDWINDOW = 0X00000100 | 0X00000200
	WS_EX_PALETTEWINDOW    = 0X00000100 | 0X00000080 | 0X00000008
	WS_EX_LAYERED          = 0X00080000
	WS_EX_NOINHERITLAYOUT  = 0X00100000
	WS_EX_LAYOUTRTL        = 0X00400000
	WS_EX_NOACTIVATE       = 0X08000000
)

/*
XWndShowWindow constants
	SW_HIDE 隐藏窗口并激活其他窗口
	SW_NORMAL 激活并显示一个窗口。如果窗口被最小化或最大化，系统将其恢复到原来的尺寸和大小。应用程序在第一次显示窗口的时候应该指定此标志。
	SW_SHOWNORMAL 激活并显示一个窗口。如果窗口被最小化或最大化，系统将其恢复到原来的尺寸和大小。应用程序在第一次显示窗口的时候应该指定此标志。
	SW_SHOWMINIMIZED 激活窗口并将其最小化
	SW_MAXIMIZE 最大化指定的窗口
	SW_SHOWMAXIMIZED 激活窗口并将其最大化
	SW_SHOWNOACTIVATE 以窗口最近一次的大小和状态显示窗口。激活窗口仍然维持激活状态。
	SW_SHOW 在窗口原来的位置以原来的尺寸激活和显示窗口。
	SW_MINIMIZE 最小化指定的窗口并且激活在Z序中的下一个顶层窗口。
	SW_SHOWMINNOACTIVE 窗口最小化，激活窗口仍然维持激活状态。
	SW_SHOWNA 以窗口原来的状态显示窗口。激活窗口仍然维持激活状态。
	SW_RESTORE 活并显示窗口。如果窗口最小化或最大化，则系统将窗口恢复到原来的尺寸和位置。在恢复最小化窗口时，应用程序应该指定这个标志。
	SW_SHOWDEFAULT 依据在STARTUPINFO结构中指定的SW_FLAG标志设定显示状态，STARTUPINFO 结构是由启动应用程序的程序传递给CreateProcess函数的。
	SW_FORCEMINIMIZE 在WindowNT5.0中最小化窗口，即使拥有窗口的线程被挂起也会最小化。在从其他线程最小化窗口时才使用这个参数。
*/
const (
	SW_HIDE            = 0
	SW_NORMAL          = 1
	SW_SHOWNORMAL      = 1
	SW_SHOWMINIMIZED   = 2
	SW_MAXIMIZE        = 3
	SW_SHOWMAXIMIZED   = 3
	SW_SHOWNOACTIVATE  = 4
	SW_SHOW            = 5
	SW_MINIMIZE        = 6
	SW_SHOWMINNOACTIVE = 7
	SW_SHOWNA          = 8
	SW_RESTORE         = 9
	SW_SHOWDEFAULT     = 10
	SW_FORCEMINIMIZE   = 11
)

type Window_position_ int32

/*
窗口位置
	window_position_error 错误
	window_position_top top
	window_position_bottom bottom
	window_position_left left
	window_position_right right
	window_position_body body
	window_position_window window 整个窗口
*/
const (
	WINDOW_POSITION_ERROR Window_position_ = -1
	WINDOW_POSITION_TOP   Window_position_ = iota + WINDOW_POSITION_ERROR
	WINDOW_POSITION_BOTTOM
	WINDOW_POSITION_LEFT
	WINDOW_POSITION_RIGHT
	WINDOW_POSITION_BODY
	WINDOW_POSITION_WINDOW
)

type Position struct {
	iRow    int
	iColumn int
}

type Window_transparent_ int32

/*
炫彩窗口透明标识
	window_transparent_false 默认窗口,不透明
	window_transparent_shaped 透明窗口,带透明通道,异型
	window_transparent_shadow 阴影窗口,带透明通道,边框阴影,窗口透明或半透明,当前未启用.
	window_transparent_simple 透明窗口,不带透明通道,指定半透明度,指定透明色.
	window_transparent_win7 WIN7玻璃窗口,需要WIN7开启特效,当前未启用.
*/
const (
	WINDOW_TRANSPARENT_FALSE Window_transparent_ = iota
	WINDOW_TRANSPARENT_SHAPED
	WINDOW_TRANSPARENT_SHADOW
	WINDOW_TRANSPARENT_SIMPLE
	WINDOW_TRANSPARENT_WIN7
)

/*
特殊ID
	XC_ID_ROOT 根节点
	XC_ID_ERROR ID错误
	XC_ID_FIRST 插入开始位置
	XC_ID_LAST 插入末尾位置
*/
const (
	XC_ID_ROOT  = 0
	XC_ID_ERROR = -1
	XC_ID_FIRST = -2
	XC_ID_LAST  = -3
)

type Menu_state_flags_ uint32

/*
弹出菜单项标识
	menu_state_flags_normal 正常
	menu_state_flags_select 选择
	menu_state_flags_check 勾选
	menu_state_flags_popup 弹出
	menu_state_flags_separator 分隔栏 ID号任意,ID号被忽略
	menu_state_flags_disable 禁用
*/
const (
	MENU_STATE_FLAGS_NORMAL    Menu_state_flags_ = 0x00
	MENU_STATE_FLAGS_SELECT    Menu_state_flags_ = 0x01
	MENU_STATE_FLAGS_CHECK     Menu_state_flags_ = 0x02
	MENU_STATE_FLAGS_POPUP     Menu_state_flags_ = 0x04
	MENU_STATE_FLAGS_SEPARATOR Menu_state_flags_ = 0x08
	MENU_STATE_FLAGS_DISABLE   Menu_state_flags_ = 0x10
)

type Menu_popup_position_ uint32

/*
弹出菜单方向
	menu_popup_position_left_top 左上角
	menu_popup_position_left_bottom 左下角
	menu_popup_position_right_top 右上角
	menu_popup_position_right_bottom 右下角
	menu_popup_position_center_left 左居中
	menu_popup_position_center_top 上居中
	menu_popup_position_center_right 右居中
	menu_popup_position_center_bottom 下居中
*/
const (
	MENU_POPUP_POSITION_LEFT_TOP Menu_popup_position_ = iota
	MENU_POPUP_POSITION_LEFT_BOTTOM
	MENU_POPUP_POSITION_RIGHT_TOP
	MENU_POPUP_POSITION_RIGHT_BOTTOM
	MENU_POPUP_POSITION_CENTER_LEFT
	MENU_POPUP_POSITION_CENTER_TOP
	MENU_POPUP_POSITION_CENTER_RIGHT
	MENU_POPUP_POSITION_CENTER_BOTTOM
)

/*
菜单ID , 当前未使用
	IDM_CLIP 剪切
	IDM_COPY 复制
	IDM_PASTE 粘贴
	IDM_DELETE 删除
	IDM_SELECTALL 全选
	IDM_DELETEALL 清空
*/
const (
	IDM_CLIP      = 1000000000
	IDM_COPY      = 1000000001
	IDM_PASTE     = 1000000002
	IDM_DELETE    = 1000000003
	IDM_SELECTALL = 1000000004
	IDM_DELETEALL = 1000000005
)

type Image_draw_type_ uint32

/*
图片绘制类型
	image_draw_type_default 默认
	image_draw_type_stretch 拉伸
	image_draw_type_adaptive 自适应,九宫格
	image_draw_type_tile 平铺
	image_draw_type_fixed_ratio 当图片超出显示范围时,按照原始比例压缩显示图片
	image_draw_type_adaptive_border 九宫格不绘制中间区域
*/
const (
	IMAGE_DRAW_TYPE_DEFAULT Image_draw_type_ = iota
	IMAGE_DRAW_TYPE_STRETCH
	IMAGE_DRAW_TYPE_ADAPTIVE
	IMAGE_DRAW_TYPE_TILE
	IMAGE_DRAW_TYPE_FIXED_RATIO
	IMAGE_DRAW_TYPE_ADAPTIVE_BORDER
)

type Common_state3_ uint32

/*
common_state3_ 普通三种状态
	common_state3_leave 离开
	common_state3_stay 停留
	common_state3_down 按下
*/

const (
	COMMON_STATE3_LEAVE Common_state3_ = iota
	COMMON_STATE3_STAY
	COMMON_STATE3_DOWN
)

type Button_state_ uint32

/*
button_state_ 按钮状态
	button_state_leave 离开状态
	button_state_stay 停留状态
	button_state_down 按下状态
	button_state_check 选中状态
	button_state_disable 禁用状态
*/
const (
	BUTTON_STATE_LEAVE Button_state_ = iota
	BUTTON_STATE_STAY
	BUTTON_STATE_DOWN
	BUTTON_STATE_CHECK
	BUTTON_STATE_DISABLE
)

type Button_type_ uint32

/*
button_type_ 按钮类型(用于区分功能)
	button_type_default 默认类型
	button_type_check 复选按钮
	button_type_radio 单选按钮
	button_type_close 窗口关闭按钮
	button_type_min 窗口最小化按钮
	button_type_max 窗口最大化还原按钮
*/
const (
	BUTTON_TYPE_DEFAULT Button_type_ = iota
	BUTTON_TYPE_CHECK
	BUTTON_TYPE_RADIO
	BUTTON_TYPE_CLOSE
	BUTTON_TYPE_MIN
	BUTTON_TYPE_MAX
)

type Button_style_ uint32

/*
button_style_ 按钮样式(用于区分外观)
	button_style_default 默认风格
	button_style_check 复选按钮
	button_style_radio 单选按钮
	button_style_scrollbar_up 水平滚动条,上按钮
	button_style_scrollbar_down 水平滚动条,下按钮
	button_style_scrollbar_left 水平滚动条,左按钮
	button_style_scrollbar_right 水平滚动条,右按钮
	button_style_scrollbar_slider 滚动条,滑块
	button_style_tabBar_button TabBar上的按钮
	button_style_toolBar_left ToolBar左滚动按钮
	button_style_toolBar_right ToolBar右滚动按钮
	button_style_pane_close 窗格关闭按钮
	button_style_pane_lock 窗格锁定按钮
	button_style_pane_menu 窗格下拉菜单按钮
	button_style_pane_dockH 框架窗口左边或右边码头上按钮
	button_style_pane_dockV 框架窗口上边或下边码头上按钮
*/
const (
	BUTTON_STYLE_DEFAULT Button_style_ = iota
	BUTTON_STYLE_CHECK
	BUTTON_STYLE_RADIO
	BUTTON_STYLE_SCROLLBAR_UP
	BUTTON_STYLE_SCROLLBAR_DOWN
	BUTTON_STYLE_SCROLLBAR_LEFT
	BUTTON_STYLE_SCROLLBAR_RIGHT
	BUTTON_STYLE_SCROLLBAR_SLIDER
	BUTTON_STYLE_TABBAR_BUTTON
	BUTTON_STYLE_TOOLBAR_LEFT
	BUTTON_STYLE_TOOLBAR_RIGHT
	BUTTON_STYLE_PANE_CLOSE
	BUTTON_STYLE_PANE_LOCK
	BUTTON_STYLE_PANE_MENU
	BUTTON_STYLE_PANE_DOCKH
	BUTTON_STYLE_PANE_DOCKV
)

type ComboBox_state_ uint32

/*
ComboBox状态
	comboBox_state_leave 鼠标离开状态
	comboBox_state_stay 鼠标停留状态
	comboBox_state_down 按下状态
*/
const (
	COMBOBOX_STATE_LEAVE ComboBox_state_ = iota
	COMBOBOX_STATE_STAY
	COMBOBOX_STATE_DOWN
)

type List_item_state_ uint32

/*
List项状态
	list_item_state_leave 项鼠标离开状态
	list_item_state_stay 项鼠标停留状态
	list_item_state_select 项选择状态
*/
const (
	LIST_ITEM_STATE_LEAVE List_item_state_ = iota
	LIST_ITEM_STATE_STAY
	LIST_ITEM_STATE_SELECT
)

type Tree_item_state_ uint32

/*
Tree项状态
	tree_item_state_leave 项鼠标离开状态
	tree_item_state_select 项选择状态
*/
const (
	TREE_ITEM_STATE_LEAVE Tree_item_state_ = iota
	TREE_ITEM_STATE_SELECT
)

type Button_icon_align_ uint32

/*
按钮图标对齐方式
	button_icon_align_left 图标在左边
	button_icon_align_top 图标在顶部
	button_icon_align_right 图标在右边
	button_icon_align_bottom 图标在底部
*/
const (
	BUTTON_ICON_ALIGN_LEFT Button_icon_align_ = iota
	BUTTON_ICON_ALIGN_TOP
	BUTTON_ICON_ALIGN_RIGHT
	BUTTON_ICON_ALIGN_BOTTOM
)

type Layout_size_type_ uint32

/*
布局大小类型
	layout_size_type_fixed 指定大小
	layout_size_type_fill fill 填充父
	layout_size_type_auto auto 自动大小,根据内容计算大小
	layout_size_type_weight weight 比例,按照比例分配剩余空间
	layout_size_type_disable disable 不使用
*/
const (
	LAYOUT_SIZE_TYPE_FIXED Layout_size_type_ = iota
	LAYOUT_SIZE_TYPE_FILL
	LAYOUT_SIZE_TYPE_AUTO
	LAYOUT_SIZE_TYPE_WEIGHT
	LAYOUT_SIZE_TYPE_DISABLE
)

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

type MessageBox_flags_ uint32

/*
弹出消息框
	messageBox_flags_other 其他
	messageBox_flags_ok 确定按钮
	messageBox_flags_cancel 取消按钮
*/
const (
	MESSAGEBOX_FLAGS_OTHER  MessageBox_flags_ = 0x00
	MESSAGEBOX_FLAGS_OK     MessageBox_flags_ = 0x01
	MESSAGEBOX_FLAGS_CANCEL MessageBox_flags_ = 0x02
)

type propertyGrid_item_type_ uint32

/*
属性网格项类型
	propertyGrid_item_type_text 默认,字符串类型
	propertyGrid_item_type_edit 编辑框
	propertyGrid_item_type_edit_color 颜色选择元素
	propertyGrid_item_type_edit_file 文件选择编辑框
	propertyGrid_item_type_edit_set 设置
	propertyGrid_item_type_comboBox 组合框
	propertyGrid_item_type_group 组
*/
const (
	PROPERTYGRID_ITEM_TYPE_TEXT propertyGrid_item_type_ = iota
	PROPERTYGRID_ITEM_TYPE_EDIT
	PROPERTYGRID_ITEM_TYPE_EDIT_COLOR
	PROPERTYGRID_ITEM_TYPE_EDIT_FILE
	PROPERTYGRID_ITEM_TYPE_EDIT_SET
	PROPERTYGRID_ITEM_TYPE_COMBOBOX
	PROPERTYGRID_ITEM_TYPE_GROUP
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

type Zorder_ uint32

/*
Z序位置
	zorder_top 最上面
	zorder_bottom 最下面
	zorder_before 指定目标下面
	zorder_after 指定目标上面
*/
const (
	ZORDER_TOP Zorder_ = iota
	ZORDER_BOTTOM
	ZORDER_BEFORE
	ZORDER_AFTER
)

type Window_state_flag_ uint32

/*
组合状态 - 7
	window_state_flag_nothing  无
	window_state_whole_leave  整个窗口
	window_state_body_leave  窗口-body
	window_state_top_leave  窗口-top
	window_state_bottom_leave  窗口-bottom
	window_state_left_leave  窗口-left
	window_state_right_leave  窗口-right
*/
const (
	WINDOW_STATE_FLAG_NOTHING Window_state_flag_ = 0x0000
	WINDOW_STATE_WHOLE_LEAVE  Window_state_flag_ = 0x0001
	WINDOW_STATE_BODY_LEAVE   Window_state_flag_ = 0x0002
	WINDOW_STATE_TOP_LEAVE    Window_state_flag_ = 0x0004
	WINDOW_STATE_BOTTOM_LEAVE Window_state_flag_ = 0x0008
	WINDOW_STATE_LEFT_LEAVE   Window_state_flag_ = 0x0010
	WINDOW_STATE_RIGHT_LEAVE  Window_state_flag_ = 0x0020
)

type Element_state_flag_ uint32

/*
组合状态 - 8
	element_state_flag_nothing  无
	element_state_flag_enable  启用
	element_state_flag_disable  禁用
	element_state_flag_focus  焦点
	element_state_flag_focus_no  无焦点
	element_state_flag_leave  鼠标离开
	element_state_flag_stay  为扩展模块保留
	element_state_flag_down  为扩展模块保留
*/
const (
	ELEMENT_STATE_FLAG_NOTHING  Element_state_flag_ = Element_state_flag_(WINDOW_STATE_FLAG_NOTHING)
	ELEMENT_STATE_FLAG_ENABLE   Element_state_flag_ = 0x0001
	ELEMENT_STATE_FLAG_DISABLE  Element_state_flag_ = 0x0002
	ELEMENT_STATE_FLAG_FOCUS    Element_state_flag_ = 0x0004
	ELEMENT_STATE_FLAG_FOCUS_NO Element_state_flag_ = 0x0008
	ELEMENT_STATE_FLAG_LEAVE    Element_state_flag_ = 0x0010
	ELEMENT_STATE_FLAG_STAY     Element_state_flag_ = 0x0020
	ELEMENT_STATE_FLAG_DOWN     Element_state_flag_ = 0x0040
)

type Button_state_flag_ uint32

/*
组合状态 - 5
	button_state_flag_leave 鼠标离开
	button_state_flag_stay 鼠标停留
	button_state_flag_down 鼠标按下
	button_state_flag_check 选中
	button_state_flag_check_no 未选中
*/
const (
	BUTTON_STATE_FLAG_LEAVE    Button_state_flag_ = Button_state_flag_(ELEMENT_STATE_FLAG_LEAVE)
	BUTTON_STATE_FLAG_STAY     Button_state_flag_ = Button_state_flag_(ELEMENT_STATE_FLAG_STAY)
	BUTTON_STATE_FLAG_DOWN     Button_state_flag_ = Button_state_flag_(ELEMENT_STATE_FLAG_DOWN)
	BUTTON_STATE_FLAG_CHECK    Button_state_flag_ = 0x0080
	BUTTON_STATE_FLAG_CHECK_NO Button_state_flag_ = 0x0100
)

type ComboBox_state_flag_ uint32

/*
组合状态 - 3
	comboBox_state_flag_leave  鼠标离开
	comboBox_state_flag_stay  鼠标停留
	comboBox_state_flag_down  鼠标按下
*/
const (
	COMBOBOX_STATE_FLAG_LEAVE ComboBox_state_flag_ = ComboBox_state_flag_(ELEMENT_STATE_FLAG_LEAVE)
	COMBOBOX_STATE_FLAG_STAY  ComboBox_state_flag_ = ComboBox_state_flag_(ELEMENT_STATE_FLAG_STAY)
	COMBOBOX_STATE_FLAG_DOWN  ComboBox_state_flag_ = ComboBox_state_flag_(ELEMENT_STATE_FLAG_DOWN)
)

type ListBox_state_flag_ uint32

/*
组合状态 - 3
	listBox_state_flag_item_leave  项鼠标离开
	listBox_state_flag_item_stay  项鼠标停留
	listBox_state_flag_item_select  项选择
*/
const (
	LISTBOX_STATE_FLAG_ITEM_LEAVE  ListBox_state_flag_ = 0x0080
	LISTBOX_STATE_FLAG_ITEM_STAY   ListBox_state_flag_ = 0x0100
	LISTBOX_STATE_FLAG_ITEM_SELECT ListBox_state_flag_ = 0x0200
)

type List_state_flag_ uint32

/*
组合状态 - 3
	list_state_flag_item_leave  项鼠标离开
	list_state_flag_item_stay  项鼠标停留
	list_state_flag_item_select  项选择
*/
const (
	LIST_STATE_FLAG_ITEM_LEAVE  List_state_flag_ = 0x0080
	LIST_STATE_FLAG_ITEM_STAY   List_state_flag_ = 0x0100
	LIST_STATE_FLAG_ITEM_SELECT List_state_flag_ = 0x0200
)

type ListView_state_flag_ uint32

/*
组合状态 - 3
	listView_state_flag_item_leave  项鼠标离开
	listView_state_flag_item_stay  项鼠标停留
	listView_state_flag_item_select  项选择
*/
const (
	LISTVIEW_STATE_FLAG_ITEM_LEAVE  ListView_state_flag_ = 0x0080
	LISTVIEW_STATE_FLAG_ITEM_STAY   ListView_state_flag_ = 0x0100
	LISTVIEW_STATE_FLAG_ITEM_SELECT ListView_state_flag_ = 0x0200
)

type Tree_state_flag_ uint32

/*
组合状态 - 2
	tree_state_flag_item_leave 项鼠标离开
	tree_state_flag_item_select 项选择
*/
const (
	TREE_STATE_FLAG_ITEM_LEAVE  Tree_state_flag_ = 0x0080
	TREE_STATE_FLAG_ITEM_SELECT Tree_state_flag_ = 0x0100
)

type MonthCal_state_flag_ uint32

/*
组合状态 - 10
	monthCal_state_flag_leave  离开状态
	monthCal_state_flag_item_leave  项-离开
	monthCal_state_flag_item_stay  项-停留
	monthCal_state_flag_item_down  项-按下
	monthCal_state_flag_item_select  项-选择
	monthCal_state_flag_item_today  项-今天
	monthCal_state_flag_item_other  项-上月及下月
	monthCal_state_flag_item_last_month  项-上月
	monthCal_state_flag_item_cur_month  项-当月
	monthCal_state_flag_item_next_month  项-下月
*/
const (
	MONTHCAL_STATE_FLAG_LEAVE           MonthCal_state_flag_ = MonthCal_state_flag_(ELEMENT_STATE_FLAG_LEAVE)
	MONTHCAL_STATE_FLAG_ITEM_LEAVE      MonthCal_state_flag_ = 0x0080
	MONTHCAL_STATE_FLAG_ITEM_STAY       MonthCal_state_flag_ = 0x0100
	MONTHCAL_STATE_FLAG_ITEM_DOWN       MonthCal_state_flag_ = 0x0200
	MONTHCAL_STATE_FLAG_ITEM_SELECT     MonthCal_state_flag_ = 0x0400
	MONTHCAL_STATE_FLAG_ITEM_TODAY      MonthCal_state_flag_ = 0x0800
	MONTHCAL_STATE_FLAG_ITEM_OTHER      MonthCal_state_flag_ = 0x1000
	MONTHCAL_STATE_FLAG_ITEM_LAST_MONTH MonthCal_state_flag_ = 0x2000
	MONTHCAL_STATE_FLAG_ITEM_CUR_MONTH  MonthCal_state_flag_ = 0x4000
	MONTHCAL_STATE_FLAG_ITEM_NEXT_MONTH MonthCal_state_flag_ = 0x8000
)

type MonthCal_button_type_ uint32

/*
月历元素上的按钮类型 - 5
	monthCal_button_type_today 今天按钮
	monthCal_button_type_last_year 上一年
	monthCal_button_type_next_year 下一年
	monthCal_button_type_last_month 上一月
	monthCal_button_type_next_month 下一月
*/
const (
	MONTHCAL_BUTTON_TYPE_TODAY MonthCal_button_type_ = iota
	MONTHCAL_BUTTON_TYPE_LAST_YEAR
	MONTHCAL_BUTTON_TYPE_NEXT_YEAR
	MONTHCAL_BUTTON_TYPE_LAST_MONTH
	MONTHCAL_BUTTON_TYPE_NEXT_MONTH
)

/*
月历元素项数据
数据成员
 int  nDay
  日期
 int  nType
  1上月,2当月,3下月
 int  nState
  组合状态 monthCal_state_flag_
 RECT  rcItem
  项坐标
*/
type MonthCal_item_i struct {
	nDay   int
	nType  int
	nState int
	rcItem RECT
}
