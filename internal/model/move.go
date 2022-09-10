package model

type ByCoords []placement

type Move struct {
	PlayerId    int
	PlacedTiles []placement
}

type placement struct {
	Coords Coords
	Tile   letter
}

func (a ByCoords) Less(i, j int) bool {
	if a[i].Coords.Row == a[j].Coords.Row {
		return a[i].Coords.Column < a[j].Coords.Column
	}

	if a[i].Coords.Column == a[j].Coords.Column {
		return a[i].Coords.Row < a[j].Coords.Row
	}

	return false
}

func (a ByCoords) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByCoords) Len() int {
	return len(a)
}
