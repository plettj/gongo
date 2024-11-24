package tui

import tea "github.com/charmbracelet/bubbletea"

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
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
		case "r":
			m.ClearBoard()
		case "enter", " ":
			played := m.PlayMove(m.Cursor[0], m.Cursor[1])

			if played {
				if m.Cursor[1] < 19-1 {
					m.Cursor[1]++
				} else if m.Cursor[0] < 19-1 {
					m.Cursor[0]++
				} else {
					m.Cursor[1]--
				}
			}
		}
	case tea.MouseMsg:
		// [0,0] is the top-left cell, A19
		x := (msg.X - m.Offsets[0]) / 2 // One cell is two characters wide
		y := msg.Y - m.Offsets[1]

		switch tea.MouseEvent(msg).Action {
		case tea.MouseActionMotion:
			if 0 <= x && x < 19 && 0 <= y && y < 19 {
				m.Cursor[0] = uint8(x)
				m.Cursor[1] = uint8(y)
			}
		case tea.MouseActionPress:
			if 0 <= x && x < 19 && 0 <= y && y < 19 {
				m.PlayMove(m.Cursor[0], m.Cursor[1])
			}
		}
	}

	return m, nil
}
