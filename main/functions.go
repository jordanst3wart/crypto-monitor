package main

import (
	"crypto-monitor/main/cryptoExchanges"
	"crypto-monitor/main/fiatCurrencyExchange"
	"log"
)

type Four struct {
	name, url, currency, crypto string
}

func requestToExchange(exchange CryptoExchanges.CryptoExchange, urlList []Four, ch chan CryptoExchanges.CryptoData) {
	for _, v := range urlList {
		go exchange.RequestUpdate(v.name, v.url, ch, v.currency, v.crypto)
	}
}

type startData struct {
	exchange string
	list     []Four
}

func convertHelper(conversion float64, dto CryptoExchanges.CryptoData) CryptoExchanges.CryptoData {

	last, _ := dto.Coin.LastFloat()
	high, _ := dto.Coin.HighFloat()
	low, _ := dto.Coin.LowFloat()
	volume, _ := dto.Coin.VolumeFloat()
	ask, _ := dto.Coin.AskFloat()
	bid, _ := dto.Coin.BidFloat()

	tmpCoin := CryptoExchanges.BTCMarket{
		last * conversion,
		high * conversion,
		low * conversion,
		volume * conversion,
		bid * conversion,
		ask * conversion}

	return CryptoExchanges.CryptoData{
		dto.Name,
		tmpCoin,
		dto.Error,
		"AUD",
		dto.Crypto,
	}
}

func CheckArbitrage(exchange1 CryptoExchanges.CryptoData, exchange2 CryptoExchanges.CryptoData) float64 {
	bid, _ := exchange1.Coin.BidFloat()
	ask, _ := exchange2.Coin.AskFloat()
	return bid / ask
}

func ConvertCurrency(crypto CryptoExchanges.CryptoData, exchangeRate fiatCurrencyExchange.ExchangeRates) CryptoExchanges.CryptoData {
	switch crypto.Currency {
	case "USD":
		return convertHelper(exchangeRate.Rates["USD2AUD"], crypto)
	case "GBP":
		return convertHelper(exchangeRate.Rates["GBP2AUD"], crypto)
	case "AUD":
		return crypto
	default:
		log.Println("Unknown currency trying to be converted")
		return CryptoExchanges.CryptoData{}
	}
}

func DeduplicateStrings(input []string) []string {
	uniqueStrings := make([]string, 0, len(input))
	seenMap := make(map[string]bool)

	for _, str := range input {
		if _, exists := seenMap[str]; !exists {
			seenMap[str] = true
			uniqueStrings = append(uniqueStrings, str)
		}
	}

	return uniqueStrings
}

func exchangeMutex(data startData, ch chan CryptoExchanges.CryptoData) {
	switch data.exchange {
	case "CoinfloorTickerAndBitstamp":
		var resseObjectCoinfloorAndBitstamp CryptoExchanges.CoinfloorTickerAndBitstamp
		requestToExchange(resseObjectCoinfloorAndBitstamp, data.list, ch)
	case "IndependentReserve":
		var responseObjectIndependentReserve CryptoExchanges.IndependentReserve
		requestToExchange(responseObjectIndependentReserve, data.list, ch)
	case "GeminiTickerBTC":
		var responseObjectGeminiBTC CryptoExchanges.GeminiTickerBTC
		requestToExchange(responseObjectGeminiBTC, data.list, ch)
	case "GeminiTickerETH":
		var responseObjectGeminiETH CryptoExchanges.GeminiTickerETH
		requestToExchange(responseObjectGeminiETH, data.list, ch)
	case "BTCMarket":
		var responseObjectBTC CryptoExchanges.BTCMarket
		requestToExchange(responseObjectBTC, data.list, ch)
	case "ACXTicker":
		var responseObjectACX CryptoExchanges.ACXTicker
		requestToExchange(responseObjectACX, data.list, ch)
	case "Coinjar":
		var responseObjectCoinjar CryptoExchanges.Coinjar
		requestToExchange(responseObjectCoinjar, data.list, ch)
	default:
		log.Println("Invalid key in startData")
	}
}

func ExchangeDataList() []startData {
	return []startData{{"CoinfloorTickerAndBitstamp", []Four{
		{"CoinCorner_BTC", "https://api.coincorner.com/api/Ticker?Coin=BTC&Currency=GBP", "GBP", "BTC"},
		{"CoinCorner_ETH", "https://api.coincorner.com/api/Ticker?Coin=ETH&Currency=GBP", "GBP", "ETH"},
		/*{"Coinfloor_BCH","https://webapi.coinfloor.co.uk/bist/BCH/GBP/ticker/", "GBP", "BCH"},  no longer supported */
		{"Bitstamp_BTC", "https://www.bitstamp.net/api/v2/ticker/btcusd/", "USD", "BTC"},
		{"Bitstamp_XRP", "https://www.bitstamp.net/api/v2/ticker/xrpusd/", "USD", "XRP"},
		{"Bitstamp_LTC", "https://www.bitstamp.net/api/v2/ticker/ltcusd/", "USD", "LTC"},
		{"Bitstamp_ETH", "https://www.bitstamp.net/api/v2/ticker/ethusd/", "USD", "ETH"},
		{"Bitstamp_BCH", "https://www.bitstamp.net/api/v2/ticker/bchusd/", "USD", "BCH"}}},
		startData{"IndependentReserve", []Four{
			{"IndependentReserve_BTC", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xbt&secondaryCurrencyCode=aud", "AUD", "BTC"},
			{"IndependentReserve_ETH", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=eth&secondaryCurrencyCode=aud", "AUD", "ETH"},
			{"IndependentReserve_BCH", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=bch&secondaryCurrencyCode=aud", "AUD", "BCH"},
			{"IndependentReserve_XRP", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xrp&secondaryCurrencyCode=aud", "AUD", "XRP"},
			{"IndependentReserve_LTC", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=ltc&secondaryCurrencyCode=aud", "AUD", "LTC"}}},
		startData{"GeminiTickerBTC", []Four{
			{"GEMINI_BTC", "https://api.gemini.com/v1/pubticker/btcusd", "USD", "BTC"}}},
		startData{"GeminiTickerETH", []Four{
			{"GEMINI_ETH", "https://api.gemini.com/v1/pubticker/ethusd", "USD", "ETH"}}},
		startData{"BTCMarket", []Four{
			{"BTCMarket_AUD_BTC", "https://api.btcmarkets.net/market/BTC/AUD/tick", "AUD", "BTC"},
			{"BTCMarket_AUD_ETH", "https://api.btcmarkets.net/market/ETH/AUD/tick", "AUD", "ETH"},
			{"BTCMarket_AUD_BCH", "https://api.btcmarkets.net/market/BCHABC/AUD/tick", "AUD", "BCH"},
			{"BTCMarket_AUD_XRP", "https://api.btcmarkets.net/market/XRP/AUD/tick", "AUD", "XRP"},
			{"BTCMarket_AUD_LTC", "https://api.btcmarkets.net/market/LTC/AUD/tick", "AUD", "LTC"}}},
		startData{"Coinjar", []Four{
			{"Coinjar_AUD_BTC", "https://data.exchange.coinjar.com/products/BTCAUD/ticker", "AUD", "BTC"},
			{"Coinjar_AUD_ETH", "https://data.exchange.coinjar.com/products/ETHAUD/ticker", "AUD", "ETH"},
			{"Coinjar_AUD_XRP", "https://data.exchange.coinjar.com/products/XRPAUD/ticker", "AUD", "XRP"},
			{"Coinjar_AUD_LTC", "https://data.exchange.coinjar.com/products/LTCAUD/ticker", "AUD", "LTC"}}}}
}
