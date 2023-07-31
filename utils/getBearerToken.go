package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetToken() (string, error) {
	clientID := "oCdQoZoI96ERE9HY3sQ7JmbACfBf55RY"
	rapidAPIKey := "9ac629f999msh3885b4618d8e175p12ccd4jsn29191fc28479"

	url := "https://bravenewcoin.p.rapidapi.com/oauth/token"

	payload := strings.NewReader(`{
		"audience": "https://api.bravenewcoin.com",
		"client_id": "` + clientID + `",
		"grant_type": "client_credentials"
	}`)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", rapidAPIKey)
	req.Header.Add("X-RapidAPI-Host", "bravenewcoin.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make the API request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read API response: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request returned status code: %d, body: %s", res.StatusCode, string(body))
	}

	var response struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal API response: %v", err)
	}

	return response.AccessToken, nil
}
