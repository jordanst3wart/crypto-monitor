package CryptoExchanges

import (
	"encoding/json"
)

type IndependentReserve struct {
	Last   float64 `json:"LastPrice"`
	High   float64 `json:"DayHighestPrice"`
	Low    float64 `json:"DayLowestPrice"`
	Volume float64 `json:"DayVolumeXbt"`
	Bid    float64 `json:"CurrentHighestBidPrice"`
	Ask    float64 `json:"CurrentLowestOfferPrice"`
}

func (b IndependentReserve) RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b, err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b, err, currency, crypto}
	}
}

func (b IndependentReserve) VolumeFloat() (float64, error) {
	return b.Volume, nil
}

func (b IndependentReserve) AskFloat() (float64, error) {
	return b.Ask, nil
}

func (b IndependentReserve) BidFloat() (float64, error) {
	return b.Bid, nil
}

func (b IndependentReserve) HighFloat() (float64, error) {
	return b.High, nil
}

func (b IndependentReserve) LowFloat() (float64, error) {
	return b.Low, nil
}

func (b IndependentReserve) LastFloat() (float64, error) {
	return b.Last, nil
}
