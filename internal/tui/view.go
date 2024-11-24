package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m *Model) View() string {
	titleStyle := lipgloss.NewStyle() // Could be bold, unsure how it looks though.
	themedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6731F1"))
	mutedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#aaaaaa", Dark: "#555555"})
	emphasizedStyle := lipgloss.NewStyle().
		Bold(true)

	turn := "Black"
	if m.Game.Turn == 2 {
		turn = "White"
	}

	title := titleStyle.Render(fmt.Sprintf("%s%s%s",
		"Gon",
		themedStyle.Render("g"),
		titleStyle.Render("o tui"), // Wrapped in titleStyle due to undocumented lipgloss bug.
	))

	version := mutedStyle.Render(" - v1.2")

	topText := fmt.Sprintf("\n%s%s\n\n%s's turn.\n\n", title, version, emphasizedStyle.Render(turn))

	boardText := "     A B C D E F G H J K L M N O P Q R S T\n"
	boardText += fmt.Sprint(themedStyle.Render("   ╭───────────────────────────────────────╮")) + "\n"
	for y := 0; y < 19; y++ {
		boardText += fmt.Sprintf("%2d %s", 19-y, themedStyle.Render("│"))
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
			boardText += cell
		}
		boardText += fmt.Sprintf(" %s\n", themedStyle.Render("│"))
	}
	boardText += fmt.Sprint(themedStyle.Render("   ╰───────────────────────────────────────╯")) + "\n"

	bottomText := fmt.Sprint(mutedStyle.Render(" r: restart • q: exit\n"))

	return fmt.Sprintf("%s%s%s", topText, boardText, bottomText)
}
