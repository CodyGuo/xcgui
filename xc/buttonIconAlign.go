package xc

type BUTTON_ICON_ALIGN_ uint32

/* 按钮图标对齐方式
button_icon_align_left 图标在左边
button_icon_align_top 图标在顶部
button_icon_align_right 图标在右边
button_icon_align_bottom 图标在底部
*/
const (
	BUTTON_ICON_ALIGN_LEFT BUTTON_ICON_ALIGN_ = iota
	BUTTON_ICON_ALIGN_TOP
	BUTTON_ICON_ALIGN_RIGHT
	BUTTON_ICON_ALIGN_BOTTOM
)
