package xc

type image_draw_type_ uint32

/* 图片绘制类型
image_draw_type_default 默认
image_draw_type_stretch 拉伸
image_draw_type_adaptive 自适应
image_draw_type_tile 平铺
*/
const (
	IMAGE_DRAW_TYPE_DEFAULT image_draw_type_ = iota
	IMAGE_DRAW_TYPE_STRETCH
	IMAGE_DRAW_TYPE_ADAPTIVE
	IMAGE_DRAW_TYPE_TILE
)
