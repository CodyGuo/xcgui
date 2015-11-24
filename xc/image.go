package xc

import (
	"syscall"
)

var (
	XImage_LoadFile                *syscall.Proc
	XImage_LoadFileAdaptive        *syscall.Proc
	XImage_LoadFileRect            *syscall.Proc
	XImage_LoadResAdaptive         *syscall.Proc
	XImage_LoadRes                 *syscall.Proc
	XImage_LoadZip                 *syscall.Proc
	XImage_LoadZipAdaptive         *syscall.Proc
	XImage_LoadZipRect             *syscall.Proc
	XImage_LoadMemory              *syscall.Proc
	XImage_LoadMemoryRect          *syscall.Proc
	XImage_LoadMemoryAdaptive      *syscall.Proc
	XImage_LoadFromImage           *syscall.Proc
	XImage_LoadFileFromExtractIcon *syscall.Proc
	XImage_LoadFileFromHICON       *syscall.Proc
	XImage_LoadFileFromHBITMAP     *syscall.Proc
	XImage_IsStretch               *syscall.Proc
	XImage_IsAdaptive              *syscall.Proc
	XImage_IsTile                  *syscall.Proc
	XImage_SetDrawType             *syscall.Proc
	XImage_SetDrawTypeAdaptive     *syscall.Proc
	XImage_SetTranColor            *syscall.Proc
	XImage_SetTranColorEx          *syscall.Proc
	XImage_EnableTranColor         *syscall.Proc
	XImage_EnableAutoDestroy       *syscall.Proc
	XImage_EnableCenter            *syscall.Proc
	XImage_IsCenter                *syscall.Proc
	XImage_GetDrawType             *syscall.Proc
	XImage_GetWidth                *syscall.Proc
	XImage_GetHeight               *syscall.Proc
	XImage_Release                 *syscall.Proc
	XImage_Destroy                 *syscall.Proc
)

func init() {
	XImage_LoadFile = XCDLL.MustFindProc("XImage_LoadFile")
	XImage_LoadFileAdaptive = XCDLL.MustFindProc("XImage_LoadFileAdaptive")
	XImage_LoadFileRect = XCDLL.MustFindProc("XImage_LoadFileRect")
	XImage_LoadResAdaptive = XCDLL.MustFindProc("XImage_LoadResAdaptive")
	XImage_LoadRes = XCDLL.MustFindProc("XImage_LoadRes")
	XImage_LoadZip = XCDLL.MustFindProc("XImage_LoadZip")
	XImage_LoadZipAdaptive = XCDLL.MustFindProc("XImage_LoadZipAdaptive")
	XImage_LoadZipRect = XCDLL.MustFindProc("XImage_LoadZipRect")
	XImage_LoadMemory = XCDLL.MustFindProc("XImage_LoadMemory")
	XImage_LoadMemoryRect = XCDLL.MustFindProc("XImage_LoadMemoryRect")
	XImage_LoadMemoryAdaptive = XCDLL.MustFindProc("XImage_LoadMemoryAdaptive")
	XImage_LoadFromImage = XCDLL.MustFindProc("XImage_LoadFromImage")
	XImage_LoadFileFromExtractIcon = XCDLL.MustFindProc("XImage_LoadFileFromExtractIcon")
	XImage_LoadFileFromHICON = XCDLL.MustFindProc("XImage_LoadFileFromHICON")
	XImage_LoadFileFromHBITMAP = XCDLL.MustFindProc("XImage_LoadFileFromHBITMAP")
	XImage_IsStretch = XCDLL.MustFindProc("XImage_IsStretch")
	XImage_IsAdaptive = XCDLL.MustFindProc("XImage_IsAdaptive")
	XImage_IsTile = XCDLL.MustFindProc("XImage_IsTile")
	XImage_SetDrawType = XCDLL.MustFindProc("XImage_SetDrawType")
	XImage_SetDrawTypeAdaptive = XCDLL.MustFindProc("XImage_SetDrawTypeAdaptive")
	XImage_SetTranColor = XCDLL.MustFindProc("XImage_SetTranColor")
	XImage_SetTranColorEx = XCDLL.MustFindProc("XImage_SetTranColorEx")
	XImage_EnableTranColor = XCDLL.MustFindProc("XImage_EnableTranColor")
	XImage_EnableAutoDestroy = XCDLL.MustFindProc("XImage_EnableAutoDestroy")
	XImage_EnableCenter = XCDLL.MustFindProc("XImage_EnableCenter")
	XImage_IsCenter = XCDLL.MustFindProc("XImage_IsCenter")
	XImage_GetDrawType = XCDLL.MustFindProc("XImage_GetDrawType")
	XImage_GetWidth = XCDLL.MustFindProc("XImage_GetWidth")
	XImage_GetHeight = XCDLL.MustFindProc("XImage_GetHeight")
	XImage_Release = XCDLL.MustFindProc("XImage_Release")
	XImage_Destroy = XCDLL.MustFindProc("XImage_Destroy")
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-22 15:05:47
// @Function: XImageLoadFile
// @Description: 加载图片从文件.
// @Calls: XImage_LoadFile
// @Input: pImageName 图片文件. bStretch 是否拉伸图片.
// @Return: 图片句柄.
// *******************************************
func XImageLoadFile(pImageName string, bStretch bool) HIMAGE {
	ret, _, _ := XImage_LoadFile.Call(
		StringToUintPtr(pImageName),
		uintptr(BoolToBOOL(bStretch)))

	return HIMAGE(ret)
}
