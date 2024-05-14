package sudoku

type Table [9][9]uint8

type Sudoku struct {
	table Table
}

func New() *Sudoku {
	return &Sudoku{}
}

func (s *Sudoku) SetBoard(table Table) error {

  s.table = table

	return nil
}
