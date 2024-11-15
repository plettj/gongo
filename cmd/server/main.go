package main

import (
	"log"
	"net/http"

	"gongo/internal/api"
)

func main() {
	router := api.NewRouter()
	log.Println("Server is running on port 8080.")
	log.Fatal(http.ListenAndServe(":8080", router))
}