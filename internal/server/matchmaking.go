/*
All game matchmaking logic and connection processes.
*/
package server

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"

	"gongo/internal/engine"
	"gongo/internal/utils"
)

// MatchmakingResponse is the response sent to a player when a match is found
type MatchmakingResponse struct {
	GameID      string `json:"game_id"`
	PlayerToken string `json:"player_token"`
}

// Global variables
var (
	matchmakingQueues = make(map[int]chan chan MatchmakingResponse)
	games             = make(map[string]*Game) // In-memory game storage. TODO: Make persistent via PostgreSQL!!
	mutex             sync.Mutex
)

// Initialize matchmaking queues and start matchmaking goroutines
func init() {
	sizes := []int{9, 13, 19}

	for _, size := range sizes {
		matchmakingQueues[size] = make(chan chan MatchmakingResponse)
		go startMatchmaking(size, matchmakingQueues[size])
	}
}

// startMatchmaking matches players and creates games
func startMatchmaking(boardSize int, queue chan chan MatchmakingResponse) {
	var waitingPlayer chan MatchmakingResponse
	for {
		player := <-queue
		if waitingPlayer == nil {
			// No waiting player, set current player as waiting
			waitingPlayer = player
		} else {
			// Match found, create game
			gameId := utils.GetSnowflake()
			token1 := utils.GenerateToken()
			token2 := utils.GenerateToken()

			// Randomly assign colors
			colors := []string{"black", "white"}
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(colors), func(i, j int) { colors[i], colors[j] = colors[j], colors[i] })

			// Initialize game
			game := &Game{
				Id:      gameId,
				Board:   engine.NewBoard(byte(boardSize)),
				Players: [2]string{token1, token2},
				Moves:   []Move{},
			}

			// Store game
			mutex.Lock()
			games[gameId] = game
			mutex.Unlock()

			// Respond to both players
			match1 := MatchmakingResponse{GameID: gameId, PlayerToken: token1}
			match2 := MatchmakingResponse{GameID: gameId, PlayerToken: token2}
			waitingPlayer <- match1
			player <- match2

			// Reset waiting player
			waitingPlayer = nil
		}
	}
}

// RequestMatch handles POST /api/games for matchmaking
func RequestMatch(w http.ResponseWriter, r *http.Request) {
	var req struct {
		BoardSize int `json:"board_size"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate board size
	mutex.Lock()
	queue, ok := matchmakingQueues[req.BoardSize]
	mutex.Unlock()
	if !ok {
		http.Error(w, "Unsupported board size", http.StatusBadRequest)
		return
	}

	responseChan := make(chan MatchmakingResponse)

	// Add to matchmaking queue
	queue <- responseChan

	// Wait for match to be found
	match := <-responseChan

	// Respond with game details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(match)
}

// GetGameState handles GET /api/games/{gameId}
func GetGameState(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameId"]

	mutex.Lock()
	game, exists := games[gameID]
	mutex.Unlock()

	if !exists {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	// Return game state
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}
