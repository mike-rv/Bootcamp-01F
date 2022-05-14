package main

import "fmt"

func main() {
	fmt.Println(cap("hello! hOw aRe you? how+are+things_+4yOu?"))
}

func cap(sentence string) string {
	rSentence := []rune(sentence)

	letter := 0

	for i, char := range rSentence {
		if char >= 'a' && char <= 'z' && letter == 0 {
			rSentence[i] = rSentence[i] - 32' '
			letter++
		} else if char >= 'A' && char <= 'Z' && letter != 0 {
			rSentence[i] = rSentence[i] + 32
		}
		if char >= ' ' && char <= '@' || char >= '[' && char <= '\'' {
			letter = 0
		}
	}
	return string(rSentence)
}
