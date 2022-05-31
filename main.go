package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Chdir("os-tutorlal")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(os.Getwd())
}
