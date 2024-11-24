package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// TODO: To decrease the total number of board representations in this repo,
//       Redesign this to be also the serializable representation of the board.
type Game struct {
	Size  int
	Board []byte    // 0 is empty, 1 is black, 2 is white
	Moves [][2]byte // List of moves made
	Turn  byte
}

type Model struct {
	Game       Game
	Cursor     [2]int // Which cell is currently selected
	MouseEvent tea.MouseEvent
	Offsets    [2]int // Offset of board from top-left of view
}

func (m *Model) CursorCell() int {
	return m.Cursor[0] + m.Cursor[1]*m.Game.Size
}

func NewModel(size int) *Model {
	model := Model{
		Game: Game{
			Size:  size,
			Board: make([]byte, size*size),
			Moves: [][2]byte{},
			Turn:  1,
		},
		Cursor:  [2]int{size / 2, size / 2},
		Offsets: [2]int{4, 7},
	}

	if size == 13 {
		model.Game.Board = GONGO_13[:]
	}

	return &model
}
