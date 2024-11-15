package main

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"gongo/internal/api"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Game struct {
	ID string `json:"id"` // Gongo Snowflake ID
	Size int `json:"size"` // Go board size
	Board int `json:"board"` // Go board in binary
}

// use godot package to load/read the .env file and
// return the value of the key
func getEnv(key string) string {
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file.")
  }

  return os.Getenv(key)
}

// Generates a Gongo Snowflake ID based on Epoch, machineId, and a sequence number
func makeGameID() string {
	machineId, err := strconv.ParseInt(getEnv("MACHINE_ID"), 10, 64)
	if err != nil {
    panic(err)
	}

	// Get the current time in seconds since 2024
	epoch := int64(time.Since(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC))) / 1000000000
	const sequence int64 = 0 // TODO: Have the sequence counter increment on the server

	var snowflake int64 = (epoch << 22) | (machineId << 12) | sequence

	str := base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%08x", snowflake)))

	log.Println("Long String: ", str)

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(snowflake))

	testerStr := base64.StdEncoding.EncodeToString(b)
	log.Println("Tester String: ", testerStr)

	return str[:8]
}

func main() {
    router := api.NewRouter()
    log.Println("Server is running on port 8080")
    log.Println("*Hello World!*")
		str := makeGameID()
		log.Println(str)
    log.Fatal(http.ListenAndServe(":8080", router))
}