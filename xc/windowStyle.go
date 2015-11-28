package xc

// 窗口默认大小
const CW_USEDEFAULT = ^0x7fffffff

// type xc_window_style_ uint32

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
	XC_WINDOW_STYLE_NOTHING     = 0x00000000
	XC_WINDOW_STYLE_CAPTION     = 0x00000001
	XC_WINDOW_STYLE_BORDER      = 0x00000002
	XC_WINDOW_STYLE_CENTER      = 0x00000004
	XC_WINDOW_STYLE_DRAG_BORDER = 0x00000008
	XC_WINDOW_STYLE_DRAG_WINDOW = 0x00000010
	XC_WINDOW_STYLE_DEFAULT     = 0x00000001 | 0x00000002 | 0x00000004 | 0x00000008
	XC_WINDOW_STYLE_MODAL       = 0x00000001 | 0x00000004 | 0x00000008
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
