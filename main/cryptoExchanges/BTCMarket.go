package CryptoExchanges

import "encoding/json"

type BTCMarket struct {
	Last   float64 `json:"lastPrice"`
	High   float64 `json:"-"`
	Low    float64 `json:"-"`
	Volume float64 `json:"volume24h"`
	Bid    float64 `json:"bestBid"`
	Ask    float64 `json:"bestAsk"`
}

func (b BTCMarket) VolumeFloat() (float64, error) {
	return b.Volume, nil
}

func (b BTCMarket) AskFloat() (float64, error) {
	return b.Ask, nil
}

func (b BTCMarket) BidFloat() (float64, error) {
	return b.Bid, nil
}

func (b BTCMarket) HighFloat() (float64, error) {
	return b.High, nil
}

func (b BTCMarket) LowFloat() (float64, error) {
	return b.Low, nil
}

func (b BTCMarket) LastFloat() (float64, error) {
	return b.Last, nil
}

func (b BTCMarket) RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b, err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b, err, currency, crypto}
	}
}
