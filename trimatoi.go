package piscine

func TrimAtoi(s string) int {
	result := 0
	negative := 1
	for _, c := range s { //"012 345"
		if c >= '0' && c <= '9' {
			result *= 10
			result += int(c - '0')
		}
		if c == '-' && result == 0 {
			negative *= -1
		}
	}
	return result * negative
}
