package xc

type button_type_ uint32

/* 按钮类型(用于区分功能)
   button_type_default 默认类型
   button_type_check 复选按钮
   button_type_radio 单选按钮
   button_type_close 窗口关闭按钮
   button_type_min 窗口最小化按钮
   button_type_max 窗口最大化还原按钮
*/
const (
    button_type_default button_type_ = iota
    button_type_check
    button_type_radio
    button_type_close
    button_type_min
    button_type_max
)
