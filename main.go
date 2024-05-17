package main

import "github.com/kaputi/sudokugo/sudoku"

func main() {
	su := sudoku.New()
	err := su.Generate(3)

	if err != nil {
		panic(err)
	}

}
