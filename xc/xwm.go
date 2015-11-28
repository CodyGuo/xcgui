package xc

//
/*
Window message constants Windows 标准消息
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
	WM_COMMAND 本消息用于实现MFC的上下文敏感帮助,按下<F1键>后消息被映射到CWinApp::OnHelp。该函数会向最外层框架窗口发送本消息,本消息响应过程是自顶向下的,对应的消息响应函数为:ON_WM_HELPINFO"
	WM_COMMNOTIFY Win3.1中,当串口事件产生时,通讯设备驱动程序发送消息本消息给系统,指示输入输出队列的状态
	WM_COMPACTING 显示内存已经很少了
	WM_COMPAREITEM 可发送本消息来确定组合框(CBS_SORT)或列表框(LBS_SORT)中新增项的相对位置
	WM_CONTEXTMENU 当用户在某窗口中点击右键就发送本消息给该窗口,设置右键菜单
	WM_COPY 应用程序发送本消息给一个编辑框或组合框,以便用CF_TEXT格式复制当前选择的文本到剪贴板
	WM_COPYDATA 当一个应用程序传递数据给另一个应用程序时发送本消息
	WM_CREATE 新建一个窗口
	WM_CTLCOLORBTN 设置按钮的背景色
	WM_CTLCOLORDLG 设置对话框的背景色,通常是在WinnApp中使用SetDialogBkColor函数实现
	WM_CTLCOLOREDIT 当一个编辑框控件将要被绘制时,发送本消息给其父窗;通过响应本消息,所有者窗口可通过使用给定的相关显示设备的句柄来设置编辑框的文本和背景色
	WM_CTLCOLORLISTBOX 当一个列表框控件将要被绘制前,发送本消息给其父窗;通过响应本消息,所有者窗口可通过使用给定的相关显示设备的句柄来设置列表框的文本和背景色
	WM_CTLCOLORMSGBOX 系统绘制消息框前发送本消息给消息框的所有者窗口,通过响应本消息,所有者窗口可通过使用给定的相关显示设备的句柄来设置消息框的文本和背景色
	WM_CTLCOLORSCROLLBAR 设置滚动条的背景色
	WM_CTLCOLORSTATIC 设置一个静态控件的背景色
	WM_CUT 应用程序发送本消息给一个编辑框或组合框来删除当前选择的文本
	WM_DEADCHAR \"死字符\"消息,当使用TranslateMessage函数翻译WM_KEYUP消息时,发送本消息给拥有键盘焦点的窗口,注:德语键盘上,有些按键只是给字符添加音标的,并不产生字符,故称\"死字符\"
	WM_DELETEITEM 当列表框或组合框被销毁或通过LB_DELETESTRING、LB_RESETCONTENT、CB_DELETESTRING或CB_RESETCONTENT消息删除某些项时,会发送本消息给这些控件的所有者
	WM_DESTROY 窗口销毁消息.
	WM_DESTROYCLIPBOARD 当调用EmptyClipboard函数时,发送本消息给剪贴板的所有者
	WM_DEVICECHANGE 当设备的硬件配置改变时,发送本消息给应用程序或设备驱动程序
	WM_DEVMODECHANGE 改变设备模式设置(\"win.ini\")时,处理本消息的应用程序可重新初始化它们的设备模式设置
	WM_DISPLAYCHANGE 当显示器的分辨率改变后,发送本消息给所有窗口
	WM_DRAWCLIPBOARD 当剪贴板的内容变化时,发送本消息给剪贴板观察链的首个窗口;它允许用剪贴板观察窗口来显示剪贴板的新内容
	WM_DRAWITEM 按钮、组合框、列表框、菜单的外观改变时会发送本消息给这些控件的所有者
	WM_DROPFILES 鼠标拖放时,放下事件产生时发送本消息,比如:文件拖放功能
	WM_ENABLE 使一个窗口处于可用状态
	WM_ENDSESSION 关机或注销时系统会发出WM_QUERYENDSESSION消息,然后将本消息发送给应用程序,通知程序会话结束
	WM_ENTERIDLE 当一个模态对话框或菜单进入空闲状态时,发送本消息给它的所有者,一个模态对话框或菜单进入空闲状态就是在处理完一条或几条先前的消息后,没有消息在消息列队中等待
	WM_ENTERMENULOOP 发送本消息通知应用程序的主窗口已进入菜单循环模式
	WM_ENTERSIZEMOVE 当某窗口进入移动或调整大小的模式循环时,本消息发送到该窗口
	WM_ERASEBKGND 当一个窗口的背景必须被擦除时本消息会被触发(如:窗口大小改变时)
	WM_EXITMENULOOP 发送本消息通知应用程序的主窗口已退出菜单循环模式
	WM_EXITSIZEMOVE 窗口退出移动或调整大小模式循环改，详情参加MSDN.
	WM_FONTCHANGE 当系统的字体资源库变化时发送本消息给所有顶级窗口
	WM_GETDLGCODE 发送本消息给某个与对话框程序关联的控件,系统控制方位键和TAB键使输入进入该控件,通过响应本消息应用程序可把它当成一个特殊的输入控件并能处理它
	WM_GETFONT 得到当前控件绘制其文本所用的字体
	WM_GETHOTKEY 确定某热键与某窗口是否相关联
	WM_GETICON 本消息发送给某个窗口,用于返回与某窗口有关联的大图标或小图标的句柄
	WM_GETMINMAXINFO 当窗口将要改变大小或位置时,由系统发送本消息给窗口,用户拖动一个可重置大小的窗口时便会发出本消息
	WM_GETTEXT 复制窗口的文本到缓冲区
	WM_GETTEXTLENGTH 得到窗口的文本长度(不含结束符)
	WM_HANDHELDFIRST 消息含义未知,搜索了整个Visual Studio 6的目录也只有其定义,却未见其使用的相关代码
	WM_HANDHELDLAST 消息含义未知,搜索了整个Visual Studio 6的目录也只有其定义,却未见其使用的相关代码
	WM_HELP 按下<F1>后,若某菜单是激活的,就发送本消息给此窗口关联的菜单;否则就发送给有焦点的窗口;若当前都没有焦点,就把本消息发送给当前激活的窗口
	WM_HOTKEY 当用户按下由RegisterHotKey函数注册的热键时,发送本消息
	WM_HSCROLL 当窗口的标准水平滚动条产生一个滚动事件时,发送本消息给该窗口
	WM_HSCROLLCLIPBOARD 本消息通过一个剪贴板观察窗口发送给剪贴板的所有者,它发生在当剪贴板包含CFOWNERDISPALY格式的数据,并且有个事件在剪贴板观察窗的水平滚动条上,所有者应滚动剪贴板图像并更新滚动条的值
	WM_ICONERASEBKGND 本消息发送给某个最小化的窗口,仅当它在画图标前它的背景必须被重画
	WM_INITDIALOG 在某对话框程序被显示前发送本消息给该对话框程序,通常用本消息对控件进行一些初始化工作和执行其它任务
	WM_INITMENU 当一个菜单将被激活时发送本消息,它发生在用户点击了某菜单项或按下某菜单键。它允许程序在显示前更改菜单
	WM_INITMENUPOPUP 当一个下拉菜单或子菜单将被激活时发送本消息,它允许程序在它显示前更改菜单,却不更改全部菜单
	WM_INPUT
	WM_INPUTLANGCHANGE 切换输入法后,系统会发送本消息给受影响的顶层窗口
	WM_INPUTLANGCHANGEREQUEST 当用户通过过单击任务栏上的语言指示符或某快捷键组合选择改变输入法时系统会向焦点窗口发送本消息
	WM_KEYDOWN 窗口键盘按键消息.
	WM_KEYUP 当一个非系统按键被释放弹起时(<ALT>键没有被按下),会发送本消息给拥有键盘焦点的窗口
	WM_KILLFOCUS 窗口失去焦点.
	WM_MDIACTIVATE 发送本消息给多文档应用程序的客户窗口通知客户窗口激活另一个MDI子窗口,当客户窗口收到本消息后,它发出WM_MDIACTIVE消息给MDI子窗口(未激活)来激活它
	WM_MDICASCADE 发送本消息给MDI客户窗口,以层叠方式重新排列所有MDI子窗口
	WM_MDICREATE 发送本消息给多文档应用程序的客户窗口来创建一个MDI子窗口
	WM_MDIDESTROY 发送本消息给多文档应用程序的客户窗口来关闭一个MDI子窗口
	WM_MDIGETACTIVE 发送本消息给MDI客户窗口以找到激活的子窗口句柄
	WM_MDIICONARRANGE 发送本消息给MDI客户窗口重新排列所有最小化的MDI子窗口
	WM_MDIMAXIMIZE 发送本消息给MDI客户窗口来最大化一个MDI子窗口
	WM_MDINEXT 发送本消息给MDI客户窗口,激活下一个或前一个窗口
	WM_MDIREFRESHMENU 发送本消息给多文档应用程序的客户窗口,根据当前MDI子窗口更新MDI框架窗口的菜单
	WM_MDIRESTORE 发送本消息给MDI客户窗口,让子窗口从最大最小化恢复到原来的大小
	WM_MDISETMENU 发送本消息给MDI客户窗口,用MDI菜单代替子窗口的菜单
	WM_MDITILE 发送本消息给MDI客户窗口,以平铺方式重新排列所有MDI子窗口
	WM_MEASUREITEM 按钮、组合框、列表框、列表控件、菜单项被创建时会发送本消息给这些控件的所有者
	WM_GETOBJECT \"oleacc.dll\"(COM组件)(Microsoft Active Accessibility:方便残疾人使用电脑的一种技术)发送本消息激活服务程序以便获取它所包含的关联对象的信息
	WM_CHANGEUISTATE
	WM_UPDATEUISTATE
	WM_QUERYUISTATE
	WM_UNINITMENUPOPUP 当一条下拉菜单或子菜单被销毁时,发送本消息
	WM_MENURBUTTONUP 本消息允许程序为菜单项提供一个感知上下文的菜单(即快捷菜单),要为菜单项显示一下上下文菜单,请使用TPM_RECURSE标识调用TrackPopupMenuEx函数
	WM_MENUCOMMAND 当用户在一个菜单上作出选择时,会发送本消息,菜单要具有MNS_NOTIFYBYPOS风格(在MENUINFO结构体中设置)
	WM_MENUGETOBJECT 当鼠标光标进入或离开菜单项时,本消息发送给支持拖放的菜单的拥有者,相关结构体:MENUGETOBJECTINFO,注:菜单要具有MNS_DRAGDROP风格
	WM_MENUDRAG 当用户拖动菜单项时,发送本消息给拖放菜单的拥有者,可让菜单支持拖拽,可使用OLE拖放传输协议启动拖放操作,注:菜单要具有MNS_DRAGDROP风格
	WM_APPCOMMAND
	WM_MENUCHAR 当菜单已被激活且用户按下了某菜单字符键(菜单字符键用括号括着、带下划线,不同于快捷键),发送本消息给菜单的所有者
	WM_MENUSELECT 当用户选择一条菜单项时,发送本消息给菜单的所有者(一般是窗口)
	WM_MOVE 移动一个窗口
	WM_MOVING 当用户在移动窗口时发送本消息,通过本消息应用程序以监视窗口大小和位置,也可修改它们
	WM_NCACTIVATE 本消息发送给某窗口,在窗口的非客户区被激活时重绘窗口
	WM_NCCALCSIZE 当某窗口的客户区的大小和位置须被计算时发送本消息
	WM_NCCREATE 当某窗口首次被创建时,本消息在WM_CREATE消息发送前发送
	WM_NCDESTROY 窗口非客户区销毁消息.
	WM_NCHITTEST 当用户在在非客户区移动鼠标、按住或释放鼠标时发送本消息(击中测试);若鼠标没有被捕获,则本消息在窗口得到光标之后发出,否则消息发送到捕获到鼠标的窗口
	WM_NCLBUTTONDBLCLK 当用户双击鼠标左键的同时光标在某窗口的非客户区内时,会发送本消息
	WM_NCLBUTTONDOWN 当光标在某窗口的非客户区内的同时按下鼠标左键,会发送本消息
	WM_NCLBUTTONUP 当用户释放鼠标左键的同时光标在某窗口的非客户区内时,会发送本消息
	WM_NCMBUTTONDBLCLK 当用户双击鼠标中键的同时光标在某窗口的非客户区内时,会发送本消息
	WM_NCMBUTTONDOWN 当用户按下鼠标中键的同时光标在某窗口的非客户区内时,会发送本消息
	WM_NCMBUTTONUP 当用户释放鼠标中键的同时光标在某窗口的非客户区内时,会发送本消息
	WM_NCXBUTTONDOWN
	WM_NCXBUTTONUP
	WM_NCXBUTTONDBLCLK
	WM_NCMOUSEHOVER
	WM_NCMOUSELEAVE
	WM_NCMOUSEMOVE 当光标在某窗口的非客户区内移动时,发送本消息给该窗口
	WM_NCPAINT 当窗口框架(非客户区)必须被被重绘时,应用程序发送本消息给该窗口
	WM_NCRBUTTONDBLCLK 当用户双击鼠标右键的同时光标在某窗口的非客户区内时,会发送本消息
	WM_NCRBUTTONDOWN 当用户按下鼠标右键的同时光标在某窗口的非客户区内时,会发送本消息
	WM_NCRBUTTONUP 当用户释放鼠标右键的同时光标在某窗口的非客户区内时,会发送本消息
	WM_NEXTDLGCTL 发送本消息给一个对话框程序窗口过程,以便在各控件间设置键盘焦点位置
	WM_NEXTMENU 当使用左箭头光标键或右箭头光标键在菜单条与系统菜单之间切换时,会发送本消息给应用程序,相关结构体:MDINEXTMENU
	WM_NOTIFY 当某控件的某事件已发生或该控件需得到一些信息时,发送本消息给其父窗
	WM_NOTIFYFORMAT 公用控件、自定义控件和其父窗通过本消息判断控件在WM_NOTIFY通知消息中是使用ANSI还是UNICODE,使用本消息能使某个控件与它的父控件间进行相互通信
	WM_NULL 空消息,可检测程序是否有响应等
	WM_PAINT 窗口绘制消息
	WM_PAINTCLIPBOARD 当剪贴板包含CF_OWNERDIPLAY格式的数据,并且剪贴板观察窗口的客户区需要重画时,触发发送本消息
	WM_PAINTICON 当一个最小化的窗口图标将被重绘时发送本消息
	WM_PALETTECHANGED 本消息在一个拥有焦点的窗口实现它的逻辑调色板后,发送本消息给所有顶级并重叠的窗口,以此来改变系统调色板
	WM_PALETTEISCHANGING 当一个应用程序正要实现它的逻辑调色板时,发本消息通知所有的应用程序
	WM_PARENTNOTIFY 当MDI子窗口被创建或被销毁,或用户按了一下鼠标键而光标在子窗口上时,发送本消息给其父窗
	WM_PASTE 应用程序发送本消息给编辑框或组合框,以便从剪贴板中得到数据
	WM_PENWINFIRST 指定首个Pen Window消息,参见:PENWIN.H/WINUSER.H
	WM_PENWINLAST 指定末个Pen Window消息,参见:PENWIN.H/WINUSER.H
	WM_POWER 当系统将要进入暂停状态时发送本消息(适用于16位的windows)
	WM_POWERBROADCAST 本消息发送给应用程序来通知它有关电源管理事件,比如待机休眠时会发送本消息
	WM_PRINT 发送本消息给一个窗口请求在指定的设备上下文中绘制自身,可用于窗口截图,但对子控件截图时得到的是与子控件等大的黑块
	WM_PRINTCLIENT 送本消息给一个窗口请求在指定的设备上下文中绘制其客户区(最通常是在一个打印机设备上下文中)
	WM_QUERYDRAGICON 本消息发送给最小化的窗口(iconic),当该窗口将被拖放而其窗口类中没有定义图标,应用程序能返回一个图标或光标的句柄。当用户拖放图标时系统会显示这个图标或光标
	WM_QUERYENDSESSION 关机或注销时系统会按优先级给各进程发送WM_QUERYENDSESSION,告诉应用程序要关机或注销了
	WM_QUERYNEWPALETTE 本消息发送给将要收到焦点的窗口,本消息能使窗口在收到焦点时同时有机会实现逻辑调色板
	WM_QUERYOPEN 最小化的窗口即将被恢复以前的大小位置
	WM_QUEUESYNC 本消息由基于计算机的训练程序发送,通过WH_JOURNALPALYBACK的HOOK程序分离出用户输入消息
	WM_QUIT 关闭消息循环结束程序的运行
	WM_RENDERALLFORMATS 应用程序退出时在程序退出时,系统会给当前程序发送该消息,要求提供所有格式的剪帖板数据,避免造成数据丢失
	WM_RENDERFORMAT 应用程序需要系统剪切板数据时,触发发送本消息
	WM_SETCURSOR 窗口设置鼠标光标.
	WM_SETFOCUS 窗口获得焦点.
	WM_SETFONT 指定控件所用字体
	WM_SETHOTKEY 为某窗口关联一个热键
	WM_SETICON 应用程序发送本消息让一个新的大图标或小图标与某窗口相关联
	WM_SETREDRAW 设置窗口是否能重绘
	WM_SETTEXT 设置一个窗口的文本
	WM_SETTINGCHANGE 当SystemParametersInfo函数改变一个系统范围内的设置时，系统向所有顶层窗口发送WM_SETTINGCHANGE消息。系统仅在 SystemParametersInfo的调用者指定了SPIF_SENDCHANGE标志时 才发送此消息
	WM_SHOWWINDOW 发送本消息给一个窗口,以便隐藏或显示该窗口
	WM_SIZE 窗口大小改变消息.
	WM_SIZECLIPBOARD 当剪贴板包含CF_OWNERDIPLAY格式的数据,并且剪贴板观察窗口的客户区域的大小已改变时,本消息通过剪贴板观察窗口发送给剪贴板的所有者
	WM_SIZING 当用户正在调整窗口大小时,发送本消息给窗口;通过本消息应用程序可监视窗口大小和位置,也可修改它们
	WM_SPOOLERSTATUS 每当打印管理列队增加或减少一条作业时就会发出本消息
	WM_STYLECHANGED 当调用SetWindowLong函数改变一个或多个窗口的风格后,发送本消息给那个窗口
	WM_STYLECHANGING 当调用SetWindowLong函数将要改变一个或多个窗口的风格时,发送本消息给那个窗口
	WM_SYSCHAR 当WM_SYSKEYDOWN消息被TranslateMessage函数翻译后,发送本消息给拥有焦点的窗口,注:<ALT>键被按下
	WM_SYSCOLORCHANGE 当系统颜色改变时,发送本消息给所有顶级窗口
	WM_SYSCOMMAND 当用户选择一条系统菜单命令、用户最大化或最小化或还原或关闭时，窗口会收到本消息
	WM_SYSDEADCHAR \"死字符\"消息,当使用TranslateMessage函数翻译WM_SYSKEYDOWN消息时,发送本消息给拥有键盘焦点的窗口,注:德语键盘上,有些按键只是给字符添加音标的,并不产生字符,故称\"死字符\"
	WM_SYSKEYDOWN 当用户按住<ALT>键的同时又按下其它键时,发送本消息给拥有焦点的窗口
	WM_SYSKEYUP 当用户释放一个按键的同时<ALT>键还按着时,发送本消息给拥有焦点的窗口
	WM_TCARD 程序已初始化windows帮助例程时会发送本消息给应用程序
	WM_THEMECHANGED
	WM_TIMECHANGE 当系统的时间变化时发送本消息给所有顶级窗口
	WM_TIMER 窗口定时器消息.
	WM_UNDO 应用程序发送本消息给编辑框或组合框,以撤消最后一次操作
	WM_USER 用于帮助应用程序自定义私有消息,通常形式为:WM_USER + X
	WM_USERCHANGED 当用户已登入或退出后发送本消息给所有窗口;当用户登入或退出时系统更新用户的具体设置信息,在用户更新设置时系统马上发送本消息
	WM_VKEYTOITEM LBS_WANTKEYBOARDINPUT风格的列表框会发出本消息给其所有者,以便响应WM_KEYDOWN消息
	WM_VSCROLL 当窗口的标准垂直滚动条产生一个滚动事件时,发送本消息给该窗口
	WM_VSCROLLCLIPBOARD 当剪贴板查看器的垂直滚动条被单击时,触发发送本消息
	WM_WINDOWPOSCHANGED 本消息会发送给那些大小和位置(Z_Order)已被改变的窗口,以调用SetWindowPos函数或其它窗口管理函数
	WM_WINDOWPOSCHANGING 本消息会发送给那些大小和位置(Z_Order)将被改变的窗口,以调用SetWindowPos函数或其它窗口管理函数
	WM_WININICHANGE 读写\"win.ini\"时会发送本消息给所有顶层窗口,通知其它进程该文件已被更改
	WM_KEYFIRST 用于WinCE系统,本消息在使用GetMessage和PeekMessage函数时,用于过滤键盘消息
	WM_KEYLAST 用于WinCE系统,本消息在使用GetMessage和PeekMessage函数时,用于过滤键盘消息
	WM_SYNCPAINT 当避免联系独立的GUI线程时,本消息用于同步刷新,本消息由系统确定是否发送
	WM_MOUSEACTIVATE 当鼠标光标在某个未激活窗口内,而用户正按着鼠标的某个键时,会发送本消息给当前窗口
	WM_MOUSEMOVE 窗口鼠标移动消息.
	WM_LBUTTONDOWN 窗口鼠标左键按下消息
	WM_LBUTTONUP 窗口鼠标左键弹起消息.
	WM_LBUTTONDBLCLK 窗口鼠标左键双击消息.
	WM_RBUTTONDOWN 窗口鼠标右键按下消息.
	WM_RBUTTONUP 窗口鼠标右键弹起消息.
	WM_RBUTTONDBLCLK 窗口鼠标右键双击消息.
	WM_MBUTTONDOWN 按下鼠标中键
	WM_MBUTTONUP 释放鼠标中键
	WM_MBUTTONDBLCLK 双击鼠标中键
	WM_MOUSEWHEEL 窗口鼠标滚轮滚动消息.
	WM_MOUSEFIRST 鼠标移动时发生(与WM_MOUSEMOVE等值),常用于判断鼠标消息的范围,比如:if(message >= WM_MOUSEFIRST)&&(message <= WM_MOUSELAST)
	WM_XBUTTONDOWN
	WM_XBUTTONUP 当光标在窗口客户区时，用户释放鼠标左键时发出的消息。如果鼠标没有捕获，这个消息被送到光标下的窗口。否则，该消息发布到捕获鼠标的窗口。
	WM_XBUTTONDBLCLK 表示左键双击事件.
	WM_MOUSELAST WM_MBUTTONDBLCLK的别名,通常用于判断鼠标消息的范围,对应的还有WM_MOUSEFIRST,例如:if(message > =  WM_MOUSEFIRST)&&(message <= WM_MOUSELAST)
	WM_MOUSEHOVER 窗口鼠标进入消息
	WM_MOUSELEAVE 窗口鼠标离开消息.
	WM_CLIPBOARDUPDATE 处理WM_CLIPBOARDUPDATE消息,当剪贴板的内容发生变化后,消息处理程序会收到这条消息通知.
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

/*
窗口事件
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

/*
重绘元素
	XWM_REDRAW 窗口重绘延时 ----不公开-----内部自定义消息
	XWM_REDRAW_ELE 重绘元素 wParam:元素句柄, lParam:RECT*基于窗口坐标
*/
const (
	XWM_REDRAW     = WM_APP + 1007
	XWM_REDRAW_ELE = 0x7000 + 3
)

/*
窗口消息-包含系统非客户区消息
	XWM_EVENT_ALL 事件投递 -------不公开-------不需要注册
	XWM_WINDPROC 注册窗口处理过程
	XWM_DRAW_T 窗口绘制,内部使用, wParam:0, lParam:0
*/
const (
	XWM_EVENT_ALL = WM_APP + 1000
	XWM_WINDPROC  = 0x7000
	XWM_DRAW_T    = 0x7000 + 1
)
