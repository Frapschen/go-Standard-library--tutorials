package io_tutorial

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func Test_io(t *testing.T) {
	//从 src 读取数据到 dst
	r := strings.NewReader("some io.Reader stream to be read\n")
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

	// buff 数据拷贝
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)

	// buf is used here...
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	// ... reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}

	//拷贝指定的byte
	r = strings.NewReader("some io.Reader stream to be read")
	if _, err := io.CopyN(os.Stdout, r, 4); err != nil {
		log.Fatal(err)
	}

	//pipe
	rr, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n")
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, rr); err != nil {
		log.Fatal(err)
	}

	//读取所有内容
	r = strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)

	//至少读取
	r = strings.NewReader("some io.Reader stream to be read\n")
	buf = make([]byte, 14)
	if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// buffer smaller than minimal read size.
	shortBuf := make([]byte, 3)
	if _, err := io.ReadAtLeast(r, shortBuf, 4); err != nil {
		fmt.Println("error:", err)
	}

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadAtLeast(r, longBuf, 64); err != nil {
		fmt.Println("error:", err)
	}

	//填满buffer
	r = strings.NewReader("some io.Reader stream to be read\n")
	buf = make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// minimal read size bigger than io.Reader stream
	longBuf = make([]byte, 64)
	if _, err := io.ReadFull(r, longBuf); err != nil {
		fmt.Println("error:", err)
	}

	//将string 写入 Writer
	if _, err := io.WriteString(os.Stdout, "Hello World"); err != nil {
		log.Fatal(err)
	}
}
