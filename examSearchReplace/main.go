package main

import (
	"fmt"
	"os"
	// "github.com/01-edu/z01"
)

// Search and Replace

func main() {
	a := os.Args[1]
	b := os.Args[2]
	c := os.Args[3]
	// z01.PrintRune(sNr(a, b, c))
	str := ""             // declare empty string
	for i, r := range a { // range thru args1 with index and rune
		if a[i] != b[0] { // if byte (individual char of args1 string) is not equal to byte of args2
			str += string(r) // concatanate the rune converted to string to new string
		} else {
			str += c // otherwise replace with args3
		}
	}
	fmt.Println(str)
}

// package main

// import "fmt"

// func main() {
// 	a := "abcd"
// 	b := "z"
// 	c := "l"
// 	fmt.Println(snd(a, b, c))
// }

// func snd(a, b, c string) string {
// 	x := []rune(b)
// 	y := []rune(c)
// 	var d string
// 	dr := []rune(d)
// 	for _, r := range a {
// 		for _, br := range x {
// 			if br == r {
// 				for _, cr := range y {
// 					r = cr
// 					dr = append(dr, r)
// 					break
// 				}
// 			} else {
// 				dr = append(dr, r)
// 				break
// 			}
// 		}
// 	}
// 	return string(dr)
// }
