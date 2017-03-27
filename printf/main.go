package main

import "fmt"

func main() {
	peter := "Peter"
	peter = fmt.Sprintf("\"%s\"", peter)
	fmt.Println(peter)
}
