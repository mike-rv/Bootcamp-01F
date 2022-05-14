package piscine

func ConcatParams(args []string) string {
	// \/ declare string index count at 0 (1st element)
	var result string = args[0]
	// \/ for range over string index starting from 2nd element to last element
	for _, word := range args[1:] {
		// /\ word are the individual elements of the string index hence the string type
		result = result + "\n" + word
	} // /\ for every iteration each element from the the 1st to last is returned with a new line
	return result
}

// argsfrom1 := args[0:]

// 	return ""
// // var answer []string
// 	for _, word := range args {
// 		count := 0
// 		return word
// 		count++
// 	}
// // 	return ""
// 	// args := os.Args[1:]
// 		for _, word := range args {
// 			for _, char1 := range word {
// 				return string(char1)
// 				// fmt.Println(rune(char))
// 			}
// 		}
// 		return ""
// answer := make([]string, len(args))
// 	for i := 0; i < len(args); i++ {
// 		return answer[i]
// 	}
// 	return ""
