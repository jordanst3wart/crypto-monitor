package structs

import "github.com/shopspring/decimal"

// todo write interface to be used

/*
	bid_aud
	offer_aud
*/

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

type GenericCryptoResponse struct {
	Last   string
	High   string
	Low    string
	Volume string
	Bid    string
	Ask    string
}

type IndepentReserve struct {
	Last   string   `json:"LastPrice"`
	High   string   `json:"DayHighestPrice"`
	Low    string   `json:"DayLowestPrice"`
	Volume string   `json:"DayVolumeXbt"`
	Bid    string   `json:"CurrentHighestBidPrice"`
	Ask    string   `json:"CurrentLowestOfferPrice"`
}

type BTCMarket struct {
	Last   string    `json:"lastPrice"`
	High   string    `json:"-"`
	Low    string    `json:"-"`
	Volume string    `json:"volume24h"`
	Bid    string    `json:"bestBid"`
	Ask    string    `json:"bestAsk"`
}

func (b BTCMarket) volume() string {
	return b.Volume
}

func (b BTCMarket) ask() string {
	return b.Ask
}

func (b BTCMarket) bid() string {
	return b.Bid
}

func (b BTCMarket) high() string {
	return b.High
}

func (b BTCMarket) low() string {
	return b.Low
}

func (b BTCMarket) last() string {
	return b.Last
}

func (b IndepentReserve) volume() string {
	return b.Volume
}

func (b IndepentReserve) ask() string {
	return b.Ask
}

func (b IndepentReserve) bid() string {
	return b.Bid
}

func (b IndepentReserve) high() string {
	return b.High
}

func (b IndepentReserve) low() string {
	return b.Low
}

func (b IndepentReserve) last() string {
	return b.Last
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

type CryptoExchange interface {
	volume() string
	ask() string
	bid() string
	high() string
	low() string
	last() string
}
