package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Helper for getting hoshi points (star points) designed by me.
func isHoshi(x, y, size int) bool {
	if size%2 == 0 {
		return false
	} else if size >= 15 {
		return (x == 3 || x == size/2 || x == size-4) && (y == 3 || y == size/2 || y == size-4)
	} else if size == 13 {
		return (x == 3 || x == size-4) && (y == 3 || y == size-4) || (x == y && x == size/2)
	} else if size >= 7 {
		return (x == 2 || x == size-3) && (y == 2 || y == size-3) || (size > 7 && x == y && x == size/2)
	} else {
		return x%2 == 1 && y%2 == 1
	}
}

// TODO: Extract myself some *helper functions* :sparkles:
func (m *Model) View() string {
	size := int(m.Board.Size)

	titleStyle := lipgloss.NewStyle() // Could be bold, unsure how it looks though.
	themedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6731F1"))
	mutedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#aaaaaa", Dark: "#555555"})
	emphasizedStyle := lipgloss.NewStyle().
		Bold(true)
	selectedCellStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#FFFFFF"))
	markedCellStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#6731F1"))
	inactiveTabStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#555555"))
	activeTabStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFFFFF"))

	turn := "Black"
	if m.Board.Turn == 2 {
		turn = "White"
	}

	title := titleStyle.Render(fmt.Sprintf("%s%s%s",
		"Gon",
		themedStyle.Render("g"),
		titleStyle.Render("o tui"), // Wrapped in titleStyle due to undocumented lipgloss bug.
	))

	version := mutedStyle.Render(" - v1.4")

	topText := fmt.Sprintf("\n%s%s\n\n%s's turn.\n\n", title, version, emphasizedStyle.Render(turn))

	boardBorder := strings.Repeat("───", size)
	boardLetters := "   "
	for i := 0; i < size; i++ {
		letter := 'A' + rune(i)
		if letter >= 'I' {
			letter++
		}
		boardLetters += fmt.Sprintf("  %c", letter)
	}

	boardText := boardLetters + "\n"
	boardText += themedStyle.Render("   ╭"+boardBorder+"╮") + "\n"
	for y := 0; y < size; y++ {
		boardText += fmt.Sprintf("%2d %s", size-y, themedStyle.Render("│"))
		for x := 0; x < size; x++ {
			cell := "─┼─"

			if x == 0 && y == 0 {
				cell = " ┌─"
			} else if x == 0 && y == size-1 {
				cell = " └─"
			} else if x == size-1 && y == 0 {
				cell = "─┐ "
			} else if x == size-1 && y == size-1 {
				cell = "─┘ "
			} else if x == 0 {
				cell = " ├─"
			} else if x == size-1 {
				cell = "─┤ "
			} else if y == 0 {
				cell = "─┬─"
			} else if y == size-1 {
				cell = "─┴─"
			} else if isHoshi(x, y, size) {
				cell = "─┿─"
			}

			if m.Board.Board[x+y*size] != 0 {
				switch m.Board.Board[x+y*size] {
				case 1:
					cell = "⚫"
				case 2:
					cell = "⚪"
				}
				if x < size-1 {
					cell += "╶"
				} else {
					cell += " "
				}
			}

			cell = mutedStyle.Render(cell)
			if x == m.Cursor[0] && y == m.Cursor[1] {
				cell = selectedCellStyle.Render(cell)
			} else {
				for _, marked := range m.Board.Flagged {
					if x == int(marked.X) && y == int(marked.Y) {
						cell = markedCellStyle.Render(cell)
						break
					}
				}
			}

			boardText += cell
		}
		boardText += themedStyle.Render("│") + "\n"
	}
	boardText += themedStyle.Render("   ╰"+boardBorder+"╯") + "\n"

	// TODO: Set up tabs to actually work. Include two settings:
	//       - Board size
	//       - Whether to play against a bot
	tabs := ""
	for i, tab := range TABS {
		if tab == m.ActiveTab {
			tabs += activeTabStyle.Render(tab)
		} else {
			tabs += inactiveTabStyle.Render(tab)
		}
		if i < len(TABS)-1 {
			tabs += inactiveTabStyle.Render(" | ")
		}
	}

	commands := mutedStyle.Render("r: restart • q: exit\n")
	bottomText := fmt.Sprintf("\n%s\n\n%s\n", tabs, commands)

	return fmt.Sprintf("%s%s%s", topText, boardText, bottomText)
}
