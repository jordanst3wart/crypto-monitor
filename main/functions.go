package main

import (
	"crypto-monitor/main/cryptoExchanges"
	"crypto-monitor/main/fiatCurrencyExchange"
	"log"
	"math"
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
	exchange   string
	cryptoList []Four
}

func convertHelper(conversion float64, dto CryptoExchanges.CryptoData) CryptoExchanges.CryptoData {

	last, _ := dto.Coin.LastFloat()
	high, _ := dto.Coin.HighFloat()
	low, _ := dto.Coin.LowFloat()
	volume, _ := dto.Coin.VolumeFloat()
	ask, _ := dto.Coin.AskFloat()
	bid, _ := dto.Coin.BidFloat()

	return CryptoExchanges.CryptoData{
		Name: dto.Name,
		Coin: CryptoExchanges.BTCMarket{
			Last:   last * conversion,
			High:   high * conversion,
			Low:    low * conversion,
			Volume: volume * conversion,
			Bid:    bid * conversion,
			Ask:    ask * conversion},
		Error:    dto.Error,
		Currency: "AUD",
		Crypto:   dto.Crypto,
	}
}

func isZero(f float64) bool {
	epsilon := 1e-10
	return math.Abs(f) < epsilon
}

func CheckArbitrage(exchange1 CryptoExchanges.CryptoData, exchange2 CryptoExchanges.CryptoData) float64 {
	bid, _ := exchange1.Coin.BidFloat()
	ask, _ := exchange2.Coin.AskFloat()
	if isZero(ask) {
		log.Printf("Exchange %s ask value was %v changing value to 1.0", exchange2.Name, ask)
		return 1.0
	}
	return bid / ask
}

func ConvertCurrency(crypto CryptoExchanges.CryptoData, exchangeRate fiatCurrencyExchange.ExchangeRates) CryptoExchanges.CryptoData {
	switch crypto.Currency {
	case "USD":
		return convertHelper(exchangeRate.Rates["USD2AUD"], crypto)
	case "GBP":
		return convertHelper(exchangeRate.Rates["GBP2AUD"], crypto)
	case "EUR":
		return convertHelper(exchangeRate.Rates["EUR2AUD"], crypto)
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
	case "Bitstamp":
		var response CryptoExchanges.Bitstamp
		requestToExchange(response, data.cryptoList, ch)
	case "IndependentReserve":
		var response CryptoExchanges.IndependentReserve
		requestToExchange(response, data.cryptoList, ch)
	case "GeminiTickerBTC":
		var response CryptoExchanges.GeminiTickerBTC
		requestToExchange(response, data.cryptoList, ch)
	case "GeminiTickerETH":
		var response CryptoExchanges.GeminiTickerETH
		requestToExchange(response, data.cryptoList, ch)
	case "BTCMarket":
		var response CryptoExchanges.BTCMarket
		requestToExchange(response, data.cryptoList, ch)
	case "Coinjar":
		var response CryptoExchanges.Coinjar
		requestToExchange(response, data.cryptoList, ch)
	case "Binance":
		var response CryptoExchanges.Binance
		requestToExchange(response, data.cryptoList, ch)
	default:
		log.Println("Invalid key in exchange list")
	}
}

func ExchangeDataList() []startData {
	return []startData{{"Bitstamp", []Four{
		{"Bitstamp_USD_BTC", "https://www.bitstamp.net/api/v2/ticker/btcusd/", "USD", "BTC"},
		{"Bitstamp_USD_XRP", "https://www.bitstamp.net/api/v2/ticker/xrpusd/", "USD", "XRP"},
		{"Bitstamp_USD_LTC", "https://www.bitstamp.net/api/v2/ticker/ltcusd/", "USD", "LTC"},
		{"Bitstamp_USD_ETH", "https://www.bitstamp.net/api/v2/ticker/ethusd/", "USD", "ETH"},
		{"Bitstamp_USD_SOL", "https://www.bitstamp.net/api/v2/ticker/solusd/", "USD", "SOL"},
		{"Bitstamp_USD_ADA", "https://www.bitstamp.net/api/v2/ticker/adausd/", "USD", "ADA"},
		{"Bitstamp_USD_AVAX", "https://www.bitstamp.net/api/v2/ticker/avaxusd/", "USD", "AVAX"},
		{"Bitstamp_USD_DOT", "https://www.bitstamp.net/api/v2/ticker/dotusd/", "USD", "DOT"},
		{"Bitstamp_USD_MATIC", "https://www.bitstamp.net/api/v2/ticker/maticusd/", "USD", "MATIC"},
		{"Bitstamp_USD_LINK", "https://www.bitstamp.net/api/v2/ticker/linkusd/", "USD", "LINK"},
		{"Bitstamp_USD_UNI", "https://www.bitstamp.net/api/v2/ticker/uniusd/", "USD", "UNI"},
		{"Bitstamp_USD_MANA", "https://www.bitstamp.net/api/v2/ticker/manausd/", "USD", "MANA"},
		{"Bitstamp_USD_SUI", "https://www.bitstamp.net/api/v2/ticker/suiusd/", "USD", "SUI"},
		{"Bitstamp_USD_DOGE", "https://www.bitstamp.net/api/v2/ticker/dogeusd/", "USD", "DOGE"},
		{"Bitstamp_USD_BAT", "https://www.bitstamp.net/api/v2/ticker/batusd/", "USD", "BAT"},
		{"Bitstamp_USD_XLM", "https://www.bitstamp.net/api/v2/ticker/xlmusd/", "USD", "XLM"},
		{"Bitstamp_USD_INJ", "https://www.bitstamp.net/api/v2/ticker/injusd/", "USD", "INJ"},
		{"Bitstamp_USD_AAVE", "https://www.bitstamp.net/api/v2/ticker/aaveusd/", "USD", "AAVE"},
	}},
		startData{"IndependentReserve", []Four{
			{"IndependentReserve_AUD_BTC", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xbt&secondaryCurrencyCode=aud", "AUD", "BTC"},
			{"IndependentReserve_AUD_ETH", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=eth&secondaryCurrencyCode=aud", "AUD", "ETH"},
			{"IndependentReserve_AUD_XRP", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xrp&secondaryCurrencyCode=aud", "AUD", "XRP"},
			{"IndependentReserve_AUD_LTC", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=ltc&secondaryCurrencyCode=aud", "AUD", "LTC"},
			{"IndependentReserve_AUD_SOL", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=sol&secondaryCurrencyCode=aud", "AUD", "SOL"},
			{"IndependentReserve_AUD_ADA", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=ada&secondaryCurrencyCode=aud", "AUD", "ADA"},
			{"IndependentReserve_AUD_MATIC", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=matic&secondaryCurrencyCode=aud", "AUD", "MATIC"},
			{"IndependentReserve_AUD_XLM", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xlm&secondaryCurrencyCode=aud", "AUD", "XLM"},
			{"IndependentReserve_AUD_DOGE", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=doge&secondaryCurrencyCode=aud", "AUD", "DOGE"},
			{"IndependentReserve_AUD_LINK", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=link&secondaryCurrencyCode=aud", "AUD", "LINK"},
			{"IndependentReserve_AUD_AAVE", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=aave&secondaryCurrencyCode=aud", "AUD", "AAVE"},
			{"IndependentReserve_AUD_DOT", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=dot&secondaryCurrencyCode=aud", "AUD", "DOT"},
			{"IndependentReserve_AUD_MANA", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=mana&secondaryCurrencyCode=aud", "AUD", "MANA"},
			{"IndependentReserve_AUD_UNI", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=uni&secondaryCurrencyCode=aud", "AUD", "UNI"},
			{"IndependentReserve_AUD_BAT", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=bat&secondaryCurrencyCode=aud", "AUD", "BAT"},
		}},
		startData{"GeminiTickerBTC", []Four{
			{"GEMINI_BTC", "https://api.gemini.com/v1/pubticker/btcusd", "USD", "BTC"}}},
		startData{"GeminiTickerETH", []Four{
			{"GEMINI_ETH", "https://api.gemini.com/v1/pubticker/ethusd", "USD", "ETH"}}},
		startData{"BTCMarket", []Four{
			{"BTCMarket_AUD_BTC", "https://api.btcmarkets.net/market/BTC/AUD/tick", "AUD", "BTC"},
			{"BTCMarket_AUD_ETH", "https://api.btcmarkets.net/market/ETH/AUD/tick", "AUD", "ETH"},
			{"BTCMarket_AUD_XRP", "https://api.btcmarkets.net/market/XRP/AUD/tick", "AUD", "XRP"},
			{"BTCMarket_AUD_LTC", "https://api.btcmarkets.net/market/LTC/AUD/tick", "AUD", "LTC"},
			{"BTCMarket_AUD_SOL", "https://api.btcmarkets.net/market/SOL/AUD/tick", "AUD", "SOL"},
			{"BTCMarket_AUD_ADA", "https://api.btcmarkets.net/market/ADA/AUD/tick", "AUD", "ADA"},
			{"BTCMarket_AUD_XLM", "https://api.btcmarkets.net/market/XLM/AUD/tick", "AUD", "XLM"},
			{"BTCMarket_AUD_AVAX", "https://api.btcmarkets.net/market/AVAX/AUD/tick", "AUD", "AVAX"},
			{"BTCMarket_AUD_LINK", "https://api.btcmarkets.net/market/LINK/AUD/tick", "AUD", "LINK"},
			{"BTCMarket_AUD_ALGO", "https://api.btcmarkets.net/market/ALGO/AUD/tick", "AUD", "ALGO"},
			{"BTCMarket_AUD_AAVE", "https://api.btcmarkets.net/market/AAVE/AUD/tick", "AUD", "AAVE"},
			{"BTCMarket_AUD_SUI", "https://api.btcmarkets.net/market/SUI/AUD/tick", "AUD", "SUI"},
			{"BTCMarket_AUD_DOT", "https://api.btcmarkets.net/market/DOT/AUD/tick", "AUD", "DOT"},
			{"BTCMarket_AUD_OMG", "https://api.btcmarkets.net/market/OMG/AUD/tick", "AUD", "OMG"},
			{"BTCMarket_AUD_POWR", "https://api.btcmarkets.net/market/POWR/AUD/tick", "AUD", "POWR"},
			{"BTCMarket_AUD_UNI", "https://api.btcmarkets.net/market/UNI/AUD/tick", "AUD", "UNI"},
			{"BTCMarket_AUD_MANA", "https://api.btcmarkets.net/market/MANA/AUD/tick", "AUD", "MANA"},
			{"BTCMarket_AUD_BAT", "https://api.btcmarkets.net/market/BAT/AUD/tick", "AUD", "BAT"},
		}},
		startData{"Coinjar", []Four{
			{"Coinjar_AUD_BTC", "https://data.exchange.coinjar.com/products/BTCAUD/ticker", "AUD", "BTC"},
			{"Coinjar_AUD_ETH", "https://data.exchange.coinjar.com/products/ETHAUD/ticker", "AUD", "ETH"},
			{"Coinjar_AUD_XRP", "https://data.exchange.coinjar.com/products/XRPAUD/ticker", "AUD", "XRP"},
			{"Coinjar_AUD_LTC", "https://data.exchange.coinjar.com/products/LTCAUD/ticker", "AUD", "LTC"},
			{"Coinjar_AUD_SOL", "https://data.exchange.coinjar.com/products/SOLAUD/ticker", "AUD", "SOL"},
			{"Coinjar_AUD_INJ", "https://data.exchange.coinjar.com/products/INJAUD/ticker", "AUD", "INJ"},
			{"Coinjar_AUD_UNI", "https://data.exchange.coinjar.com/products/UNIAUD/ticker", "AUD", "UNI"},
			{"Coinjar_AUD_LINK", "https://data.exchange.coinjar.com/products/LINKAUD/ticker", "AUD", "LINK"},
			{"Coinjar_AUD_DOGE", "https://data.exchange.coinjar.com/products/DOGEAUD/ticker", "AUD", "DOGE"},
			{"Coinjar_AUD_SHIB", "https://data.exchange.coinjar.com/products/SHIBAUD/ticker", "AUD", "SHIB"},
			{"Coinjar_AUD_ADA", "https://data.exchange.coinjar.com/products/ADAAUD/ticker", "AUD", "ADA"},
			{"Coinjar_AUD_ALGO", "https://data.exchange.coinjar.com/products/ALGOAUD/ticker", "AUD", "ALGO"},
			{"Coinjar_AUD_MATIC", "https://data.exchange.coinjar.com/products/MATICAUD/ticker", "AUD", "MATIC"},
			{"Coinjar_AUD_XLM", "https://data.exchange.coinjar.com/products/XLMAUD/ticker", "AUD", "XLM"},
			{"Coinjar_AUD_OMG", "https://data.exchange.coinjar.com/products/OMGAUD/ticker", "AUD", "OMG"},
			{"Coinjar_AUD_AAVE", "https://data.exchange.coinjar.com/products/AAVEAUD/ticker", "AUD", "AAVE"},
			{"Coinjar_AUD_DOT", "https://data.exchange.coinjar.com/products/DOTAUD/ticker", "AUD", "DOT"},
		}},
		startData{"Binance", []Four{
			{"Binance_EUR_BTC", "https://api.binance.com/api/v3/ticker/24hr?symbol=BTCEUR", "EUR", "BTC"},
			{"Binance_EUR_ETH", "https://api.binance.com/api/v3/ticker/24hr?symbol=ETHEUR", "EUR", "ETH"},
			{"Binance_EUR_XRP", "https://api.binance.com/api/v3/ticker/24hr?symbol=XRPEUR", "EUR", "XRP"},
			{"Binance_EUR_LTC", "https://api.binance.com/api/v3/ticker/24hr?symbol=LTCEUR", "EUR", "LTC"},
			{"Binance_EUR_SOL", "https://api.binance.com/api/v3/ticker/24hr?symbol=SOLEUR", "EUR", "SOL"},
			// USDT {"Binance_EUR_INJ", "https://api.binance.com/api/v3/ticker/24hr?symbol=INJEUR", "EUR", "INJ"},
			// {"Binance_EUR_AAVE", "https://api.binance.com/api/v3/ticker/24hr?symbol=AAVEEUR", "EUR", "AAVE"},
			//maybe USDT {"Binance_EUR_UNI", "https://api.binance.com/api/v3/ticker/24hr?symbol=UNIEUR", "EUR", "UNI"},
			{"Binance_EUR_LINK", "https://api.binance.com/api/v3/ticker/24hr?symbol=LINKEUR", "EUR", "LINK"},
			{"Binance_EUR_DOGE", "https://api.binance.com/api/v3/ticker/24hr?symbol=DOGEEUR", "EUR", "DOGE"},
			{"Binance_EUR_SHIB", "https://api.binance.com/api/v3/ticker/24hr?symbol=SHIBEUR", "EUR", "SHIB"},
			{"Binance_EUR_ADA", "https://api.binance.com/api/v3/ticker/24hr?symbol=ADAEUR", "EUR", "ADA"},
			// USDT {"Binance_EUR_ALGO", "https://api.binance.com/api/v3/ticker/24hr?symbol=ALGOEUR", "EUR", "ALGO"},
			{"Binance_EUR_MATIC", "https://api.binance.com/api/v3/ticker/24hr?symbol=MATICEUR", "EUR", "MATIC"},
			{"Binance_EUR_XLM", "https://api.binance.com/api/v3/ticker/24hr?symbol=XLMEUR", "EUR", "XLM"},
			// maybe USDT {"Binance_EUR_OMG", "https://api.binance.com/api/v3/ticker/24hr?symbol=OMGEUR", "EUR", "OMG"},
			{"Binance_EUR_DOT", "https://api.binance.com/api/v3/ticker/24hr?symbol=DOTEUR", "EUR", "DOT"},
		}},
	}
}
