package piscine

func CountIf(f func(string) bool, tab []string) int {
	// var integer int
	count := 0
	for _, word := range tab {
		if f(word) {
			count++
		}
	}
	return count
}

// Write a function CountIf that returns, the number of elements
// of a string slice, for which the f function returns true.
