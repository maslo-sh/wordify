package model

import mapset "github.com/deckarep/golang-set/v2"

type Player struct {
	Name  string
	Score int
	Rack  Rack
}

type Rack mapset.Set[Tile]
