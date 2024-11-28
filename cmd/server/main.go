package main

import (
	"gongo/internal/server"
	"log"
	"net/http"
)

func main() {
	router := server.NewRouter()
	log.Println("Server is running on port 8080.")
	log.Fatal(http.ListenAndServe(":8080", router))

	// How to start my new router system.
	// apiRouter := router.PathPrefix("/api").Subrouter()
	// apiRouter.HandleFunc("/games", RequestMatch).Methods("POST")
	// apiRouter.HandleFunc("/games/{gameId}", GetGameState).Methods("GET")
}
