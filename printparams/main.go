package main

import (
	"os"

	// "fmt"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for _, word := range args {
		for _, char1 := range word {
			z01.PrintRune((char1))
			// fmt.Println(rune(char))
		}
		z01.PrintRune('\n')
	}
}
