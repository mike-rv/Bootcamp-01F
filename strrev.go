package piscine

func StrRev(s string) string {
	var r string
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return r
}

// chars := []rune(s)
// 	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
// 		chars[i], chars[j] = chars[j], chars[i]
// 	}
// 	return string(chars)
