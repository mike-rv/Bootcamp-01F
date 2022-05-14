package piscine

func NRune(s string, n int) rune {
	r := []rune(s)
	// l := len(r)
	if n > 0 && n <= len(r) {
		return r[n-1]
	}
	return 0
}

// var nr rune
// 	for _, nr := range r {
// 		nr = nr + r[n-1]
// 	}
// 	return nr - 1

// var nr rune
// 	for _, nr := range r {
// 		nr = nr + r[n-1]
// 	}
// 	return nr - 1
// result :=

// r := []rune(s)

// if r > 0 || len(r) {
// 	return  r[n-1]
// }
// return 0

// if n <= 0 || n >= len(s) {
// 	return 0
// }
// r := []rune(s)
// result := r[n-1]
// return result
