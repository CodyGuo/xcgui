#ifndef wkeBrowser_h__
#define wkeBrowser_h__



typedef char utf8;

typedef enum
{
    WKE_PROXY_NONE,
	WKE_PROXY_HTTP,
	WKE_PROXY_SOCKS4,
	WKE_PROXY_SOCKS4A,
	WKE_PROXY_SOCKS5,
	WKE_PROXY_SOCKS5HOSTNAME
} wkeProxyType;
XC_API unsigned int WINAPI XWeb_GetVersion();
XC_API const utf8* WINAPI XWeb_GetVersionString();
XC_API void* WINAPI XWeb_Create(int x,int y,int cx,int cy,HWND hParent);	
XC_API void WINAPI XWeb_MoveWindow(void* hWeb,int x,int y,int cx,int cy);	
XC_API void WINAPI XWeb_SetUserAgentA(void* hWeb,const utf8* userAgent);	
XC_API void WINAPI XWeb_SetUserAgentW(void* hWeb,const wchar_t* userAgent);	
XC_API void WINAPI XWeb_Sleep(void* hWeb);	
XC_API void WINAPI XWeb_Wake(void* hWeb);	
XC_API bool WINAPI XWeb_IsAwake(void* hWeb);	
XC_API void WINAPI XWeb_SetProxy(void* hWeb,int proxyType = WKE_PROXY_SOCKS5,const char* hostName = "127.0.0.1",unsigned short port = 1080,const char* username = NULL,const char* pwd = NULL);	
XC_API void WINAPI XWeb_LoadUrl(void* hWeb,const wchar_t* url);	
XC_API void WINAPI XWeb_PostUrl(void* hWeb,const wchar_t* url, const char* postData,int postLen);	
XC_API void WINAPI XWeb_LoadHtmlFromText(void* hWeb,const wchar_t* htmlData);	
XC_API __int64 WINAPI XWeb_RunJs(void* hWeb,const wchar_t* jsText);	
XC_API void* WINAPI XWeb_GlobalExec(void* hWeb);	
XC_API void WINAPI XWeb_Zoom(void* hWeb,float f);	
XC_API float WINAPI XWeb_GetZoom(void* hWeb);	
XC_API void WINAPI XWeb_ZoomReset(void* hWeb);	
XC_API void WINAPI XWeb_SetEditable(void* hWeb,bool editable);	
XC_API const utf8* WINAPI XWeb_GetStringA(void* string);	
XC_API const wchar_t* WINAPI XWeb_GetStringW(void* string);	
XC_API void WINAPI XWeb_SetStringA(void* string, const utf8* str, size_t len);	
XC_API void WINAPI XWeb_SetStringW(void* string, const wchar_t* str, size_t len);	
XC_API const wchar_t* WINAPI XWeb_GetCookie(void* hWeb);	
XC_API void WINAPI XWeb_SetCookieEnabled(void* hWeb,bool enable);	
XC_API bool WINAPI XWeb_IsCookieEnabled(void* hWeb);	
XC_API void WINAPI XWeb_GoBack(void* hWeb);	
XC_API void WINAPI XWeb_SetMediaVolume(void* hWeb,float volume);	
XC_API float WINAPI XWeb_GetMediaVolume(void* hWeb);	
XC_API void WINAPI XWeb_GoForward(void* hWeb);	
XC_API void WINAPI XWeb_EditorSelectAll(void* hWeb);	
XC_API void WINAPI XWeb_EditorCopy(void* hWeb);	
XC_API void WINAPI XWeb_EditorCut(void* hWeb);	
XC_API void WINAPI XWeb_EditorPaste(void* hWeb);	
XC_API void WINAPI XWeb_EditorDelete(void* hWeb);	
XC_API bool WINAPI XWeb_CanGoForward(void* hWeb);	
XC_API bool WINAPI XWeb_CanGoBack(void* hWeb);	
XC_API void WINAPI XWeb_OnTitleChanged(void* hWeb,void* callback,void* param = NULL);	
XC_API void WINAPI XWeb_OnURLChanged(void* hWeb,void* callback,void* param = NULL);	
XC_API void WINAPI XWeb_OnNavigation(void* hWeb,void* callback,void* param = NULL);	
XC_API bool WINAPI XWeb_IsLoadingSucceeded(void* hWeb);	
XC_API bool WINAPI XWeb_IsLoadingFailed(void* hWeb);	
XC_API bool WINAPI XWeb_IsLoadingCompleted(void* hWeb);	
XC_API bool WINAPI XWeb_IsDocumentReady(void* hWeb);	
XC_API void WINAPI XWeb_StopLoading(void* hWeb);	
XC_API void WINAPI XWeb_Reload(void* hWeb);	
XC_API void WINAPI XWeb_JsBindFuction(const char* functionName,void* fn,unsigned int argCount);	
XC_API void WINAPI XWeb_JsBindGetter(const char* functionName,void* fn);	
XC_API void WINAPI XWeb_JsBindSetter(const char* functionName,void* fn);	
XC_API int WINAPI XWeb_JsToInt(void* es,__int64 v);	
XC_API float WINAPI XWeb_jsToFloat(void* es,__int64 v);	
XC_API double WINAPI XWeb_jsToDouble(void* es,__int64 v);	
XC_API bool WINAPI XWeb_jsToBoolean(void* es,__int64 v);	
XC_API const wchar_t* WINAPI XWeb_JsToTempStringW(void* es,__int64 v);	
XC_API const utf8* WINAPI XWeb_JsToTempStringA(void* es,__int64 v);	
XC_API __int64 WINAPI XWeb_JsArg(void* es,int argIdx);	
XC_API __int64 WINAPI XWeb_JsInt(int n);	
XC_API __int64 WINAPI XWeb_JsDouble(double d);	
XC_API __int64 WINAPI XWeb_JsBoolean(bool b);	
XC_API __int64 WINAPI XWeb_JsUndefined();
XC_API __int64 WINAPI XWeb_JsNull();
XC_API __int64 WINAPI XWeb_JsTrue();
XC_API __int64 WINAPI XWeb_JsFalse();
XC_API __int64 WINAPI XWeb_JsStringA(void* es,const utf8* str);
XC_API __int64 WINAPI XWeb_jsStringW(void* es, const wchar_t* str);
XC_API __int64 WINAPI XWeb_JsGlobalObject(void* es);	
XC_API __int64 WINAPI XWeb_JsGet(void* es,__int64 object,const char* prop);	
XC_API void WINAPI XWeb_JsSet(void* es,__int64 object,const char* prop,__int64 v);
XC_API __int64 WINAPI XWeb_JsGetAt(void* es,__int64 object,int index);
XC_API void WINAPI XWeb_JsSetAt(void* es,__int64 object,int index,__int64 v);
	



#endif // wkeBrowser_h__