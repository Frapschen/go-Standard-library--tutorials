package os_tutorial

import (
	"log"
	"os"
	"testing"
	"time"
)

func Test_file(t *testing.T) {
	//切换工作目录
	//err := os.Chdir("../fmt-tutorial")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(os.Getwd())

	//修改文件权限
	if err := os.Chmod("example-file", 0644); err != nil {
		log.Fatal(err)
	}

	//修改文件的所属
	//os.Chown()

	//修改文件的时间信息
	mtime := time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC)
	atime := time.Date(2007, time.March, 2, 4, 5, 6, 0, time.UTC)
	if err := os.Chtimes("example-file", atime, mtime); err != nil {
		log.Fatal(err)
	}

	//
}
