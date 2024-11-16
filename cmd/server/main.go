package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	board [361]int // 0 is empty, 1 is white, 2 is black
	cursor [2]int // which cell is currently selected
	moves [][3]int // list of moves made
}

func initialModel() model {
	return model{
		board: [361]int{},
		cursor: [2]int{0, 0},
		moves: [][3]int{},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func main() {
	// router := api.NewRouter()
	// log.Println("Server is running on port 8080.")
	// log.Fatal(http.ListenAndServe(":8080", router))
}