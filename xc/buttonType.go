package xc

type BUTTON_TYPE_ uint32

/* button_type_ 按钮类型(用于区分功能)
   button_type_default 默认类型
   button_type_check 复选按钮
   button_type_radio 单选按钮
   button_type_close 窗口关闭按钮
   button_type_min 窗口最小化按钮
   button_type_max 窗口最大化还原按钮
*/
const (
    BUTTON_TYPE_DEFAULT BUTTON_TYPE_ = iota
    BUTTON_TYPE_CHECK
    BUTTON_TYPE_RADIO
    BUTTON_TYPE_CLOSE
    BUTTON_TYPE_MIN
    BUTTON_TYPE_MAX
)
