package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Hardcoded quotes for different moods
var quotes = map[string]string{
	"sad":     "Tears come from the heart and not from the brain. — Leonardo da Vinci",
	"happy":   "Happiness is not something ready made. It comes from your own actions. — Dalai Lama",
	"tired":   "Rest when you're weary. Refresh and renew yourself. — Ralph Marston",
	"excited": "The excitement of learning separates youth from old age. — Ezra Taft Benson",
	"angry":   "Speak when you are angry and you will make the best speech you will ever regret. — Ambrose Bierce",
}

// Structs for parsing incoming JSON
type MCPRequest struct {
	Tool  string `json:"tool"`
	Input struct {
		Mood string `json:"mood"`
	} `json:"input"`
}

type MCPResponse struct {
	Output struct {
		Quote string `json:"quote"`
	} `json:"output"`
}

// Handle /messages/ endpoint
func handleSSE(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Pick a default mood or a query param
	mood := r.URL.Query().Get("mood")
	if mood == "" {
		mood = "happy"
	}
	quote, ok := quotes[mood]
	if !ok {
		quote = "Sorry, I don't have a quote for that mood yet."
	}

	// Build response
	resp := MCPResponse{}
	resp.Output.Quote = quote

	respBytes, _ := json.Marshal(resp)
	fmt.Fprintf(w, "data: %s\n\n", respBytes)
}

func main() {
	http.HandleFunc("/messages/", handleSSE)
	log.Println("✅ Mood Quote Server running at http://localhost:8080/messages/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
