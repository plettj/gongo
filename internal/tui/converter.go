/*
Conversion utilities for converting between engine Board and TUI Game.
*/
package tui

import (
	"gongo/internal/engine"
)

func BoardToGame(board *engine.Board) *Game {
	game := Game{
		Board: [19 * 19]uint8{},
		Turn:  board.Turn,
	}

	for i := 0; i < 19; i++ {
		for j := 0; j < 19; j++ {
			stone := board.GetStone(engine.Loc{X: uint8(i + 1), Y: uint8(j + 1)})
			if stone&engine.COLOR_MASK == engine.BLACK {
				game.Board[i+19*j] = 1
			} else if stone&engine.COLOR_MASK == engine.WHITE {
				game.Board[i+19*j] = 2
			}
		}
	}

	return &game
}

func GameToBoard(game *Game) *engine.Board {
	board := engine.NewBoard(19)
	board.Turn = game.Turn

	// TODO: Once Zobrist hashes are implemented, use the tui.Game.Moves to initialize them.
	for i := 0; i < 19; i++ {
		for j := 0; j < 19; j++ {
			cell := game.Board[i+19*j]
			loc := engine.Loc{X: uint8(i + 1), Y: uint8(j + 1)}
			if cell == 1 {
				board.SetStone(loc, engine.BLACK)
			} else if cell == 2 {
				board.SetStone(loc, engine.WHITE)
			}
		}
	}

	return board
}
