package board

type Board struct {
	size  uint8
	board [361]uint8
	prev  *Board
}

func NewBoard(size uint8) *Board {
	return &Board{size: size}
}

func (b *Board) CheckStone(x, y uint8) uint8 {
	return b.board[x+y*b.size]
}

func (b *Board) SetStone(x, y, color uint8) {
	b.board[x+y*b.size] = color
}

func (b *Board) RemoveStone(x, y uint8) {
	b.board[x+y*b.size] = 0
}

func (b *Board) GetMoves() [361]bool {
	var moves [361]bool
	for i := range b.board {
		moves[i] = b.board[i] == 0
	}
	return moves
}

// TODO: Define a spec for a board state that stores all board functions.

// 1. "Static Baduk (or Board) Notation" (SBN) will be my new spec for go games, equivalent to FEN and therefore supplementary to SGF.

// 2. More research into existing Go engines is required, before tackling the next tasks:

// 3. The board will have internal SBN states for the board state, and internal SGFs for the game state.
