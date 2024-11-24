package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// FIXME: This is crazy but I'm not kidding. Right-clicking executes "a" "s" "s" in that order.
		//        Update it's because right-clicking in a windows cmd is supposed to be "paste" which is a
		//        serialized command, and essentially BubbleTea doesn't know how to handle that, so it
		//        serializes the input into a `msg` (regrettably).
		//        It was pure coincidence that the only three valid types that came through were "ass" ⚰️
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
		x := (msg.X - m.Offsets[0]) / 3 // One cell is 3 characters wide
		y := msg.Y - m.Offsets[1]

		switch tea.MouseEvent(msg).Action {
		case tea.MouseActionMotion:
			if 0 <= x && x < 19 && 0 <= y && y < 19 {
				m.Cursor[0] = byte(x)
				m.Cursor[1] = byte(y)
			}
		case tea.MouseActionPress:
			if tea.MouseEvent(msg).Button == tea.MouseButtonLeft && 0 <= x && x < 19 && 0 <= y && y < 19 {
				m.PlayMove(m.Cursor[0], m.Cursor[1])
			}
		}
	}

	return m, nil
}
