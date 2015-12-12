# xcgui
炫彩界面库，用的免费版dll。<br>
xc目录下为原API可以直接使用，进度中为原API的记录。<br>
有钱的朋友可以去官方网站购买正式版，支持正版将有更多功能。<br>
官方网站：[www.xcgui.com](http://www.xcgui.com "xcgui 官方网站")<br>

# DOC
帮助文档在doc目录下，或者访问：[godoc.org](https://godoc.org/github.com/CodyGuo/xcgui/xc "xcgui api")<br>
友情提醒：64位编译环境下需要设置 set GOARCH=386 ，可写入bat编译。<br>
支持与walk混写，参考[托盘](https://github.com/CodyGuo/xcgui/tree/master/examples/notifyicon) [浏览图片](https://github.com/CodyGuo/xcgui/tree/master/examples/imageviewer)例子。

# GOAL

  Hopefully support following widgets and methods enough to run general application.

(炫彩界面库模块)

    + UI窗口模块列表
        - FrameWindow                   : 100% (  7/  7) -框架窗口
        - Menu                          : 100% ( 25/ 25) -弹出菜单
        - modalWinow                    : 100% (  5/  5) -模态窗口
        - window                        : 100% ( 63/ 63) -基础窗口

    + UI元素模块列表
        - Button                        : 100% ( 32/ 32) -按钮
        - ComboBox                      : 100% ( 19/ 19) -下拉组合框
        - Element                       : 100% ( 91/ 91) -基础元素
        - ListBox                       : 100% ( 31/ 31) -列表框
        - list                          : 100% ( 36/ 36) -列表
        - ListView                      : 100% ( 33/ 33) -列表视图
        - MenuBar                       : 100% (  5/  5) -菜单条
        - pane                          : 100% (  9/  9) -Pane元素
        - ProgressBar                   : 100% (  8/  8) -进度条
        - RichEdit                      : 100% ( 42/ 42) -富文本编辑框
        - ScrollBar                     : 100% ( 16/ 16) -滚动条
        - ScrollView                    : 100% ( 29/ 29) -滚动视图
        - SliderBar                     : 100% ( 11/ 11) -滑动条元素
        - tabBar                        : 100% ( 21/ 21) -TabBar元素
        - TextLink                      : 100% (  6/  6) -静态文本连接按钮
        - ToolBar                       : 100% ( 11/ 11) -工具条
        - Tree                          : 100% ( 31/ 31) -列表树元素

    + 形状对象
        - Shape                         : 100% ( 15/ 15) -形状对象
        - ShapeText                     : 100% ( 13/ 13) -形状对象文本
        - ShapePicture                  : 100% (  6/  6) -形状对象图片
        - ShapeGif                      : 100% (  6/  6) -形状对象GIF
        - ShapeRect                     : 100% (  8/  8) -矩形形状对象
        - ShapeEllipse                  : 100% (  5/  5) -圆形(形状对象)
        - ShapeGroupBox                 : 100% ( 10/ 10) -组框(形状对象)
        - ShapeLine                     : 100% (  3/  3) -直线(形状对象)

    + 数据适配器
        - AdapterListView               : 100% ( 24/ 24) -数据适配器-列表视元素
        - AdapterTable                  : 100% ( 25/ 25) -数据适配器-XList-XListBox
        - AdapterMap                    : 100% (  7/  7) -数据适配器-单列Map
        - AdapterTree                   : 100% ( 18/ 18) -数据适配器-树元素

    + 其他模块
        - BkInfoManager                 : 100% (  9/  9) -背景内容管理器
        - Draw                          : 100% ( 59/ 59) -图形绘制
        - FontX                         : 100% (  7/  7) -炫彩字体
        - Image                         : 100% ( 31/ 31) -图片操作
        - LayoutObject                  : 100% ( 35/ 35) -布局对象

    + 事件
        - wm                            :  100% (210/210) -窗口事件
        - xe                            :  100% ( 76/ 76) -元素事件

    + 宏定义
        - XC_OBJECT_TYPE                :  100% ( 54/ 54) -接口句柄类型
        - HXCGUI                        :  100% ( 22/ 22) -句柄列表
        - xc_window_style               :  100% (  8/  8) -炫彩窗口样式
        - window_position_              :  100% (  7/  7) -窗口位置
        - window_transparent_           :  100% (  5/  5) -炫彩窗口透明标识
        - ID                            :  100% (  4/  4) -特殊ID
        - menu_state_flags              :  100% (  6/  6) -弹出菜单
        - menu_popup_position           :  100% (  8/  8) -弹出菜单
        - IDM                           :  100% (  6/  6) -弹出菜单 - 菜单ID
        - image_draw_type_              :  100% (  4/  4) -图片绘制类型
        - common_state3_                :  100% (  3/  3) -普通三种状态
        - button_state_                 :  100% (  5/  5) -按钮状态
        - button_type_                  :  100% (  6/  6) -按钮类型(用于区分功能)
        - button_style_                 :  100% ( 16/ 16) -按钮样式(用于区分外观)
        - comboBox_state_               :  100% (  3/  3) -ComboBox状态
        - list_item_state_              :  100% (  3/  3) -List项状态
        - tree_item_state_              :  100% (  2/  2) -Tree项状态
        - button_icon_align_            :  100% (  4/  4) -按钮图标对齐方式
        - layout_size_type_             :  100% (  5/  5) -布局大小类型
        - list_drawItemBk_flags_        :  100% (  4/  4) -List,ListBox,Tree,项背景绘制标志位
        - messageBox_flags_             :  100% (  3/  3) -弹出消息框

    + API                               :  100% ( 21/ 21) -全局API
    + UI                                :  100% ( 11/ 11) -UI设计器支持
    + RegEventC                         :  100% (  7/  7) -注册事件C - 在window和Element中已定义
    -----------------------------------------------------------------------------------------

    Total progress                      :  100% (1315/1315)
