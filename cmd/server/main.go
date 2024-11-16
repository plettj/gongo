package main

import (
	"fmt"
	"os"

	"gongo/internal/board"
	"gongo/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	sgf := board.FileToString("_files/sgf/40799184-110-plettj-gucci_gang.sgf")
	sgf = board.ParseSgf(sgf)

	fmt.Println(sgf)

	if false {
		p := tea.NewProgram(tui.InitialModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("An error occurred while trying to show my new Go game player: %v", err)
			os.Exit(1)
		}
	}

	// router := api.NewRouter()
	// log.Println("Server is running on port 8080.")
	// log.Fatal(http.ListenAndServe(":8080", router))
}
