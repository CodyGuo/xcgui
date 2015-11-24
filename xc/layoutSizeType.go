package xc

type LAYOUT_SIZE_TYPE_ uint32

/* 布局大小类型
layout_size_type_fixed 指定大小
layout_size_type_fill fill 填充父
layout_size_type_auto auto 自动大小,根据内容计算大小
layout_size_type_weight weight 比例,按照比例分配剩余空间
layout_size_type_disable disable 不使用
*/
const (
	LAYOUT_SIZE_TYPE_FIXED LAYOUT_SIZE_TYPE_ = iota
	LAYOUT_SIZE_TYPE_FILL
	LAYOUT_SIZE_TYPE_AUTO
	LAYOUT_SIZE_TYPE_WEIGHT
	LAYOUT_SIZE_TYPE_DISABLE
)
