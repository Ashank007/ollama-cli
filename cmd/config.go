package cmd

import (
	"os"
)

const configFile = ".ollama_config"

func GetAPIKey() (string, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func SaveAPIKey(key string) error {
	return os.WriteFile(configFile, []byte(key), 0600)
}

