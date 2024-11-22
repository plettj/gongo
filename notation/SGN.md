# Static Go Notation v0.1

_Currently Deprecated - using SGF[4]_

This is the official specification of **Static Go Notation (SGN)**, version `0.1`.

**SGN** is a human-readable file format designed to represent a snapshot of a [_Go_](<https://en.wikipedia.org/wiki/Go_(game)>) game, exactly the amount of information necessary to resume the game from that point. It is analogous to the [FEN](https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation) in chess.

SGN uses the file extension `.sgn`.

## Purpose and Function

SGN is a standard notation for representing a snapshot of a Go games in a human- and computer-readable manner. It is meant to be used used when transporting games from point A to point B, and includes exactly the information necessary to restart the game as-is.

If you're looking for a complete Go representation capable of storing full games and more, see the [General Go Notation](GGN.md)

---

# Specification

Here follows the concise but complete specification for the Static Go Notation.
