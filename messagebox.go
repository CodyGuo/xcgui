package xcgui

import (
    "github.com/codyguo/xcgui/xc"
)

type MsgBoxStyle uint

const (
    MsgBoxOK                MsgBoxStyle = xc.MB_OK
    MsgBoxOKCancel          MsgBoxStyle = xc.MB_OKCANCEL
    MsgBoxAbortRetryIgnore  MsgBoxStyle = xc.MB_ABORTRETRYIGNORE
    MsgBoxYesNoCancel       MsgBoxStyle = xc.MB_YESNOCANCEL
    MsgBoxYesNo             MsgBoxStyle = xc.MB_YESNO
    MsgBoxRetryCancel       MsgBoxStyle = xc.MB_RETRYCANCEL
    MsgBoxCancelTryContinue MsgBoxStyle = xc.MB_CANCELTRYCONTINUE
    MsgBoxIconHand          MsgBoxStyle = xc.MB_ICONHAND
    MsgBoxIconQuestion      MsgBoxStyle = xc.MB_ICONQUESTION
    MsgBoxIconExclamation   MsgBoxStyle = xc.MB_ICONEXCLAMATION
    MsgBoxIconAsterisk      MsgBoxStyle = xc.MB_ICONASTERISK
    MsgBoxUserIcon          MsgBoxStyle = xc.MB_USERICON
    MsgBoxIconWarning       MsgBoxStyle = xc.MB_ICONWARNING
    MsgBoxIconError         MsgBoxStyle = xc.MB_ICONERROR
    MsgBoxIconInformation   MsgBoxStyle = xc.MB_ICONINFORMATION
    MsgBoxIconStop          MsgBoxStyle = xc.MB_ICONSTOP
    MsgBoxDefButton1        MsgBoxStyle = xc.MB_DEFBUTTON1
    MsgBoxDefButton2        MsgBoxStyle = xc.MB_DEFBUTTON2
    MsgBoxDefButton3        MsgBoxStyle = xc.MB_DEFBUTTON3
    MsgBoxDefButton4        MsgBoxStyle = xc.MB_DEFBUTTON4
)

func MsgBox(hParent xc.HWINDOW, title, message string, style MsgBoxStyle) int {
    return int(xc.MessageBox(
        xc.HWND(hParent),
        message,
        title,
        uint32(style)))
}
