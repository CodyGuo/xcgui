package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xFont_Create            *syscall.Proc
	xFont_Create2           *syscall.Proc
	xFont_CreateEx          *syscall.Proc
	xFont_CreateFromHFONT   *syscall.Proc
	xFont_EnableAutoDestroy *syscall.Proc
	xFont_Destroy           *syscall.Proc
	xFont_GetHFONT          *syscall.Proc
)

func init() {
	// Functions
	xFont_Create = xcDLL.MustFindProc("XFont_Create")
	xFont_Create2 = xcDLL.MustFindProc("XFont_Create2")
	xFont_CreateEx = xcDLL.MustFindProc("XFont_CreateEx")
	xFont_CreateFromHFONT = xcDLL.MustFindProc("XFont_CreateFromHFONT")
	xFont_EnableAutoDestroy = xcDLL.MustFindProc("XFont_EnableAutoDestroy")
	xFont_Destroy = xcDLL.MustFindProc("XFont_Destroy")
	xFont_GetHFONT = xcDLL.MustFindProc("XFont_GetHFONT")
}

/*
创建炫彩字体,当字体句柄与元素关联后,会自动释放.

参数:
	size 字体大小.
返回:
	字体句柄.
*/
func XFont_Create(size int) HFONTX {
	ret, _, _ := xFont_Create.Call(uintptr(size))

	return HFONTX(ret)
}

/*
创建炫彩字体.

参数:
	pName 字体名称.*uint16
	size 字体大小.
	bBold 是否为粗体.
	bItalic 是否为斜体.
	bUnderline 是否有下划线.
	bStrikeOut 是否有删除线.
返回:
	字体句柄.
*/
func XFont_Create2(pName string, size int, bBold, bItalic, bUnderline, bStrikeOut bool) HFONTX {
	ret, _, _ := xFont_Create2.Call(
		StringToUintPtr(pName),
		// uintptr(unsafe.Pointer(pName)),
		uintptr(size),
		uintptr(BoolToBOOL(bBold)),
		uintptr(BoolToBOOL(bItalic)),
		uintptr(BoolToBOOL(bUnderline)),
		uintptr(BoolToBOOL(bStrikeOut)))

	return HFONTX(ret)
}

/*
创建炫彩字体.

参数:
	pFontInfo 字体信息.
返回:
	字体句柄.
*/
func XFont_CreateEx(pFontInfo *LOGFONTW) HFONTX {
	ret, _, _ := xFont_CreateEx.Call(uintptr(unsafe.Pointer(pFontInfo)))

	return HFONTX(ret)
}

/*
创建炫彩字体从现有HFONT字体.

参数:
	hFont 字体句柄.
返回:
	返回炫彩字体.
*/
func XFont_CreateFromHFONT(hFont HFONT) HFONTX {
	ret, _, _ := xFont_CreateFromHFONT.Call(uintptr(hFont))

	return HFONTX(ret)
}

/*
是否自动销毁.

参数:
	hFontX 字体句柄.
	bEnable 是否启用.
*/
func XFont_EnableAutoDestroy(hFontX HFONTX, bEnable bool) {
	xFont_EnableAutoDestroy.Call(
		uintptr(hFontX),
		uintptr(BoolToBOOL(bEnable)))
}

/*
销毁炫彩字体.

参数:
	hFontX 字体句柄.
*/
func XFont_Destroy(hFontX HFONTX) {
	xFont_Destroy.Call(uintptr(hFontX))
}

/*
获取字体HFONT句柄..

参数:
	hFontX 字体句柄.
返回:
	返回HFONT句柄.
*/
func XFont_GetHFONT(hFontX HFONTX) HFONT {
	ret, _, _ := xFont_GetHFONT.Call(uintptr(hFontX))

	return HFONT(ret)
}
