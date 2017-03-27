package main

import (
	"fmt"
	"reflect"
)

type student struct {
	name int
	age  int
}

func reflectType(input student) {
	t := reflect.TypeOf(input)
	fmt.Println(t)
	val := reflect.New(t)
	fmt.Printf("val type is: ")
	fmt.Println(val.Elem())
	//fmt.Printf("val.name is %s\n", val.name)
}

func main() {
	stu := student{7, 6}
	reflectType(stu)
}
