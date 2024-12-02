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

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	// TODO: Remove the /game endpoint and fully replace it with the /games endpoint.
	apiRouter.HandleFunc("/game", getGame).Methods("GET")

	apiRouter.HandleFunc("/games", RequestMatch).Methods("POST")
	apiRouter.HandleFunc("/games/{gameId}", GetGameState).Methods("GET")
	return apiRouter
}

// GET /game endpoint
func getGame(w http.ResponseWriter, r *http.Request) {
	board := engine.NewBoard(19)

	game := Game{
		Board: board.Serialize(),
		Id:    utils.GetSnowflake(),
		// TODO: Replace the testing values below with actually relevant data.
		Players: [2]string{"blackUsername", "whiteUsername"},
		Moves:   []Move{{X: 1, Y: 2, Color: 1}, {X: 3, Y: 4, Color: 2}, {X: 5, Y: 5, Color: 0}, {X: 9, Y: 9, Color: 2}},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(game); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
