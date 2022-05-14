package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(k string) {
	for _, value := range k {
		z01.PrintRune(value)
	}
}

// x := []rune(k)
// for i := 0; i < len(x); i++ {
// z01.PrintRune(x[i])
// }
