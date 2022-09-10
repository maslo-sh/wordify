package model

type MoveRegister struct {
	Move   Move
	Points int
}

func NewMoveRegister(move Move, points int) *MoveRegister {
	return &MoveRegister{move, points}
}
