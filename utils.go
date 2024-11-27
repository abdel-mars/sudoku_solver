package piscine

const SIZE = 9

func Isvalid(board [SIZE][SIZE]int, row, col, num int) bool {
	// Check row and column
	for i := 0; i < SIZE; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}
	// Check 3x3 subgrid
	startRow, startCol := (row/3)*3, (col/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

// Validates that the initial Sudoku board respects Sudoku rules
func IsValidSudoku(board [SIZE][SIZE]int) bool {
	for row := 0; row < SIZE; row++ {
		for col := 0; col < SIZE; col++ {
			num := board[row][col]
			if num != 0 {
				board[row][col] = 0 // Temporarily clear the cell
				if !Isvalid(board, row, col, num) {
					return false
				}
				board[row][col] = num // Restore the cell
			}
		}
	}
	return true
}

// Parses input arguments into a Sudoku board
func ParseInput(args []string) ([SIZE][SIZE]int, bool) {
	var board [SIZE][SIZE]int
	if len(args) != SIZE {
		return board, false
	}
	for i, row := range args {
		if len(row) != SIZE {
			return board, false
		}
		for j, c := range row {
			if c == '.' || c == '0' {
				board[i][j] = 0
			} else if c >= '1' && c <= '9' {
				board[i][j] = int(c - '0')
			} else {
				return board, false
			}
		}
	}
	return board, true
}

func Solve(board *[SIZE][SIZE]int) bool {
	for row := 0; row < SIZE; row++ {
		for col := 0; col < SIZE; col++ {
			if board[row][col] == 0 { // Empty cell
				for num := 1; num <= SIZE; num++ { // Try numbers 1 through 9
					if Isvalid(*board, row, col, num) {
						board[row][col] = num
						if Solve(board) { // Recursive call
							return true
						}
						board[row][col] = 0 // Backtrack
					}
				}
				return false // No valid number found
			}
		}
	}
	return true // Puzzle solved
}
