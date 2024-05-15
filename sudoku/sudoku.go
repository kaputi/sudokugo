package sudoku

import (
	"errors"
	"fmt"
	"math/rand"
)

type Board [9][9]uint8

type Coord struct {
	row int
	col int
}

func NewCoord(row, col int) Coord {
	return Coord{row: row, col: col}
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

func getCoords(board Board, filled bool) []Coord {
	spaces := []Coord{}

	for rowIdx, row := range board {
		for colIdx, cell := range row {
			if filled && cell != 0 {
				spaces = append(spaces, NewCoord(rowIdx, colIdx))
			}
			if !filled && cell == 0 {
				spaces = append(spaces, NewCoord(rowIdx, colIdx))
			}
		}
	}

	return spaces
}

func GetEmpty(board Board) []Coord {
	emptySpaces := getCoords(board, false)
	return emptySpaces
}

func GetFilled(board Board) []Coord {
	filledSpaces := getCoords(board, true)
	return filledSpaces
}

func Validate(board Board) bool {
	// TODO: check for only one posible solution
	_, err := Solve(board)
	return err == nil
}

func solveRecurse(emtySpaceIdx int, emptySpaces *[]Coord, board *Board) bool {
	if emtySpaceIdx == len((*emptySpaces)) {
		return true
	}

	emptySpace := (*emptySpaces)[emtySpaceIdx]

	for number := uint8(1); number <= 9; number++ {
		// for _, number := range numbers {
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

// TODO: make to work with generics
// func scrambleSlice(slice *[]interface{}) {
// 	for i := range *slice {
// 		j := rand.Intn(i + 1)
// 		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
// 	}
// }

func SwapNumbersRandomly(board *Board) {
	numbers := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// scramble numbers
	// TODO: use func
	for i := range numbers {
		j := rand.Intn(i + 1)
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	for rowIdx, row := range *board {
		for colIdx, cell := range row {
			(*board)[rowIdx][colIdx] = numbers[cell-1]
		}
	}

}

func getNextEmpty(board Board) (int, int) {
	for rowIdx, row := range board {
		for colIdx, cell := range row {
			if cell == 0 {
				return rowIdx, colIdx
			}
		}
	}

	return -1, -1
}

func countSolutionsHelper(board *Board, counter *int) bool {

	row, col := getNextEmpty(*board)

	for value := uint8(1); value <= 9; value++ {
		if !isValidCell(value, row, col, *board) {
			continue
		}

		(*board)[row][col] = value

		if row == 8 && col == 8 {
			*counter++
			break
		}

		if checkFullBoard(*board) {
			*counter++
			break
		} else if countSolutionsHelper(board, counter) {
			return true
		}

	}

	(*board)[row][col] = 0
	return false
}

func CountSolutions(board Board) int {
	counter := 0

	countSolutionsHelper(&board, &counter)

	return counter
}

func removeCoords(board *Board, coord Coord) bool {

	val := (*board)[coord.row][coord.col]
	(*board)[coord.row][coord.col] = 0

	solutions := CountSolutions(*board)

	if solutions != 1 {
		board[coord.row][coord.col] = val
		return false
	}

	return true
}

func removeNumbers(board *Board, dificulty int) {

	dificulty = max(0, min(6, dificulty))

	filledSpaces := GetFilled(*board)

	filledLen := len(filledSpaces)

	clues := 30 - dificulty

	// pick a random filled space to remove
	coord := filledSpaces[rand.Intn(filledLen)]
	if removeCoords(board, coord) {
		filledLen--
	} else {
		removeNumbers(board, dificulty)
		return
	}

	// TODO: final sudoku has 20 numbers, this needs to be adjusted somehow for dificulty
	if filledLen <= clues {
		return
	}

	diagonals := getDiagonals(coord, 1)
	//scramble diagonals
	// TODO: use fn
	for i := range diagonals {
		j := rand.Intn(i + 1)
		diagonals[i], diagonals[j] = diagonals[j], diagonals[i]
	}

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

	removeNumbers(board, dificulty)
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

func CreatePuzzle(dificulty int) Board {
	puzzle := CreateSolvedBoard()

	removeNumbers(&puzzle, dificulty)

	return puzzle
}

func CreateSolvedBoard() Board {
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
	SwapNumbersRandomly(&solved)

	return solved
}

func PrintBoard(board Board) {

	fmt.Println(" -------------------------------------------------")
	fmt.Println("   0 1 2 | 3 4 5 | 6 7 8")
	fmt.Println("   =====================")
	for rowIdx, row := range board {
		if rowIdx%3 == 0 && rowIdx != 0 {
			fmt.Println(" | ---------------------")
		}
		for colIdx, cell := range row {
			if colIdx == 0 {
				fmt.Printf("%v| ", rowIdx)
			}
			fmt.Printf("%v ", cell)
			if (colIdx+1)%3 == 0 && colIdx != len(row)-1 {
				fmt.Print("| ")
			}
		}
		fmt.Print("\n")
	}

	fmt.Print("\n")
}

func checkFullBoard(board Board) bool {
	for _, rowIdx := range board {
		for _, cell := range rowIdx {
			if cell == 0 {
				return false
			}
		}
	}

	return true
}

/*
binary map
0 = 0000
1 = 0001
2 = 0010
3 = 0011
4 = 0100
5 = 0101
6 = 0110
7 = 0111
8 = 1000
9 = 1001
*/
