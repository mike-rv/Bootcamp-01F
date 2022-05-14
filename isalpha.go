package piscine

func IsAlpha(s string) bool {
	for _, letter := range s {
		if letter < '0' || letter > '9' && letter < 'A' || letter > 'Z' && letter < 'a' || letter > 'z' {
			return false
		}
	}
	return true
}

// Write a function that returns true if the string passed as the parameter only contains
// alphanumerical characters or is empty, otherwise, and returns false.

//
//if letter >= '0' || letter <= '9' || letter >= 'A' || letter <= 'Z' || letter >= 'a' || letter <= 'z' || letter != ' '
