package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatingTripleWordMultiplyingFieldOnStandardBoard(t *testing.T) {
	assert := assert.New(t)

	board := NewBoard("../resources/boards/standard.txt")

	assert.Equalf(multiplier{TRIPLE, STANDARD}, board[Coords{0, 0}].MultiplierType, "")
	assert.Equalf(multiplier{TRIPLE, STANDARD}, board[Coords{14, 14}].MultiplierType, "")
}

func TestCreatingDoubleWordMultiplyingFieldOnStandardBoard(t *testing.T) {
	assert := assert.New(t)

	board := NewBoard("../resources/boards/standard.txt")

	assert.Equalf(multiplier{DOUBLE, STANDARD}, board[Coords{1, 1}].MultiplierType, "")
	assert.Equalf(multiplier{DOUBLE, STANDARD}, board[Coords{13, 13}].MultiplierType, "")
	assert.Equalf(multiplier{DOUBLE, STANDARD}, board[Coords{7, 7}].MultiplierType, "")
}
