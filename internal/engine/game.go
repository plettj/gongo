package engine

import "gongo/internal/utils"

type Game struct {
	ID    string `json:"id"`    // Gongo Snowflake ID
	Size  int    `json:"size"`  // Go board size
	Board int    `json:"board"` // Go board in binary
}

func NewGame() *Game {
	return &Game{
		ID:    utils.GetSnowflake(),
		Size:  19,
		Board: 0,
	}
}