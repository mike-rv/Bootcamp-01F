package piscine

func IsPrintable(s string) bool {
	for _, letter := range s {
		if letter < ' ' || letter > '~' {
			return false
		}
	}
	return true
}

// Write a function that returns true if the string passed as a parameter
// contains only printable characters, otherwise, returns false.
