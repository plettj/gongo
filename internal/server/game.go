/*
All server-side game play and validation logic.
*/
package server

// Serializable game struct for live games.
type Game struct {
	Id      string    `json:"id"`      // Game Snowflake ID
	Board   string    `json:"board"`   // Game board in serialized engine format
	Players [2]string `json:"players"` // [Black, White] where values are usernames
	Moves   []Move    `json:"moves"`   // Game history
}

// Serializable move struct.
type Move struct {
	X     int  `json:"x"`
	Y     int  `json:"y"`
	Color byte `json:"color"` // 0: Pass, 1: Black, 2: White
}
