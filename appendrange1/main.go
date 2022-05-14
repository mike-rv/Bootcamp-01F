package main

import "fmt"

func main() {
	aR := appendrange1(5, 10)
	fmt.Println(aR)
	aR = appendrange1(10, 5)
	fmt.Println(aR)
}

func appendrange1(min, max int) []int {
	var intIndex []int
	if min > max {
		return nil
	}
	for i := min; i < max; i++ {
		intIndex = append(intIndex, i)
	}
	return intIndex
}

// cannnot use make function; use *append* instead

// for min > max, declare intIndex with no value type []int
// use a for loop with min as initializer
// and max as end test condition.
// for code block, append i to intIndex = intIndex
// every iteration int index increases
