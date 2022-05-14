package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) < 1 {
		return // return when thr length is 0
	}
	if os.Args[1] == "" {
		fmt.Print()
	} else if os.Args[1] == "\\n" {
		fmt.Print(strings.Replace(os.Args[1], "\\n", "", 1))
		fmt.Println()
	} else {
		newLine(os.Args[1])
	}
}

func ascii(str string) string {
	input := str
	readFile, err := os.Open("standard.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()
	num := 0
	art := ""
	for i := 0; i < 8; i++ {
		for _, runes := range input {
			asciiVal := ((runes - 32) * 9)
			txtLineNum := int(asciiVal) + 1
			art += fileLines[txtLineNum+num]
		}
		art += ("\n")
		num++
	}
	return art
}
func newLine(str string) {
	splitstr := strings.Split(str, "\\n")
	for i := 0; i < len(splitstr); i++ {
		if splitstr[i] == "" {
			fmt.Print("\n")
		} else {
			fmt.Print(ascii(splitstr[i]))
		}
	}
}
