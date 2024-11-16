package utils

import (
	"encoding/base64"
	"log"
	"strconv"
	"time"

	"gongo/internal/config"
)

// Gongo's fully custom url-safe hashed Snowflake ID generation module.

// Generate a Gongo Snowflake ID
func GetSnowflake() string {
	machineId, err := strconv.ParseInt(config.GetEnv("MACHINE_ID"), 10, 64)
	if err != nil || machineId < 0 || machineId > 15 {
		log.Fatalf("Invalid MACHINE_ID environment variable. Must be between 0 and 15: %v", err)
	}

	epoch := int64(time.Since(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)).Milliseconds())
	if epoch < 0 || epoch > (1<<35)-1 {
		log.Fatalf("Epoch exceeds 35-bit limit.") // Needs fixing in year 2159.
	}

	const sequence int64 = 0 // TODO: Implement sequence counter via batching.

	// Create the Gongo 48-bit Snowflake ID in binary
	snowflake := rotate48bit((((epoch << 13) | (machineId << 9) | sequence) & 0xFFFFFFFFFFFF) ^ 14444999999999)

	return encode48bitToUrl(snowflake)
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

// Encodes the 48-bit ID in a URL-safe Base64 format
func encode48bitToUrl(id int64) string {
	bytes := make([]byte, 6)
	for i := 0; i < 6; i++ {
		bytes[i] = byte(id >> uint(40-i*8))
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
