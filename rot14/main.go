// package main

// import (
// 	"github.com/01-edu/z01"
// 	"os"
// )

// func main() {
// 	str := os.Args[1]

// 	for _, r := range str {
// 		if r == ' ' {
// 			z01.PrintRune(' ')
// 		}
// 		if r >= 'A' && r < 'M' || r >= 'a' && r < 'm' {
// 			z01.PrintRune(r + 14)
// 		}
// 		if r >= 'M' && r <= 'Z' || r >= 'm' && r <= 'z' {
// 			z01.PrintRune(r - 12)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }
/*
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	str := os.Args[1]

	for _, r := range str {
		if r == ' ' {
			z01.PrintRune(' ')
		}
		if r > 'a' && r < 'm' || r > 'A' && r < 'M' {
			z01.PrintRune(r + 14)
		}
		if r > 'm' && r < 'z' || r > 'M' && r < 'Z' {
			z01.PrintRune(r - 12)
		}
	}
	z01.PrintRune('\n')
}

package main

import (
	"github.com/01-edu/z01"
	// "piscine"
)
*/
// func main() {
// 	result := Rot14("Hello! How are You?")

// 	for _, r := range result {
// 		z01.PrintRune(r)
// 	}
// 	z01.PrintRune('\n')
// }

// func Rot14(s string) []rune {
// 	runes := []rune{}
// 	for _, v := range s {
// 		if v < 'a' && v > 'z' || v < 'A' && v > 'Z' {
// 			runes = append(runes, v)
// 		}
// 		if v > 'a' && v < 'm' || v > 'A' && v < 'M' {
// 			runes = append(runes, v+14)
// 		} else if v > 'm' && v < 'z' || v > 'M' && v < 'Z' {
// 			runes = append(runes, v-12)
// 		}
// 	}
// 	return runes
// }
package main

import (
	"piscine"
	"github.com/01-edu/z01"
)

func main() {
	result := piscine.Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
And its output :

$ go run .
Vszzc! Vck ofs Mci?
$