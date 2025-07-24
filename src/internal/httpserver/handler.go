package httpserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/osesantos/gomind/src/internal/model"
	"github.com/osesantos/gomind/src/internal/ollama"
)

func Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ask", askHandler)
	return mux
}

func askHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.HttpAskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Printf("💬 Received question: %s", req.Question)

	// For now we'll use the ollama package to process the question
	// In the future, This will call the brain package to process the question
	ollamaResp, err := ollama.Prompt(&http.Client{}, req.Question)
	if err != nil {
		log.Printf("❌ Error processing question: %v", err)
		http.Error(w, "Failed to process question", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Response from Ollama")

	resp := model.HttpAskReply{
		Answer: ollamaResp,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
