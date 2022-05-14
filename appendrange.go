package piscine

func AppendRange(min, max int) []int {
	var answer []int
	if min > max {
		return answer[0:0]
	}
	for i := min; i < max; i++ {
		answer = append(answer, i)
	}
	return answer
}
