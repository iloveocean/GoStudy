package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	couts := make(map[string]int)
	files := os.Args[1:]
	var dupFiles []string
	if len(files) == 0 {
		myCountLinesFromStdin(os.Stdin, couts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dump02: %v\n", err)
				continue
			}
			if hasDup := myCountLinesFromFile(f, couts); hasDup {
				dupFiles = append(dupFiles, file)
			}
			f.Close()
		}
	}
	for line, n := range couts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	for _, dupfile := range dupFiles {
		fmt.Printf("file: %s contains duplicated line(s).\n", dupfile)
	}
}

func myCountLinesFromStdin(f *os.File, counts map[string]int) {
	input := bufio.NewReader(f)
	text, err := input.ReadString('\n')
	if err != nil {
		fmt.Printf("read string error: %v\n", err)
		return
	}
	fmt.Printf("input length is: %d\n", len(text))
	text = text[0 : len(text)-1]
	fmt.Println(text)
	units := strings.Split(text, " ")
	for _, myunit := range units {
		counts[myunit]++
	}
}

func myCountLinesFromFile(f *os.File, counts map[string]int) bool {
	var hasDupLine bool
	input := bufio.NewScanner(f)
	for input.Scan() {
		if _, ok := counts[input.Text()]; ok {
			hasDupLine = true
		}
		counts[input.Text()]++
	}
	return hasDupLine
}
