package xc

type button_icon_align_ uint32

/* 按钮图标对齐方式
   button_icon_align_left 图标在左边
   button_icon_align_top 图标在顶部
   button_icon_align_right 图标在右边
   button_icon_align_bottom 图标在底部
*/
const (
    button_icon_align_left button_icon_align_ = iota
    button_icon_align_top
    button_icon_align_right
    button_icon_align_bottom
)
