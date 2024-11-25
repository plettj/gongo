package tui

import (
	"gongo/internal/engine"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Board      engine.Board
	Cursor     [2]int // Which cell is currently selected
	MouseEvent tea.MouseEvent
	Offsets    [2]int // Offset of board from top-left of view
	ActiveTab  string
	Settings   map[string]int
}

const (
	ON  = 1
	OFF = 0
)

var TABS = [2]string{"Game", "Settings"}
var SETTINGS = [2]string{"Bot", "Board Size"}

func (m *Model) CursorCell() int {
	return m.Cursor[0] + m.Cursor[1]*int(m.Board.Size)
}

func NewModel(size int) *Model {
	model := Model{
		Board:     *engine.NewBoard(byte(size)),
		Cursor:    [2]int{size / 2, size / 2},
		Offsets:   [2]int{4, 7},
		ActiveTab: TABS[0],
		Settings:  map[string]int{"Bot": OFF, "Board Size": size},
	}

	return &model
}
