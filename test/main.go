package main

import "fmt"

func main() {
	a := "abcd"
	b := "z"
	c := "l"
	fmt.Println(snd(a, b, c))
}

func snd(a, b, c string) string {
	x := []rune(b)
	y := []rune(c)
	var d string
	dr := []rune(d)
	for _, r := range a {
		for _, br := range x {
			if br == r {
				for _, cr := range y {
					r = cr
					dr = append(dr, r)
					break
				}
			} else {
				dr = append(dr, r)
				break
			}
		}
	}
	return string(dr)
}
