package xc

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

import (
	"github.com/lxn/walk"
)

func downLoadXCGUIDll() {
	mw, _ := walk.NewMainWindow()
	ret := walk.MsgBox(mw, "下载提示", "没有找到XCGUI.dll，是否要联网下载？", walk.MsgBoxIconQuestion+walk.MsgBoxOKCancel)
	if ret != 1 {
		walk.MsgBox(mw, "警告信息", "请在重新运行程序前将XCGUI.dll放到运行目录.", walk.MsgBoxIconWarning)
		mw.Close()
		os.Exit(1)
	}

	resp, err := http.Get("http://git.oschina.net/CodyGuo/xcgui/raw/master/examples/lib/XCGUI.dll")
	if err != nil {
		log.Println("[ERROR] 下载失败，请检查网络.")
		walk.MsgBox(mw, "错误信息", "XCGUI.dll 下载失败，请检查网络.", walk.MsgBoxIconError)
		mw.Close()
	}

	if resp.StatusCode == http.StatusOK {
		log.Println("[INFO] 正在下载 XCGUI.dll .")

		downFile, err := os.Create("XCGUI.dll")
		if err != nil {
			log.Fatal(err)
		}
		// 不要忘记关闭打开的文件.
		defer downFile.Close()

		body, err := ioutil.ReadAll(resp.Body)
		io.Copy(downFile, bytes.NewReader(body))
		log.Println("[INFO] XCGUI.dll 下载成功.正在运行主程序.")

		walk.MsgBox(mw, "提示信息", "下载成功, 请点击确定继续运行程序.", walk.MsgBoxIconInformation)
		mw.Close()
	} else {
		log.Printf("[ERROR] 下载失败,%s.\n", resp.Status)

		walk.MsgBox(mw, "错误信息", "[ERROR] 下载失败,请检查网络.", walk.MsgBoxIconError)
		mw.Close()
	}

	mw.Run()
}
