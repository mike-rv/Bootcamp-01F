package main

import "github.com/01-edu/z01"

type points struct {
	x string
	y string
}

func setPoint(ptr *points) {
	ptr.x = "42"
	ptr.y = "21"
}

func main() {
	points := points{}
	setPoint(&points)
	msg := "x = " + points.x + "," + " y = " + points.y + "\n"

	for _, pointy := range msg {
		z01.PrintRune(pointy)
	}
}

// Create a new directory called point.

//     The code below must be copied into a file called main.go inside the point directory.

//     The necessary changes must be applied so that the program works.

// $ go run .
// x = 42, y = 21
// $
