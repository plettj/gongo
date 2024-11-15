package main

import (
	"log"

	"gongo/internal/utils"
)

type Game struct {
	ID string `json:"id"` // Gongo Snowflake ID
	Size int `json:"size"` // Go board size
	Board int `json:"board"` // Go board in binary
}

func main() {
	// router := api.NewRouter()
	// log.Println("Server is running on port 8080")
	log.Println("*Hello World!*")
	snowflake := utils.GetSnowflake()
	log.Println("Generated Snowflake ID:", snowflake)
	// log.Fatal(http.ListenAndServe(":8080", router))
}