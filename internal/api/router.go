package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"gongo/internal/engine"
)

// Set up the API routes and return a mux router.
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/game", getGame).Methods("GET") // Define the GET /game endpoint
	return router
}

// GET /game endpoint
func getGame(w http.ResponseWriter, r *http.Request) {
	game := engine.NewGame()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(game); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
