package sudoku

import "fmt"

type Board [9][9]uint8

type Coord struct {
	row int
	col int
}

type Theme map[string]string

func NewCoord(row, col int) Coord {
	return Coord{row: row, col: col}
}

type Cell struct {
	coord Coord
	value uint8
}

func NewCell(row, col int, value uint8) Cell {
	return Cell{coord: NewCoord(row, col), value: value}
}

const (
	_ = iota
	PLAY_LAYER
	NOTE1_LAYER
	NOTE2_LAYER
)

type Sudoku struct {
	puzzle   Board
	solution Board

	layer int

	// showCurrentErrors  bool
	// showSolutionErrors bool

	theme Theme

	placed             []Cell
	notes1             []Cell
	notes2             []Cell
	errorsWithCurrent  []Cell
	errorsWithSolution []Cell
}

func New() *Sudoku {
	return &Sudoku{
		theme: map[string]string{
			"grid":      "blue",
			"note1Grid": "green",
			"note2Grid": "red",
			"puzzle":    "yellow",
			"placed":    "green",
			"error":     "red",
			"note":      "cyan",
		},
		layer: PLAY_LAYER,
	}
}

func (s *Sudoku) SetBoard(board Board) error {
	solution, err := Solve(board)
	if err != nil {
		return err
	}

	s.puzzle = board
	s.solution = solution

	return nil
}

func (s *Sudoku) Generate(dificulty int) error {
	newBoard, solution := CreatePuzzle(dificulty)

	s.puzzle = newBoard
	s.solution = solution

	return nil
}

func (s *Sudoku) SetCell(row, col int, value uint8) {
	if s.puzzle[row][col] != 0 {
		return
	}

	isErrorWithSoluton := s.solution[row][col] != value

	cell := NewCell(row, col, value)

	if isErrorWithSoluton {
		s.errorsWithSolution = append(s.errorsWithSolution, cell)
		return
	}

	// TODO:  iserrorwithcurrent

	s.placed = append(s.placed, cell)
}

func (s *Sudoku) SetNote(row, col, noteLayer int, value uint8) {
	if s.puzzle[row][col] != 0 {
		return
	}

	cell := NewCell(row, col, value)

	if noteLayer == 1 {
		s.notes1 = append(s.notes1, cell)
	} else {
		s.notes2 = append(s.notes2, cell)
	}
}

func (s *Sudoku) ClearCell(row, col int) {
	cell := NewCell(row, col, 0)

	for i, c := range s.placed {
		if c.coord == cell.coord {
			s.placed = append(s.placed[:i], s.placed[i+1:]...)
			return
		}
	}

	for i, c := range s.errorsWithCurrent {
		if c.coord == cell.coord {
			s.errorsWithCurrent = append(s.errorsWithCurrent[:i], s.errorsWithCurrent[i+1:]...)
			return
		}
	}

	for i, c := range s.errorsWithSolution {
		if c.coord == cell.coord {
			s.errorsWithSolution = append(s.errorsWithSolution[:i], s.errorsWithSolution[i+1:]...)
			return
		}
	}
}

func (s *Sudoku) getCellValue(row, col int) string {
	// return s.puzzle[row][col]
	if s.puzzle[row][col] != 0 {
		return color(s.theme["puzzle"], fmt.Sprintf("%d", s.puzzle[row][col]))
	}

	if s.layer == NOTE1_LAYER {
		for _, cell := range s.notes1 {
			if cell.coord.row == row && cell.coord.col == col {
				return color(s.theme["note"], fmt.Sprintf("%d", cell.value))
			}
		}
	}

	if s.layer == NOTE2_LAYER {
		for _, cell := range s.notes2 {
			if cell.coord.row == row && cell.coord.col == col {
				return color(s.theme["note"], fmt.Sprintf("%d", cell.value))
			}
		}
	}

	if s.layer == PLAY_LAYER {
		for _, cell := range s.placed {
			if cell.coord.row == row && cell.coord.col == col {
				return color(s.theme["placed"], fmt.Sprintf("%d", cell.value))
			}
		}
	}

	// TODO: erroes
	// currentErrorTheme := s.theme["placed"]
	// solutionErrorTheme := s.theme["placed"]
	// if s.showCurrentErrors {
	// 	currentErrorTheme = s.theme["error"]
	// }
	// if s.showSolutionErrors {
	// 	solutionErrorTheme = s.theme["error"]
	// }

	return " "
}

func (s *Sudoku) GetBoardStrings() []string {

	gridLayer := "grid"
	if s.layer == NOTE1_LAYER {
		gridLayer = "note1Grid"
	}
	if s.layer == NOTE2_LAYER {
		gridLayer = "note2Grid"
	}

	topRow := color(s.theme[gridLayer], "┌───┬───┬───┐ ┌───┬───┬───┐ ┌───┬───┬───┐")
	midRow := color(s.theme[gridLayer], "├───┼───┼───┤ ├───┼───┼───┤ ├───┼───┼───┤")
	botRow := color(s.theme[gridLayer], "└───┴───┴───┘ └───┴───┴───┘ └───┴───┴───┘")
	split := color(s.theme[gridLayer], "│")
	dSplit := color(s.theme[gridLayer], "│ │")

	boardString := []string{
		topRow,
	}

	for rowIdx, row := range s.puzzle {
		if rowIdx%3 != 0 {
			boardString = append(boardString, midRow)
		} else {
			if rowIdx != 0 {
				boardString = append(boardString, botRow)
				boardString = append(boardString, topRow)
			}
		}

		rowStr := ""
		for colIdx := range row {
			if colIdx == 0 {
				rowStr += split
			} else {
				if colIdx%3 == 0 {
					rowStr += dSplit
				} else {
					rowStr += split
				}
			}

			rowStr += fmt.Sprintf(" %v ", s.getCellValue(rowIdx, colIdx))
		}
		rowStr += split

		boardString = append(boardString, rowStr)
		if rowIdx == 8 {
			boardString = append(boardString, botRow)
		}
	}

	return boardString

}
