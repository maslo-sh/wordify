package model

type Player struct {
	User  *User
	Score int
	Rack  Rack
}

type Rack Bag
