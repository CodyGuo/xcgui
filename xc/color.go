package xc

type COLORREF uint32

func RGB(r, g, b byte) COLORREF {
    return COLORREF(uint32(r) | uint32(g)<<8 | uint32(b)<<16)
}

func (c COLORREF) R() byte {
    return byte(c & 0xff)
}

func (c COLORREF) G() byte {
    return byte((c >> 8) & 0xff)
}

func (c COLORREF) B() byte {
    return byte((c >> 16) & 0xff)
}
