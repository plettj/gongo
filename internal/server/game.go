/*
All server-side game play and validation logic.
*/
package server

import "gongo/internal/engine"

// Serializable game struct.
type Game struct {
	Id      string        `json:"id"`
	Board   *engine.Board `json:"board"`
	Players [2]string     `json:"players"`
	Moves   []Move        `json:"moves"`
}

// Serializable move struct.
type Move struct {
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Color string `json:"color"`
}
