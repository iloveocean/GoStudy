package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Printf("json.Unmarshal met error: %v\n", err)
		return
	}
	m := f.(map[string]interface{})
	for k, v := range m {
		fmt.Printf("key is: %s ", k)
		vType := reflect.TypeOf(v)
		switch concrete := v.(type) {
		case []interface{}:
			fmt.Printf("value is array\n")
			for i, element := range concrete {
				fmt.Printf("element %d of this array is %v\n", i, element)
			}
		default:
			fmt.Printf("value is: %v type is: %v\n", v, vType)
		}
	}
}
