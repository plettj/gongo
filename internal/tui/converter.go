/*
Conversion utilities for converting between engine Board and TUI Game.
*/
package tui

import (
	"gongo/internal/engine"
)

func BoardToGame(board *engine.Board) *Game {
	game := Game{
		Size:  int(board.Size),
		Board: [19 * 19]byte{},
		Turn:  board.Turn,
	}

	for i := 0; i < 19; i++ {
		for j := 0; j < 19; j++ {
			stone := board.GetStone(engine.Loc{X: byte(i + 1), Y: byte(j + 1)})
			index := i + 19*j
			if stone&engine.COLOR_MASK == engine.BLACK {
				game.Board[index] = 1
			} else if stone&engine.COLOR_MASK == engine.WHITE {
				game.Board[index] = 2
			}
		}
	}

	return &game
}

func GameToBoard(game *Game) *engine.Board {
	board := engine.NewBoard(19)
	board.Size = byte(game.Size)
	board.Turn = game.Turn

	// TODO: Once Zobrist hashes are implemented, use the tui.Game.Moves to initialize them.
	for i := 0; i < 19; i++ {
		for j := 0; j < 19; j++ {
			cell := game.Board[i+19*j]
			loc := engine.Loc{X: byte(i + 1), Y: byte(j + 1)}
			if cell == 1 {
				board.SetStone(loc, engine.BLACK)
			} else if cell == 2 {
				board.SetStone(loc, engine.WHITE)
			}
		}
	}

	return board
}
