/*
Conversion utilities for converting between engine Board and TUI Game.
*/
package tui

import (
	"gongo/internal/engine"
)

// FIXME: Massively improve this architecture to avoid so much reallocation and conversion.

func BoardToGame(board *engine.Board) *Game {
	size := int(board.Size)
	game := Game{
		Size:   size,
		Board:  make([]byte, size*size),
		Turn:   board.Turn,
		Marked: [][2]byte{},
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			stone := board.GetStone(engine.Loc{X: byte(i + 1), Y: byte(j + 1)})
			index := i + size*j
			if stone&engine.COLOR_MASK == engine.BLACK {
				game.Board[index] = 1
			} else if stone&engine.COLOR_MASK == engine.WHITE {
				game.Board[index] = 2
			}
		}
	}

	// TESTING ONLY
	for i := range board.Marked {
		game.Marked = append(game.Marked, [2]byte{board.Marked[i].X - 1, board.Marked[i].Y - 1})
	}

	return &game
}

func GameToBoard(game *Game) *engine.Board {
	size := game.Size
	board := engine.NewBoard(byte(size))
	board.Turn = game.Turn

	// TODO: Once Zobrist hashes are implemented, use the tui.Game.Moves to initialize them.
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			cell := game.Board[i+size*j]
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
