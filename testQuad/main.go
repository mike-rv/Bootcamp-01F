package main

import "fmt"

func main() {
	quad(5, 3)
}

func quad(x, y int) {
	for i := 0; i <= x-2; i++ {
		for j := 0; j <= y-2; j++ {
			if i == 0 {
				fmt.Printf("o")
			} else if i == x {
				fmt.Printf("o\n")
			} else if i != x || i != 0 {
				fmt.Printf("-")
			} else if j == 0 {
				fmt.Printf("|")
			} else if j == j {
				fmt.Printf("|")
			} else if j != x || j != 0 {
				fmt.Printf(" ")
			}
		}
	}
}
