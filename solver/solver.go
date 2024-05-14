package solver

import (
	"errors"

	"github.com/kaputi/sudokugo/sudoku"
)

type EmptySpaces struct {
	row int
	col int
}

func Solve(table sudoku.Table) (sudoku.Table, error) {
	emptySpaces := []EmptySpaces{}

	// get all empty spaces coords
	for rowIdx, row := range table {
		for colIdx, cell := range row {
			if cell == 0 {
				emptySpaces = append(emptySpaces, EmptySpaces{row: rowIdx, col: colIdx})
			}
		}
	}

	if solveRecurse(0, &emptySpaces, &table) {
		return table, nil
	}

	return table, errors.New("unsolvable")
}

func Validate(table sudoku.Table) bool {
	_, err := Solve(table)
	return err == nil
}

func solveRecurse(emtySpaceIdx int, emptySpaces *[]EmptySpaces, table *sudoku.Table) bool {
	if emtySpaceIdx == len((*emptySpaces)) {
		return true
	}

	emptySpace := (*emptySpaces)[emtySpaceIdx]

	for number := uint8(1); number <= 9; number++ {
		if isValidCell(number, emptySpace.row, emptySpace.col, *(table)) {
			(*table)[emptySpace.row][emptySpace.col] = number

			if solveRecurse(emtySpaceIdx+1, emptySpaces, table) {
				return true
			}

			(*table)[emptySpace.row][emptySpace.col] = 0
		}
	}

	return false
}

func isValidCell(number uint8, row int, col int, table sudoku.Table) bool {
	// check that number doesnt repeat in row or column
	for i := 0; i < 9; i++ {
		if table[row][i] == number || table[i][col] == number {
			return false
		}
	}

	// check that number doesnt repeat in the 3x3 grid
	firstRowIdxInGrid := row - row%3
	firstColIdxInGrid := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if table[firstRowIdxInGrid+i][firstColIdxInGrid+j] == number {
				return false
			}
		}
	}

	return true
}
