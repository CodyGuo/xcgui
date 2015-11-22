/**************************************************************************\
*
* Copyright (c)  炫彩界面库. All Rights Reserved.
*
* 炫彩界面库 - 官方网站 : http://www.xcgui.com
*
\**************************************************************************/
#pragma once

//v1.8.9.5
////////////输入输出///////////////////
#define in_
#define out_
#define in_out_
#define XC_API extern "C"

#define  XC_NAME    _T("name")
#define  XC_NAME2   _T("name2")
#define  XC_NAME3   _T("name3")
#define  XC_NAME4   _T("name4")
#define  XC_NAME5   _T("name5")
#define  XC_NAME6   _T("name6")
#define  XC_NAME7   _T("name7")
#define  XC_NAME8   _T("name8")
#define  XC_NAME9   _T("name9")
#define  XC_NAME10   _T("name10")
/// @defgroup groupObjectType 接口句柄类型
/// @ingroup macroDefine
/// @{
/// 炫彩对象类型
enum XC_OBJECT_TYPE  ///<炫彩对象类型
{
	XC_ERROR =-1,      ///<错误类型
	XC_WINDOW = 1,     ///<窗口
	XC_MODALWINDOW ,   ///<模态窗口
	XC_FRAMEWND    ,   ///<框架窗口
	XC_FLOATWND     ,  ///<浮动窗口
	XC_COMBOBOXWINDOW   =   11,   //comboBoxWindow_        组合框弹出下拉列表窗口
	XC_POPUPMENUWINDOW  ,         //popupMenuWindow_       弹出菜单主窗口
	XC_POPUPMENUCHILDWINDOW ,     //popupMenuChildWindow_  弹出菜单子窗口

	//元素
	XC_ELE      = 21, ///<基础元素
	XC_BUTTON   ,     ///<按钮
	XC_RICHEDIT   ,   ///<富文本编辑框
	XC_COMBOBOX    ,  ///<下拉组合框元素
	XC_SCROLLBAR   ,  ///<滚动条元素
	XC_SCROLLVIEW  ,  ///<滚动视图元素
	XC_LIST        ,  ///<列表元素
	XC_LISTBOX    ,   ///<列表框元素
	XC_LISTVIEW   ,   ///<列表视图,大图标
	XC_TREE       ,   ///<树元素
	XC_MENUBAR    ,   ///<菜单条
	XC_SLIDERBAR     ,  ///<滑动条
	XC_PROGRESSBAR   ,  ///<进度条
	XC_TOOLBAR       ,  ///<工具条
	XC_MONTHCAL      ,  ///<月历元素
	XC_DATETIME      ,  ///<时间元素
	XC_PROPERTYGRID  ,  ///<属性网格
	XC_RICHEDIT_COLOR   ,  ///<颜色选择元素
	XC_RICHEDIT_SET,
	XC_TABBAR,          ///<tab条
	XC_TEXTLINK,        ///<文本链接按钮

	XC_PANE          ,  ///<布局窗格
	XC_PANE_SPLIT    ,  ///<布局窗格拖动分割条
	XC_MENUBAR_BUTTON       ,   ///<菜单条上的按钮
	XC_TOOLBAR_BUTTON       ,   ///<工具条上按钮
	XC_PROPERTYPAGE_LABEL   ,   ///<属性页标签按钮
	XC_PIER                 ,   ///<窗格停靠码头
	XC_BUTTON_MENU          ,   ///<弹出菜单按钮
	XC_VIRTUAL_ELE          ,   ///<虚拟元素
	XC_RICHEDIT_FILE        ,   ///<RichEditFile 文件选择编辑框
	XC_RICHEDIT_FOLDER      ,   ///<RichEditFolder  文件夹选择编辑框
	XC_LIST_HEADER          ,   ///<列表头元素

	XC_SHAPE       =61,  ///<形状对象
	XC_SHAPE_TEXT,       ///<形状对象-文本
	XC_SHAPE_PICTURE,    ///<形状对象-图片
	XC_SHAPE_RECT,       ///<形状对象-矩形
	XC_SHAPE_ELLIPSE,    ///<形状对象-圆
	XC_SHAPE_LINE,       ///<形状对象-直线
	XC_SHAPE_GROUPBOX,   ///<形状对象-组框
	XC_SHAPE_GIF,        ///<形状对象-GIF

	//其他类型
	XC_MENU       = 81,  ///<弹出菜单
	XC_IMAGE         ,   ///<图片操作
	XC_HDRAW         ,   ///<绘图操作
	XC_FONT          ,   ///<炫彩字体
	XC_FLASH         ,   ///<flash   
	XC_PANE_CELL     ,  
	XC_WEB           ,

	XC_LAYOUT_OBJECT = 101,  ///<布局对象LayoutObject
	XC_ADAPTER          ,
	XC_ADAPTER_TABLE    ,    ///<数据适配器AdapterTable
	XC_ADAPTER_TREE     ,    ///<数据适配器AdapterTree
	XC_ADAPTER_LISTVIEW ,    ///<数据适配器AdapterListView
	XC_ADAPTER_MAP,          ///<数据适配器AdapterMap

	//无实体对象,只是用来判断布局
	XC_LAYOUT_LISTVIEW     = 111,
	XC_LAYOUT_LIST     ,
	XC_LAYOUT_OBJECT_GROUP ,
	XC_LAYOUT_OBJECT_ITEM  ,
	XC_LAYOUT_PANEL,

	XC_BKINFOM,
	//XC_BUTTON_EX,
};

///@}

//#ifdef _DEBUG
#define DECLARE_HANDLEX(name) struct name##__ { int unused; }; typedef struct name##__ *name
//#else
//#define	DECLARE_HANDLEX(name) typedef void* name
//#endif

/// @defgroup groupHandle 句柄列表
/// @ingroup macroDefine
/// @{
typedef void*   HXCGUI;   ///<资源句柄
DECLARE_HANDLEX(HWINDOW); ///<窗口资源句柄
DECLARE_HANDLEX(HELE);    ///<元素资源句柄
DECLARE_HANDLEX(HMENUX);  ///<菜单资源句柄
DECLARE_HANDLEX(HDRAW);   ///<图形绘制资源句柄
DECLARE_HANDLEX(HIMAGE);  ///<图片资源句柄
//DECLARE_HANDLEX(HARRAY_ITERATOR); ///<数组迭代器
DECLARE_HANDLEX(HFONTX);    ///<炫彩字体句柄
DECLARE_HANDLEX(HBKINFOM);    ///<背景内容管理器句柄

///@}

#ifndef _INC_SHELLAPI
DECLARE_HANDLE(HDROP);
#endif

typedef   void  bkInfo_i;  //背景信息

//窗口样式----------------------------
/// @defgroup groupWindowStyle 炫彩窗口样式
/// @ingroup macroDefine
/// @{
enum  xc_window_style_   
{
	xc_window_style_nothing =  0x00000000,      ///<什么也没有
	xc_window_style_caption=   0x00000001,      ///<top布局,如果指定该属性,默认为绑定标题栏元素
	xc_window_style_border=    0x00000002,      ///<边框,指定默认上下左右布局大小,如果没有指定,那么边框布局大小为0
	xc_window_style_center=    0x00000004,      ///<窗口居中
	xc_window_style_drag_border= 0x00000008,    ///<拖动窗口边框
	xc_window_style_drag_window =0x00000010,    ///<拖动窗口
	xc_window_style_allow_maxWindow=   0x00000020,   ///允许窗口最大化

	///窗口默认样式
	xc_window_style_default=(xc_window_style_caption | xc_window_style_border | xc_window_style_center | xc_window_style_drag_border | xc_window_style_allow_maxWindow),
	xc_window_style_default2=(xc_window_style_caption | xc_window_style_border | xc_window_style_drag_border | xc_window_style_allow_maxWindow),

	///模态窗口样式
	xc_window_style_modal=(xc_window_style_caption | xc_window_style_center |  xc_window_style_border),
};

///@}

#ifdef _DEBUG
#define XTRACE   xtrace
#define XTRACEW  xtracew
#define XERROR   XTRACE 
#define XERRORW  XTRACEW
#else
#define XTRACE
#define XTRACEW
#define XERROR 
#define XERRORW
#endif

#define  XC_WINDOW_TOP            1 //上
#define  XC_WINDOW_BOTTOM         2 //下
#define  XC_WINDOW_LEFT           3 //左
#define  XC_WINDOW_RIGHT          4 //右
#define  XC_WINDOW_TOPLEFT        5 //左上角
#define  XC_WINDOW_TOPRIGHT       6 //右上角
#define  XC_WINDOW_BOTTOMLEFT     7 //左下角
#define  XC_WINDOW_BOTTOMRIGHT    8 //右下角
#define  XC_WINDOW_CAPTION        9 //标题栏移动窗口区域
#define  XC_WINDOW_BODY           10

/// @defgroup groupWindow_position 窗口位置
/// @ingroup macroDefine
/// @{
enum   window_position_
{
	window_position_error=-1,  ///<错误
	window_position_top=0,     ///<top
	window_position_bottom,    ///<bottom
	window_position_left,      ///<left
	window_position_right,     ///<right
	window_position_body,      ///<body
	window_position_window,    ///<window 整个窗口 
};

///@}

//透明窗口
/// @defgroup groupWindowTransparent 炫彩窗口透明标识
/// @ingroup macroDefine
/// @{
enum  window_transparent_
{
	window_transparent_false=0,   ///<默认窗口,不透明
	window_transparent_shaped,    ///<透明窗口,带透明通道,异型
	window_transparent_shadow,    ///<阴影窗口,带透明通道,边框阴影,窗口透明或半透明,当前未启用.
	window_transparent_simple,    ///<透明窗口,不带透明通道,指定半透明度,指定透明色.
	window_transparent_win7,      ///<WIN7玻璃窗口,需要WIN7开启特效,当前未启用.
};

///@}

/// @defgroup groupID_fixed  特殊ID
/// @ingroup macroDefine
/// @{
#define  XC_ID_ROOT            0   ///<根节点
#define  XC_ID_ERROR          -1   ///<ID错误
#define  XC_ID_FIRST          -2   ///<插入开始位置
#define  XC_ID_LAST           -3   ///<插入末尾位置

/// @}


/// @defgroup groupMenu 弹出菜单
/// @ingroup macroDefine
/// @{

//弹出菜单项标识
///@name 弹出菜单项标识
///@{
enum   menu_state_flags_
{
	menu_state_flags_normal=   0x00,   ///<正常
	menu_state_flags_select=   0x01,   ///<选择 
	menu_state_flags_check=    0x02,   ///<勾选
	menu_state_flags_popup=    0x04,   ///<弹出
	menu_state_flags_separator=0x08,   ///<分隔栏 ID号任意,ID号被忽略
	menu_state_flags_disable=  0x10,   ///<禁用
};

///@}

//弹出菜单方向
///@name 弹出菜单方向
///@{
enum  menu_popup_position_
{
	menu_popup_position_left_top=0,      ///<左上角
	menu_popup_position_left_bottom,     ///<左下角
	menu_popup_position_right_top,       ///<右上角
	menu_popup_position_right_bottom,    ///<右下角
	menu_popup_position_center_left,     ///<左居中
	menu_popup_position_center_top,      ///<上居中
	menu_popup_position_center_right,    ///<右居中
	menu_popup_position_center_bottom,   ///<下居中
};

///@}

//菜单ID--------------------------------------
///@name 菜单ID , 当前未使用
///@{
#define  IDM_CLIP          1000000000    ///<剪切
#define  IDM_COPY          1000000001    ///<复制
#define  IDM_PASTE         1000000002    ///<粘贴
#define  IDM_DELETE        1000000003    ///<删除
#define  IDM_SELECTALL     1000000004    ///<全选
#define  IDM_DELETEALL     1000000005    ///<清空

///@}

///@}

/// @defgroup groupImageDrawType 图片绘制类型
/// @ingroup macroDefine
/// @{
enum  image_draw_type_
{
	image_draw_type_default=0,  ///<默认
	image_draw_type_stretch,    ///<拉伸
	image_draw_type_adaptive,   ///<自适应,九宫格
	image_draw_type_tile,       ///<平铺
	image_draw_type_fixed_ratio,  ///<当图片超出显示范围时,按照原始比例压缩显示图片
	image_draw_type_adaptive_border,  ///<九宫格不绘制中间区域
};

///@}


//状态--------------------------------------
/// @defgroup groupCommonState3 普通三种状态
/// @ingroup  macroDefine
/// @{
enum  common_state3_
{
	common_state3_leave=0,  ///<离开
	common_state3_stay,     ///<停留
	common_state3_down,     ///<按下
};

///@}

/// @defgroup groupButtonStateFlag 按钮状态
/// @ingroup  macroDefine
/// @{
enum  button_state_
{
	button_state_leave=0,   ///<离开状态
	button_state_stay,      ///<停留状态
	button_state_down,      ///<按下状态
	button_state_check,     ///<选中状态
	button_state_disable,   ///<禁用状态
};

///@}

/// @defgroup groupButtonType 按钮类型(用于区分功能)
/// @ingroup macroDefine
/// @{
enum button_type_
{
	button_type_default=0, ///<默认类型
	button_type_check,     ///<复选按钮
	button_type_radio,     ///<单选按钮
	button_type_close,     ///<窗口关闭按钮
	button_type_min,       ///<窗口最小化按钮
	button_type_max,       ///<窗口最大化还原按钮
};

///@}

/// @defgroup groupButtonStyle 按钮样式(用于区分外观)
/// @ingroup macroDefine
/// @{
enum  button_style_
{
	button_style_default         = 0,  ///<默认风格
	button_style_check            ,  ///<复选按钮
	button_style_radio            ,  ///<单选按钮
	button_style_scrollbar_up     ,  ///<水平滚动条,上按钮
	button_style_scrollbar_down   ,  ///<水平滚动条,下按钮
	button_style_scrollbar_left   ,  ///<水平滚动条,左按钮
	button_style_scrollbar_right  ,  ///<水平滚动条,右按钮
	button_style_scrollbar_slider ,  ///<滚动条,滑块
	button_style_tabBar_button    ,  ///<TabBar上的按钮
	button_style_toolBar_left     ,  ///<ToolBar左滚动按钮
	button_style_toolBar_right   ,  ///<ToolBar右滚动按钮

	button_style_pane_close     , ///<窗格关闭按钮
	button_style_pane_lock      , ///<窗格锁定按钮
	button_style_pane_menu      , ///<窗格下拉菜单按钮
	button_style_pane_dockH    , ///<框架窗口左边或右边码头上按钮
	button_style_pane_dockV    , ///<框架窗口上边或下边码头上按钮
};

///@}

/// @defgroup   groupComboBoxState ComboBox状态
/// @ingroup macroDefine
/// @{
enum  comboBox_state_
{
	comboBox_state_leave=0,   ///<鼠标离开状态
	comboBox_state_stay=1,    ///<鼠标停留状态
	comboBox_state_down=2,    ///<按下状态
};

///@}

/// @defgroup groupListItemState List项状态
/// @ingroup macroDefine
/// @{
enum  list_item_state_
{
	list_item_state_leave=0,   ///<项鼠标离开状态
	list_item_state_stay=1,    ///<项鼠标停留状态
	list_item_state_select=2,  ///<项选择状态
};

///@}


/// @defgroup groupTreeItemState Tree项状态
/// @ingroup macroDefine
/// @{
enum  tree_item_state_
{
	tree_item_state_leave=0,   ///<项鼠标离开状态
	tree_item_state_select=1,  ///<项选择状态
};

///@}


//按钮图标对齐方式
/// @defgroup groupButtonIconAlign 按钮图标对齐方式
/// @ingroup macroDefine
/// @{
enum  button_icon_align_
{
	button_icon_align_left=0,  ///<图标在左边
	button_icon_align_top,     ///<图标在顶部
	button_icon_align_right,   ///<图标在右边
	button_icon_align_bottom,  ///<图标在底部
};

///@}

/// @defgroup groupLayoutSize 布局大小类型
/// @ingroup macroDefine
/// @{
enum  layout_size_type_
{
	layout_size_type_fixed=0,   ///<指定大小
	layout_size_type_fill,      ///<fill 填充父
	layout_size_type_auto,      ///<auto 自动大小,根据内容计算大小
	layout_size_type_weight,    ///<weight 比例,按照比例分配剩余空间
	layout_size_type_disable,   ///<disable 不使用
};

///@}

/// @defgroup  groupListDrawItemBkFlags  List,ListBox,Tree,项背景绘制标志位
/// @ingroup macroDefine
/// @{
enum  list_drawItemBk_flags_
{
	list_drawItemBk_flags_nothing= 0x000,    ///<不绘制
	list_drawItemBk_flags_leave=   0x001,    ///<绘制鼠标离开状态项背景
	list_drawItemBk_flags_stay=    0x002,    ///<绘制鼠标选择状态项背景
	list_drawItemBk_flags_select=  0x004,    ///<绘制鼠标停留状态项项背景
	list_drawItemBk_flags_group_stay= 0x008  ///<绘制鼠标停留状态项背景,当项为组时
};

/// @}

//弹出消息框类型
/// @defgroup groupMessageBox 弹出消息框
/// @ingroup macroDefine
/// @{
enum  messageBox_flags_
{
	messageBox_flags_other= 0x00,    ///<其他
	messageBox_flags_ok=    0x01,    ///<确定按钮
	messageBox_flags_cancel=0x02,    ///<取消按钮
};

///@}

///@name 项类型
///@{
enum  propertyGrid_item_type_
{
	propertyGrid_item_type_text=0,      ///<默认,字符串类型
	propertyGrid_item_type_edit,        ///<编辑框
	propertyGrid_item_type_edit_color,  ///<颜色选择元素
	propertyGrid_item_type_edit_file,   ///<文件选择编辑框
	propertyGrid_item_type_edit_set,    ///<设置
	propertyGrid_item_type_comboBox,    ///<组合框
	propertyGrid_item_type_group,       ///<组
};

///@}

//////////////框架窗口///////////////////////////////////////////////
//窗格菜单 当前未使用
#define  IDM_LOCK          1000000006    ///<锁定
#define  IDM_DOCK          1000000007    ///<停靠
#define  IDM_FLOAT         1000000008    ///<浮动
#define  IDM_HIDE          1000000009    ///<隐藏

enum  layout_align_
{
	layout_align_left=0,
	layout_align_top,
	layout_align_right,
	layout_align_bottom,
	layout_align_center,
	layout_align_equidistant,
};

enum align_type_
{
	align_error=-1,
	align_left=0,
	align_top,
	align_right,
	align_bottom,
	align_center,
};

enum dock_align_  //停靠码头方向
{
	dock_align_left=0,
	dock_align_top,
	dock_align_right,
	dock_align_bottom,
};

//窗格状态
enum pane_state_
{
	pane_state_any=0,
	pane_state_lock,
	pane_state_dock,
	pane_state_float,
};


typedef void (CALLBACK *pFunDebugError)(const char *pInfo); //错误回调
/////////////////////////////////////////////////////////
//////////////定义结构体/////////////////////////////////
/////////////////////////////////////////////////////////
struct BorderSize_  //4条边的大小
{
	int  leftSize;
	int  topSize;
	int  rightSize;
	int  bottomSize;
};
typedef  BorderSize_  SpaceSize_ , PaddingSize_;

struct  Position
{
	int  iRow;
	int  iColumn;
};

struct  template_info_i  //模板项
{
	int  info;
};

struct listBox_item_i
{
	int     nUserData;  ///<用户绑定数据
	int     nHeight;    ///<项默认高度
	int     nSelHeight; ///<项选中时高度
	//	BOOL    bSelect;    ///<是否选中
	list_item_state_  nState;  ///<状态
	int     index;      ///<项索引
	RECT    rcItem;     ///<项坐标
	HXCGUI  hLayout;     ///<布局对象
	template_info_i  *pInfo;   ///<模板信息
};

struct listBox_item_Info_i
{
	int     nUserData;  //用户绑定数据
	int     nHeight;
	int     nSelHeight;
};

struct  listView_item_id_i
{
	int  iGroup;   ///<组索引
	int  iItem;    ///<项索引
};

struct list_item_i
{
	int     index;
	int     iSubItem;
	int     nUserData;
	//	BOOL    bSelect;
	list_item_state_  nState;
	RECT    rcItem;
	HXCGUI  hLayout;     ///<布局对象
	template_info_i  *pInfo;   ///<模板信息
};

struct list_header_item_i
{
	int       nState;
	int       nUserData;
	BOOL      bSelect;
	RECT      rcItem;
};

struct tree_item_i
{
	int     nID;
	int     nDepth;
	tree_item_state_  nState;
	int     nHeight;
	int     nSelHeight;
	int     nUserData;
	BOOL    bExpand;
	RECT    rcItem;
	HXCGUI  hLayout;     ///<布局对象
	template_info_i  *pInfo;   ///<模板信息
};

struct listView_item_i
{
	int     iGroup;             ///<项所述组索引 -1没有组
	int     iItem;              ///<项在数组中位置索引,如果此致为-1,那么为组
	int     nUserData;           ///<用户绑定数据
	list_item_state_  nState;    ///<状态  
	RECT    rcItem;              ///<整个区域,包含边框
	HXCGUI  hLayout;           ///<布局对象
	template_info_i  *pInfo;   ///<模板信息
};

//布局信息
struct layout_info_i
{
	layout_size_type_  widthType;
	layout_size_type_  heightType;
	short  width;    //宽度
	short  height;   //高度
};

struct  menu_popupWnd_i 
{
	HWINDOW hWindow;    //窗口句柄
	int     nParentID;  //父项ID
};

///菜单背景自绘结构
struct menu_drawBackground_i
{
	HMENUX  hMenu;      ///<菜单句柄
	HWINDOW hWindow;    ///<当前弹出菜单项的窗口句柄
	int     nParentID;  ///<父项ID
};

///菜单项自绘结构
struct  menu_drawItem_i
{
	HMENUX     hMenu;       ///<菜单句柄
	HWINDOW    hWindow;     ///<当前弹出菜单项的窗口句柄
	int        nID;         ///<ID
	int        nState;	    ///<状态 @ref menu_state_flags_
	RECT       rcItem;      ///<坐标
	HIMAGE     hIcon;       ///<菜单项图标
	const wchar_t    *pText;   ///<文本
};

///背景形状类型
enum  bkInfo_type_  
{
	bkInfo_type_fill=1,          ///<填充
	bkInfo_type_border,          ///<边框
	bkInfo_type_image,           ///<图片
	bkInfo_type_triangle,        ///<三角形
	bkInfo_type_fill_triangle,   ///<填充三角形
	bkInfo_type_ellipse,         ///<圆形
	bkInfo_type_fill_Ellipse,    ///<填充圆形
	bkInfo_type_roundRect,       ///<圆角矩形
	bkInfo_type_fill_RoundRect,  ///<填充圆角矩形
	bkInfo_type_line,            ///<直线
	bkInfo_type_gradient,        ///<渐变
};



///////////////////////////////////////////////////////////////////////////////////
/////////////////////窗口消息//////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////

//窗口消息
/**
@addtogroup wndMSG

<hr>
<h2>Windows 标准消息</h2>
- @ref WM_PAINT 窗口绘制消息
- @ref WM_CLOSE 窗口关闭消息.
- @ref WM_DESTROY 窗口销毁消息.
- @ref WM_NCDESTROY 窗口非客户区销毁消息.
- @ref WM_MOUSEMOVE 窗口鼠标移动消息.
- @ref WM_LBUTTONDOWN 窗口鼠标左键按下消息
- @ref WM_LBUTTONUP 窗口鼠标左键弹起消息.
- @ref WM_RBUTTONDOWN 窗口鼠标右键按下消息.
- @ref WM_RBUTTONUP 窗口鼠标右键弹起消息.
- @ref WM_LBUTTONDBLCLK 窗口鼠标左键双击消息.
- @ref WM_RBUTTONDBLCLK 窗口鼠标右键双击消息.
- @ref WM_MOUSEWHEEL 窗口鼠标滚轮滚动消息.
- @ref WM_EXITSIZEMOVE 窗口退出移动或调整大小模式循环改，详情参见MSDN.
- @ref WM_MOUSEHOVER 窗口鼠标进入消息
- @ref WM_MOUSELEAVE 窗口鼠标离开消息.
- @ref WM_SIZE 窗口大小改变消息.
- @ref WM_TIMER 窗口定时器消息.
- @ref WM_SETFOCUS 窗口获得焦点.
- @ref WM_KILLFOCUS 窗口失去焦点.
- @ref WM_KEYDOWN 窗口键盘按键消息.
- @ref WM_CAPTURECHANGED 窗口鼠标捕获改变消息.
- @ref WM_SETCURSOR  窗口设置鼠标光标.
- @ref WM_CHAR       窗口字符消息.
- @ref WM_DROPFILES  拖动文件到窗口.
- @ref other 其他Windows系统消息,包含你自定义的Windows消息.

<hr>
@anchor WM_PAINT WM_PAINT 窗口绘制消息
@code   int CALLBACK OnWndDrawWindow(HDRAW hDraw,BOOL *pbHandled); @endcode
@param  hDraw   图形绘制句柄.

<hr>
@anchor WM_CLOSE WM_CLOSE 窗口关闭消息.
@code   int CALLBACK OnWndClose(BOOL *pbHandled); @endcode

<hr>
@anchor WM_DESTROY WM_DESTROY 窗口销毁消息.
@code   int CALLBACK OnWndDestroy(BOOL *pbHandled);  @endcode

<hr>
@anchor WM_NCDESTROY WM_NCDESTROY 窗口非客户区销毁消息.
@code   int CALLBACK OnWndNCDestroy(BOOL *pbHandled); @endcode

<hr>
@anchor WM_MOUSEMOVE WM_MOUSEMOVE 窗口鼠标移动消息.
@code   int CALLBACK OnWndMouseMove(UINT nFlags,POINT *pPt,BOOL *pbHandled);  @endcode
@param  flags   请参见MSDN WM_MOUSEMOVE wParam参数.
@param  pPt     鼠标坐标点.

<hr>
@anchor WM_LBUTTONDOWN WM_LBUTTONDOWN 窗口鼠标左键按下消息
@code   int CALLBACK OnWndLButtonDown(UINT nFlags,POINT *pPt,BOOL *pbHandled); @endcode
@param  flags   请参见MSDN WM_LBUTTONDOWN.
@param  pPt     鼠标坐标点.

<hr>
@anchor WM_LBUTTONUP WM_LBUTTONUP 窗口鼠标左键弹起消息.
@code   int CALLBACK OnWndLButtonUp(UINT nFlags,POINT *pPt,BOOL *pbHandled); @endcode
@param  flags   请参见MSDN WM_LBUTTONUP.
@param  pPt     鼠标坐标点.

<hr>
@anchor WM_RBUTTONDOWN WM_RBUTTONDOWN 窗口鼠标右键按下消息.
@code   int CALLBACK OnWndRButtonDown(UINT nFlags,POINT *pPt,BOOL *pbHandled); @endcode
@param  flags   请参见MSDN WM_RBUTTONDOWN.
@param  pPt     鼠标坐标点.

<hr>
@anchor WM_RBUTTONUP WM_RBUTTONUP 窗口鼠标右键弹起消息.
@code   int CALLBACK OnWndRButtonUp(UINT nFlags,POINT *pPt,BOOL *pbHandled); @endcode
@param  flags   请参见MSDN WM_RBUTTONUP.
@param  pPt     鼠标坐标点.

<hr>
@anchor WM_LBUTTONDBLCLK WM_LBUTTONDBLCLK 窗口鼠标左键双击消息.
@code   int CALLBACK OnWndLButtonDBClick(UINT nFlags,POINT *pPt,BOOL *pbHandled); @endcode
@param  flags   请参见MSDN WM_LBUTTONDBLCLK.
@param  pPt     鼠标坐标点.

<hr>
@anchor WM_RBUTTONDBLCLK WM_RBUTTONDBLCLK 窗口鼠标右键双击消息.
@code   int CALLBACK OnWndRButtonDBClick(UINT nFlags,POINT *pPt,BOOL *pbHandled); @endcode
@param  flags   请参见MSDN WM_RBUTTONDBLCLK.
@param  pPt     鼠标坐标点.

<hr>
@anchor WM_MOUSEWHEEL WM_MOUSEWHEEL 窗口鼠标滚轮滚动消息.
@code   int CALLBACK OnWndMouseWheel(UINT nFlags,POINT *pPt,BOOL *pbHandled); @endcode
@param  flags   请参见MSDN WM_MOUSEWHEEL消息wParam参数.
@param  pPt     鼠标坐标点.

<hr>
@anchor WM_EXITSIZEMOVE WM_EXITSIZEMOVE 窗口退出移动或调整大小模式循环改，详情参见MSDN.
@code   int CALLBACK OnWndExitSizeMove(BOOL *pbHandled); @endcode

<hr>
@anchor WM_MOUSEHOVER WM_MOUSEHOVER 窗口鼠标进入消息.
@code   int CALLBACK OnWndMouseHover(UINT nFlags,POINT *pPt,BOOL *pbHandled); @endcode
@param  flags  请参见MSDN WM_MOUSEHOVER消息wParam参数.
@param  pPt    鼠标位置

<hr>
@anchor WM_MOUSELEAVE WM_MOUSELEAVE 窗口鼠标离开消息.
@code   int CALLBACK OnWndMouseLeave(BOOL *pbHandled); @endcode

<hr>
@anchor WM_SIZE WM_SIZE 窗口大小改变消息.
@code   int CALLBACK OnWndSize(UINT nFlags,SIZE *pSize,BOOL *pbHandled); @endcode
@param  flags   请参见MSDN WM_SIZE消息wParam参数.
@param  pSize   窗口大小.

<hr>
@anchor WM_TIMER WM_TIMER 窗口定时器消息.
@code   int CALLBACK OnWndTimer(UINT nIDEvent,BOOL *pbHandled); @endcode
@param  nIDEnent 定时器标示符.
@param  uElapse  指定所经过的系统启动以来的毫秒数,这是由GetTickCount函数的返回值.

<hr>
@anchor WM_SETFOCUS WM_SETFOCUS 窗口获得焦点.
@code   int CALLBACK OnWndSetFocus(BOOL *pbHandled); @endcode

<hr>
@anchor WM_KILLFOCUS WM_KILLFOCUS 窗口失去焦点.
@code   int CALLBACK OnWndKillFocus(BOOL *pbHandled); @endcode

<hr>
@anchor WM_KEYDOWN WM_KEYDOWN 窗口键盘按键消息.
@code   int CALLBACK OnWndKeyDown(WPARAM wParam,LPARAM lParam,BOOL *pbHandled); @endcode
@note   wParam,lParam:请参见MSDN WM_KEYDOWN.

<hr>
@anchor WM_CAPTURECHANGED WM_CAPTURECHANGED 窗口鼠标捕获改变消息.
@code   int CALLBACK OnWndCaptureChanged(HWND hWnd,BOOL *pbHandled); @endcode
@param  hWnd    获得鼠标捕获的窗口句柄.

<hr>
@anchor  WM_SETCURSOR WM_SETCURSOR 窗口设置鼠标光标.
@code    int CALLBACK OnWndSetCursor(WPARAM wParam,LPARAM lParam,BOOL *pbHandled); @endcode
@note    wParam,lParam:请参见MSDN WM_SETCURSOR.

<hr>
@anchor  WM_CHAR WM_CHAR 窗口字符输入消息.
@code    int CALLBACK OnWndChar(WPARAM wParam,LPARAM lParam,BOOL *pbHandled); @endcode
@note    wParam,lParam:请参见MSDN WM_CHAR.

<hr>
@anchor  WM_DROPFILES WM_DROPFILES 拖动文件到窗口消息.
@code    int CALLBACK OnDropFiles(HDROP hDropInfo , BOOL *pbHandled); @endcode
@note    请参见MSDN WM_DROPFILES.

<hr>
@anchor other 其他Windows系统消息,包含你自定义的Windows消息.
@code   int CALLBACK OnWndOther(WPARAM wParam,LPARAM lParam,BOOL *pbHandled); @endcode
@param  wParam  消息参数.
@param  lParam  消息参数.
@note   参数信息请参见MSDN.
@{
*/

//wParam:left-top坐标组合; lParam:right-bottom坐标组合; 如果这2个参数为空,那么重绘整个窗口
#define  XWM_REDRAW            WM_APP+1007  //窗口重绘延时 ----不公开-----内部自定义消息

//重绘元素,内部使用
#define  XWM_REDRAW_ELE       0x7000+3 //重绘元素 wParam:元素句柄, lParam:RECT*基于窗口坐标

//窗口消息-包含系统非客户区消息
#define  XWM_EVENT_ALL        WM_APP+1000 //事件投递 -------不公开-------不需要注册

#define  XWM_WINDPROC         0x7000    //注册窗口处理过程
#define  XWM_DRAW_T           0x7000+1  //窗口绘制,内部使用, wParam:0, lParam:0


///@brief 菜单弹出
///@code  int CALLBACK OnWndMenuPopup(HMENUX hMenu, BOOL *pbHandled); @endcode
#define  XWM_MENU_POPUP      0x7000+5

///@brief 菜单弹出窗口
///@code  int CALLBACK OnWndMenuPopupWnd(HMENUX hMenu,menu_popupWnd_i *pInfo,BOOL *pbHandled); @endcode
#define  XWM_MENU_POPUP_WND     0x7000+6    

///@brief 菜单选择
///@code  int CALLBACK OnWndMenuSelect(int nID,BOOL *pbHandled); @endcode
#define  XWM_MENU_SELECT       0x7000+7 //菜单选择 wParam:菜单项ID, lParam:0

///@brief 菜单退出
///@code  int CALLBACK OnWndMenuExit(BOOL *pbHandled); @endcode
#define  XWM_MENU_EXIT         0x7000+8 //菜单退出 wParam:0, lParam:0

///@brief 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
///@code  int CALLBACK OnWndMenuDrawBackground(HDRAW hDraw,menu_drawBackground_i *pInfo,BOOL *pbHandled); @endcode
#define  XWM_MENU_DRAW_BACKGROUND   0x7000+9 //绘制菜单背景

/// @brief 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
/// @code  int CALLBACK OnMenuDrawItem(HDRAW hDraw,menu_drawItem_i* pInfo,BOOL *pbHandled); @endcode
#define  XWM_MENU_DRAWITEM    0x7000+10  //绘制菜单项 

#define  XWM_COMBOBOX_POPUP_DROPLIST   0x7000+11  //弹出下拉组框列表,内部使用

///@}

///////////////////////////////////////////////////////////////////////////////////
/////////////////元素事件//////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////
//元素事件
/// @addtogroup eleEvents
/// @{


/// @brief 元素处理过程事件.
/// @code  int CALLBACK OnEventProc(HELE hEventEle, UINT nEvent, WPARAM wParam, LPARAM lParam, BOOL *pbHandled) @endcode
#define  XE_ELEPROCE         1     ///元素事件处理过程

/// @brief 元素绘制事件
/// @code  int CALLBACK OnDraw(HDRAW hDraw,BOOL *pbHandled);  @endcode
#define  XE_PAINT            2

/// @brief 该元素及子元素绘制完成事件.启用该功能需要调用XEle_EnableEvent_XE_PAINT_END();
/// @code  int CALLBACK OnPaintEnd(HDRAW hDraw,BOOL *pbHandled); @endcode
#define  XE_PAINT_END        3     //元素及子元素绘制完成后触发,需要启用该功能,XEle_EnableEvent_XE_PAINT_END();

/// @brief 滚动视图绘制事件.
/// @code int CALLBACK OnDrawScrollView(HDRAW hDraw,BOOL *pbHandled); @endcode
#define  XE_PAINT_SCROLLVIEW   4

/// @brief 元素鼠标移动事件.
/// @code  int CALLBACK OnMouseMove(UINT nFlags, POINT *pPt, BOOL *pbHandled);  @endcode
#define  XE_MOUSEMOVE        5

/// @brief 元素鼠标进入事件.
/// @code  int CALLBACK OnMouseStay(BOOL *pbHandled); @endcode
#define  XE_MOUSESTAY        6   ///停留

/// @brief 元素鼠标悬停事件.
/// @code  int CALLBACK OnMouseHover(UINT nFlags, POINT *pPt, BOOL *pbHandled); @endcode
#define  XE_MOUSEHOVER       7   ///悬停

/// @brief 元素鼠标离开事件.
/// @param  nFlags 将获得鼠标停留状态的元素.
/// @code  int CALLBACK OnMouseLeave(HELE hEleStay,BOOL *pbHandled); @endcode
#define  XE_MOUSELEAVE       8   ///离开

/// @brief 元素鼠标滚轮滚动事件.
/// @code  int CALLBACK OnMouseWheel(UINT nFlags,POINT *pPt,BOOL *pbHandled); @endcode
/// @param flags      见MSDN中WM_MOUSEWHEEL消息wParam参数说明.
#define  XE_MOUSEWHEEL         9  //鼠标滚轮  wParam:标识,lParam:POINT坐标

/// @brief 鼠标左键按下事件.
/// @code  int CALLBACK OnLButtonDown(UINT nFlags, POINT *pPt,BOOL *pbHandled);  @endcode
#define  XE_LBUTTONDOWN        10

/// @brief 鼠标左键弹起事件.
/// @code  int CALLBACK OnLButtonUp(UINT nFlags, POINT *pPt,BOOL *pbHandled); @endcode
#define  XE_LBUTTONUP          11

/// @brief 鼠标右键按下事件.
/// @code  int CALLBACK OnRButtonDOWN(UINT nFlags, POINT *pPt,BOOL *pbHandled);  @endcode
#define  XE_RBUTTONDOWN        12

/// @brief 鼠标右键弹起事件.
/// @code int CALLBACK OnRButtonUp(UINT nFlags, POINT *pPt,BOOL *pbHandled);  @endcode
#define  XE_RBUTTONUP          13

/// @brief 鼠标左键双击事件.
/// @code   int CALLBACK OnLButtonDBClick(UINT nFlags, POINT *pPt,BOOL *pbHandled); @endcode
#define  XE_LBUTTONDBCLICK     14

#define  XE_RBUTTONDBCLICK     15

/// @brief 元素获得焦点事件.
/// @code int CALLBACK OnSetFocus(BOOL *pbHandled); @endcode
#define  XE_SETFOCUS           31

/// @brief 元素失去焦点事件.
/// @code int CALLBACK OnKillFocus(BOOL *pbHandled); @endcode
#define  XE_KILLFOCUS          32

/// @brief 元素销毁事件.
/// @code int CALLBACK OnDestroy(BOOL *pbHandled); @endcode
#define  XE_DESTROY            33   //元素销毁

/// @brief 按钮点击事件.
/// @code int CALLBACK OnBtnClick(BOOL *pbHandled); @endcode
#define  XE_BNCLICK            34

/// @brief 按钮选中事件.
/// @code int CALLBACK OnButtonCheck(BOOL bCheck,BOOL *pbHandled); @endcode
#define  XE_BUTTON_CHECK       35  //按钮选中事件

/// @brief 元素大小改变事件.
/// @code int CALLBACK *OnSize(BOOL *pbHandled); @endcode
#define  XE_SIZE               36

/// @brief 元素显示隐藏事件.
/// @code  int CALLBACK OnShow(BOOL bShow,BOOL *pbHandled); @endcode
#define  XE_SHOW               37  //wParam:TRUE或FALSE, lParam:0

/// @brief 元素设置字体事件.
/// @code  int CALLBACK OnSetFont(BOOL *pbHandled); @endcode
#define  XE_SETFONT           38

/// @brief 元素按键事件.
/// @code  int CALLBACK OnEventKeyDown(WPARAM wParam,LPARAM lParam,BOOL *pbHandled); @endcode
/// @note  wParam,lParam:请参见MSDN WM_KEYDOWN. 
#define  XE_KEYDOWN            39 //wParam和lParam参数与标准消息相同

/// @brief 通过TranslateMessage函数翻译的字符事件.
/// @code  int CALLBACK OnEventChar(WPARAM wParam,LPARAM lParam,BOOL *pbHandled); @endcode
/// @note   wParam,lParam:请参见MSDN WM_KEYDOWN. 
#define  XE_CHAR               40   //wParam和lParam参见MSDN

/// @brief 元素设置鼠标捕获.
/// @code int CALLBACK OnSetCapture(BOOL *pbHandled); @endcode
#define  XE_SETCAPTURE      51

/// @brief 元素失去鼠标捕获.
/// @code int CALLBACK OnKillCapture(BOOL *pbHandled); @endcode
#define  XE_KILLCAPTURE     52

/// @brief 设置鼠标光标
/// @code int CALLBACK OnSetCursor(WPARAM wParam,LPARAM lParam,BOOL *pbHandled) @endcode
#define  XE_SETCURSOR            53  // SetCursor

/// @brief 滚动视图元素水平滚动事件,滚动视图触发.
/// @code  int CALLBACK OnScrollViewScrollH(int pos,BOOL *pbHandled); @endcode
/// @param pos         当前滚动点.
#define  XE_SCROLLVIEW_SCROLL_H    54  //滚动视图 滚动事件 wParam:滚动点,lParam:0 (滚动视图触发,表明滚动视图已滚动完成)

/// @brief 滚动视图元素垂直滚动事件,滚动视图触发.
/// @code  int CALLBACK OnScrollViewScrollV(int pos,BOOL *pbHandled); @endcode
/// @param pos         当前滚动点.
#define  XE_SCROLLVIEW_SCROLL_V    55  //滚动视图 滚动事件 wParam:滚动点,lParam:0 (滚动视图触发,表明滚动视图已滚动完成)

/// @brief 滚动条元素滚动事件,滚动条触发.
/// @code  int CALLBACK OnSBarScroll(int pos,BOOL *pbHandled); @endcode
/// @param pos   当前滚动点.
#define  XE_SBAR_SCROLL        56   //滚动条滚动事件 wParam:滚动点,lParam:0 (滚动条触发)

///@brief 菜单弹出
///@code  int CALLBACK OnWndMenuPopup(HMENUX hMenu, BOOL *pbHandled); @endcode
#define  XE_MENU_POPUP       57

///@brief 菜单弹出窗口
///@code  int CALLBACK OnWndMenuPopupWnd(HMENUX hMenu,menu_popupWnd_i* pInfo,BOOL *pbHandled); @endcode
#define  XE_MENU_POPUP_WND     58

/// @brief 弹出菜单项选择事件.
/// @code  int CALLBACK OnEventMenuSelect(int nItem,BOOL *pbHandled); @endcode
/// @param nItem          菜单项id.
#define  XE_MENU_SELECT      59  //菜单项选择 wParam:菜单ID,lParam:0

///@brief 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
///@code  int CALLBACK OnMenuDrawBackground(HDRAW hDraw,menu_drawBackground_i *pInfo,BOOL *pbHandled); @endcode
#define  XE_MENU_DRAW_BACKGROUND   60 //绘制菜单背景

/// @brief 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
/// @code  int CALLBACK OnMenuDrawItem(HDRAW hDraw,menu_drawItem_i* pInfo,BOOL *pbHandled); @endcode
#define  XE_MENU_DRAWITEM    61  //绘制菜单项 

/// @brief 弹出菜单退出事件.
/// @code  int CALLBACK OnMenuExit(BOOL *pbHandled); @endcode
#define  XE_MENU_EXIT        62 //菜单退出

/// @brief 滑动条元素,滑块位置改变事件.
/// @code  int CALLBACK OnSliderBarChange(int pos,BOOL *pbHandled); @endcode
#define  XE_SLIDERBAR_CHANGE   63

/// @brief 进度条元素,进度改变事件.
/// @code  int CALLBACK OnProgressBarChange(int pos,BOOL *pbHandled); @endcode
#define  XE_PROGRESSBAR_CHANGE  64

/// @brief 组合框下拉列表项选择事件.
/// @code  int CALLBACK OnComboBoxSelect(int iItem,BOOL *pbHandled);  @endcode
#define  XE_COMBOBOX_SELECT       71   //组合框项选择  

/// @brief 组合框下拉列表弹出事件.
/// @code  int CALLBACK OnComboBoxPopupList(HWINDOW hWindow,HELE hListBox,BOOL *pbHandled);  @endcode
#define  XE_COMBOBOX_POPUP_LIST   72

/// @brief 组合框下拉列表退出事件.
/// @code  int CALLBACK OnComboBoxExitList(BOOL *pbHandled);  @endcode
#define  XE_COMBOBOX_EXIT_LIST    73

/// @brief 列表框元素,项模板创建.
/// @code int CALLBACK OnListBoxTemplateCreate(listBox_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_LISTBOX_TEMP_CREATE     81

/// @brief 列表框元素,项模板创建完成事件.
/// @code int CALLBACK OnListBoxTemplateCreateEnd(listBox_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_LISTBOX_TEMP_CREATE_END     82

/// @brief 列表框元素,项模板销毁.
/// @code int CALLBACK OnListBoxTemplateDestroy(listBox_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_LISTBOX_TEMP_DESTROY    83

/// @brief 列表框元素,项模板调整坐标.
/// @code  int CALLBACK OnListBoxTemplateAdjustCoordinate(listBox_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_LISTBOX_TEMP_ADJUST_COORDINATE  84

/// @brief 列表框元素,项绘制事件.
/// @code  int CALLBACK OnListBoxDrawItem(HDRAW hDraw,listBox_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_LISTBOX_DRAWITEM     85

/// @brief 列表框元素,项选择事件.
/// @code  int CALLBACK OnListBoxSelect(int iItem,BOOL *pbHandled);  @endcode
#define  XE_LISTBOX_SELECT       86

/// @brief 列表元素,项模板创建.
/// @code  int CALLBACK OnListTemplateCreate(list_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_LIST_TEMP_CREATE     101

/// @brief 列表元素,项模板创建完成事件.
/// @code  int CALLBACK OnListTemplateCreateEnd(list_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_LIST_TEMP_CREATE_END     102

/// @brief 列表元素,项模板销毁.
/// @code int CALLBACK OnListTemplateDestroy(list_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_LIST_TEMP_DESTROY    103

/// @brief 列表元素,项模板调整坐标.
/// @code  typedef int CALLBACK OnListTemplateAdjustCoordinate(list_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_LIST_TEMP_ADJUST_COORDINATE  104

/// @brief 列表元素,绘制项.
/// @code  int CALLBACK OnListDrawItem(HDRAW hDraw,list_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_LIST_DRAWITEM                105

/// @brief 列表元素,项选择事件.
/// @code  int CALLBACK OnListSelect(int iItem,BOOL *pbHandled); @endcode
#define  XE_LIST_SELECT                  106

/// @brief 列表元素绘制列表头项.
/// @code  int CALLBACK OnListHeaderDrawItem(HDRAW hDraw, list_header_item_i* pItem, BOOL *pbHandled); @endcode
#define  XE_LIST_HEADER_DRAWITEM         107

#define  XE_LIST_HEADER_CLICK            108

#define  XE_LIST_HEADER_WIDTH_CHANGE     109

/// @brief 树元素,项模板创建.
/// @code  int CALLBACK OnTreeTemplateCreate(tree_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_TREE_TEMP_CREATE             121  

/// @brief 树元素,项模板创建完成.
/// @code  int CALLBACK OnTreeTemplateCreateEnd(tree_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_TREE_TEMP_CREATE_END         122

/// @brief 树元素,项模板销毁.
/// @code  int CALLBACK OnTreeTemplateDestroy(tree_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_TREE_TEMP_DESTROY            123 

/// @brief 树元素,项模板,调整项坐标.
/// @code  int CALLBACK OnTreeTemplateAdjustCoordinate(tree_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_TREE_TEMP_ADJUST_COORDINATE  124  

/// @brief 树元素,绘制项.
/// @code  int CALLBACK OnTreeDrawItem(HDRAW hDraw,tree_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_TREE_DRAWITEM                125  

/// @brief 树元素,项选择事件.
/// @code  int CALLBACK OnTreeSelect(int nItemID,BOOL *pbHandled); @endcode
/// @param nItemID  项ID.
#define  XE_TREE_SELECT                 126

/// @brief 树元素,项展开收缩事件.
/// @code int CALLBACK OnTreeExpand(int id,BOOL bExpand,BOOL *pbHandled); @endcode
#define  XE_TREE_EXPAND                 127

/// @brief 列表视元素,项模板创建.
/// @code  int CALLBACK OnListViewTemplateCreate(listView_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_LISTVIEW_TEMP_CREATE           141

/// @brief 列表视元素,项模板创建完成事件.
/// @code  int CALLBACK OnListViewTemplateCreateEnd(listView_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_LISTVIEW_TEMP_CREATE_END           142

/// @brief 列表视元素,项模板销毁.
/// @code  int CALLBACK OnListViewTemplateDestroy(listView_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_LISTVIEW_TEMP_DESTROY           143

/// @brief 列表视元素,项模板调整坐标.
/// @code  int CALLBACK OnListViewTemplateAdjustCoordinate(listView_item_i* pItem,BOOL *pbHandled); @endcode
#define  XE_LISTVIEW_TEMP_ADJUST_COORDINATE   144

/// @brief 列表视元素,自绘项.
/// @code int CALLBACK OnListViewDrawItem(HDRAW hDraw,listView_item_i* pItem,BOOL *pbHandled);  @endcode
#define  XE_LISTVIEW_DRAWITEM              145

/// @brief 列表视元素,项选择事件.
/// @code int CALLBACK OnListViewSelect(int iGroup,int iItem,BOOL *pbHandled); @endcode
#define  XE_LISTVIEW_SELECT            146

/// @brief 列表视元素,组展开收缩事件.
/// @code int CALLBACK OnListViewExpand(int iGroup,BOOL bExpand,BOOL *pbHandled); @endcode
#define  XE_LISTVIEW_EXPAND         147


/// @brief 属性网格元素 项值改变事件
/// @code   int CALLBACK OnPGridValueChange(int nItemID,BOOL *pbHandled); @endcode
/// @param  nItemID  项ID.
#define  XE_PGRID_VALUE_CHANGE     161

/// @brief 富文本元素 用户修改内容事件,只有当用户操作时才会触发,需要开启后有效, XRichEdit_EnableEvent_XE_RICHEDIT_CHANGE()；
/// XRichEdit_SetText()、XRichEdit_InsertString()不会触发此事件
/// @code   int CALLBACK OnRichEditChange(BOOL *pbHandled); @endcode
#define  XE_RICHEDIT_CHANGE        162

//int CALLBACK OnRichEditSet(BOOL *pbHandled);
#define XE_RICHEDIT_SET     201

//int CALLBACK OnPGridItemSet(int nItemID, BOOL *pbHandled);
#define XE_PGRID_ITEM_SET   202

/// @brief  TabBar标签按钮选择改变事件
/// @code   int CALLBACK OnTabBarSelect(int iItem, BOOL *pbHandled); @endcode
/// @param  iItem  标签位置索引.
#define  XE_TABBAR_SELECT    221

/// @brief  TabBar标签按钮删除事件
/// @code   int CALLBACK OnTabBarDelete(int iItem, BOOL *pbHandled); @endcode
/// @param  iItem  标签位置索引.
#define  XE_TABBAR_DELETE    222

/// @brief  月历元素日期改变事件, 未开放
/// @code   int CALLBACK OnCalendarChange(BOOL *pbHandled); @endcode
#define  XE_MONTHCAL_CHANGE   231   //月历 日期改变事件

/// @brief  日期时间元素,内容改变事件, 未开放
/// @code   int CALLBACK OnDateTimeChange(BOOL *pbHandled); @endcode
#define  XE_DATETIME_CHANGE    241   //日期时间元素  改变事件

/// @brief  日期时间元素,弹出月历卡片事件, 未开放
/// @code   int CALLBACK OnDateTimePopupMonthCal(HWINDOW hMonthCalWnd,HELE hMonthCal,BOOL *pbHandled); @endcode
#define  XE_DATETIME_POPUP_MONTHCAL     242   //日期时间元素  弹出月历事件

/// @brief  日期时间元素,弹出的月历卡片退出事件, 未开放
/// @code   int CALLBACK OnDateTimeExitMonthCal(HWINDOW hMonthCalWnd,HELE hMonthCal,BOOL *pbHandled); @endcode
#define  XE_DATETIME_EXIT_MONTHCAL      243   //日期时间元素  退出月历事件

/// @brief  文件拖放事件.
/// @code   int CALLBACK OnDropFiles(HDROP hDropInfo, BOOL *pbHandled); @endcode
#define  XE_DROPFILES                 250  

#define  XE_LISTVIEW_DRAG_INSERT
#define  XE_PANE_LOACK
#define  XE_PANE_DOCK
#define  XE_PANE_FLOAT
#define  XE_EDIT_CHANGE

///@}

///////////////////////////////////////////////////////////////////////////////////
/////////////事件注册//////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////
#define XC_PARAM_int            1
#define XC_PARAM_uint           2
#define XC_PARAM_BOOL           3
#define XC_PARAM_CHAR           4
#define XC_PARAM_BYTE           5
#define XC_PARAM_HXCGUI         6
#define XC_PARAM_HWINDOW        7
#define XC_PARAM_HELE           8
#define XC_PARAM_HMENUX         9
#define XC_PARAM_HIMAGE         10
#define XC_PARAM_HDRAW          11
#define XC_PARAM_P_RECT         12
#define XC_PARAM_P_POINT        13
#define XC_PARAM_P_SIZE         14
#define XC_PARAM_P_BOOL         15
#define XC_PARAM_P_listBox_item_i      16
#define XC_PARAM_P_list_header_item_i  17
#define XC_PARAM_P_list_item_i         18
#define XC_PARAM_P_listView_item_i     19
#define XC_PARAM_P_tree_item_i         20
#define XC_PARAM_P_menu_popupWnd       21
#define XC_PARAM_P_menu_drawBackground 22
#define XC_PARAM_P_menu_drawItem       23
#define XC_PARAM_HDROP   24



typedef int(CALLBACK *pFun_int_uint_uint_uint_bPtr)(UINT, WPARAM, LPARAM, BOOL*);
typedef int(CALLBACK *pFun_int_w_uint_uint_uint_bPtr)(HWINDOW, UINT, WPARAM, LPARAM, BOOL*); //窗口
typedef int(CALLBACK *pFun_int_e_uint_uint_uint_bPtr)(HELE, UINT, WPARAM, LPARAM, BOOL*); //元素
typedef int(CALLBACK *pFun_int_e_e_uint_uint_uint_bPtr)(HELE, HELE, UINT, WPARAM, LPARAM, BOOL*); //元素

typedef int(CALLBACK *pFun_int_uint_ptPtr_bPtr)(UINT, POINT*, BOOL*);
typedef int(CALLBACK *pFun_int_w_uint_ptPtr_bPtr)(HWINDOW, UINT, POINT*, BOOL*);
typedef int(CALLBACK *pFun_int_e_uint_ptPtr_bPtr)(HELE, UINT, POINT*, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_uint_ptPtr_bPtr)(HELE, HELE, UINT, POINT*, BOOL*);

typedef int(CALLBACK *pFun_int_hDraw_bPtr)(HDRAW, BOOL*);
typedef int(CALLBACK *pFun_int_w_hDraw_bPtr)(HWINDOW, HDRAW, BOOL*);
typedef int(CALLBACK *pFun_int_e_hDraw_bPtr)(HELE, HDRAW, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_hDraw_bPtr)(HELE, HELE, HDRAW, BOOL*);

typedef int(CALLBACK *pFun_int_bPtr)(BOOL*);
typedef int(CALLBACK *pFun_int_w_bPtr)(HWINDOW, BOOL*);
typedef int(CALLBACK *pFun_int_e_bPtr)(HELE, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_bPtr)(HELE, HELE, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_e_bPtr)(HELE, HELE, BOOL*);

typedef int(CALLBACK *pFun_int_w_e_bPtr)(HWINDOW, HELE, BOOL*);
typedef int(CALLBACK *pFun_int_e_w_e_bPtr)(HELE, HWINDOW, HELE, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_w_e_bPtr)(HELE, HELE, HWINDOW, HELE, BOOL*);

typedef int(CALLBACK *pFun_int_uint_sizePtr_bPtr)(UINT,SIZE*,BOOL*);
typedef int(CALLBACK *pFun_int_w_uint_sizePtr_bPtr)(HWINDOW, UINT,SIZE*,BOOL*);


typedef int(CALLBACK *pFun_int_hMenu_bPtr)(HMENUX, BOOL*);
typedef int(CALLBACK *pFun_int_w_hMenu_bPtr)(HWINDOW, HMENUX, BOOL*);
typedef int(CALLBACK *pFun_int_e_hMenu_bPtr)(HELE, HMENUX, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_hMenu_bPtr)(HELE, HELE, HMENUX, BOOL*);

typedef int(CALLBACK *pFun_int_hMenu_pMenuPopupWnd_bPtr)(HMENUX, menu_popupWnd_i*, BOOL*);
typedef int(CALLBACK *pFun_int_w_hMenu_pMenuPopupWnd_bPtr)(HWINDOW, HMENUX, menu_popupWnd_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_hMenu_pMenuPopupWnd_bPtr)(HELE, HMENUX, menu_popupWnd_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_hMenu_pMenuPopupWnd_bPtr)(HELE, HELE, HMENUX, menu_popupWnd_i*, BOOL*);

typedef int(CALLBACK *pFun_int_hDraw_pMenuDrawBackground_bPtr)(HDRAW, menu_drawBackground_i*, BOOL*);
typedef int(CALLBACK *pFun_int_w_hDraw_pMenuDrawBackground_bPtr)(HWINDOW, HDRAW, menu_drawBackground_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_hDraw_pMenuDrawBackground_bPtr)(HELE, HDRAW, menu_drawBackground_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_hDraw_pMenuDrawBackground_bPtr)(HELE, HELE, HDRAW, menu_drawBackground_i*, BOOL*);

typedef int(CALLBACK *pFun_int_hDraw_pMenuDrawItem_bPtr)(HDRAW, menu_drawItem_i*, BOOL*);
typedef int(CALLBACK *pFun_int_w_hDraw_pMenuDrawItem_bPtr)(HWINDOW, HDRAW, menu_drawItem_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_hDraw_pMenuDrawItem_bPtr)(HELE, HDRAW, menu_drawItem_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_hDraw_pMenuDrawItem_bPtr)(HELE, HELE, HDRAW, menu_drawItem_i*, BOOL*);

typedef int(CALLBACK *pFun_int_uint_uint_bPtr)(WPARAM, LPARAM, BOOL*); //元素
typedef int(CALLBACK *pFun_int_w_uint_uint_bPtr)(HWINDOW, WPARAM, LPARAM, BOOL*);

typedef int(CALLBACK *pFun_int_uint_bPtr)(UINT, BOOL*);

typedef int(CALLBACK *pFun_int_int_int_bPtr)(int, int, BOOL*);
typedef int(CALLBACK *pFun_int_e_int_int_bPtr)(HELE, int, int, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_int_int_bPtr)(HELE, HELE, int, int, BOOL*);

typedef int(CALLBACK *pFun_int_int_bPtr)(int, BOOL*);
typedef int(CALLBACK *pFun_int_w_int_bPtr)(HWINDOW, int, BOOL*);
typedef int(CALLBACK *pFun_int_e_int_bPtr)(HELE, int, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_int_bPtr)(HELE, HELE, int, BOOL*);

typedef int(CALLBACK *pFun_int_pListBoxItem_bPtr)(listBox_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_pListBoxItem_bPtr)(HELE, listBox_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_pListBoxItem_bPtr)(HELE, HELE, listBox_item_i*, BOOL*);

typedef int(CALLBACK *pFun_int_hDraw_pListBoxItem_bPtr)(HDRAW, listBox_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_hDraw_pListBoxItem_bPtr)(HELE, HDRAW, listBox_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_hDraw_pListBoxItem_bPtr)(HELE, HELE, HDRAW, listBox_item_i*, BOOL*);

typedef int(CALLBACK *pFun_int_pListItem_bPtr)(list_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_pListItem_bPtr)(HELE, list_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_pListItem_bPtr)(HELE, HELE, list_item_i*, BOOL*);

typedef int(CALLBACK *pFun_int_hDraw_pListItem_bPtr)(HDRAW,list_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_hDraw_pListItem_bPtr)(HELE, HDRAW,list_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_hDraw_pListItem_bPtr)(HELE, HELE, HDRAW,list_item_i*, BOOL*);

typedef int(CALLBACK *pFun_int_hDraw_pListHeaderItem_bPtr)(HDRAW,list_header_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_hDraw_pListHeaderItem_bPtr)(HELE, HDRAW,list_header_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_hDraw_pListHeaderItem_bPtr)(HELE, HELE, HDRAW,list_header_item_i*, BOOL*);

typedef int(CALLBACK *pFun_int_pTreeItem_bPtr)(tree_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_pTreeItem_bPtr)(HELE, tree_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_pTreeItem_bPtr)(HELE, HELE, tree_item_i*, BOOL*);

typedef int(CALLBACK *pFun_int_hDraw_pTreeItem_bPtr)(HDRAW,tree_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_hDraw_pTreeItem_bPtr)(HELE, HDRAW,tree_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_hDraw_pTreeItem_bPtr)(HELE, HELE, HDRAW,tree_item_i*, BOOL*);

typedef int(CALLBACK *pFun_int_pListViewItem_bPtr)(listView_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_pListViewItem_bPtr)(HELE, listView_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_pListViewItem_bPtr)(HELE, HELE, listView_item_i*, BOOL*);

typedef int(CALLBACK *pFun_int_hDraw_pListViewItem_bPtr)(HDRAW,listView_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_hDraw_pListViewItem_bPtr)(HELE, HDRAW, listView_item_i*, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_hDraw_pListViewItem_bPtr)(HELE, HELE, HDRAW, listView_item_i*, BOOL*);

typedef int(CALLBACK *pFun_int_e_uint_uint_bPtr)(HELE, WPARAM, LPARAM, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_uint_uint_bPtr)(HELE, HELE, WPARAM, LPARAM, BOOL*);

typedef int(CALLBACK *pFun_int_hDrop_bPtr)(HDROP, BOOL*);
typedef int(CALLBACK *pFun_int_w_hDrop_bPtr)(HWINDOW, HDROP, BOOL*);
typedef int(CALLBACK *pFun_int_e_hDrop_bPtr)(HELE, HDROP, BOOL*);
typedef int(CALLBACK *pFun_int_e_e_hDrop_bPtr)(HELE, HELE, HDROP, BOOL*);

class CEventBase
{
public:
	CEventBase(){ m_nEvent=XC_ID_ERROR; m_pFunCStyle=NULL; m_regType=XC_ID_ERROR; }
	UINT     m_nEvent;      //事件类型
	void*    m_pFunCStyle;
	int      m_regType;   //注册类型, CPP,CPP1,CPP2

	virtual  int OnWndProc(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message = 0){ return 0; }
	virtual  int OnEleProc(HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent = 0){ return 0; }
	virtual  void Release(){ delete this; }
	virtual  void* GetThis(){ return NULL; }
	virtual  char* GetFunName(){ return NULL; }
};

template<typename T_ClassName>
class CEventType
{
public:
	typedef int(T_ClassName::*pFun_int_uint_uint_uint_bPtr)(UINT, WPARAM, LPARAM, BOOL*); //窗口
	typedef int(T_ClassName::*pFun_int_w_uint_uint_uint_bPtr)(HWINDOW, UINT, WPARAM, LPARAM, BOOL*); //窗口
	typedef int(T_ClassName::*pFun_int_e_uint_uint_uint_bPtr)(HELE, UINT, WPARAM, LPARAM, BOOL*); //元素
	typedef int(T_ClassName::*pFun_int_e_e_uint_uint_uint_bPtr)(HELE, HELE, UINT, WPARAM, LPARAM, BOOL*); //元素

	typedef int(T_ClassName::*pFun_int_uint_ptPtr_bPtr)(UINT, POINT*, BOOL*);
	typedef int(T_ClassName::*pFun_int_w_uint_ptPtr_bPtr)(HWINDOW, UINT, POINT*, BOOL*);

	typedef int(T_ClassName::*pFun_int_hDraw_bPtr)(HDRAW, BOOL*);
	typedef int(T_ClassName::*pFun_int_w_hDraw_bPtr)(HWINDOW, HDRAW, BOOL*);
	typedef int(T_ClassName::*pFun_int_e_hDraw_bPtr)(HELE, HDRAW, BOOL*);
	typedef int(T_ClassName::*pFun_int_e_e_hDraw_bPtr)(HELE, HELE, HDRAW, BOOL*);

	typedef int(T_ClassName::*pFun_int_bPtr)(BOOL*);
	typedef int(T_ClassName::*pFun_int_e_bPtr)(HELE, BOOL*);
	typedef int(T_ClassName::*pFun_int_e_e_bPtr)(HELE, HELE, BOOL*);
	typedef int(T_ClassName::*pFun_int_e_e_e_bPtr)(HELE, HELE, HELE, BOOL*);

	typedef int(T_ClassName::*pFun_intuint_ptPtr_bPtr)(UINT,POINT*,BOOL*);
	typedef int(T_ClassName::*pFun_int_e_uint_ptPtr_bPtr)(HELE, UINT, POINT*, BOOL*);
	typedef int(T_ClassName::*pFun_int_e_e_uint_ptPtr_bPtr)(HELE, HELE, UINT, POINT*, BOOL*);

	typedef int(T_ClassName::*pFun_int_int_bPtr)(int, BOOL*);
	typedef int(T_ClassName::*pFun_int_e_int_bPtr)(HELE, int, BOOL*);
	typedef int(T_ClassName::*pFun_int_e_e_int_bPtr)(HELE, HELE, int, BOOL*);

	typedef int(T_ClassName::*pFun_int_w_e_bPtr)(HWINDOW, HELE, BOOL*);
	typedef int(T_ClassName::*pFun_int_e_w_e_bPtr)(HELE, HWINDOW, HELE, BOOL*);
	typedef int(T_ClassName::*pFun_int_e_e_w_e_bPtr)(HELE, HELE, HWINDOW, HELE, BOOL*);

	typedef int(T_ClassName::*pFun_int_uint_bPtr)(UINT, BOOL*);

	typedef int(T_ClassName::*pFun_int_int_int_bPtr)(int, int, BOOL*);

	typedef int(T_ClassName::*pFun_int_uint_uint_bPtr)(WPARAM, LPARAM, BOOL*);
	typedef int(T_ClassName::*pFun_int_e_uint_uint_bPtr)(HELE, WPARAM, LPARAM, BOOL*);
	typedef int(T_ClassName::*pFun_int_w_uint_uint_bPtr)(HWINDOW, WPARAM, LPARAM, BOOL*);

	typedef int(T_ClassName::*pFun_int_uint_sizePtr_bPtr)(UINT,SIZE*,BOOL*);

	typedef int(T_ClassName::*pFun_int_pListBoxItem_bPtr)(listBox_item_i*, BOOL*);
	typedef int(T_ClassName::*pFun_int_e_pListBoxItem_bPtr)(HELE, listBox_item_i*, BOOL*);
	typedef int(T_ClassName::*pFun_int_e_e_pListBoxItem_bPtr)(HELE, HELE, listBox_item_i*, BOOL*);

	typedef int(T_ClassName::*pFun_int_hDraw_pListBoxItem_bPtr)(HDRAW, listBox_item_i*, BOOL*);

	typedef int(T_ClassName::*pFun_int_pListItem_bPtr)(list_item_i*, BOOL*);
	typedef int(T_ClassName::*pFun_int_hDraw_pListItem_bPtr)(HDRAW,list_item_i*, BOOL*);

	typedef int(T_ClassName::*pFun_int_hDraw_pListHeaderItem_bPtr)(HDRAW,list_header_item_i*, BOOL*);

	typedef int(T_ClassName::*pFun_int_pTreeItem_bPtr)(tree_item_i*, BOOL*);
	typedef int(T_ClassName::*pFun_int_hDraw_pTreeItem_bPtr)(HDRAW,tree_item_i*, BOOL*);

	typedef int(T_ClassName::*pFun_int_pListViewItem_bPtr)(listView_item_i*, BOOL*);
	typedef int(T_ClassName::*pFun_int_hDraw_pListViewItem_bPtr)(HDRAW,listView_item_i*, BOOL*);

	typedef int(T_ClassName::*pFun_int_hMenu_bPtr)(HMENUX, BOOL*);
	typedef int(T_ClassName::*pFun_int_hMenu_pMenuPopupWnd_bPtr)(HMENUX, menu_popupWnd_i*, BOOL*);
	typedef int(T_ClassName::*pFun_int_hDraw_pMenuDrawItem_bPtr)(HDRAW,menu_drawItem_i*,BOOL*);
	typedef int(T_ClassName::*pFun_int_hDraw_pMenuDrawBackground_bPtr)(HDRAW,menu_drawBackground_i*, BOOL*);

	typedef int(T_ClassName::*pFun_int_hDrop_bPtr)(HDROP, BOOL*);
	typedef int(T_ClassName::*pFun_int_w_hDrop_bPtr)(HWINDOW, HDROP, BOOL*);
	typedef int(T_ClassName::*pFun_int_e_hDrop_bPtr)(HELE, HDROP, BOOL*);
	typedef int(T_ClassName::*pFun_int_e_e_hDrop_bPtr)(HELE, HELE, HDROP, BOOL*);
};


template<typename T_ClassName>
class CEventHandleWnd : public CEventType < T_ClassName > //处理事件，根据事件类型自动匹配
{
public:
	//HWINDOW
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_w_hDraw_bPtr pMember){
		return (pThis->*pMember)(hWindow, (HDRAW)wParam, pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_uint_uint_uint_bPtr pMember){
		return (pThis->*pMember)(message, wParam, lParam, pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_uint_ptPtr_bPtr pMember){
		return (pThis->*pMember)(wParam, (POINT*)lParam, pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_hDraw_bPtr pMember){
		return (pThis->*pMember)((HDRAW)wParam, pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_bPtr pMember){
		return (pThis->*pMember)(pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_uint_sizePtr_bPtr pMember){
		return (pThis->*pMember)(wParam,(SIZE*)lParam, pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_uint_uint_bPtr pMember){
		return (pThis->*pMember)(wParam, lParam, pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_w_uint_uint_bPtr pMember){
		return (pThis->*pMember)(hWindow, wParam, lParam, pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_uint_bPtr pMember){
		return (pThis->*pMember)(wParam, pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_int_bPtr pMember){
		return (pThis->*pMember)((int)wParam, pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_hMenu_bPtr pMember){
		return (pThis->*pMember)((HMENUX)wParam, pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_hMenu_pMenuPopupWnd_bPtr pMember){
		return (pThis->*pMember)((HMENUX)wParam, (menu_popupWnd_i*)lParam, pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_hDraw_pMenuDrawItem_bPtr pMember){
		return (pThis->*pMember)((HDRAW)wParam, (menu_drawItem_i*)lParam, pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_w_hDraw_pMenuDrawItem_bPtr pMember){
		return (pThis->*pMember)(hWindow,(HDRAW)wParam, (menu_drawItem_i*)lParam, pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_w_uint_ptPtr_bPtr pMember){
		return (pThis->*pMember)(hWindow,(int)wParam, (POINT*)lParam, pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_hDraw_pMenuDrawBackground_bPtr pMember){
		return (pThis->*pMember)((HDRAW)wParam, (menu_drawBackground_i*)lParam, pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_hDrop_bPtr pMember){
		return (pThis->*pMember)((HDROP)wParam, pbHandled);
	}
	BOOL  Handle(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message, T_ClassName* pThis, pFun_int_w_hDrop_bPtr pMember){
		return (pThis->*pMember)(hWindow,(HDROP)wParam, pbHandled);
	}
};

template<typename T_ClassName>
class CEventHandleEle : public CEventType < T_ClassName > //处理事件，根据事件类型自动匹配
{
public:
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_uint_uint_uint_bPtr pMember){
		return (pThis->*pMember)(hEventEle, nEvent, wParam, lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_hDraw_bPtr pMember){
		if(0==nRegType)
			return (pThis->*pMember)((HDRAW)wParam, pbHandled);
		return FALSE;
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_hDraw_bPtr pMember){
		if(0==nRegType)
			return (pThis->*pMember)((HELE)wParam,(HDRAW)lParam, pbHandled);
		if(1==nRegType)
			return (pThis->*pMember)(hEventEle, (HDRAW)wParam, pbHandled);
		return FALSE;
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_e_hDraw_bPtr pMember){
		if(2==nRegType)
			return (pThis->*pMember)(hEle, hEventEle, (HDRAW)wParam, pbHandled);
		return FALSE;
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_bPtr pMember){
		if(0==nRegType)
			return (pThis->*pMember)(pbHandled);
		return FALSE;
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_bPtr pMember){
		if(0==nRegType)
			return (pThis->*pMember)((HELE)wParam, pbHandled);
		if(1==nRegType)
			return (pThis->*pMember)(hEventEle, pbHandled);
		return FALSE;
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_e_bPtr pMember){
		if(0==nRegType)
			return (pThis->*pMember)((HELE)wParam, (HELE)lParam, pbHandled);;
		if(1==nRegType)
			return (pThis->*pMember)(hEventEle, (HELE)wParam, pbHandled);
		if(2==nRegType)
			return (pThis->*pMember)(hEle, hEventEle, pbHandled);
		return FALSE;
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_e_e_bPtr pMember){
		if(1==nRegType)
			return (pThis->*pMember)(hEventEle, (HELE)wParam, (HELE)lParam, pbHandled);
		if(2==nRegType)
			return (pThis->*pMember)(hEle, hEventEle, (HELE)wParam, pbHandled);
		return FALSE;
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_uint_ptPtr_bPtr pMember){
		return (pThis->*pMember)(wParam, (POINT*)lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_uint_ptPtr_bPtr pMember){
		return (pThis->*pMember)(hEventEle, wParam, (POINT*)lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_e_uint_ptPtr_bPtr pMember){
		return (pThis->*pMember)(hEle, hEventEle, wParam, (POINT*)lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_int_int_bPtr pMember){
		return (pThis->*pMember)((int)wParam, (int)lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_int_bPtr pMember){
		return (pThis->*pMember)((int)wParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_w_e_bPtr pMember){
		return (pThis->*pMember)((HWINDOW)wParam, (HELE)lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_w_e_bPtr pMember){
		return (pThis->*pMember)(hEventEle,(HWINDOW)wParam, (HELE)lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_e_w_e_bPtr pMember){
		return (pThis->*pMember)(hEle, hEventEle,(HWINDOW)wParam, (HELE)lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_uint_bPtr pMember){
		return (pThis->*pMember)(wParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_uint_uint_bPtr pMember){
		return (pThis->*pMember)(wParam, lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_uint_uint_bPtr pMember){
		return (pThis->*pMember)(hEventEle, wParam, lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_pListBoxItem_bPtr pMember){
		return (pThis->*pMember)((listBox_item_i*)wParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_pListBoxItem_bPtr pMember){
		return (pThis->*pMember)(hEventEle,(listBox_item_i*)wParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_e_pListBoxItem_bPtr pMember){
		return (pThis->*pMember)(hEle, hEventEle, (listBox_item_i*)wParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_hDraw_pListBoxItem_bPtr pMember){
		return (pThis->*pMember)((HDRAW)wParam,(listBox_item_i*)lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_pListItem_bPtr pMember){
		return (pThis->*pMember)((list_item_i*)wParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_hDraw_pListItem_bPtr pMember){
		return (pThis->*pMember)((HDRAW)wParam,(list_item_i*)lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_hDraw_pListHeaderItem_bPtr pMember){
		return (pThis->*pMember)((HDRAW)wParam,(list_header_item_i*)lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_pTreeItem_bPtr pMember){
		return (pThis->*pMember)((tree_item_i*)wParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_hDraw_pTreeItem_bPtr pMember){
		return (pThis->*pMember)((HDRAW)wParam,(tree_item_i*)lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_pListViewItem_bPtr pMember){
		return (pThis->*pMember)((listView_item_i*)wParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_hDraw_pListViewItem_bPtr pMember){
		return (pThis->*pMember)((HDRAW)wParam,(listView_item_i*)lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_int_bPtr pMember){
		return (pThis->*pMember)(hEventEle, wParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_e_uint_uint_uint_bPtr pMember){
		return (pThis->*pMember)(hEle, hEventEle, nEvent, wParam, lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_e_int_bPtr pMember){
		return (pThis->*pMember)(hEle, hEventEle, wParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_hDraw_pMenuDrawBackground_bPtr pMember){
		return (pThis->*pMember)((HDRAW)wParam,(menu_drawBackground_i*)lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_hDraw_pMenuDrawItem_bPtr pMember){
		return (pThis->*pMember)((HDRAW)wParam,(menu_drawItem_i*)lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_hMenu_bPtr pMember){
		return (pThis->*pMember)((HMENUX)wParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_hMenu_pMenuPopupWnd_bPtr pMember){
		return (pThis->*pMember)((HMENUX)wParam, (menu_popupWnd_i*)lParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_hDrop_bPtr pMember){
		return (pThis->*pMember)((HDROP)wParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_hDrop_bPtr pMember){
		return (pThis->*pMember)(hEventEle,(HDROP)wParam, pbHandled);
	}
	BOOL  Handle(int nRegType, HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent, T_ClassName* pThis, pFun_int_e_e_hDrop_bPtr pMember){
		return (pThis->*pMember)(hEle,hEventEle,(HDROP)wParam, pbHandled);
	}
};

//省略1参数验证
template<typename T_ClassName>
class CEventWndTest : public  CEventType < T_ClassName >
{
public:
	void  Test(UINT nEvent, pFun_int_uint_uint_uint_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_uint;
		pArray[1]=XC_PARAM_uint;
		pArray[2]=XC_PARAM_uint;
	}
	void  Test(UINT nEvent, pFun_int_uint_ptPtr_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_uint;
		pArray[1]=XC_PARAM_P_POINT;
	}
	void  Test(UINT nEvent, pFun_int_hDraw_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDRAW;
	}
	void  Test(UINT nEvent, pFun_int_bPtr,BYTE* pArray){
	}
	void  Test(UINT nEvent, pFun_int_uint_sizePtr_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_uint;
		pArray[1]=XC_PARAM_P_SIZE;
	}
	void  Test(UINT nEvent, pFun_int_uint_uint_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_uint;
		pArray[1]=XC_PARAM_uint;
	}
	void  Test(UINT nEvent, pFun_int_uint_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_uint;
	}
	void  Test(UINT nEvent, pFun_int_int_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_int;
	}
	void  Test(UINT nEvent, pFun_int_hMenu_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HMENUX;
	}
	void  Test(UINT nEvent, pFun_int_hMenu_pMenuPopupWnd_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HMENUX;
		pArray[1]=XC_PARAM_P_menu_popupWnd;
	}
	void Test(UINT nEvent,pFun_int_hDraw_pMenuDrawBackground_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDRAW;
		pArray[1]=XC_PARAM_P_menu_drawBackground;
	};
	void Test(UINT nEvent,pFun_int_hDraw_pMenuDrawItem_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDRAW;
		pArray[1]=XC_PARAM_P_menu_drawItem;
	}
	void Test(UINT nEvent,pFun_int_hDrop_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDROP;
	}
};

template<typename T_ClassName>
class CEventWndTest2 : public  CEventType < T_ClassName >
{
public:
	void  Test(UINT nEvent, pFun_int_w_hDraw_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDRAW;
	}
	void  Test(UINT nEvent, pFun_int_w_uint_ptPtr_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_uint;
		pArray[1]=XC_PARAM_P_POINT;
	}
	void Test(UINT nEvent,pFun_int_w_hDrop_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDROP;
	}
	void Test(UINT nEvent,pFun_int_w_uint_uint_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_uint;
		pArray[1]=XC_PARAM_uint;
	}
};

//省略2参数验证
template<typename T_ClassName>
class CEventEleTest : public  CEventType < T_ClassName >
{
public:
	void  Test(UINT nEvent, pFun_int_hDraw_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDRAW;
	}
	void  Test(UINT nEvent, pFun_int_bPtr,BYTE* pArray){
	}
	void  Test(UINT nEvent, pFun_int_e_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HELE;
	}
	void  Test(UINT nEvent, pFun_int_uint_ptPtr_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_uint;
		pArray[1]=XC_PARAM_P_POINT;
	}
	void  Test(UINT nEvent, pFun_int_int_int_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_int;
		pArray[1]=XC_PARAM_int;
	}
	void  Test(UINT nEvent, pFun_int_int_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_int;
	}
	void  Test(UINT nEvent, pFun_int_uint_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_uint;
	}
	void  Test(UINT nEvent, pFun_int_uint_uint_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_uint;
		pArray[1]=XC_PARAM_uint;
	}
	void  Test(UINT nEvent, pFun_int_w_e_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HWINDOW;
		pArray[1]=XC_PARAM_HELE;
	}
	void  Test(UINT nEvent, pFun_int_pListBoxItem_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_P_listBox_item_i;
	}
	void  Test(UINT nEvent, pFun_int_hDraw_pListBoxItem_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDRAW;
		pArray[1]=XC_PARAM_P_listBox_item_i;
	}
	void  Test(UINT nEvent, pFun_int_pListItem_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_P_list_item_i;
	}
	void  Test(UINT nEvent, pFun_int_hDraw_pListItem_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDRAW;
		pArray[1]=XC_PARAM_P_list_item_i;
	}
	void  Test(UINT nEvent, pFun_int_hDraw_pListHeaderItem_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDRAW;
		pArray[1]=XC_PARAM_P_list_header_item_i;
	}
	void  Test(UINT nEvent,pFun_int_pTreeItem_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_P_tree_item_i;
	}
	void  Test(UINT nEvent,pFun_int_hDraw_pTreeItem_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDRAW;
		pArray[1]=XC_PARAM_P_tree_item_i;
	}
	void  Test(UINT nEvent,pFun_int_pListViewItem_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_P_listView_item_i;
	}
	void  Test(UINT nEvent,pFun_int_hDraw_pListViewItem_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDRAW;
		pArray[1]=XC_PARAM_P_listView_item_i;
	}
	void  Test(UINT nEvent,pFun_int_hDraw_pMenuDrawBackground_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDRAW;
		pArray[1]=XC_PARAM_P_menu_drawBackground;
	}
	void  Test(UINT nEvent,pFun_int_hDraw_pMenuDrawItem_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDRAW;
		pArray[1]=XC_PARAM_P_menu_drawItem;
	}
	void  Test(UINT nEvent,pFun_int_hMenu_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HMENUX;
	}
	void  Test(UINT nEvent,pFun_int_hMenu_pMenuPopupWnd_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HMENUX;
		pArray[1]=XC_PARAM_P_menu_popupWnd;
	}
	void  Test(UINT nEvent,pFun_int_hDrop_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDROP;
	}
};

//省略1参数验证
template<typename T_ClassName>
class CEventEleTest1 : public  CEventType < T_ClassName >
{
public:
	void  Test(UINT nEvent, pFun_int_e_uint_uint_uint_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_uint;
		pArray[1]=XC_PARAM_uint;
		pArray[2]=XC_PARAM_uint;
	}
	void  Test(UINT nEvent, pFun_int_e_int_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_int;
	}
	void  Test(UINT nEvent, pFun_int_e_hDraw_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDRAW;
	}
	void  Test(UINT nEvent, pFun_int_e_bPtr,BYTE* pArray){
	}
	void  Test(UINT nEvent, pFun_int_e_e_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HELE;
		pArray[1]=XC_PARAM_HELE;
	}
	void  Test(UINT nEvent, pFun_int_e_w_e_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HWINDOW;
		pArray[1]=XC_PARAM_HELE;
	}
	void  Test(UINT nEvent, pFun_int_e_uint_ptPtr_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_uint;
		pArray[1]=XC_PARAM_P_POINT;
	}
	void  Test(UINT nEvent, pFun_int_e_uint_uint_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_uint;
		pArray[1]=XC_PARAM_uint;
	}
	void  Test(UINT nEvent, pFun_int_e_pListBoxItem_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_P_listBox_item_i;
	}
	void  Test(UINT nEvent,pFun_int_e_hDrop_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDROP;
	}
};

//全参数验证
template<typename T_ClassName>
class CEventEleTest2 : public  CEventType < T_ClassName >
{
public:
	void  Test(UINT nEvent, pFun_int_e_e_uint_uint_uint_bPtr, BYTE* pArray){
		pArray[0]=XC_PARAM_uint;
		pArray[1]=XC_PARAM_uint;
		pArray[2]=XC_PARAM_uint;
	}
	void  Test(UINT nEvent, pFun_int_e_e_hDraw_bPtr, BYTE* pArray){
		pArray[0]=XC_PARAM_HDRAW;
	}
	void  Test(UINT nEvent, pFun_int_e_e_int_bPtr, BYTE* pArray){
		pArray[0]=XC_PARAM_int;
	}
	void  Test(UINT nEvent, pFun_int_e_e_bPtr, BYTE* pArray){
	}
	void  Test(UINT nEvent, pFun_int_e_e_e_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HELE;
		pArray[1]=XC_PARAM_HELE;
		pArray[2]=XC_PARAM_HELE;
	}
	void  Test(UINT nEvent, pFun_int_e_e_w_e_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HWINDOW;
		pArray[1]=XC_PARAM_HELE;
	}
	void  Test(UINT nEvent, pFun_int_e_e_uint_ptPtr_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_uint;
		pArray[1]=XC_PARAM_P_POINT;
	}
	void  Test(UINT nEvent, pFun_int_e_e_pListBoxItem_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_P_listBox_item_i;
	}
	void  Test(UINT nEvent,pFun_int_e_e_hDrop_bPtr,BYTE* pArray){
		pArray[0]=XC_PARAM_HDROP;
	}
};

template<typename T_ClassName, typename T_Member>
class CEventWnd : public  CEventBase
{
public:
	CEventWnd(){
		m_nEvent = 0; m_pThis = NULL; m_pMember = NULL; m_name[0]=NULL;
	}
	BOOL  Bind(char* pName,int nEvent, T_ClassName*  pThis, T_Member pMember){
		CEventWndTest<T_ClassName> EventTest;  //事件格式验证
		BYTE array[4]={0};
		EventTest.Test(nEvent, pMember,array);
		if(FALSE==XWnd_RegEventTest(nEvent,array))  return FALSE;

		m_nEvent = nEvent;  m_pThis = pThis;  m_pMember = pMember;
		m_regType=0;
		strcpy_s(m_name,pName);
		return TRUE;
	}
	BOOL  Bind2(char* pName,int nEvent, T_ClassName*  pThis, T_Member pMember){
		CEventWndTest2<T_ClassName> EventTest;  //事件格式验证
		BYTE array[4]={0};
		EventTest.Test(nEvent, pMember,array);
		if(FALSE==XWnd_RegEventTest(nEvent,array))  return FALSE;

		m_nEvent = nEvent;  m_pThis = pThis;  m_pMember = pMember;
		m_regType=1;
		strcpy_s(m_name,pName);
		return TRUE;
	}
	BOOL Remove(char* pName,HWINDOW hWindow, int nEvent, T_ClassName*  pThis, T_Member pMember){
		m_nEvent = nEvent;  m_pThis = pThis;  m_pMember = pMember;
		strcpy_s(m_name,pName);
		return XWnd_RemoveEventEx(hWindow,this);
	}
private:
	T_ClassName*   m_pThis;
	T_Member  m_pMember;
	char      m_name[MAX_PATH];
private:
	CEventHandleWnd<T_ClassName>   m_EventHandle;
	virtual void Release(){ delete this; }
	virtual int OnWndProc(HWINDOW hWindow, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT message){
		return m_EventHandle.Handle(hWindow, wParam, lParam, pbHandled, message, m_pThis, m_pMember);
	}
	virtual  void* GetThis(){ return m_pThis; }
	virtual  char* GetFunName(){ return m_name; }
};

template<typename T_ClassName, typename T_Member>
class CEventEle : public  CEventBase
{
public:
	CEventEle(){
		m_nEvent = 0; m_pThis = NULL; m_pMember = NULL;
		m_name[0]=0;
	}
	BOOL  Bind(char* pName,int nEvent, T_ClassName*  pThis, T_Member pMember){
		CEventEleTest<T_ClassName> EventTest;  //事件格式验证
		BYTE array[4]={0};
		EventTest.Test(nEvent, pMember,array);
		if(FALSE==XEle_RegEventTest(nEvent,array))  return FALSE;

		m_nEvent = nEvent;  m_pThis = pThis;  m_pMember = pMember;
		m_regType=0;
		strcpy_s(m_name,pName);
		return TRUE;
	}
	BOOL  Bind2(char* pName,int nEvent, T_ClassName*  pThis, T_Member pMember){
		CEventEleTest1<T_ClassName> EventTest;  //事件格式验证
		BYTE array[4]={0};
		EventTest.Test(nEvent, pMember,array);
		if(FALSE==XEle_RegEventTest(nEvent,array))  return FALSE;

		m_nEvent = nEvent;  m_pThis = pThis;  m_pMember = pMember;
		m_regType=1;
		strcpy_s(m_name,pName);
		return TRUE;
	}
	BOOL  Bind3(char* pName,int nEvent, T_ClassName*  pThis, T_Member pMember){
		CEventEleTest2<T_ClassName> EventTest;  //事件格式验证
		BYTE array[4]={0};
		EventTest.Test(nEvent, pMember,array);
		if(FALSE==XEle_RegEventTest(nEvent,array))  return FALSE;

		m_nEvent = nEvent;  m_pThis = pThis;  m_pMember = pMember;
		m_regType=2;
		strcpy_s(m_name,pName);
		return TRUE;
	}
	BOOL Remove(char* pName,HELE hEle, int nEvent, T_ClassName*  pThis, T_Member pMember)
	{
		m_nEvent = nEvent;  m_pThis = pThis;  m_pMember = pMember;
		strcpy_s(m_name,pName);
		return XEle_RemoveEventEx(hEle,this);
	}
private:
	T_ClassName*   m_pThis;
	T_Member  m_pMember;
	char  m_name[MAX_PATH];

	CEventHandleEle<T_ClassName>   m_EventHandle;
	virtual void Release(){ delete this; }
	virtual int OnEleProc(HELE hEle, HELE hEventEle, WPARAM wParam, LPARAM lParam, BOOL *pbHandled, UINT nEvent){
		return m_EventHandle.Handle(m_regType,hEle, hEventEle, wParam, lParam, pbHandled, nEvent, m_pThis, m_pMember);
	}
	virtual  void* GetThis(){ return m_pThis; }
	virtual  char* GetFunName(){ return m_name; }
};

/////////////窗口//////////////////////////////////////////////////
template<typename T_ClassName, typename T_Member>
void _XWnd_RegEventCPP1(HWINDOW hWindow, int nEvent, T_ClassName *pThis, T_Member pFun, char* pFunName, char* pFile, int pLine)
{
	CEventWnd< T_ClassName, T_Member > *pEvent = new CEventWnd < T_ClassName, T_Member > ;
	if (pEvent->Bind2(pFunName,nEvent, pThis, pFun)){
		XWnd_RegEventEx(hWindow, pEvent);  return;
	}	delete pEvent;
}

// 注册窗口事件,将类成员函数作为事件回调函数.回调函数不省略参数.
#define XWnd_RegEventCPP1(hWindow, nEvent,pFunMember)  _XWnd_RegEventCPP1(hWindow, nEvent, this, pFunMember,#pFunMember,__FILE__,__LINE__);

template<typename T_ClassName, typename T_Member>
void _XWnd_RegEventCPP(HWINDOW hWindow, int nEvent, T_ClassName *pThis, T_Member pFun, char* pFunName, char* pFile, int pLine)
{
	CEventWnd< T_ClassName, T_Member > *pEvent = new CEventWnd < T_ClassName, T_Member > ;
	if (pEvent->Bind(pFunName,nEvent, pThis, pFun)){
		XWnd_RegEventEx(hWindow, pEvent);  return;
	}	delete pEvent;
}

// 注册窗口事件,将类成员函数作为事件回调函数.回调函数省略参数窗口自身句柄hWindow.
#define XWnd_RegEventCPP(hWindow, nEvent,pFunMember)  _XWnd_RegEventCPP(hWindow, nEvent, this, pFunMember,#pFunMember,__FILE__,__LINE__);

template<typename T_ClassName, typename T_Member>
BOOL _Wnd_RemoveEventCPP(HWINDOW hWindow, int nEvent, T_ClassName *pThis, T_Member pFun,char* pFunName)
{
	CEventWnd< T_ClassName, T_Member > Event;
	return Event.Remove(pFunName,hWindow,nEvent, pThis, pFun);
}
//移除事件C++
#define XWnd_RemoveEventCPP(hWindow,nEvent,pFunMember)  _Wnd_RemoveEventCPP(hWindow,nEvent,this,pFunMember,#pFunMember)


////////////////元素///////////////////////////////////////////////////
template<typename T_ClassName, typename T_Member>
void _XEle_RegEventCPP2(HELE hEle, int nEvent, T_ClassName *pThis, T_Member pFun,char* pFunName)
{
	CEventEle< T_ClassName, T_Member > *pEvent = new CEventEle < T_ClassName, T_Member > ;
	if (pEvent->Bind3(pFunName,nEvent, pThis, pFun)){
		XEle_RegEventEx(hEle, pEvent);  return;
	}  delete pEvent;
}

// 注册元素事件,将类成员函数作为事件回调函数.回调函数不省略参数
#define XEle_RegEventCPP2(hEle, nEvent,pFunMember)  _XEle_RegEventCPP2(hEle, nEvent, this, pFunMember,#pFunMember);

template<typename T_ClassName, typename T_Member>
void _XEle_RegEventCPP1(HELE hEle, int nEvent, T_ClassName *pThis, T_Member pFun,char* pFunName)
{
	CEventEle< T_ClassName, T_Member > *pEvent = new CEventEle < T_ClassName, T_Member > ;
	if (pEvent->Bind2(pFunName,nEvent, pThis, pFun)){
		XEle_RegEventEx(hEle, pEvent);  return;
	}  delete pEvent;
}

// 注册元素事件,将类成员函数作为事件回调函数.回调函数省略元素自身句柄hEle,
#define XEle_RegEventCPP1(hEle, nEvent,pFunMember)  _XEle_RegEventCPP1(hEle, nEvent, this, pFunMember,#pFunMember);

template<typename T_ClassName, typename T_Member>
void _XEle_RegEventCPP(HELE hEle, int nEvent, T_ClassName *pThis, T_Member pFun,char* pFunName)
{
	CEventEle< T_ClassName, T_Member > *pEvent = new CEventEle < T_ClassName, T_Member > ;
	if (pEvent->Bind(pFunName,nEvent, pThis, pFun)){

		XEle_RegEventEx(hEle, pEvent);  return;
	}  delete pEvent;
}


// 注册元素事件,将类成员函数作为事件回调函数.回调函数省略元素自身句柄hEle,省略触发事件元素句柄hEventEle.
#define XEle_RegEventCPP(hEle, nEvent,pFunMember)  _XEle_RegEventCPP(hEle, nEvent, this, pFunMember, #pFunMember);

template<typename T_ClassName, typename T_Member>
BOOL _XEle_RemoveEventCPP(HELE hEle, int nEvent, T_ClassName *pThis, T_Member pFun,char* pFunName)
{
	CEventEle< T_ClassName, T_Member > Event;
	return Event.Remove(pFunName,hEle,nEvent, pThis, pFun);
}

//移除事件C++
#define XEle_RemoveEventCPP(hEle,nEvent,pFunMember)  _XEle_RemoveEventCPP(hEle,nEvent,this,pFunMember,#pFunMember)

//////////////C方式注册///////////////////////////////////////////////
//将函数注册到指定消息,不省略参数
template<typename T_Fun>
void XWnd_RegEventCT1(HWINDOW hWindow, int nEvent, T_Fun *pFun)
{
	XWnd_RegEventC1(hWindow, nEvent, pFun);
}

//将函数注册到指定消息,类成员函数省略窗口自身句柄hWindow
template<typename T_Fun>
void XWnd_RegEventCT(HWINDOW hWindow, int nEvent, T_Fun *pFun)
{
	XWnd_RegEventC(hWindow, nEvent, pFun);
}

//将函数注册到指定事件,不省略参数
template<typename T_Fun>
void XEle_RegEventCT2(HELE hEle, int nEvent, T_Fun *pFun)
{
	XEle_RegEventC2(hEle, nEvent, pFun);
}

//将函数注册到指定事件,类成员函数省略元素自身句柄hEle
template<typename T_Fun>
void XEle_RegEventCT1(HELE hEle, int nEvent, T_Fun *pFun)
{
	XEle_RegEventC1(hEle, nEvent, pFun);
}

//将函数注册到指定事件,类成员函数省略元素自身句柄hEle,省略触发事件元素句柄hEventEle
template<typename T_Fun>
void XEle_RegEventCT(HELE hEle, int nEvent, T_Fun *pFun)
{
	XEle_RegEventC(hEle, nEvent, pFun);
}

//hEle      接收事件的元素句柄
//hEventEle 触发事件的元素句柄
//////////////////////////////////////////////////////////////////////


XC_API int WINAPI XC_UnicodeToAnsi(const wchar_t *pIn, int inLen, char *pOut, int outLen);
XC_API int WINAPI XC_AnsiToUnicode(const char *pIn, int inLen, wchar_t *pOut, int outLen);
XC_API void WINAPI XC_DebugToFileInfo(const char *pInfo);//打印调试信息到文件 xcgui_debug.txt
XC_API BOOL WINAPI XC_IsHELE(HXCGUI hEle); //检查句柄
XC_API BOOL WINAPI XC_IsHWINDOW(HXCGUI hWindow); //检查句柄
XC_API BOOL WINAPI XC_IsShape(HXCGUI hShape); //检查句柄
XC_API BOOL WINAPI XC_IsHXCGUI(HXCGUI hXCGUI,XC_OBJECT_TYPE nType);
XC_API HWINDOW WINAPI XC_hWindowFromHWnd(HWND hWnd);
XC_API BOOL WINAPI XC_IsSViewExtend(HELE hEle);  //判断元素是否从滚动视图元素扩展的新元素,包含滚动视图元素
XC_API XC_OBJECT_TYPE WINAPI XC_GetObjectType(HXCGUI hXCGUI);
XC_API HXCGUI WINAPI XC_GetObjectByID(int nID); //通过ID获取对象句柄
XC_API int  WINAPI XC_GetResIDValue(const wchar_t *pName);  //获取资源ID值
XC_API BOOL WINAPI XC_RectInRect(RECT *pRect1, RECT *pRect2);
XC_API void WINAPI XC_CombineRect(RECT *pDest, RECT *pSrc1, RECT *pSrc2);
XC_API void WINAPI XC_ShowLayoutFrame(BOOL bShow);
XC_API void WINAPI XC_EnableDebugFile(BOOL bEnable);
XC_API void WINAPI XC_SetLayoutFrameColor(COLORREF color);
XC_API void WINAPI XC_EnableErrorMessageBox(BOOL bEnabel);
XC_API HXCGUI WINAPI XC_LoadLayout(const wchar_t *pFileName, HXCGUI hParent=NULL);
XC_API HXCGUI WINAPI XC_LoadLayoutFromString(const char *pStringXML, HXCGUI hParent=NULL);
XC_API BOOL WINAPI XC_LoadResource(const wchar_t *pFileName,const wchar_t *pDir=NULL);
XC_API BOOL WINAPI XC_LoadResourceFromString(const char *pStringXML,const wchar_t *pDir=NULL);
XC_API template_info_i* WINAPI XC_LoadTemplate(XC_OBJECT_TYPE nType,const wchar_t* pFileName);  //加载模板 返回模板对象
XC_API template_info_i* WINAPI XC_LoadTemplateFromString(XC_OBJECT_TYPE nType,const char* pStringXML);  //加载模板,从内存, 返回模板对象
XC_API void WINAPI XC_TemplateDestroy(template_info_i* pInfo);
XC_API void WINAPI XC_GetTextSize(const wchar_t* pString,int length,HFONTX hFontX,out_ SIZE* pOutSize);
XC_API HFONTX WINAPI XC_GetDefaultFont(); //获取默认字体
XC_API void  WINAPI XC_InitFont(LOGFONTW *pFont,wchar_t *pName,int size,BOOL bBold=FALSE,BOOL bItalic=FALSE,BOOL bUnderline=FALSE,BOOL bStrikeOut=FALSE);
XC_API  void* WINAPI XC_Malloc(int size);
XC_API  void  WINAPI XC_Free(void *p);
XC_API void WINAPI _XC_SetType(HXCGUI hXCGUI,XC_OBJECT_TYPE nType);
XC_API void WINAPI _XC_AddType(HXCGUI hXCGUI,XC_OBJECT_TYPE nType);
XC_API void WINAPI _XC_BindData(HXCGUI hXCGUI,int data);
XC_API int  WINAPI _XC_GetBindData(HXCGUI hXCGUI);
XC_API BOOL WINAPI XInitXCGUI(wchar_t *pText = NULL);
XC_API void WINAPI XRunXCGUI();
XC_API void WINAPI XExitXCGUI();
XC_API HBKINFOM WINAPI XBkInfoM_Create();   //创建背景内容管理器
XC_API void WINAPI XBkInfoM_Destroy(HBKINFOM hBkInfoM);  //销毁背景内容管理器
XC_API int  WINAPI XBkInfoM_SetBkInfo(HBKINFOM hBkInfoM, const wchar_t *pText);  //添加背景内容
XC_API void WINAPI XBkInfoM_AddBorder(HBKINFOM hBkInfoM, COLORREF color,BYTE alpha, int width);  //添加背景内容 - 边框
XC_API void WINAPI XBkInfoM_AddFill(HBKINFOM hBkInfoM, COLORREF color,BYTE alpha);  //添加背景内容 - 填充
XC_API void WINAPI XBkInfoM_AddImage(HBKINFOM hBkInfoM, HIMAGE hImage); //添加背景内天 - 图片
XC_API int  WINAPI XBkInfoM_GetCount(HBKINFOM hBkInfoM); //获取背景内容数量
XC_API void WINAPI XBkInfoM_Clear(HBKINFOM hBkInfoM);  //清空背景内容
XC_API BOOL WINAPI XBkInfoM_Draw(HBKINFOM hBkInfoM, HDRAW hDraw, RECT* pRect);  //绘制背景内容
XC_API HELE WINAPI XBtn_Create(int x, int y, int cx, int cy,const wchar_t *pName, HXCGUI hParent = NULL);
XC_API BOOL WINAPI XBtn_IsCheck(HELE hEle); //按钮是否被选中
XC_API BOOL WINAPI XBtn_SetCheck(HELE hEle, BOOL bCheck);
XC_API void WINAPI XBtn_SetState(HELE hEle,common_state3_ nState);
XC_API common_state3_  WINAPI XBtn_GetState(HELE hEle);
XC_API button_state_   WINAPI XBtn_GetStateEx(HELE hEle);
XC_API void WINAPI XBtn_SetStyle(HELE hEle, button_style_  nStyle);
XC_API button_style_  WINAPI XBtn_GetStyle(HELE hEle);
XC_API void WINAPI XBtn_SetType(HELE hEle,button_type_ nType);
XC_API button_type_  WINAPI XBtn_GetType(HELE hEle);
XC_API void WINAPI XBtn_SetGroupID(HELE hEle,int nID);
XC_API int  WINAPI XBtn_GetGroupID(HELE hEle);
XC_API void WINAPI XBtn_SetBindEle(HELE hEle,HELE hBindEle);
XC_API HELE WINAPI XBtn_GetBindEle(HELE hEle);
XC_API void WINAPI XBtn_SetTextAlign(HELE hEle,int nFlags);
XC_API int  WINAPI XBtn_GetTextAlign(HELE hEle);
XC_API void WINAPI XBtn_SetIconAlign(HELE hEle,button_icon_align_ align);
XC_API void WINAPI XBtn_SetOffset(HELE hEle,int x,int y);
XC_API void WINAPI XBtn_SetOffsetIcon(HELE hEle,int x,int y);
XC_API void WINAPI XBtn_SetIconSpace(HELE hEle, int size);
XC_API void WINAPI XBtn_GetText(HELE hEle,out_ wchar_t *pOut,int nOutLen);
XC_API void WINAPI XBtn_SetText(HELE hEle,in_ const wchar_t *pName);
XC_API void WINAPI XBtn_SetIcon(HELE hEle,HIMAGE hImage);
XC_API void WINAPI XBtn_SetIconDisable(HELE hEle,HIMAGE hImage);
XC_API void WINAPI XBtn_AddAnimationFrame(HELE hEle,HIMAGE hImage,UINT uElapse);
XC_API void WINAPI XBtn_EnableAnimation(HELE hEle,BOOL bEnable,BOOL bLoopPlay=FALSE);
XC_API void WINAPI XBtn_AddBkBorder(HELE hEle, button_state_ nState,COLORREF color,BYTE alpha, int width);
XC_API void WINAPI XBtn_AddBkFill(HELE hEle, button_state_ nState,COLORREF color,BYTE alpha);
XC_API void WINAPI XBtn_AddBkImage(HELE hEle, button_state_ nState,HIMAGE hImage);
XC_API int  WINAPI XBtn_GetBkInfoCount(HELE hEle, button_state_ nState);
XC_API void WINAPI XBtn_ClearBkInfo(HELE hEle, button_state_ nState);
XC_API HBKINFOM WINAPI XBtn_GetBkInfoManager(HELE hEle,button_state_ nState);
XC_API HELE WINAPI XComboBox_Create(int x, int y, int cx, int cy, HXCGUI hParent = NULL);
XC_API BOOL  WINAPI XComboBox_SetSelItem(HELE hEle,int iIndex);
XC_API void WINAPI XComboBox_GetButtonRect(HELE hEle,out_ RECT* pRect);
XC_API void WINAPI XComboBox_SetButtonSize(HELE hEle,int size);
XC_API void WINAPI XComboBox_SetDropHeight(HELE hEle,int height); //设置下拉列表高度
XC_API int  WINAPI XComboBox_GetDropHeight(HELE hEle);
XC_API void WINAPI XComboBox_BindApapter(HELE hEle,HXCGUI hAdapter);
XC_API void WINAPI XComboBox_SetItemTemplateXML(HELE hEle,const wchar_t* pXmlFile);
XC_API void WINAPI XComboBox_SetItemTemplateXMLFromString(HELE hEle,const char* pStringXML);
XC_API void WINAPI XComboBox_EnableDrawButton(HELE hEle,BOOL bEnable);
XC_API void WINAPI XComboBox_EnableEdit(HELE hEle,BOOL bEdit);  //启动编辑内容
XC_API void WINAPI XComboBox_AddBkBorder(HELE hEle,comboBox_state_ nState,COLORREF color,BYTE alpha, int width);
XC_API void WINAPI XComboBox_AddBkFill(HELE hEle,comboBox_state_ nState,COLORREF color,BYTE alpha);
XC_API void WINAPI XComboBox_AddBkImage(HELE hEle,comboBox_state_ nState,HIMAGE hImage);
XC_API int  WINAPI XComboboX_GetBkInfoCount(HELE hEle, comboBox_state_ nState);
XC_API void WINAPI XComboBox_ClearBkInfo(HELE hEle,comboBox_state_ nState);
XC_API HBKINFOM WINAPI XComboBox_GetBkInfoManager(HELE hEle, comboBox_state_ nState);
XC_API int WINAPI XComboBox_GetSelItem(HELE hEle);
XC_API comboBox_state_ WINAPI XComboBox_GetState(HELE hEle);
XC_API void WINAPI XAdapter_Destroy(HXCGUI hAdapter);
XC_API HXCGUI WINAPI XAdapterListView_Create();
XC_API int  WINAPI XAdapterListView_Group_AddColumn(HXCGUI hAdapter,const wchar_t *pName);
XC_API int  WINAPI XAdapterListView_Group_AddItemText(HXCGUI hAdapter,const wchar_t *pValue);
XC_API int  WINAPI XAdapterListView_Group_AddItemTextEx(HXCGUI hAdapter,const wchar_t* pName,const wchar_t *pValue);
XC_API int  WINAPI XAdapterListView_Group_AddItemImage(HXCGUI hAdapter,HIMAGE hImage);
XC_API int  WINAPI XAdapterListView_Group_AddItemImageEx(HXCGUI hAdapter,const wchar_t* pName,HIMAGE hImage);
XC_API BOOL WINAPI XAdapterListView_Group_SetText(HXCGUI hAdapter,int iGroup,int iColumn,const wchar_t *pValue);
XC_API BOOL WINAPI XAdapterListView_Group_SetTextEx(HXCGUI hAdapter,int iGroup,const wchar_t *pName,const wchar_t *pValue);
XC_API BOOL WINAPI XAdapterListView_Group_SetImage(HXCGUI hAdapter,int iGroup,int iColumn,HIMAGE hImage);
XC_API BOOL WINAPI XAdapterListView_Group_SetImageEx(HXCGUI hAdapter,int iGroup,const wchar_t *pName,HIMAGE hImage);
XC_API int  WINAPI XAdapterListView_Item_AddColumn(HXCGUI hAdapter,const wchar_t *pName);  //增加列
XC_API int  WINAPI XAdapterListView_Item_AddItemText(HXCGUI hAdapter,int iGroup,const wchar_t *pValue);
XC_API int  WINAPI XAdapterListView_Item_AddItemTextEx(HXCGUI hAdapter,int iGroup,const wchar_t *pName,const wchar_t *pValue);
XC_API int  WINAPI XAdapterListView_Item_AddItemImage(HXCGUI hAdapter,int iGroup,HIMAGE hImage);
XC_API int  WINAPI XAdapterListView_Item_AddItemImageEx(HXCGUI hAdapter,int iGroup,const wchar_t *pName,HIMAGE hImage);
XC_API BOOL WINAPI XAdapterListView_Item_SetText(HXCGUI hAdapter,int iGroup,int iItem,int iColumn,const wchar_t* pValue);
XC_API BOOL WINAPI XAdapterListView_Item_SetTextEx(HXCGUI hAdapter,int iGroup,int iItem,const wchar_t *pName,const wchar_t* pValue);
XC_API BOOL WINAPI XAdapterListView_Item_SetImage(HXCGUI hAdapter,int iGroup,int iItem,int iColumn,HIMAGE hImage);
XC_API BOOL WINAPI XAdapterListView_Item_SetImageEx(HXCGUI hAdapter,int iGroup,int iItem,const wchar_t *pName,HIMAGE hImage);
XC_API BOOL WINAPI XAdapterListView_Group_DeleteItem(HXCGUI hAdapter,int iGroup); //删除组
XC_API BOOL WINAPI XAdapterListView_Item_DeleteItem(HXCGUI hAdapter,int iGroup,int iItem);  //删除项
XC_API void WINAPI XAdapterListView_DeleteAll(HXCGUI hAdapter);
XC_API BOOL WINAPI XAdapterListView_Item_GetTextEx(HXCGUI hAdapter,int iGroup,int iItem,const wchar_t* pName,out_ wchar_t* pOut,int nOutLen);
XC_API HIMAGE WINAPI XAdapterListView_Item_GetImageEx(HXCGUI hAdapter,int iGroup,int iItem,const wchar_t* pName);
XC_API HXCGUI WINAPI XAdapterTable_Create();
XC_API int WINAPI XAdapterTable_AddColumn(HXCGUI hAdapter,const wchar_t* pName);
XC_API int WINAPI XAdapterTable_SetColumn(HXCGUI hAdapter, const wchar_t*  pColName);
XC_API int WINAPI XAdapterTable_AddItemText(HXCGUI hAdapter,const wchar_t* pValue);
XC_API int WINAPI XAdapterTable_AddItemTextEx(HXCGUI hAdapter,const wchar_t* pName,const wchar_t* pValue);
XC_API int WINAPI XAdapterTable_AddItemImage(HXCGUI hAdapter,HIMAGE hImage);
XC_API int WINAPI XAdapterTable_AddItemImageEx(HXCGUI hAdapter,const wchar_t* pName,HIMAGE hImage);
XC_API BOOL WINAPI XAdapterTable_InsertItemText(HXCGUI hAdapter,int iItem,const wchar_t* pValue);
XC_API BOOL WINAPI XAdapterTable_InsertItemTextEx(HXCGUI hAdapter,int iItem,const wchar_t* pName,const wchar_t* pValue);
XC_API BOOL WINAPI XAdapterTable_InsertItemImage(HXCGUI hAdapter,int iItem,HIMAGE hImage);
XC_API BOOL WINAPI XAdapterTable_InsertItemImageEx(HXCGUI hAdapter,int iItem,const wchar_t* pName,HIMAGE hImage);
XC_API BOOL WINAPI XAdapterTable_SetItemText(HXCGUI hAdapter,int iItem,int iColumn,const wchar_t* pValue);
XC_API BOOL WINAPI XAdapterTable_SetItemTextEx(HXCGUI hAdapter,int iItem,const wchar_t* pName,const wchar_t* pValue);
XC_API BOOL WINAPI XAdapterTable_SetItemImage(HXCGUI hAdapter,int iItem,int iColumn,HIMAGE hImage);
XC_API BOOL WINAPI XAdapterTable_SetItemImageEx(HXCGUI hAdapter,int iItem,const wchar_t* pName,HIMAGE hImage);
XC_API BOOL WINAPI XAdapterTable_DeleteItem(HXCGUI hAdapter,int iItem);
XC_API BOOL WINAPI XAdapterTable_DeleteItemEx(HXCGUI hAdapter,int iItem,int nCount);
XC_API void WINAPI XAdapterTable_DeleteItemAll(HXCGUI hAdapter);
XC_API void WINAPI XAdapterTable_DeleteColumnAll(HXCGUI hAdapter);
XC_API int WINAPI XAdapterTable_GetCount(HXCGUI hAdapter);
XC_API int WINAPI XAdapterTable_GetCountColumn(HXCGUI hAdapter);
XC_API BOOL WINAPI XAdapterTable_GetItemText(HXCGUI hAdapter,int iItem,int iColumn,out_ wchar_t* pOut,int nOutLen);
XC_API HIMAGE WINAPI XAdapterTable_GetItemImage(HXCGUI hAdapter,int iItem,int iColumn);
XC_API BOOL WINAPI XAdapterTable_GetItemTextEx(HXCGUI hAdapter,int iItem,const wchar_t* pName,out_ wchar_t* pOut,int nOutLen);
XC_API HIMAGE WINAPI XAdapterTable_GetItemImageEx(HXCGUI hAdapter,int iItem,const wchar_t* pName);
XC_API HXCGUI WINAPI XAdapterTree_Create();
XC_API int WINAPI XAdapterTree_AddColumn(HXCGUI hAdapter,const wchar_t* pName);
XC_API int WINAPI XAdapterTree_SetColumn(HXCGUI hAdapter,const wchar_t* pColName);
XC_API int WINAPI XAdapterTree_InsertItemText(HXCGUI hAdapter,const wchar_t* pValue,int nParentID=XC_ID_ROOT,int insertID=XC_ID_LAST);
XC_API int WINAPI XAdapterTree_InsertItemTextEx(HXCGUI hAdapter,const wchar_t* pName,const wchar_t* pValue,int nParentID=XC_ID_ROOT,int insertID=XC_ID_LAST);
XC_API int WINAPI XAdapterTree_InsertItemImage(HXCGUI hAdapter,HIMAGE hImage,int nParentID=XC_ID_ROOT,int insertID=XC_ID_LAST);
XC_API int WINAPI XAdapterTree_InsertItemImageEx(HXCGUI hAdapter,const wchar_t* pName,HIMAGE hImage,int nParentID=XC_ID_ROOT,int insertID=XC_ID_LAST);
XC_API int WINAPI XAdapterTree_GetCount(HXCGUI hAdapter);
XC_API int WINAPI XAdapterTree_GetCountColumn(HXCGUI hAdapter);
XC_API BOOL WINAPI XAdapterTree_SetItemText(HXCGUI hAdapter,int nID,int iColumn,const wchar_t* pValue);
XC_API BOOL WINAPI XAdapterTree_SetItemTextEx(HXCGUI hAdapter,int nID,const wchar_t* pName,const wchar_t* pValue);
XC_API BOOL WINAPI XAdapterTree_SetItemImage(HXCGUI hAdapter,int nID,int iColumn,HIMAGE hImage);
XC_API BOOL WINAPI XAdapterTree_SetItemImageEx(HXCGUI hAdapter,int nID,const wchar_t* pName,HIMAGE hImage);
XC_API BOOL WINAPI XAdapterTree_GetItemTextEx(HXCGUI hAdapter,int nID,const wchar_t* pName,out_ wchar_t* pOut,int nOutLen);
XC_API HIMAGE WINAPI XAdapterTree_GetItemImageEx(HXCGUI hAdapter,int nID,const wchar_t* pName);
XC_API BOOL WINAPI XAdapterTree_DeleteItem(HXCGUI hAdapter,int nID);
XC_API void WINAPI XAdapterTree_DeleteItemAll(HXCGUI hAdapter);
XC_API void WINAPI XAdapterTree_DeleteColumnAll(HXCGUI hAdapter);
XC_API HXCGUI WINAPI XAdapterMap_Create();
XC_API BOOL WINAPI XAdapterMap_AddItemText(HXCGUI hAdapter,const wchar_t* pName,const wchar_t* pValue);
XC_API BOOL WINAPI XAdapterMap_AddItemImage(HXCGUI hAdapter,const wchar_t* pName,HIMAGE hImage);
XC_API BOOL WINAPI XAdapterMap_DeleteItem(HXCGUI hAdapter,const wchar_t* pName);
XC_API int  WINAPI XAdapterMap_GetCount(HXCGUI hAdapter);
XC_API BOOL WINAPI XAdapterMap_GetItemText(HXCGUI hAdapter,const wchar_t* pName,out_ wchar_t* pOut,int nOutLen);
XC_API HIMAGE WINAPI XAdapterMap_GetItemImage(HXCGUI hAdapter,const wchar_t* pName);
XC_API void WINAPI _XC_DebugPrintf(int  level, const char *pFormat, ...); 
XC_API void WINAPI _XC_DebugPrintfW(int  level, const wchar_t *pFormat, ...);
XC_API void WINAPI xtrace(const char *pFormat, ...);
XC_API void WINAPI xtracew(const wchar_t *pFormat, ...);
XC_API HDRAW WINAPI XDraw_Create(HDC hdc); //创建
XC_API void  WINAPI XDraw_Destroy(HDRAW hDraw); //销毁
XC_API void  WINAPI XDraw_SetOffset(HDRAW hDraw, int x, int y); //设置坐标偏移量
XC_API void  WINAPI XDraw_GetOffset(HDRAW hDraw, out_ int *pX, out_ int *pY); //获取坐标偏移量
XC_API void  WINAPI XDraw_RestoreGDIOBJ(HDRAW hDraw);  //还原状态,释放用户绑定的GDI对象
XC_API HDC   WINAPI XDraw_GetHDC(HDRAW hDraw);
XC_API void WINAPI XDraw_SetBrushColor(HDRAW hDraw, COLORREF color,BYTE alpha=254);
XC_API void WINAPI XDraw_SetTextAlign(HDRAW hDraw, int nFlag);
XC_API void WINAPI XDraw_SetTextVertical(HDRAW hDraw, BOOL bVertical);
XC_API void WINAPI XDraw_SetFontX(HDRAW hDraw, HFONTX  hFontx);
XC_API void WINAPI XDraw_SetFont(HDRAW hDraw, int size);
XC_API void WINAPI XDraw_SetFont2(HDRAW hDraw, const wchar_t* pName,int size);
XC_API void WINAPI XDraw_SetFont3(HDRAW hDraw, LOGFONT* pLogFont);
XC_API void WINAPI XDraw_SetLineWidth(HDRAW hDraw, int nWidth);
XC_API int  WINAPI XDraw_SetBkMode(HDRAW hDraw, BOOL bTransparent);
XC_API HBRUSH WINAPI XDraw_CreateSolidBrush(HDRAW hDraw, COLORREF crColor);
XC_API HPEN WINAPI XDraw_CreatePen(HDRAW hDraw, int fnPenStyle, int nWidth, COLORREF crColor);
XC_API HRGN WINAPI XDraw_CreateRectRgn(HDRAW hDraw, int nLeftRect, int nTopRect, int nRightRect, int nBottomRect);
XC_API HRGN WINAPI XDraw_CreateRoundRectRgn(HDRAW hDraw, int nLeftRect, int nTopRect, int nRightRect, int nBottomRect, int nWidthEllipse, int nHeightEllipse);
XC_API HRGN WINAPI XDraw_CreatePolygonRgn(HDRAW hDraw, POINT *pPt, int cPoints, int fnPolyFillMode);
XC_API int WINAPI XDraw_SelectClipRgn(HDRAW hDraw, HRGN hRgn);
XC_API void WINAPI XDraw_FillRect(HDRAW hDraw, RECT *pRect);
XC_API void WINAPI XDraw_FillRectColor(HDRAW hDraw, RECT *pRect,COLORREF color,BYTE alpha=255);
XC_API BOOL WINAPI XDraw_FillRgn(HDRAW hDraw, HRGN hrgn, HBRUSH hbr);
XC_API void WINAPI XDraw_FillEllipse(HDRAW hDraw, RECT *pRect);
XC_API void WINAPI XDraw_DrawEllipse(HDRAW hDraw, RECT *pRect);
XC_API void WINAPI XDraw_FillRoundRect(HDRAW hDraw,RECT *pRect,int nWidth,int nHeight); //填充圆角矩形
XC_API void WINAPI XDraw_DrawRoundRect(HDRAW hDraw,RECT *pRect,int nWidth,int nHeight); //绘制圆角矩形边框
XC_API BOOL WINAPI XDraw_Rectangle(HDRAW hDraw, int nLeftRect, int nTopRect, int nRightRect, int nBottomRect);
XC_API void WINAPI XDraw_DrawGroupBox_Rect(HDRAW hDraw, RECT *pRect,const wchar_t* pName,COLORREF textColor,BYTE textAlpha,POINT* pOffset);
XC_API void WINAPI XDraw_DrawGroupBox_RoundRect(HDRAW hDraw, RECT *pRect,const wchar_t* pName,COLORREF textColor,BYTE textAlpha,POINT* pOffset,int nWidth,int nHeight);
XC_API void WINAPI XDraw_GradientFill2(HDRAW hDraw, COLORREF color1, BYTE alpha1, COLORREF color2, BYTE alpha2, RECT *pRect, int mode);
XC_API BOOL WINAPI XDraw_GradientFill4(HDRAW hDraw, COLORREF color1, COLORREF color2, COLORREF color3, COLORREF color4, RECT *pRect, int mode);
XC_API BOOL WINAPI XDraw_FrameRgn(HDRAW hDraw, HRGN hrgn, HBRUSH hbr, int nWidth, int nHeight);
XC_API void  WINAPI XDraw_FrameRect(HDRAW hDraw, RECT *pRect);
XC_API void WINAPI XDraw_FocusRect(HDRAW hDraw, RECT *pRect);
XC_API void WINAPI XDraw_DrawLine(HDRAW hDraw,int x1,int y1,int x2,int y2);
XC_API BOOL WINAPI XDraw_MoveToEx(HDRAW hDraw, int X, int Y, LPPOINT lpPoint = NULL);
XC_API BOOL WINAPI XDraw_LineTo(HDRAW hDraw, int nXEnd, int nYEnd);
XC_API BOOL WINAPI XDraw_Polyline(HDRAW hDraw, in_ POINT *pArrayPt, int arrayPtSize);
XC_API void WINAPI XDraw_Dottedline(HDRAW hDraw, int x1, int y1, int x2, int y2);  //绘制虚线
XC_API COLORREF WINAPI XDraw_SetPixel(HDRAW hDraw, int X, int Y, COLORREF crColor);
XC_API void WINAPI XDraw_Check(HDRAW hDraw, int x, int y, COLORREF color, BOOL bCheck);
XC_API BOOL WINAPI XDraw_DrawIconEx(HDRAW hDraw, int xLeft, int yTop, HICON hIcon, int cxWidth, int cyWidth, UINT istepIfAniCur, HBRUSH hbrFlickerFreeDraw, UINT diFlags);
XC_API BOOL WINAPI XDraw_BitBlt(HDRAW hDrawDest, int nXDest, int nYDest, int nWidth, int nHeight, HDC hdcSrc, int nXSrc, int nYSrc, DWORD dwRop);
XC_API BOOL WINAPI XDraw_BitBlt2(HDRAW hDrawDest, int nXDest, int nYDest, int nWidth, int nHeight, HDRAW hDrawSrc, int nXSrc, int nYSrc, DWORD dwRop);
XC_API BOOL WINAPI XDraw_AlphaBlend(HDRAW hDraw, int nXOriginDest, int nYOriginDest, int nWidthDest, int nHeightDest, HDC hdcSrc, int nXOriginSrc,int nYOriginSrc, int nWidthSrc, int nHeightSrc, int alpha);
XC_API void WINAPI XDraw_TriangularArrow(HDRAW hDraw, int align, int x, int y, int width, int height);
XC_API void WINAPI XDraw_Image(HDRAW hDraw, HIMAGE hImage, int x, int y);
XC_API void WINAPI XDraw_Image2(HDRAW hDraw, HIMAGE hImage, int x, int y, int width, int height);
XC_API void WINAPI XDraw_ImageStretch(HDRAW hDraw, HIMAGE hImage, int x, int y, int width, int height);
XC_API void WINAPI XDraw_ImageAdaptive(HDRAW hDraw, HIMAGE hImage, RECT *pRect, BOOL bOnlyBorder=FALSE);
XC_API void WINAPI XDraw_ImageExTile(HDRAW hDraw, HIMAGE hImage, RECT *pRect, int flag = 0);
XC_API void WINAPI XDraw_ImageSuper(HDRAW hDraw, HIMAGE hImage, RECT *pRect,BOOL bClip=FALSE);
XC_API void WINAPI XDraw_ImageSuper2(HDRAW hDraw, HIMAGE hImage, RECT *pRcDest, RECT *pSrcRect);
XC_API void WINAPI XDraw_DrawText(HDRAW hDraw, const wchar_t * lpString, int nCount, RECT* lpRect);
XC_API void WINAPI XDraw_DrawTextUnderline(HDRAW hDraw, const wchar_t * lpString, int nCount, RECT* lpRect,COLORREF colorLine,BYTE alphaLine=254);
XC_API void WINAPI XDraw_TextOut(HDRAW hDraw, int nXStart, int nYStart, const wchar_t * lpString, int cbString);
XC_API void WINAPI XDraw_TextOutA(HDRAW hDraw, int nXStart, int nYStart, const char * lpString);
XC_API void WINAPI XDraw_SetAlpha(HDRAW hDraw,BYTE alpha);  //设置透明通道度
XC_API void WINAPI XDraw_SetAlphaEx(HDC hdc,BYTE alpha);  //设置透明通道度
XC_API HELE WINAPI XEle_Create(int x, int y, int cx, int cy, HXCGUI hParent = NULL);
XC_API void WINAPI XEle_RegEventEx(HELE hEle, CEventBase* pEvent); //内部使用
XC_API BOOL WINAPI XEle_RemoveEventEx(HELE hEle, CEventBase *pEvent); //内部使用
XC_API void WINAPI XEle_RegEventC(HELE hEle, int nEvent, void *pFun); //注册事件C方式
XC_API void WINAPI XEle_RegEventC1(HELE hEle, int nEvent, void *pFun); 
XC_API void WINAPI XEle_RegEventC2(HELE hEle, int nEvent, void *pFun); 
XC_API BOOL WINAPI XEle_RemoveEventC(HELE hEle, int nEvent, void *pFun); //移除事件函数C方式
XC_API int  WINAPI XEle_SendEvent(HELE hEle, HELE hEventEle, int nEvent, WPARAM wParam, LPARAM lParam);
XC_API BOOL WINAPI XEle_PostEvent(HELE hEle, HELE hEventEle, int nEvent, WPARAM wParam, LPARAM lParam);
XC_API BOOL WINAPI XEle_IsShow(HELE hEle);
XC_API BOOL WINAPI XEle_IsEnable(HELE hEle);
XC_API BOOL WINAPI XEle_IsEnableFocus(HELE hEle);
XC_API BOOL WINAPI XEle_IsDrawFocus(HELE hEle);
XC_API BOOL WINAPI XEle_IsEnableEvent_XE_PAINT_END(HELE hEle);
XC_API BOOL WINAPI XEle_IsMouseThrough(HELE hEle);
XC_API BOOL WINAPI XEle_IsBkTransparent(HELE hEle);
XC_API BOOL WINAPI XEle_IsKeyTab(HELE hEle);
XC_API BOOL WINAPI XEle_IsSwitchFocus(HELE hEle);
XC_API BOOL WINAPI XEle_IsEnable_XE_MOUSEWHEEL(HELE hEle);
XC_API void WINAPI XEle_Enable(HELE hEle, BOOL bEnable);
XC_API void WINAPI XEle_EnableFocus(HELE hEle, BOOL bEnable);
XC_API void WINAPI XEle_EnableDrawFocus(HELE hEle, BOOL bEnable);
XC_API void WINAPI XEle_EnableEvent_XE_PAINT_END(HELE hEle,BOOL bEnable);
XC_API void WINAPI XEle_EnableBkTransparent(HELE hEle,BOOL bEnable);
XC_API void WINAPI XEle_EnableMouseThrough(HELE hEle,BOOL bEnable);
XC_API void WINAPI XEle_EnableKeyTab(HELE hEle,BOOL bEnable);
XC_API void WINAPI XEle_EnableSwitchFocus(HELE hEle,BOOL bEnable);
XC_API void WINAPI XEle_EnableEvent_XE_MOUSEWHEEL(HELE hEle,BOOL bEnable);
XC_API BOOL WINAPI XEle_SetRect(HELE hEle, RECT *pRect,BOOL bRedraw=FALSE);
XC_API BOOL WINAPI XEle_SetRectLogic(HELE hEle, RECT *pRect, BOOL bRedraw=FALSE); //逻辑模式坐标
XC_API void WINAPI XEle_GetRect(HELE hEle, RECT *pRect);   //相对与父坐标,人眼观察模式
XC_API void WINAPI XEle_GetRectLogic(HELE hEle, RECT *pRect); //相对与父坐标,逻辑模式
XC_API void WINAPI XEle_GetClientRect(HELE hEle, RECT *pRect);  //左上角为0，0坐标
XC_API void WINAPI XEle_GetWndClientRect(HELE hEle,out_ RECT *pRect);
XC_API int  WINAPI XEle_GetWidth(HELE hEle);
XC_API int  WINAPI XEle_GetHeight(HELE hEle);
XC_API void WINAPI XEle_RectWndClientToEleClient(HELE hEle, in_out_ RECT *pRect);
XC_API void WINAPI XEle_PointWndClientToEleClient(HELE hEle, in_out_ POINT *pPt); 
XC_API void WINAPI XEle_RectClientToWndClient(HELE hEle,in_out_ RECT *pRect);
XC_API void WINAPI XEle_PointClientToWndClient(HELE hEle, in_out_ POINT *pPt);
XC_API BOOL WINAPI XEle_AddEle(HELE hEle, HELE hChildEle);
XC_API BOOL WINAPI XEle_InsertEle(HELE hEle, HELE hChildEle,HELE hDestEle);  //插入到指定元素前面
XC_API void WINAPI XEle_RemoveEle(HELE hEle);
XC_API BOOL WINAPI XEle_AddShape(HELE hEle,HXCGUI hShape);
XC_API void WINAPI XEle_ShowEle(HELE hEle, BOOL bShow);
XC_API XC_OBJECT_TYPE WINAPI XEle_GetType(HELE hEle);
XC_API HWND WINAPI XEle_GetHWND(HELE hEle);
XC_API HWINDOW WINAPI XEle_GetHWINDOW(HELE hEle);
XC_API void WINAPI XEle_SetCursor(HELE hEle, HCURSOR hCursor);
XC_API HCURSOR WINAPI XEle_GetCursor(HELE hEle);
XC_API void WINAPI XEle_SetID(HELE hEle, int id);
XC_API int  WINAPI XEle_GetID(HELE hEle);
XC_API HELE WINAPI XEle_GetParentEle(HELE hEle);
XC_API HXCGUI WINAPI XEle_GetParent(HELE hEle);
XC_API void WINAPI XEle_SetTextColor(HELE hEle, COLORREF color,BYTE alpha=254);
XC_API COLORREF WINAPI XEle_GetTextColor(HELE hEle);
XC_API void WINAPI XEle_SetFocusBorderColor(HELE hEle, COLORREF color,BYTE alpha=255);
XC_API COLORREF WINAPI XEle_GetFocusBorderColor(HELE hEle);
XC_API void WINAPI XEle_SetFont(HELE hEle,HFONTX hFontx);
XC_API HFONTX WINAPI XEle_GetFont(HELE hEle);
XC_API void WINAPI XEle_SetAlpha(HELE hEle,BYTE alpha);
XC_API int  WINAPI XEle_GetChildCount(HELE hEle);
XC_API HELE WINAPI XEle_GetChildByIndex(HELE hEle,int index);
XC_API HELE WINAPI XEle_GetChildByID(HELE hEle,int id);
XC_API int  WINAPI XEle_GetChildShapeCount(HELE hEle);
XC_API HXCGUI WINAPI XEle_GetChildShapeByIndex(HELE hEle, int index);
XC_API HELE WINAPI XEle_HitChildEle(HELE hEle,in_ POINT  *pPt);  //判断坐标点在哪个元素上
XC_API void WINAPI XEle_BindLayoutObject(HELE hEle,HXCGUI hLayout);
XC_API HXCGUI WINAPI XEle_GetLayoutObject(HELE hEle);
XC_API void WINAPI XEle_SetUserData(HELE hEle,int nData);
XC_API int  WINAPI XEle_GetUserData(HELE hEle);
XC_API void WINAPI XEle_GetContentSize(HELE hEle,out_ SIZE* pSize);
XC_API void WINAPI XEle_SetCapture(HELE hEle,BOOL b);
XC_API void WINAPI XEle_SetLayoutWidth(HELE hEle,layout_size_type_ nType,int nWidth);
XC_API void WINAPI XEle_SetLayoutHeight(HELE hEle,layout_size_type_ nType,int nHeight);
XC_API void WINAPI XEle_GetLayoutWidth(HELE hEle, out_ layout_size_type_ *pType,out_ int *pWidth);
XC_API void WINAPI XEle_GetLayoutHeight(HELE hEle, out_ layout_size_type_ *pType,out_ int *pHeight);
XC_API void WINAPI XEle_RedrawEle(HELE hEle);
XC_API void WINAPI XEle_RedrawRect(HELE hEle,RECT *pRect);
XC_API void WINAPI XEle_Destroy(HELE hEle);  //销毁
XC_API void WINAPI XEle_AddBkBorder(HELE hEle,COLORREF color,BYTE alpha, int width);
XC_API void WINAPI XEle_AddBkFill(HELE hEle,COLORREF color,BYTE alpha);
XC_API void WINAPI XEle_AddBkImage(HELE hEle,HIMAGE hImage);
XC_API int  WINAPI XEle_GetBkInfoCount(HELE hEle);
XC_API void WINAPI XEle_ClearBkInfo(HELE hEle);
XC_API HBKINFOM WINAPI XEle_GetBkInfoManager(HELE hEle);
XC_API BOOL WINAPI XEle_DrawFocus(HELE hEle,HDRAW hDraw,RECT* pRect);
XC_API void WINAPI XEle_EnableTransparentChannel(HELE hEle,BOOL bEnable);
XC_API void WINAPI XEle_SetToolTip(HELE hEle, const wchar_t* pText);
XC_API void WINAPI XEle_GetToolTip(HELE hEle, out_ wchar_t  *pOut, int nOutLen);
XC_API void WINAPI XEle_EnableToolTip(HELE hEle, BOOL bEnable);
XC_API void WINAPI XEle_AdjustLayoutObject(HELE hEle);
XC_API void WINAPI XEle_AdjustLayout(HELE hEle);
XC_API BOOL WINAPI XEle_RegEventTest(int nEvent,BYTE* pParamType); //检测回调函数参数是否正确
XC_API HFONTX WINAPI XFont_Create(int size);//创建字体
XC_API HFONTX WINAPI XFont_Create2(const wchar_t *pName=L"宋体",int size=12,BOOL bBold=FALSE,BOOL bItalic=FALSE,BOOL bUnderline=FALSE,BOOL bStrikeOut=FALSE);
XC_API HFONTX WINAPI XFont_CreateEx(LOGFONTW *pFontInfo);
XC_API HFONTX WINAPI XFont_CreateFromHFONT(HFONT hFont);
XC_API void   WINAPI XFont_EnableAutoDestroy(HFONTX hFontX,BOOL bEnable);
XC_API void   WINAPI XFont_Destroy(HFONTX hFontX);  //销毁字体
XC_API HFONT  WINAPI XFont_GetHFONT(HFONTX hFontX);
XC_API HWINDOW WINAPI XFrameWnd_Create(int x, int y, int cx, int cy, wchar_t *pTitle, HWND hWndParent, int XCStyle);  //创建窗口
XC_API HWINDOW WINAPI XFrameWnd_CreateEx(DWORD dwExStyle, wchar_t* lpClassName, wchar_t* lpWindowName, DWORD dwStyle, int x, int y, int cx, int cy, HWND hWndParent, int XCStyle);  //创建窗口
XC_API void WINAPI XFrameWnd_GetLayoutAreaRect(HWINDOW hWindow,out_ RECT *pRect); //获取客户区布局区域坐标
XC_API void WINAPI XFrameWnd_SetView(HWINDOW hWindow,HELE hEle);
XC_API void WINAPI XFrameWnd_SetPaneSplitBarColor(HWINDOW hWindow, COLORREF color, BYTE alpha=255);
XC_API BOOL WINAPI XFrameWnd_AddPane(HWINDOW hWindow,HELE hPaneDest,HELE hPaneNew,align_type_ align);
XC_API BOOL WINAPI XFrameWnd_MergePane(HWINDOW hWindow,HELE hPaneDest,HELE hPaneNew);
XC_API HIMAGE WINAPI XImage_LoadFile(const wchar_t *pImageName,BOOL bStretch=FALSE);//从文件中加载图片
XC_API HIMAGE WINAPI XImage_LoadFileAdaptive(const wchar_t *pImageName,int leftSize,int topSize,int rightSize,int bottomSize);//从文件中加载图片
XC_API HIMAGE WINAPI XImage_LoadFileRect(const wchar_t *pImageName,int x,int y,int cx,int cy);  //加载图片,指定区域
XC_API HIMAGE WINAPI XImage_LoadResAdaptive(int id,const wchar_t *pType,int leftSize,int topSize,int rightSize,int bottomSize);//从资源中加载图片
XC_API HIMAGE WINAPI XImage_LoadRes(int id,const wchar_t *pType,BOOL bStretch=FALSE);//从资源中加载图片
XC_API HIMAGE WINAPI XImage_LoadZip(const wchar_t *pZipFileName,const wchar_t *pImageName,const wchar_t *pPassword=NULL,BOOL bStretch=FALSE);//从ZIP中加载图片
XC_API HIMAGE WINAPI XImage_LoadZipAdaptive(const wchar_t *pZipFileName,const wchar_t *pImageName,const wchar_t *pPassword,int x1,int x2,int y1,int y2);//从ZIP中加载图片
XC_API HIMAGE WINAPI XImage_LoadZipRect(const wchar_t *pZipFileName,const wchar_t *pImageName,const wchar_t *pPassword,int x,int y,int cx,int cy);
XC_API HIMAGE WINAPI XImage_LoadMemory(void* pBuffer,int nSize,BOOL bStretch);
XC_API HIMAGE WINAPI XImage_LoadMemoryRect(void* pBuffer,int nSize,int x,int y,int cx,int cy,BOOL bStretch);
XC_API HIMAGE WINAPI XImage_LoadMemoryAdaptive(void* pBuffer,int nSize,int leftSize,int topSize,int rightSize,int bottomSize);
XC_API HIMAGE WINAPI XImage_LoadFromImage(void *pImage);
XC_API HIMAGE WINAPI XImage_LoadFileFromExtractIcon(const wchar_t *pImageName);
XC_API HIMAGE WINAPI XImage_LoadFileFromHICON(HICON hIcon);
XC_API HIMAGE WINAPI XImage_LoadFileFromHBITMAP(HBITMAP hBitmap);
XC_API BOOL WINAPI XImage_IsStretch(HIMAGE hImage);   //是否拉伸图片
XC_API BOOL WINAPI XImage_IsAdaptive(HIMAGE hImage);  //是否为自适应图片
XC_API BOOL WINAPI XImage_IsTile(HIMAGE hImage);      //是否为平铺图片
XC_API BOOL WINAPI XImage_SetDrawType(HIMAGE hImage,image_draw_type_ nType);//设置图片绘制类型
XC_API BOOL WINAPI XImage_SetDrawTypeAdaptive(HIMAGE hImage,int leftSize,int topSize,int rightSize,int bottomSize);//设置图片自适应
XC_API void WINAPI XImage_SetTranColor(HIMAGE hImage,COLORREF color); //设置透明色
XC_API void WINAPI XImage_SetTranColorEx(HIMAGE hImage,COLORREF color,BYTE tranColor); //设置透明色
XC_API void WINAPI XImage_EnableTranColor(HIMAGE hImage,BOOL bEnable); //启用透明色
XC_API void WINAPI XImage_EnableAutoDestroy(HIMAGE hImage,BOOL bEnable); //启用或关闭自动销毁,当与UI元素关联时有效
XC_API void WINAPI XImage_EnableCenter(HIMAGE hImage,BOOL bCenter); //bCenter
XC_API BOOL WINAPI XImage_IsCenter(HIMAGE hImage);
XC_API image_draw_type_ WINAPI XImage_GetDrawType(HIMAGE hImage); //获取图片绘制类型
XC_API int WINAPI XImage_GetWidth(HIMAGE hImage);
XC_API int WINAPI XImage_GetHeight(HIMAGE hImage);
XC_API void WINAPI XImage_Release(HIMAGE hImage);
XC_API void WINAPI XImage_Destroy(HIMAGE hImage);
XC_API HIMAGE WINAPI _Image_AddRef(HIMAGE hImage); //增加引用计数
XC_API void   WINAPI _Image_Release(HIMAGE hImage); //释放引用计数
XC_API HXCGUI WINAPI XLayout_Create();
XC_API void WINAPI XLayout_Destroy(HXCGUI hLayout);
XC_API void WINAPI XLayout_AddEle(HXCGUI hLayout,HELE hEle);
XC_API void WINAPI XLayout_AddLayoutObject(HXCGUI hLayout,HXCGUI hLayoutObject);
XC_API void WINAPI XLayout_AddShape(HXCGUI hLayout,HXCGUI hShape);
XC_API void WINAPI XLayout_RemoveChild(HXCGUI hLayout,HXCGUI hChild);
XC_API void WINAPI XLayout_AdjustLayout(HXCGUI hLayout);
XC_API void WINAPI XLayout_SetRectFrame(HXCGUI hLayout,RECT *pRect);
XC_API void WINAPI XLayout_GetRect(HXCGUI hLayout,out_ RECT *pRect);
XC_API void WINAPI XLayout_GetRectFrame(HXCGUI hLayout,out_ RECT *pRect);
XC_API HWINDOW WINAPI XLayout_GetWindow(HXCGUI hLayout);
XC_API HELE WINAPI XLayout_GetEle(HXCGUI hLayout);
XC_API HXCGUI WINAPI XLayout_GetParentLayout(HXCGUI hLayout);
XC_API HXCGUI WINAPI XLayout_GetParent(HXCGUI hLayout);
XC_API void WINAPI XLayout_SetID(HXCGUI hLayout,int nID);
XC_API int  WINAPI XLayout_GetID(HXCGUI hLayout);
XC_API void WINAPI XLayout_SetHorizon(HXCGUI hLayout,BOOL bHorizon);
XC_API void WINAPI XLayout_SetAlignH(HXCGUI hLayout,layout_align_ nAlign);
XC_API void WINAPI XLayout_SetAlignV(HXCGUI hLayout,layout_align_ nAlign);
XC_API void WINAPI XLayout_SetPadding(HXCGUI hLayout,int left,int top,int right,int bottom);
XC_API void WINAPI XLayout_SetSpace(HXCGUI hLayout,int nSpace);
XC_API void WINAPI XLayout_SetLayoutWidth(HXCGUI hLayout,layout_size_type_ nType,int nWidth);
XC_API void WINAPI XLayout_SetLayoutHeight(HXCGUI hLayout,layout_size_type_ nType,int nHeight);
XC_API void WINAPI XLayout_GetLayoutWidth(HXCGUI hLayout,out_ layout_size_type_ *pType,out_ int *pWidth);
XC_API void WINAPI XLayout_GetLayoutHeight(HXCGUI hLayout,out_ layout_size_type_ *pType,out_ int *pHeight);
XC_API int WINAPI XLayout_GetWidth(HXCGUI hLayout);
XC_API int WINAPI XLayout_GetHeight(HXCGUI hLayout);
XC_API int WINAPI XLayout_GetWidthIn(HXCGUI hLayout);
XC_API int WINAPI XLayout_GetHeightIn(HXCGUI hLayout);
XC_API void WINAPI XLayout_GetContentSize(HXCGUI hLayout,out_ SIZE* pSize);
XC_API void WINAPI XLayout_ShowLayout(HXCGUI hLayout,BOOL bShow);
XC_API int WINAPI XLayout_GetChildCount(HXCGUI hLayout);
XC_API XC_OBJECT_TYPE WINAPI XLayout_GetChildType(HXCGUI hLayout,int index);
XC_API HXCGUI WINAPI XLayout_GetChild(HXCGUI hLayout,int index);
XC_API void WINAPI XLayout_Draw(HXCGUI hLayout,HDRAW hDraw);
XC_API void WINAPI Layout_SetEle(HXCGUI hLayout,HELE hParent);
XC_API HELE WINAPI XListBox_Create(int x, int y, int cx, int cy,HXCGUI hParent=NULL);
XC_API void WINAPI XListBox_SetDrawItemBkFlags(HELE hEle, int nFlags);
XC_API BOOL WINAPI XListBox_SetItemData(HELE hEle, int iItem, int nUserData);
XC_API int  WINAPI XListBox_GetItemData(HELE hEle,int iItem);
XC_API BOOL WINAPI XListBox_SetItemInfo(HELE hEle,int iItem,in_ listBox_item_Info_i *pItem);
XC_API BOOL WINAPI XListBox_GetItemInfo(HELE hEle,int iItem,out_ listBox_item_Info_i *pItem);
XC_API void WINAPI XListBox_AddItemBkBorder(HELE hEle,list_item_state_ nState, COLORREF color,BYTE alpha, int width);
XC_API void WINAPI XListBox_AddItemBkFill(HELE hEle,list_item_state_ nState, COLORREF color,BYTE alpha);
XC_API void WINAPI XListBox_AddItemBkImage(HELE hEle,list_item_state_ nState, HIMAGE hImage);
XC_API int  WINAPI XListBox_GetItemBkInfoCount(HELE hEle, list_item_state_ nState);
XC_API void WINAPI XListBox_ClearItemBkInfo(HELE hEle,list_item_state_ nState);
XC_API HBKINFOM WINAPI XListBox_GetItemBkInfoManager(HELE hEle,list_item_state_ nState);
XC_API BOOL WINAPI XListBox_SetSelectItem(HELE hEle,int iItem);
XC_API int  WINAPI XListBox_GetSelectItem(HELE hEle);
XC_API BOOL WINAPI XListBox_CancelSelectItem(HELE hEle,int iItem);
XC_API BOOL WINAPI XListBox_CancelSelectAll(HELE hEle);
XC_API int  WINAPI XListBox_GetSelectAll(HELE hEle,out_ int *pArray,int nArraySize);
XC_API int  WINAPI XListBox_GetSelectCount(HELE hEle);
XC_API int  WINAPI XListBox_GetItemMouseStay(HELE hEle);
XC_API BOOL WINAPI XListBox_SelectAll(HELE hEle);
XC_API void WINAPI XListBox_SetItemHeightDefault(HELE hEle,int nHeight,int nSelHeight);
XC_API void WINAPI XListBox_GetItemHeightDefault(HELE hEle,out_ int *pHeight,out_ int *pSelHeight);
XC_API int  WINAPI XListBox_GetItemIndexFromHXCGUI(HELE hEle,HXCGUI hXCGUI);
XC_API int  WINAPI XListBox_HitTest(HELE hEle,POINT *pPt);
XC_API int  WINAPI XListBox_HitTestOffset(HELE hEle,POINT *pPt); //自动添加滚动视图偏移量
XC_API BOOL WINAPI XListBox_SetItemTemplateXML(HELE hEle,const wchar_t* pXmlFile);
XC_API BOOL WINAPI XListBox_SetItemTemplateXMLFromString(HELE hEle,const char* pStringXML);
XC_API HXCGUI WINAPI XListBox_GetTemplateObject(HELE hEle,int iItem,int nTempItemID);
XC_API void WINAPI XListBox_EnableMultiSel(HELE hEle,BOOL bEnable);
XC_API void WINAPI XListBox_BindAdapter(HELE hEle,HXCGUI hAdapter);
XC_API HXCGUI WINAPI XListBox_GetAdapter(HELE hEle);
XC_API HELE WINAPI XList_Create(int x,int y,int cx,int cy,HXCGUI hParent=NULL);
XC_API int WINAPI XList_AddColumn(HELE hEle,int width); //增加列
XC_API int WINAPI XList_InsertColumn(HELE hEle,int width,int iItem);
XC_API void WINAPI XList_EnableMultiSel(HELE hEle, BOOL bEnable);
XC_API void WINAPI XList_SetDrawItemBkFlags(HELE hEle,int style);
XC_API void WINAPI XList_SetColumnWidth(HELE hEle,int iItem,int width);
XC_API void WINAPI XList_SetColumnMinWidth(HELE hEle,int iItem,int width);
XC_API BOOL WINAPI XList_DeleteColumn(HELE hEle,int iItem);
XC_API void WINAPI XList_DeleteColumnAll(HELE hEle);
XC_API BOOL WINAPI XList_SetItemData(HELE hEle,int iItem,int iSubItem,int data);
XC_API int  WINAPI XList_GetItemData(HELE hEle,int iItem,int iSubItem);
XC_API BOOL WINAPI XList_SetSelectItem(HELE hEle,int iItem,BOOL bSelect);
XC_API int  WINAPI XList_GetSelectItem(HELE hEle);
XC_API int  WINAPI XList_GetSelectItemCount(HELE hEle);
XC_API void WINAPI XList_SelectAll(HELE hEle);
XC_API int  WINAPI XList_GetSelectAll(HELE hEle,out_ int *pArray,int nArraySize);
XC_API HELE WINAPI XList_GetHeaderHELE(HELE hEle);
XC_API void WINAPI XList_BindAdapter(HELE hEle,HXCGUI hAdapter);
XC_API void WINAPI XList_BindAdapterHeader(HELE hEle,HXCGUI hAdapter);
XC_API HXCGUI WINAPI XList_GetAdapter(HELE hEle);
XC_API BOOL WINAPI XList_SetItemTemplateXML(HELE hEle, const wchar_t* pXmlFile);
XC_API BOOL WINAPI XList_SetItemTemplateXMLFromString(HELE hEle, const char* pStringXML);
XC_API HXCGUI WINAPI XList_GetTemplateObject(HELE hEle, int iItem,int nTempItemID); //通过模板项ID,获取实例化模板项ID对应的对象.
XC_API int    WINAPI XList_GetItemIndexFromHXCGUI(HELE hEle, HXCGUI hXCGUI);
XC_API void WINAPI XList_SetHeaderHeight(HELE hEle, int height);
XC_API int  WINAPI XList_GetHeaderHeight(HELE hEle);
XC_API void WINAPI XList_AddItemBkBorder(HELE hEle, list_item_state_ nState, COLORREF color,BYTE alpha, int width);
XC_API void WINAPI XList_AddItemBkFill(HELE hEle, list_item_state_ nState, COLORREF color,BYTE alpha);
XC_API void WINAPI XList_AddItemBkImage(HELE hEle, list_item_state_ nState, HIMAGE hImage);
XC_API int  WINAPI XList_GetItemBkInfoCount(HELE hEle, list_item_state_ nState);
XC_API void WINAPI XList_ClearItemBkInfo(HELE hEle, list_item_state_ nState);
XC_API HBKINFOM WINAPI XList_GetItemBkInfoManager(HELE hEle, list_item_state_ nState);
XC_API void WINAPI XList_SetItemHeightDefault(HELE hEle,int nHeight,int nSelHeight);
XC_API void WINAPI XList_GetItemHeightDefault(HELE hEle,out_ int *pHeight,out_ int *pSelHeight);
XC_API BOOL WINAPI XList_HitTest(HELE hEle,POINT *pPt,out_ int *piItem,out_ int *piSubItem);
XC_API BOOL WINAPI XList_HitTestOffset(HELE hEle,POINT *pPt,out_ int *piItem,out_ int *piSubItem);
XC_API void WINAPI XList_RefreshData(HELE hEle);
XC_API HELE WINAPI XListView_Create(int x, int y, int cx, int cy,HXCGUI hParent=NULL);
XC_API void WINAPI XListView_BindAdapter(HELE hEle,HXCGUI hAdapter);
XC_API HXCGUI WINAPI XListView_GetAdapter(HELE hEle);
XC_API BOOL WINAPI XListView_SetItemTemplateXML(HELE hEle,const wchar_t* pXmlFile);
XC_API BOOL WINAPI XListView_SetItemTemplateXMLFromString(HELE hEle,const char* pStringXML);
XC_API HXCGUI WINAPI XListView_GetTemplateObject(HELE hEle,int iGroup,int iItem,int nTempItemID);
XC_API HXCGUI WINAPI XListView_GetTemplateObjectGroup(HELE hEle,int iGroup,int nTempItemID);
XC_API BOOL WINAPI XListView_GetItemIDFromHXCGUI(HELE hEle, HXCGUI hXCGUI,out_ int* piGroup,out_ int* piItem);
XC_API BOOL WINAPI XListView_HitTest(HELE hEle,in_ POINT* pPt,out_ int* pOutGroup,out_ int* pOutItem);
XC_API int  WINAPI XListView_HitTestOffset(HELE hEle,in_ POINT *pPt,out_ int* pOutGroup,out_ int* pOutItem); //自动添加滚动视图偏移量
XC_API void WINAPI XListView_EnableMultiSel(HELE hEle,BOOL bEnable);
XC_API void WINAPI XListView_SetDrawItemBkFlags(HELE hEle, int nFlags);
XC_API BOOL WINAPI XListView_SetSelectItem(HELE hEle,int iGroup,int iItem);
XC_API BOOL WINAPI XListView_GetSelectItem(HELE hEle,int *piGroup,int *piItem);
XC_API int  WINAPI XListView_GetSelectItemCount(HELE hEle);
XC_API int  WINAPI XListView_GetSelectItemAll(HELE hEle,out_ listView_item_id_i *pArray,int nArraySize);
XC_API void WINAPI XListView_SetSelectItemAll(HELE hEle);
XC_API void WINAPI XListView_CancelSelectItemAll(HELE hEle);
XC_API void WINAPI XListView_SetColumnSpace(HELE hEle,int space); //设置列间隔
XC_API void WINAPI XListView_SetRowSpace(HELE hEle,int space); //设置行间隔
XC_API void WINAPI XListView_SetAlignSizeLeft(HELE hEle,int size);
XC_API void WINAPI XListView_SetAlignSizeTop(HELE hEle,int size);
XC_API void WINAPI XListView_SetItemSize(HELE hEle,int width,int height);
XC_API void WINAPI XListView_GetItemSize(HELE hEle,SIZE* pSize);
XC_API void WINAPI XListView_SetGroupHeight(HELE hEle,int height);
XC_API int  WINAPI XListView_GetGroupHeight(HELE hEle);
XC_API void WINAPI XListView_AddItemBkBorder(HELE hEle,list_item_state_ nState, COLORREF color,BYTE alpha, int width);
XC_API void WINAPI XListView_AddItemBkFill(HELE hEle,list_item_state_ nState, COLORREF color,BYTE alpha);
XC_API void WINAPI XListView_AddItemBkImage(HELE hEle,list_item_state_ nState, HIMAGE hImage);
XC_API int  WINAPI XListView_GetItemBkInfoCount(HELE hEle, list_item_state_ nState);
XC_API void WINAPI XListView_ClearItemBkInfo(HELE hEle,list_item_state_ nState);
XC_API HBKINFOM WINAPI XListView_GetItemBkInfoManager(HELE hEle,list_item_state_ nState);
XC_API BOOL WINAPI XListView_ExpandGroup(HELE hEle,int iGroup,BOOL bExpand);
XC_API HELE WINAPI XMenuBar_Create(int x, int y, int cx, int cy, HXCGUI hParent = NULL);
XC_API int  WINAPI XMenuBar_AddButton(HELE hEle, wchar_t *pText);     //增加弹出菜单按钮
XC_API void WINAPI XMenuBar_SetButtonHeight(HELE hEle,int height);
XC_API HMENUX WINAPI XMenuBar_GetMenu(HELE hEle, int nIndex);
XC_API BOOL WINAPI XMenuBar_DeleteButton(HELE hEle,int nIndex); //删除菜单项并且销毁,同时该按钮下的弹出菜单也被销毁
XC_API HMENUX WINAPI XMenu_Create();
XC_API void WINAPI XMenu_AddItem(HMENUX hMenu, int nID,const wchar_t *pText, int parentId = XC_ID_ROOT, int nFlags = 0); //添加菜单项
XC_API void WINAPI XMenu_AddItemIcon(HMENUX hMenu, int nID,const wchar_t *pText,int nParentID,HIMAGE hImage,int nFlags=0);
XC_API void WINAPI XMenu_InsertItem(HMENUX hMenu, int nID, const wchar_t *pText, int nFlags, int insertID);
XC_API void WINAPI XMenu_InsertItemIcon(HMENUX hMenu, int nID, const wchar_t *pText, HIMAGE hIcon, int nFlags, int insertID);
XC_API void WINAPI XMenu_SetAutoDestroy(HMENUX hMenu, BOOL bAuto); //是否自动销毁对象,默认弹出菜单关闭后自动销毁
XC_API void WINAPI XMenu_EnableDrawBackground(HMENUX hMenu,BOOL bEnable); //是否有用户绘制菜单背景
XC_API void WINAPI XMenu_EnableDrawItem(HMENUX hMenu, BOOL bEnable);
XC_API BOOL WINAPI XMenu_Popup(HMENUX hMenu, HWND hParentWnd, int x, int y, HELE hParentEle=NULL, menu_popup_position_ nPosition=menu_popup_position_left_top); //弹出菜单
XC_API void WINAPI XMenu_DestroyMenu(HMENUX hMenu);
XC_API void WINAPI XMenu_CloseMenu(HMENUX hMenu);
XC_API void WINAPI XMenu_SetBkImage(HMENUX hMenu, HIMAGE hImage);
XC_API BOOL WINAPI XMenu_SetItemText(HMENUX hMenu,int nID,const wchar_t *pText); //设置项文本
XC_API BOOL WINAPI XMenu_GetItemText(HMENUX hMenu, int nID,out_ wchar_t *pOut, int nOutLen);
XC_API int  WINAPI XMenu_GetItemTextLength(HMENUX hMenu, int nID);
XC_API BOOL WINAPI XMenu_SetItemIcon(HMENUX hMenu, int nID, HIMAGE hIcon);
XC_API BOOL WINAPI XMenu_SetItemFlags(HMENUX hMenu,int nID,int uFlags);   //设置项属性
XC_API void WINAPI XMenu_SetItemHeight(HMENUX hMenu,int height); //设置菜单项高度
XC_API int  WINAPI XMenu_GetItemHeight(HMENUX hMenu);
XC_API void WINAPI XMenu_SetBorderColor(HMENUX hMenu, COLORREF crColor,BYTE alpha=255);
XC_API int  WINAPI XMenu_GetLeftWidth(HMENUX hMenu); //获取左侧宽度
XC_API int  WINAPI XMenu_GetLeftSpaceText(HMENUX hMenu);  //获取菜单项文本左间隔
XC_API int  WINAPI XMenu_GetItemCount(HMENUX hMenu); //获取菜单项数量,包含子菜单项
XC_API BOOL WINAPI XMenu_SetItemCheck(HMENUX hMenu, int nID,BOOL bCheck);
XC_API BOOL WINAPI XMenu_IsItemCheck(HMENUX hMenu, int nID);
XC_API HWINDOW WINAPI XModalWnd_Create(int nWidth,int nHeight,wchar_t *pTitle,HWND hWndParent,int XCStyle=xc_window_style_modal);
XC_API HWINDOW WINAPI XModalWnd_CreateEx(DWORD dwExStyle, wchar_t* lpClassName, wchar_t* lpWindowName, DWORD dwStyle, int x, int y, int cx, int cy, HWND hWndParent, int XCStyle=xc_window_style_modal);  //创建窗口
XC_API void WINAPI XModalWnd_EnableAutoClose(HWINDOW hWindow,BOOL bEnable);
XC_API int  WINAPI XModalWnd_DoModal(HWINDOW hWindow);     //启动模态窗口
XC_API void WINAPI XModalWnd_EndModal(HWINDOW hWindow,int nResult); //终止
XC_API HELE WINAPI XPane_Create(const wchar_t *pName,int nWidth,int nHeight,HWINDOW hFrameWnd=NULL);
XC_API void WINAPI XPane_SetView(HELE hEle, HELE hView);
XC_API void WINAPI XPane_GetTitle(HELE hEle, out_ wchar_t* pOut,int nOutLen);
XC_API BOOL WINAPI XPane_IsShowPane(HELE hEle); //判断窗格是否隐藏
XC_API void  WINAPI XPane_HidePane(HELE hEle);     //隐藏窗格
XC_API void  WINAPI XPane_ShowPane(HELE hEle);     //隐藏-显示窗格
XC_API void  WINAPI XPane_DockPane(HELE hEle);      //停靠窗格,自动隐藏
XC_API void  WINAPI XPane_LockPane(HELE hEle);      //锁定窗格
XC_API void  WINAPI XPane_FloatPane(HELE hEle);     //弹出窗格,当在码头上
XC_API HELE WINAPI XProgBar_Create(int x, int y, int cx, int cy,HXCGUI hParent=NULL);
XC_API void WINAPI XProgBar_SetRange(HELE hEle,int range);
XC_API int WINAPI XProgBar_GetRange(HELE hEle);
XC_API void WINAPI XProgBar_SetSpaceTwo(HELE hEle,int leftSize,int rightSize);
XC_API void WINAPI XProgBar_SetPos(HELE hEle,int pos);
XC_API int  WINAPI XProgBar_GetPos(HELE hEle);
XC_API void WINAPI XProgBar_SetHorizon(HELE hEle,BOOL bHorizon);
XC_API void WINAPI XProgBar_SetImageLoad(HELE hEle,HIMAGE hImage);
XC_API HELE WINAPI XPGrid_Create(int x,int y,int cx,int cy,HXCGUI hParent=NULL);
XC_API int  WINAPI XPGrid_AddItem(HELE hEle,const wchar_t *pName,propertyGrid_item_type_ nType,int nParentID=XC_ID_ROOT); //添加项
XC_API int  WINAPI XPGrid_AddItemString(HELE hEle,const wchar_t *pName,const wchar_t *pValue,int nParentID=XC_ID_ROOT);
XC_API void WINAPI XPGrid_DeleteAll(HELE hEle);
XC_API HELE WINAPI XPGrid_GetItemHELE(HELE hEle,int nItemID);
XC_API void WINAPI XPGrid_SetWidth(HELE hEle,int nWidth);
XC_API BOOL WINAPI XPGrid_SetItemValue(HELE hEle,int nItemID,const wchar_t* pValue);
XC_API BOOL WINAPI XPGrid_SetItemValueInt(HELE hEle,int nItemID,int nValue);
XC_API const wchar_t* WINAPI XPGrid_GetItemValue(HELE hEle,int nItemID);
XC_API int  WINAPI XPGrid_HitTest(HELE hEle,POINT *pPt,BOOL *pbExpandButton);
XC_API int  WINAPI XPGrid_HitTestOffset(HELE hEle,POINT *pPt,BOOL *pbExpandButton);
XC_API BOOL WINAPI XPGrid_ExpandItem(HELE hEle,int nItemID,BOOL bExpand);//展开项
XC_API HELE WINAPI XRichEditColor_Create(int x, int y, int cx, int cy, HXCGUI hParent);
XC_API void WINAPI XRichEditColor_SetColor(HELE hEle,COLORREF color);//设置颜色
XC_API COLORREF WINAPI XRichEditColor_GetColor(HELE hEle); //获取颜色RGB值
XC_API HELE WINAPI XRichEditSet_Create(int x, int y, int cx, int cy, HXCGUI hParent);
XC_API HELE WINAPI XRichEditFile_Create(int x, int y, int cx, int cy, HXCGUI hParent = NULL);
XC_API void WINAPI XRichEditFile_SetOpenFileType(HELE hEle,const wchar_t *pType); //设置打开文件类型
XC_API void WINAPI XRichEditFile_SetDefaultFile(HELE hEle,const wchar_t *pFile); //设置默认目录文件
XC_API void WINAPI XRichEditFile_SetRelativeDir(HELE hEle,const wchar_t *pDir); //TODO:设置相对路径
XC_API HELE WINAPI XRichEditFolder_Create(int x, int y, int cx, int cy, HXCGUI hParent);
XC_API void WINAPI XRichEditFolder_SetDefaultDir(HELE hEle,const wchar_t* pDir);  //设置默认目录
XC_API HELE WINAPI XRichEdit_Create(int x, int y, int cx, int cy, HXCGUI hParent = NULL);
XC_API void WINAPI XRichEdit_InsertString(HELE hEle, const wchar_t *pString, const LOGFONTW *pFont=NULL,COLORREF color=0);
XC_API BOOL WINAPI XRichEdit_InsertImage(HELE hEle, HIMAGE hImage,const wchar_t *pImagePath);
XC_API BOOL WINAPI XRichEdit_InsertGif(HELE hEle, HIMAGE hImage,const wchar_t *pImagePath);
XC_API void WINAPI XRichEdit_InsertStringEx(HELE hEle, int iRow, int iColumn, const wchar_t *pString, const LOGFONTW *pFont=NULL,COLORREF color=0);
XC_API BOOL WINAPI XRichEdit_InsertImageEx(HELE hEle, int iRow,int iColumn,HIMAGE hImage,const wchar_t *pImagePath);
XC_API BOOL WINAPI XRichEdit_InsertGifEx(HELE hEle, int iRow,int iColumn,HIMAGE hImage,const wchar_t *pImagePath);
XC_API void WINAPI XRichEdit_EnableReadOnly(HELE hEle, BOOL bEnable);
XC_API void WINAPI XRichEdit_EnableMultiLine(HELE hEle, BOOL bEnable);
XC_API void WINAPI XRichEdit_EnablePassword(HELE hEle, BOOL bEnable);
XC_API void WINAPI XRichEdit_EnableEvent_XE_RICHEDIT_CHANGE(HELE hEle,BOOL bEnable);
XC_API void WINAPI XRichEdit_EnableAutoWrap(HELE hEle, BOOL bEnable);
XC_API void WINAPI XRichEdit_EnableAutoCancelSel(HELE hEle, BOOL bEnable);
XC_API void WINAPI XRichEdit_SetText(HELE hEle,const wchar_t* pString);
XC_API void WINAPI XRichEdit_SetTextInt(HELE hEle,int nVaule);
XC_API int  WINAPI XRichEdit_GetText(HELE hEle,out_ wchar_t* pOut,int len);
XC_API void  WINAPI XRichEdit_GetHTMLFormat(HELE hEle, out_ wchar_t* pOut, int len);  //获取HTML格式内容
XC_API void* WINAPI XRichEdit_GetData(HELE hEle,out_ int *pDataSize=NULL);
XC_API BOOL  WINAPI XRichEdit_InsertData(HELE hEle, void *pData, int iRow, int iColumn);
XC_API int  WINAPI XRichEdit_GetTextLength(HELE hEle);
XC_API void WINAPI XRichEdit_SetRowHeight(HELE hEle,UINT nHeight); //设置默认行高
XC_API int WINAPI XRichEdit_GetCurrentRow(HELE hEle);
XC_API int WINAPI XRichEdit_GetCurrentColumn(HELE hEle);
XC_API void WINAPI XRichEdit_SetCurrentPos(HELE hEle,int iRow,int iColumn);
XC_API int WINAPI XRichEdit_GetRowCount(HELE hEle);
XC_API int WINAPI XRichEdit_GetRowLength(HELE hEle,int iRow);
XC_API int WINAPI XRichEdit_GetSelectText(HELE hEle,out_ wchar_t* pOut,int len);
XC_API BOOL WINAPI XRichEdit_GetSelectPosition(HELE hEle,out_ Position *pBegin, out_ Position *pEnd); 
XC_API BOOL WINAPI XRichEdit_SetSelect(HELE hEle,int iStartRow,int iStartCol,int iEndRow,int iEndCol);
XC_API BOOL WINAPI XRichEdit_SetItemFontEx(HELE hEle,int beginRow,int beginColumn,int endRow,int endColumn,LOGFONTW *pFont);
XC_API BOOL WINAPI XRichEdit_SetItemColorEx(HELE hEle,int beginRow,int beginColumn,int endRow,int endColumn,COLORREF color,BYTE alpha=255);
XC_API void WINAPI XRichEdit_CancelSelect(HELE hEle);
XC_API void WINAPI XRichEdit_SetSelectBkColor(HELE hEle,COLORREF color,BYTE alpha=255);
XC_API BOOL WINAPI XRichEdit_IsEmpty(HELE hEle);
XC_API BOOL WINAPI XRichEdit_SelectAll(HELE hEle);    //选择所有内容
XC_API BOOL WINAPI XRichEdit_DeleteSelect(HELE hEle); //删除选择内容
XC_API void WINAPI XRichEdit_DeleteAll(HELE hEle);    //删除所有
XC_API BOOL WINAPI XRichEdit_ClipboardCut(HELE hEle);      //剪切
XC_API BOOL WINAPI XRichEdit_ClipboardCopy(HELE hEle);     //复制 选择的内容
XC_API BOOL WINAPI XRichEdit_ClipboardPaste(HELE hEle);    //粘贴 粘贴剪贴板内容
XC_API HELE WINAPI XSBar_Create(int x, int y, int cx, int cy, HXCGUI hParent = NULL);
XC_API void WINAPI XSBar_SetRange(HELE hEle, int range);
XC_API int  WINAPI XSBar_GetRange(HELE hEle);
XC_API void  WINAPI XSBar_ShowButton(HELE hEle, BOOL bShow);
XC_API void WINAPI XSBar_SetSliderLength(HELE hEle, int length); //设置滑块长度
XC_API void WINAPI XSBar_SetSliderMinLength(HELE hEle, int minLength); //设置滑块最小长度
XC_API BOOL WINAPI XSBar_SetHorizon(HELE hEle, BOOL bHorizon);  //水平显示或垂直
XC_API int  WINAPI XSBar_GetSliderMaxLength(HELE hEle);
XC_API BOOL WINAPI XSBar_ScrollUp(HELE hEle);       //向左滚动
XC_API BOOL WINAPI XSBar_ScrollDown(HELE hEle);     //向右滚动
XC_API BOOL WINAPI XSBar_ScrollTop(HELE hEle);      //滚动到顶部
XC_API BOOL WINAPI XSBar_ScrollBottom(HELE hEle);   //滚动到底部
XC_API BOOL WINAPI XSBar_ScrollPos(HELE hEle, int pos); //滚动到指定点
XC_API HELE WINAPI XSBar_GetButtonUp(HELE hEle);
XC_API HELE WINAPI XSBar_GetButtonDown(HELE hEle);
XC_API HELE WINAPI XSBar_GetButtonSlider(HELE hEle);
XC_API HELE WINAPI XSView_Create(int x, int y, int cx, int cy, HXCGUI hParent = NULL);
XC_API BOOL WINAPI XSView_SetTotalSize(HELE hEle, int cx, int cy); //设置视图内容总大小(不包含边框间隔)
XC_API void WINAPI XSView_GetTotalSize(HELE hEle, SIZE *pSize);  //获取设置的内容大小
XC_API BOOL WINAPI XSView_SetLineSize(HELE hEle,int nHeight,int nWidth);
XC_API void WINAPI XSView_GetLineSize(HELE hEle,out_ SIZE *pSize);
XC_API int  WINAPI XSView_GetViewPosH(HELE hEle);   //获取视口原点X坐标
XC_API int  WINAPI XSView_GetViewPosV(HELE hEle);   //获取视口原点Y坐标
XC_API int  WINAPI XSView_GetViewWidth(HELE hEle); 
XC_API int  WINAPI XSView_GetViewHeight(HELE hEle);
XC_API void WINAPI XSView_GetViewRect(HELE hEle, out_ RECT *pRect); //获取视口坐标(不含边框间隔),left为视口左间隔大小,top为视口上间隔大小,非(0,0).
XC_API HELE WINAPI XSView_GetScrollBarH(HELE hEle); //获取水平滚动条
XC_API HELE WINAPI XSView_GetScrollBarV(HELE hEle); //获取垂直滚动条
XC_API void WINAPI XSView_SetPadding(HELE hEle, int left,int top,int right,int bottom);
XC_API void WINAPI XSView_GetPadding(HELE hEle, out_ PaddingSize_* pPadding);
XC_API BOOL WINAPI XSView_ScrollPosH(HELE hEle, int pos); //滚动到指定点
XC_API BOOL WINAPI XSView_ScrollPosV(HELE hEle, int pos); //滚动到指定点
XC_API BOOL WINAPI XSView_ScrollPosXH(HELE hEle, int posX); //滚动到指定坐标
XC_API BOOL WINAPI XSView_ScrollPosYV(HELE hEle, int posY); //滚动到指定坐标
XC_API void WINAPI XSView_ShowSBarH(HELE hEle,BOOL bShow); //启用水平滚动条
XC_API void WINAPI XSView_ShowSBarV(HELE hEle,BOOL bShow); //启用垂直滚动条
XC_API void WINAPI XSView_EnableAutoShowScrollBar(HELE hEle,BOOL bEnable);	
XC_API BOOL WINAPI XSView_ScrollLeftLine(HELE hEle);    //向左滚动
XC_API BOOL WINAPI XSView_ScrollRightLine(HELE hEle);   //向右滚动
XC_API BOOL WINAPI XSView_ScrollTopLine(HELE hEle);     //向上滚动
XC_API BOOL WINAPI XSView_ScrollBottomLine(HELE hEle);    //向下滚动
XC_API BOOL WINAPI XSView_ScrollLeft(HELE hEle);    //水平滚动到左侧
XC_API BOOL WINAPI XSView_ScrollRight(HELE hEle);   //水平滚动到右侧
XC_API BOOL WINAPI XSView_ScrollTop(HELE hEle);     //滚动到顶部
XC_API BOOL WINAPI XSView_ScrollBottom(HELE hEle);  //滚动到底部
XC_API HELE WINAPI XShape_GetParentEle(HXCGUI hShape);
XC_API HXCGUI WINAPI XShape_GetParentLayout(HXCGUI hShape);
XC_API HWINDOW WINAPI XShape_GetHWINDOW(HXCGUI hShape);
XC_API HXCGUI  WINAPI XShape_GetParent(HXCGUI hShape);
XC_API void WINAPI XShape_SetID(HXCGUI hShape,int nID);
XC_API int  WINAPI XShape_GetID(HXCGUI hShape);
XC_API void WINAPI XShape_Redraw(HXCGUI hShape);
XC_API int  WINAPI XShape_GetWidth(HXCGUI hShape);
XC_API int  WINAPI XShape_GetHeight(HXCGUI hShape);
XC_API void WINAPI XShape_GetRect(HXCGUI hShape,out_ RECT* pRect);
XC_API void WINAPI XShape_SetRect(HXCGUI hShape,in_  RECT* pRect);
XC_API void WINAPI XShape_GetContentSize(HXCGUI hShape,out_ SIZE* pSize);//获取内容大小
XC_API void WINAPI XShape_ShowLayout(HXCGUI hShape,BOOL bShow);
XC_API void WINAPI XShape_AdjustLayout(HXCGUI hShape);
XC_API void WINAPI XShape_Destroy(HXCGUI hShape);
XC_API HXCGUI WINAPI XShapeText_Create(int x, int y, int cx, int cy, const wchar_t *pName,HXCGUI hParent=NULL);
XC_API void  WINAPI XShapeText_SetText(HXCGUI hTextBlock,const wchar_t *pName);
XC_API void  WINAPI XShapeText_GetText(HXCGUI hTextBlock,out_ wchar_t *pOut,int nOutLen);
XC_API int  WINAPI XShapeText_GetTextLength(HXCGUI hTextBlock);
XC_API void  WINAPI XShapeText_SetFont(HXCGUI hTextBlock,HFONTX hFontx);
XC_API void WINAPI XShapeText_SetTextColor(HXCGUI hTextBlock, COLORREF color,BYTE alpha);
XC_API COLORREF WINAPI XShapeText_GetTextColor(HXCGUI hTextBlock);
XC_API void WINAPI XShapeText_SetTextAlign(HXCGUI hTextBlock,int align);
XC_API void WINAPI XShapeText_SetOffset(HXCGUI hTextBlock,int x,int y);
XC_API void WINAPI XShapeText_SetLayoutWidth(HXCGUI hTextBlock,layout_size_type_ nType,int width);
XC_API void WINAPI XShapeText_SetLayoutHeight(HXCGUI hTextBlock,layout_size_type_ nType,int height);
XC_API void WINAPI XShapeText_GetLayoutWidth(HXCGUI hTextBlock,out_ layout_size_type_ *pType,out_ int *pWidth);
XC_API void WINAPI XShapeText_GetLayoutHeight(HXCGUI hTextBlock,out_ layout_size_type_ *pType,out_ int *pHeight);
XC_API HXCGUI WINAPI XShapePic_Create(int x,int y,int cx,int cy, HXCGUI hParent=NULL);
XC_API void WINAPI   XShapePic_SetImage(HXCGUI hShape,HIMAGE hImage);
XC_API void WINAPI XShapePic_SetLayoutWidth(HXCGUI hShape,layout_size_type_ nType,int width);
XC_API void WINAPI XShapePic_SetLayoutHeight(HXCGUI hShape,layout_size_type_ nType,int height);
XC_API void WINAPI XShapePic_GetLayoutWidth(HXCGUI hShape,out_ layout_size_type_ *pType,out_ int *pWidth);
XC_API void WINAPI XShapePic_GetLayoutHeight(HXCGUI hShape,out_ layout_size_type_ *pType,out_ int *pHeight);
XC_API HXCGUI WINAPI XShapeGif_Create(int x,int y,int cx,int cy, HXCGUI hParent=NULL);
XC_API void WINAPI XShapeGif_SetImage(HXCGUI hShape,HIMAGE hImage);
XC_API void WINAPI XShapeGif_SetLayoutWidth(HXCGUI hShape,layout_size_type_ nType,int width);
XC_API void WINAPI XShapeGif_SetLayoutHeight(HXCGUI hShape,layout_size_type_ nType,int height);
XC_API void WINAPI XShapeGif_GetLayoutWidth(HXCGUI hShape,out_ layout_size_type_ *pType,out_ int *pWidth);
XC_API void WINAPI XShapeGif_GetLayoutHeight(HXCGUI hShape,out_ layout_size_type_ *pType,out_ int *pHeight);
XC_API HXCGUI WINAPI XShapeRect_Create(int x,int y,int cx,int cy, HXCGUI hParent=NULL);
XC_API void WINAPI XShapeRect_SetBorderColor(HXCGUI hShape,COLORREF color,BYTE alpha=255);
XC_API void WINAPI XShapeRect_SetFillColor(HXCGUI hShape,COLORREF color,BYTE alpha=255);
XC_API void WINAPI XShapeRect_SetRoundAngle(HXCGUI hShape,int nWidth, int nHeight);
XC_API void WINAPI XShapeRect_GetRoundAngle(HXCGUI hShape,out_ int *pWidth,out_ int *pHeight);
XC_API void WINAPI XShapeRect_EnableBorder(HXCGUI hShape,BOOL bEnable);
XC_API void WINAPI XShapeRect_EnableFill(HXCGUI hShape,BOOL bEnable);
XC_API void WINAPI XShapeRect_EnableRoundAngle(HXCGUI hShape,BOOL bEnable);
XC_API HXCGUI WINAPI XShapeEllipse_Create(int x,int y,int cx,int cy, HXCGUI hParent=NULL);
XC_API void WINAPI XShapeEllipse_SetBorderColor(HXCGUI hShape,COLORREF color,BYTE alpha=255);
XC_API void WINAPI XShapeEllipse_SetFillColor(HXCGUI hShape,COLORREF color,BYTE alpha=255);
XC_API void WINAPI XShapeEllipse_EnableBorder(HXCGUI hShape,BOOL bEnable);
XC_API void WINAPI XShapeEllipse_EnableFill(HXCGUI hShape,BOOL bEnable);
XC_API HXCGUI WINAPI XShapeGroupBox_Create(int x,int y,int cx,int cy,const wchar_t* pName,HXCGUI hParent=NULL);
XC_API void WINAPI XShapeGroupBox_SetBorderColor(HXCGUI hShape,COLORREF color,BYTE alpha=255);
XC_API void WINAPI XShapeGroupBox_SetTextColor(HXCGUI hShape,COLORREF color,BYTE alpha=254);
XC_API void WINAPI XShapeGroupBox_SetFontX(HXCGUI hShape,HFONTX hFontX);
XC_API void WINAPI XShapeGroupBox_SetTextOffset(HXCGUI hShape,int offsetX,int offsetY);
XC_API void WINAPI XShapeGroupBox_SetRoundAngle(HXCGUI hShape,int nWidth, int nHeight);
XC_API void WINAPI XShapeGroupBox_SetText(HXCGUI hShape,const wchar_t* pText);
XC_API void WINAPI XShapeGroupBox_GetTextOffset(HXCGUI hShape,out_ int *pOffsetX,out_ int *pOffsetY);
XC_API void WINAPI XShapeGroupBox_GetRoundAngle(HXCGUI hShape,out_ int *pWidth,out_ int *pHeight);
XC_API void WINAPI XShapeGroupBox_EnableRoundAngle(HXCGUI hShape,BOOL bEnable);
XC_API HXCGUI WINAPI XShapeLine_Create(int x1,int y1,int x2,int y2, HXCGUI hParent=NULL);
XC_API void WINAPI XShapeLine_SetPosition(HXCGUI hShape,int x1,int y1,int x2,int y2);
XC_API void WINAPI XShapeLine_SetColor(HXCGUI hShape,COLORREF color,BYTE alpha);
XC_API HELE WINAPI XSliderBar_Create(int x, int y, int cx, int cy,HXCGUI hParent=NULL);
XC_API void WINAPI XSliderBar_SetRange(HELE hEle,int range);
XC_API int  WINAPI XSliderBar_GetRange(HELE hEle);
XC_API void WINAPI XSliderBar_SetButtonWidth(HELE hEle,int width);
XC_API void WINAPI XSliderBar_SetButtonHeight(HELE hEle,int height);
XC_API void WINAPI XSliderBar_SetSpaceTwo(HELE hEle,int leftSize,int rightSize);
XC_API void WINAPI XSliderBar_SetPos(HELE hEle,int pos);
XC_API int  WINAPI XSliderBar_GetPos(HELE hEle);
XC_API HELE WINAPI XSliderBar_GetButton(HELE hEle);
XC_API void WINAPI XSliderBar_SetHorizon(HELE hEle,BOOL bHorizon);
XC_API void WINAPI XSliderBar_SetImageLoad(HELE hEle,HIMAGE hImage);
XC_API HELE WINAPI XTabBar_Create(int x, int y, int cx, int cy, HXCGUI hParent = NULL);
XC_API int  WINAPI XTabBar_AddLabel(HELE hEle,const wchar_t *pName); //添加一个标签
XC_API int  WINAPI XTabBar_InsertLabel(HELE hEle,int index,wchar_t *pName);
XC_API BOOL WINAPI XTabBar_DeleteLabel(HELE hEle,int index);  //删除
XC_API void WINAPI XTabBar_DeleteLabelAll(HELE hEle);  //删除所有TAB
XC_API HELE WINAPI XTabBar_GetLabel(HELE hEle,int index);//获取标签按钮Button
XC_API HELE WINAPI XTabBar_GetLabelClose(HELE hEle,int index);
XC_API HELE WINAPI XTabBar_GetButtonLeft(HELE hEle);
XC_API HELE WINAPI XTabBar_GetButtonRight(HELE hEle);
XC_API int  WINAPI XTabBar_GetSelect(HELE hEle);  //获取选择的标签
XC_API int  WINAPI XTabBar_GetLabelSpacing(HELE hEle); //获取标签间距, 0没有间距
XC_API int  WINAPI XTabBar_GetLabelCount(HELE hEle); //获取标签项数量
XC_API void WINAPI XTabBar_SetLabelSpacing(HELE hEle,int spacing);//设置标签间距, 0没有间距
XC_API void WINAPI XTabBar_SetSelect(HELE hEle,int index);   //设置选择标签
XC_API void WINAPI XTabBar_SetUp(HELE hEle);    //左滚动
XC_API void WINAPI XTabBar_SetDown(HELE hEle);  //右滚动
XC_API void WINAPI XTabBar_EnableTile(HELE hEle, BOOL bTile);  //平铺标签,每个标签显示相同大小
XC_API void WINAPI XTabBar_EnableClose(HELE hEle, BOOL bEnable);
XC_API void WINAPI XTabBar_SetCloseSize(HELE hEle, SIZE *pSize);
XC_API void WINAPI XTabBar_SetTurnButtonSize(HELE hEle, SIZE *pSize);
XC_API BOOL WINAPI XTabBar_ShowLabel(HELE hEle,int index, BOOL bShow); //显示或隐藏指定项
XC_API HELE WINAPI XTextLink_Create(int x, int y, int cx, int cy,const wchar_t* pName, HXCGUI hParent = NULL);
XC_API void WINAPI XTextLink_EnableUnderlineLeave(HELE hEle,BOOL bEnable);
XC_API void WINAPI XTextLink_EnableUnderlineStay(HELE hEle,BOOL bEnable);
XC_API void WINAPI XTextLink_SetTextColorStay(HELE hEle,COLORREF color,BYTE alpha);
XC_API void WINAPI XTextLink_SetUnderlineColorLeave(HELE hEle,COLORREF color,BYTE alpha);
XC_API void WINAPI XTextLink_SetUnderlineColorStay(HELE hEle,COLORREF color,BYTE alpha);
XC_API HELE WINAPI XToolBar_Create(int x, int y, int cx, int cy,HXCGUI hParent=NULL);
XC_API int WINAPI XToolBar_InsertEle(HELE hEle,HELE hNewEle,int index=-1); //添加元素到工具条, -1插入到末尾
XC_API int WINAPI XToolBar_InsertSeparator(HELE hEle,int index=-1); //添加分隔符, -1插入到末尾
XC_API void WINAPI XToolBar_EnableButtonMenu(HELE hEle,BOOL bEnable);  //显示隐藏项菜单按钮
XC_API HELE WINAPI XToolBar_GetHEle(HELE hEle,int index);  //获取工具条上元素句柄
XC_API HELE WINAPI XToolBar_GetButtonLeft(HELE hEle);   //获取工具条左滚动按钮Button
XC_API HELE WINAPI XToolBar_GetButtonRight(HELE hEle);  //获取工具条右滚动按钮Button
XC_API HELE WINAPI XToolBar_GetButtonMenu(HELE hEle);   //获取弹出菜单按钮
XC_API void WINAPI XToolBar_SetSpace(HELE hEle,int nSize);
XC_API void WINAPI XToolBar_DeleteEle(HELE hEle,int index); //移除工具条上元素并销毁
XC_API void WINAPI XToolBar_DeleteAllEle(HELE hEle);
XC_API HELE WINAPI XTree_Create(int x,int y,int cx,int cy,HXCGUI hParent=NULL);
XC_API BOOL WINAPI XTree_SetItemTemplateXML(HELE hEle,const wchar_t* pXmlFile);
XC_API BOOL WINAPI XTree_SetItemTemplateXMLSel(HELE hEle,const wchar_t* pXmlFile);
XC_API BOOL WINAPI XTree_SetItemTemplateXMLFromString(HELE hEle,const char* pStringXML);
XC_API BOOL WINAPI XTree_SetItemTemplateXMLSelFromString(HELE hEle,const char* pStringXML);
XC_API void WINAPI XTree_SetDrawItemBkFlags(HELE hEle,int nFlags);
XC_API BOOL WINAPI XTree_SetItemData(HELE hEle,int nID,int nUserData);
XC_API int  WINAPI XTree_GetItemData(HELE hEle,int nID);
XC_API BOOL WINAPI XTree_SetSelectItem(HELE hEle,int nID);
XC_API int  WINAPI XTree_GetSelectItem(HELE hEle);
XC_API BOOL WINAPI XTree_IsExpand (HELE hEle,int nID);
XC_API BOOL WINAPI XTree_ExpandItem(HELE hEle,int nID,BOOL bExpand);
XC_API int  WINAPI XTree_HitTest(HELE hEle,POINT *pPt);
XC_API int  WINAPI XTree_HitTestOffet(HELE hEle,POINT *pPt);
XC_API int  WINAPI XTree_GetFirstChildItem(HELE hEle,int nID);
XC_API int  WINAPI XTree_GetNextSiblingItem(HELE hEle,int nID);
XC_API int  WINAPI XTree_GetParentItem(HELE hEle,int nID);
XC_API void WINAPI XTree_SetIndentation(HELE hEle,int nWidth);
XC_API int  WINAPI XTree_GetIndentation(HELE hEle);
XC_API void WINAPI XTree_SetItemHeight(HELE hEle,int nID,int nHeight,int nSelHeight);
XC_API void WINAPI XTree_GetItemHeight(HELE hEle,int nID,out_ int *pHeight,out_ int *pSelHeight);
XC_API void WINAPI XTree_AddItemBkBorder(HELE hEle,tree_item_state_ nState, COLORREF color,BYTE alpha, int width);
XC_API void WINAPI XTree_AddItemBkFill(HELE hEle,tree_item_state_ nState, COLORREF color,BYTE alpha);
XC_API void WINAPI XTree_AddItemBkImage(HELE hEle,tree_item_state_ nState, HIMAGE hImage);
XC_API int  WINAPI XTree_GetItemBkInfoCount(HELE hEle,tree_item_state_ nState);
XC_API void WINAPI XTree_ClearItemBkInfo(HELE hEle,tree_item_state_ nState);
XC_API HBKINFOM WINAPI XTree_GetItemBkInfoManager(HELE hEle, tree_item_state_ nState);
XC_API void WINAPI XTree_SetItemHeightDefault(HELE hEle,int nHeight,int nSelHeight);
XC_API void WINAPI XTree_GetItemHeightDefault(HELE hEle,out_ int *pHeight,out_ int *pSelHeight);
XC_API HXCGUI WINAPI XTree_GetTemplateObject(HELE hEle,int nID,int nTempItemID); //通过模板项ID,获取实例化模板项ID对应的对象.
XC_API int    WINAPI XTree_GetItemIDFromHXCGUI(HELE hEle, HXCGUI hXCGUI);
XC_API void WINAPI XTree_BindAdapter(HELE hEle,HXCGUI hAdapter);
XC_API HXCGUI WINAPI XTree_GetAdapter(HELE hEle);
XC_API void WINAPI XWnd_RegEventEx(HWINDOW hWindow, CEventBase* pEvent);
XC_API BOOL WINAPI XWnd_RemoveEventEx(HWINDOW hWindow,CEventBase *pEvent);
XC_API void WINAPI XWnd_RegEventC(HWINDOW hWindow, int nEvent, void *pFun);
XC_API void WINAPI XWnd_RegEventC1(HWINDOW hWindow, int nEvent, void *pFun);
XC_API BOOL WINAPI XWnd_RemovegEvent(HWINDOW hWindow, int nEvent, void *pFun);
XC_API HWINDOW WINAPI XWnd_Create(int x, int y, int cx, int cy, const wchar_t *pTitle, HWND hWndParent = NULL, int XCStyle = xc_window_style_default);
XC_API HWINDOW WINAPI XWnd_CreateEx(DWORD dwExStyle, const wchar_t* lpClassName, const wchar_t* lpWindowName, DWORD dwStyle, int x, int y, int cx, int cy, HWND hWndParent=NULL, int XCStyle=xc_window_style_default);
XC_API BOOL WINAPI XWnd_AddEle(HWINDOW hWindow, HELE hEle);
XC_API BOOL WINAPI XWnd_InsertEle(HWINDOW hWindow, HELE hChildEle,HELE hDestEle);
XC_API BOOL WINAPI XWnd_AddShape(HWINDOW hWindow, HXCGUI hShape);
XC_API HWND WINAPI XWnd_GetHWND(HWINDOW hWindow);
XC_API void WINAPI XWnd_EnableDragBorder(HWINDOW hWindow, BOOL bEnable);
XC_API void WINAPI XWnd_EnableDragWindow(HWINDOW hWindow, BOOL bEnable);
XC_API void WINAPI XWnd_EnableDrawBk(HWINDOW hWindow, BOOL bEnable);
XC_API void WINAPI XWnd_EnableAutoFocus(HWINDOW hWindow, BOOL bEnable); //当鼠标左键按下是否获得焦点
XC_API void WINAPI XWnd_EnableMaxWindow(HWINDOW hWindow, BOOL bEnable);
XC_API void WINAPI XWnd_RedrawWnd(HWINDOW hWindow);
XC_API void WINAPI XWnd_RedrawWndRect(HWINDOW hWindow, RECT *pRect, BOOL bImmediately = FALSE); //重绘窗口指定区域
XC_API void WINAPI XWnd_SetFocusEle(HWINDOW hWindow, HELE hFocusEle);
XC_API HELE WINAPI XWnd_GetFocusEle(HWINDOW hWindow);
XC_API void WINAPI XWnd_SetCursor(HWINDOW hWindow,HCURSOR hCursor);
XC_API HCURSOR WINAPI XWnd_GetCursor(HWINDOW hWindow);
XC_API HCURSOR WINAPI XWnd_SetCursorSys(HWINDOW hWindow,HCURSOR hCursor);
XC_API void WINAPI XWnd_SetFont(HWINDOW hWindow,HFONTX hFontx);
XC_API void WINAPI XWnd_SetID(HWINDOW hWindow, int nID);
XC_API int  WINAPI XWnd_GetID(HWINDOW hWindow);
XC_API void WINAPI XWnd_SetCaptureEle(HWINDOW hWindow, HELE hEle);
XC_API HELE WINAPI XWnd_GetCaptureEle(HWINDOW hWindow);
XC_API BOOL WINAPI XWnd_BindLayoutEle(HWINDOW hWindow,window_position_ nPosition, HELE hEle);
XC_API HELE WINAPI XWnd_GetLayoutEle(HWINDOW hWindow,window_position_ nPosition);
XC_API void WINAPI XWnd_BindLayoutObject(HWINDOW hWindow,window_position_ nPosition,HXCGUI hLayout);
XC_API HXCGUI WINAPI XWnd_GetLayoutObject(HWINDOW hWindow,window_position_ nPosition);
XC_API void WINAPI XWnd_SetLayoutSize(HWINDOW hWindow,int left, int top, int right, int bottom);
XC_API void WINAPI XWnd_GetLayoutSize(HWINDOW hWindow,out_ BorderSize_  *pBorderSize);
XC_API void WINAPI XWnd_SetDragBorderSize(HWINDOW hWindow,int left, int top, int right, int bottom); //设置拖动边框大小
XC_API void WINAPI XWnd_GetDragBorderSize(HWINDOW hWindow,out_ BorderSize_* pSize);
XC_API void WINAPI XWnd_SetMinimumSize(HWINDOW hWindow,int width,int height);
XC_API HELE WINAPI XWnd_HitChildEle(HWINDOW hWindow,POINT *pPt);
XC_API int  WINAPI XWnd_GetChildCount(HWINDOW hWindow);
XC_API HELE WINAPI XWnd_GetChildByIndex(HWINDOW hWindow,int index);
XC_API HELE WINAPI XWnd_GetChildByID(HWINDOW hWindow,int nID);
XC_API HELE WINAPI XWnd_GetEle(HWINDOW hWindow,int nID);
XC_API int  WINAPI XWnd_GetChildShapeCount(HWINDOW hWindow);
XC_API HXCGUI WINAPI XWnd_GetChildShapeByIndex(HWINDOW hWindow, int index);
XC_API void WINAPI XWnd_GetDrawRect(HWINDOW hWindow,RECT *pRcPaint);
XC_API BOOL WINAPI XWnd_ShowWindow(HWINDOW hWindow, int nCmdShow);
XC_API void WINAPI XWnd_AdjustLayout(HWINDOW hWindow);
XC_API void WINAPI XWnd_CloseWindow(HWINDOW hWindow);
XC_API void WINAPI XWnd_CreateCaret(HWINDOW hWindow, HELE hEle, int width, int height);//创建插入符
XC_API void WINAPI XWnd_SetCaretSize(HWINDOW hWindow, int width, int height); //设置插入符大小
XC_API void WINAPI XWnd_SetCaretPos(HWINDOW hWindow, int x, int y); //设置插入符位置
XC_API void WINAPI XWnd_SetCaretPosEx(HWINDOW hWindow, int x, int y, int width, int height); //设置插入符位置
XC_API void WINAPI XWnd_SetCaretColor(HWINDOW hWindow, COLORREF color); //设置插入符颜色
XC_API void WINAPI XWnd_ShowCaret(HWINDOW hWindow, BOOL bShow);  //显示插入符
XC_API void WINAPI XWnd_DestroyCaret(HWINDOW hWindow);    //销毁插入符
XC_API HELE WINAPI XWnd_GetCaretHELE(HWINDOW hWindow);
XC_API BOOL WINAPI XWnd_GetClientRect(HWINDOW hWindow, out_ RECT *pRect); //获取客户区坐标
XC_API void WINAPI XWnd_GetBodyRect(HWINDOW hWindow, out_ RECT *pRect);  //获取窗口body坐标
XC_API UINT WINAPI XWnd_SetTimer(HWINDOW hWindow, UINT nIDEvent, UINT uElapse); //设置定时器
XC_API BOOL WINAPI XWnd_KillTimer(HWINDOW hWindow, UINT nIDEvent);
XC_API HBKINFOM WINAPI XWnd_GetBkInfoManager(HWINDOW hWindow, window_position_ nPosition);
XC_API void WINAPI XWnd_SetTransparentType(HWINDOW hWindow,window_transparent_ nType);  //设置透明窗口
XC_API void WINAPI XWnd_SetTransparentAlpha(HWINDOW hWindow,BYTE alpha); //设置窗口透明度
XC_API void WINAPI XWnd_SetTransparentColor(HWINDOW hWindow,COLORREF color); //设置窗口透明色
XC_API BOOL WINAPI XWnd_RegEventTest(int nEvent,BYTE* pParamType);
XC_API void WINAPI XWnd_AdjustLayoutObject(HWINDOW hWindow);
XC_API HXCGUI WINAPI XEle_GetParentLayoutObject(HELE hEle);
