package server

import (
	"encoding/json"
	"log"
	"net/http"
	"poker-sim/internal/card"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HelloWorldHandler)
	mux.HandleFunc("/deck", s.drawCardHandler)
	mux.HandleFunc("/health", s.healthHandler)

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) drawCardHandler(w http.ResponseWriter, r *http.Request) {
	deck := card.NewDeck()
	deck.Shuffle()

	jsonResp, err := json.Marshal(deck.Cards)
	if err != nil {
		log.Fatalf("Error marshaling card to JSON. Err: %v", err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Replace with your Vite dev server's port
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	_, _ = w.Write(jsonResp)
}
