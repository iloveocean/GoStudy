package main

import "fmt"

func main() {
	var a [2]int = [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{2, 3}
	fmt.Println(a == b, b == c, a == c)

	//compile error: different type can not be
	//var d [3]int = [3]int{7, 8, 9}
	//fmt.Println(a == d)
}
