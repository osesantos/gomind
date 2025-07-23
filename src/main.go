package main

import (
	"log"
	"net/http"

	"github.com/osesantos/gomind/src/internal/httpserver"
)

func main() {
	// Start HTTP server
	log.Println("🔁 Running GoMind MCP on :4433")

	if err := http.ListenAndServe(":4433", httpserver.Router()); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
