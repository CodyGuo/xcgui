package xc

import (
	"syscall"
	"unsafe"
)

//模板项
type template_info_i struct {
	info int
}

//c++
//typedef struct {
//        lfHeight;
//        lfWidth;
//        lfEscapement;
//        lfOrientation;
//        lfWeight;
//  BYTE  lfItalic;
//  BYTE  lfUnderline;
//  BYTE  lfStrikeOut;
//  BYTE  lfCharSet;
//  BYTE  lfOutPrecision;
//  BYTE  lfClipPrecision;
//  BYTE  lfQuality;
//  BYTE  lfPitchAndFamily;
//  WCHAR lfFaceName[LF_FACESIZE];
//} LOGFONTW;

const LF_FACESIZE = 32

type LOGFONTW struct {
	lfHeight         int32
	lfWidth          int32
	lfEscapement     int32
	lfOrientation    int32
	lfWeight         int32
	lfItalic         byte
	lfUnderline      byte
	lfStrikeOut      byte
	lfCharSet        byte
	lfOutPrecision   byte
	lfClipPrecision  byte
	lfQuality        byte
	lfPitchAndFamily byte
	lfFaceName       [LF_FACESIZE]uint16
}

var (
	xc_LoadLayout             *syscall.Proc
	xc_LoadLayoutFromString   *syscall.Proc
	xc_LoadResource           *syscall.Proc
	xc_LoadResourceFromString *syscall.Proc
	xc_LoadTemplate           *syscall.Proc
	xc_LoadTemplateFromString *syscall.Proc
	xc_TemplateDestroy        *syscall.Proc
	xc_GetDefaultFont         *syscall.Proc
	xc_SetDefaultFont         *syscall.Proc
	xc_InitFont               *syscall.Proc
	xc_Malloc                 *syscall.Proc
	xc_Free                   *syscall.Proc
)

func init() {
	xc_LoadLayout = xcDLL.MustFindProc("XC_LoadLayout")
	xc_LoadLayoutFromString = xcDLL.MustFindProc("XC_LoadLayoutFromString")
	xc_LoadResource = xcDLL.MustFindProc("XC_LoadResource")
	xc_LoadResourceFromString = xcDLL.MustFindProc("XC_LoadResourceFromString")
	xc_LoadTemplate = xcDLL.MustFindProc("XC_LoadTemplate")
	xc_LoadTemplateFromString = xcDLL.MustFindProc("XC_LoadTemplateFromString")
	xc_TemplateDestroy = xcDLL.MustFindProc("XC_TemplateDestroy")
	xc_GetDefaultFont = xcDLL.MustFindProc("XC_GetDefaultFont")
	xc_SetDefaultFont = xcDLL.MustFindProc("XC_SetDefaultFont")
	xc_InitFont = xcDLL.MustFindProc("XC_InitFont")
	xc_Malloc = xcDLL.MustFindProc("XC_Malloc")
	xc_Free = xcDLL.MustFindProc("XC_Free")
}

/*
加载布局文件.

参数:
	pFileName 布局文件名.
	hParent 父对象,窗口句柄或UI元素句柄.
返回:
	返回窗口句柄或布局句柄或元素句柄.
*/
func XC_LoadLayout(pFileName string, hParent HXCGUI) HXCGUI {
	ret, _, _ := xc_LoadLayout.Call(
		StringToUintPtr(pFileName),
		uintptr(hParent))

	return HXCGUI(ret)
}

/*
加载布局文件从内存字符串.

参数:
	pStringXML 字符串指针.
	hParent 父对象,窗口句柄或UI元素句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XC_LoadLayoutFromString(pStringXML string, hParent HXCGUI) HXCGUI {
	ret, _, _ := xc_LoadLayoutFromString.Call(
		StringToUintPtr(pStringXML),
		uintptr(hParent))

	return HXCGUI(ret)
}

/*
加载资源文件.

参数:
	pFileName 资源文件名.
	pDir 指定资源目录,可选参数,如果未指定那么为程序当前目录.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XC_LoadResource(pFileName, pDir string) bool {
	ret, _, _ := xc_LoadResource.Call(
		StringToUintPtr(pFileName),
		StringToUintPtr(pDir))

	return ret == TRUE
}

/*
加载资源文件从内存字符串.

参数:
	pStringXML 字符串指针.
	pDir 指定资源目录,可选参数,如果未指定那么为程序当前目录.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XC_LoadResourceFromString(pStringXML, pDir string) bool {
	ret, _, _ := xc_LoadResourceFromString.Call(
		StringToUintPtr(pStringXML),
		StringToUintPtr(pDir))

	return ret == TRUE
}

/*
项模板文件载入.

参数:
	nType 模板类型,例如 XC_LIST , XC_LISTBOX.
	pFileName 文件名.
返回:
	返回模板信息.
*/
func XC_LoadTemplate(nType XC_OBJECT_TYPE, pFileName string) *template_info_i {
	ret, _, _ := xc_LoadTemplate.Call(
		uintptr(nType),
		StringToUintPtr(pFileName))

	return (*template_info_i)(unsafe.Pointer(ret))
}

/*
加载模板文件从内存字符串.

参数:
	nType 模板类型,例如 XC_LIST , XC_LISTBOX..
	pStringXML 字符串指针.
返回:
	返回模板信息.
*/
func XC_LoadTemplateFromString(nType XC_OBJECT_TYPE, pStringXML string) *template_info_i {
	ret, _, _ := xc_LoadTemplate.Call(
		uintptr(nType),
		StringToUintPtr(pStringXML))

	return (*template_info_i)(unsafe.Pointer(ret))
}

/*
项模板信息销毁.

参数:
	pInfo 项模板信息.
*/
func XC_TemplateDestroy(pInfo *template_info_i) {
	xc_TemplateDestroy.Call(uintptr(unsafe.Pointer(pInfo)))
}

/*
获取默认字体.

返回:
	返回默认字体句柄.
*/
func XC_GetDefaultFont() HFONTX {
	ret, _, _ := xc_GetDefaultFont.Call()

	return HFONTX(ret)
}

/*
设置默认字体.

参数:
	hFontX 炫彩字体句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XC_SetDefaultFont(hFontX HFONTX) {
	xc_SetDefaultFont.Call(uintptr(hFontX))
}

/*
初始化LOGFONTW结构体.

参数:
	pFont LOGFONTW结构体指针.
	pName 字体名称.
	size 字体大小.
	bBold 是否为粗体.
	bItalic 是否为斜体.
	bUnderline 是否有下划线.
	bStrikeOut 是否有删除线.
*/
func XC_InitFont(pFont *LOGFONTW, pName string, size int, bBold, bItalic, bUnderline, bStrikeOut bool) {
	xc_InitFont.Call(
		uintptr(unsafe.Pointer(pFont)),
		StringToUintPtr(pName),
		uintptr(size),
		uintptr(BoolToBOOL(bBold)),
		uintptr(BoolToBOOL(bItalic)),
		uintptr(BoolToBOOL(bItalic)),
		uintptr(BoolToBOOL(bStrikeOut)))
}

/*
在UI库中申请内存.

参数:
	size 大小,字节为单位.
返回:
	内存首地址.
*/
func XC_Malloc(size int) uintptr {
	ret, _, _ := xc_Malloc.Call(uintptr(size))

	return ret
}

/*
在UI库中释放内存.

参数:
	p 内存首地址.
*/
func XC_Free(p uintptr) {
	xc_Free.Call(p)
}
