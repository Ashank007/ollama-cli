package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendChatRequest(model string, history []Message) (string, error) {
	payload := ChatRequest{
		Model:    model,
		Messages: history,
		Stream:   true,
	}

	body, _ := json.Marshal(payload)
	resp, err := http.Post("http://localhost:11434/api/chat", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	var fullReply string
	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var chunk StreamChunk
		err := json.Unmarshal([]byte(line), &chunk)
		if err != nil {
			return "", fmt.Errorf("error decoding chunk: %v", err)
		}

		fmt.Print(chunk.Message.Content)
		fullReply += chunk.Message.Content

		if chunk.Done {
			break
		}
	}
	fmt.Println()
	return fullReply, nil
}

