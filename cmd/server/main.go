package main

import (
	"gongo/internal/api"
	"log"
	"net/http"
)

type Game struct {
	ID int `json:"id"` // Gongo Snowflake ID
	Size int `json:"size"` // Go board size
	Board int `json:"board"` // Go board in binary
}

func main() {
    router := api.NewRouter()
    log.Println("Server is running on port 8080")
    log.Println("*Hello World!*")
    log.Fatal(http.ListenAndServe(":8080", router))
}