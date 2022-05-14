package main

import (
	"fmt"
	// "piscine"
)

func main() {
	// fmt.Println(piscine.BasicAtoi("12345"))
	// fmt.Println(piscine.BasicAtoi("0000000012345"))
	// fmt.Println(piscine.BasicAtoi("000000"))
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
	fmt.Println(BasicAtoi("54321"))
}
func BasicAtoi(s string) int {
	var n int
	for _, r := range s {
		n = (n * 10) + int(r) - 48
	}
	return n
}

/*
func BasicAtoi(s string) int {
	b := 0
	for _, n := range s {
		a := 0
		for i := '1'; i <= n; i++ {
			a++
		}
		b = b*10 + a
	}
	return b
}
"12345"

line 20: a= 0
line 21: i= '1' <=  n='1'
line 22: a= 1
exit loop
line 24: 0 x 10 + 1 = 1, b = 1

line 20: a= 1
line 21: i= 2
*/
