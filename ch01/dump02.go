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
	if len(files) == 0 {
		coutLinesFromStdin(os.Stdin, couts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dump02: %v\n", err)
				continue
			}
			coutLinesFromFile(f, couts)
			f.Close()
		}
	}
	for line, n := range couts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func coutLinesFromStdin(f *os.File, counts map[string]int) {
	input := bufio.NewReader(f)
	text, err := input.ReadString('\n')
	if err == nil {
		fmt.Printf("input length is: %d\n", len(text))
		text = text[0 : len(text)-1]
		fmt.Println(text)
		units := strings.Split(text, " ")
		for _, myunit := range units {
			counts[myunit]++
		}
	}
}

func coutLinesFromFile(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
