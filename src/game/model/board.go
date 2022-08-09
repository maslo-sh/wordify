package model

type MultiplyType string

const (
	WORD_MULTIPLY   MultiplyType = "WORD"
	LETTER_MULTIPLY              = "LETTER"
)

type Board [][]Field

type Field struct {
	Row          rune
	Col          rune
	MultiplyType MultiplyType
	Tile         *Tile
}

func NewBoard() Board {
	board := make([][]Field, 15)
	for i := 0; i < 15; i++ {
		board[i] = make([]Field, 15)
	}
	return board
}
