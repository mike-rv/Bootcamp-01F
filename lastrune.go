package piscine

func LastRune(s string) rune {
	last := []rune(s)
	last1 := last[len(last)-1]

	return last1
}

// var rx rune
// s1 := rune(s)

// rx = s1[len(s1)-1:]

// return rx

// var lastR rune
// r := []rune(s)
// lastR = string(r[len(r)-1:])
// return lastR

// j := len(s)
// for i:= 0 ; i <= len(s) ; i++ {
// 	_, last :=
