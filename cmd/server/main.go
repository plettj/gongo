package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// TODO: Capitalized because I want to export it from a different file later.
type Model struct {
	board  [19 * 19]int // 0 is empty, 1 is black, 2 is white
	cursor [2]int       // Which cell is currently selected
	moves  [][2]int     // List of moves made
	turn   int          // 1 is black, 2 is white
}

func initialModel() Model {
	return Model{
		board:  [19 * 19]int{},
		cursor: [2]int{9, 9},
		moves:  [][2]int{},
		turn:   1,
	}
}

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

		// The "enter" and "spacebar" keys place stones.
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

func (m Model) View() string {
	turn := "Black"
	if m.turn == 2 {
		turn = "White"
	}
	s := fmt.Sprintf("\nGongo Terminal Player\n\n%s's turn to place a stone.\n\n", turn)

	// Render the board
	for y := 0; y < 19; y++ {
		for x := 0; x < 19; x++ {
			cell := "╶╴"
			switch m.board[x+y*19] {
			case 1:
				cell = "⚫" // or ●?
			case 2:
				cell = "⚪" // or ○?
			}
			if x == m.cursor[0] && y == m.cursor[1] {
				cell = "[]"
			}
			s += cell
		}
		s += "\n"
	}

	s += "\nPress q to quit.\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("An error occurred while trying to show my new Go game player: %v", err)
		os.Exit(1)
	}

	// router := api.NewRouter()
	// log.Println("Server is running on port 8080.")
	// log.Fatal(http.ListenAndServe(":8080", router))
}
