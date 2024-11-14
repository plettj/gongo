# Developer Notes

## Engine to-dos

- [x] Decide on an API protocol (gRPC, REST, GraphQL) we'll use, and router (gorilla/mux, gin, go) we'll use. Decisions: REST + gorilla/mux.
- [ ] Set up and test a first simple API request and response example to mimic the creation of a game.
- [ ] Create a bare-bones Go viewer in a nice CLI. Example: [Tetrgo](https://github.com/Broderick-Westrope/tetrigo?tab=readme-ov-file) built with [BubbleTea](https://github.com/charmbracelet/bubbletea?tab=readme-ov-file).
- [ ] Be able to load Go games into the viewer based on a standard game representation.
- [ ] Research the computationally fastest way to represent a large Go game's board, and implement it.
- [ ] Research and decide on the various rules of Go I'll be supporting.
- [ ] Write a spec for the rules my engine supports.
- [ ] Research the computationally fastest way people have made move generators for Go.
- [ ] (large) Implement a Go move generator.

## Usage

Run a go program (short for `go run <filepath>`):

```bash
just run
```
