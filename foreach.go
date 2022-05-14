package piscine

func ForEach(f func(int), a []int) {
	for _, a1 := range a {
		f(a1)
	}
}

// Write a function ForEach that, for an int slice,
// applies a function on each element of that slice.
