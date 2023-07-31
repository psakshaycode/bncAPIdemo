package handlers

import (
	"bncapidemo/models"
	"bncapidemo/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetAssetTicker(queryParams url.Values) (*models.AssetTickerResp, error) {
	baseURL := "https://bravenewcoin.p.rapidapi.com/market-cap"

	u, _ := url.Parse(baseURL)
	u.RawQuery = queryParams.Encode()
	urlStr := u.String()
	fmt.Println(urlStr)

	req, _ := http.NewRequest("GET", urlStr, nil)
	accessToken, err := utils.GetToken()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)
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
	var tickRes models.AssetTickerResp
	json.Unmarshal(body, &tickRes)

	return &tickRes, nil
}
