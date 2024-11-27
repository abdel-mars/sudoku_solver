package piscine

import (
	"piscine"
	"testing"
)

func TestValidSudokuInput(t *testing.T) {
	args := []string{
		"53..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	}

	board, valid := piscine.ParseInput(args)
	if !valid {
		t.Errorf("Expected input to be valid, but got invalid")
	}

	if !piscine.Solve(&board) {
		t.Errorf("Expected the Sudoku to be solvable")
	}
}

func TestInvalidSudokuInputLength(t *testing.T) {
	args := []string{
		"53..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
	}

	_, valid := piscine.ParseInput(args)
	if valid {
		t.Errorf("Expected input with fewer than 9 rows to be invalid")
	}
}

func TestInvalidSudokuRowLength(t *testing.T) {
	args := []string{
		"53..7...",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	}

	_, valid := piscine.ParseInput(args)
	if valid {
		t.Errorf("Expected input with a row shorter than 9 characters to be invalid")
	}
}

func TestUnsolvableSudoku(t *testing.T) {
	args := []string{
		"533..7...",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	}

	board, valid := piscine.ParseInput(args)
	if !valid {
		t.Errorf("Expected input to be valid, but got invalid")
	}

	if piscine.Solve(&board) {
		t.Errorf("Expected the Sudoku to be unsolvable")
	}
}

func TestEmptySudoku(t *testing.T) {
	args := []string{
		".........",
		".........",
		".........",
		".........",
		".........",
		".........",
		".........",
		".........",
		".........",
	}

	board, valid := piscine.ParseInput(args)
	if !valid {
		t.Errorf("Expected empty Sudoku to be valid")
	}

	if !piscine.Solve(&board) {
		t.Errorf("Expected empty Sudoku to be solvable")
	}
}
