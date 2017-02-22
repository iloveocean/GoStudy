package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("open file: %s failed, error: %v", os.Args[1], err)
		return
	}
	src := string(data)
	re, err := regexp.Compile(`<[\S\s]+?>`)
	if err != nil {
		fmt.Printf("compile reg exp failed: %v", err)
	}
	results := re.FindAllString(src, -1)
	fmt.Printf("find total %d strings\n", len(results))
	for _, result := range results {
		fmt.Println(result)
	}
}
