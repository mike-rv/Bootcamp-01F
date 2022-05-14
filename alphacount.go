package piscine

func AlphaCount(s string) int {
	count := 0
	for _, letter := range s {
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
			count++
		}
	}

	return count
}

// c := len(s.replace(" ", ""))
// r := []rune(s)
// var nr rune

// for i := 'A'; i <= 'Z'; i++ {
// 	r++
// 	' {
// 		result = result + r
// 	}
// }
// return r
// result := nr
// 			result1 := string(result)
// 			return len(result1)
