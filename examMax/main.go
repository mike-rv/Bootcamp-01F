package main

import (
	"fmt"
	"github.com/01-edu/z01"
)

func main() {
	n := []int{5, 10, 6, 4, 8}
	a := max(n)
	z01.PrintRune(a)
}
func max(n []int) rune {
	max := 0
	for i := 0; i < len(n); i++ {
		if n[i] > max {
			max = n[i]
		}
	}
	return max
}
