package main

import (
	"gongo/internal/api"
	"log"
	"net/http"
)

func main() {
	router := api.NewRouter()
	log.Println("Server is running on port 8080.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
