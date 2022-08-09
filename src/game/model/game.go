package model

type Game struct {
	FirstPlayer  *Player
	SecondPlayer *Player
	Board        Board
	LetterBag    *Bag
}

func NewGame(firstPlayer, secondPlayer *Player) *Game {

	return &Game{
		FirstPlayer:  firstPlayer,
		SecondPlayer: secondPlayer,
		Board:        NewBoard(),
		LetterBag:    NewBag(),
	}
}
