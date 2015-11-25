package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	XRichEdit_Create                         *syscall.Proc
	XRichEdit_InsertString                   *syscall.Proc
	XRichEdit_InsertImage                    *syscall.Proc
	XRichEdit_InsertGif                      *syscall.Proc
	XRichEdit_InsertStringEx                 *syscall.Proc
	XRichEdit_InsertImageEx                  *syscall.Proc
	XRichEdit_InsertGifEx                    *syscall.Proc
	XRichEdit_SetText                        *syscall.Proc
	XRichEdit_SetTextInt                     *syscall.Proc
	XRichEdit_GetText                        *syscall.Proc
	XRichEdit_GetHTMLFormat                  *syscall.Proc
	XRichEdit_GetData                        *syscall.Proc
	XRichEdit_InsertData                     *syscall.Proc
	XRichEdit_EnableReadOnly                 *syscall.Proc
	XRichEdit_EnableMultiLine                *syscall.Proc
	XRichEdit_EnablePassword                 *syscall.Proc
	XRichEdit_EnableEvent_XE_RICHEDIT_CHANGE *syscall.Proc
	XRichEdit_EnableAutoWrap                 *syscall.Proc
	XRichEdit_EnableAutoCancelSel            *syscall.Proc
	XRichEdit_GetTextLength                  *syscall.Proc
	XRichEdit_SetRowHeight                   *syscall.Proc
	XRichEdit_GetCurrentRow                  *syscall.Proc
	XRichEdit_GetCurrentColumn               *syscall.Proc
	XRichEdit_SetCurrentPos                  *syscall.Proc
	XRichEdit_GetRowCount                    *syscall.Proc
	XRichEdit_GetRowLength                   *syscall.Proc
	XRichEdit_GetSelectText                  *syscall.Proc
	XRichEdit_GetSelectPosition              *syscall.Proc
	XRichEdit_SetSelect                      *syscall.Proc
	XRichEdit_SetItemFontEx                  *syscall.Proc
	XRichEdit_SetItemColorEx                 *syscall.Proc
	XRichEdit_CancelSelect                   *syscall.Proc
	XRichEdit_SetSelectBkColor               *syscall.Proc
	XRichEdit_IsEmpty                        *syscall.Proc
	XRichEdit_SelectAll                      *syscall.Proc
	XRichEdit_DeleteSelect                   *syscall.Proc
	XRichEdit_DeleteAll                      *syscall.Proc
	XRichEdit_ClipboardCut                   *syscall.Proc
	XRichEdit_ClipboardCopy                  *syscall.Proc
	XRichEdit_ClipboardPaste                 *syscall.Proc
)

func init() {
	XRichEdit_Create = XCDLL.MustFindProc("XRichEdit_Create")
	XRichEdit_InsertString = XCDLL.MustFindProc("XRichEdit_InsertString")
	XRichEdit_InsertImage = XCDLL.MustFindProc("XRichEdit_InsertImage")
	XRichEdit_InsertGif = XCDLL.MustFindProc("XRichEdit_InsertGif")
	XRichEdit_InsertStringEx = XCDLL.MustFindProc("XRichEdit_InsertStringEx")
	XRichEdit_InsertImageEx = XCDLL.MustFindProc("XRichEdit_InsertImageEx")
	XRichEdit_InsertGifEx = XCDLL.MustFindProc("XRichEdit_InsertGifEx")
	XRichEdit_SetText = XCDLL.MustFindProc("XRichEdit_SetText")
	XRichEdit_SetTextInt = XCDLL.MustFindProc("XRichEdit_SetTextInt")
	XRichEdit_GetText = XCDLL.MustFindProc("XRichEdit_GetText")
	XRichEdit_GetHTMLFormat = XCDLL.MustFindProc("XRichEdit_GetHTMLFormat")
	XRichEdit_GetData = XCDLL.MustFindProc("XRichEdit_GetData")
	XRichEdit_InsertData = XCDLL.MustFindProc("XRichEdit_InsertData")
	XRichEdit_EnableReadOnly = XCDLL.MustFindProc("XRichEdit_EnableReadOnly")
	XRichEdit_EnableMultiLine = XCDLL.MustFindProc("XRichEdit_EnableMultiLine")
	XRichEdit_EnablePassword = XCDLL.MustFindProc("XRichEdit_EnablePassword")
	XRichEdit_EnableEvent_XE_RICHEDIT_CHANGE = XCDLL.MustFindProc("XRichEdit_EnableEvent_XE_RICHEDIT_CHANGE")
	XRichEdit_EnableAutoWrap = XCDLL.MustFindProc("XRichEdit_EnableAutoWrap")
	XRichEdit_EnableAutoCancelSel = XCDLL.MustFindProc("XRichEdit_EnableAutoCancelSel")
	XRichEdit_GetTextLength = XCDLL.MustFindProc("XRichEdit_GetTextLength")
	XRichEdit_SetRowHeight = XCDLL.MustFindProc("XRichEdit_SetRowHeight")
	XRichEdit_GetCurrentRow = XCDLL.MustFindProc("XRichEdit_GetCurrentRow")
	XRichEdit_GetCurrentColumn = XCDLL.MustFindProc("XRichEdit_GetCurrentColumn")
	XRichEdit_SetCurrentPos = XCDLL.MustFindProc("XRichEdit_SetCurrentPos")
	XRichEdit_GetRowCount = XCDLL.MustFindProc("XRichEdit_GetRowCount")
	XRichEdit_GetRowLength = XCDLL.MustFindProc("XRichEdit_GetRowLength")
	XRichEdit_GetSelectText = XCDLL.MustFindProc("XRichEdit_GetSelectText")
	XRichEdit_GetSelectPosition = XCDLL.MustFindProc("XRichEdit_GetSelectPosition")
	XRichEdit_SetSelect = XCDLL.MustFindProc("XRichEdit_SetSelect")
	XRichEdit_SetItemFontEx = XCDLL.MustFindProc("XRichEdit_SetItemFontEx")
	XRichEdit_SetItemColorEx = XCDLL.MustFindProc("XRichEdit_SetItemColorEx")
	XRichEdit_CancelSelect = XCDLL.MustFindProc("XRichEdit_CancelSelect")
	XRichEdit_SetSelectBkColor = XCDLL.MustFindProc("XRichEdit_SetSelectBkColor")
	XRichEdit_IsEmpty = XCDLL.MustFindProc("XRichEdit_IsEmpty")
	XRichEdit_SelectAll = XCDLL.MustFindProc("XRichEdit_SelectAll")
	XRichEdit_DeleteSelect = XCDLL.MustFindProc("XRichEdit_DeleteSelect")
	XRichEdit_DeleteAll = XCDLL.MustFindProc("XRichEdit_DeleteAll")
	XRichEdit_ClipboardCut = XCDLL.MustFindProc("XRichEdit_ClipboardCut")
	XRichEdit_ClipboardCopy = XCDLL.MustFindProc("XRichEdit_ClipboardCopy")
	XRichEdit_ClipboardPaste = XCDLL.MustFindProc("XRichEdit_ClipboardPaste")
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 10:32:00
// @Function: XRichEditCreate
// @Description: 创建富文本编辑框.
// @Calls: XRichEdit_Create
// @Input: x 元素x坐标. y 元素y坐标. cx 宽度. cy 高度.
//         hParent 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
// @Return: 元素句柄.
// *******************************************
func XRichEditCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := XRichEdit_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 10:35:49
// @Function: XRichEditInsertString
// @Description: 插入字符串到当前位置.
// @Calls: XRichEdit_InsertString
// @Input: hEle 元素句柄. pString 字符串指针. pFont 字体. color 颜色.
// @Return: 元素句柄.
// *******************************************
func XRichEditInsertString(hEle HELE, pString string, pFont *LOGFONTW, color COLORREF) {
	XRichEdit_InsertString.Call(
		uintptr(hEle),
		StringToUintPtr(pString),
		uintptr(unsafe.Pointer(pFont)),
		uintptr(color))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 11:04:08
// @Function: XRichEditInsertString
// @Description: 插入图片到当前位置.
// @Calls: XRichEdit_InsertString
// @Input: hEle 元素句柄. hImage 图片句柄. pImagePath 图片路径,用于图片的复制.
// @Return: 成功返回TRUE,否则相反.
// *******************************************
// func XRichEditInsertImage() {
//     XRichEdit_InsertImage
// }

// XRichEdit_InsertGif
// XRichEdit_InsertStringEx
// XRichEdit_InsertImageEx
// XRichEdit_InsertGifEx
// XRichEdit_SetText
// XRichEdit_SetTextInt
// XRichEdit_GetText
// XRichEdit_GetHTMLFormat
// XRichEdit_GetData
// XRichEdit_InsertData
// XRichEdit_EnableReadOnly
// XRichEdit_EnableMultiLine
// XRichEdit_EnablePassword
// XRichEdit_EnableEvent_XE_RICHEDIT_CHANGE
// XRichEdit_EnableAutoWrap
// XRichEdit_EnableAutoCancelSel
// XRichEdit_GetTextLength
// XRichEdit_SetRowHeight
// XRichEdit_GetCurrentRow
// XRichEdit_GetCurrentColumn
// XRichEdit_SetCurrentPos
// XRichEdit_GetRowCount
// XRichEdit_GetRowLength
// XRichEdit_GetSelectText
// XRichEdit_GetSelectPosition
// XRichEdit_SetSelect
// XRichEdit_SetItemFontEx
// XRichEdit_SetItemColorEx
// XRichEdit_CancelSelect
// XRichEdit_SetSelectBkColor
// XRichEdit_IsEmpty
// XRichEdit_SelectAll
// XRichEdit_DeleteSelect
// XRichEdit_DeleteAll
// XRichEdit_ClipboardCut
// XRichEdit_ClipboardCopy
// XRichEdit_ClipboardPaste
