package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "w":
			if m.cursor[1] > 0 {
				m.cursor[1]--
			}
		case "down", "s":
			if m.cursor[1] < 19-1 {
				m.cursor[1]++
			}
		case "left", "a":
			if m.cursor[0] > 0 {
				m.cursor[0]--
			}
		case "right", "d":
			if m.cursor[0] < 19-1 {
				m.cursor[0]++
			}
		case "enter", " ":
			if m.board[m.cursor[0]+m.cursor[1]*19] == 0 {
				m.moves = append(m.moves, m.cursor)
				m.board[m.cursor[0]+m.cursor[1]*19] = m.turn
				m.turn = 3 - m.turn

				// TODO: Refactor, along with this entire terminal viewer.
				if m.cursor[1] < 19-1 && m.board[m.cursor[0]+(m.cursor[1]+1)*19] == 0 {
					m.cursor[1]++
				} else if m.cursor[1] > 0 && m.board[m.cursor[0]+(m.cursor[1]-1)*19] == 0 {
					m.cursor[1]--
				} else if m.cursor[0] < 19-1 && m.board[(m.cursor[0]+1)+m.cursor[1]*19] == 0 {
					m.cursor[0]++
				} else if m.cursor[0] > 0 && m.board[(m.cursor[0]-1)+m.cursor[1]*19] == 0 {
					m.cursor[0]--
				}
			}
		}
	}

	return m, nil
}
