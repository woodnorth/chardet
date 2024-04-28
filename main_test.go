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
	filepath.Walk("/Users/chenxiaoxi/PycharmProjects/pythonProject/chardet/tests/windows-1250-czech/", func(path string, info fs.FileInfo, err error) error {
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
	rsp, err := http.Get("https://www.chinanews.com.cn/gn/2020/06-24/9221419.shtml")
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
