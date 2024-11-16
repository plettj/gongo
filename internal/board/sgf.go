/*
Gongo Board Smart Game Format Parser

This parser supports the following interpretation of SGF Go game records.

- FF[4] - SGF version 4
- CA[UTF-8] - Character encoding
- GM[1] - Game type (1 = Go)
- DT[YYYY-MM-DD] - Date of the game
- PC[platform: url] - Place of the game
- GN[name] - Game name
- PB[name] - Black player name
- PW[name] - White player name
- BR[rank] - Black player rank
- WR[rank] - White player rank
- TM[time] - Time limit or control
- OT[time] - Overtime control
- RE[result] - Result of the game (?, W+#, B-#, draw [TODO: idk how draw is written])
- SZ[num] - Board size (num x num)
- KM[num] - Komi
- RU[name] - Ruleset of the game (Japanese, Chinese, ... [TODO: idk all the rulesets])
- AB[moves, ...] - Add initial black stones
- W[move] - White move
- B[move] - Black move

Format of a move:
- [xy] - Place a stone at column x from left, row y from top. Eg: AB[dc] means place a black stone at location {D, size - 2}.
- [] - Pass.
*/
package board

import (
	"os"
)

func FileToString(path string) string {
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

// TODO: Implement a function that converts an sgf string (or file) to a gng string
func ParseSgf(sgf string) string {
	return sgf
}
