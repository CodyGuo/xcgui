package wke

import (
	"syscall"
)

import (
	"github.com/codyguo/xcgui/xc"
)

var (
	// Functions
	xWeb_Create             *syscall.Proc
	xWeb_MoveWindow         *syscall.Proc
	xWeb_SetProxy           *syscall.Proc
	xWeb_LoadUrl            *syscall.Proc
	xWeb_PostUrl            *syscall.Proc
	xWeb_LoadHtmlFromText   *syscall.Proc
	xWeb_RunJs              *syscall.Proc
	xWeb_Zoom               *syscall.Proc
	xWeb_GetZoom            *syscall.Proc
	xWeb_ZoomReset          *syscall.Proc
	xWeb_SetEditable        *syscall.Proc
	xWeb_GetStringA         *syscall.Proc
	xWeb_GetStringW         *syscall.Proc
	xWeb_SetStringA         *syscall.Proc
	xWeb_SetStringW         *syscall.Proc
	xWeb_GetCookie          *syscall.Proc
	xWeb_SetCookieEnabled   *syscall.Proc
	xWeb_IsCookieEnabled    *syscall.Proc
	xWeb_GoBack             *syscall.Proc
	xWeb_GoForward          *syscall.Proc
	xWeb_EditorSelectAll    *syscall.Proc
	xWeb_EditorCopy         *syscall.Proc
	xWeb_EditorCut          *syscall.Proc
	xWeb_EditorPaste        *syscall.Proc
	xWeb_EditorDelete       *syscall.Proc
	xWeb_CanGoForward       *syscall.Proc
	xWeb_CanGoBack          *syscall.Proc
	xWeb_OnTitleChanged     *syscall.Proc
	xWeb_OnURLChanged       *syscall.Proc
	xWeb_OnNavigation       *syscall.Proc
	xWeb_IsLoadingSucceeded *syscall.Proc
	xWeb_IsLoadingFailed    *syscall.Proc
	xWeb_IsLoadingCompleted *syscall.Proc
	xWeb_IsDocumentReady    *syscall.Proc
	xWeb_StopLoading        *syscall.Proc
	xWeb_Reload             *syscall.Proc
	xWeb_JsBindFuction      *syscall.Proc
	xWeb_JsBindGetter       *syscall.Proc
	xWeb_JsBindSetter       *syscall.Proc
	xWeb_JsToInt            *syscall.Proc
	xWeb_jsToFloat          *syscall.Proc
	xWeb_jsToDouble         *syscall.Proc
	xWeb_jsToBoolean        *syscall.Proc
	xWeb_JsToTempStringW    *syscall.Proc
	xWeb_JsToTempStringA    *syscall.Proc
	xWeb_JsArg              *syscall.Proc
	xWeb_JsInt              *syscall.Proc
	xWeb_JsDouble           *syscall.Proc
	xWeb_JsBoolean          *syscall.Proc
	xWeb_JsUndefined        *syscall.Proc
	xWeb_JsNull             *syscall.Proc
	xWeb_JsTrue             *syscall.Proc
	xWeb_JsFalse            *syscall.Proc
)

func init() {
	// Functions
	xWeb_Create = wkeWebDLL.MustFindProc("XWeb_Create")
	xWeb_MoveWindow = wkeWebDLL.MustFindProc("XWeb_MoveWindow")
	xWeb_SetProxy = wkeWebDLL.MustFindProc("XWeb_SetProxy")
	xWeb_LoadUrl = wkeWebDLL.MustFindProc("XWeb_LoadUrl")
	xWeb_PostUrl = wkeWebDLL.MustFindProc("XWeb_PostUrl")
	xWeb_LoadHtmlFromText = wkeWebDLL.MustFindProc("XWeb_LoadHtmlFromText")
	xWeb_RunJs = wkeWebDLL.MustFindProc("XWeb_RunJs")
	xWeb_Zoom = wkeWebDLL.MustFindProc("XWeb_Zoom")
	xWeb_GetZoom = wkeWebDLL.MustFindProc("XWeb_GetZoom")
	xWeb_ZoomReset = wkeWebDLL.MustFindProc("XWeb_ZoomReset")
	xWeb_SetEditable = wkeWebDLL.MustFindProc("XWeb_SetEditable")
	xWeb_GetStringA = wkeWebDLL.MustFindProc("XWeb_GetStringA")
	xWeb_GetStringW = wkeWebDLL.MustFindProc("XWeb_GetStringW")
	xWeb_SetStringA = wkeWebDLL.MustFindProc("XWeb_SetStringA")
	xWeb_SetStringW = wkeWebDLL.MustFindProc("XWeb_SetStringW")
	xWeb_GetCookie = wkeWebDLL.MustFindProc("XWeb_GetCookie")
	xWeb_SetCookieEnabled = wkeWebDLL.MustFindProc("XWeb_SetCookieEnabled")
	xWeb_IsCookieEnabled = wkeWebDLL.MustFindProc("XWeb_IsCookieEnabled")
	xWeb_GoBack = wkeWebDLL.MustFindProc("XWeb_GoBack")
	xWeb_GoForward = wkeWebDLL.MustFindProc("XWeb_GoForward")
	xWeb_EditorSelectAll = wkeWebDLL.MustFindProc("XWeb_EditorSelectAll")
	xWeb_EditorCopy = wkeWebDLL.MustFindProc("XWeb_EditorCopy")
	xWeb_EditorCut = wkeWebDLL.MustFindProc("XWeb_EditorCut")
	xWeb_EditorPaste = wkeWebDLL.MustFindProc("XWeb_EditorPaste")
	xWeb_EditorDelete = wkeWebDLL.MustFindProc("XWeb_EditorDelete")
	xWeb_CanGoForward = wkeWebDLL.MustFindProc("XWeb_CanGoForward")
	xWeb_CanGoBack = wkeWebDLL.MustFindProc("XWeb_CanGoBack")
	xWeb_OnTitleChanged = wkeWebDLL.MustFindProc("XWeb_OnTitleChanged")
	xWeb_OnURLChanged = wkeWebDLL.MustFindProc("XWeb_OnURLChanged")
	xWeb_OnNavigation = wkeWebDLL.MustFindProc("XWeb_OnNavigation")
	xWeb_IsLoadingSucceeded = wkeWebDLL.MustFindProc("XWeb_IsLoadingSucceeded")
	xWeb_IsLoadingFailed = wkeWebDLL.MustFindProc("XWeb_IsLoadingFailed")
	xWeb_IsLoadingCompleted = wkeWebDLL.MustFindProc("XWeb_IsLoadingCompleted")
	xWeb_IsDocumentReady = wkeWebDLL.MustFindProc("XWeb_IsDocumentReady")
	xWeb_StopLoading = wkeWebDLL.MustFindProc("XWeb_StopLoading")
	xWeb_Reload = wkeWebDLL.MustFindProc("XWeb_Reload")
	xWeb_JsBindFuction = wkeWebDLL.MustFindProc("XWeb_JsBindFuction")
	xWeb_JsBindGetter = wkeWebDLL.MustFindProc("XWeb_JsBindGetter")
	xWeb_JsBindSetter = wkeWebDLL.MustFindProc("XWeb_JsBindSetter")
	xWeb_JsToInt = wkeWebDLL.MustFindProc("XWeb_JsToInt")
	xWeb_jsToFloat = wkeWebDLL.MustFindProc("XWeb_jsToFloat")
	xWeb_jsToDouble = wkeWebDLL.MustFindProc("XWeb_jsToDouble")
	xWeb_jsToBoolean = wkeWebDLL.MustFindProc("XWeb_jsToBoolean")
	xWeb_JsToTempStringW = wkeWebDLL.MustFindProc("XWeb_JsToTempStringW")
	xWeb_JsToTempStringA = wkeWebDLL.MustFindProc("XWeb_JsToTempStringA")
	xWeb_JsArg = wkeWebDLL.MustFindProc("XWeb_JsArg")
	xWeb_JsInt = wkeWebDLL.MustFindProc("XWeb_JsInt")
	xWeb_JsDouble = wkeWebDLL.MustFindProc("XWeb_JsDouble")
	xWeb_JsBoolean = wkeWebDLL.MustFindProc("XWeb_JsBoolean")
	xWeb_JsUndefined = wkeWebDLL.MustFindProc("XWeb_JsUndefined")
	xWeb_JsNull = wkeWebDLL.MustFindProc("XWeb_JsNull")
	xWeb_JsTrue = wkeWebDLL.MustFindProc("XWeb_JsTrue")
	xWeb_JsFalse = wkeWebDLL.MustFindProc("XWeb_JsFalse")
}

func XWebCreate(x, y, cx, cy int, hParent xc.HWND) xc.HELE {
	ret, _, _ := xWeb_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return xc.HELE(ret)
}

func XWebMoveWindow(hWeb xc.HELE, x, y, cx, cy int) {
	xWeb_MoveWindow.Call(
		uintptr(hWeb),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy))
}

func XWebSetProxy(hWeb xc.HELE, proxyType WkeProxyType, hostName string, port uint32, username, pwd string) {
	xWeb_SetProxy.Call(
		uintptr(hWeb),
		uintptr(proxyType),
		xc.StringToUintPtr(hostName),
		uintptr(port),
		xc.StringToUintPtr(username),
		xc.StringToUintPtr(pwd))
}

func XWebLoadUrl(hWeb xc.HELE, url string) {
	xWeb_LoadUrl.Call(
		uintptr(hWeb),
		xc.StringToUintPtr(url))
}

func XWebPostUrl(hWeb xc.HELE, url, postData string, postLen int) {
	xWeb_PostUrl.Call(
		uintptr(hWeb),
		xc.StringToUintPtr(url),
		uintptr(postLen))
}

func XWebLoadHtmlFromText(hWeb xc.HELE, htmlData string) {
	xWeb_LoadHtmlFromText.Call(
		uintptr(hWeb),
		xc.StringToUintPtr(htmlData))
}

func XWebRunJs(hWeb xc.HELE, jsText string) {
	xWeb_RunJs.Call(
		uintptr(hWeb),
		xc.StringToUintPtr(jsText))
}

func XWebZoom(hWeb xc.HELE, f float32) {
	xWeb_Zoom.Call(
		uintptr(hWeb),
		uintptr(f))
}

func XWebGetZoom(hWeb xc.HELE) float32 {
	ret, _, _ := xWeb_GetZoom.Call(uintptr(hWeb))

	return float32(ret)
}

func XWebZoomReset(hWeb xc.HELE) {
	xWeb_ZoomReset.Call(uintptr(hWeb))
}

func XWebSetEditable(hWeb xc.HELE, editable bool) {
	xWeb_SetEditable.Call(
		uintptr(hWeb),
		uintptr(xc.BoolToBOOL(editable)))
}

func XWebGetStringA(stringA uintptr) string {
	ret, _, _ := xWeb_GetStringA.Call(stringA)

	return string(ret)
}

func XWebGetStringW(stringW uintptr) string {
	ret, _, _ := xWeb_GetStringW.Call(stringW)

	return xc.UINTptrToString(ret)
}

/*
http://baike.baidu.com/link?url=gFiwJdUF0mAgY-pLBvflyvwAdr5mXESqBBj7S_pTg2UDaB_w6Ees6O8fQJ3pUQwktt4ADu80z5eW9G4t5eHV4_
*/
func XWebSetStringA(stringA uintptr, str string, lenA int) {
	xWeb_SetStringA.Call(
		stringA,
		xc.StringToUintPtr(str),
		uintptr(lenA))
}

func XWebSetStringW(stringW uintptr, str string, lenW int) {
	xWeb_SetStringW.Call(
		stringW,
		xc.StringToUintPtr(str),
		uintptr(lenW))
}

func XWebGetCookie(hWeb xc.HELE) string {
	ret, _, _ := xWeb_GetCookie.Call(uintptr(hWeb))

	return xc.UINTptrToString(ret)
}

func XWebSetCookieEnabled(hWeb xc.HELE, enable bool) {
	xWeb_SetCookieEnabled.Call(
		uintptr(hWeb),
		uintptr(xc.BoolToBOOL(enable)))
}

func XWebIsCookieEnabled(hWeb xc.HELE) bool {
	ret, _, _ := xWeb_IsCookieEnabled.Call(uintptr(hWeb))

	return ret == xc.TRUE
}

func XWebGoBack(hWeb xc.HELE) {
	xWeb_GoBack.Call(uintptr(hWeb))
}

func XWebGoForward(hWeb xc.HELE) {
	xWeb_GoForward.Call(uintptr(hWeb))
}

func XWebEditorSelectAll(hWeb xc.HELE) {
	xWeb_EditorSelectAll.Call(uintptr(hWeb))
}

func XWebEditorCopy(hWeb xc.HELE) {
	xWeb_EditorCopy.Call(uintptr(hWeb))
}

func XWebEditorCut(hWeb xc.HELE) {
	xWeb_EditorCut.Call(uintptr(hWeb))
}

func XWebEditorPaste(hWeb xc.HELE) {
	xWeb_EditorPaste.Call(uintptr(hWeb))
}

func XWebEditorDelete(hWeb xc.HELE) {
	xWeb_EditorDelete.Call(uintptr(hWeb))
}

func XWebCanGoForward(hWeb xc.HELE) bool {
	ret, _, _ := xWeb_CanGoForward.Call(uintptr(hWeb))

	return ret == xc.TRUE
}

func XWebCanGoBack(hWeb xc.HELE) bool {
	ret, _, _ := xWeb_CanGoBack.Call(uintptr(hWeb))

	return ret == xc.TRUE
}

func XWebOnTitleChanged(hWeb xc.HELE, pFunc, param uintptr) {
	xWeb_OnTitleChanged.Call(
		uintptr(hWeb),
		pFunc,
		param)
}

func XWebOnURLChanged(hWeb xc.HELE, pFunc, param uintptr) {
	xWeb_OnURLChanged.Call(
		uintptr(hWeb),
		pFunc,
		param)
}

func XWebOnNavigation(hWeb xc.HELE, pFunc, param uintptr) {
	xWeb_OnNavigation.Call(
		uintptr(hWeb),
		pFunc,
		param)
}

func XWebIsLoadingSucceeded(hWeb xc.HELE) bool {
	ret, _, _ := xWeb_IsLoadingSucceeded.Call(uintptr(hWeb))

	return ret == xc.TRUE
}

func XWebIsLoadingFailed(hWeb xc.HELE) bool {
	ret, _, _ := xWeb_IsLoadingFailed.Call(uintptr(hWeb))

	return ret == xc.TRUE
}

func XWebIsLoadingCompleted(hWeb xc.HELE) bool {
	ret, _, _ := xWeb_IsLoadingCompleted.Call(uintptr(hWeb))

	return ret == xc.TRUE
}

func XWebIsDocumentReady(hWeb xc.HELE) bool {
	ret, _, _ := xWeb_IsDocumentReady.Call(uintptr(hWeb))

	return ret == xc.TRUE
}

func XWebStopLoading(hWeb xc.HELE) {
	xWeb_StopLoading.Call(uintptr(hWeb))
}

func XWebReload(hWeb xc.HELE) {
	xWeb_Reload.Call(uintptr(hWeb))
}

func XWebJsBindFuction(funcName string, pFunc uintptr, argCount uint32) {
	xWeb_JsBindFuction.Call(
		xc.StringToUintPtr(funcName),
		pFunc,
		uintptr(argCount))
}

func XWebJsBindGetter(funcName string, pFunc uintptr) {
	xWeb_JsBindGetter.Call(
		xc.StringToUintPtr(funcName),
		pFunc)
}

func XWebJsBindSetter(funcName string, pFunc uintptr) {
	xWeb_JsBindSetter.Call(
		xc.StringToUintPtr(funcName),
		pFunc)
}

func XWebJsToInt(es uintptr, value int64) int {
	ret, _, _ := xWeb_JsToInt.Call(
		es,
		uintptr(value))

	return int(ret)
}

func XWebjsToFloat(es uintptr, value int64) float32 {
	ret, _, _ := xWeb_jsToFloat.Call(
		es,
		uintptr(value))

	return float32(ret)
}

func XWebjsToDouble(es uintptr, value int64) float64 {
	ret, _, _ := xWeb_jsToDouble.Call(
		es,
		uintptr(value))

	return float64(ret)
}

func XWebjsToBoolean(es uintptr, value int64) bool {
	ret, _, _ := xWeb_jsToBoolean.Call(
		es,
		uintptr(value))

	return ret == xc.TRUE
}

func XWebJsToTempStringW(es uintptr, value int64) string {
	ret, _, _ := xWeb_JsToTempStringW.Call(
		es,
		uintptr(value))

	return xc.UINTptrToString(ret)
}

func XWebJsToTempStringA(es uintptr, value int64) string {
	ret, _, _ := xWeb_JsToTempStringA.Call(
		es,
		uintptr(value))

	return xc.UINTptrToString(ret)
}

func XWebJsArg(es uintptr, argIdx int) int64 {
	ret, _, _ := xWeb_JsArg.Call(
		es,
		uintptr(argIdx))

	return int64(ret)
}

func XWebJsInt(n int) int64 {
	ret, _, _ := xWeb_JsInt.Call(uintptr(n))

	return int64(ret)
}

func XWebJsDouble(d float32) int64 {
	ret, _, _ := xWeb_JsDouble.Call(uintptr(d))

	return int64(ret)
}

func XWebJsBoolean(b bool) int64 {
	ret, _, _ := xWeb_JsBoolean.Call(uintptr(xc.BoolToBOOL(b)))

	return int64(ret)
}

func XWebJsUndefined() int64 {
	ret, _, _ := xWeb_JsUndefined.Call()

	return int64(ret)
}

func XWebJsNull() int64 {
	ret, _, _ := xWeb_JsNull.Call()

	return int64(ret)
}

func XWebJsTrue() int64 {
	ret, _, _ := xWeb_JsTrue.Call()

	return int64(ret)
}

func XWebJsFalse() int64 {
	ret, _, _ := xWeb_JsFalse.Call()

	return int64(ret)
}
