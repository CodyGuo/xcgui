package xc

type MENU_STATE_FLAGS_ uint32

/* 弹出菜单项标识
menu_state_flags_normal 正常
menu_state_flags_select 选择
menu_state_flags_check 勾选
menu_state_flags_popup 弹出
menu_state_flags_separator 分隔栏 ID号任意,ID号被忽略
menu_state_flags_disable 禁用
*/
const (
	MENU_STATE_FLAGS_NORMAL    MENU_STATE_FLAGS_ = 0x00
	MENU_STATE_FLAGS_SELECT    MENU_STATE_FLAGS_ = 0x01
	MENU_STATE_FLAGS_CHECK     MENU_STATE_FLAGS_ = 0x02
	MENU_STATE_FLAGS_POPUP     MENU_STATE_FLAGS_ = 0x04
	MENU_STATE_FLAGS_SEPARATOR MENU_STATE_FLAGS_ = 0x08
	MENU_STATE_FLAGS_DISABLE   MENU_STATE_FLAGS_ = 0x10
)

type MENU_POPUP_POSITION_ uint32

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
	MENU_POPUP_POSITION_LEFT_TOP MENU_POPUP_POSITION_ = iota
	MENU_POPUP_POSITION_LEFT_BOTTOM
	MENU_POPUP_POSITION_RIGHT_TOP
	MENU_POPUP_POSITION_RIGHT_BOTTOM
	MENU_POPUP_POSITION_CENTER_LEFT
	MENU_POPUP_POSITION_CENTER_TOP
	MENU_POPUP_POSITION_CENTER_RIGHT
	MENU_POPUP_POSITION_CENTER_BOTTOM
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
