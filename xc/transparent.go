package xc

type window_transparent_ int32

/* 炫彩窗口透明标识
   window_transparent_false 默认窗口,不透明
   window_transparent_shaped 透明窗口,带透明通道,异型
   window_transparent_shadow 阴影窗口,带透明通道,边框阴影,窗口透明或半透明,当前未启用.
   window_transparent_simple 透明窗口,不带透明通道,指定半透明度,指定透明色.
   window_transparent_win7 WIN7玻璃窗口,需要WIN7开启特效,当前未启用.
*/
const (
    window_transparent_false window_transparent_ = iota
    window_transparent_shaped
    window_transparent_shadow
    window_transparent_simple
    window_transparent_win7
)
