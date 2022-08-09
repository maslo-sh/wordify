package model

import mapset "github.com/deckarep/golang-set/v2"

type Bag struct {
	Tiles mapset.Set[Tile]
}

func NewBag() *Bag {
	return &Bag{
		Tiles: createSetOfLetters(),
	}
}

func createSetOfLetters() mapset.Set[Tile] {
	tempSet := mapset.NewSet[Tile]()
	for i := 0; i < 26; i++ {
		tempSet.Add(Tile{rune(i + 97), 1})
	}
	return tempSet
}
