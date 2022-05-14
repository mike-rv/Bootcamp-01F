package main

import "fmt"

func Index(s string, toFind string) int {
	for i := range s {
		if len(toFind) < len(s[i:]) {
			if string(s[i:i+len(toFind)]) == toFind {
				return i
			}
		}
	}
	return -1
}

func main() {
	// a :=
	fmt.Println(Index("Hello!", "l"))
}
func Index2(s string, toFind string) int {

	j := []rune(s)
	l := []rune(toFind)
	n := len(j)
	k := len(l)

	for i := 0; i <= n-k; i++ {
		if toFind == s[i:i+k] {
			return (i)
		}
	}
	return -1
}
