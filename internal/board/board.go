// The spec for the two comprehensive board representations I'll be designing.
//
// - (GBN) General Baduk Notation - Covers everything (analogous to SGF) but more specifically to Go.
// - (SBN) Static Baduk Notation - Covers the minimal information required to pick up where one left off (unsure if it will support SuperKo)
//
// The general Baduk notation will be a superset of the static Baduk notation.

package board

type Color uint8

const (
	Black Color = 1
	White Color = 2
)

type PlayerInfo struct {
	Name string
	Rank string
	Team string
}

type Position struct {
	X int8 // Column index (0 to BoardSize-1), -1 for pass
	Y int8 // Row index (0 to BoardSize-1), -1 for pass
}

// Move represents a single move in the game.
type Move struct {
	Thread uint8 // Thread index, 0 for main thread (matches A in SGF)
	Color  Color
	Num    uint32 // Move number
	X      uint8  // Column index, -1 for pass
	Y      uint8  // Row index, -1 for pass
	Time   int32  // Time taken for the move (seconds, signed since incrementing/gifting exists (TODO: Unsure if should be absolute.))
	Text   string // Comment
}

// Full information of a game.
type Game struct {
	// Game Metadata
	Date      string  // Date of the game (e.g., "2024-11-17")
	GameName  string  // Name of the game or event
	Ruleset   string  // Ruleset used (e.g., "Japanese", "Chinese", "AGA")
	Result    string  // Result of the game (e.g., "B+R", "W+3.5", "Draw")
	Komi      float32 // Komi value (e.g., 6.5)
	BoardSize int8    // Board size (e.g., 19)

	// Player Information
	BlackPlayer PlayerInfo
	WhitePlayer PlayerInfo

	// Time Controls
	TimeLimit string // Main time limit in seconds (e.g., "300")
	Overtime  string // Overtime rules (e.g., "4x30 byo-yomi")

	// Initial Setup
	InitialBlack []Position // Positions of initial black stones (handicap stones)
	InitialWhite []Position // Positions of initial white stones

	// Move Sequence
	Moves []Move // Sequence of moves played during the game, also containing comments and other threads, for now

	// Other
	Copyright string // Copyright
	Event     string // Event name
	Round     string // Round number
	Place     string // Place of the game (e.g., "OGS: https://online-go.com/game/12345678")
}
