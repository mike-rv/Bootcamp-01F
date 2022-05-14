package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for apple := 'z'; apple >= 'a'; apple-- {
		z01.PrintRune(apple)
	}

	z01.PrintRune('\n')
}
