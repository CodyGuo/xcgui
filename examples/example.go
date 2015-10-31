package main

import (
    . "github.com/codyguo/xcgui"
)

func main() {
    xui := XCGUI(0, 0, 600, 400, "Hey Man", 0, 0x8|0x5)
    xui.Show(0x5)
    xui.LoopMessage()
}
