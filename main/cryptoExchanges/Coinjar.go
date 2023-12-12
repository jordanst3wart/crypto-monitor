package CryptoExchanges

import (
	"encoding/json"
	"strconv"
)

type Coinjar struct {
	Last   string `json:"last"`
	High   string `json:"-"`
	Low    string `json:"-"`
	Volume string `json:"volume_24h"`
	Bid    string `json:"bid"`
	Ask    string `json:"ask"`
}

func (b Coinjar) RequestUpdate(name string, url string, ch chan CryptoData, currency string, crypto string) {
	// RequestUpdate(v.name, v.url, ch, v.currency, v.crypto)
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoData{name, b, err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoData{name, b, err, currency, crypto}
	}
}

func (b Coinjar) VolumeFloat() (float64, error) {
	return strconv.ParseFloat(b.Volume, 64)
}

func (b Coinjar) AskFloat() (float64, error) {
	return strconv.ParseFloat(b.Ask, 64)
}

func (b Coinjar) BidFloat() (float64, error) {
	return strconv.ParseFloat(b.Bid, 64)
}

func (b Coinjar) HighFloat() (float64, error) {
	return strconv.ParseFloat(b.High, 64)
}

func (b Coinjar) LowFloat() (float64, error) {
	return strconv.ParseFloat(b.Low, 64)
}

func (b Coinjar) LastFloat() (float64, error) {
	return strconv.ParseFloat(b.Last, 64)
}
