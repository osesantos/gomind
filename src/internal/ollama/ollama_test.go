package ollama_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/osesantos/gomind/src/internal/ollama"
)

type MockRoundTripper struct {
	roundTripFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.roundTripFunc != nil {
		return m.roundTripFunc(req)
	}
	return nil, nil // Default behavior if no function is set
}

func TestPrompt(t *testing.T) {
	mockRoundTripper := &MockRoundTripper{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			// Mock response for the Ollama API
			response := `{"response": "Alice is the oldest."}`
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(response)),
				Header:     make(http.Header),
			}, nil
		},
	}

	client := &http.Client{Transport: mockRoundTripper}
	answer, err := ollama.Prompt(client, "If Alice is older than Bob, and Bob is older than Charlie, who is the oldest?")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedAnswer := "Alice is the oldest."
	if answer != expectedAnswer {
		t.Errorf("Expected answer %q, got %q", expectedAnswer, answer)
	}
}

func TestPromptErrorWithBadJson(t *testing.T) {
	mockRoundTripper := &MockRoundTripper{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			// Mock response with bad JSON
			response := `{"response": "Alice is the oldest."` // Missing closing brace
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(response)),
				Header:     make(http.Header),
			}, nil
		},
	}

	client := &http.Client{Transport: mockRoundTripper}
	_, err := ollama.Prompt(client, "If Alice is older than Bob, and Bob is older than Charlie, who is the oldest?")
	if err == nil {
		t.Fatal("Expected an error due to bad JSON, but got none")
	}
}

func TestPromptErrorWithBadResponse(t *testing.T) {
	mockRoundTripper := &MockRoundTripper{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			// Mock response with a non-200 status code
			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(strings.NewReader("Internal Server Error")),
				Header:     make(http.Header),
			}, nil
		},
	}

	client := &http.Client{Transport: mockRoundTripper}
	_, err := ollama.Prompt(client, "If Alice is older than Bob, and Bob is older than Charlie, who is the oldest?")
	if err == nil {
		t.Fatal("Expected an error due to bad response, but got none")
	}
}
