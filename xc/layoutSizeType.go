package xc

type layout_size_type_ uint32

/* 布局大小类型
   layout_size_type_fixed 指定大小
   layout_size_type_fill fill 填充父
   layout_size_type_auto auto 自动大小,根据内容计算大小
   layout_size_type_weight weight 比例,按照比例分配剩余空间
   layout_size_type_disable disable 不使用
*/
const (
    layout_size_type_fixed layout_size_type_ = iota
    layout_size_type_fill
    layout_size_type_auto
    layout_size_type_weight
    layout_size_type_disable
)
