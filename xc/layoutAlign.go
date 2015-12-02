package xc

// 对齐方式
type Layout_align_ int

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
	Layout_align_left Layout_align_ = iota
	Layout_align_top
	Layout_align_right
	Layout_align_bottom
	Layout_align_center
	Layout_align_equidistant
)
