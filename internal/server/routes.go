package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"poker-sim/internal/card"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HelloWorldHandler)
	mux.HandleFunc("/deck", s.drawCardHandler)
	mux.HandleFunc("/health", s.healthHandler)
	mux.HandleFunc("/test", s.testCard)

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

	_, err = w.Write(jsonResp)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) drawCardHandler(w http.ResponseWriter, r *http.Request) {
	deck := card.NewDeck()
	deck.Shuffle()

	jsonResp, err := json.Marshal(deck.Cards)
	if err != nil {
		log.Fatalf("Error marshaling card to JSON. Err: %v", err)
		fmt.Printf("Error marshaling card to JSON. Err: %v", err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	_, err = w.Write(jsonResp)
	if err != nil {
		fmt.Println(err)

	}
}

func (s *Server) testCard(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 5000; i++ {
		s.drawCardHandler(w, r)
	}
}
