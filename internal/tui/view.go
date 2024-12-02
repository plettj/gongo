package tui

import (
	"fmt"
)

func (m *Model) View() string {

	turn := "Black"
	if m.Board.Turn == 2 {
		turn = "White"
	}

	title := fmt.Sprintf("%s%s%s",
		"Gon",
		Styles["themed"].Render("g"),
		"o tui",
	)

	version := Styles["muted"].Render(" - v1.4")

	topText := fmt.Sprintf("\n%s%s\n\n%s's turn.\n\n", title, version, Styles["emphasized"].Render(turn))

	boardText := DrawBoard(&m.Board, m.Cursor, Styles["themed"], Styles["muted"], Styles["cell-selected"], Styles["cell-flagged"])

	// TODO: Set up tabs to actually work. Include two settings:
	//       - Board size
	//       - Whether to play against a bot
	tabs := ""
	for i, tab := range TABS {
		if tab == m.ActiveTab {
			tabs += Styles["tab-active"].Render(tab)
		} else {
			tabs += Styles["tab-inactive"].Render(tab)
		}
		if i < len(TABS)-1 {
			tabs += Styles["tab-inactive"].Render(" | ")
		}
	}

	commands := Styles["muted"].Render("r: restart • q: exit\n")
	//bottomText := fmt.Sprintf("\n%s\n\n%s\n", tabs, commands)
	bottomText := fmt.Sprintf("\n%s\n", commands)

	return fmt.Sprintf("%s%s%s", topText, boardText, bottomText)
}
