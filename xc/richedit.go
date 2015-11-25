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
func XRichEditInsertImage(hEle HELE, hImage HIMAGE, pImagePath *string) bool {
	ret, _, _ := XRichEdit_InsertImage.Call(
		uintptr(hEle),
		uintptr(hImage),
		uintptr(unsafe.Pointer(pImagePath)))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 20:10:08
// @Function: XRichEditInsertGif
// @Description: 插入GIF动画图片到当前位置.
// @Calls: XRichEdit_InsertGif
// @Input: hEle 元素句柄. hImage 图片句柄. pImagePath 图片路径,用于图片的复制.
// @Return: 成功返回TRUE,否则相反.
// *******************************************
func XRichEditInsertGif(hEle HELE, hImage HIMAGE, pImagePath *string) bool {
	ret, _, _ := XRichEdit_InsertGif.Call(
		uintptr(hEle),
		uintptr(hImage),
		uintptr(unsafe.Pointer(pImagePath)))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 20:11:34
// @Function: XRichEditInsertStringEx
// @Description: 插入字符串.
// @Calls: XRichEdit_InsertStringEx
// @Input: hEle 元素句柄. iRow 行索引. iColumn 列索引. pString 字符串指针. pFont 字体. color 颜色.
// @Return:
// *******************************************
func XRichEditInsertStringEx(hEle HELE, iRow, iColumn int, pString *string, pFont *LOGFONTW, color COLORREF) {
	XRichEdit_InsertStringEx.Call(
		uintptr(hEle),
		uintptr(iRow),
		uintptr(iColumn),
		uintptr(unsafe.Pointer(pString)),
		uintptr(unsafe.Pointer(pFont)),
		uintptr(color))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 20:15:32
// @Function: XRichEditInsertImageEx
// @Description: 插入图片.
// @Calls: XRichEdit_InsertImageEx
// @Input: hEle 元素句柄. iRow 插入的行位置,如果值为-1插入到末尾行. iColumn 插入的列位置,如果值为-1插入末尾列.
//         hImage 图片句柄. pImagePath 图片路径,用于图片的复制.
// @Return: 成功返回TRUE,否则相反.
// *******************************************
func XRichEditInsertImageEx(hEle HELE, iRow, iColumn int, hImage HIMAGE, pImagePath *string) bool {
	ret, _, _ := XRichEdit_InsertImageEx.Call(
		uintptr(hEle),
		uintptr(iRow),
		uintptr(iColumn),
		uintptr(hImage),
		uintptr(unsafe.Pointer(pImagePath)))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 20:47:02
// @Function: XRichEditInsertGifEx
// @Description: 插入GIF动画图片.
// @Calls: XRichEdit_InsertGifEx
// @Input: hEle 元素句柄. iRow 插入的行位置,如果值为-1插入到末尾行. iColumn 插入的列位置,如果值为-1插入末尾列.
//         hImage 图片句柄. pImagePath 图片路径,用于图片的复制.
// @Return: 成功返回TRUE,否则相反.
// *******************************************
func XRichEditInsertGifEx(hEle HELE, iRow, iColumn int, hImage HIMAGE, pImagePath *string) bool {
	ret, _, _ := XRichEdit_InsertGifEx.Call(
		uintptr(hEle),
		uintptr(iRow),
		uintptr(iColumn),
		uintptr(hImage),
		uintptr(unsafe.Pointer(pImagePath)))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 20:48:33
// @Function: XRichEditSetText
// @Description: 设置文本内容.
// @Calls: XRichEdit_SetText
// @Input: hEle 元素句柄. pString 字符串指针.
// @Return:
// *******************************************
func XRichEditSetText(hEle HELE, pString *string) {
	XRichEdit_SetText.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pString)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 20:50:10
// @Function: XRichEditSetTextInt
// @Description: 设置内容.
// @Calls: XRichEdit_SetTextInt
// @Input: hEle 元素句柄. nVaule 整型值.
// @Return:
// *******************************************
func XRichEditSetTextInt(hEle HELE, nVaule int) {
	XRichEdit_SetTextInt.Call(
		uintptr(hEle),
		uintptr(nVaule))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 20:51:23
// @Function: XRichEditGetText
// @Description: 获取文本内容.
// @Calls: XRichEdit_GetText
// @Input: hEle 元素句柄. pOut 接收内容缓冲区. len 缓冲区长度,字符为单位.
// @Return: 返回内容长度,不包含空终止符.
// *******************************************
func XRichEditGetText(hEle HELE, pOut *uint16, len int) int {
	ret, _, _ := XRichEdit_GetText.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(len))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 20:54:14
// @Function: XRichEditGetHTMLFormat
// @Description: 获取HTML格式内容.
// @Calls: XRichEdit_GetHTMLFormat
// @Input: hEle 元素句柄. pOut 接收内容缓冲区. len 缓冲区长度,字符为单位.
// @Return: 返回内容长度,不包含空终止符.
// *******************************************
func XRichEditGetHTMLFormat(hEle HELE, pOut *uint16, len int) {
	XRichEdit_GetHTMLFormat.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(len))

}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 20:55:09
// @Function: XRichEditGetData
// @Description: 获取富文本数据,包含字体,图片,颜色信息.
// @Calls: XRichEdit_GetData
// @Input: hEle 元素句柄. pDataSize 返回数据大小,字节为单位.
// @Return: 返回数据内存指针,不用之后需要释放内容XC_Free().
// *******************************************
func XRichEditGetData(hEle HELE, pDataSize *int) {
	XRichEdit_GetData.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pDataSize)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 20:57:12
// @Function: XRichEditInsertData
// @Description: 插入数据内容,包含字体,图片,颜色信息.
// @Calls: XRichEdit_InsertData
// @Input: hEle 元素句柄. pData 数据指针. iRow 插入的行位置,如果值为-1插入到末尾行. iColumn 插入的列位置,如果值为-1插入末尾列.
// @Return: 成功返回TRUE,否则相反.
// *******************************************
func XRichEditInsertData(hEle HELE, pData *interface{}, iRow, iColumn int) bool {
	ret, _, _ := XRichEdit_InsertData.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pData)),
		uintptr(iRow),
		uintptr(iColumn))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:03:44
// @Function: XRichEditEnableReadOnly
// @Description: 启用只读模式.
// @Calls: XRichEdit_EnableReadOnly
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XRichEditEnableReadOnly(hEle HELE, bEnable bool) {
	XRichEdit_EnableReadOnly.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:05:11
// @Function: XRichEditEnableMultiLine
// @Description: 启用多行模式.
// @Calls: XRichEdit_EnableMultiLine
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XRichEditEnableMultiLine(hEle HELE, bEnable bool) {
	XRichEdit_EnableMultiLine.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:05:58
// @Function: XRichEditEnablePassword
// @Description: 启用密码模式.
// @Calls: XRichEdit_EnablePassword
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XRichEditEnablePassword(hEle HELE, bEnable bool) {
	XRichEdit_EnablePassword.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:06:41
// @Function: XRichEditEnableEvent_XE_RICHEDIT_CHANGE
// @Description: 启用XE_RICHEDIT_CHANGE事件.
// @Calls: XRichEdit_EnableEvent_XE_RICHEDIT_CHANGE
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XRichEditEnableEvent_XE_RICHEDIT_CHANGE(hEle HELE, bEnable bool) {
	XRichEdit_EnableEvent_XE_RICHEDIT_CHANGE.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:07:27
// @Function: XRichEditEnableAutoWrap
// @Description: 启用自动换行.
// @Calls: XRichEdit_EnableAutoWrap
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XRichEditEnableAutoWrap(hEle HELE, bEnable bool) {
	XRichEdit_EnableAutoWrap.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:08:12
// @Function: XRichEditEnableAutoCancelSel
// @Description: 当失去焦点时,自动取消选择内容.
// @Calls: XRichEdit_EnableAutoCancelSel
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XRichEditEnableAutoCancelSel(hEle HELE, bEnable bool) {
	XRichEdit_EnableAutoCancelSel.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:09:01
// @Function: XRichEditGetTextLength
// @Description: 获取文本内容长度.
// @Calls: XRichEdit_GetTextLength
// @Input: hEle 元素句柄.
// @Return: 长度,不包含空终止符.
// *******************************************
func XRichEditGetTextLength(hEle HELE) int {
	ret, _, _ := XRichEdit_GetTextLength.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:10:09
// @Function: XRichEditSetRowHeight
// @Description: 设置默认行高.
// @Calls: XRichEdit_SetRowHeight
// @Input: hEle 元素句柄. nHeight 高度.
// @Return:
// *******************************************
func XRichEditSetRowHeight(hEle HELE, nHeight uint) {
	XRichEdit_SetRowHeight.Call(
		uintptr(hEle),
		uintptr(nHeight))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:12:09
// @Function: XRichEditGetCurrentRow
// @Description: 获取当前行索引.
// @Calls: XRichEdit_GetCurrentRow
// @Input: hEle 元素句柄.
// @Return: 当前行索引.
// *******************************************
func XRichEditGetCurrentRow(hEle HELE) int {
	ret, _, _ := XRichEdit_GetCurrentRow.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:15:58
// @Function: XRichEditGetCurrentColumn
// @Description: 获取当前列索引.
// @Calls: XRichEdit_GetCurrentColumn
// @Input: hEle 元素句柄.
// @Return: 当前列索引.
// *******************************************
func XRichEditGetCurrentColumn(hEle HELE) int {
	ret, _, _ := XRichEdit_GetCurrentColumn.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:16:51
// @Function: XRichEditSetCurrentPos
// @Description: 设置当前输入位置.
// @Calls: XRichEdit_SetCurrentPos
// @Input: hEle 元素句柄. iRow 行索引,-1代表末尾行. iColumn 列索引,-1代表末尾列.
// @Return:
// *******************************************
func XRichEditSetCurrentPos(hEle HELE, iColumn, iRow int) {
	XRichEdit_SetCurrentPos.Call(
		uintptr(hEle),
		uintptr(iColumn),
		uintptr(iRow))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:18:41
// @Function: XRichEditGetRowCount
// @Description: 获取行数量.
// @Calls: XRichEdit_GetRowCount
// @Input: hEle 元素句柄.
// @Return: 行数量.
// *******************************************
func XRichEditGetRowCount(hEle HELE) int {
	ret, _, _ := XRichEdit_GetRowCount.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:20:09
// @Function: XRichEditGetRowLength
// @Description: 获取指定行长度.
// @Calls: XRichEdit_GetRowLength
// @Input: hEle 元素句柄. iRow 行索引.
// @Return: 行长度.
// *******************************************
func XRichEditGetRowLength(hEle HELE, iRow int) int {
	ret, _, _ := XRichEdit_GetRowLength.Call(
		uintptr(hEle),
		uintptr(iRow))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:21:40
// @Function: XRichEditGetSelectText
// @Description: 获取选择文本.
// @Calls: XRichEdit_GetSelectText
// @Input: hEle 元素句柄. pOut 接收内容缓冲区. len 缓冲区长度,字符为单位.
// @Return: 返回内容长度,字符为单位,不包含空终止符.
// *******************************************
func XRichEditGetSelectText(hEle HELE, pOut *uint16, len int) int {
	ret, _, _ := XRichEdit_GetSelectText.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(len))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:23:43
// @Function: XRichEditGetSelectPosition
// @Description: 获取选择内容位置.
// @Calls: XRichEdit_GetSelectPosition
// @Input: hEle 元素句柄. pBegin 开始位置点. pEnd 结束位置点.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XRichEditGetSelectPosition(hEle HELE, pBegin *Position, pEnd *Position) bool {
	ret, _, _ := XRichEdit_GetSelectPosition.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pBegin)),
		uintptr(unsafe.Pointer(pEnd)))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:28:29
// @Function: XRichEditSetSelect
// @Description: 设置选择内容.
// @Calls: XRichEdit_SetSelect
// @Input: hEle 元素句柄. iStartRow 开始行索引. iStartCol 开始行列索引. iEndRow 结束行索引. iEndCol 结束列索引.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XRichEditSetSelect(hEle HELE, iStartRow, iStartCol, iEndRow, iEndCol int) bool {
	ret, _, _ := XRichEdit_SetSelect.Call(
		uintptr(hEle),
		uintptr(iStartRow),
		uintptr(iStartCol),
		uintptr(iEndRow),
		uintptr(iEndCol))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:31:06
// @Function: XRichEditSetItemFontEx
// @Description: 设置指定内容字体.
// @Calls: XRichEdit_SetItemFontEx
// @Input: hEle 元素句柄. beginRow 开始行索引. beginColumn 开始行列索引. endRow 结束行索引. endColumn 结束行列索引. pFont 字体信息.
// @Return: 成功返回TRUE,否则相反.
// *******************************************
func XRichEditSetItemFontEx(hEle HELE, beginRow, beginColumn, endRow, endColumn int, pFont *LOGFONTW) bool {
	ret, _, _ := XRichEdit_SetItemFontEx.Call(
		uintptr(hEle),
		uintptr(beginRow),
		uintptr(beginColumn),
		uintptr(endRow),
		uintptr(endColumn),
		uintptr(unsafe.Pointer(pFont)))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:33:57
// @Function: XRichEditSetItemColorEx
// @Description: 设置指定内容字体颜色.
// @Calls: XRichEdit_SetItemColorEx
// @Input: hEle 元素句柄. beginRow 开始行索引. beginColumn 开始行列索引.
//         endRow 结束行索引. endColumn 结束行列索引. color RGB颜色. alpha 透明度.
// @Return: 成功返回TRUE,否则相反.
// *******************************************
func XRichEditSetItemColorEx(hEle HELE, beginRow, beginColumn, endRow, endColumn int, color COLORREF, alpha byte) bool {
	ret, _, _ := XRichEdit_SetItemColorEx.Call(
		uintptr(hEle),
		uintptr(beginRow),
		uintptr(beginColumn),
		uintptr(endRow),
		uintptr(endColumn),
		uintptr(color),
		uintptr(alpha))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:36:20
// @Function: XRichEditCancelSelect
// @Description: 取消选择.
// @Calls: XRichEdit_CancelSelect
// @Input: hEle 元素句柄.
// @Return:
// *******************************************
func XRichEditCancelSelect(hEle HELE) {
	XRichEdit_CancelSelect.Call(uintptr(hEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:37:46
// @Function: XRichEditSetSelectBkColor
// @Description: 设置选择内容背景颜色.
// @Calls: XRichEdit_SetSelectBkColor
// @Input: hEle 元素句柄. color 颜色值. alpha 透明度.
// @Return:
// *******************************************
func XRichEditSetSelectBkColor(hEle HELE, color COLORREF, alpha byte) {
	XRichEdit_SetSelectBkColor.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:37:46
// @Function: XRichEditIsEmpty
// @Description: 内容是否为空.
// @Calls: XRichEdit_IsEmpty
// @Input: hEle 元素句柄.
// @Return: 如果为空返回TRUE否则返回FALSE.
// *******************************************
func XRichEditIsEmpty(hEle HELE) bool {
	ret, _, _ := XRichEdit_IsEmpty.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:43:26
// @Function: XRichEditSelectAll
// @Description: 选择所有内容.
// @Calls: XRichEdit_SelectAll
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XRichEditSelectAll(hEle HELE) bool {
	ret, _, _ := XRichEdit_SelectAll.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:44:17
// @Function: XRichEditDeleteSelect
// @Description: 删除选择内容.
// @Calls: XRichEdit_DeleteSelect
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XRichEditDeleteSelect(hEle HELE) bool {
	ret, _, _ := XRichEdit_DeleteSelect.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:45:01
// @Function: XRichEditDeleteAll
// @Description: 删除所有内容.
// @Calls: XRichEdit_DeleteAll
// @Input: hEle 元素句柄.
// @Return:
// *******************************************
func XRichEditDeleteAll(hEle HELE) {
	XRichEdit_DeleteAll.Call(uintptr(hEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:49:55
// @Function: XRichEditClipboardCut
// @Description: 剪切选择内容到剪贴板.
// @Calls: XRichEdit_ClipboardCut
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XRichEditClipboardCut(hEle HELE) bool {
	ret, _, _ := XRichEdit_ClipboardCut.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:50:53
// @Function: XRichEditClipboardCopy
// @Description: 复制选择内容到剪贴板.
// @Calls: XRichEdit_ClipboardCopy
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XRichEditClipboardCopy(hEle HELE) bool {
	ret, _, _ := XRichEdit_ClipboardCopy.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:52:12
// @Function: XRichEditClipboardPaste
// @Description: 从剪贴板粘贴内容到当前位置.
// @Calls: XRichEdit_ClipboardPaste
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XRichEditClipboardPaste(hEle HELE) bool {
	ret, _, _ := XRichEdit_ClipboardPaste.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}
