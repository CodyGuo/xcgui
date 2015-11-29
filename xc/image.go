package xc

import (
	"syscall"
	"unsafe"
)

var (
	xImage_LoadFile                *syscall.Proc
	xImage_LoadFileAdaptive        *syscall.Proc
	xImage_LoadFileRect            *syscall.Proc
	xImage_LoadResAdaptive         *syscall.Proc
	xImage_LoadRes                 *syscall.Proc
	xImage_LoadZip                 *syscall.Proc
	xImage_LoadZipAdaptive         *syscall.Proc
	xImage_LoadZipRect             *syscall.Proc
	xImage_LoadMemory              *syscall.Proc
	xImage_LoadMemoryRect          *syscall.Proc
	xImage_LoadMemoryAdaptive      *syscall.Proc
	xImage_LoadFromImage           *syscall.Proc
	xImage_LoadFileFromExtractIcon *syscall.Proc
	xImage_LoadFileFromHICON       *syscall.Proc
	xImage_LoadFileFromHBITMAP     *syscall.Proc
	xImage_IsStretch               *syscall.Proc
	xImage_IsAdaptive              *syscall.Proc
	xImage_IsTile                  *syscall.Proc
	xImage_SetDrawType             *syscall.Proc
	xImage_SetDrawTypeAdaptive     *syscall.Proc
	xImage_SetTranColor            *syscall.Proc
	xImage_SetTranColorEx          *syscall.Proc
	xImage_EnableTranColor         *syscall.Proc
	xImage_EnableAutoDestroy       *syscall.Proc
	xImage_EnableCenter            *syscall.Proc
	xImage_IsCenter                *syscall.Proc
	xImage_GetDrawType             *syscall.Proc
	xImage_GetWidth                *syscall.Proc
	xImage_GetHeight               *syscall.Proc
	xImage_Release                 *syscall.Proc
	xImage_Destroy                 *syscall.Proc
)

func init() {
	xImage_LoadFile = XCDLL.MustFindProc("XImage_LoadFile")
	xImage_LoadFileAdaptive = XCDLL.MustFindProc("XImage_LoadFileAdaptive")
	xImage_LoadFileRect = XCDLL.MustFindProc("XImage_LoadFileRect")
	xImage_LoadResAdaptive = XCDLL.MustFindProc("XImage_LoadResAdaptive")
	xImage_LoadRes = XCDLL.MustFindProc("XImage_LoadRes")
	xImage_LoadZip = XCDLL.MustFindProc("XImage_LoadZip")
	xImage_LoadZipAdaptive = XCDLL.MustFindProc("XImage_LoadZipAdaptive")
	xImage_LoadZipRect = XCDLL.MustFindProc("XImage_LoadZipRect")
	xImage_LoadMemory = XCDLL.MustFindProc("XImage_LoadMemory")
	xImage_LoadMemoryRect = XCDLL.MustFindProc("XImage_LoadMemoryRect")
	xImage_LoadMemoryAdaptive = XCDLL.MustFindProc("XImage_LoadMemoryAdaptive")
	xImage_LoadFromImage = XCDLL.MustFindProc("XImage_LoadFromImage")
	xImage_LoadFileFromExtractIcon = XCDLL.MustFindProc("XImage_LoadFileFromExtractIcon")
	xImage_LoadFileFromHICON = XCDLL.MustFindProc("XImage_LoadFileFromHICON")
	xImage_LoadFileFromHBITMAP = XCDLL.MustFindProc("XImage_LoadFileFromHBITMAP")
	xImage_IsStretch = XCDLL.MustFindProc("XImage_IsStretch")
	xImage_IsAdaptive = XCDLL.MustFindProc("XImage_IsAdaptive")
	xImage_IsTile = XCDLL.MustFindProc("XImage_IsTile")
	xImage_SetDrawType = XCDLL.MustFindProc("XImage_SetDrawType")
	xImage_SetDrawTypeAdaptive = XCDLL.MustFindProc("XImage_SetDrawTypeAdaptive")
	xImage_SetTranColor = XCDLL.MustFindProc("XImage_SetTranColor")
	xImage_SetTranColorEx = XCDLL.MustFindProc("XImage_SetTranColorEx")
	xImage_EnableTranColor = XCDLL.MustFindProc("XImage_EnableTranColor")
	xImage_EnableAutoDestroy = XCDLL.MustFindProc("XImage_EnableAutoDestroy")
	xImage_EnableCenter = XCDLL.MustFindProc("XImage_EnableCenter")
	xImage_IsCenter = XCDLL.MustFindProc("XImage_IsCenter")
	xImage_GetDrawType = XCDLL.MustFindProc("XImage_GetDrawType")
	xImage_GetWidth = XCDLL.MustFindProc("XImage_GetWidth")
	xImage_GetHeight = XCDLL.MustFindProc("XImage_GetHeight")
	xImage_Release = XCDLL.MustFindProc("XImage_Release")
	xImage_Destroy = XCDLL.MustFindProc("XImage_Destroy")
}

/*
加载图片从文件.

参数:
	pImageName 图片文件.
	bStretch 是否拉伸图片
返回:
	图片句柄.
*/
func XImageLoadFile(pImageName *uint16, bStretch bool) HIMAGE {
	ret, _, _ := xImage_LoadFile.Call(
		uintptr(unsafe.Pointer(pImageName)),
		uintptr(BoolToBOOL(bStretch)))

	return HIMAGE(ret)
}

/*
加载图片从文件,自适应图片.

参数:
	pImageName 图片文件.
	leftSize 坐标.
	topSize 坐标.
	rightSize 坐标.
	bottomSize 坐标.
返回:
	图片句柄.
*/
func XImageLoadFileAdaptive(pImageName *uint16, leftSize, topSize, rightSize, bottomSize int) HIMAGE {
	ret, _, _ := xImage_LoadFileAdaptive.Call(
		uintptr(unsafe.Pointer(pImageName)),
		uintptr(leftSize),
		uintptr(topSize),
		uintptr(rightSize),
		uintptr(bottomSize))

	return HIMAGE(ret)
}

/*
加载图片,指定区位置及大小.

参数:
	pImageName 图片文件.
	x 坐标.
	y 坐标.
	cx 宽度.
	cy 高度.
返回:
	图片句柄.
*/
func XImageLoadFileRect(pImageName *uint16, x, y, cx, cy int) HIMAGE {
	ret, _, _ := xImage_LoadFileRect.Call(
		uintptr(unsafe.Pointer(pImageName)),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy))

	return HIMAGE(ret)
}

/*
加载图片从资源,自适应图片.

参数:
	id 资源ID.
	pType 资源类型.
	leftSize 坐标.
	topSize 坐标.
	rightSize 坐标.
	bottomSize 坐标.
返回:
	图片句柄.
*/
func XImageLoadResAdaptive(id int, pType *uint16, leftSize, topSize, rightSize, bottomSize int) HIMAGE {
	ret, _, _ := xImage_LoadResAdaptive.Call(
		uintptr(id),
		uintptr(unsafe.Pointer(pType)),
		uintptr(leftSize),
		uintptr(topSize),
		uintptr(rightSize),
		uintptr(bottomSize))

	return HIMAGE(ret)
}

/*
加载图片从资源.

参数:
	id 资源ID.
	pType 资源类型.
	bStretch 是否拉伸图片
返回:
	图片句柄.
*/
func XImageLoadRes(id int, pType *uint16, bStretch bool) HIMAGE {
	ret, _, _ := xImage_LoadRes.Call(
		uintptr(id),
		uintptr(unsafe.Pointer(pType)),
		uintptr(BoolToBOOL(bStretch)))

	return HIMAGE(ret)
}

/*
加载图片从ZIP压缩包.

参数:
	pZipFileName ZIP压缩包文件名.
	pImageName 图片文件名.
	pPassword ZIP压缩包密码.
	bStretch 是否拉伸图片
返回:
	图片句柄.
*/
func XImageLoadZip(pZipFileName, pImageName, pPassword *uint16, bStretch bool) HIMAGE {
	ret, _, _ := xImage_LoadZip.Call(
		uintptr(unsafe.Pointer(pZipFileName)),
		uintptr(unsafe.Pointer(pImageName)),
		uintptr(unsafe.Pointer(pPassword)),
		uintptr(BoolToBOOL(bStretch)))

	return HIMAGE(ret)
}

/*
加载图片从ZIP压缩包,自适应图片.

参数:
	pZipFileName ZIP压缩包文件名.
	pImageName 图片文件名.
	pPassword ZIP压缩包密码,如果没有填NULL.
	x1 坐标.
	x2 坐标.
	y1 坐标.
	y2 坐标.
返回:
	图片句柄.
*/
func XImageLoadZipAdaptive(pZipFileName, pImageName, pPassword *uint16, x1, x2, y1, y2 int) HIMAGE {
	ret, _, _ := xImage_LoadZipAdaptive.Call(
		uintptr(unsafe.Pointer(pZipFileName)),
		uintptr(unsafe.Pointer(pImageName)),
		uintptr(unsafe.Pointer(pPassword)),
		uintptr(x1),
		uintptr(x2),
		uintptr(y1),
		uintptr(y2))

	return HIMAGE(ret)
}

/*
加载ZIP图片,指定区位置及大小.

参数:
	pZipFileName ZIP文件.
	pImageName 图片名称
	pPassword 密码
	x 坐标.
	y 坐标.
	cx 宽度.
	cy 高度.
返回:
	图片句柄.
*/
func XImageLoadZipRect(pZipFileName, pImageName, pPassword *uint16, x, y, cx, cy int) HIMAGE {
	ret, _, _ := xImage_LoadZipRect.Call(
		uintptr(unsafe.Pointer(pZipFileName)),
		uintptr(unsafe.Pointer(pImageName)),
		uintptr(unsafe.Pointer(pPassword)),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy))

	return HIMAGE(ret)
}

/*
加载流图片,指定区位置及大小.

参数:
	pBuffer 图片缓冲区
	nSize 图片缓冲区大小
	bStretch 是否拉伸图片
返回:
	图片句柄.
*/
func XImageLoadMemory(pBuffer uintptr, nSize int, bStretch bool) HIMAGE {
	ret, _, _ := xImage_LoadMemory.Call(
		pBuffer,
		uintptr(nSize),
		uintptr(BoolToBOOL(bStretch)))

	return HIMAGE(ret)
}

/*
加载流图片,指定区位置及大小.

参数:
	pBuffer 图片缓冲区
	nSize 图片缓冲区大小
	x 坐标.
	y 坐标.
	cx 宽度.
	cy 高度.
	bStretch 是否拉伸图片
返回:
	图片句柄.
*/
func XImageLoadMemoryRect(pBuffer uintptr, nSize int, x, y, cx, cy int, bStretch bool) HIMAGE {
	ret, _, _ := xImage_LoadMemoryRect.Call(
		pBuffer,
		uintptr(nSize),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(BoolToBOOL(bStretch)))

	return HIMAGE(ret)
}

/*
加载流图片压缩包,自适应图片.

参数:
	pBuffer 图片缓冲区
	nSize 图片缓冲区大小
	leftSize 坐标.
	topSize 坐标.
	rightSize 坐标.
	bottomSize 坐标.
返回:
	图片句柄.
*/
func XImageLoadMemoryAdaptive(pBuffer uintptr, nSize, leftSize, topSize, rightSize, bottomSize int) HIMAGE {
	ret, _, _ := xImage_LoadMemoryAdaptive.Call(
		pBuffer,
		uintptr(nSize),
		uintptr(leftSize),
		uintptr(topSize),
		uintptr(rightSize),
		uintptr(bottomSize))

	return HIMAGE(ret)
}

/*
加载图片从GDI+的Image对象.

参数:
	pImage GDI图片对象指针 Image*.
返回:
	成功返回炫彩图片句柄,失败返回FALSE.
*/
func XImageLoadFromImage(pImage uintptr) HIMAGE {
	ret, _, _ := xImage_LoadFromImage.Call(pImage)
	if ret == FALSE {
		return 0
	}

	return HIMAGE(ret)
}

/*
加载文件图标,从一个EXE文件或DLL文件或图标文件;例如:*.exe文件的图标.

参数:
	pImageName 文件名.
返回:
	成功返回炫彩图片句柄,失败返回FALSE.
*/
func XImageLoadFileFromExtractIcon(pImageName *uint16) HIMAGE {
	ret, _, _ := xImage_LoadFileFromExtractIcon.Call(uintptr(unsafe.Pointer(pImageName)))
	if ret == FALSE {
		return 0
	}

	return HIMAGE(ret)
}

/*
创建一个炫彩图片句柄,从一个现有的图标句柄HICON.

参数:
	hIcon 图标句柄,如果你不使用可以释放 DestroyIcon().
返回:
	成功返回炫彩图片句柄,失败返回FALSE.
*/
func XImageLoadFileFromHICON(hIcon HICON) HIMAGE {
	ret, _, _ := xImage_LoadFileFromHICON.Call(uintptr(hIcon))
	if ret == FALSE {
		return 0
	}

	return HIMAGE(ret)
}

/*
创建一个炫彩图片句柄,从一个现有的位图句柄HBITMAP.

参数:
	hBitmap 位图句柄,如果你不使用可以释放 DeleteObject().
返回:
	成功返回炫彩图片句柄,失败返回FALSE.
*/
func XImageLoadFileFromHBITMAP(hBitmap HBITMAP) HIMAGE {
	ret, _, _ := xImage_LoadFileFromHBITMAP.Call(uintptr(hBitmap))
	if ret == FALSE {
		return 0
	}

	return HIMAGE(ret)
}

/*
是否为拉伸图片句柄

参数:
	hImage 图片句柄.
返回:
	如果是返回TRUE,否则相反.
*/
func XImageIsStretch(hImage HIMAGE) bool {
	ret, _, _ := xImage_IsStretch.Call(uintptr(hImage))

	return ret == TRUE
}

/*
是否为自适应图片句柄

参数:
	hImage 图片句柄.
返回:
	如果是返回TRUE,否则相反.
*/
func XImageIsAdaptive(hImage HIMAGE) bool {
	ret, _, _ := xImage_IsAdaptive.Call(uintptr(hImage))

	return ret == TRUE
}

/*
是否为平铺图片

参数:
	hImage 图片句柄.
返回:
	如果是返回TRUE,否则相反.
*/
func XImageIsTile(hImage HIMAGE) bool {
	ret, _, _ := xImage_IsTile.Call(uintptr(hImage))

	return ret == TRUE
}

/*
设置图片绘制类型

参数:
	hImage 图片句柄.
	nType 图片绘制类型.
返回:
	如果是返回TRUE,否则相反.
*/
func XImageSetDrawType(hImage HIMAGE, nType IMAGE_DRAW_TYPE_) bool {
	ret, _, _ := xImage_SetDrawType.Call(
		uintptr(hImage),
		uintptr(nType))

	return ret == TRUE
}

/*
设置图片自适应

参数:
	hImage 图片句柄.
	leftSize 坐标.
	topSize 坐标.
	rightSize 坐标.
	bottomSize 坐标.
返回:
	如果是返回TRUE,否则相反.
*/
func XImageSetDrawTypeAdaptive(hImage HIMAGE, leftSize, topSize, rightSize, bottomSize int) bool {
	ret, _, _ := xImage_SetDrawTypeAdaptive.Call(
		uintptr(hImage),
		uintptr(leftSize),
		uintptr(topSize),
		uintptr(rightSize),
		uintptr(bottomSize))

	return ret == TRUE
}

/*
指定图片透明颜色.

参数:
	hImage 图片句柄.
	color RGB颜色.
*/
func XImageSetTranColor(hImage HIMAGE, color COLORREF) {
	xImage_SetTranColor.Call(
		uintptr(hImage),
		uintptr(color))
}

/*
指定图片透明颜色及透明度.

参数:
	hImage 图片句柄.
	color RGB颜色.
	tranColor 透明色的透明度.
*/
func XImageSetTranColorEx(hImage HIMAGE, color COLORREF, tranColor byte) {
	xImage_SetTranColorEx.Call(
		uintptr(hImage),
		uintptr(color),
		uintptr(tranColor))
}

/*
启用或关闭图片透明色.

参数:
	hImage 图片句柄.
	bEnable 启用TRUE,关闭FALSE.
*/
func XImageEnableTranColor(hImage HIMAGE, bEnable bool) {
	xImage_EnableTranColor.Call(
		uintptr(hImage),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用或关闭自动销毁,当与UI元素关联时有效

参数:
	hImage 图片句柄.
	bEnable 启用自动销毁TRUE,关闭自动销毁FALSE.
*/
func XImageEnableAutoDestroy(hImage HIMAGE, bEnable bool) {
	xImage_EnableAutoDestroy.Call(
		uintptr(hImage),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用或关闭图片居中显示，默认属性图片有效。

参数:
	hImage 图片句柄.
	bCenter 是否居中显示.
*/
func XImageEnableCenter(hImage HIMAGE, bCenter bool) {
	xImage_EnableCenter.Call(
		uintptr(hImage),
		uintptr(BoolToBOOL(bCenter)))
}

/*
判断图片是否居中显示

参数:
	hImage 图片句柄.
返回:
	如果居中显示返回TRUE，否则相反.
*/
func XImageIsCenter(hImage HIMAGE) bool {
	ret, _, _ := xImage_IsCenter.Call(uintptr(hImage))

	return ret == TRUE
}

/*
获取图片绘制类型

参数:
	hImage 图片句柄.
返回:
	图片绘制类型.
*/
func XImageGetDrawType(hImage HIMAGE) IMAGE_DRAW_TYPE_ {
	ret, _, _ := xImage_GetDrawType.Call(uintptr(hImage))

	return IMAGE_DRAW_TYPE_(ret)
}

/*
获取图片宽度.

参数:
	hImage 图片句柄.
返回:
	图片宽度.
*/
func XImageGetWidth(hImage HIMAGE) int {
	ret, _, _ := xImage_GetWidth.Call(uintptr(hImage))

	return int(ret)
}

/*
获取图片高度.

参数:
	hImage 图片句柄.
返回:
	图片高度.
*/
func XImageGetHeight(hImage HIMAGE) int {
	ret, _, _ := xImage_GetHeight.Call(uintptr(hImage))

	return int(ret)
}

/*
释放引用计数.

参数:
	hImage 图片句柄.
*/
func XImageRelease(hImage HIMAGE) {
	xImage_Release.Call(uintptr(hImage))
}

/*
销毁图片接口.

参数:
	hImage 图片句柄.
*/
func XImageDestroy(hImage HIMAGE) {
	xImage_Destroy.Call(uintptr(hImage))
}
