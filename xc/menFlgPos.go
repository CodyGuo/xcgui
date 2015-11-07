package xc

type menu_state_flags uint32

/* 弹出菜单项标识
   menu_state_flags_normal 正常
   menu_state_flags_select 选择
   menu_state_flags_check 勾选
   menu_state_flags_popup 弹出
   menu_state_flags_separator 分隔栏 ID号任意,ID号被忽略
   menu_state_flags_disable 禁用
*/
const (
    menu_state_flags_normal    menu_state_flags = 0x00
    menu_state_flags_select    menu_state_flags = 0x01
    menu_state_flags_check     menu_state_flags = 0x02
    menu_state_flags_popup     menu_state_flags = 0x04
    menu_state_flags_separator menu_state_flags = 0x08
    menu_state_flags_disable   menu_state_flags = 0x10
)

type menu_popup_position uint32

/* 弹出菜单方向
   menu_popup_position_left_top 左上角
   menu_popup_position_left_bottom 左下角
   menu_popup_position_right_top 右上角
   menu_popup_position_right_bottom 右下角
   menu_popup_position_center_left 左居中
   menu_popup_position_center_top 上居中
   menu_popup_position_center_right 右居中
   menu_popup_position_center_bottom 下居中
*/
const (
    menu_popup_position_left_top menu_popup_position = iota
    menu_popup_position_left_bottom
    menu_popup_position_right_top
    menu_popup_position_right_bottom
    menu_popup_position_center_left
    menu_popup_position_center_top
    menu_popup_position_center_right
    menu_popup_position_center_bottom
)

/* 菜单ID , 当前未使用
   IDM_CLIP 剪切
   IDM_COPY 复制
   IDM_PASTE 粘贴
   IDM_DELETE 删除
   IDM_SELECTALL 全选
   IDM_DELETEALL 清空
*/
const (
    IDM_CLIP      = 1000000000
    IDM_COPY      = 1000000001
    IDM_PASTE     = 1000000002
    IDM_DELETE    = 1000000003
    IDM_SELECTALL = 1000000004
    IDM_DELETEALL = 1000000005
)
