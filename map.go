package piscine

func Map(f func(int) bool, a []int) []bool {
	var boolindex []bool
	for _, number := range a {
		boolindex = append(boolindex, f(number))
	}
	return boolindex
}
