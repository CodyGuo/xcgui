package xc

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func downLoadXCGUIDll() {
	resp, err := http.Get("http://git.oschina.net/CodyGuo/xcgui/raw/master/examples/lib/XCGUI.dll")
	if err != nil {
		log.Fatal(err)
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
	} else {
		log.Printf("[ERROR] 下载失败,%s.\n", resp.Status)
	}

}
