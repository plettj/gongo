# Go SGF Parser

This module is a [Smart Game Format (SGF)](https://en.wikipedia.org/wiki/Smart_Game_Format#About_the_format) parser and writer for my Go game implementation.

**References**

- [rooklift/sgf](https://pkg.go.dev/github.com/rooklift/sgf) - SGF parser written in Go.
- [seehun/go-sgf](https://pkg.go.dev/seehuhn.de/go/sgf) - SGF parser written in Go.

### SGF Fields (to be) Supported

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
- ON[text] - Opening (Fuseki) information (rarely used)
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

### Module Architecture

(initially phrased as a prompt to ChatGPT, which was very unwise)

- `board.go`: Defines the SGF-compatible board structure (Board), which contains every property one might have during sgf stored in a reasonable way such as key<>value pairs when necessary or a Tree when necessary. Includes Zobrist hashes for positions that have already appeared.
- `test/board.go`: Test file for testing board.go, specifically just its methods tho, any other methods you've added to the board as you go. Throughout this file and others, the Board value should only contain valid SGF file information ever; never can be invalid.
- `node.go`: Represents an individual game \*Tree node (move). It has decodeMove, encodeMove, and any others you see as relevant.
- `test/node.go`: Tests the methods that are in node.go.
- `read.go`: Reads an sgf file which ofc can contain one or many sgf files. Has a parse method, and then the private parseGameTree and parseNode used internally. Also has next, backup, and peek. (but only implement them if you found you needed them in the main parse method. Also defines an error type.
- `test/read.go`: Tests the methods that are in read.go.
- `scanner.go`: Breaks a string into tokens basically.
- `test/scanner.go`: Tests the methods that are in scanner.go.
- `write.go`: Writes a \*Board into a .sgf file. Includes a method that takes a list of games via many boards. Functionally you can consider it is the reverse of read.go, similar approach.
- `test/write.go`: Essentially analogous to test/read.go, tests the writing of sgf files.
