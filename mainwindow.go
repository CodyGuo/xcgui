package xcgui

import (
    "github.com/codyguo/xcgui/xc"
)

var (
    winStyle uint32 = xc.XC_WINDOW_STYLE_DEFAULT
)

type MainWindow struct {
    WindowBase
}

func NewMainWindow(width, height int, title string) (*MainWindow, error) {
    mw := new(MainWindow)

    if err := InitWindow(
        mw,
        nil,
        width,
        height,
        title,
        winStyle); err != nil {
        return nil, err
    }

    succeeded := false

    go func() {
        if !succeeded {
            newErrorNoPanic("NewMainWindow")
        }
    }()

    succeeded = true

    return mw, nil
}
