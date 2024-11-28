# Developer Notes

## Server MVP _PLAN OF ATTACK_

This is my general plan for writing a very simple guest-based Go server, that still has appropriate matchmaking.

API endpoints:

- POST /api/games - allows client to request a match with settings (handles matchmaking).
- GET /api/games/{game-id} - gets the game state (for current/future spectators only).

Websocket endpoint(s):

- /api/ws/{game-id} - for mediating the real-time game between two players.

Server-side matching:

- Manage queues based on settings (currently, only board size)

Token generation and security:

- Generate game verification tokens upon match assignment.
- Authenticate the players on their tokens on every move. (Client has to store these somehow.)

For some of the technical stuff I'll be referencing [this](https://chatgpt.com/share/67466543-5480-800f-a21c-4a80b0d351ba) plan.

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

## Miscellaneous

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
- Rules for [Tromp-Taylor](https://senseis.xmp.net/?TrompTaylorRules)
- Resource for [Japanese vs Chinese](https://polgote.com/en/blog/go-rules-japanese-vs-chinese/) rules.
- Reached out to [Karl Fogel](https://red-bean.com/people.html) from red-bean about SGF and he pointed me towards [Arno](mailto:ahollosi@xmp.net).
- // I should create an [XML](https://www.w3schools.com/xml/) definition for GGN and SGN.
- // I should create an [ANTLR](https://github.com/antlr/grammars-v4/blob/master/pgn/PGN.g4) standard for GGN and SGN.
- Programming the scoring logic: [thread](https://www.reddit.com/r/cbaduk/comments/15tsaxj/comment/jwn5kku/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button).
- Great simple Go AI reference code base: [wally](https://github.com/maksimKorzh/wally/blob/main/wally.py)
- Consider [ONNX](https://onnx.ai/) standard for my method!
- Consider a "transformer network" on top of the [reinforcement learning](https://www.youtube.com/playlist?list=PLdAoL1zKcqTXFJniO3Tqqn6xMBBL07EDc) base.

### Longer-term To-dos

- Implement a simple [GTP](https://www.lysator.liu.se/~gunnar/gtp/gtp2-spec-draft2/gtp2-spec.html) interface.
- Make my custom-written snowflake generator a public go package.
- Create example `.ggn` and `.sgn` files.
- Define on a `.md` file the exact specification of game rulesets my online player will be capable of.
- Be able to load Go games into the viewer based on [sgf](https://red-bean.com/sgf/go.html). Examples [here](https://red-bean.com/sgf/examples/). SGF's [way of storing branches](https://red-bean.com/sgf/var.html) comes from [this](https://en.wikipedia.org/wiki/Newick_format). Summary of Go SGF: [here](https://en.wikipedia.org/wiki/Smart_Game_Format#About_the_format), or examples in this repo at `_files/sgf`.
- Implement a very **basically intelligent** Go move generator.
  Computing if a group is pass-alive: https://senseis.xmp.net/?BensonsAlgorithm
  Storing moves of a game in Go: https://red-bean.com/sgf/go.html
  General go programming resource: https://senseis.xmp.net/
  Example simple go-playing program: https://github.com/maksimKorzh/wally

### Domains

- 13go.com ($4,200)
- ahgo.com ($9,700)
- axgo.com ($9,700)
- dkgo.com ($8,300)
- ffgo.com ($3,900)
- g-go.com ($1,300)
- l-go.com ($900)

- noogg.com
- baduk.gg
- oh-go.com
- ahgo.gg
- gongo.gg
- gongo.dev
- gongo.ong
- gong.ong
- ogogn.com
- playgo.gg
- josigo.com
