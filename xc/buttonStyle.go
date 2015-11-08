package xc

type button_style_ uint32

/* 按钮样式(用于区分外观)
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
    BUTTON_STYLE_DEFAULT button_style_ = iota
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
