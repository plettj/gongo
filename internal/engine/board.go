/*
Fast board implementation for Go games.

Unsure of the best way. Below are my options.
- [361]uint8s (current)
- [19][19]uint8s
- 3 [19]uint32s
*/

package engine

// Fast internal Go board representation.
type Board struct {
	size  uint8
	board [361]uint8 // 0 = empty, 1 = black, 2 = white, 3 = ko
}

// Individual board location.
type Loc struct {
	X, Y uint8
}

// Representation of any stone state.
type Stone struct {
	X, Y  uint8
	Color uint8
}

func NewBoard(size uint8) *Board {
	return &Board{size: size}
}

func (b *Board) CheckStone(l Loc) uint8 {
	return b.board[l.X+l.Y*b.size]
}

func (b *Board) SetStone(s Stone) {
	b.board[s.X+s.Y*b.size] = s.Color
}

func (b *Board) RemoveStone(l Loc) {
	b.board[l.X+l.Y*b.size] = 0
}

func (b *Board) GetMoves() [361]bool {
	var moves [361]bool
	for i := range b.board {
		moves[i] = b.board[i] == 0
	}
	return moves
}
