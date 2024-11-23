/*
File for expanded user interface actions like connecting to a board engine.
*/

package tui

func (m *Model) PlayRandomMove() {
	board := GameToBoard(&m.Game)
	board.MakeRandomMove()
	m.Game = *BoardToGame(board)
}
