package tui

type Model struct {
	board  [19 * 19]int // 0 is empty, 1 is black, 2 is white
	cursor [2]int       // Which cell is currently selected
	moves  [][2]int     // List of moves made
	turn   int          // 1 is black, 2 is white
}

func InitialModel() Model {
	return Model{
		board:  [19 * 19]int{},
		cursor: [2]int{9, 9},
		moves:  [][2]int{},
		turn:   1,
	}
}
