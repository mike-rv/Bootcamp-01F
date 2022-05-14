// hello
// displayfirstparam
// displaylastparam
// paramcount
// displayalrevm
// displayalpham
// printdigits
// countdown
// strlen
// max
// wdmatch
// firstrune
// lastrune
// rot13
// lastword
// reduceint
// chunk
// searchreplace
// tabmult
// alphamirror
// compare
// doop
// findprevprime
// foldint
// romannumbers
// ispowerof2
// union
// inter
// printbits
// swapbits
// reversebits
// piglatin
// repeatalpja
// sortwordarr
// printhex
// gcd
// hiddenp
// rostring
// split
// revwstr
// fprime
// addprimesum
// atoibase
// itoa
// printmemory
// rpncalc
// brackets
// listsize
// options
// listremoveif
// itoabase
// brainfuck

package main

import (
	// "github.com/01-edu/z01"
	// "piscine"
	"fmt"
)

func main() {
	u := []int{5, 4, 6, 7, 5, 12, 45, 33, 30, 78, 56, 34, 54}
	z := sort(u)
	fmt.Println(z[7])
	fmt.Println(z)
}

func sort(nums []int) []int {
	// newNums := []int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

//
