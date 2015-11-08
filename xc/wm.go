package xc

//
/* Window message constants Windows 标准消息
   WM_APP 用于帮助应用程序自定义私有消息,通常形式为:WM_APP + X
   WM_ACTIVATE 一个窗口被激活或失去激活状态
   WM_ACTIVATEAPP 窗口进程激活状态改动,正被激活的窗口属于不同的应用程序
   WM_AFXFIRST 指定首个AFX消息(MFC)
   WM_AFXLAST 指定末个afx消息
   WM_ASKCBFORMATNAME 通过剪贴板观察窗口发送本消息给剪贴板的所有者,以请求一个CF_OWNERDISPLAY格式的剪贴板的名字
   WM_CANCELJOURNAL 当用户取消程序日志激活状态时,发送本消息给那个应用程序。该消息使用空窗口句柄发送
   WM_CANCELMODE 发送本消息来取消某种正在进行的模态(操作)(如鼠示捕获),例如:启动一个模态窗口时,父窗会收到本消息;该消息无参数
   WM_CAPTURECHANGED 当它失去捕获的鼠标时,发送本消息给窗口
   WM_CHANGECBCHAIN 当一个窗口从剪贴板观察链中移去时,发送本消息给剪贴板观察链的首个窗口
   WM_CHAR 窗口字符消息.按下某按键,并已发出WM_KEYDOWN、WM_KEYUP消息,本消息包含被按下的按键的字符码
   WM_CHARTOITEM LBS_WANTKEYBOARDINPUT风格的列表框会发送本消息给其所有者,以便响应WM_CHAR消息
   WM_CHILDACTIVATE 点击窗口标题栏或当窗口被激活、移动、大小改变时,会发送本消息给MDI子窗口
   WM_CLEAR 应用程序发送本消息给编辑框或组合框,以清除当前选择的内容
   WM_CLOSE 窗口关闭消息.
   WM_COMMAND
   WM_COMMNOTIFY
   WM_COMPACTING
   WM_COMPAREITEM
   WM_CONTEXTMENU
   WM_COPY
   WM_COPYDATA
   WM_CREATE
   WM_CTLCOLORBTN
   WM_CTLCOLORDLG
   WM_CTLCOLOREDIT
   WM_CTLCOLORLISTBOX
   WM_CTLCOLORMSGBOX
   WM_CTLCOLORSCROLLBAR
   WM_CTLCOLORSTATIC
   WM_CUT
   WM_DEADCHAR
   WM_DELETEITEM
   WM_DESTROY 窗口销毁消息.
   WM_DESTROYCLIPBOARD
   WM_DEVICECHANGE
   WM_DEVMODECHANGE
   WM_DISPLAYCHANGE
   WM_DRAWCLIPBOARD
   WM_DRAWITEM
   WM_DROPFILES
   WM_ENABLE
   WM_ENDSESSION
   WM_ENTERIDLE
   WM_ENTERMENULOOP
   WM_ENTERSIZEMOVE
   WM_ERASEBKGND
   WM_EXITMENULOOP
   WM_EXITSIZEMOVE 窗口退出移动或调整大小模式循环改，详情参加MSDN.
   WM_FONTCHANGE
   WM_GETDLGCODE
   WM_GETFONT
   WM_GETHOTKEY
   WM_GETICON
   WM_GETMINMAXINFO
   WM_GETTEXT
   WM_GETTEXTLENGTH
   WM_HANDHELDFIRST
   WM_HANDHELDLAST
   WM_HELP
   WM_HOTKEY
   WM_HSCROLL
   WM_HSCROLLCLIPBOARD
   WM_ICONERASEBKGND
   WM_INITDIALOG
   WM_INITMENU
   WM_INITMENUPOPUP
   WM_INPUT
   WM_INPUTLANGCHANGE
   WM_INPUTLANGCHANGEREQUEST
   WM_KEYDOWN 窗口键盘按键消息.
   WM_KEYUP
   WM_KILLFOCUS 窗口失去焦点.
   WM_MDIACTIVATE
   WM_MDICASCADE
   WM_MDICREATE
   WM_MDIDESTROY
   WM_MDIGETACTIVE
   WM_MDIICONARRANGE
   WM_MDIMAXIMIZE
   WM_MDINEXT
   WM_MDIREFRESHMENU
   WM_MDIRESTORE
   WM_MDISETMENU
   WM_MDITILE
   WM_MEASUREITEM
   WM_GETOBJECT
   WM_CHANGEUISTATE
   WM_UPDATEUISTATE
   WM_QUERYUISTATE
   WM_UNINITMENUPOPUP
   WM_MENURBUTTONUP
   WM_MENUCOMMAND
   WM_MENUGETOBJECT
   WM_MENUDRAG
   WM_APPCOMMAND
   WM_MENUCHAR
   WM_MENUSELECT
   WM_MOVE
   WM_MOVING
   WM_NCACTIVATE
   WM_NCCALCSIZE
   WM_NCCREATE
   WM_NCDESTROY 窗口非客户区销毁消息.
   WM_NCHITTEST
   WM_NCLBUTTONDBLCLK
   WM_NCLBUTTONDOWN
   WM_NCLBUTTONUP
   WM_NCMBUTTONDBLCLK
   WM_NCMBUTTONDOWN
   WM_NCMBUTTONUP
   WM_NCXBUTTONDOWN
   WM_NCXBUTTONUP
   WM_NCXBUTTONDBLCLK
   WM_NCMOUSEHOVER
   WM_NCMOUSELEAVE
   WM_NCMOUSEMOVE
   WM_NCPAINT
   WM_NCRBUTTONDBLCLK
   WM_NCRBUTTONDOWN
   WM_NCRBUTTONUP
   WM_NEXTDLGCTL
   WM_NEXTMENU
   WM_NOTIFY
   WM_NOTIFYFORMAT
   WM_NULL
   WM_PAINT 窗口绘制消息
   WM_PAINTCLIPBOARD
   WM_PAINTICON
   WM_PALETTECHANGED
   WM_PALETTEISCHANGING
   WM_PARENTNOTIFY
   WM_PASTE
   WM_PENWINFIRST
   WM_PENWINLAST
   WM_POWER
   WM_POWERBROADCAST
   WM_PRINT
   WM_PRINTCLIENT
   WM_QUERYDRAGICON
   WM_QUERYENDSESSION
   WM_QUERYNEWPALETTE
   WM_QUERYOPEN
   WM_QUEUESYNC
   WM_QUIT
   WM_RENDERALLFORMATS
   WM_RENDERFORMAT
   WM_SETCURSOR 窗口设置鼠标光标.
   WM_SETFOCUS 窗口获得焦点.
   WM_SETFONT
   WM_SETHOTKEY
   WM_SETICON
   WM_SETREDRAW
   WM_SETTEXT
   WM_SETTINGCHANGE
   WM_SHOWWINDOW
   WM_SIZE 窗口大小改变消息.
   WM_SIZECLIPBOARD
   WM_SIZING
   WM_SPOOLERSTATUS
   WM_STYLECHANGED
   WM_STYLECHANGING
   WM_SYSCHAR
   WM_SYSCOLORCHANGE
   WM_SYSCOMMAND
   WM_SYSDEADCHAR
   WM_SYSKEYDOWN
   WM_SYSKEYUP
   WM_TCARD
   WM_THEMECHANGED
   WM_TIMECHANGE
   WM_TIMER 窗口定时器消息.
   WM_UNDO
   WM_USER
   WM_USERCHANGED
   WM_VKEYTOITEM
   WM_VSCROLL
   WM_VSCROLLCLIPBOARD
   WM_WINDOWPOSCHANGED
   WM_WINDOWPOSCHANGING
   WM_WININICHANGE
   WM_KEYFIRST
   WM_KEYLAST
   WM_SYNCPAINT
   WM_MOUSEACTIVATE
   WM_MOUSEMOVE 窗口鼠标移动消息.
   WM_LBUTTONDOWN 窗口鼠标左键按下消息
   WM_LBUTTONUP 窗口鼠标左键弹起消息.
   WM_LBUTTONDBLCLK 窗口鼠标左键双击消息.
   WM_RBUTTONDOWN 窗口鼠标右键按下消息.
   WM_RBUTTONUP 窗口鼠标右键弹起消息.
   WM_RBUTTONDBLCLK 窗口鼠标右键双击消息.
   WM_MBUTTONDOWN
   WM_MBUTTONUP
   WM_MBUTTONDBLCLK
   WM_MOUSEWHEEL 窗口鼠标滚轮滚动消息.
   WM_MOUSEFIRST
   WM_XBUTTONDOWN
   WM_XBUTTONUP
   WM_XBUTTONDBLCLK
   WM_MOUSELAST
   WM_MOUSEHOVER 窗口鼠标进入消息
   WM_MOUSELEAVE 窗口鼠标离开消息.
   WM_CLIPBOARDUPDATE
*/
const (
    WM_APP                    = 32768
    WM_ACTIVATE               = 6
    WM_ACTIVATEAPP            = 28
    WM_AFXFIRST               = 864
    WM_AFXLAST                = 895
    WM_ASKCBFORMATNAME        = 780
    WM_CANCELJOURNAL          = 75
    WM_CANCELMODE             = 31
    WM_CAPTURECHANGED         = 533
    WM_CHANGECBCHAIN          = 781
    WM_CHAR                   = 258
    WM_CHARTOITEM             = 47
    WM_CHILDACTIVATE          = 34
    WM_CLEAR                  = 771
    WM_CLOSE                  = 16
    WM_COMMAND                = 273
    WM_COMMNOTIFY             = 68 /* OBSOLETE */
    WM_COMPACTING             = 65
    WM_COMPAREITEM            = 57
    WM_CONTEXTMENU            = 123
    WM_COPY                   = 769
    WM_COPYDATA               = 74
    WM_CREATE                 = 1
    WM_CTLCOLORBTN            = 309
    WM_CTLCOLORDLG            = 310
    WM_CTLCOLOREDIT           = 307
    WM_CTLCOLORLISTBOX        = 308
    WM_CTLCOLORMSGBOX         = 306
    WM_CTLCOLORSCROLLBAR      = 311
    WM_CTLCOLORSTATIC         = 312
    WM_CUT                    = 768
    WM_DEADCHAR               = 259
    WM_DELETEITEM             = 45
    WM_DESTROY                = 2
    WM_DESTROYCLIPBOARD       = 775
    WM_DEVICECHANGE           = 537
    WM_DEVMODECHANGE          = 27
    WM_DISPLAYCHANGE          = 126
    WM_DRAWCLIPBOARD          = 776
    WM_DRAWITEM               = 43
    WM_DROPFILES              = 563
    WM_ENABLE                 = 10
    WM_ENDSESSION             = 22
    WM_ENTERIDLE              = 289
    WM_ENTERMENULOOP          = 529
    WM_ENTERSIZEMOVE          = 561
    WM_ERASEBKGND             = 20
    WM_EXITMENULOOP           = 530
    WM_EXITSIZEMOVE           = 562
    WM_FONTCHANGE             = 29
    WM_GETDLGCODE             = 135
    WM_GETFONT                = 49
    WM_GETHOTKEY              = 51
    WM_GETICON                = 127
    WM_GETMINMAXINFO          = 36
    WM_GETTEXT                = 13
    WM_GETTEXTLENGTH          = 14
    WM_HANDHELDFIRST          = 856
    WM_HANDHELDLAST           = 863
    WM_HELP                   = 83
    WM_HOTKEY                 = 786
    WM_HSCROLL                = 276
    WM_HSCROLLCLIPBOARD       = 782
    WM_ICONERASEBKGND         = 39
    WM_INITDIALOG             = 272
    WM_INITMENU               = 278
    WM_INITMENUPOPUP          = 279
    WM_INPUT                  = 0X00FF
    WM_INPUTLANGCHANGE        = 81
    WM_INPUTLANGCHANGEREQUEST = 80
    WM_KEYDOWN                = 256
    WM_KEYUP                  = 257
    WM_KILLFOCUS              = 8
    WM_MDIACTIVATE            = 546
    WM_MDICASCADE             = 551
    WM_MDICREATE              = 544
    WM_MDIDESTROY             = 545
    WM_MDIGETACTIVE           = 553
    WM_MDIICONARRANGE         = 552
    WM_MDIMAXIMIZE            = 549
    WM_MDINEXT                = 548
    WM_MDIREFRESHMENU         = 564
    WM_MDIRESTORE             = 547
    WM_MDISETMENU             = 560
    WM_MDITILE                = 550
    WM_MEASUREITEM            = 44
    WM_GETOBJECT              = 0X003D
    WM_CHANGEUISTATE          = 0X0127
    WM_UPDATEUISTATE          = 0X0128
    WM_QUERYUISTATE           = 0X0129
    WM_UNINITMENUPOPUP        = 0X0125
    WM_MENURBUTTONUP          = 290
    WM_MENUCOMMAND            = 0X0126
    WM_MENUGETOBJECT          = 0X0124
    WM_MENUDRAG               = 0X0123
    WM_APPCOMMAND             = 0X0319
    WM_MENUCHAR               = 288
    WM_MENUSELECT             = 287
    WM_MOVE                   = 3
    WM_MOVING                 = 534
    WM_NCACTIVATE             = 134
    WM_NCCALCSIZE             = 131
    WM_NCCREATE               = 129
    WM_NCDESTROY              = 130
    WM_NCHITTEST              = 132
    WM_NCLBUTTONDBLCLK        = 163
    WM_NCLBUTTONDOWN          = 161
    WM_NCLBUTTONUP            = 162
    WM_NCMBUTTONDBLCLK        = 169
    WM_NCMBUTTONDOWN          = 167
    WM_NCMBUTTONUP            = 168
    WM_NCXBUTTONDOWN          = 171
    WM_NCXBUTTONUP            = 172
    WM_NCXBUTTONDBLCLK        = 173
    WM_NCMOUSEHOVER           = 0X02A0
    WM_NCMOUSELEAVE           = 0X02A2
    WM_NCMOUSEMOVE            = 160
    WM_NCPAINT                = 133
    WM_NCRBUTTONDBLCLK        = 166
    WM_NCRBUTTONDOWN          = 164
    WM_NCRBUTTONUP            = 165
    WM_NEXTDLGCTL             = 40
    WM_NEXTMENU               = 531
    WM_NOTIFY                 = 78
    WM_NOTIFYFORMAT           = 85
    WM_NULL                   = 0
    WM_PAINT                  = 15
    WM_PAINTCLIPBOARD         = 777
    WM_PAINTICON              = 38
    WM_PALETTECHANGED         = 785
    WM_PALETTEISCHANGING      = 784
    WM_PARENTNOTIFY           = 528
    WM_PASTE                  = 770
    WM_PENWINFIRST            = 896
    WM_PENWINLAST             = 911
    WM_POWER                  = 72
    WM_POWERBROADCAST         = 536
    WM_PRINT                  = 791
    WM_PRINTCLIENT            = 792
    WM_QUERYDRAGICON          = 55
    WM_QUERYENDSESSION        = 17
    WM_QUERYNEWPALETTE        = 783
    WM_QUERYOPEN              = 19
    WM_QUEUESYNC              = 35
    WM_QUIT                   = 18
    WM_RENDERALLFORMATS       = 774
    WM_RENDERFORMAT           = 773
    WM_SETCURSOR              = 32
    WM_SETFOCUS               = 7
    WM_SETFONT                = 48
    WM_SETHOTKEY              = 50
    WM_SETICON                = 128
    WM_SETREDRAW              = 11
    WM_SETTEXT                = 12
    WM_SETTINGCHANGE          = 26
    WM_SHOWWINDOW             = 24
    WM_SIZE                   = 5
    WM_SIZECLIPBOARD          = 779
    WM_SIZING                 = 532
    WM_SPOOLERSTATUS          = 42
    WM_STYLECHANGED           = 125
    WM_STYLECHANGING          = 124
    WM_SYSCHAR                = 262
    WM_SYSCOLORCHANGE         = 21
    WM_SYSCOMMAND             = 274
    WM_SYSDEADCHAR            = 263
    WM_SYSKEYDOWN             = 260
    WM_SYSKEYUP               = 261
    WM_TCARD                  = 82
    WM_THEMECHANGED           = 794
    WM_TIMECHANGE             = 30
    WM_TIMER                  = 275
    WM_UNDO                   = 772
    WM_USER                   = 1024
    WM_USERCHANGED            = 84
    WM_VKEYTOITEM             = 46
    WM_VSCROLL                = 277
    WM_VSCROLLCLIPBOARD       = 778
    WM_WINDOWPOSCHANGED       = 71
    WM_WINDOWPOSCHANGING      = 70
    WM_WININICHANGE           = 26
    WM_KEYFIRST               = 256
    WM_KEYLAST                = 264
    WM_SYNCPAINT              = 136
    WM_MOUSEACTIVATE          = 33
    WM_MOUSEMOVE              = 512
    WM_LBUTTONDOWN            = 513
    WM_LBUTTONUP              = 514
    WM_LBUTTONDBLCLK          = 515
    WM_RBUTTONDOWN            = 516
    WM_RBUTTONUP              = 517
    WM_RBUTTONDBLCLK          = 518
    WM_MBUTTONDOWN            = 519
    WM_MBUTTONUP              = 520
    WM_MBUTTONDBLCLK          = 521
    WM_MOUSEWHEEL             = 522
    WM_MOUSEFIRST             = 512
    WM_XBUTTONDOWN            = 523
    WM_XBUTTONUP              = 524
    WM_XBUTTONDBLCLK          = 525
    WM_MOUSELAST              = 525
    WM_MOUSEHOVER             = 0X2A1
    WM_MOUSELEAVE             = 0X2A3
    WM_CLIPBOARDUPDATE        = 0x031D
)

/* 窗口事件
   XWM_MENU_POPUP   0x7000+5   菜单弹出
   XWM_MENU_POPUP_WND   0x7000+6   菜单弹出窗口
   XWM_MENU_SELECT   0x7000+7   菜单选择
   XWM_MENU_EXIT   0x7000+8   菜单退出
   XWM_MENU_DRAW_BACKGROUND   0x7000+9   绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
   XWM_MENU_DRAWITEM   0x7000+10   绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
*/
const (
    XWM_MENU_POPUP           = 0x7000 + 5
    XWM_MENU_POPUP_WND       = 0x7000 + 6
    XWM_MENU_SELECT          = 0x7000 + 7
    XWM_MENU_EXIT            = 0x7000 + 8
    XWM_MENU_DRAW_BACKGROUND = 0x7000 + 9
    XWM_MENU_DRAWITEM        = 0x7000 + 10
)
