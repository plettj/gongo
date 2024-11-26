set windows-shell := ["powershell.exe", "-c"]
# set dotenv-filename := ".env.local"

build-server:
  go build -o bin/server ./cmd/server

run:
  go run ./cmd/server/main.go

tui:
  go run ./cmd/tui/main.go