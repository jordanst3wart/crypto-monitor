package structs

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

/*
	bid_aud
	offer_aud
*/

// the `json:""` are called struct tags
// by default the json is parsed input the matching key
// currency struct
type CurrencyExchangeAPI struct {
	Base string            `json:"base"`
	Rates CurrencyExchange `json:"rates"`
	Date string            `json:"date"`
}

type CurrencyExchange struct {
	GDP float64 `json:"GBP"`
	AUD float64 `json:"AUD"`
	EUR float64 `json:"EUR"`
	USD float64 `json:"USD"`
}

type CryptoDTO struct {
	Name  string
	Coin  CryptoExchange
	Error error
	Currency string
	Crypto string
}
//

// flatten struct
// https://stackoverflow.com/questions/24642575/go-golang-flatten-a-nested-json
type GeminiTickerBTC struct {
	Bid string            `json:"bid"`
	Ask string            `json:"ask"`
	BTCVolume	          `json:"volume"`
	Last string           `json:"last"`
	High   string         `json:"-"`
	Low    string         `json:"-"`
}

type BTCVolume struct {
	Value  string    `json:"BTC"`
	Volume    string    `json:"USD"`
	Timestamp int    `json:"timestamp"`
}

type GeminiTickerETH struct {
	Bid string            `json:"bid"`
	Ask string            `json:"ask"`
	ETHVolume	          `json:"volume"`
	Last string           `json:"last"`
	High   string         `json:"-"`
	Low    string         `json:"-"`
}

type ETHVolume struct {
	Value    string    `json:"ETH"`
	Volume   string    `json:"USD"`
	Timestamp int      `json:"timestamp"`
}

type ACXTicker struct {
	ACXNested `json:"ticker"`
}

type ACXNested struct {
	Bid string            `json:"buy"`
	Ask string            `json:"sell"`
	Volume string	      `json:"vol"`
	Last string           `json:"last"`
	High   string         `json:"high"`
	Low    string         `json:"low"`
}

type CoinfloorTickerAndBitstamp struct {
	Last   string   `json:"last"`
	High   string   `json:"high"`
	Low    string   `json:"low"`
	Volume string   `json:"volume"`
	Bid    string   `json:"bid"`
	Ask    string   `json:"ask"`
	Vwap   string   `json:"vwap"`
}

type IndepentReserve struct {
	Last   float64   `json:"LastPrice"`
	High   float64   `json:"DayHighestPrice"`
	Low    float64   `json:"DayLowestPrice"`
	Volume float64   `json:"DayVolumeXbt"`
	Bid    float64   `json:"CurrentHighestBidPrice"`
	Ask    float64   `json:"CurrentLowestOfferPrice"`
}

type BTCMarket struct {
	Last   float64   `json:"lastPrice"`
	High   float64   `json:"-"`
	Low    float64   `json:"-"`
	Volume float64   `json:"volume24h"`
	Bid    float64   `json:"bestBid"`
	Ask    float64   `json:"bestAsk"`
}

type Coinjar struct {
	Last   string    `json:"last"`
	High   string    `json:"-"`
	Low    string    `json:"-"`
	Volume string    `json:"volume_24h"`
	Bid    string    `json:"bid"`
	Ask    string    `json:"ask"`
}

// TODO put logic in here for RequestUpdate to reduce repeating code
func requestWrapper(url string) ([]byte, error) {
	var responseData []byte
	resp, err := http.Get(url)
	if err != nil {
		return responseData, err
	}

	return ioutil.ReadAll(resp.Body)
}

func (b BTCMarket) VolumeFloat() string {
	return strconv.FormatFloat(b.Volume, 'f', -1, 64)
}

func (b BTCMarket) AskFloat() string {
	return strconv.FormatFloat(b.Ask, 'f', -1, 64)
}

func (b BTCMarket) BidFloat() string {
	return strconv.FormatFloat(b.Bid, 'f', -1, 64)
}

func (b BTCMarket) HighFloat() string {
	return strconv.FormatFloat(b.High, 'f', -1, 64)
}

func (b BTCMarket) LowFloat() string {
	return strconv.FormatFloat(b.Low, 'f', -1, 64)
}

func (b BTCMarket) LastFloat() string {
	return strconv.FormatFloat(b.Last, 'f', -1, 64)
}

func (b BTCMarket) RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b,err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b,err, currency, crypto}
	}
}

func (b IndepentReserve) RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b,err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b,err, currency, crypto}
	}
}

func (b IndepentReserve) VolumeFloat() string {
	return strconv.FormatFloat(b.Volume, 'f', -1, 64)
}

func (b IndepentReserve) AskFloat() string {
	return strconv.FormatFloat(b.Ask, 'f', -1, 64)
}

func (b IndepentReserve) BidFloat() string {
	return strconv.FormatFloat(b.Bid, 'f', -1, 64)
}

func (b IndepentReserve) HighFloat() string {
	return strconv.FormatFloat(b.High, 'f', -1, 64)
}

func (b IndepentReserve) LowFloat() string {
	return strconv.FormatFloat(b.Low, 'f', -1, 64)
}

func (b IndepentReserve) LastFloat() string {
	return strconv.FormatFloat(b.Last, 'f', -1, 64)
}

func (b GeminiTickerETH) RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b,err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b,err, currency, crypto}
	}
}

func (b GeminiTickerETH) VolumeFloat() string {
	return b.Volume
}

func (b GeminiTickerETH) AskFloat() string {
	return b.Ask
}

func (b GeminiTickerETH) BidFloat() string {
	return b.Bid
}

func (b GeminiTickerETH) HighFloat() string {
	return b.High
}

func (b GeminiTickerETH) LowFloat() string {
	return b.Low
}

func (b GeminiTickerETH) LastFloat() string {
	return b.Last
}

func (b GeminiTickerBTC) RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b,err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b,err, currency, crypto}
	}
}

func (b GeminiTickerBTC) VolumeFloat() string {
	return b.Volume
}

func (b GeminiTickerBTC) AskFloat() string {
	return b.Ask
}

func (b GeminiTickerBTC) BidFloat() string {
	return b.Bid
}

func (b GeminiTickerBTC) HighFloat() string {
	return b.High
}

func (b GeminiTickerBTC) LowFloat() string {
	return b.Low
}

func (b GeminiTickerBTC) LastFloat() string {
	return b.Last
}

func (b CoinfloorTickerAndBitstamp) RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b,err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b,err, currency, crypto}
	}
}

func (b CoinfloorTickerAndBitstamp) VolumeFloat() string {
	return b.Volume
}

func (b CoinfloorTickerAndBitstamp) AskFloat() string {
	return b.Ask
}

func (b CoinfloorTickerAndBitstamp) BidFloat() string {
	return b.Bid
}

func (b CoinfloorTickerAndBitstamp) HighFloat() string {
	return b.High
}

func (b CoinfloorTickerAndBitstamp) LowFloat() string {
	return b.Low
}

func (b CoinfloorTickerAndBitstamp) LastFloat() string {
	return b.Last
}

func (b ACXTicker) RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b,err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b,err, currency, crypto}
	}
}

func (b ACXTicker) VolumeFloat() string {
	return b.Volume
}

func (b ACXTicker) AskFloat() string {
	return b.Ask
}

func (b ACXTicker) BidFloat() string {
	return b.Bid
}

func (b ACXTicker) HighFloat() string {
	return b.High
}

func (b ACXTicker) LowFloat() string {
	return b.Low
}

func (b ACXTicker) LastFloat() string {
	return b.Last
}

func (b Coinjar) RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string) {
	// RequestUpdate(v.name, v.url, ch, v.currency, v.crypto)
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b,err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b,err, currency, crypto}
	}
}

func (b Coinjar) VolumeFloat() string {
	return b.Volume
}

func (b Coinjar) AskFloat() string {
	return b.Ask
}

func (b Coinjar) BidFloat() string {
	return b.Bid
}

func (b Coinjar) HighFloat() string {
	return b.High
}

func (b Coinjar) LowFloat() string {
	return b.Low
}

func (b Coinjar) LastFloat() string {
	return b.Last
}

type CryptoExchange interface {
	VolumeFloat() string
	AskFloat() string
	BidFloat() string
	HighFloat() string
	LowFloat() string
	LastFloat() string
	RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string)
	// RequestUpdate(v.name, v.url, ch, v.currency, v.crypto)
}
