package model

type Coords struct {
	Row    int
	Column int
}

func NewCoords(row, col int) Coords {
	return Coords{row, col}
}
