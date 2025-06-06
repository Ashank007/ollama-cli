package cmd

import (
	"fmt"
)

func PickModel() (string, error) {
	fmt.Println("📦 Fetching available models...")
	models, err := listModels()
	if err != nil {
		return "", fmt.Errorf("could not fetch models: %v", err)
	}
	if len(models) == 0 {
		return "", fmt.Errorf("⚠️ No models found. Run `ollama pull <model>` first.")
	}

	fmt.Println("🧠 Available models:")
	for i, m := range models {
		fmt.Printf("[%d] %s\n", i+1, m)
	}

	fmt.Print("Pick a model number: ")
	var choice int
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(models) {
		return "", fmt.Errorf("invalid choice")
	}

	return models[choice-1], nil
}

