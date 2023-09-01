package main

type DataRanks struct {
	Data []struct {
		Name           string  `_json:"name"`
		FullName       string  `_json:"fullname"`
		Code           string  `_json:"code"`
		MarketBalue    int     `_json:"market_value"`
		MarketValueUsd int     `_json:"market_value_usd"`
		Marketcap      int     `_json:"marketcap"`
		Turnoverrate   float32 `_json:"turnoverrate"`
	} `_json:"data"`
}
