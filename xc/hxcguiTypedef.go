package xc

// HXCGUI: 其他句柄
type HXCGUI uintptr

// 句柄列表
type (
    // 系统窗口句柄
    HWND HXCGUI
    // 窗口句柄
    HWINDOW HXCGUI
    // UI元素句柄
    HELE HXCGUI

    HMENUX          HXCGUI
    HSTRING         HXCGUI
    HDRAW           HXCGUI
    HARRAY          HXCGUI
    HIMAGE          HXCGUI
    HXMLRES         HXCGUI
    HARRAY_ITERATOR HXCGUI
    HFONTX          HXCGUI
    HPANE_GROUP     HXCGUI
    HBKINFOM        HXCGUI
)
