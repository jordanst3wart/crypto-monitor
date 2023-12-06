package structs

import (
	"encoding/json"
	errors "github.com/pkg/errors"
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
	Base  string           `json:"base"`
	Rates CurrencyExchange `json:"rates"`
	Date  string           `json:"date"`
}

type CurrencyExchange struct {
	GBP float64 `json:"GBP"`
	AUD float64 `json:"AUD"`
	EUR float64 `json:"EUR"`
	USD float64 `json:"USD"`
}

type CryptoDTO struct {
	Name     string
	Coin     CryptoExchange
	Error    error
	Currency string
	Crypto   string
}

//

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

type CoinfloorTickerAndBitstamp struct {
	Last   string `json:"last"`
	High   string `json:"high"`
	Low    string `json:"low"`
	Volume string `json:"volume"`
	Bid    string `json:"bid"`
	Ask    string `json:"ask"`
	Vwap   string `json:"vwap"`
}

type IndepentReserve struct {
	Last   float64 `json:"LastPrice"`
	High   float64 `json:"DayHighestPrice"`
	Low    float64 `json:"DayLowestPrice"`
	Volume float64 `json:"DayVolumeXbt"`
	Bid    float64 `json:"CurrentHighestBidPrice"`
	Ask    float64 `json:"CurrentLowestOfferPrice"`
}

type BTCMarket struct {
	Last   float64 `json:"lastPrice"`
	High   float64 `json:"-"`
	Low    float64 `json:"-"`
	Volume float64 `json:"volume24h"`
	Bid    float64 `json:"bestBid"`
	Ask    float64 `json:"bestAsk"`
}

type Coinjar struct {
	Last   string `json:"last"`
	High   string `json:"-"`
	Low    string `json:"-"`
	Volume string `json:"volume_24h"`
	Bid    string `json:"bid"`
	Ask    string `json:"ask"`
}

// TODO put logic in here for RequestUpdate to reduce repeating code
func requestWrapper(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		err = errors.Wrap(err, " get request failed")
	}
	responseData, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		err = errors.Wrap(err, err2.Error()+"response data is: "+string(responseData))
	}
	return responseData, err
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

func (b IndepentReserve) RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b, err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b, err, currency, crypto}
	}
}

func (b IndepentReserve) VolumeFloat() (float64, error) {
	return b.Volume, nil
}

func (b IndepentReserve) AskFloat() (float64, error) {
	return b.Ask, nil
}

func (b IndepentReserve) BidFloat() (float64, error) {
	return b.Bid, nil
}

func (b IndepentReserve) HighFloat() (float64, error) {
	return b.High, nil
}

func (b IndepentReserve) LowFloat() (float64, error) {
	return b.Low, nil
}

func (b IndepentReserve) LastFloat() (float64, error) {
	return b.Last, nil
}

func (b GeminiTickerETH) RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b, err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b, err, currency, crypto}
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

func (b GeminiTickerBTC) RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b, err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b, err, currency, crypto}
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

func (b ACXTicker) RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string) {
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b, err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b, err, currency, crypto}
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

func (b Coinjar) RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string) {
	// RequestUpdate(v.name, v.url, ch, v.currency, v.crypto)
	responseData, err := requestWrapper(url)
	if err != nil {
		ch <- CryptoDTO{name, b, err, currency, crypto}
		//return b, err
	} else {
		err = json.Unmarshal(responseData, &b)
		ch <- CryptoDTO{name, b, err, currency, crypto}
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

type CryptoExchange interface {
	VolumeFloat() (float64, error)
	AskFloat() (float64, error)
	BidFloat() (float64, error)
	HighFloat() (float64, error)
	LowFloat() (float64, error)
	LastFloat() (float64, error)
	RequestUpdate(name string, url string, ch chan CryptoDTO, currency string, crypto string)
	// RequestUpdate(v.name, v.url, ch, v.currency, v.crypto)
}
