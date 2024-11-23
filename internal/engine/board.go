/*
Fast board implementation for Go games.

Unsure of the best way. Below are my options.
- [361]uint8s (current)
- [19][19]uint8s
- 3 [19]uint32s
*/

package engine

const (
	EMPTY uint8 = 0
	BLACK uint8 = 1
	WHITE uint8 = 2
	EDGE  uint8 = 3 // Edge of the board.
	MARK  uint8 = 7 // Marked by our exploration algorithm.
	// TODO: Consider a representation for the Ko square.
)

// Fast internal Go board representation.
type Board struct {
	size  uint8
	board [441]uint8
	turn  uint8
	// TODO: Consider storing the chains and liberties in the board.
}

// Individual board location.
type Loc struct {
	X, Y uint8
}

type Group struct {
	Color     uint8 // EMPTY, BLACK, WHITE
	Stones    []Loc
	Liberties []Loc
}

func (l Loc) String() string {
	return string('{' + l.X + ',' + l.Y + '}')
}

func NewBoard(size uint8) *Board {
	return &Board{size: size}
}

func (b *Board) GetStone(l Loc) uint8 {
	return b.board[l.X+l.Y*b.size]
}

func (b *Board) SetStone(l Loc, color uint8) {
	b.board[l.X+l.Y*b.size] = color
}

func (b *Board) RemoveStone(l Loc) {
	b.board[l.X+l.Y*b.size] = 0
}

func (b *Board) MakeMove(l Loc) bool {
	if b.GetStone(l) != 0 {
		return false
	}

	// IT APPEARS I'VE DISCOVERED THE FIRST MAJOR CHALLENGE OF THIS PROJECT: PROGRAMMING MOVE PLACEMENT (aka "rules").
	// Great reference point is this python implementation: https://github.com/maksimKorzh/wally/blob/main/wally.py

	// 0. Compute if this move is a suicide move. (boolean "suicide")

	// 1. Compute if any stones are being captured. If so, remove that chain. If not and "suicide" is true, return false. https://qr.ae/p2EkE1

	// 1.a. If we get here and "suicide" is true, set the KO value.

	// 2. Place the stone.

	b.SetStone(Loc{X: l.X, Y: l.Y}, b.turn)

	return true
}

// Traverse across a group of like locations.
// Pre: Must be a stone at `l`.
func (b *Board) GetGroup(l Loc) Group {
	c := b.GetStone(l)
	g := Group{Color: c}

	// Traverse the group with color `c`, marking cells, and incrementing liberties.

	// TODO: Do this!

	return g
}

func (b *Board) GetMoves() [361]bool {
	var moves [361]bool
	for i := range b.board {
		moves[i] = b.board[i]&0b11 == 0
	}
	return moves
}
