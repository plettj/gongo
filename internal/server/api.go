/*
All API endpoints for the game server.
*/
package server

import (
	"encoding/json"
	"net/http"

	"gongo/internal/engine"
	"gongo/internal/utils"

	"github.com/gorilla/mux"
)

// Set up the API routes and return a mux router.
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/game", getGame).Methods("GET")
	return router
}

// GET /game endpoint
func getGame(w http.ResponseWriter, r *http.Request) {
	board := engine.NewBoard(19)

	game := Game{
		Board: board,
		Id:    utils.GetSnowflake(),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(game); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
