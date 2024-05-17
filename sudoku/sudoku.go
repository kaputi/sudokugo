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
	placed   Board
	notes1   Board
	notes2   Board

	layer              int
	showCurrentErrors  bool
	showSolutionErrors bool
	theme              Theme
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

// TODO: esto no me gusta, puede ser con un iota como layer
func (s *Sudoku) SetErrorOption(option string) {
	switch option {
	case "none":
		s.showCurrentErrors = false
		s.showSolutionErrors = false
	case "current":
		s.showCurrentErrors = true
		s.showSolutionErrors = false
	case "solution":
		s.showCurrentErrors = false
		s.showSolutionErrors = true
	}
}

func (s *Sudoku) SetLayerOption(layer int) {
	s.layer = layer
}

func (s *Sudoku) SetBoard(board Board) error {
	solution, err := Solve(board)
	if err != nil {
		return err
	}

	s.puzzle = board
	s.placed = board
	s.solution = solution

	return nil
}

func (s *Sudoku) Generate(dificulty int) error {
	newBoard, solution := CreatePuzzle(dificulty)

	s.puzzle = newBoard
	s.placed = newBoard
	s.solution = solution

	return nil
}

func (s *Sudoku) SetCell(row, col int, value uint8) {
	if s.puzzle[row][col] != 0 {
		return
	}

	if s.layer != PLAY_LAYER {
		s.setNote(row, col, s.layer, value)
		return
	}

	s.placed[row][col] = value
}

func (s *Sudoku) setNote(row, col, noteLayer int, value uint8) {
	switch noteLayer {
	case NOTE1_LAYER:
		s.notes1[row][col] = value
	case NOTE2_LAYER:
		s.notes2[row][col] = value
	}
}

func (s *Sudoku) ClearCell(row, col int) {
	s.SetCell(row, col, 0)
}

func (s *Sudoku) getCellValue(row, col int) string {
	puzzleValue := s.puzzle[row][col]
	solutionValue := s.solution[row][col]
	placedValue := s.placed[row][col]

	if s.layer == PLAY_LAYER {
		if placedValue != 0 && s.showSolutionErrors && placedValue != solutionValue {
			return color(s.theme["error"], fmt.Sprintf("%d", placedValue))
		}
		if placedValue != 0 && s.showCurrentErrors && !isValidCell(placedValue, row, col, s.placed) {
			return color(s.theme["error"], fmt.Sprintf("%d", placedValue))
		}

		if puzzleValue != 0 {
			return color(s.theme["puzzle"], fmt.Sprintf("%d", puzzleValue))
		}
		if placedValue != 0 {
			return fmt.Sprintf("%d", placedValue)
		}
	} else {
		var noteBoard Board
		if s.layer == NOTE1_LAYER {
			noteBoard = s.notes1
		} else {
			noteBoard = s.notes2
		}

		if puzzleValue != 0 {
			return color(s.theme["puzzle"], fmt.Sprintf("%d", puzzleValue))
		}
		if placedValue != 0 {
			return fmt.Sprintf("%d", placedValue)
		}
		if noteBoard[row][col] != 0 {
			return color(s.theme["note"], fmt.Sprintf("%d", noteBoard[row][col]))
		}
	}

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

func (s *Sudoku) Display() {
	strings := s.GetBoardStrings()
	for _, row := range strings {
		fmt.Println(row)
	}
}
