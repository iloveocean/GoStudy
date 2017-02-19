package main

import (
	"fmt"
	"os"
)

func main() {
	//s, sep := "", ""
	var s string
	for i, sep := range os.Args[1:] {
		s += sep + " "
		fmt.Printf("index: %d\tvalue: %s\n", i, sep)
	}
	fmt.Println(s)
}
