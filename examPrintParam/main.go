package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args1 := len(os.Args[1:])
	a := itoa(args1)
	for _, v := range a {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')

}

func itoa(n int) []rune {
	output := []rune{}
	for n != 0 {
		digit := n % 10
		output = append([]rune{rune(digit + 48)}, output...)
		n /= 10
	}
	return output
}
