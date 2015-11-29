package xc

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
)
