package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	// Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	// Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	// Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	var twoD [][]int
	for i := 0; i < len(slice); i++ {
		if i < size {
			twoD = append(twoD, slice[i, ])
		}
	}
}

// a := make([][]uint8, size) // initialize a slice of dy slices
// for i := 0; i < size; i++ {
// 	a[i] = make([]uint8, size) // initialize a slice of dx unit8 in each of dy slices
// 	fmt.Println(a[i])
// }
/*
	define 2d array
	for loop with i += size
	define end // i  + size
	if end > size
	end = size
	append 2d array with slice[i : end]
*/