package sudoku

import (
	"errors"
)

func isValidCell(number uint8, row int, col int, board Board) bool {
	// check that number doesnt repeat in row or column
	for i := 0; i < 9; i++ {
		if board[row][i] == number && i != col {
			return false
		}

		if board[i][col] == number && i != row {
			return false
		}
	}

	// check that number doesnt repeat in the 3x3 grid
	firstRowIdxInGrid := row - row%3
	firstColIdxInGrid := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			checRow := firstRowIdxInGrid + i
			checkCol := firstColIdxInGrid + j
			if checRow == row && checkCol == col {
				continue
			}
			if board[checRow][checkCol] == number {
				return false
			}

		}
	}

	return true
}

func CountSolutions(board Board) int {
	solved := board
	counter := 0

	solveHelper(&board, &solved, &counter, true)

	return counter
}

func Solve(board Board) (Board, error) {
	solved := board
	counter := 0

	solveHelper(&board, &solved, &counter, false)

	if counter == 0 {
		return board, errors.New("unsolvable")
	}

	if counter > 1 {
		return solved, errors.New("invalid, multiple solutions found")
	}

	return solved, nil
}

func solveHelper(board *Board, solved *Board, counter *int, countAll bool) bool {
	row, col := getNextEmpty(*board)

	for value := uint8(1); value <= 9; value++ {
		if !isValidCell(value, row, col, *board) {
			continue
		}

		(*board)[row][col] = value

		// if board is full
		if boardIsFull(*board) {
			*counter++
			if *counter > 1 && !countAll {
				return false
			}
			*solved = *board
			break
		} else if solveHelper(board, solved, counter, countAll) {
			return true
		}

	}

	(*board)[row][col] = 0
	return false
}

func boardIsFull(board Board) bool {
	for _, rowIdx := range board {
		for _, cell := range rowIdx {
			if cell == 0 {
				return false
			}
		}
	}

	return true
}
