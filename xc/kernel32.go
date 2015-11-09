package xc

import (
    "syscall"
    "unsafe"
)

const MAX_PATH = 260

// Error codes
const (
    ERROR_SUCCESS             = 0
    ERROR_INVALID_FUNCTION    = 1
    ERROR_FILE_NOT_FOUND      = 2
    ERROR_INVALID_PARAMETER   = 87
    ERROR_INSUFFICIENT_BUFFER = 122
    ERROR_MORE_DATA           = 234
)

// GlobalAlloc flags
const (
    GHND          = 0x0042
    GMEM_FIXED    = 0x0000
    GMEM_MOVEABLE = 0x0002
    GMEM_ZEROINIT = 0x0040
    GPTR          = 0x004
)

// Predefined locale ids
const (
    LOCALE_CUSTOM_DEFAULT     LCID = 0x0c00
    LOCALE_CUSTOM_UI_DEFAULT  LCID = 0x1400
    LOCALE_CUSTOM_UNSPECIFIED LCID = 0x1000
    LOCALE_INVARIANT          LCID = 0x007f
    LOCALE_USER_DEFAULT       LCID = 0x0400
    LOCALE_SYSTEM_DEFAULT     LCID = 0x0800
)

// LCTYPE constants
const (
    LOCALE_SDECIMAL          LCTYPE = 14
    LOCALE_STHOUSAND         LCTYPE = 15
    LOCALE_SISO3166CTRYNAME  LCTYPE = 0x5a
    LOCALE_SISO3166CTRYNAME2 LCTYPE = 0x68
    LOCALE_SISO639LANGNAME   LCTYPE = 0x59
    LOCALE_SISO639LANGNAME2  LCTYPE = 0x67
)

var (
    // Library
    libkernel32 *syscall.DLL

    // Functions
    closeHandle            *syscall.Proc
    fileTimeToSystemTime   *syscall.Proc
    getConsoleTitle        *syscall.Proc
    getConsoleWindow       *syscall.Proc
    getLastError           *syscall.Proc
    getLocaleInfo          *syscall.Proc
    getLogicalDriveStrings *syscall.Proc
    getModuleHandle        *syscall.Proc
    getNumberFormat        *syscall.Proc
    getThreadLocale        *syscall.Proc
    getThreadUILanguage    *syscall.Proc
    getVersion             *syscall.Proc
    globalAlloc            *syscall.Proc
    globalFree             *syscall.Proc
    globalLock             *syscall.Proc
    globalUnlock           *syscall.Proc
    moveMemory             *syscall.Proc
    mulDiv                 *syscall.Proc
    setLastError           *syscall.Proc
    systemTimeToFileTime   *syscall.Proc
    getProfileString       *syscall.Proc
)

type (
    ATOM      uint16
    HANDLE    uintptr
    HGLOBAL   HANDLE
    HINSTANCE HANDLE
    LCID      uint32
    LCTYPE    uint32
    LANGID    uint16
)

type FILETIME struct {
    DwLowDateTime  uint32
    DwHighDateTime uint32
}

type NUMBERFMT struct {
    NumDigits     uint32
    LeadingZero   uint32
    Grouping      uint32
    LpDecimalSep  *uint16
    LpThousandSep *uint16
    NegativeOrder uint32
}

type SYSTEMTIME struct {
    WYear         uint16
    WMonth        uint16
    WDayOfWeek    uint16
    WDay          uint16
    WHour         uint16
    WMinute       uint16
    WSecond       uint16
    WMilliseconds uint16
}

func init() {
    // Library
    libkernel32 = syscall.MustLoadDLL("kernel32.dll")

    // Functions
    closeHandle = libkernel32.MustFindProc("CloseHandle")
    fileTimeToSystemTime = libkernel32.MustFindProc("FileTimeToSystemTime")
    getConsoleTitle = libkernel32.MustFindProc("GetConsoleTitleW")
    getConsoleWindow = libkernel32.MustFindProc("GetConsoleWindow")
    getLastError = libkernel32.MustFindProc("GetLastError")
    getLocaleInfo = libkernel32.MustFindProc("GetLocaleInfoW")
    getLogicalDriveStrings = libkernel32.MustFindProc("GetLogicalDriveStringsW")
    getModuleHandle = libkernel32.MustFindProc("GetModuleHandleW")
    getNumberFormat = libkernel32.MustFindProc("GetNumberFormatW")
    getProfileString = libkernel32.MustFindProc("GetProfileStringW")
    getThreadLocale = libkernel32.MustFindProc("GetThreadLocale")
    getThreadUILanguage = libkernel32.MustFindProc("GetThreadUILanguage")
    getVersion = libkernel32.MustFindProc("GetVersion")
    globalAlloc = libkernel32.MustFindProc("GlobalAlloc")
    globalFree = libkernel32.MustFindProc("GlobalFree")
    globalLock = libkernel32.MustFindProc("GlobalLock")
    globalUnlock = libkernel32.MustFindProc("GlobalUnlock")
    moveMemory = libkernel32.MustFindProc("RtlMoveMemory")
    mulDiv = libkernel32.MustFindProc("MulDiv")
    setLastError = libkernel32.MustFindProc("SetLastError")
    systemTimeToFileTime = libkernel32.MustFindProc("SystemTimeToFileTime")

}

func CloseHandle(hObject HANDLE) bool {
    ret, _, _ := closeHandle.Call(
        uintptr(hObject))

    return ret != 0
}

func FileTimeToSystemTime(lpFileTime *FILETIME, lpSystemTime *SYSTEMTIME) bool {
    ret, _, _ := fileTimeToSystemTime.Call(
        uintptr(unsafe.Pointer(lpFileTime)),
        uintptr(unsafe.Pointer(lpSystemTime)))

    return ret != 0
}

func GetConsoleTitle(lpConsoleTitle *uint16, nSize uint32) uint32 {
    ret, _, _ := getConsoleTitle.Call(
        uintptr(unsafe.Pointer(lpConsoleTitle)),
        uintptr(nSize))

    return uint32(ret)
}

func GetConsoleWindow() HWND {
    ret, _, _ := getConsoleWindow.Call()

    return HWND(ret)
}

func GetLastError() uint32 {
    ret, _, _ := getLastError.Call()

    return uint32(ret)
}

func GetLocaleInfo(Locale LCID, LCType LCTYPE, lpLCData *uint16, cchData int32) int32 {
    ret, _, _ := getLocaleInfo.Call(
        uintptr(Locale),
        uintptr(LCType),
        uintptr(unsafe.Pointer(lpLCData)),
        uintptr(cchData))

    return int32(ret)
}

func GetLogicalDriveStrings(nBufferLength uint32, lpBuffer *uint16) uint32 {
    ret, _, _ := getLogicalDriveStrings.Call(
        uintptr(nBufferLength),
        uintptr(unsafe.Pointer(lpBuffer)))

    return uint32(ret)
}

func GetModuleHandle(lpModuleName *uint16) HINSTANCE {
    ret, _, _ := getModuleHandle.Call(
        uintptr(unsafe.Pointer(lpModuleName)))

    return HINSTANCE(ret)
}

func GetNumberFormat(Locale LCID, dwFlags uint32, lpValue *uint16, lpFormat *NUMBERFMT, lpNumberStr *uint16, cchNumber int32) int32 {
    ret, _, _ := getNumberFormat.Call(
        uintptr(Locale),
        uintptr(dwFlags),
        uintptr(unsafe.Pointer(lpValue)),
        uintptr(unsafe.Pointer(lpFormat)),
        uintptr(unsafe.Pointer(lpNumberStr)),
        uintptr(cchNumber))

    return int32(ret)
}

func GetProfileString(lpAppName, lpKeyName, lpDefault *uint16, lpReturnedString uintptr, nSize uint32) bool {
    ret, _, _ := getProfileString.Call(
        uintptr(unsafe.Pointer(lpAppName)),
        uintptr(unsafe.Pointer(lpKeyName)),
        uintptr(unsafe.Pointer(lpDefault)),
        lpReturnedString,
        uintptr(nSize))
    return ret != 0
}

func GetThreadLocale() LCID {
    ret, _, _ := getThreadLocale.Call()

    return LCID(ret)
}

func GetThreadUILanguage() LANGID {
    ret, _, _ := getThreadUILanguage.Call()

    return LANGID(ret)
}

func GetVersion() int64 {
    ret, _, _ := getVersion.Call()
    return int64(ret)
}

func GlobalAlloc(uFlags uint32, dwBytes uintptr) HGLOBAL {
    ret, _, _ := globalAlloc.Call(
        uintptr(uFlags),
        dwBytes)

    return HGLOBAL(ret)
}

func GlobalFree(hMem HGLOBAL) HGLOBAL {
    ret, _, _ := globalFree.Call(uintptr(hMem))

    return HGLOBAL(ret)
}

func GlobalLock(hMem HGLOBAL) unsafe.Pointer {
    ret, _, _ := globalLock.Call(uintptr(hMem))

    return unsafe.Pointer(ret)
}

func GlobalUnlock(hMem HGLOBAL) bool {
    ret, _, _ := globalUnlock.Call(uintptr(hMem))

    return ret != 0
}

func MoveMemory(destination, source unsafe.Pointer, length uintptr) {
    moveMemory.Call(
        uintptr(unsafe.Pointer(destination)),
        uintptr(source),
        uintptr(length))
}

func MulDiv(nNumber, nNumerator, nDenominator int32) int32 {
    ret, _, _ := mulDiv.Call(
        uintptr(nNumber),
        uintptr(nNumerator),
        uintptr(nDenominator))

    return int32(ret)
}

func SetLastError(dwErrorCode uint32) {
    setLastError.Call(uintptr(dwErrorCode))
}

func SystemTimeToFileTime(lpSystemTime *SYSTEMTIME, lpFileTime *FILETIME) bool {
    ret, _, _ := systemTimeToFileTime.Call(
        uintptr(unsafe.Pointer(lpSystemTime)),
        uintptr(unsafe.Pointer(lpFileTime)))

    return ret != 0
}
