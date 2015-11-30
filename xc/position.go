package xc

type WINDOW_POSITION_ int32

/*
窗口位置
	window_position_error 错误
	window_position_top top
	window_position_bottom bottom
	window_position_left left
	window_position_right right
	window_position_body body
	window_position_window window 整个窗口
*/
const (
	WINDOW_POSITION_ERROR WINDOW_POSITION_ = -1
	WINDOW_POSITION_TOP   WINDOW_POSITION_ = iota + WINDOW_POSITION_ERROR
	WINDOW_POSITION_BOTTOM
	WINDOW_POSITION_LEFT
	WINDOW_POSITION_RIGHT
	WINDOW_POSITION_BODY
	WINDOW_POSITION_WINDOW
)

type Position struct {
	iRow    int
	iColumn int
}
