# General Go Notation v0.1

This is the official specification of **General Go Notation (GGN)**, version `0.1`.

**GGN** is a human-readable file format designed to represent a complete game of [_Go_](<https://en.wikipedia.org/wiki/Go_(game)>) (also known as _Baduk_ or _Weiqi_), including annotations, variations, and metadata. It is an alternative to the [Smart Game Format (SGF)](https://www.red-bean.com/sgf/) for Go, [GM[1]](https://www.red-bean.com/sgf/go.html).

It is being developed in light of [GM[1]](https://www.red-bean.com/sgf/go.html) with inspiration from the [Portable Game Notation (PGN)](https://en.wikipedia.org/wiki/Portable_Game_Notation) in Chess for three primary reasons:

- To address the increasing demands of fully-featured online Go platforms.
- To innovate on SGF for Go, in the modern world of software.
- To ease the path for new contributors to the Go community.

GGN uses the file extension `.ggn`.

## Purpose and Function

GGN is a standard notation for representing complete Go games in a human- and computer-readable manner. It is meant to be used when transporting games from point A to point B, and includes all necessary data to describe the individual game, comments, annotations, time, metadata, and other relevant circumstances in full.

GGN also defines notation for representing many Go games, much like the [PGN](https://en.wikipedia.org/wiki/Portable_Game_Notation) in Chess.

If you're looking for a shorter, simpler representation, see the [Static Go Notation](SGN.md) specification.

## Converting between GGN and SGF

GGN also includes a concise specification for converting files _to_ and _from_ SGF, for backwards compatibility.

_TODO: I plan to develop an online tool to convert files between GGN and SGF._

---

# Specification

Here follows the concise but complete specification for the General Go Notation.

<details>
<summary>General structure of property specification.</summary>
<br>

The way GGN specifies properties is intentionally concise, while fully descriptive. Some properties have unique formats, but the general property format is as follows.

**Human Attribute Name:** `Code`

`type` _OR_ [`list`, `of`, `values`]

Description of the meaning or usage of the attribute.

<details>
<summary>Example of <code>Code</code></summary>
<br>

```
[Code: value]
```

</details>

---

</details>

### Required Properties

These properties must be included in all valid `.ggn` files.

---

## Acknowledgements

Both GGN and SGN are primarily inspired by the effectiveness of [PGN](https://en.wikipedia.org/wiki/Portable_Game_Notation) and [FEN](https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation) in Chess, but of course use properties from [SGF](https://www.red-bean.com/sgf/) and [GM[1]](https://www.red-bean.com/sgf/go.html), the current standard for representing Go games.

At the time of writing, GGN and SGN are developed solely by [Josiah Plett](https://plett.dev/). Hopefully that can change soon.

## License

[GPL3](../LICENSE).
