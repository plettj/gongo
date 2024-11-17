package board

// TODO: Define a spec for a board state that stores all board functions.

// 1. "Static Baduk (or Board) Notation" (SBN) will be my new spec for go games, equivalent to FEN and therefore supplementary to SGF.

// 2. More research into existing Go engines is required, before tackling the next tasks:

// 3. The board will have internal SBN states for the board state, and internal SGFs for the game state.

// BRAINSTORM 1: Fixed length array of ints.
type Board1 struct {
	size  int
	board [361]int // Increased memory overhead for smaller boards is negligible.
}

// BRAINSTORM 2: Binary.
type Board2 struct {
	size uint16
	w    [6]uint64 // 361 (384) bits for white stones
	b    [6]uint64 // 361 (384) bits for black stones
}

func NewBoard2(size uint16) *Board2 {
	return &Board2{size: size}
}

func (b *Board2) Set(x, y uint16, color bool) {
	// No bounds checking for perf reasons.
	index := x + y*b.size
	arrIndex := index / 64
	bitIndex := index % 64

	if color {
		b.w[arrIndex] |= 1 << bitIndex
	} else {
		b.b[bitIndex] |= 1 << bitIndex
	}
}

// BRAINSTORM 3: Large-size Bitboard, array of uint8s.
type Board3 struct {
	size  uint8
	board [361]uint8
}

// Because the uint8s have extra bits, we can store more relative board information in the same byte.
// Example potential location representations:
// 0b00000000: Empty
// 0b00000001: Black
// 0b00000010: White
// 0b00000101: Black with top liberty
// 0b00001001: Black with right liberty
// 0b00010001: Black with bottom liberty
// 0b00100001: Black with left liberty
// 0b00000111: Black with top and right liberties
// ...
// We're still not using the full byte, so we can store more information in the future.

// We'll also generally need to store a location type with methods:
type Location struct {
	x, y uint8
}

// We should also consider storing the wall of the board within the bitboard too...
