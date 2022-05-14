package piscine

func SplitWhiteSpaces(str string) []string {
	TextToString := ""
	t := []string{}
	for i, v := range str {
		if i == lent2(str)-1 && string(v) != " " && string(v) != "\t" && string(v) != "\n" {
			TextToString += string(v)
			t = append(t, TextToString)
		} else if string(v) != " " && string(v) != "\t" && string(v) != "\n" {
			TextToString += string(v)
		} else {
			if len(TextToString) >= 1 {
				t = append(t, TextToString)
			}
			TextToString = ""
		}
	}
	return t
}

func lent2(d string) int {
	inc := 0
	for range d {
		inc++
	}
	return inc
}

// Write a function that separates the words
// of a string and puts them in a string slice.

// The separators are spaces, tabs and newlines.
// package piscine

// for i := 0; i < len(s); i++ {
// 		result = s[i] + 1

// func Capitalize(s string) string {
//     runeArray := []rune(s)
//     for i, char := range runeArray {
//         if isNumberOrAlph(char) {
//             if i == 0  isNumberOrAlph(runeArray[i-1]) == false {
//                 if runeArray[i] >= 'a' && runeArray[i] <= 'z' {
//                     runeArray[i] = char - 32
//                 }
//             } else {
//                 if runeArray[i] >= 'A' && runeArray[i] <= 'Z' {
//                     runeArray[i] = char + 32
//                 }
//             }
//         }
//     }
//     return string(runeArray)
// }

// func isNumberOrAlph(r rune) bool {
//     if r >= 'a' && r <= 'z'  r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
//         return true
//     }
//     return false
// }
