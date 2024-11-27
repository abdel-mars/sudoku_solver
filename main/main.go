package main

import (
	"fmt"
	"os"
	"piscine"
)

const SIZE = 9

func main() {
	args := os.Args[1:]
	if len(args) != SIZE {
		fmt.Println("Error")
		return
	}
	board, valid := piscine.ParseInput(args)
	if !valid || !piscine.IsValidSudoku(board) {
		fmt.Println("Error")
		return
	}
	if !piscine.Solve(&board) {
		fmt.Println("Error")
		return
	}
	// Print solved Sudoku board
	for _, row := range board {
		for _, num := range row {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}

// Solves the Sudoku puzzle using backtracking
// func Solve(board *[SIZE][SIZE]int) bool {
// 	for row := 0; row < SIZE; row++ {
// 		for col := 0; col < SIZE; col++ {
// 			if board[row][col] == 0 { // Empty cell
// 				for num := 1; num <= SIZE; num++ { // Try numbers 1 through 9
// 					if piscine.Isvalid(*board, row, col, num) {
// 						board[row][col] = num
// 						if Solve(board) { // Recursive call
// 							return true
// 						}
// 						board[row][col] = 0 // Backtrack
// 					}
// 				}
// 				return false // No valid number found
// 			}
// 		}
// 	}
// 	return true // Puzzle solved
// }
