/*
Gongo Board Smart Game Format Parser

This parser should support the following interpretation of SGF Go game records.

Resource: https://en.wikipedia.org/wiki/Smart_Game_Format#About_the_format
          Examples in this repository, too.

Lines beginning with // are lines I am choosing not to support.

- FF[4] - File Format: SGF version 4
- CA[UTF-8] - Character encoding
- GM[1] - Game type (1 = Go)
- DT[YYYY-MM-DD] - Date of the game
- PC[platform: url] - Place of the game
- CP[text] - Copyright information
- EV[name] - Event name
- RO[name] - Round number
- GN[name] - Game name
- PB[name] - Black player name
- PW[name] - White player name
- BR[rank] - Black player rank
- WR[rank] - White player rank
- BT[team] - Black team name
- WT[team] - White team name
- TM[time] - Time limit or control, in seconds
- OT[time] - Overtime control
// ON[text] - Opening (Fuseki) information (rarely used)
- PL[color] - Color of player to play next
- HA[num] - Handicap, for Baduk rules http://english.baduk.or.kr/sub02_02.htm
- RE[result] - Result of the game (?, W+#, B+#, W+R (resignation), Draw/void/"No result" [TODO: idk how draw is written])
- SZ[#x#] - Board size (num x num)
- KM[num] - Komi
- RU[name] - Ruleset of the game (Japanese, Chinese, ... [TODO: idk all the rulesets])
- AB[moves, ...] - Add initial black stones
- W[move] - White move
- B[move] - Black move
- C[text] - Comment

Format of a move:
- [xy] - Place a stone at column x from left, row y from top. Eg: AB[dc] means place a black stone at location {D, size - 2}.
- [] - Pass.

Format of branches:
- (; ... ) - Branches. Eg: (;W[dd];W[ee]) means two branches.

Format of time taken per move:
- Unspecified. I would like to create a new standard.

Format of total prisoners taken so far:
- Unspecified. I would like to create a new standard.
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
