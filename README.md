<div align="center" style="padding: 0px 50px"><img width="100%" src="gallery/gongo-banner.png" /></div>

<div align="center" style="text-align: center; width: 100%">
<h1>gongo - The Go Go Engine</h1>
</div>

This repo is a server-side [Go](<https://en.wikipedia.org/wiki/Go_(game)>) (game) engine written in [Go](https://go.dev/) (language).

I'm building it primarily to prepare for my [upcoming job](https://www.maximahq.com/) in the Go language, so I'll be documenting some of that learning here.

---

## Engine to-dos

- [ ] Decide on a routing protocol the engine API will use, and implement one endpoint for said protocol.
- [ ] Create a bare-bones Go viewer in a nice CLI.
- [ ] Be able to load Go games into the viewer based on a standard game representation.
- [ ] Research the computationally fastest way to represent a large Go game's board, and implement it.
- [ ] Research and decide on the various rules of Go I'll be supporting.
- [ ] Write a spec for the rules my engine supports.
- [ ] Research the computationally fastest way people have made move generators for Go.
- [ ] (large) Implement a Go move generator.

---

## Project Infrastructure

- **Go Engine Code:** Referencing [GNU Go](https://www.gnu.org/software/gnugo/gnugo_4.html#SEC39) and [KataGo](https://github.com/lightvector/KataGo/blob/master/cpp/README.md)
- **Architecture:** Server/Client (until [Go WASM](https://github.com/golang/go/issues/28631) supports [threads](https://caniuse.com/wasm-threads))
- **Hosting:** [Railway](https://railway.app/)
- **Framework:** None (saved by [the crowd](https://www.reddit.com/r/golang/comments/q3r8qo/do_you_guys_use_frameworks_with_go_for_backend/))

## Engine

Game engines will always have [tree traversal](https://en.wikipedia.org/wiki/Tree_traversal) and [evaluation](https://www.chessprogramming.org/Evaluation) at their roots. But beyond that, anything goes.

A top engine can either be _heuristic-based_ or learn dynamically from _self-play_. Learning through self-play is the [zeitgeist for Go engines](https://en.wikipedia.org/wiki/AlphaGo#History) like [AlphaGo](https://www.nature.com/articles/nature24270.epdf?author_access_token=VJXbVjaSHxFoctQQ4p2k4tRgN0jAjWel9jnR3ZoTv0PVW4gB86EEpGqTRDtpIz-2rmo8-KG06gqVobU5NSCFeHILHcVFUeMsbvwS-lxjqQGg98faovwjxeTUgZAUMnRQ), while heuristic-based is preferred for Chess engines like (my) [Hagnus Miemann](https://github.com/plettj/hagnusmiemann) or [Stockfish](https://github.com/official-stockfish/Stockfish).

I hope to implement _self-play_ learning in this Go engine.

## Notable Learnings

- Followed Go's official [tour](https://go.dev/tour/list) to learn the basics.
- Learned Go code structure from [these](https://go.dev/doc/modules/layout#server-project) [four](https://developer20.com/how-to-structure-go-code/) [layout](https://skife.org/golang/2013/03/24/go_dev_env.html) [standards](https://github.com/golang-standards/project-layout).

---

_Project approach is inspired by [Lucas](https://github.com/Strophox)'s [Tetrs in Rust](https://github.com/Strophox/tetrs?tab=readme-ov-file)._

[^1]: As stated in the [first sentence](https://en.wikipedia.org/wiki/AlphaGo#History) of AlphaGo's Wikipedia article.
