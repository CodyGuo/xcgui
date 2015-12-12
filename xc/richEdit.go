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
	xRichEdit_EnableAutoSelAll               *syscall.Proc
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
	xRichEdit_EnableAutoSelAll = xcDLL.MustFindProc("XRichEdit_EnableAutoSelAll")
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

/*
创建富文本编辑框.

参数:
	x 元素x坐标.
	y 元素y坐标.
	cx 宽度.
	cy 高度.
	hParent 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
返回:
	元素句柄.
*/
func XRichEditCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xRichEdit_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
插入字符串到当前位置.

参数:
	hEle 元素句柄.
	pString 字符串指针. *uint16
	pFont 字体.
	color 颜色.
*/
func XRichEditInsertString(hEle HELE, pString string, pFont *LOGFONTW, color COLORREF) {
	xRichEdit_InsertString.Call(
		uintptr(hEle),
		StringToUintPtr(pString),
		// uintptr(unsafe.Pointer(pString)),
		uintptr(unsafe.Pointer(pFont)),
		uintptr(color))
}

/*
插入图片到当前位置.

参数:
	hEle 元素句柄.
	hImage 图片句柄.
	pImagePath 图片路径,用于图片的复制.*uint16
返回:
	成功返回TRUE,否则相反.
*/
func XRichEditInsertImage(hEle HELE, hImage HIMAGE, pImagePath string) bool {
	ret, _, _ := xRichEdit_InsertImage.Call(
		uintptr(hEle),
		uintptr(hImage),
		StringToUintPtr(pImagePath))
	// uintptr(unsafe.Pointer(pImagePath)))

	return ret == TRUE
}

/*
插入GIF动画图片到当前位置.

参数:
	hEle 元素句柄.
	hImage 图片句柄.
	pImagePath 图片路径,用于图片的复制.*uint16
返回:
	成功返回TRUE,否则相反.
*/
func XRichEditInsertGif(hEle HELE, hImage HIMAGE, pImagePath string) bool {
	ret, _, _ := xRichEdit_InsertGif.Call(
		uintptr(hEle),
		uintptr(hImage),
		StringToUintPtr(pImagePath))
	// uintptr(unsafe.Pointer(pImagePath)))

	return ret == TRUE
}

/*
插入字符串.

参数:
	hEle 元素句柄.
	iRow 行索引.
	iColumn 列索引.
	pString 字符串指针.*uint16
	pFont 字体.
	color 颜色.
*/
func XRichEditInsertStringEx(hEle HELE, iRow, iColumn int, pString string, pFont *LOGFONTW, color COLORREF) {
	xRichEdit_InsertStringEx.Call(
		uintptr(hEle),
		uintptr(iRow),
		uintptr(iColumn),
		StringToUintPtr(pString),
		// uintptr(unsafe.Pointer(pString)),
		uintptr(unsafe.Pointer(pFont)),
		uintptr(color))
}

/*
插入图片.

参数:
	hEle 元素句柄.
	iRow 插入的行位置,如果值为-1插入到末尾行.
	iColumn 插入的列位置,如果值为-1插入末尾列.
	hImage 图片句柄.
	pImagePath 图片路径,用于图片的复制.*uint16
返回:
	成功返回TRUE,否则相反.
*/
func XRichEditInsertImageEx(hEle HELE, iRow, iColumn int, hImage HIMAGE, pImagePath string) bool {
	ret, _, _ := xRichEdit_InsertImageEx.Call(
		uintptr(hEle),
		uintptr(iRow),
		uintptr(iColumn),
		uintptr(hImage),
		StringToUintPtr(pImagePath))
	// uintptr(unsafe.Pointer(pImagePath)))

	return ret == TRUE
}

/*
插入GIF动画图片.

参数:
	hEle 元素句柄.
	iRow 插入的行位置,如果值为-1插入到末尾行.
	iColumn 插入的列位置,如果值为-1插入末尾列.
	hImage 图片句柄.
	pImagePath 图片路径,用于图片的复制.*uint16
返回:
	成功返回TRUE,否则相反.
*/
func XRichEditInsertGifEx(hEle HELE, iRow, iColumn int, hImage HIMAGE, pImagePath string) bool {
	ret, _, _ := xRichEdit_InsertGifEx.Call(
		uintptr(hEle),
		uintptr(iRow),
		uintptr(iColumn),
		uintptr(hImage),
		StringToUintPtr(pImagePath))
	// uintptr(unsafe.Pointer(pImagePath)))

	return ret == TRUE
}

/*
设置文本内容.

参数:
	hEle 元素句柄.
	pString 字符串指针.*uint16
*/
func XRichEditSetText(hEle HELE, pString string) {
	xRichEdit_SetText.Call(
		uintptr(hEle),
		StringToUintPtr(pString))
	// uintptr(unsafe.Pointer(pString)))
}

/*
设置内容.

参数:
	hEle 元素句柄.
	nVaule 整型值.
*/
func XRichEditSetTextInt(hEle HELE, nVaule int) {
	xRichEdit_SetTextInt.Call(
		uintptr(hEle),
		uintptr(nVaule))
}

/*
获取文本内容.

参数:
	hEle 元素句柄.
	pOut 接收内容缓冲区.
	len 缓冲区长度,字符为单位.
返回:
	返回内容长度,不包含空终止符.
*/
func XRichEditGetText(hEle HELE, pOut *uint16, len int) int {
	ret, _, _ := xRichEdit_GetText.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(len))

	return int(ret)
}

/*
获取HTML格式内容.

参数:
	hEle 元素句柄.
	pOut 接收内容缓冲区.
	len 缓冲区长度,字符为单位
*/
func XRichEditGetHTMLFormat(hEle HELE, pOut *uint16, len int) {
	xRichEdit_GetHTMLFormat.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(len))

}

/*
获取富文本数据,包含字体,图片,颜色信息.

参数:
	hEle 元素句柄.
	pDataSize 返回数据大小,字节为单位.
返回:
	返回数据内存指针,不用之后需要释放内容XC_Free().
*/
func XRichEditGetData(hEle HELE, pDataSize *int) {
	xRichEdit_GetData.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pDataSize)))
}

/*
插入数据内容,包含字体,图片,颜色信息.

参数:
	hEle 元素句柄.
	pData 数据指针.
	iRow 插入的行位置,如果值为-1插入到末尾行.
	iColumn 插入的列位置,如果值为-1插入末尾列.
返回:
	成功返回TRUE,否则相反.
*/
func XRichEditInsertData(hEle HELE, pData uintptr, iRow, iColumn int) bool {
	ret, _, _ := xRichEdit_InsertData.Call(
		uintptr(hEle),
		pData,
		uintptr(iRow),
		uintptr(iColumn))

	return ret == TRUE
}

/*
启用只读模式.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XRichEditEnableReadOnly(hEle HELE, bEnable bool) {
	xRichEdit_EnableReadOnly.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用多行模式.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XRichEditEnableMultiLine(hEle HELE, bEnable bool) {
	xRichEdit_EnableMultiLine.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用密码模式.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XRichEditEnablePassword(hEle HELE, bEnable bool) {
	xRichEdit_EnablePassword.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用XE_RICHEDIT_CHANGE事件.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XRichEditEnableEvent_XE_RICHEDIT_CHANGE(hEle HELE, bEnable bool) {
	xRichEdit_EnableEvent_XE_RICHEDIT_CHANGE.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用自动换行.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XRichEditEnableAutoWrap(hEle HELE, bEnable bool) {
	xRichEdit_EnableAutoWrap.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
当失去焦点时,自动取消选择内容.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XRichEditEnableAutoCancelSel(hEle HELE, bEnable bool) {
	xRichEdit_EnableAutoCancelSel.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
当获得焦点时,自动选择所有内容.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XRichEditEnableAutoSelAll(hEle HELE, bEnable bool) {
	xRichEdit_EnableAutoSelAll.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
获取文本内容长度.

参数:
	hEle 元素句柄.
返回:
	长度,不包含空终止符.
*/
func XRichEditGetTextLength(hEle HELE) int {
	ret, _, _ := xRichEdit_GetTextLength.Call(uintptr(hEle))

	return int(ret)
}

/*
设置默认行高.

参数:
	hEle 元素句柄.
	nHeight 高度.
*/
func XRichEditSetRowHeight(hEle HELE, nHeight uint) {
	xRichEdit_SetRowHeight.Call(
		uintptr(hEle),
		uintptr(nHeight))
}

/*
获取当前行索引.

参数:
	hEle 元素句柄.
返回:
	当前行索引.
*/
func XRichEditGetCurrentRow(hEle HELE) int {
	ret, _, _ := xRichEdit_GetCurrentRow.Call(uintptr(hEle))

	return int(ret)
}

/*
获取当前列索引.

参数:
	hEle 元素句柄.
返回:
	当前列索引.
*/
func XRichEditGetCurrentColumn(hEle HELE) int {
	ret, _, _ := xRichEdit_GetCurrentColumn.Call(uintptr(hEle))

	return int(ret)
}

/*
设置当前输入位置.

参数:
	hEle 元素句柄.
	iRow 行索引,-1代表末尾行.
	iColumn 列索引,-1代表末尾列.
*/
func XRichEditSetCurrentPos(hEle HELE, iColumn, iRow int) {
	xRichEdit_SetCurrentPos.Call(
		uintptr(hEle),
		uintptr(iColumn),
		uintptr(iRow))
}

/*
获取行数量.

参数:
	hEle 元素句柄.
返回:
	行数量.
*/
func XRichEditGetRowCount(hEle HELE) int {
	ret, _, _ := xRichEdit_GetRowCount.Call(uintptr(hEle))

	return int(ret)
}

/*
获取指定行长度.

参数:
	hEle 元素句柄.
	iRow 行索引.
返回:
	行长度.
*/
func XRichEditGetRowLength(hEle HELE, iRow int) int {
	ret, _, _ := xRichEdit_GetRowLength.Call(
		uintptr(hEle),
		uintptr(iRow))

	return int(ret)
}

/*
获取选择文本.

参数:
	hEle 元素句柄.
	pOut 接收内容缓冲区.
	len 缓冲区长度,字符为单位.
返回:
	返回内容长度,字符为单位,不包含空终止符.
*/
func XRichEditGetSelectText(hEle HELE, pOut *uint16, reLen int) int {
	ret, _, _ := xRichEdit_GetSelectText.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(reLen))

	return int(ret)
}

/*
获取选择内容位置.

参数:
	hEle 元素句柄.
	pBegin 开始位置点.
	pEnd 结束位置点.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XRichEditGetSelectPosition(hEle HELE, pBegin, pEnd *Position) bool {
	ret, _, _ := xRichEdit_GetSelectPosition.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pBegin)),
		uintptr(unsafe.Pointer(pEnd)))

	return ret == TRUE
}

/*
设置选择内容.

参数:
	hEle 元素句柄.
	iStartRow 开始行索引.
	iStartCol 开始行列索引.
	iEndRow 结束行索引.
	iEndCol 结束列索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XRichEditSetSelect(hEle HELE, iStartRow, iStartCol, iEndRow, iEndCol int) bool {
	ret, _, _ := xRichEdit_SetSelect.Call(
		uintptr(hEle),
		uintptr(iStartRow),
		uintptr(iStartCol),
		uintptr(iEndRow),
		uintptr(iEndCol))

	return ret == TRUE
}

/*
设置指定内容字体.

参数:
	hEle 元素句柄.
	beginRow 开始行索引.
	beginColumn 开始行列索引.
	endRow 结束行索引.
	endColumn 结束行列索引.
	pFont 字体信息.
返回:
	成功返回TRUE,否则相反.
*/
func XRichEditSetItemFontEx(hEle HELE, beginRow, beginColumn, endRow, endColumn int, pFont *LOGFONTW) bool {
	ret, _, _ := xRichEdit_SetItemFontEx.Call(
		uintptr(hEle),
		uintptr(beginRow),
		uintptr(beginColumn),
		uintptr(endRow),
		uintptr(endColumn),
		uintptr(unsafe.Pointer(pFont)))

	return ret == TRUE
}

/*
设置指定内容字体颜色.

参数:
	hEle 元素句柄.
	beginRow 开始行索引.
	beginColumn 开始行列索引.
	endRow 结束行索引.
	endColumn 结束行列索引.
	color RGB颜色.
	alpha 透明度..
返回:
	成功返回TRUE,否则相反.
*/
func XRichEditSetItemColorEx(hEle HELE, beginRow, beginColumn, endRow, endColumn int, color COLORREF, alpha byte) bool {
	ret, _, _ := xRichEdit_SetItemColorEx.Call(
		uintptr(hEle),
		uintptr(beginRow),
		uintptr(beginColumn),
		uintptr(endRow),
		uintptr(endColumn),
		uintptr(color),
		uintptr(alpha))

	return ret == TRUE
}

/*
取消选择.

参数:
	hEle 元素句柄.
*/
func XRichEditCancelSelect(hEle HELE) {
	xRichEdit_CancelSelect.Call(uintptr(hEle))
}

/*
设置选择内容背景颜色.

参数:
	hEle 元素句柄.
	color 颜色值.
	alpha 透明度.
*/
func XRichEditSetSelectBkColor(hEle HELE, color COLORREF, alpha byte) {
	xRichEdit_SetSelectBkColor.Call(
		uintptr(hEle),
		uintptr(color),
		uintptr(alpha))
}

/*
内容是否为空.

参数:
	hEle 元素句柄.
返回:
	如果为空返回TRUE否则返回FALSE.
*/
func XRichEditIsEmpty(hEle HELE) bool {
	ret, _, _ := xRichEdit_IsEmpty.Call(uintptr(hEle))

	return ret == TRUE
}

/*
选择所有内容.

参数:
	hEle 元素句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XRichEditSelectAll(hEle HELE) bool {
	ret, _, _ := xRichEdit_SelectAll.Call(uintptr(hEle))

	return ret == TRUE
}

/*
删除选择内容.

参数:
	hEle 元素句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XRichEditDeleteSelect(hEle HELE) bool {
	ret, _, _ := xRichEdit_DeleteSelect.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

/*
删除所有内容.

参数:
	hEle 元素句柄.
*/
func XRichEditDeleteAll(hEle HELE) {
	xRichEdit_DeleteAll.Call(uintptr(hEle))
}

/*
剪切选择内容到剪贴板.

参数:
	hEle 元素句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XRichEditClipboardCut(hEle HELE) bool {
	ret, _, _ := xRichEdit_ClipboardCut.Call(uintptr(hEle))

	return ret == TRUE
}

/*
复制选择内容到剪贴板.

参数:
	hEle 元素句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XRichEditClipboardCopy(hEle HELE) bool {
	ret, _, _ := xRichEdit_ClipboardCopy.Call(uintptr(hEle))

	return ret == TRUE
}

/*
从剪贴板粘贴内容到当前位置.

参数:
	hEle 元素句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XRichEditClipboardPaste(hEle HELE) bool {
	ret, _, _ := xRichEdit_ClipboardPaste.Call(uintptr(hEle))

	return ret == TRUE
}
