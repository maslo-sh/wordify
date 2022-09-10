package api

import (
	"wordify/internal/model"
)

type GameRequest struct {
	Players []*model.Player `json:"players"`
}
