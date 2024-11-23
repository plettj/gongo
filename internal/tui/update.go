package tui

import tea "github.com/charmbracelet/bubbletea"

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "w":
			if m.Cursor[1] > 0 {
				m.Cursor[1]--
			}
		case "down", "s":
			if m.Cursor[1] < 19-1 {
				m.Cursor[1]++
			}
		case "left", "a":
			if m.Cursor[0] > 0 {
				m.Cursor[0]--
			}
		case "right", "d":
			if m.Cursor[0] < 19-1 {
				m.Cursor[0]++
			}
		case "enter", " ":
			g := m.Game
			if g.Board[m.CursorCell()] == 0 {
				g.Moves = append(g.Moves, m.Cursor)
				g.Board[m.CursorCell()] = g.Turn
				g.Turn = 3 - g.Turn

				// TODO: Refactor this when I re-implement the Update function
				if m.Cursor[1] < 19-1 {
					m.Cursor[1]++
				} else if m.Cursor[1] > 0 {
					m.Cursor[1]--
				} else if m.Cursor[0] < 19-1 {
					m.Cursor[0]++
				} else if m.Cursor[0] > 0 {
					m.Cursor[0]--
				}
			}
		}
	}

	return m, nil
}
