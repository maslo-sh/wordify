package model

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type factor int8

const (
	TRIPLE   factor = 3
	DOUBLE   factor = 2
	STANDARD factor = 1
)

type multiplier struct {
	WordMultiplier   factor
	LetterMultiplier factor
}

type Board map[Coords]Field

type Field struct {
	Tile           letter
	MultiplierType multiplier
}

func NewBoard(filename string) Board {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	content, err := ioutil.ReadAll(io.Reader(file))

	rows := strings.Split(string(content), "\n")

	board := make(Board)

	for i, row := range rows {
		fields := strings.Split(row, " ")
		for j, field := range fields {
			fieldCoords := NewCoords(i, j)
			board[fieldCoords] = NewField(getMultiplierFromPattern(field))
		}
	}

	return board
}

func getMultiplierFromPattern(pattern string) multiplier {
	fact := factor(pattern[0] - 48)
	typeOfMultiplier := string(pattern[1])

	switch typeOfMultiplier {
	case "W":
		return multiplier{fact, 1}

	case "L":
		return multiplier{1, fact}
	default:
		return multiplier{1, 1}

	}
}

func NewField(multiplierType multiplier) Field {
	return Field{MultiplierType: multiplierType}
}
