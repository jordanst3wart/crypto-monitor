package CryptoExchanges

import (
	"encoding/json"
	"strconv"
)

type GeminiTickerETH struct {
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	ETHVolume `json:"volume"`
	Last      string `json:"last"`
	High      string `json:"-"`
	Low       string `json:"-"`
}

type ETHVolume struct {
	Value     string `json:"ETH"`
	Volume    string `json:"USD"`
	Timestamp int    `json:"timestamp"`
}

func (b GeminiTickerETH) RequestUpdate(name string, url string, ch chan CryptoData, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoData{name, b, err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoData{name, b, err, currency, crypto}
	}
}

func (b GeminiTickerETH) VolumeFloat() (float64, error) {
	return strconv.ParseFloat(b.Volume, 64)
}

func (b GeminiTickerETH) AskFloat() (float64, error) {
	return strconv.ParseFloat(b.Ask, 64)
}

func (b GeminiTickerETH) BidFloat() (float64, error) {
	return strconv.ParseFloat(b.Bid, 64)
}

func (b GeminiTickerETH) HighFloat() (float64, error) {
	return strconv.ParseFloat(b.High, 64)
}

func (b GeminiTickerETH) LowFloat() (float64, error) {
	return strconv.ParseFloat(b.Low, 64)
}

func (b GeminiTickerETH) LastFloat() (float64, error) {
	return strconv.ParseFloat(b.Last, 64)
}
