package tui

import "fmt"

func (m Model) View() string {
	turn := "Black"
	if m.turn == 2 {
		turn = "White"
	}
	s := fmt.Sprintf("\nGongo Terminal Player\n\n%s's turn to place a stone.\n\n", turn)

	// Render the board
	for y := 0; y < 19; y++ {
		for x := 0; x < 19; x++ {
			cell := " -"
			switch m.board[x+y*19] {
			case 1:
				cell = "⚫"
			case 2:
				cell = "⚪"
			}
			if x == m.cursor[0] && y == m.cursor[1] {
				if turn == "Black" {
					cell = "◾" // or ⬛
				} else {
					cell = "◽" // or ⬜
				}
			}
			s += cell
		}
		s += "\n"
	}

	s += "\nPress q to quit.\n"

	return s
}
