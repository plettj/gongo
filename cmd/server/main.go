package main

import (
	"encoding/base64"
	"log"
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

// Load environment variables from .env file
func getEnv(key string) string {
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file.")
  }

  return os.Getenv(key)
}

// getSnowflake generates a Gongo Snowflake ID with the following structure:
// - bits 0-8: Sequence counter
// - bits 9-12: Machine ID
// - bits 13-47: Epoch time in seconds since the start of 2024
func getSnowflake() string {
	machineId, err := strconv.ParseInt(getEnv("MACHINE_ID"), 10, 64)
	if err != nil || machineId < 0 || machineId > 15 {
		log.Fatalf("Invalid MACHINE_ID environment variable. Must be between 0 and 15: %v", err)
	}

	epoch := int64(time.Since(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)).Milliseconds())
	if epoch < 0 || epoch > (1<<35)-1 {
		log.Fatalf("Epoch exceeds 35-bit limit.") // Needs fixing in year 2159.
	}

	// TODO: Implement sequence counter via batching.
	const sequence int64 = 0

	hash, err := strconv.ParseInt(getEnv("SNOWFLAKE_HASH"), 16, 64)
	if err != nil {
		log.Fatalf("Invalid SNOWFLAKE_HASH environment variable: %v", err)
	}

	// Create the Gongo 48-bit Snowflake ID in binary
	snowflake := rotate48bit((((epoch << 13) | (machineId << 9) | sequence) & 0xFFFFFFFFFFFF) ^ hash)

	return encode48bitToUrl(snowflake ^ hash)
}

// Rotates a 48-bit number as though it were a 6x8 matrix
func rotate48bit(snowflake int64) int64 {
	var reordered int64

	for i := 0; i < 6; i++ {
		g := (snowflake >> (i * 8)) & 0xFF
		var expanded int64
		for j := 0; j < 8; j++ {
			bit := (g >> j) & 1
			expanded |= int64(bit) << (j * 6)
		}
		reordered |= expanded << i
	}

	return reordered
}

func encode48bitToUrl(id int64) string {
	bytes := make([]byte, 6)
	for i := 0; i < 6; i++ {
		bytes[i] = byte(id >> uint(40 - i * 8))
	}

	return base64.URLEncoding.EncodeToString(bytes)
}

func main() {
    // router := api.NewRouter()
    // log.Println("Server is running on port 8080")
    log.Println("*Hello World!*")
		str := getSnowflake()
		log.Println(str)
    // log.Fatal(http.ListenAndServe(":8080", router))
}