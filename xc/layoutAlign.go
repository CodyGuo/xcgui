package xc

// 对齐方式
type Layout_align_ uint32

/*
对齐方式
	layout_align_left 左对齐
	layout_align_top 上对齐
	layout_align_right 右对齐
	layout_align_bottom 下对齐
	layout_align_center 居中对齐
	layout_align_equidistant  等距
*/
const (
	LAYOUT_ALIGN_LEFT Layout_align_ = iota
	LAYOUT_ALIGN_TOP
	LAYOUT_ALIGN_RIGHT
	LAYOUT_ALIGN_BOTTOM
	LAYOUT_ALIGN_CENTER
	LAYOUT_ALIGN_EQUIDISTANT
)
