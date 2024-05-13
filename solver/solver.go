package solver

import (
	"errors"
)

type EmptySpaces struct {
	row int
	col int
}

func solveRecurse(emptySpaceIdx int, emptySpaces *[]EmptySpaces, problem *[][]uint8) bool {
	if emptySpaceIdx == len((*emptySpaces)) {
		return true
	}

	emptySpace := (*emptySpaces)[emptySpaceIdx]

	for number := uint8(1); number <= 9; number++ {
		if isValid(number, emptySpace.row, emptySpace.col, problem) {
			(*problem)[emptySpace.row][emptySpace.col] = number

			if solveRecurse(emptySpaceIdx+1, emptySpaces, problem) {
				return true
			}

			(*problem)[emptySpace.row][emptySpace.col] = 0
		}
	}

	return false
}

func Solve(problem [][]uint8) ([][]uint8, error) {

	emptySpaces := []EmptySpaces{}

	for rowIdx, row := range problem {
		for colIdx, col := range row {
			if col == 0 {
				emptySpaces = append(emptySpaces, EmptySpaces{row: rowIdx, col: colIdx})
			}
		}
	}

	if solveRecurse(0, &emptySpaces, &problem) {
		return problem, nil
	}


	return problem, errors.New("no solution found")
}

func isValid(number uint8, row, col int, board *[][]uint8) bool {
	// chcek that number doesnt repeat in row and column
	for i := 0; i < 9; i++ {
		if (*board)[row][i] == number || (*board)[i][col] == number {
			return false
		}
	}

	// check 3x3 matrix
	// we need to find the first position of the 3x3 matrix
	firstRow := row - row%3
	firstCol := col - col%3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (*board)[firstRow+i][firstCol+j] == number {
				return false
			}
		}
	}

	return true
}
