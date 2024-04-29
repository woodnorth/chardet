package chardet

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestX(t *testing.T) {
	fd, _ := os.Open("/Users/chenxiaoxi/PycharmProjects/pythonProject/chardet/tests/EUC-JP/_mozilla_bug426271_text-euc-jp.html")
	data, _ := io.ReadAll(fd)
	fmt.Println(len(data))
	r := Detect(data)
	fmt.Println(r)
}

func TestY(t *testing.T) {
	filepath.Walk("/Users/chenxiaoxi/PycharmProjects/pythonProject/chardet/tests/utf-8/", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fd, _ := os.Open(path)
		data, _ := io.ReadAll(fd)
		r := Detect(data)
		fmt.Println(r)
		return nil
	})

}

func TestZ(t *testing.T) {
	req, err := http.NewRequest("GET", "https://zhuanlan.zhihu.com/p/641187337", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	data, err := io.ReadAll(rsp.Body)
	if err != nil {
		return
	}
	fmt.Println(DetectBest(data))
	fmt.Println(DetectAll(data, false, true))
}
