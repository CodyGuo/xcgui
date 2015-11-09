package xcgui

type Size struct {
    Width, Height int
}

func minSize(a, b Size) Size {
    var s Size

    if a.Width < b.Width {
        s.Width = a.Width
    } else {
        s.Width = b.Width
    }

    if a.Height < b.Height {
        s.Height = a.Height
    } else {
        s.Height = b.Height
    }

    return s
}

func maxSize(a, b Size) Size {
    var s Size

    if a.Width > b.Width {
        s.Width = a.Width
    } else {
        s.Width = b.Width
    }

    if a.Height > b.Height {
        s.Height = a.Height
    } else {
        s.Height = b.Height
    }

    return s
}
