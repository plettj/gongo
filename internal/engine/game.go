package engine

import "gongo/internal/utils"

// GameSer is a serializable representation of a Go game.
type GameSer struct {
	ID    string    `json:"id"`    // Gongo Snowflake ID
	Size  uint8     `json:"size"`  // Go board size
	Board [6]uint64 `json:"board"` // Go board in binary
}

func NewGame() *GameSer {
	return &GameSer{
		ID:    utils.GetSnowflake(),
		Size:  19,
		Board: [6]uint64{0, 0, 0, 0, 0, 0},
	}
}
