package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xDraw_Create                 *syscall.Proc
	xDraw_Destroy                *syscall.Proc
	xDraw_SetOffset              *syscall.Proc
	xDraw_GetOffset              *syscall.Proc
	xDraw_RestoreGDIOBJ          *syscall.Proc
	xDraw_GetHDC                 *syscall.Proc
	xDraw_SetBrushColor          *syscall.Proc
	xDraw_SetTextVertical        *syscall.Proc
	xDraw_SetTextAlign           *syscall.Proc
	xDraw_SetFontX               *syscall.Proc
	xDraw_SetFont                *syscall.Proc
	xDraw_SetFont2               *syscall.Proc
	xDraw_SetFont3               *syscall.Proc
	xDraw_SetLineWidth           *syscall.Proc
	xDraw_SetBkMode              *syscall.Proc
	xDraw_CreateSolidBrush       *syscall.Proc
	xDraw_CreatePen              *syscall.Proc
	xDraw_CreateRectRgn          *syscall.Proc
	xDraw_CreateRoundRectRgn     *syscall.Proc
	xDraw_CreatePolygonRgn       *syscall.Proc
	xDraw_SelectClipRgn          *syscall.Proc
	xDraw_FillRect               *syscall.Proc
	xDraw_FillRectColor          *syscall.Proc
	xDraw_FillRgn                *syscall.Proc
	xDraw_FillEllipse            *syscall.Proc
	xDraw_DrawEllipse            *syscall.Proc
	xDraw_FillRoundRect          *syscall.Proc
	xDraw_DrawRoundRect          *syscall.Proc
	xDraw_Rectangle              *syscall.Proc
	xDraw_DrawGroupBox_Rect      *syscall.Proc
	xDraw_DrawGroupBox_RoundRect *syscall.Proc
	xDraw_GradientFill2          *syscall.Proc
	xDraw_GradientFill4          *syscall.Proc
	xDraw_FrameRgn               *syscall.Proc
	xDraw_FrameRect              *syscall.Proc
	xDraw_DrawLine               *syscall.Proc
	xDraw_FocusRect              *syscall.Proc
	xDraw_MoveToEx               *syscall.Proc
	xDraw_LineTo                 *syscall.Proc
	xDraw_Polyline               *syscall.Proc
	xDraw_Dottedline             *syscall.Proc
	xDraw_SetPixel               *syscall.Proc
	xDraw_Check                  *syscall.Proc
	xDraw_DrawIconEx             *syscall.Proc
	xDraw_BitBlt                 *syscall.Proc
	xDraw_BitBlt2                *syscall.Proc
	xDraw_AlphaBlend             *syscall.Proc
	xDraw_TriangularArrow        *syscall.Proc
	xDraw_Image                  *syscall.Proc
	xDraw_Image2                 *syscall.Proc
	xDraw_ImageStretch           *syscall.Proc
	xDraw_ImageAdaptive          *syscall.Proc
	xDraw_ImageExTile            *syscall.Proc
	xDraw_ImageSuper             *syscall.Proc
	xDraw_ImageSuper2            *syscall.Proc
	xDraw_DrawText               *syscall.Proc
	xDraw_DrawTextUnderline      *syscall.Proc
	xDraw_TextOut                *syscall.Proc
	xDraw_TextOutA               *syscall.Proc
)

func init() {
	// Functions
	xDraw_Create = XCDLL.MustFindProc("XDraw_Create")
	xDraw_Destroy = XCDLL.MustFindProc("XDraw_Destroy")
	xDraw_SetOffset = XCDLL.MustFindProc("XDraw_SetOffset")
	xDraw_GetOffset = XCDLL.MustFindProc("XDraw_GetOffset")
	xDraw_RestoreGDIOBJ = XCDLL.MustFindProc("XDraw_RestoreGDIOBJ")
	xDraw_GetHDC = XCDLL.MustFindProc("XDraw_GetHDC")
	xDraw_SetBrushColor = XCDLL.MustFindProc("XDraw_SetBrushColor")
	xDraw_SetTextVertical = XCDLL.MustFindProc("XDraw_SetTextVertical")
	xDraw_SetTextAlign = XCDLL.MustFindProc("XDraw_SetTextAlign")
	xDraw_SetFontX = XCDLL.MustFindProc("XDraw_SetFontX")
	xDraw_SetFont = XCDLL.MustFindProc("XDraw_SetFont")
	xDraw_SetFont2 = XCDLL.MustFindProc("XDraw_SetFont2")
	xDraw_SetFont3 = XCDLL.MustFindProc("XDraw_SetFont3")
	xDraw_SetLineWidth = XCDLL.MustFindProc("XDraw_SetLineWidth")
	xDraw_SetBkMode = XCDLL.MustFindProc("XDraw_SetBkMode")
	xDraw_CreateSolidBrush = XCDLL.MustFindProc("XDraw_CreateSolidBrush")
	xDraw_CreatePen = XCDLL.MustFindProc("XDraw_CreatePen")
	xDraw_CreateRectRgn = XCDLL.MustFindProc("XDraw_CreateRectRgn")
	xDraw_CreateRoundRectRgn = XCDLL.MustFindProc("XDraw_CreateRoundRectRgn")
	xDraw_CreatePolygonRgn = XCDLL.MustFindProc("XDraw_CreatePolygonRgn")
	xDraw_SelectClipRgn = XCDLL.MustFindProc("XDraw_SelectClipRgn")
	xDraw_FillRect = XCDLL.MustFindProc("XDraw_FillRect")
	xDraw_FillRectColor = XCDLL.MustFindProc("XDraw_FillRectColor")
	xDraw_FillRgn = XCDLL.MustFindProc("XDraw_FillRgn")
	xDraw_FillEllipse = XCDLL.MustFindProc("XDraw_FillEllipse")
	xDraw_DrawEllipse = XCDLL.MustFindProc("XDraw_DrawEllipse")
	xDraw_FillRoundRect = XCDLL.MustFindProc("XDraw_FillRoundRect")
	xDraw_DrawRoundRect = XCDLL.MustFindProc("XDraw_DrawRoundRect")
	xDraw_Rectangle = XCDLL.MustFindProc("XDraw_Rectangle")
	xDraw_DrawGroupBox_Rect = XCDLL.MustFindProc("XDraw_DrawGroupBox_Rect")
	xDraw_DrawGroupBox_RoundRect = XCDLL.MustFindProc("XDraw_DrawGroupBox_RoundRect")
	xDraw_GradientFill2 = XCDLL.MustFindProc("XDraw_GradientFill2")
	xDraw_GradientFill4 = XCDLL.MustFindProc("XDraw_GradientFill4")
	xDraw_FrameRgn = XCDLL.MustFindProc("XDraw_FrameRgn")
	xDraw_FrameRect = XCDLL.MustFindProc("XDraw_FrameRect")
	xDraw_DrawLine = XCDLL.MustFindProc("XDraw_DrawLine")
	xDraw_FocusRect = XCDLL.MustFindProc("XDraw_FocusRect")
	xDraw_MoveToEx = XCDLL.MustFindProc("XDraw_MoveToEx")
	xDraw_LineTo = XCDLL.MustFindProc("XDraw_LineTo")
	xDraw_Polyline = XCDLL.MustFindProc("XDraw_Polyline")
	xDraw_Dottedline = XCDLL.MustFindProc("XDraw_Dottedline")
	xDraw_SetPixel = XCDLL.MustFindProc("XDraw_SetPixel")
	xDraw_Check = XCDLL.MustFindProc("XDraw_Check")
	xDraw_DrawIconEx = XCDLL.MustFindProc("XDraw_DrawIconEx")
	xDraw_BitBlt = XCDLL.MustFindProc("XDraw_BitBlt")
	xDraw_BitBlt2 = XCDLL.MustFindProc("XDraw_BitBlt2")
	xDraw_AlphaBlend = XCDLL.MustFindProc("XDraw_AlphaBlend")
	xDraw_TriangularArrow = XCDLL.MustFindProc("XDraw_TriangularArrow")
	xDraw_Image = XCDLL.MustFindProc("XDraw_Image")
	xDraw_Image2 = XCDLL.MustFindProc("XDraw_Image2")
	xDraw_ImageStretch = XCDLL.MustFindProc("XDraw_ImageStretch")
	xDraw_ImageAdaptive = XCDLL.MustFindProc("XDraw_ImageAdaptive")
	xDraw_ImageExTile = XCDLL.MustFindProc("XDraw_ImageExTile")
	xDraw_ImageSuper = XCDLL.MustFindProc("XDraw_ImageSuper")
	xDraw_ImageSuper2 = XCDLL.MustFindProc("XDraw_ImageSuper2")
	xDraw_DrawText = XCDLL.MustFindProc("XDraw_DrawText")
	xDraw_DrawTextUnderline = XCDLL.MustFindProc("XDraw_DrawTextUnderline")
	xDraw_TextOut = XCDLL.MustFindProc("XDraw_TextOut")
	xDraw_TextOutA = XCDLL.MustFindProc("XDraw_TextOutA")
}

func XDrawCreate() {
	xDraw_Create
}

//
// XDraw_Destroy
// XDraw_SetOffset
// XDraw_GetOffset
// XDraw_RestoreGDIOBJ
// XDraw_GetHDC
// XDraw_SetBrushColor
// XDraw_SetTextVertical
// XDraw_SetTextAlign
// XDraw_SetFontX
// XDraw_SetFont
// XDraw_SetFont2
// XDraw_SetFont3
// XDraw_SetLineWidth
// XDraw_SetBkMode
// XDraw_CreateSolidBrush
// XDraw_CreatePen
// XDraw_CreateRectRgn
// XDraw_CreateRoundRectRgn
// XDraw_CreatePolygonRgn
// XDraw_SelectClipRgn
// XDraw_FillRect
// XDraw_FillRectColor
// XDraw_FillRgn
// XDraw_FillEllipse
// XDraw_DrawEllipse
// XDraw_FillRoundRect
// XDraw_DrawRoundRect
// XDraw_Rectangle
// XDraw_DrawGroupBox_Rect
// XDraw_DrawGroupBox_RoundRect
// XDraw_GradientFill2
// XDraw_GradientFill4
// XDraw_FrameRgn
// XDraw_FrameRect
// XDraw_DrawLine
// XDraw_FocusRect
// XDraw_MoveToEx
// XDraw_LineTo
// XDraw_Polyline
// XDraw_Dottedline
// XDraw_SetPixel
// XDraw_Check
// XDraw_DrawIconEx
// XDraw_BitBlt
// XDraw_BitBlt2
// XDraw_AlphaBlend
// XDraw_TriangularArrow
// XDraw_Image
// XDraw_Image2
// XDraw_ImageStretch
// XDraw_ImageAdaptive
// XDraw_ImageExTile
// XDraw_ImageSuper
// XDraw_ImageSuper2
// XDraw_DrawText
// XDraw_DrawTextUnderline
// XDraw_TextOut
// XDraw_TextOutA
