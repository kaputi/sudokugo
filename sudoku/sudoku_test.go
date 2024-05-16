package sudoku

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSudoku(t *testing.T) {
	board := Board{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		// -------------------
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		// -------------------
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	// assert.Equal(t, true, Validate(board))

	expected := Board{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		// -------------------
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		// -------------------
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	solved, err := Solve(board)

	assert.Equal(t, expected, solved)
	assert.Nil(t, err)

	worstCase := Board{
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		// -------------------
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		// -------------------
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	_, err = Solve(worstCase)
	assert.EqualError(t, err, "invalid, multiple solutions found")

	unsolvable := Board{
		{2, 0, 0, 9, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 6, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0},
		// -------------------
		{5, 0, 2, 6, 0, 0, 4, 0, 7},
		{0, 0, 0, 0, 0, 4, 1, 0, 0},
		{0, 0, 0, 0, 9, 8, 0, 2, 3},
		// -------------------
		{0, 0, 0, 0, 0, 3, 0, 8, 0},
		{0, 0, 5, 0, 1, 0, 0, 0, 0},
		{0, 0, 7, 0, 0, 0, 0, 0, 0},
	}

	// assert.Equal(t, false, Validate(unsolvable))

	_, err = Solve(unsolvable)
	assert.EqualError(t, err, "unsolvable")

	newBoard, newSolutionExpects := CreatePuzzle(6)
	clues := len(GetFilled(newBoard))
	fmt.Println("clues: ", clues)
	PrintBoard(newBoard)

	newBoardSolution, err := Solve(newBoard)
	assert.Equal(t, newSolutionExpects, newBoardSolution)
	assert.Nil(t, err)

	twoSolutions := Board{
		{2, 9, 5, 7, 4, 3, 8, 6, 1},
		{4, 3, 1, 8, 6, 5, 9, 0, 0},
		{8, 7, 6, 1, 9, 2, 5, 4, 3},
		{3, 8, 7, 4, 5, 9, 2, 1, 6},
		{6, 1, 2, 3, 8, 7, 4, 9, 5},
		{5, 4, 9, 2, 1, 6, 7, 3, 8},
		{7, 6, 3, 5, 2, 4, 1, 8, 9},
		{9, 2, 8, 6, 7, 1, 3, 5, 4},
		{1, 5, 4, 9, 3, 8, 6, 0, 0},
	}

	count := CountSolutions(twoSolutions)
	assert.Equal(t, 2, count)

	sudoku := New()

	err = sudoku.Generate(4)
	fmt.Println(err)

	strings := sudoku.GetBoardStrings()

  for _, row := range strings{
    fmt.Println(row)
  }
}
