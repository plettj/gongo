/*
File for expanded user interface actions like connecting to a board engine.
*/

package tui

import "gongo/internal/engine"

func (m *Model) PlayRandomMove() bool {
	played := m.Board.MakeRandomMove()

	return played
}

// Pre: 0 <= x < 19, 0 <= x < 19
func (m *Model) PlayMove(x int, y int) bool {
	legal := m.Board.MakeMove(engine.Loc{X: byte(x + 1), Y: byte(y + 1)})

	if legal && m.Settings["Bot"] == ON {
		m.PlayRandomMove()
	}

	return legal
}

func (m *Model) ClearBoard() {
	for i := range m.Board.Board {
		m.Board.Board[i] = 0
	}
	m.Board.Turn = 1
}
