package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello word!")
	var scanValue string
	fmt.Fscan(os.Stdin, &scanValue)
	fmt.Println(scanValue)
}
