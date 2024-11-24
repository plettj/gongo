package tui

import (
	"fmt"
)

func (m *Model) View() string {
	turn := "Black"
	if m.Game.Turn == 2 {
		turn = "White"
	}
	s := fmt.Sprintf("\nGongo Terminal Player 1.2\n\n%s's turn.\n\n", turn)

	// Render the board
	s += "     A B C D E F G H J K L M N O P Q R S T\n"
	s += "   ┌───────────────────────────────────────┐\n"
	for y := 0; y < 19; y++ {
		s += fmt.Sprintf("%2d │", 19-y)
		for x := 0; x < 19; x++ {
			cell := " -"
			if (x+3)%6 == 0 && (y+3)%6 == 0 {
				cell = " +"
			}
			switch m.Game.Board[x+y*19] {
			case 1:
				cell = "⚫"
			case 2:
				cell = "⚪"
			}
			if uint8(x) == m.Cursor[0] && uint8(y) == m.Cursor[1] {
				if turn == "Black" {
					cell = "◾" // or ⬛?
				} else {
					cell = "◽" // or ⬜?
				}
			}
			s += cell
		}
		s += " │\n"
	}
	s += "   └───────────────────────────────────────┘\n"

	s += "\nPress r to restart. Press q to quit.\n"

	return s
}
