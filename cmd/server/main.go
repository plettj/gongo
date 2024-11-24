package main

import (
	"fmt"
	"os"

	"gongo/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	BOARD_SIZE = 13
)

func main() {
	p := tea.NewProgram(tui.NewModel(BOARD_SIZE), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("An error occurred while trying to show my new Go game player: %v", err)
		os.Exit(1)
	}

	// router := api.NewRouter()
	// log.Println("Server is running on port 8080.")
	// log.Fatal(http.ListenAndServe(":8080", router))
}
