package CryptoExchanges

import (
	"encoding/json"
	"strconv"
)

type Bitstamp struct {
	Last   string `json:"last"`
	High   string `json:"high"`
	Low    string `json:"low"`
	Volume string `json:"volume"`
	Bid    string `json:"bid"`
	Ask    string `json:"ask"`
	Vwap   string `json:"vwap"`
}

func (b Bitstamp) RequestUpdate(name string, url string, ch chan CryptoData, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoData{name, b, err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoData{name, b, err, currency, crypto}
	}
}

func (b Bitstamp) VolumeFloat() (float64, error) {
	return strconv.ParseFloat(b.Volume, 64)
}

func (b Bitstamp) AskFloat() (float64, error) {
	return strconv.ParseFloat(b.Ask, 64)
}

func (b Bitstamp) BidFloat() (float64, error) {
	return strconv.ParseFloat(b.Bid, 64)
}

func (b Bitstamp) HighFloat() (float64, error) {
	return strconv.ParseFloat(b.High, 64)
}

func (b Bitstamp) LowFloat() (float64, error) {
	return strconv.ParseFloat(b.Low, 64)
}

func (b Bitstamp) LastFloat() (float64, error) {
	return strconv.ParseFloat(b.Last, 64)
}
