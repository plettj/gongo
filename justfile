set windows-shell := ["powershell.exe", "-c"]
# set dotenv-filename := ".env.local"

build:
  go build -o bin/server ./cmd/server

run:
  go run ./cmd/server/main.go