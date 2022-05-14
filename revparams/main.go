package main

import (
	//"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for i := len(args) - 1; i >= 0; i-- {
		for _, word2 := range args[i] {
			z01.PrintRune(word2)
			// fmt.Println(s[i])
		}
		z01.PrintRune('\n')
	}
}

// args := os.Args[1:]
// for _, word := range args {
// 	for _, char1 := range word {
// 		// char1 = len(word) - 1 - char1
// 		z01.PrintRune(char1)
// 		// fmt.Println(char1(len(char1) - 1 - i))
// 	}
// }
// z01.PrintRune('\n')
// args := os.Args
// revchar1 := [len(args)]int{}
// for i := range args {
// 	revchar1 = [len(args)]int{}
// 	// char1 = len(word) - 1 - char1
// 	z01.PrintRune(rune(revchar1[len(args)-1-i] == args[i]))
// 	// fmt.Println(char1(len(char1) - 1 - i))
// }
// // z01.PrintRune('\n')
