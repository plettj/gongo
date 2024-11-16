package board

type Board struct {
	board [19 * 19]int // 0 is empty, 1 is black, 2 is white
}

// TODO: Define a spec for a board state that stores all board functions.

// 1. "Static Baduk Notation" (SBN) will be my new spec for go games, equivalent to FEN and therefore supplementary to SGF.

// 2. More research into existing Go engines is required, before tackling the next tasks:

// 3. The board will have internal SBN states for the board state, and internal SGFs for the game state.
