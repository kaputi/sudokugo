package state

import "github.com/kaputi/sudokugo/sudoku"

type State struct {
	puzzle    sudoku.Board
	soluion   sudoku.Board

	correct   []sudoku.Coord
	incorrect []sudoku.Coord


}
