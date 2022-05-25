package os_tutorial

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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

	//清除环境变量
	os.Clearenv()

	//将一个dir打开问一个文件目录
	fs := os.DirFS("/Users/fraps/daocloud/github-code/go-Standard-library-tutorials/os")
	fmt.Println(fs)

	//返回所有环境变量
	envs := os.Environ()
	fmt.Println(envs)

	//返回执行当前进程的文件
	fmt.Println(os.Executable())

	//os.Exit(1)

	//字符串替代
	mapper := func(placeholderName string) string {
		switch placeholderName {
		case "DAY_PART":
			return "morning"
		case "NAME":
			return "Gopher"
		}

		return ""
	}
	fmt.Println(os.Expand("Good ${DAY_PART}, $NAME!", mapper))

	//用环境变量替换
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")
	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))

	//返回 有效用户组id
	fmt.Println(os.Getegid())

	//获取环境变量
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")
	fmt.Printf("%s lives in %s.\n", os.Getenv("NAME"), os.Getenv("BURROW"))

	//获取 用户id
	fmt.Println(os.Geteuid())

	//返回 用户组id
	fmt.Println(os.Getgid())

	fmt.Println(os.Getgroups())

	fmt.Println(os.Getpagesize())

	fmt.Println(os.Getpid())

	fmt.Println(os.Getppid())

	fmt.Println(os.Getuid())

	fmt.Println(os.Getwd())

	fmt.Println(os.Hostname())

	//这个两个函数 比较有意思 检测错误？
	//fmt.Println(os.IsExist())
	//fmt.Println(os.IsNotExist())

	//c 是否是文件夹分割符号
	fmt.Println(os.IsPathSeparator('/'))

	// 是否有权限？
	//fmt.Println(os.IsPermission())

	//timeout error
	//fmt.Println(os.IsTimeout())

	//修改文件的所属
	//fmt.Println(os.Lchown())

	//修改文件的Link
	//fmt.Println(os.Link())

	//查找是否有特定的环境变量
	show := func(key string) {
		val, ok := os.LookupEnv(key)
		if !ok {
			fmt.Printf("%s not set\n", key)
		} else {
			fmt.Printf("%s=%s\n", key, val)
		}
	}

	os.Setenv("SOME_KEY", "value")
	os.Setenv("EMPTY_KEY", "")

	show("SOME_KEY")
	show("EMPTY_KEY")
	show("MISSING_KEY")

	//创建文件夹
	err := os.Mkdir("testdir", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	err = os.WriteFile("testdir/testfile.txt", []byte("Hello, Gophers!"), 0660)
	if err != nil {
		log.Fatal(err)
	}

	//创建文件路径
	err = os.MkdirAll("test/subdir", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	err = os.WriteFile("test/subdir/testfile.txt", []byte("Hello, Gophers!"), 0660)
	if err != nil {
		log.Fatal(err)
	}

	//创建临时文件夹 dir "./" 当前目录下
	dir, err := os.MkdirTemp("./", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir) // clean up

	file := filepath.Join(dir, "tmpfile")
	if err := os.WriteFile(file, []byte("content"), 0666); err != nil {
		log.Fatal(err)
	}

	logsDir, err := os.MkdirTemp("", "*-logs")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(logsDir) // clean up

	// Logs can be cleaned out earlier if needed by searching
	// for all directories whose suffix ends in *-logs.
	globPattern := filepath.Join(os.TempDir(), "*-logs")
	matches, err := filepath.Glob(globPattern)
	if err != nil {
		log.Fatalf("Failed to match %q: %v", globPattern, err)
	}

	for _, match := range matches {
		if err := os.RemoveAll(match); err != nil {
			log.Printf("Failed to remove %q: %v", match, err)
		}
	}
}
