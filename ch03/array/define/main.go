package main

import "fmt"

type Currency int

const (
	_ int = iota
	USD
	EUR
	GBP
	RMB
)

func main() {
	money := [...]string{USD: "USD", EUR: "EUR", GBP: "GBP", RMB: "RMB"}
	fmt.Printf("money length is: %d\n", len(money))
	fmt.Printf("money[2] value is: %s\n", money[2])
}
