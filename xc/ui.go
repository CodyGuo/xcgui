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
	xc_InitFont = xcDLL.MustFindProc("XC_InitFont")
	xc_Malloc = xcDLL.MustFindProc("XC_Malloc")
	xc_Free = xcDLL.MustFindProc("XC_Free")
}

func XC_LoadLayout(pFileName string, hParent HXCGUI) HXCGUI {
	ret, _, _ := xc_LoadLayout.Call(
		StringToUintPtr(pFileName),
		uintptr(hParent))
	return HXCGUI(ret)
}

func XC_LoadLayoutFromString(pStringXML string, hParent HXCGUI) HXCGUI {
	ret, _, _ := xc_LoadLayoutFromString.Call(
		StringToUintPtr(pStringXML),
		uintptr(hParent))
	return HXCGUI(ret)
}

func XC_LoadResource(pFileName string, pDir string) bool {
	ret, _, _ := xc_LoadResource.Call(
		StringToUintPtr(pFileName),
		StringToUintPtr(pDir))
	return ret == TRUE
}

func XC_LoadResourceFromString(pStringXML string, pDir string) bool {
	ret, _, _ := xc_LoadResourceFromString.Call(
		StringToUintPtr(pStringXML),
		StringToUintPtr(pDir))
	return ret == TRUE
}

func XC_LoadTemplate(nType XC_OBJECT_TYPE, pFileName string) *template_info_i {
	ret, _, _ := xc_LoadTemplate.Call(
		uintptr(nType),
		StringToUintPtr(pFileName))
	return (*template_info_i)(unsafe.Pointer(ret))
}

func XC_LoadTemplateFromString(nType XC_OBJECT_TYPE, pStringXML string) *template_info_i {
	ret, _, _ := xc_LoadTemplate.Call(
		uintptr(nType),
		StringToUintPtr(pStringXML))

	return (*template_info_i)(unsafe.Pointer(ret))
}

func XC_TemplateDestroy(pInfo *template_info_i) {
	xc_TemplateDestroy.Call(uintptr(unsafe.Pointer(pInfo)))
}

func XC_GetDefaultFont() HFONTX {
	ret, _, _ := xc_GetDefaultFont.Call()
	return HFONTX(ret)
}

func XC_InitFont(pFont *LOGFONTW, pName string, size int, bBold bool, bItalic bool, bUnderline bool, bStrikeOut bool) {
	xc_InitFont.Call(
		uintptr(unsafe.Pointer(pFont)),
		StringToUintPtr(pName),
		uintptr(size),
		uintptr(BoolToBOOL(bBold)),
		uintptr(BoolToBOOL(bItalic)),
		uintptr(BoolToBOOL(bItalic)),
		uintptr(BoolToBOOL(bStrikeOut)))
}

func XC_Malloc(size int) {
	xc_Malloc.Call(uintptr(size))
}

func XC_Free(p *interface{}) {
	xc_Free.Call(uintptr(unsafe.Pointer(p)))
}
