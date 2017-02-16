package main

import (
	"fmt"
	"os"
)

func main() {
	//s, sep := "", ""
	var s string
	for _, sep := range os.Args[1:] {
		s += sep + " "
	}
	fmt.Println(s)
}
