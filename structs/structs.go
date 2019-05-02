package structs

import (
	"encoding/json"
	"github.com/shopspring/decimal" // could probably just use float64
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
	GDP decimal.Decimal `json:"GBP"`
	AUD decimal.Decimal `json:"AUD"`
	EUR decimal.Decimal `json:"EUR"`
	USD decimal.Decimal `json:"USD"`
}

type CryptoDTO struct {
	name  string
	coin  CryptoExchange
	error error
}

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

func (b BTCMarket) volume() string {
	return strconv.FormatFloat(b.Volume, 'f', -1, 64)
}

func (b BTCMarket) ask() string {
	return strconv.FormatFloat(b.Ask, 'f', -1, 64)
}

func (b BTCMarket) bid() string {
	return strconv.FormatFloat(b.Bid, 'f', -1, 64)
}

func (b BTCMarket) high() string {
	return strconv.FormatFloat(b.High, 'f', -1, 64)
}

func (b BTCMarket) low() string {
	return strconv.FormatFloat(b.Low, 'f', -1, 64)
}

func (b BTCMarket) last() string {
	return strconv.FormatFloat(b.Last, 'f', -1, 64)
}

func (b BTCMarket) RequestUpdate(name string, url string, ch chan CryptoDTO) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b,err}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b,err}
	}
}

func (b IndepentReserve) RequestUpdate(name string, url string, ch chan CryptoDTO) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b,err}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b,err}
	}
}

func (b IndepentReserve) volume() string {
	return strconv.FormatFloat(b.Volume, 'f', -1, 64)
}

func (b IndepentReserve) ask() string {
	return strconv.FormatFloat(b.Ask, 'f', -1, 64)
}

func (b IndepentReserve) bid() string {
	return strconv.FormatFloat(b.Bid, 'f', -1, 64)
}

func (b IndepentReserve) high() string {
	return strconv.FormatFloat(b.High, 'f', -1, 64)
}

func (b IndepentReserve) low() string {
	return strconv.FormatFloat(b.Low, 'f', -1, 64)
}

func (b IndepentReserve) last() string {
	return strconv.FormatFloat(b.Last, 'f', -1, 64)
}

func (b GeminiTickerETH) RequestUpdate(name string, url string, ch chan CryptoDTO) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b,err}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b,err}
	}
}

func (b GeminiTickerETH) volume() string {
	return b.Volume
}

func (b GeminiTickerETH) ask() string {
	return b.Ask
}

func (b GeminiTickerETH) bid() string {
	return b.Bid
}

func (b GeminiTickerETH) high() string {
	return b.High
}

func (b GeminiTickerETH) low() string {
	return b.Low
}

func (b GeminiTickerETH) last() string {
	return b.Last
}

func (b GeminiTickerBTC) RequestUpdate(name string, url string, ch chan CryptoDTO) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b,err}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b,err}
	}
}

func (b GeminiTickerBTC) volume() string {
	return b.Volume
}

func (b GeminiTickerBTC) ask() string {
	return b.Ask
}

func (b GeminiTickerBTC) bid() string {
	return b.Bid
}

func (b GeminiTickerBTC) high() string {
	return b.High
}

func (b GeminiTickerBTC) low() string {
	return b.Low
}

func (b GeminiTickerBTC) last() string {
	return b.Last
}

func (b CoinfloorTickerAndBitstamp) RequestUpdate(name string, url string, ch chan CryptoDTO) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b,err}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b,err}
	}
}

func (b CoinfloorTickerAndBitstamp) volume() string {
	return b.Volume
}

func (b CoinfloorTickerAndBitstamp) ask() string {
	return b.Ask
}

func (b CoinfloorTickerAndBitstamp) bid() string {
	return b.Bid
}

func (b CoinfloorTickerAndBitstamp) high() string {
	return b.High
}

func (b CoinfloorTickerAndBitstamp) low() string {
	return b.Low
}

func (b CoinfloorTickerAndBitstamp) last() string {
	return b.Last
}

func (b ACXTicker) RequestUpdate(name string, url string, ch chan CryptoDTO) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b,err}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b,err}
	}
}

func (b ACXTicker) volume() string {
	return b.Volume
}

func (b ACXTicker) ask() string {
	return b.Ask
}

func (b ACXTicker) bid() string {
	return b.Bid
}

func (b ACXTicker) high() string {
	return b.High
}

func (b ACXTicker) low() string {
	return b.Low
}

func (b ACXTicker) last() string {
	return b.Last
}

func (b Coinjar) RequestUpdate(name string, url string, ch chan CryptoDTO) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b,err}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b,err}
	}
}

func (b Coinjar) volume() string {
	return b.Volume
}

func (b Coinjar) ask() string {
	return b.Ask
}

func (b Coinjar) bid() string {
	return b.Bid
}

func (b Coinjar) high() string {
	return b.High
}

func (b Coinjar) low() string {
	return b.Low
}

func (b Coinjar) last() string {
	return b.Last
}

type CryptoExchange interface {
	volume() string
	ask() string
	bid() string
	high() string
	low() string
	last() string
	//exchange.RequestUpdate(v.name, v.url, ch)
	RequestUpdate(name string, url string, ch chan CryptoDTO)
}
