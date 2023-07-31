package models

type Asset struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Symbol string `json:"symbol"`
	Type   string `json:"type"`
	Url    string `json:"url"`
}

type Assets struct {
	Content []Asset `json:"content"`
}

type AssetTickerResp interface {
}
