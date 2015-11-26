package xc

import (
	"syscall"
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

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-22 15:05:47
// @Function: XImageLoadFile
// @Description: 加载图片从文件.
// @Calls: xImage_LoadFile
// @Input: pImageName 图片文件. bStretch 是否拉伸图片.
// @Return: 图片句柄.
// *******************************************
func XImageLoadFile(pImageName string, bStretch bool) HIMAGE {
	ret, _, _ := xImage_LoadFile.Call(
		StringToUintPtr(pImageName),
		uintptr(BoolToBOOL(bStretch)))

	return HIMAGE(ret)
}
