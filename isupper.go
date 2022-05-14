package piscine

func IsUpper(s string) bool {
	for _, letter := range s {
		if letter < 'A' || letter > 'Z' {
			return false
		}
	}
	return true
}

// Write a function that returns true if the string passed as parameter
// contains only uppercase characters, otherwise, returns false.

// count := 0

// 	for _, letter := range s {
// 		count++
// 		if letter  >= 'A' && r <= 'Z' && != '!'{
// 			return true
// 		} else {
// 			return false
// 		}
// 	}
// 	return false

// r := []rune(s)
// 	for i := 0 ; i < len(r) ; i++ {
// 		if r >= 'A' && r <= 'Z' &&  {
