// package main

// import "fmt"

// func main() {
// 	a := []rune("faya")
// 	b := []rune("fgvvfdxcacpolhyghbred")
// 	c := []rune{}

// 	for i := 0; i < len(a); i++ {
// 		for j,v := range b {
// 			if a[i] == v {
// 				c = append(c, a[i]) //ffaayyaa
// 				b = b[j:]
// 				break
// 			}
// 		}
// 	}
// 	// e := unique(c)
// 	for _, v := range c {
// 		fmt.Print(string(rune(v)))
// 	}
// 	fmt.Println()
// }

/*
func unique(s []rune) []rune {
	inResult := make(map[rune]bool)
	var result []rune
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
*/

package main

/*
• The Check Point Exam (monday 7 march 2022) Program;
filename: ChkPontExamp.go; author: Kamal H. Zada; Date: 7 march 2022;
• This main program takes two strings as Arguments
  The program displays the first string only if chars in the first string are contained in the second string;
  The order of chars in the first string must be in order in the second string;
  At the end do z01.PrintRune('\n')

Examples:
Go Run . "123" "123"
output: 123

Go Run . "faya" "fra in your backalong"
output: faya

Go Run . "faya" " a front at your convenience"
displays nothing.
*/

import (
	"fmt"
	"os"
)

func main() {
	var arguments []string
	var numOfArgs int = 0
	var inpStr1 string
	var inpStr2 string
	var str1Len int = 0
	var str2Len int = 0
	var idx1 int = 0
	var idx2 int = 0
	var chrFound bool = false
	var allOK bool = true

	fmt.Println("The Check Point Exam Question 2 entered.")
	arguments = os.Args
	numOfArgs = len(os.Args)
	if numOfArgs == 3 {
		/*correct number of Arguments*/
		fmt.Println("Arg 1:", arguments[1], "; Arg 2:", arguments[2], "; Number of Arguments:", numOfArgs)
		inpStr1 = os.Args[1]
		inpStr2 = os.Args[2]
		str1Len = len(inpStr1)
		str2Len = len(inpStr2)
		for idx1 < str1Len && allOK {
			for idx2 < str2Len && !chrFound {
				if inpStr1[idx1] == inpStr2[idx2] {
					chrFound = true
				} else {
					chrFound = false
					idx2++
				} /*if*/
			} /*while loop*/
			/*ask why it dropped out of the while loop*/
			switch {
			case chrFound:
				/*fmt.Println("Char at IDX1:", idx1, "found in string 2 at Pos(IDX2):", idx2)*/
				chrFound = false
				allOK = true
				idx1++
			case idx2 == str2Len:
				chrFound = false
				allOK = false
			default:
			} /*switch*/
		} /*outer for loop*/
		if allOK {
			fmt.Println("All OK, Input string 1:", inpStr1)
		} else {
			fmt.Println("All Not OK, Input string cannot be displayed.")
		} /*if*/
	} else {
		/*Error: wrong number of Arguments*/
	} /*if*/
} /*main*/
*/
