package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"gongo/internal/tui"
)

func main() {
	p := tea.NewProgram(tui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("An error occurred while trying to show my new Go game player: %v", err)
		os.Exit(1)
	}

	// router := api.NewRouter()
	// log.Println("Server is running on port 8080.")
	// log.Fatal(http.ListenAndServe(":8080", router))
}
