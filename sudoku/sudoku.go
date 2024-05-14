package sudoku

import "errors"

type Board [9][9]uint8

type EmptySpaces struct {
	row int
	col int
}

type Sudoku struct {
	board Board
}

func New() *Sudoku {
	return &Sudoku{}
}

func (s *Sudoku) SetBoard(board Board) error {

	if !Validate(board) {
		return errors.New("invalid")
	}

	s.board = board

	return nil
}

func Solve(board Board) (Board, error) {

	emptySpaces := GetEmpty(board)

	if solveRecurse(0, &emptySpaces, &board) {
		return board, nil
	}

	return board, errors.New("unsolvable")
}

func GetEmpty(board Board) []EmptySpaces {
	emptySpaces := []EmptySpaces{}

	for rowIdx, row := range board {
		for colIdx, cell := range row {
			if cell == 0 {
				emptySpaces = append(emptySpaces, EmptySpaces{row: rowIdx, col: colIdx})
			}
		}
	}

	return emptySpaces
}

func Validate(board Board) bool {
	_, err := Solve(board)
	return err == nil
}

func solveRecurse(emtySpaceIdx int, emptySpaces *[]EmptySpaces, board *Board) bool {
	if emtySpaceIdx == len((*emptySpaces)) {
		return true
	}

	emptySpace := (*emptySpaces)[emtySpaceIdx]

	for number := uint8(1); number <= 9; number++ {
		if isValidCell(number, emptySpace.row, emptySpace.col, *(board)) {
			(*board)[emptySpace.row][emptySpace.col] = number

			if solveRecurse(emtySpaceIdx+1, emptySpaces, board) {
				return true
			}

			(*board)[emptySpace.row][emptySpace.col] = 0
		}
	}

	return false
}

func isValidCell(number uint8, row int, col int, board Board) bool {
	// check that number doesnt repeat in row or column
	for i := 0; i < 9; i++ {
		if board[row][i] == number || board[i][col] == number {
			return false
		}
	}

	// check that number doesnt repeat in the 3x3 grid
	firstRowIdxInGrid := row - row%3
	firstColIdxInGrid := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[firstRowIdxInGrid+i][firstColIdxInGrid+j] == number {
				return false
			}
		}
	}

	return true
}
