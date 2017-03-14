package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {

}

func main() {
	fmt.Println("vim-go")
}
