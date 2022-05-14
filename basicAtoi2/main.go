package main

import (
	"fmt"
	// "piscine"
)

func main() {
	// fmt.Println(piscine.BasicAtoi("12345"))
	// fmt.Println(piscine.BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("Hello"))
	fmt.Println(BasicAtoi("12 345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
	fmt.Println(BasicAtoi("543210000"))

}
func BasicAtoi(s string) int {
	var n int
	for _, r := range s {
		if r >= ' ' && r <= '/' || r >= ':' && r <= '~' {
			return 0
		} else {
			n = (n * 10) + int(r) - 48
		}
	}
	return n
}
