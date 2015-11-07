package xc

type xc_window_style_ uint32

/* 炫彩窗口样式
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
    XC_WINDOW_STYLE_NOTHING     xc_window_style_ = 0x00000000
    XC_WINDOW_STYLE_CAPTION     xc_window_style_ = 0x00000001
    XC_WINDOW_STYLE_BORDER      xc_window_style_ = 0x00000002
    XC_WINDOW_STYLE_CENTER      xc_window_style_ = 0x00000004
    XC_WINDOW_STYLE_DRAG_BORDER xc_window_style_ = 0x00000008
    XC_WINDOW_STYLE_DRAG_WINDOW xc_window_style_ = 0x00000010
    XC_WINDOW_STYLE_DEFAULT     xc_window_style_ = 0x00000001 | 0x00000002 | 0x00000004 | 0x00000008
    XC_WINDOW_STYLE_MODAL       xc_window_style_ = 0x00000001 | 0x00000004 | 0x00000008
)
