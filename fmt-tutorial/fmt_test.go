package fmt_tutorial

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

/*
fmt包主要实现了c语言的 printf scanf.
*/
func Test_print_go(t *testing.T) {
	//Sprint系列 将string输出到一个变量
	sp1 := fmt.Sprint("Sprint: learn fmt")
	sp2 := fmt.Sprintln("Sprintln: kearn fmt")
	sp3 := fmt.Sprintf("Sprintf: learn fmt %v", "test")
	fmt.Println(sp1, sp2, sp3)

	//Fprint系列 指定一个 io.Writer 作为参数输出
	fmt.Fprint(os.Stdout, "Fprint learn")
	fmt.Fprintln(os.Stdout, "Fprintln learn")
	fmt.Fprintf(os.Stdout, "Fprintf: learn fmt %v", "test")

	//print系列
	fmt.Print("Print learn")
	fmt.Println("Println learn")
	fmt.Printf("Printf :learn fmt %v", "test")
}

func Test_scan_go(t *testing.T) {

	//fscan系列
	file, _ := ioutil.ReadFile("fscan.txt")
	var fscan1, fscan2, fscan3 string
	fmt.Fscan(bytes.NewReader(file), &fscan1, &fscan2, &fscan3)
	fmt.Println(fscan1, fscan2, fscan3)

	//fscanln 遇到换行和EOF停止
	file, _ = ioutil.ReadFile("fscanln.txt")
	fscan1 = ""
	fscan2 = ""
	fscan3 = ""
	fmt.Fscanln(bytes.NewReader(file), &fscan1, &fscan2, &fscan3)
	fmt.Println(fscan1, fscan2, fscan3)

	//Fscanf 可以转换成指定的格式
	var (
		i int
		b bool
		s string
	)
	r := strings.NewReader("5 true gophers")
	n, err := fmt.Fscanf(r, "%d %t %s", &i, &b, &s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}
	fmt.Println(i, b, s)
	fmt.Println(n)

	//Scan系列 从 os.studin 读取数据
	//fmt.Scan(&fscan1, &fscan2, fscan3)
	//fmt.Scanln()
	//fmt.Scanf()

	//Sscan系列
	fscan1 = ""
	fscan2 = ""
	fscan3 = ""
	fmt.Sscan("a analogous b", &fscan1, &fscan2, &fscan3)
	fmt.Println(fmt.Println(fscan1, fscan2, fscan3))

	fscan1 = ""
	fscan2 = ""
	fscan3 = ""
	fmt.Sscanln("a analogous\nb", &fscan1, &fscan2, &fscan3)
	fmt.Println(fmt.Println(fscan1, fscan2, fscan3))

	//fmt.Sscanf() //可以指定格式了

}
