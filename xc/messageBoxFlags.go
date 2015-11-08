package xc

type messageBox_flags_ uint32

/* 弹出消息框
messageBox_flags_other 其他
messageBox_flags_ok 确定按钮
messageBox_flags_cancel 取消按钮
*/
const (
    MESSAGEBOX_FLAGS_OTHER  messageBox_flags_ = 0x00
    MESSAGEBOX_FLAGS_OK     messageBox_flags_ = 0x01
    MESSAGEBOX_FLAGS_CANCEL messageBox_flags_ = 0x02
)
