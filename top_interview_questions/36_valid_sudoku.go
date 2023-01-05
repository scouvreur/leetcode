package main

import "testing"

var sampleInput = [][]byte{
	{0x35, 0x33, 0x2e, 0x2e, 0x37, 0x2e, 0x2e, 0x2e, 0x2e},
	{0x36, 0x2e, 0x2e, 0x31, 0x39, 0x35, 0x2e, 0x2e, 0x2e},
	{0x2e, 0x39, 0x38, 0x2e, 0x2e, 0x2e, 0x2e, 0x36, 0x2e},
	{0x38, 0x2e, 0x2e, 0x2e, 0x36, 0x2e, 0x2e, 0x2e, 0x33},
	{0x34, 0x2e, 0x2e, 0x38, 0x2e, 0x33, 0x2e, 0x2e, 0x31},
	{0x37, 0x2e, 0x2e, 0x2e, 0x32, 0x2e, 0x2e, 0x2e, 0x36},
	{0x2e, 0x36, 0x2e, 0x2e, 0x2e, 0x2e, 0x32, 0x38, 0x2e},
	{0x2e, 0x2e, 0x2e, 0x34, 0x31, 0x39, 0x2e, 0x2e, 0x35},
	{0x2e, 0x2e, 0x2e, 0x2e, 0x38, 0x2e, 0x2e, 0x37, 0x39},
}

var validBoard = [][]string{
	{"5", "3", ".", ".", "7", ".", ".", ".", "."},
	{"6", ".", ".", "1", "9", "5", ".", ".", "."},
	{".", "9", "8", ".", ".", ".", ".", "6", "."},
	{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
	{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
	{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
	{".", "6", ".", ".", ".", ".", "2", "8", "."},
	{".", ".", ".", "4", "1", "9", ".", ".", "5"},
	{".", ".", ".", ".", "8", ".", ".", "7", "9"},
}

var invalidBoard = [][]string{
	{"8", "3", ".", ".", "7", ".", ".", ".", "."},
	{"6", ".", ".", "1", "9", "5", ".", ".", "."},
	{".", "9", "8", ".", ".", ".", ".", "6", "."},
	{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
	{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
	{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
	{".", "6", ".", ".", ".", ".", "2", "8", "."},
	{".", ".", ".", "4", "1", "9", ".", ".", "5"},
	{".", ".", ".", ".", "8", ".", ".", "7", "9"},
}

var invalidBoard2 = [][]string{
	{".", ".", ".", ".", "5", ".", ".", "1", "."},
	{".", "4", ".", "3", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", "3", ".", ".", "1"},
	{"8", ".", ".", ".", ".", ".", ".", "2", "."},
	{".", ".", "2", ".", "7", ".", ".", ".", "."},
	{".", "1", "5", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", "2", ".", ".", "."},
	{".", "2", ".", "9", ".", ".", ".", ".", "."},
	{".", ".", "4", ".", ".", ".", ".", ".", "."},
}

func stringBoard(board [][]byte) [][]string {
	parsedBoard := [][]string{}

	for _, byteRow := range board {
		row := []string{}
		for _, cell := range byteRow {
			row = append(row, string(cell))
		}
		parsedBoard = append(parsedBoard, row)
	}

	return parsedBoard
}

// Remove empty cells from rows.
func removeEmpties(row []string) []string {
	nonEmptyCells := []string{}

	for _, cell := range row {
		if cell != "." {
			nonEmptyCells = append(nonEmptyCells, cell)
		}
	}

	return nonEmptyCells
}

// Check hoizontal rows.
func checkRows(board [][]string) bool {
	rowsValid := true

	for _, row := range board {
		var rowUniqueValues = make(map[string]bool)

		for index, value := range removeEmpties(row) {
			if index == 0 {
				rowUniqueValues[value] = true
			} else {
				// case when value already in the map
				if rowUniqueValues[value] {
					rowsValid = false
				}

				rowUniqueValues[value] = true
			}

			if !rowsValid {
				break
			}

			// fmt.Println(value, rowUniqueValues, rowsValid)
		}

		if !rowsValid {
			break
		}
	}

	return rowsValid
}

func checkColumns(board [][]string) bool {
	columns := [][]string{}

	for i := 0; i < 9; i++ {
		column := []string{}
		for _, row := range board {
			column = append(column, row[i])
		}
		columns = append(columns, column)
	}

	return checkRows(columns)
}

func checkGrids(board [][]string) bool {
	grids := [][]string{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grid := []string{}

			for _, row := range board[i*3 : (i+1)*3] {
				grid = append(grid, row[j*3:(j+1)*3]...)
			}
			grids = append(grids, grid)
		}
	}

	return checkRows(grids)
}

func isValidSudoku(board [][]byte) bool {
	parsedBoard := stringBoard(board)
	if !checkRows(parsedBoard) {
		return false
	} else if !checkColumns(parsedBoard) {
		return false
	} else if !checkGrids(parsedBoard) {
		return false
	} else {
		return true
	}
}

func main() {}

func TestIsValidSudoku(t *testing.T) {
	// fmt.Println(checkRows(validBoard))   // true
	// fmt.Println(checkRows(invalidBoard)) // true

	// fmt.Println(checkColumns(validBoard))   // true
	// fmt.Println(checkColumns(invalidBoard)) // false

	// fmt.Println(checkGrids(validBoard)) // true
	// fmt.Println(checkGrids(validBoard)) // true

	// fmt.Println(checkRows(invalidBoard2))    // true
	// fmt.Println(checkColumns(invalidBoard2)) // true
	// fmt.Println(checkGrids(invalidBoard2)) // false
}
