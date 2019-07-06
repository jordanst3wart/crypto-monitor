package main

import (
	"crypto-monitor/structs"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
	Gets exchange rate, every 5 minutes
*/

func fiatCurrencyExchangeRates(ch chan ExchangeRates) {
	exchangeMap := make(map[string]float64)
	var responseObject structs.CurrencyExchangeAPI
	// TODO need to cache
	for {
		resp, err := http.Get("https://api.exchangeratesapi.io/latest?base=USD")
		if err != nil {
			// should wrap error
			ch <- ExchangeRates{nil, err}
			return
		}
		// https://fixer.io/product

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			ch <- ExchangeRates{nil, err}
			return
		}
		err = json.Unmarshal(responseData, &responseObject)
		if err != nil {
			ch <- ExchangeRates{nil, err}
			return
		} else {
			exchangeMap["USD2AUD"] = responseObject.Rates.AUD / responseObject.Rates.USD
			exchangeMap["GBP2AUD"] = responseObject.Rates.AUD / responseObject.Rates.GDP
			ch <- ExchangeRates{exchangeMap, err}
		}
		time.Sleep(5 * time.Minute)
	}
}

type Four struct {
	name, url, currency, crypto string
}

func requestToExchange(exchange structs.CryptoExchange, urlList []Four, ch chan structs.CryptoDTO){
	for _, v := range urlList {
		go exchange.RequestUpdate(v.name, v.url, ch, v.currency, v.crypto)
		//ch<-CryptoDTO{v.name,val, err}
		//groupList = append(groupList, CryptoDTO{v.name,val, err})
	}
	//return groupList
}

type ExchangeRates struct {
	rates map[string]float64
	err error
}

type startData struct {
	exchange string
	list []Four
}

func convertHelper(conversion float64, dto structs.CryptoDTO) (structs.CryptoDTO){

	last, _   :=dto.Coin.LastFloat()
	high, _   := dto.Coin.HighFloat()
	low, _    := dto.Coin.LowFloat()
	volume, _ := dto.Coin.VolumeFloat()
	ask, _    := dto.Coin.AskFloat()
	bid, _    := dto.Coin.BidFloat()

	tmpCoin := structs.BTCMarket{
		last * conversion,
		high * conversion,
		low * conversion,
		volume * conversion,
		bid * conversion,
		ask * conversion}

	return structs.CryptoDTO{
		dto.Name,
		tmpCoin,
		dto.Error,
		"AUD",
		dto.Crypto,
	}
}

func CheckArbitage(exchange1 structs.CryptoDTO, exchange2 structs.CryptoDTO) float64 {
	bid, _ := exchange1.Coin.BidFloat()
	ask, _ := exchange2.Coin.AskFloat()
	return bid / ask
}

func ConvertCurrency(crypto structs.CryptoDTO, exchangeRate ExchangeRates) structs.CryptoDTO {
	switch crypto.Currency {
	case "USD":
		return convertHelper(exchangeRate.rates["USD2AUD"],crypto)
	case "GBP":
		return convertHelper(exchangeRate.rates["GBP2AUD"],crypto)
	case "AUD":
		return crypto
	default:
		log.Println("Unknown currency trying to be converted")
		return structs.CryptoDTO{}
	}
}

func UniqueStrings(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}


func calculate(data startData, ch chan structs.CryptoDTO) {
	switch data.exchange {
	case "CoinfloorTickerAndBitstamp":
		//log.Println("Requesting data from CoinfloorTickerAndBitstamp")
		var resseObjectCoinfloorAndBitstamp structs.CoinfloorTickerAndBitstamp
		requestToExchange(resseObjectCoinfloorAndBitstamp, data.list, ch)
	case "IndepentReserve":
		//log.Println("Requesting data from IndepentReserve")
		var responseObjectIndependentReserve structs.IndepentReserve
		requestToExchange(responseObjectIndependentReserve, data.list, ch)
	case "GeminiTickerBTC":
		//log.Println("Requesting data from IndepentReserve")
		var responseObjectGeminiBTC structs.GeminiTickerBTC
		requestToExchange(responseObjectGeminiBTC, data.list, ch)
	case "GeminiTickerETH":
		//log.Println("Requesting data from GeminiTickerETH")
		var responseObjectGeminiETH structs.GeminiTickerETH
		requestToExchange(responseObjectGeminiETH, data.list, ch)
	case "BTCMarket":
		//log.Println("Requesting data from BTCMarket")
		var responseObjectBTC structs.BTCMarket
		requestToExchange(responseObjectBTC, data.list, ch)
	case "ACXTicker":
		//log.Println("Requesting data from ACX")
		var responseObjectACX structs.ACXTicker
		requestToExchange(responseObjectACX, data.list, ch)
	case "Coinjar":
		//log.Println("Requesting data from Coinjar")
		var responseObjectCoinjar structs.Coinjar
		requestToExchange(responseObjectCoinjar, data.list, ch)
	default:
		log.Println("Invalid key in startData")
	}
}

func getStartData() []startData {
	return []startData{{"CoinfloorTickerAndBitstamp", []Four{
		{"Coinfloor_BTC", "https://webapi.coinfloor.co.uk/bist/XBT/GBP/ticker/", "GBP", "BTC"},
		{"Coinfloor_ETH","https://webapi.coinfloor.co.uk/bist/ETH/GBP/ticker/", "GBP", "ETH"},
		/*{"Coinfloor_BCH","https://webapi.coinfloor.co.uk/bist/BCH/GBP/ticker/", "GBP", "BCH"},  no longer supported */
		{"Bitstamp_BTC","https://www.bitstamp.net/api/v2/ticker/btcusd/", "USD", "BTC"},
		{"Bitstamp_XRP","https://www.bitstamp.net/api/v2/ticker/xrpusd/", "USD", "XRP"},
		{"Bitstamp_LTC","https://www.bitstamp.net/api/v2/ticker/ltcusd/", "USD", "LTC"},
		{"Bitstamp_ETH","https://www.bitstamp.net/api/v2/ticker/ethusd/", "USD", "ETH"},
		{"Bitstamp_BCH","https://www.bitstamp.net/api/v2/ticker/bchusd/", "USD", "BCH"}}},
		startData{"IndepentReserve",[]Four{
			{"IndependentReserve_BTC", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xbt&secondaryCurrencyCode=aud", "AUD","BTC"},
			{"IndependentReserve_ETH","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=eth&secondaryCurrencyCode=aud", "AUD", "ETH"},
			{"IndependentReserve_BCH","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=bch&secondaryCurrencyCode=aud", "AUD", "BCH"},
			{"IndependentReserve_XRP","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xrp&secondaryCurrencyCode=aud", "AUD", "XRP"},
			{"IndependentReserve_LTC","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=ltc&secondaryCurrencyCode=aud", "AUD", "LTC"}}},
		startData{"GeminiTickerBTC", []Four{
			{"GEMINI_BTC", "https://api.gemini.com/v1/pubticker/btcusd", "USD", "BTC"}}},
		startData{"GeminiTickerETH",[]Four{
			{"GEMINI_ETH", "https://api.gemini.com/v1/pubticker/ethusd", "USD", "ETH"}}},
		startData{"BTCMarket", []Four{
			{"BTCMarket_AUD_BTC", "https://api.btcmarkets.net/market/BTC/AUD/tick", "AUD", "BTC"},
			{"BTCMarket_AUD_ETH", "https://api.btcmarkets.net/market/ETH/AUD/tick", "AUD", "ETH"},
			{"BTCMarket_AUD_BCH","https://api.btcmarkets.net/market/BCHABC/AUD/tick", "AUD", "BCH"},
			{"BTCMarket_AUD_XRP", "https://api.btcmarkets.net/market/XRP/AUD/tick", "AUD", "XRP"},
			{"BTCMarket_AUD_LTC","https://api.btcmarkets.net/market/LTC/AUD/tick", "AUD", "LTC"}}},
		startData{"ACXTicker",[]Four{
			{"ACX_AUD_BTC", "https://acx.io:443/api/v2/tickers/btcaud.json", "AUD", "BTC"},
			{"ACX_AUD_ETH", "https://acx.io:443/api/v2/tickers/ethaud.json", "AUD", "ETH"},
			{"ACX_AUD_BCH","https://acx.io:443/api/v2/tickers/bchaud.json", "AUD", "BCH"},
			{"ACX_AUD_LTC", "https://acx.io:443/api/v2/tickers/ltcaud.json", "AUD", "LTC"},
			{"ACX_AUD_XRP","https://acx.io:443/api/v2/tickers/xrpaud.json","AUD", "XRP"}}},
		startData{"Coinjar",[]Four{
			{"Coinjar_AUD_BTC", "https://data.exchange.coinjar.com/products/BTCAUD/ticker", "AUD", "BTC"},
			{"Coinjar_AUD_ETH", "https://data.exchange.coinjar.com/products/ETHAUD/ticker", "AUD", "ETH"},
			{"Coinjar_AUD_XRP","https://data.exchange.coinjar.com/products/XRPAUD/ticker","AUD", "XRP"},
			{"Coinjar_AUD_LTC", "https://data.exchange.coinjar.com/products/LTCAUD/ticker", "AUD", "LTC"}}}}
}
