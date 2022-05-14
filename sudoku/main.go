package main

import (
	"fmt"
	"os"
)

func main() {
	if CheckError() == true {
		fmt.Println("Error", "\n")
	} else {

		board := sodukuArray()

		if solveBoard(board) {
			printBoard(board)
		} else {
			fmt.Println("Error", "\n")
		}
	}
}

func CheckError() bool {
	args := os.Args

	argumentLen := len(args)

	if argumentLen != 10 {
		return true
	}

	for i := 1; i < len(args); i++ {
		if len(args[i]) != 9 {
			return true
		}
	}
	return false
}

func sodukuArray() [][]int {
	args := os.Args
	// // Error Messages
	// if len(args) < 10 {
	// 	fmt.Print("Error")
	// } else if len(args[1: ]) < 9 {
	// 	fmt.Print("Error")
	// } else if args[i: ] != string {
	// 	fmt.Print("Error")
	// }
	a := [][]int{}
	for i := 1; i <= 9; i++ {
		b := []int{}
		for j := 0; j < 9; j++ {
			r := []rune(args[i])

			if r[j] == '.' {
				b = append(b, 0)
			} else if r[j] >= '0' && r[j] <= '9' {
				b = append(b, int(r[j]-48))
			}

		}
		a = append(a, b)
	}
	return a
}

func printBoard(board [][]int) {
	for row := 0; row < 9; row++ {
		fmt.Print(board[row][0])
		for col := 1; col < 9; col++ {
			fmt.Printf("%2d", board[row][col])
		}
		fmt.Println()
	}
}

func isNumberInRow(board [][]int, number int, row int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == number {
			return true
		}
	}
	return false
}

func isNumberInCol(board [][]int, number int, col int) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == number {
			return true
		}
	}
	return false
}

func isNumberInBox(board [][]int, number int, row int, col int) bool {
	localBoxRow := row - row%3
	localBoxCol := col - col%3
	for i := localBoxRow; i < localBoxRow+3; i++ {
		for j := localBoxCol; j < localBoxCol+3; j++ {
			if board[i][j] == number {
				return true
			}
		}
	}
	return false
}

func isValidPlacement(board [][]int, number int, row int, col int) bool {
	if isNumberInRow(board, number, row) == false &&
		isNumberInCol(board, number, col) == false &&
		isNumberInBox(board, number, row, col) == false {
		return true
	}
	return false
}

func solveBoard(board [][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				for numToTry := 1; numToTry <= 9; numToTry++ {
					if isValidPlacement(board, numToTry, row, col) {
						board[row][col] = numToTry
						if solveBoard(board) {
							return true
						} else {
							board[row][col] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}
