package sudoku

import (
	"fmt"
	"math/rand"
)

func swapNumbersRandomly(board *Board) {
	numbers := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	scrambleSlice(&numbers)

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

func scrambleSlice[T any](slice *[]T) {
	for i := range *slice {
		j := rand.Intn(i + 1)
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	}
}

var colorMap = map[string]string{
	"reset":   "\033[0m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"gray":    "\033[37m",
	"white":   "\033[97m",
}

var inertedColorMal = map[string]string{
	"red":     "\033[30m\033[41m",
	"green":   "\033[30m\033[42m",
	"yellow":  "\033[30m\033[43m",
	"blue":    "\033[30m\033[44m",
	"magenta": "\033[30m\033[45m",
	"cyan":    "\033[30m\033[46m",
	"gray":    "\033[30m\033[47m",
	"white":   "\033[90m\033[107m",
}

// Default	\033[39m	\033[49m
// Black	\033[30m	\033[40m
// Dark red	\033[31m	\033[41m
// Dark green	\033[32m	\033[42m
// Dark yellow (Orange-ish)	\033[33m	\033[43m
// Dark blue	\033[34m	\033[44m
// Dark magenta	\033[35m	\033[45m
// Dark cyan	\033[36m	\033[46m
// Light gray	\033[37m	\033[47m
// Dark gray	\033[90m	\033[100m
// Red	\033[91m	\033[101m
// Green	\033[92m	\033[101m
// Orange	\033[93m	\033[103m
// Blue	\033[94m	\033[104m
// Magenta	\033[95m	\033[105m
// Cyan	\033[96m	\033[106m
// White	\033[97m	\033[107m

func color(color, str string) string {
	if _, ok := colorMap[color]; !ok {
		return str
	}
	return colorMap[color] + str + colorMap["reset"]
}

func PrintBoard(board Board) {

	fmt.Print("\n")

	var split = "|"
	var cSplit = color("blue", "|")
	var doubleSplit = color("blue", "||")
	var barSplit = color("blue", "++")
	var barEdge = color("blue", "+")

	for rowIdx, row := range board {
		if rowIdx == 0 {
			fmt.Println(color("blue", "     +---+---+---++---+---+---++---+---+---+"))
		}

		if rowIdx%3 != 0 {
			fmt.Printf("     %v", barEdge)
			fmt.Printf(" - + - + - %v - + - + - %v - + - + - ", barSplit, barSplit)
			fmt.Println(barEdge)
		}

		if rowIdx%3 == 0 && rowIdx != 0 {
			fmt.Println(color("blue", "     +-=-+-=-+-=-++-=-+-=-+-=-++-=-+-=-+-=-+"))
		}

		for colIdx, cell := range row {
			if colIdx == 0 {
				fmt.Print("     ")
				fmt.Print(cSplit)
			}

			if colIdx%3 == 0 && colIdx != 0 {
				fmt.Printf(" %v", doubleSplit)
			}

			if cell == 0 {
				fmt.Print("  ")
			} else {
				a := fmt.Sprintf(" %v", cell)
				fmt.Print(color("yellow", a))
			}

			if colIdx != 8 && colIdx%3 != 2 {
				fmt.Printf(" %v", split)
			}

			if colIdx == 8 {
				fmt.Printf(" %v", cSplit)
			}
		}
		fmt.Printf("\n")

		if rowIdx == 8 {
			fmt.Println(color("blue", "     +---+---+---++---+---+---++---+---+---+"))
		}
	}

	fmt.Print("\n")
}
