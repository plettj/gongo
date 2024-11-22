# Developer Notes

## Engine to-dos

- [x] Decide on an API protocol (gRPC, REST, GraphQL) we'll use, and router (gorilla/mux, gin, go) we'll use. Decisions: REST + gorilla/mux.
- [x] Set up and test a first simple API request and response example to mimic the creation of a game. [Go REST](https://medium.com/@Moesif/building-a-restful-api-with-go-dbd6e7aecf87) tutorial. [Snowflake ID](https://en.wikipedia.org/wiki/Snowflake_ID) specification.
- [x] Create a bare-bones Go viewer in a nice CLI. Example: [Tetrgo](https://github.com/Broderick-Westrope/tetrigo?tab=readme-ov-file) built with [BubbleTea](https://github.com/charmbracelet/bubbletea?tab=readme-ov-file).
- [ ] Research/experiment the computationally fastest way to represent a large Go game's board, and implement it.
      Representing Ko even with SuperKo: [My Forum Question](https://forums.online-go.com/t/is-there-ever-more-than-1-move-that-violates-positional-superko/53724) about it. [Discussion w/resources](https://forums.online-go.com/t/superko-rules/32466/4). [Situational Example](https://online-go.com/demo/view/580802). [Positional Example](https://online-go.com/demo/view/580801).
- [ ] Research/experiment and decide on the various rules of Go I'll be supporting.
- [ ] Research/experiment the computationally fastest way people have made move generators for Go.
- [ ] (large) Implement a very **basically intelligent** Go move generator.
      Computing if a group is pass-alive: https://senseis.xmp.net/?BensonsAlgorithm
      Storing moves of a game in Go: https://red-bean.com/sgf/go.html
      General go programming resource: https://senseis.xmp.net/
- [ ] Be able to load Go games into the viewer based on [sgf](https://red-bean.com/sgf/go.html). Examples [here](https://red-bean.com/sgf/examples/).
      SGF's [way of storing branches](https://red-bean.com/sgf/var.html) comes from [this](https://en.wikipedia.org/wiki/Newick_format).
      Summary of Go SGF [here](https://en.wikipedia.org/wiki/Smart_Game_Format#About_the_format) or examples in the repo.

## Usage

Run a go program (short for `go run <filepath>`):

```bash
just run
```

Get a new dependency:

```bash
go get <path-to-file>
```

Clean up unused dependencies:

```bash
go mod tidy
```

### Miscellaneous

<details>
<summary>My VSCode settings</summary>
<br>

VSCode `settings.json` Golang entry:

```json
"[go]": {
    "editor.insertSpaces": true,
    "editor.formatOnSave": true,
    "editor.defaultFormatter": "golang.go",
  },
```

Note that I'm not disabling the import organization, as despite how [annoying](https://stackoverflow.com/questions/19560334/how-to-disable-golang-unused-import-error) it is, it's better to have it than to [not have it](https://stackoverflow.com/a/61316426/8360465).

</details>

- Turns out [online-go](https://online-go.com/) is ubiquitously the best, but then there's [gokgs](https://www.gokgs.com/), [fox weiqi](https://www.foxwq.com/), and [pandanet](https://pandanet-igs.com/) roughly in that order.
- For documenting my APIs I should use [OAS](https://spec.openapis.org/oas/latest.html).
- Rules for [Go](https://en.wikipedia.org/wiki/Rules_of_Go)
- Rules for [Weiqi](https://www.cs.cmu.edu/~wjh/go/rules/Chinese.html).
- Rules for [AGA](https://www.cs.cmu.edu/~wjh/go/rules/AGA.html)
- Rules for [Baduk](http://english.baduk.or.kr/sub02_02.htm)
- Rules for [Ing's SST](https://www.cs.cmu.edu/~wjh/go/rules/KSS.html)
- Rules for [New Zealand](https://go.org.nz/index.php/about-go/new-zealand-rules-of-go)
- Resource for [Japanese vs Chinese](https://polgote.com/en/blog/go-rules-japanese-vs-chinese/) rules.
- Reached out to [Karl Fogel](https://red-bean.com/people.html) from red-bean about SGF and he pointed me towards [Arno](mailto:ahollosi@xmp.net).
- // I should create an [XML](https://www.w3schools.com/xml/) definition for GGN and SGN.
- // I should create an [ANTLR](https://github.com/antlr/grammars-v4/blob/master/pgn/PGN.g4) standard for GGN and SGN.
- Programming the scoring logic: [thread](https://www.reddit.com/r/cbaduk/comments/15tsaxj/comment/jwn5kku/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button).

### Longer-term To-dos

- Make my custom-written snowflake generator a public go package.
- Create example `.ggn` and `.sgn` files.
- Define on a `.md` file the exact specification of games my online player will be capable of.
