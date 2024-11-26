package main

import (
	"fmt"
	"os"

	"gongo/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	initialBoardSize = 13
)

func main() {
	p := tea.NewProgram(tui.NewModel(initialBoardSize), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("An error occurred while trying to show my new Go game player: %v", err)
		os.Exit(1)
	}
}
