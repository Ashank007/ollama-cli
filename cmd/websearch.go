package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SearchResult struct {
	OrganicResults []struct {
		Title string `json:"title"`
		Snippet string `json:"snippet"`
		Link string `json:"link"`
	} `json:"organic_results"`
}

func SearchWeb(query string) (string, error) {
	apiKey, err := GetAPIKey()
	if err != nil {
		return "", fmt.Errorf("🔑 API key not found. Use 'set-key' command to store your SerpAPI key.")
	}

	params := url.Values{}
	params.Add("q", query)
	params.Add("api_key", apiKey)
	params.Add("engine", "google")

	endpoint := "https://serpapi.com/search.json?" + params.Encode()
	resp, err := http.Get(endpoint)
	if err != nil {
		return "", fmt.Errorf("failed to perform search: %v", err)
	}
	defer resp.Body.Close()

	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding result: %v", err)
	}

	var summary string
	for i, r := range result.OrganicResults {
		if i >= 3 {
			break
		}
		summary += fmt.Sprintf("🔗 %s\n%s\n%s\n\n", r.Title, r.Snippet, r.Link)
	}
	return summary, nil
}

