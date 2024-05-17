package sudoku

import (
	"math/rand"
)

func removeCoords(board *Board, coord Coord) bool {

	valBackup := (*board)[coord.row][coord.col]
	(*board)[coord.row][coord.col] = 0

	_, err := Solve(*board)

	if err != nil {
		board[coord.row][coord.col] = valBackup
		return false
	}

	return true
}

func removeNumbers(board *Board, dificulty int, iterations *int) {
	dificulty = max(0, min(6, dificulty))

	filledSpaces := GetFilled(*board)

	filledLen := len(filledSpaces)

	clues := 30 - dificulty

	// pick a random filled space to remove
	coord := filledSpaces[rand.Intn(filledLen)]
	if removeCoords(board, coord) {
		filledLen--
	} else {
		*iterations++
		if *iterations > 20 {
			return
		}
		// fmt.Println("iterations: ", *iterations)
		removeNumbers(board, dificulty, iterations)
		return
	}

	*iterations = 0

	// TODO: final sudoku has 20 numbers, this needs to be adjusted somehow for dificulty
	if filledLen <= clues {
		return
	}

	diagonals := getDiagonals(coord, 1)
	scrambleSlice(&diagonals)

	amountOfDiagonalsToRemove := 4
	// after removing 20 numbers, only remove 2 of the diagonals
	if filledLen <= 60 {
		amountOfDiagonalsToRemove = 2
	}

	// remove empty diagonals
	finalDiagonals := []Coord{}
	for _, coord := range diagonals {
		if board[coord.row][coord.col] != 0 && len(finalDiagonals) < amountOfDiagonalsToRemove {
			finalDiagonals = append(finalDiagonals, coord)
		}
	}

	for _, coord := range finalDiagonals {
		if removeCoords(board, coord) {
			filledLen--
			if filledLen <= clues {
				return
			}
		}
	}

	removeNumbers(board, dificulty, iterations)
}

func getDiagonals(coord Coord, amount int) []Coord {
	diagonals := []Coord{}

	for i := 1; i <= amount; i++ {
		if coord.row+i < 9 && coord.col+i < 9 {
			diagonals = append(diagonals, NewCoord(coord.row+i, coord.col+i))
		}
		if coord.row+i < 9 && coord.col-i > 0 {
			diagonals = append(diagonals, NewCoord(coord.row+i, coord.col-i))
		}
		if coord.row-i > 0 && coord.col+i < 9 {
			diagonals = append(diagonals, NewCoord(coord.row-i, coord.col+i))
		}
		if coord.row-i > 0 && coord.col-i > 0 {
			diagonals = append(diagonals, NewCoord(coord.row-i, coord.col-i))
		}
	}

	return diagonals
}

func createSolvedBoard() Board {
	board := Board{}

	// fill 9 random spaces with numbers from 1 to 9
	for i := uint8(0); i < 9; i++ {
		row := rand.Intn(9)
		col := rand.Intn(9)

		for board[row][col] != 0 {
			row = rand.Intn(9)
			col = rand.Intn(9)
		}

		board[row][col] = i
	}

	// solve the board
	solved, _ := Solve(board)

	// the solver try to put smaller numbers first so boards are always low on first coords
	// so i mantain the board structure but swap numbers
	swapNumbersRandomly(&solved)

	return solved
}

func CreatePuzzle(dificulty int) (Board, Board) {
	solved := createSolvedBoard()
	puzzle := solved

	iteraitons := 0
	removeNumbers(&puzzle, dificulty, &iteraitons)

	return puzzle, solved
}
