package handlers

import (
	"bncapidemo/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetAssets(queryParams url.Values) (*models.Assets, error) {
	baseURL := "https://bravenewcoin.p.rapidapi.com/asset"

	u, _ := url.Parse(baseURL)
	u.RawQuery = queryParams.Encode()
	urlStr := u.String()
	fmt.Println(urlStr)

	req, _ := http.NewRequest("GET", urlStr, nil)

	req.Header.Add("X-RapidAPI-Key", "9ac629f999msh3885b4618d8e175p12ccd4jsn29191fc28479")
	req.Header.Add("X-RapidAPI-Host", "bravenewcoin.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make the API request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request returned status code: %d, body: %s", res.StatusCode, string(body))
	}

	var assetList models.Assets
	if err := json.Unmarshal(body, &assetList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
	}

	return &assetList, nil
}
