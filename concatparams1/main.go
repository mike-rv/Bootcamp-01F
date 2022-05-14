package main

import (
	"fmt"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams1(test))
}

func ConcatParams1(args []string) string {
	var result string = args[0]
	for _, word := range args[1:] {
		result = result + "\n" + word
	}
	return result
}
