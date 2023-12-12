package CryptoExchanges

import (
	"encoding/json"
	"strconv"
)

type ACXTicker struct {
	ACXNested `json:"ticker"`
}

type ACXNested struct {
	Bid    string `json:"buy"`
	Ask    string `json:"sell"`
	Volume string `json:"vol"`
	Last   string `json:"last"`
	High   string `json:"high"`
	Low    string `json:"low"`
}

func (b ACXTicker) RequestUpdate(name string, url string, ch chan CryptoData, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoData{name, b, err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoData{name, b, err, currency, crypto}
	}
}

func (b ACXTicker) VolumeFloat() (float64, error) {
	return strconv.ParseFloat(b.Volume, 64)
}

func (b ACXTicker) AskFloat() (float64, error) {
	return strconv.ParseFloat(b.Ask, 64)
}

func (b ACXTicker) BidFloat() (float64, error) {
	return strconv.ParseFloat(b.Bid, 64)
}

func (b ACXTicker) HighFloat() (float64, error) {
	return strconv.ParseFloat(b.High, 64)
}

func (b ACXTicker) LowFloat() (float64, error) {
	return strconv.ParseFloat(b.Low, 64)
}

func (b ACXTicker) LastFloat() (float64, error) {
	return strconv.ParseFloat(b.Last, 64)
}
