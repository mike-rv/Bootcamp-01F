package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	l := 0
	l = len(args)

	if l > 1 {
		fmt.Println("Too many arguments")
	} else if l == 0 {
		fmt.Println("File name missing")
	} else if args[0] == "quest8.txt" {

		input, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf(string(input))

	}
}
