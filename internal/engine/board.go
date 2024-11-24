/*
Fast board implementation for Go games.

Unsure of the best way. Below are my options.
- [361]bytes (current)
- [19][19]bytes (presumably no difference)
- 3 [19]uint32s (fast since it's all bitwise)
*/

package engine

import (
	"math/rand"
	"time"
)

const (
	EMPTY byte = 0
	BLACK byte = 1
	WHITE byte = 2
	MARK  byte = 4 // Marked by our exploration algorithm.
	EDGE  byte = 7 // Edge of the board (always marked).
)

const (
	COLOR_MASK byte = 0b11
)

func opponent(color byte) byte {
	return 3 - color
}

// Fast internal Go board representation.
type Board struct {
	Size  byte
	Board []byte // size*size array with 1 padding on each side for the edge.
	Turn  byte
	// TODO: Consider a representation for the Ko square.
	// TODO: Consider storing the chains and liberties in the board.
	// TODO: Consider storing a pointer to a scoring object.
	// TODO: Zobrist hashes.
}

// Individual board location.
type Loc struct {
	X, Y byte
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

// Get index of Loc on a one-dimensional grid
func (l *Loc) Linear(width byte) uint16 {
	return uint16(l.X) + uint16(l.Y)*uint16(1+width+1)
}

func (l *Loc) String() string {
	return string('{' + l.X + ',' + l.Y + '}')
}

type Group struct {
	Color     byte // EMPTY, BLACK, WHITE
	Stones    []Loc
	Liberties []Loc
}

func NewBoard(size byte) *Board {
	board := Board{Size: size, Turn: BLACK}
	len := int(size + 2)

	board.Board = make([]byte, len*len)

	for i := range board.Board {
		if i < len || i%len == 0 || i%len == len-1 || i >= len*len-len {
			board.Board[i] = EDGE
		} else {
			board.Board[i] = EMPTY
		}
	}

	return &board
}

func (b *Board) GetStone(l Loc) byte {
	return b.Board[l.Linear(b.Size)]
}

// Get stone value then mark it.
func (b *Board) GetAndMarkStone(l Loc) byte {
	stone := b.Board[l.Linear(b.Size)]
	b.SetMark(l)
	return stone
}

// Sets the MARK value of a location.
func (b *Board) SetMark(l Loc) {
	b.Board[l.Linear(b.Size)] |= MARK
}

// Clears the MARK value of a location.
func (b *Board) UnsetMark(l Loc) {
	b.Board[l.Linear(b.Size)] &= ^MARK
}

// Get unmarked values adjacent to a location.
func (b *Board) GetUnmarkedAdjacent(l Loc) []byte {
	adjs := []byte{}

	for _, v := range l.Adjacent() {
		stone := b.GetAndMarkStone(v)
		if stone&MARK == 0 {
			adjs = append(adjs, stone)
		}
	}

	return adjs
}

func (b *Board) SetStone(l Loc, color byte) {
	b.Board[l.Linear(b.Size)] = color
}

func (b *Board) UnsetStone(l Loc) {
	b.Board[l.Linear(b.Size)] = 0
}

func (b *Board) MakeMove(l Loc) bool {
	if b.GetStone(l) != 0 {
		return false // Cannot place a stone on top of another stone
	}

	// Simulate the move (needs to later be undone)
	b.SetStone(l, b.Turn)
	opp := opponent(b.Turn)

	/*
		group := b.GetGroup(l, true)
		suicide := len(group.Liberties) == 0

		if suicide && len(group.Stones) > 1 {
			// FIXME: Dependent on ruleset. https://qr.ae/p2EkE1

			// TODO: Use `defer` to always unset / unmark, then modify usage. https://go.dev/doc/effective_go#defer
			b.UnsetStone(l)
			return false // Cannot play a multi-stone suicide
		}

		adjs := l.Adjacent()
		deadGroups := []Group{}
		toUnmark := []Loc{}

		for _, v := range adjs {
			stone := b.GetStone(v)
			isOpp := stone&COLOR_MASK == opp
			isMarked := stone&MARK == MARK

			if !isOpp || isMarked {
				continue // Not a new group
			}

			// Leave group stones marked for preventing duplicated groups
			group := b.GetGroup(v, false)

			toUnmark = append(toUnmark, group.Stones...)

			if len(group.Liberties) == 0 {
				deadGroups = append(deadGroups, group)
			}
		}

		for _, v := range toUnmark {
			b.UnsetMark(v)
		}

		if suicide && len(deadGroups) == 0 {
			// FIXME: Dependent on ruleset. https://qr.ae/p2EkE1
			b.UnsetStone(l)
			return false // Cannot play self-atari if no groups are being killed
		}

		for _, g := range deadGroups {
			for _, location := range g.Stones {
				b.UnsetStone(location)
			}
		}*/

	b.Turn = opp

	return true
}

// Traverse across a group of like locations.
// Pre: There must be a stone at `l`.
func (b *Board) GetGroup(l Loc, unmark bool) Group {
	c := b.GetStone(l)
	g := Group{Color: c & COLOR_MASK}

	// Flood-fill exploration for efficient grouping
	active := []Loc{l}
	newActive := []Loc{}

	for len(active) > 0 {
		newActive = newActive[:0]

		for _, v := range active {
			adjLocs := l.Adjacent()

			for i, stone := range b.GetUnmarkedAdjacent(v) {
				s := stone & COLOR_MASK
				switch {
				case s == EMPTY:
					g.Liberties = append(g.Liberties, adjLocs[i])
				case s == c&COLOR_MASK:
					g.Stones = append(g.Stones, adjLocs[i])
				case s != EDGE:
					newActive = append(newActive, adjLocs[i])
				}
			}
		}

		active = newActive
	}

	// Conditionally unmark stones
	if unmark {
		for _, v := range g.Stones {
			b.UnsetMark(v)
		}
	}

	// Always unmark liberties
	for _, v := range g.Liberties {
		b.UnsetMark(v)
	}

	return g
}

// A random move generator, to be used when I need a "working" computer player.
func (b *Board) MakeRandomMove() bool {
	rand.Seed(time.Now().UnixNano())

	size := int(b.Size)

	// Has a 63.2% change of finding the "final" move
	for tries := 0; tries < size*size; tries++ {
		x := byte(rand.Intn(size) + 1)
		y := byte(rand.Intn(size) + 1)
		loc := Loc{X: x, Y: y}

		if b.MakeMove(loc) {
			return true
		}
	}

	return false
}
