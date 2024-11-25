package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// TODO: To decrease the total number of board representations in this repo,
//       Redesign this to be also the serializable representation of the board.
type Game struct {
	Size   int
	Board  []byte // 0 is empty, 1 is black, 2 is white
	Turn   byte
	Ko     [2]byte   // (-1, -1) is no ko.
	Marked [][2]byte // List of marked cells
}

type Model struct {
	Game       Game
	Cursor     [2]int // Which cell is currently selected
	MouseEvent tea.MouseEvent
	Offsets    [2]int // Offset of board from top-left of view
	ActiveTab  string
}

var TABS = [2]string{"Game", "Settings"}

func (m *Model) CursorCell() int {
	return m.Cursor[0] + m.Cursor[1]*m.Game.Size
}

func NewModel(size int) *Model {
	model := Model{
		Game: Game{
			Size:   size,
			Board:  make([]byte, size*size),
			Marked: [][2]byte{},
			Turn:   1,
		},
		Cursor:    [2]int{size / 2, size / 2},
		Offsets:   [2]int{4, 7},
		ActiveTab: TABS[0],
	}

	if size == 19 {
		model.Game.Board = GONGO_19[:]
	}

	return &model
}
