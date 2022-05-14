package piscine

func MakeRange(min, max int) []int {
	var sliceofintszero []int
	if min >= max {
		return sliceofintszero[0:0]
	}
	sliceofints := make([]int, (max - min))
	count := min
	for i := 0; i < max-min; i++ {
		sliceofints[i] = count
		count++
	}
	return sliceofints
}

// Write a function that takes an int min and an int max as parameters.
// The function must return a slice of ints with all the values between min and max.

// Min is included, and max is excluded.

// If min is greater than or equal to max, a nil slice is returned.

// append is not allowed for this exercise.

// if min < max {
// 	answer := make([]int, max)
// for i := min - 1; i < max-1; i++ {
// 	answer[i] = 1 + i
// 	return answer[min-1 : max-1]
// }

// return nil
