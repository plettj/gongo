/*
File for expanded user interface actions like connecting to a board engine.
*/

package tui

import "gongo/internal/engine"

func (m *Model) PlayRandomMove() bool {
	board := GameToBoard(&m.Game)
	played := board.MakeRandomMove()

	if played {
		m.Game = *BoardToGame(board)
	}

	return played
}

// Pre: 0 <= x < 19, 0 <= x < 19
func (m *Model) PlayMove(x int, y int) bool {
	board := GameToBoard(&m.Game)
	legal := board.MakeMove(engine.Loc{X: byte(x + 1), Y: byte(y + 1)})

	if legal {
		m.Game = *BoardToGame(board)

		// Bot plays move
		m.PlayRandomMove()
	}

	return legal
}

func (m *Model) ClearBoard() {
	for i := range m.Game.Board {
		m.Game.Board[i] = 0
	}
	m.Game.Turn = 1
}
