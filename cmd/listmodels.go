package cmd

import (
	"net/http"
  "encoding/json"
)

func listModels() ([]string, error) {
	resp, err := http.Get("http://localhost:11434/api/tags")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Models []struct {
			Name string `json:"name"`
		} `json:"models"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	models := []string{}
	for _, m := range result.Models {
		models = append(models, m.Name)
	}
	return models, nil
}
