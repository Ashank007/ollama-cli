package cmd

import "time"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type StreamChunk struct {
	Message Message `json:"message"`
	Done    bool    `json:"done"`
}

type ChatSession struct {
	Name      string    `json:"name"`
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	Timestamp time.Time `json:"timestamp"`
}


