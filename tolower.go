package piscine

func ToLower(s string) string {
	m := ""
	for _, letter := range s {
		if letter >= 65 && letter <= 90 {
			m += string(letter + 32)
		} else {
			m += string(letter)
		}
	}
	return m
}
