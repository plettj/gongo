package tui

import (
	"fmt"
	"gongo/internal/engine"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func DrawBoard(b *engine.Board, cursor [2]int, border, lines, selected, flagged lipgloss.Style) string {
	size := int(b.Size)

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
	boardText += border.Render("   ╭"+boardBorder+"╮") + "\n"
	for y := 0; y < size; y++ {
		boardText += fmt.Sprintf("%2d %s", size-y, border.Render("│"))
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

			c := b.GetStone(engine.Loc{X: byte(x + 1), Y: byte(y + 1)}) & engine.COLOR_MASK
			if c != engine.EMPTY {
				switch c {
				case engine.BLACK:
					cell = "⚫"
				case engine.WHITE:
					cell = "⚪"
				}
				if x < size-1 {
					cell += "╶"
				} else {
					cell += " "
				}
			}

			cell = lines.Render(cell)
			if x == cursor[0] && y == cursor[1] {
				cell = selected.Render(cell)
			} else {
				for _, f := range b.Flagged {
					if x == int(f.X) && y == int(f.Y) {
						cell = flagged.Render(cell)
						break
					}
				}
			}

			boardText += cell
		}
		boardText += border.Render("│") + "\n"
	}
	boardText += border.Render("   ╰"+boardBorder+"╯") + "\n"

	return boardText
}

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
