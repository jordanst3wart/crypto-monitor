package CryptoExchanges

import (
	"encoding/json"
	"strconv"
)

// flatten struct
// https://stackoverflow.com/questions/24642575/go-golang-flatten-a-nested-json
type GeminiTickerBTC struct {
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	BTCVolume `json:"volume"`
	Last      string `json:"last"`
	High      string `json:"-"`
	Low       string `json:"-"`
}

type BTCVolume struct {
	Value     string `json:"BTC"`
	Volume    string `json:"USD"`
	Timestamp int    `json:"timestamp"`
}

func (b GeminiTickerBTC) RequestUpdate(name string, url string, ch chan CryptoData, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoData{name, b, err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoData{name, b, err, currency, crypto}
	}
}

func (b GeminiTickerBTC) VolumeFloat() (float64, error) {
	return strconv.ParseFloat(b.Volume, 64)
}

func (b GeminiTickerBTC) AskFloat() (float64, error) {
	return strconv.ParseFloat(b.Ask, 64)
}

func (b GeminiTickerBTC) BidFloat() (float64, error) {
	return strconv.ParseFloat(b.Bid, 64)
}

func (b GeminiTickerBTC) HighFloat() (float64, error) {
	return strconv.ParseFloat(b.High, 64)
}

func (b GeminiTickerBTC) LowFloat() (float64, error) {
	return strconv.ParseFloat(b.Low, 64)
}

func (b GeminiTickerBTC) LastFloat() (float64, error) {
	return strconv.ParseFloat(b.Last, 64)
}
