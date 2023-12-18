package CryptoExchanges

import (
	"encoding/json"
	"strconv"
)

// https://api.binance.com/api/v3/ticker/24hr?symbol=BTCEUR
/*
{"symbol":"ETHEUR",
"priceChange":"-45.53000000",
"priceChangePercent":"-2.234",
"weightedAvgPrice":"2024.62838819",
"prevClosePrice":"2037.95000000",
"lastPrice":"1992.11000000",
"lastQty":"0.09160000",
"bidPrice":"1991.91000000",
"bidQty":"0.60000000",
"askPrice":"1991.96000000",
"askQty":"0.01380000",
"openPrice":"2037.64000000",
"highPrice":"2051.45000000",
"lowPrice":"1983.40000000",
"volume":"2100.83030000",
"quoteVolume":"4253400.66414800",
"openTime":1702791048715,
"closeTime":1702877448715,
"firstId":91788110,
"lastId":91799704,
"count":11595
}

*/

type Binance struct {
	Last   string `json:"lastPrice"`
	High   string `json:"highPrice"`
	Low    string `json:"lowPrice"`
	Volume string `json:"volume"`
	Bid    string `json:"bidPrice"`
	Ask    string `json:"askPrice"`
}

func (b Binance) RequestUpdate(name string, url string, ch chan CryptoData, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoData{name, b, err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoData{name, b, err, currency, crypto}
	}
}

func (b Binance) VolumeFloat() (float64, error) {
	return strconv.ParseFloat(b.Volume, 64)
}

func (b Binance) AskFloat() (float64, error) {
	return strconv.ParseFloat(b.Ask, 64)
}

func (b Binance) BidFloat() (float64, error) {
	return strconv.ParseFloat(b.Bid, 64)
}

func (b Binance) HighFloat() (float64, error) {
	return strconv.ParseFloat(b.High, 64)
}

func (b Binance) LowFloat() (float64, error) {
	return strconv.ParseFloat(b.Low, 64)
}

func (b Binance) LastFloat() (float64, error) {
	return strconv.ParseFloat(b.Last, 64)
}
