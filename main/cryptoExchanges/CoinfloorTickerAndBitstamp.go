package CryptoExchanges

import (
	"encoding/json"
	"strconv"
)

type CoinfloorTickerAndBitstamp struct {
	Last   string `json:"last"`
	High   string `json:"high"`
	Low    string `json:"low"`
	Volume string `json:"volume"`
	Bid    string `json:"bid"`
	Ask    string `json:"ask"`
	Vwap   string `json:"vwap"`
}

func (b CoinfloorTickerAndBitstamp) RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b, err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b, err, currency, crypto}
	}
}

func (b CoinfloorTickerAndBitstamp) VolumeFloat() (float64, error) {
	return strconv.ParseFloat(b.Volume, 64)
}

func (b CoinfloorTickerAndBitstamp) AskFloat() (float64, error) {
	return strconv.ParseFloat(b.Ask, 64)
}

func (b CoinfloorTickerAndBitstamp) BidFloat() (float64, error) {
	return strconv.ParseFloat(b.Bid, 64)
}

func (b CoinfloorTickerAndBitstamp) HighFloat() (float64, error) {
	return strconv.ParseFloat(b.High, 64)
}

func (b CoinfloorTickerAndBitstamp) LowFloat() (float64, error) {
	return strconv.ParseFloat(b.Low, 64)
}

func (b CoinfloorTickerAndBitstamp) LastFloat() (float64, error) {
	return strconv.ParseFloat(b.Last, 64)
}
