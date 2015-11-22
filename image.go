package xcgui

import (
    // "fmt"
    "path/filepath"
    "strings"
)
import (
    "github.com/codyguo/xcgui/xc"
)

type Image interface {
    Size() Size
}

func NewImageFromFile(filePath string) (xc.HIMAGE, error) {
    if strings.HasSuffix(filePath, ".ico") {
        return NewIconFromFile(filePath)
    }

    return NewIconFromFile(filePath)
}

func NewIconFromFile(filePath string) (hIcon xc.HIMAGE, err error) {
    absFilePath, err := filepath.Abs(filePath)
    if err != nil {
        return 0, wrapError(err)
    }

    hIcon = xc.HIMAGE(xc.XImageLoadFile(
        absFilePath,
        false))

    if hIcon == 0 {
        return 0, lastError("LoadImage")
    }

    return hIcon, nil
}
