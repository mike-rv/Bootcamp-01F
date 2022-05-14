package main

import "fmt"

func main() {
	aR := makerange1(5, 10)
	fmt.Println(aR)
	aR = makerange1(10, 5)
	fmt.Println(aR)
}

func makerange1(min, max int) []int {
	if min >= max {
		return nil
	}
	sliceOfInts1 := make([]int, max-min)
	count := min
	for i := 0; i < max-min; i++ {
		sliceOfInts1[i] = count
		count++
	}
	return sliceOfInts1
}

// min < max declare make function sliceOfInts type []int, max - min.
// declare a count value of min
// make a for loop with condition max - min
// inside codeblock, sliceOfInts with [] of i = count
// count will increase for every iteration
