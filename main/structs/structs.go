package structs


import  "github.com/shopspring/decimal"


// todo write interface to be used

/*
	bid_aud
	offer_aud
*/

type GeminiTickerBTC struct {
	Bid string            `json:"bid"`
	Ask string            `json:"ask"`
	Volume BTCVolume	  `json:"volume"`
	Last string           `json:"last"`
}

type BTCVolume struct {
	Value  string    `json:"BTC"`
	USD    string    `json:"USD"`
	Timestamp int    `json:"timestamp"`
}

type GeminiTickerETH struct {
	Bid string            `json:"bid"`
	Ask string            `json:"ask"`
	Volume ETHVolume	  `json:"volume"`
	Last string           `json:"last"`
}

type ETHVolume struct {
	Value    string    `json:"ETH"`
	USD      string    `json:"USD"`
	Timestamp int      `json:"timestamp"`
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