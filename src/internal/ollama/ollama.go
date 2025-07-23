package ollama

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/osesantos/gomind/src/internal/model"
)

const (
	MODEL  = "mistral"
	STREAM = false
	// TODO: Change this to the actual URL of your Ollama server inside the cluster
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
func Prompt(question string) (string, error) {
	jsonData, err := json.Marshal(model.OllamaRequest{
		Model:  MODEL,
		Prompt: question,
		Stream: STREAM,
	})
	if err != nil {
		return "", err
	}

	client := http.Client{}
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
