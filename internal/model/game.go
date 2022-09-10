package model

import (
	mapset "github.com/deckarep/golang-set/v2"
	"math/rand"
	"sort"
	"wordify/internal/dictionaries"
	"wordify/internal/errors"
)

type Game struct {
	GameId     int
	Players    []*Player
	Board      Board
	LetterBag  Bag
	Dictionary *dictionaries.Dictionary
}

func NewGame() *Game {

	return &Game{
		GameId:    rand.Int(),
		Players:   make([]*Player, 2),
		Board:     NewBoard("../resources/boards/standard.txt"),
		LetterBag: NewBag(),
	}
}

func (g *Game) RegisterMove(playerId int, move Move) (string, error) {
	word, err := createWordFromMove(move)
	if err != nil {
		return "", err
	}
	points := getPointsForMove(move)

	_ = NewMoveRegister(move, points)

	return word, nil
}

func getPointsForMove(move Move) int {
	// TODO: counting points for move

	return 0
}

func createWordFromMove(move Move) (string, error) {
	err := checkIfPlacementIsCorrect(move)
	if err != nil {
		return "", err
	}
	sort.Sort(ByCoords(move.PlacedTiles))
	word := ""

	for _, element := range move.PlacedTiles {
		word += string(element.Tile)
	}

	return word, nil
}

func checkIfPlacementIsCorrect(move Move) error {
	occupiedColumnsSet := mapset.NewSet[int]()
	occupiedRowsSet := mapset.NewSet[int]()

	for _, place := range move.PlacedTiles {
		occupiedColumnsSet.Add(place.Coords.Column)
		occupiedRowsSet.Add(place.Coords.Row)
	}

	if occupiedColumnsSet.Cardinality() != 1 && occupiedRowsSet.Cardinality() != 1 {
		return errors.WrongPlacementError{Message: "Wrong placement tiles"}
	}

	return nil
}
