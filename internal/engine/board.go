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
	MARK  uint8 = 4 // Marked by our exploration algorithm.
	// TODO: Consider a representation for the Ko square.
)

// Individual board location.
type Loc struct {
	X, Y uint8
}

// Get values adjacent to a location.
// Return: [N, E, S, W]
func (l *Loc) Adjacent() [4]Loc {
	return [4]Loc{
		{l.X, l.Y - 1},
		{l.X + 1, l.Y},
		{l.X, l.Y + 1},
		{l.X - 1, l.Y},
	}
}

func (l *Loc) String() string {
	return string('{' + l.X + ',' + l.Y + '}')
}

type Group struct {
	Color     uint8 // EMPTY, BLACK, WHITE
	Stones    []Loc
	Liberties []Loc
}

// Fast internal Go board representation.
type Board struct {
	size  uint8
	board [441]uint8
	turn  uint8
	// TODO: Consider storing the chains and liberties in the board.
}

func NewBoard(size uint8) *Board {
	return &Board{size: size}
}

func (b *Board) GetStone(l Loc) uint8 {
	return b.board[l.X+l.Y*b.size]
}

// Get stone value then mark it.
func (b *Board) GetAndMarkStone(l Loc) uint8 {
	stone := b.board[l.X+l.Y*b.size]
	b.board[l.X+l.Y*b.size] |= MARK
	return stone
}

// Sets the MARK value of a location.
func (b *Board) SetMark(l Loc) {
	b.board[l.X+l.Y*b.size] |= MARK
}

// Clears the MARK value of a location.
func (b *Board) UnsetMark(l Loc) {
	b.board[l.X+l.Y*b.size] &= ^MARK
}

// Get unmarked values adjacent to a location.
func (b *Board) GetUnmarkedAdjacent(l Loc) []uint8 {
	adjs := []uint8{}
	for _, v := range l.Adjacent() {
		stone := b.GetAndMarkStone(v)
		if stone&MARK == 0 {
			adjs = append(adjs, stone)
		}
	}
	return adjs
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
// Pre: There must be a stone at `l`.
func (b *Board) GetGroup(l Loc) Group {
	c := b.GetStone(l)
	g := Group{Color: c}

	// Flood-fill exploration for efficient grouping
	active := []Loc{l}
	newActive := []Loc{}
	for len(active) > 0 {
		newActive = newActive[:0]
		for _, v := range active {
			adjLocs := l.Adjacent()
			adjs := b.GetUnmarkedAdjacent(v)
			for i, stone := range adjs {
				s := stone & 0b11
				switch {
				case s == EMPTY:
					g.Liberties = append(g.Liberties, adjLocs[i])
				case s == c:
					g.Stones = append(g.Stones, adjLocs[i])
				case s != EDGE:
					newActive = append(newActive, adjLocs[i])
				}
			}
		}
		active = newActive
	}

	return g
}

func (b *Board) GetMoves() [361]bool {
	var moves [361]bool
	for i := range b.board {
		moves[i] = b.board[i]&0b11 == 0
	}
	return moves
}
