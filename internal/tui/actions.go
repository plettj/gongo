/*
File for expanded user interface actions like connecting to a board engine.
*/

package tui

import "gongo/internal/engine"

func (m *Model) PlayRandomMove() {
	board := GameToBoard(&m.Game)
	board.MakeRandomMove()
	m.Game = *BoardToGame(board)
}

// Pre: 0 <= x < 19, 0 <= x < 19
func (m *Model) PlayMove(x byte, y byte) bool {
	board := GameToBoard(&m.Game)
	legal := board.MakeMove(engine.Loc{X: x + 1, Y: y + 1})

	if legal {
		m.Game = *BoardToGame(board)
	}

	return legal
}

func (m *Model) ClearBoard() {
	for i := range m.Game.Board {
		m.Game.Board[i] = 0
	}
	m.Game.Turn = 1
}
