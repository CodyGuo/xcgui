package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xRichEdit_Create                         *syscall.Proc
	xRichEdit_InsertString                   *syscall.Proc
	xRichEdit_InsertImage                    *syscall.Proc
	xRichEdit_InsertGif                      *syscall.Proc
	xRichEdit_InsertStringEx                 *syscall.Proc
	xRichEdit_InsertImageEx                  *syscall.Proc
	xRichEdit_InsertGifEx                    *syscall.Proc
	xRichEdit_SetText                        *syscall.Proc
	xRichEdit_SetTextInt                     *syscall.Proc
	xRichEdit_GetText                        *syscall.Proc
	xRichEdit_GetHTMLFormat                  *syscall.Proc
	xRichEdit_GetData                        *syscall.Proc
	xRichEdit_InsertData                     *syscall.Proc
	xRichEdit_EnableReadOnly                 *syscall.Proc
	xRichEdit_EnableMultiLine                *syscall.Proc
	xRichEdit_EnablePassword                 *syscall.Proc
	xRichEdit_EnableEvent_XE_RICHEDIT_CHANGE *syscall.Proc
	xRichEdit_EnableAutoWrap                 *syscall.Proc
	xRichEdit_EnableAutoCancelSel            *syscall.Proc
	xRichEdit_GetTextLength                  *syscall.Proc
	xRichEdit_SetRowHeight                   *syscall.Proc
	xRichEdit_GetCurrentRow                  *syscall.Proc
	xRichEdit_GetCurrentColumn               *syscall.Proc
	xRichEdit_SetCurrentPos                  *syscall.Proc
	xRichEdit_GetRowCount                    *syscall.Proc
	xRichEdit_GetRowLength                   *syscall.Proc
	xRichEdit_GetSelectText                  *syscall.Proc
	xRichEdit_GetSelectPosition              *syscall.Proc
	xRichEdit_SetSelect                      *syscall.Proc
	xRichEdit_SetItemFontEx                  *syscall.Proc
	xRichEdit_SetItemColorEx                 *syscall.Proc
	xRichEdit_CancelSelect                   *syscall.Proc
	xRichEdit_SetSelectBkColor               *syscall.Proc
	xRichEdit_IsEmpty                        *syscall.Proc
	xRichEdit_SelectAll                      *syscall.Proc
	xRichEdit_DeleteSelect                   *syscall.Proc
	xRichEdit_DeleteAll                      *syscall.Proc
	xRichEdit_ClipboardCut                   *syscall.Proc
	xRichEdit_ClipboardCopy                  *syscall.Proc
	xRichEdit_ClipboardPaste                 *syscall.Proc
)

func init() {
	xRichEdit_Create = xcDLL.MustFindProc("XRichEdit_Create")
	xRichEdit_InsertString = xcDLL.MustFindProc("XRichEdit_InsertString")
	xRichEdit_InsertImage = xcDLL.MustFindProc("XRichEdit_InsertImage")
	xRichEdit_InsertGif = xcDLL.MustFindProc("XRichEdit_InsertGif")
	xRichEdit_InsertStringEx = xcDLL.MustFindProc("XRichEdit_InsertStringEx")
	xRichEdit_InsertImageEx = xcDLL.MustFindProc("XRichEdit_InsertImageEx")
	xRichEdit_InsertGifEx = xcDLL.MustFindProc("XRichEdit_InsertGifEx")
	xRichEdit_SetText = xcDLL.MustFindProc("XRichEdit_SetText")
	xRichEdit_SetTextInt = xcDLL.MustFindProc("XRichEdit_SetTextInt")
	xRichEdit_GetText = xcDLL.MustFindProc("XRichEdit_GetText")
	xRichEdit_GetHTMLFormat = xcDLL.MustFindProc("XRichEdit_GetHTMLFormat")
	xRichEdit_GetData = xcDLL.MustFindProc("XRichEdit_GetData")
	xRichEdit_InsertData = xcDLL.MustFindProc("XRichEdit_InsertData")
	xRichEdit_EnableReadOnly = xcDLL.MustFindProc("XRichEdit_EnableReadOnly")
	xRichEdit_EnableMultiLine = xcDLL.MustFindProc("XRichEdit_EnableMultiLine")
	xRichEdit_EnablePassword = xcDLL.MustFindProc("XRichEdit_EnablePassword")
	xRichEdit_EnableEvent_XE_RICHEDIT_CHANGE = xcDLL.MustFindProc("XRichEdit_EnableEvent_XE_RICHEDIT_CHANGE")
	xRichEdit_EnableAutoWrap = xcDLL.MustFindProc("XRichEdit_EnableAutoWrap")
	xRichEdit_EnableAutoCancelSel = xcDLL.MustFindProc("XRichEdit_EnableAutoCancelSel")
	xRichEdit_GetTextLength = xcDLL.MustFindProc("XRichEdit_GetTextLength")
	xRichEdit_SetRowHeight = xcDLL.MustFindProc("XRichEdit_SetRowHeight")
	xRichEdit_GetCurrentRow = xcDLL.MustFindProc("XRichEdit_GetCurrentRow")
	xRichEdit_GetCurrentColumn = xcDLL.MustFindProc("XRichEdit_GetCurrentColumn")
	xRichEdit_SetCurrentPos = xcDLL.MustFindProc("XRichEdit_SetCurrentPos")
	xRichEdit_GetRowCount = xcDLL.MustFindProc("XRichEdit_GetRowCount")
	xRichEdit_GetRowLength = xcDLL.MustFindProc("XRichEdit_GetRowLength")
	xRichEdit_GetSelectText = xcDLL.MustFindProc("XRichEdit_GetSelectText")
	xRichEdit_GetSelectPosition = xcDLL.MustFindProc("XRichEdit_GetSelectPosition")
	xRichEdit_SetSelect = xcDLL.MustFindProc("XRichEdit_SetSelect")
	xRichEdit_SetItemFontEx = xcDLL.MustFindProc("XRichEdit_SetItemFontEx")
	xRichEdit_SetItemColorEx = xcDLL.MustFindProc("XRichEdit_SetItemColorEx")
	xRichEdit_CancelSelect = xcDLL.MustFindProc("XRichEdit_CancelSelect")
	xRichEdit_SetSelectBkColor = xcDLL.MustFindProc("XRichEdit_SetSelectBkColor")
	xRichEdit_IsEmpty = xcDLL.MustFindProc("XRichEdit_IsEmpty")
	xRichEdit_SelectAll = xcDLL.MustFindProc("XRichEdit_SelectAll")
	xRichEdit_DeleteSelect = xcDLL.MustFindProc("XRichEdit_DeleteSelect")
	xRichEdit_DeleteAll = xcDLL.MustFindProc("XRichEdit_DeleteAll")
	xRichEdit_ClipboardCut = xcDLL.MustFindProc("XRichEdit_ClipboardCut")
	xRichEdit_ClipboardCopy = xcDLL.MustFindProc("XRichEdit_ClipboardCopy")
	xRichEdit_ClipboardPaste = xcDLL.MustFindProc("XRichEdit_ClipboardPaste")
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 10:32:00
// @Function: XRichEditCreate
// @Description: 创建富文本编辑框.
// @Calls: xRichEdit_Create
// @Input: x 元素x坐标. y 元素y坐标. cx 宽度. cy 高度.
//         hParent 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
// @Return: 元素句柄.
// *******************************************
func XRichEditCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xRichEdit_Create.Call(
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
// @Calls: xRichEdit_InsertString
// @Input: hEle 元素句柄. pString 字符串指针. pFont 字体. color 颜色.
// @Return: 元素句柄.
// *******************************************
func XRichEditInsertString(hEle HELE, pString string, pFont *LOGFONTW, color COLORREF) {
	xRichEdit_InsertString.Call(
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
// @Calls: xRichEdit_InsertString
// @Input: hEle 元素句柄. hImage 图片句柄. pImagePath 图片路径,用于图片的复制.
// @Return: 成功返回TRUE,否则相反.
// *******************************************
func XRichEditInsertImage(hEle HELE, hImage HIMAGE, pImagePath *string) bool {
	ret, _, _ := xRichEdit_InsertImage.Call(
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
// @Calls: xRichEdit_InsertGif
// @Input: hEle 元素句柄. hImage 图片句柄. pImagePath 图片路径,用于图片的复制.
// @Return: 成功返回TRUE,否则相反.
// *******************************************
func XRichEditInsertGif(hEle HELE, hImage HIMAGE, pImagePath *string) bool {
	ret, _, _ := xRichEdit_InsertGif.Call(
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
// @Calls: xRichEdit_InsertStringEx
// @Input: hEle 元素句柄. iRow 行索引. iColumn 列索引. pString 字符串指针. pFont 字体. color 颜色.
// @Return:
// *******************************************
func XRichEditInsertStringEx(hEle HELE, iRow, iColumn int, pString *string, pFont *LOGFONTW, color COLORREF) {
	xRichEdit_InsertStringEx.Call(
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
// @Calls: xRichEdit_InsertImageEx
// @Input: hEle 元素句柄. iRow 插入的行位置,如果值为-1插入到末尾行. iColumn 插入的列位置,如果值为-1插入末尾列.
//         hImage 图片句柄. pImagePath 图片路径,用于图片的复制.
// @Return: 成功返回TRUE,否则相反.
// *******************************************
func XRichEditInsertImageEx(hEle HELE, iRow, iColumn int, hImage HIMAGE, pImagePath *string) bool {
	ret, _, _ := xRichEdit_InsertImageEx.Call(
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
// @Calls: xRichEdit_InsertGifEx
// @Input: hEle 元素句柄. iRow 插入的行位置,如果值为-1插入到末尾行. iColumn 插入的列位置,如果值为-1插入末尾列.
//         hImage 图片句柄. pImagePath 图片路径,用于图片的复制.
// @Return: 成功返回TRUE,否则相反.
// *******************************************
func XRichEditInsertGifEx(hEle HELE, iRow, iColumn int, hImage HIMAGE, pImagePath *string) bool {
	ret, _, _ := xRichEdit_InsertGifEx.Call(
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
// @Calls: xRichEdit_SetText
// @Input: hEle 元素句柄. pString 字符串指针.
// @Return:
// *******************************************
func XRichEditSetText(hEle HELE, pString *string) {
	xRichEdit_SetText.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pString)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 20:50:10
// @Function: XRichEditSetTextInt
// @Description: 设置内容.
// @Calls: xRichEdit_SetTextInt
// @Input: hEle 元素句柄. nVaule 整型值.
// @Return:
// *******************************************
func XRichEditSetTextInt(hEle HELE, nVaule int) {
	xRichEdit_SetTextInt.Call(
		uintptr(hEle),
		uintptr(nVaule))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 20:51:23
// @Function: XRichEditGetText
// @Description: 获取文本内容.
// @Calls: xRichEdit_GetText
// @Input: hEle 元素句柄. pOut 接收内容缓冲区. len 缓冲区长度,字符为单位.
// @Return: 返回内容长度,不包含空终止符.
// *******************************************
func XRichEditGetText(hEle HELE, pOut *uint16, len int) int {
	ret, _, _ := xRichEdit_GetText.Call(
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
// @Calls: xRichEdit_GetHTMLFormat
// @Input: hEle 元素句柄. pOut 接收内容缓冲区. len 缓冲区长度,字符为单位.
// @Return: 返回内容长度,不包含空终止符.
// *******************************************
func XRichEditGetHTMLFormat(hEle HELE, pOut *uint16, len int) {
	xRichEdit_GetHTMLFormat.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(len))

}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 20:55:09
// @Function: XRichEditGetData
// @Description: 获取富文本数据,包含字体,图片,颜色信息.
// @Calls: xRichEdit_GetData
// @Input: hEle 元素句柄. pDataSize 返回数据大小,字节为单位.
// @Return: 返回数据内存指针,不用之后需要释放内容XC_Free().
// *******************************************
func XRichEditGetData(hEle HELE, pDataSize *int) {
	xRichEdit_GetData.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pDataSize)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 20:57:12
// @Function: XRichEditInsertData
// @Description: 插入数据内容,包含字体,图片,颜色信息.
// @Calls: xRichEdit_InsertData
// @Input: hEle 元素句柄. pData 数据指针. iRow 插入的行位置,如果值为-1插入到末尾行. iColumn 插入的列位置,如果值为-1插入末尾列.
// @Return: 成功返回TRUE,否则相反.
// *******************************************
func XRichEditInsertData(hEle HELE, pData *interface{}, iRow, iColumn int) bool {
	ret, _, _ := xRichEdit_InsertData.Call(
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
// @Calls: xRichEdit_EnableReadOnly
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XRichEditEnableReadOnly(hEle HELE, bEnable bool) {
	xRichEdit_EnableReadOnly.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:05:11
// @Function: XRichEditEnableMultiLine
// @Description: 启用多行模式.
// @Calls: xRichEdit_EnableMultiLine
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XRichEditEnableMultiLine(hEle HELE, bEnable bool) {
	xRichEdit_EnableMultiLine.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:05:58
// @Function: XRichEditEnablePassword
// @Description: 启用密码模式.
// @Calls: xRichEdit_EnablePassword
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XRichEditEnablePassword(hEle HELE, bEnable bool) {
	xRichEdit_EnablePassword.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:06:41
// @Function: XRichEditEnableEvent_XE_RICHEDIT_CHANGE
// @Description: 启用XE_RICHEDIT_CHANGE事件.
// @Calls: xRichEdit_EnableEvent_XE_RICHEDIT_CHANGE
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XRichEditEnableEvent_XE_RICHEDIT_CHANGE(hEle HELE, bEnable bool) {
	xRichEdit_EnableEvent_XE_RICHEDIT_CHANGE.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:07:27
// @Function: XRichEditEnableAutoWrap
// @Description: 启用自动换行.
// @Calls: xRichEdit_EnableAutoWrap
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XRichEditEnableAutoWrap(hEle HELE, bEnable bool) {
	xRichEdit_EnableAutoWrap.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:08:12
// @Function: XRichEditEnableAutoCancelSel
// @Description: 当失去焦点时,自动取消选择内容.
// @Calls: xRichEdit_EnableAutoCancelSel
// @Input: hEle 元素句柄. bEnable 是否启用.
// @Return:
// *******************************************
func XRichEditEnableAutoCancelSel(hEle HELE, bEnable bool) {
	xRichEdit_EnableAutoCancelSel.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:09:01
// @Function: XRichEditGetTextLength
// @Description: 获取文本内容长度.
// @Calls: xRichEdit_GetTextLength
// @Input: hEle 元素句柄.
// @Return: 长度,不包含空终止符.
// *******************************************
func XRichEditGetTextLength(hEle HELE) int {
	ret, _, _ := xRichEdit_GetTextLength.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:10:09
// @Function: XRichEditSetRowHeight
// @Description: 设置默认行高.
// @Calls: xRichEdit_SetRowHeight
// @Input: hEle 元素句柄. nHeight 高度.
// @Return:
// *******************************************
func XRichEditSetRowHeight(hEle HELE, nHeight uint) {
	xRichEdit_SetRowHeight.Call(
		uintptr(hEle),
		uintptr(nHeight))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:12:09
// @Function: XRichEditGetCurrentRow
// @Description: 获取当前行索引.
// @Calls: xRichEdit_GetCurrentRow
// @Input: hEle 元素句柄.
// @Return: 当前行索引.
// *******************************************
func XRichEditGetCurrentRow(hEle HELE) int {
	ret, _, _ := xRichEdit_GetCurrentRow.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:15:58
// @Function: XRichEditGetCurrentColumn
// @Description: 获取当前列索引.
// @Calls: xRichEdit_GetCurrentColumn
// @Input: hEle 元素句柄.
// @Return: 当前列索引.
// *******************************************
func XRichEditGetCurrentColumn(hEle HELE) int {
	ret, _, _ := xRichEdit_GetCurrentColumn.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:16:51
// @Function: XRichEditSetCurrentPos
// @Description: 设置当前输入位置.
// @Calls: xRichEdit_SetCurrentPos
// @Input: hEle 元素句柄. iRow 行索引,-1代表末尾行. iColumn 列索引,-1代表末尾列.
// @Return:
// *******************************************
func XRichEditSetCurrentPos(hEle HELE, iColumn, iRow int) {
	xRichEdit_SetCurrentPos.Call(
		uintptr(hEle),
		uintptr(iColumn),
		uintptr(iRow))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:18:41
// @Function: XRichEditGetRowCount
// @Description: 获取行数量.
// @Calls: xRichEdit_GetRowCount
// @Input: hEle 元素句柄.
// @Return: 行数量.
// *******************************************
func XRichEditGetRowCount(hEle HELE) int {
	ret, _, _ := xRichEdit_GetRowCount.Call(uintptr(hEle))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:20:09
// @Function: XRichEditGetRowLength
// @Description: 获取指定行长度.
// @Calls: xRichEdit_GetRowLength
// @Input: hEle 元素句柄. iRow 行索引.
// @Return: 行长度.
// *******************************************
func XRichEditGetRowLength(hEle HELE, iRow int) int {
	ret, _, _ := xRichEdit_GetRowLength.Call(
		uintptr(hEle),
		uintptr(iRow))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:21:40
// @Function: XRichEditGetSelectText
// @Description: 获取选择文本.
// @Calls: xRichEdit_GetSelectText
// @Input: hEle 元素句柄. pOut 接收内容缓冲区. len 缓冲区长度,字符为单位.
// @Return: 返回内容长度,字符为单位,不包含空终止符.
// *******************************************
func XRichEditGetSelectText(hEle HELE, pOut *uint16, len int) int {
	ret, _, _ := xRichEdit_GetSelectText.Call(
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
// @Calls: xRichEdit_GetSelectPosition
// @Input: hEle 元素句柄. pBegin 开始位置点. pEnd 结束位置点.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XRichEditGetSelectPosition(hEle HELE, pBegin *Position, pEnd *Position) bool {
	ret, _, _ := xRichEdit_GetSelectPosition.Call(
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
// @Calls: xRichEdit_SetSelect
// @Input: hEle 元素句柄. iStartRow 开始行索引. iStartCol 开始行列索引. iEndRow 结束行索引. iEndCol 结束列索引.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XRichEditSetSelect(hEle HELE, iStartRow, iStartCol, iEndRow, iEndCol int) bool {
	ret, _, _ := xRichEdit_SetSelect.Call(
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
// @Calls: xRichEdit_SetItemFontEx
// @Input: hEle 元素句柄. beginRow 开始行索引. beginColumn 开始行列索引. endRow 结束行索引. endColumn 结束行列索引. pFont 字体信息.
// @Return: 成功返回TRUE,否则相反.
// *******************************************
func XRichEditSetItemFontEx(hEle HELE, beginRow, beginColumn, endRow, endColumn int, pFont *LOGFONTW) bool {
	ret, _, _ := xRichEdit_SetItemFontEx.Call(
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
// @Calls: xRichEdit_SetItemColorEx
// @Input: hEle 元素句柄. beginRow 开始行索引. beginColumn 开始行列索引.
//         endRow 结束行索引. endColumn 结束行列索引. color RGB颜色. alpha 透明度.
// @Return: 成功返回TRUE,否则相反.
// *******************************************
func XRichEditSetItemColorEx(hEle HELE, beginRow, beginColumn, endRow, endColumn int, color COLORREF, alpha byte) bool {
	ret, _, _ := xRichEdit_SetItemColorEx.Call(
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
// @Calls: xRichEdit_CancelSelect
// @Input: hEle 元素句柄.
// @Return:
// *******************************************
func XRichEditCancelSelect(hEle HELE) {
	xRichEdit_CancelSelect.Call(uintptr(hEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:37:46
// @Function: XRichEditSetSelectBkColor
// @Description: 设置选择内容背景颜色.
// @Calls: xRichEdit_SetSelectBkColor
// @Input: hEle 元素句柄. color 颜色值. alpha 透明度.
// @Return:
// *******************************************
func XRichEditSetSelectBkColor(hEle HELE, color COLORREF, alpha byte) {
	xRichEdit_SetSelectBkColor.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:37:46
// @Function: XRichEditIsEmpty
// @Description: 内容是否为空.
// @Calls: xRichEdit_IsEmpty
// @Input: hEle 元素句柄.
// @Return: 如果为空返回TRUE否则返回FALSE.
// *******************************************
func XRichEditIsEmpty(hEle HELE) bool {
	ret, _, _ := xRichEdit_IsEmpty.Call(uintptr(hEle))

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
// @Calls: xRichEdit_SelectAll
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XRichEditSelectAll(hEle HELE) bool {
	ret, _, _ := xRichEdit_SelectAll.Call(uintptr(hEle))

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
// @Calls: xRichEdit_DeleteSelect
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XRichEditDeleteSelect(hEle HELE) bool {
	ret, _, _ := xRichEdit_DeleteSelect.Call(uintptr(hEle))

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
// @Calls: xRichEdit_DeleteAll
// @Input: hEle 元素句柄.
// @Return:
// *******************************************
func XRichEditDeleteAll(hEle HELE) {
	xRichEdit_DeleteAll.Call(uintptr(hEle))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-25 21:49:55
// @Function: XRichEditClipboardCut
// @Description: 剪切选择内容到剪贴板.
// @Calls: xRichEdit_ClipboardCut
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XRichEditClipboardCut(hEle HELE) bool {
	ret, _, _ := xRichEdit_ClipboardCut.Call(uintptr(hEle))

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
// @Calls: xRichEdit_ClipboardCopy
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XRichEditClipboardCopy(hEle HELE) bool {
	ret, _, _ := xRichEdit_ClipboardCopy.Call(uintptr(hEle))

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
// @Calls: xRichEdit_ClipboardPaste
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XRichEditClipboardPaste(hEle HELE) bool {
	ret, _, _ := xRichEdit_ClipboardPaste.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}
