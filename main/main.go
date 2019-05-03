package main


// go run pokemon.go

import (
	"crypto-monitor/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


// TODO simplify

func currencyExchangeRates(ch chan ExchangeRates) {
	exchangeMap := make(map[string]float64)

	resp, err := http.Get("https://api.exchangeratesapi.io/latest?base=USD")
	if err != nil {
		// should wrap error
		ch <- ExchangeRates{nil, err}
		return
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- ExchangeRates{nil, err}
		return
	}
	var responseObject structs.CurrencyExchangeAPI
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		ch <- ExchangeRates{nil, err}
		return
	} else {
		exchangeMap["USD2AUD"] = responseObject.Rates.AUD / responseObject.Rates.USD
		exchangeMap["GBP2AUD"] = responseObject.Rates.AUD / responseObject.Rates.GDP
		ch <-ExchangeRates{exchangeMap, err}
	}
}

type Four struct {
	name, url, currency, crypto string
}

func requestToExchange(exchange structs.CryptoExchange, urlList []Four, ch chan structs.CryptoDTO){
	for _, v := range urlList {
		go exchange.RequestUpdate(v.name, v.url, ch)
		//ch<-CryptoDTO{v.name,val, err}
		//groupList = append(groupList, CryptoDTO{v.name,val, err})
	}
	//return groupList
}


type ExchangeRates struct {
	rates map[string]float64
	err error
}

func main() {
	start := time.Now()

	//DEBUG := true
	chRates := make(chan ExchangeRates)
	go currencyExchangeRates(chRates)
	//groupList := []structs.CryptoDTO{}
	ch := make(chan structs.CryptoDTO)
	// log error at main level
	/*if err != nil {
		// possibly send email
		log.Fatal(err.Error())
	}
	if DEBUG {
		log.Println("Got values from exchange")
		log.Println(val)
	}*/

	urlList := []Four{
		{"Coinfloor_BTC", "https://webapi.coinfloor.co.uk:8090/bist/XBT/GBP/ticker/", "GBP", "BTC"},
		{"Coinfloor_ETH","https://webapi.coinfloor.co.uk:8090/bist/ETH/GBP/ticker/", "GBP", "ETH"},
		{"Coinfloor_BCH","https://webapi.coinfloor.co.uk:8090/bist/BCH/GBP/ticker/", "GBP", "BCH"},
		{"Bitstamp_BTC","https://www.bitstamp.net/api/v2/ticker/btcusd/", "USD", "BTC"},
		{"Bitstamp_XRP","https://www.bitstamp.net/api/v2/ticker/xrpusd/", "USD", "XRP"},
		{"Bitstamp_LTC","https://www.bitstamp.net/api/v2/ticker/ltcusd/", "USD", "LTC"},
		{"Bitstamp_ETH","https://www.bitstamp.net/api/v2/ticker/ethusd/", "USD", "ETH"},
		{"Bitstamp_BCH","https://www.bitstamp.net/api/v2/ticker/bchusd/", "USD", "BCH"}}

	var resseObjectCoinfloorAndBitstamp structs.CoinfloorTickerAndBitstamp
	requestToExchange(resseObjectCoinfloorAndBitstamp, urlList, ch)

	var responseObjectIndependentReserve structs.IndepentReserve
	urlList2 := []Four{
		{"IndependentReserve_BTC", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xbt&secondaryCurrencyCode=aud", "AUD","BTC"},
		{"IndependentReserve_ETH","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=eth&secondaryCurrencyCode=aud", "AUD", "ETH"},
		{"IndependentReserve_BCH","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=bch&secondaryCurrencyCode=aud", "AUD", "BCH"},
		{"IndependentReserve_XRP","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xrp&secondaryCurrencyCode=aud", "AUD", "XRP"},
		{"IndependentReserve_LTC","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=ltc&secondaryCurrencyCode=aud", "AUD", "LTC"}}
	requestToExchange(responseObjectIndependentReserve, urlList2, ch)


	urlList2a := []Four{
		{"GEMINI_BTC", "https://api.gemini.com/v1/pubticker/btcusd", "USD", "BTC"}}
	var responseObjectGeminiBTC structs.GeminiTickerBTC
	//btc, err1 := responseObjectGeminiBTC.RequestUpdate("https://api.gemini.com/v1/pubticker/btcusd")
	requestToExchange(responseObjectGeminiBTC, urlList2a, ch)
	urlList2b := []Four{
		{"GEMINI_ETH", "https://api.gemini.com/v1/pubticker/ethusd", "USD", "ETH"}}
	var responseObjectGeminiETH structs.GeminiTickerETH
	requestToExchange(responseObjectGeminiETH, urlList2b, ch)
	//eth, err2 := responseObjectGeminiETH.RequestUpdate("https://api.gemini.com/v1/pubticker/ethusd")
	//groupList = append(groupList, CryptoDTO{"GEMINI_USD_BTC", btc,err1})
	//groupList = append(groupList, CryptoDTO{"GEMINI_USD_ETH", eth,err2})


	var responseObjectBTC structs.BTCMarket
	urlList3 := []Four{
		{"BTCMarket_AUD_BTC", "https://api.btcmarkets.net/market/BTC/AUD/tick", "AUD", "BTC"},
		{"BTCMarket_AUD_ETH", "https://api.btcmarkets.net/market/ETH/AUD/tick", "AUD", "ETH"},
		{"BTCMarket_AUD_BCH","https://api.btcmarkets.net/market/BCHABC/AUD/tick", "AUD", "BCH"},
		{"BTCMarket_AUD_XRP", "https://api.btcmarkets.net/market/XRP/AUD/tick", "AUD", "XRP"},
		{"BTCMarket_AUD_LTC","https://api.btcmarkets.net/market/LTC/AUD/tick", "AUD", "LTC"}}
	requestToExchange(responseObjectBTC, urlList3, ch)

	var responseObjectACX structs.ACXTicker
	urlList4 := []Four{
		{"ACX_AUD_BTC", "https://acx.io:443/api/v2/tickers/btcaud.json", "AUD", "BTC"},
		{"ACX_AUD_ETH", "https://acx.io:443/api/v2/tickers/ethaud.json", "AUD", "ETH"},
		{"ACX_AUD_BCH","https://acx.io:443/api/v2/tickers/bchaud.json", "AUD", "BCH"},
		{"ACX_AUD_XRP", "https://acx.io:443/api/v2/tickers/ltcaud.json", "AUD", "XRP"},
		{"ACX_AUD_LTC","https://acx.io:443/api/v2/tickers/xrpaud.json","AUD", "LTC"}}
	requestToExchange(responseObjectACX, urlList4, ch)

	var responseObjectCoinjar structs.Coinjar
	urlList5 := []Four{
		{"Coinjar_AUD_BTC", "https://data.exchange.coinjar.com/products/BTCAUD/ticker", "AUD", "BTC"},
		{"Coinjar_AUD_ETH", "https://data.exchange.coinjar.com/products/ETHAUD/ticker", "AUD", "ETH"},
		{"Coinjar_AUD_XRP","https://data.exchange.coinjar.com/products/XRPAUD/ticker","AUD", "XRP"},
		{"Coinjar_AUD_LTC", "https://data.exchange.coinjar.com/products/LTCAUD/ticker", "AUD", "LTC"}}
	requestToExchange(responseObjectCoinjar, urlList5, ch)

	// check for errors
	/*for _, v := range groupList {
		if v.error != nil {
			log.Println(v.name, v.error.Error())
		}
	}*/
	fmt.Println("%.2fs elapsed\n", time.Since(start).Seconds())


	for range urlList {
		// Use the response (<-ch).body
		//val:=<-ch
		//val.error
		log.Println(<-ch)
		//bid() string
		//ask()
	}

	for range urlList2 {
		// Use the response (<-ch).body
		log.Println(<-ch)
	}

	for range urlList2a {
		// Use the response (<-ch).body
		log.Println(<-ch)
	}

	for range urlList2b {
		// Use the response (<-ch).body
		log.Println(<-ch)
	}

	for range urlList3 {
		// Use the response (<-ch).body
		log.Println(<-ch)
	}

	for range urlList4 {
		// Use the response (<-ch).body
		log.Println(<-ch)
	}

	for range urlList5 {
		// Use the response (<-ch).body
		log.Println(<-ch)
	}
	log.Println(<-chRates)
	//val:= <- chRates
	/*if DEBUG {
		log.Println("got values from Crypto exchange")
		log.Println(groupList)
	}*/
	fmt.Println("%.2fs elapsed\n", time.Since(start).Seconds())
	// 14.40s and 8s elapsed before async
	// 1.6s after async



	// send email if arbitage found
}