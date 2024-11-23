/*
Conversion utilities for converting between engine Board and TUI Game.
*/
package engine

import (
	"gongo/internal/tui"
)

func BoardToGame(board Board) *tui.Game {
	game := tui.Game{
		Board: [19 * 19]uint8{},
		Turn:  board.Turn,
	}

	for i := 1; i < 19+1; i++ {
		for j := 1; j < 19+1; j++ {
			stone := board.GetStone(Loc{X: uint8(i), Y: uint8(j)})
			if stone&COLOR_MASK == BLACK {
				game.Board[(i-1)+19*(j-1)] = 1
			} else if stone&COLOR_MASK == WHITE {
				game.Board[(i-1)+19*(j-1)] = 2
			}
		}
	}

	return &game
}

func GameToBoard(game tui.Game) *Board {
	board := NewBoard(19)
	board.Turn = game.Turn

	// TODO: Once Zobrist hashes are implemented, use the tui.Game.Moves to initialize them.
	for i := 1; i < 19+1; i++ {
		for j := 1; j < 19+1; j++ {
			cell := game.Board[(i-1)+19*(j-1)]
			loc := Loc{X: uint8(i), Y: uint8(j)}
			if cell == 1 {
				board.SetStone(loc, BLACK)
			} else if cell == 2 {
				board.SetStone(loc, WHITE)
			}
		}
	}

	return board
}
