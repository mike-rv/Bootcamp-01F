package piscine

func ToUpper(s string) string {
	m := ""
	for _, letter := range s {
		if letter >= 97 && letter <= 122 {
			m += string(letter - 32)
		} else {
			m += string(letter)
		}
	}
	return m
}

// Write a function that capitalizes each letter in a string.
