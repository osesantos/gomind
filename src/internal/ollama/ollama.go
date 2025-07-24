package ollama

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/osesantos/gomind/src/internal/model"
)

// TODO: Change the these const to the actual URL of your Ollama server inside the cluster, use Environment variables to set it.
const (
	// MODEL is the name of the model to use for Ollama.
	MODEL = "mistral"
	// STREAM indicates whether the response should be streamed or not.
	STREAM = false
	// URL is the URL of the Ollama API endpoint.
	URL = "http://localhost:11434/api/generate"
)

// This method is used to prompt the Ollama LLM with a question.
// e.g. of a curl request:
// curl -s http://localhost:11434/api/generate \
// -d '{
// "model": "mistral",
// "prompt": "If Alice is older than Bob, and Bob is older than Charlie, who is the oldest?",
// "stream": false
// }
func Prompt(client *http.Client, question string) (string, error) {
	jsonData, err := json.Marshal(model.OllamaRequest{
		Model:  MODEL,
		Prompt: question,
		Stream: STREAM,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer req.Body.Close()

	var ollamaReply model.OllamaReply
	if err := json.NewDecoder(resp.Body).Decode(&ollamaReply); err != nil {
		return "", err
	}

	return ollamaReply.Response, nil
}
