package main


// go run pokemon.go

import (
	"crypto-monitor/structs"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


// TODO make http requests async
// TODO simplify

func currencyExchangeRates() (map[string]decimal.Decimal, error) {
	exchangeMap := make(map[string]decimal.Decimal)

	resp, err := http.Get("https://api.exchangeratesapi.io/latest?base=USD")
	if err != nil {
		// should wrap error
		return nil, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var responseObject structs.CurrencyExchangeAPI
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}
	exchangeMap["USD2AUD"] = responseObject.Rates.AUD.Div(responseObject.Rates.USD)
	exchangeMap["GBP2AUD"] = responseObject.Rates.AUD.Div(responseObject.Rates.GDP)
	return exchangeMap, err
}

type Pair struct {
	name, url string
}

func requestToExchange(exchange structs.CryptoExchange, urlList []Pair, ch chan structs.CryptoDTO){
	for _, v := range urlList {
		go exchange.RequestUpdate(v.name, v.url, ch)
		//ch<-CryptoDTO{v.name,val, err}
		//groupList = append(groupList, CryptoDTO{v.name,val, err})
	}
	//return groupList
}


func main() {
	start := time.Now()

	DEBUG := true
	val, err := currencyExchangeRates()
	//groupList := []structs.CryptoDTO{}
	ch := make(chan structs.CryptoDTO)
	// log error at main level
	if err != nil {
		// possibly send email
		log.Fatal(err.Error())
	}
	if DEBUG {
		log.Println("Got values from exchange")
		log.Println(val)
	}

	urlList := []Pair{
		{"Coinfloor_GBP_BTC", "https://webapi.coinfloor.co.uk:8090/bist/XBT/GBP/ticker/"},
		{"Coinfloor_GBP_ETH","https://webapi.coinfloor.co.uk:8090/bist/ETH/GBP/ticker/"},
		{"Coinfloor_GBP_BCH","https://webapi.coinfloor.co.uk:8090/bist/BCH/GBP/ticker/"},
		{"Coinfloor_GBP_BCH","https://webapi.coinfloor.co.uk:8090/bist/BCH/GBP/ticker/"},
		{"Bitstamp_USD_BTC","https://www.bitstamp.net/api/v2/ticker/btcusd/"},
		{"Bitstamp_USD_XRP","https://www.bitstamp.net/api/v2/ticker/xrpusd/"},
		{"Bitstamp_USD_LTC","https://www.bitstamp.net/api/v2/ticker/ltcusd/"},
		{"Bitstamp_USD_ETH","https://www.bitstamp.net/api/v2/ticker/ethusd/"},
		{"Bitstamp_USD_BCH","https://www.bitstamp.net/api/v2/ticker/bchusd/"}}

	var resseObjectCoinfloorAndBitstamp structs.CoinfloorTickerAndBitstamp
	requestToExchange(resseObjectCoinfloorAndBitstamp, urlList, ch)

	var responseObjectIndependentReserve structs.IndepentReserve
	urlList2 := []Pair{
		{"IndependentReserve_AUD_BTC", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xbt&secondaryCurrencyCode=aud"},
		{"IndependentReserve_AUD_ETH","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=eth&secondaryCurrencyCode=aud"},
		{"IndependentReserve_AUD_BCH","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=bch&secondaryCurrencyCode=aud"},
		{"IndependentReserve_AUD_XRP","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xrp&secondaryCurrencyCode=aud"},
		{"IndependentReserve_AUD_LTC","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=ltc&secondaryCurrencyCode=aud"}}
	requestToExchange(responseObjectIndependentReserve, urlList2, ch)


	urlList2a := []Pair{
		{"GEMINI_USD_BTC", "https://api.gemini.com/v1/pubticker/btcusd"}}
	var responseObjectGeminiBTC structs.GeminiTickerBTC
	//btc, err1 := responseObjectGeminiBTC.RequestUpdate("https://api.gemini.com/v1/pubticker/btcusd")
	requestToExchange(responseObjectGeminiBTC, urlList2a, ch)
	urlList2b := []Pair{
		{"GEMINI_USD_ETH", "https://api.gemini.com/v1/pubticker/ethusd"}}
	var responseObjectGeminiETH structs.GeminiTickerETH
	requestToExchange(responseObjectGeminiETH, urlList2b, ch)
	//eth, err2 := responseObjectGeminiETH.RequestUpdate("https://api.gemini.com/v1/pubticker/ethusd")
	//groupList = append(groupList, CryptoDTO{"GEMINI_USD_BTC", btc,err1})
	//groupList = append(groupList, CryptoDTO{"GEMINI_USD_ETH", eth,err2})


	var responseObjectBTC structs.BTCMarket
	urlList3 := []Pair{
		{"BTCMarket_AUD_BTC", "https://api.btcmarkets.net/market/BTC/AUD/tick"},
		{"BTCMarket_AUD_ETH", "https://api.btcmarkets.net/market/ETH/AUD/tick"},
		{"BTCMarket_AUD_BCH","https://api.btcmarkets.net/market/BCHABC/AUD/tick"},
		{"BTCMarket_AUD_XRP", "https://api.btcmarkets.net/market/XRP/AUD/tick"},
		{"BTCMarket_AUD_LTC","https://api.btcmarkets.net/market/LTC/AUD/tick"}}
	requestToExchange(responseObjectBTC, urlList3, ch)

	var responseObjectACX structs.ACXTicker
	urlList4 := []Pair{
		{"ACX_AUD_BTC", "https://acx.io:443/api/v2/tickers/btcaud.json"},
		{"ACX_AUD_ETH", "https://acx.io:443/api/v2/tickers/ethaud.json"},
		{"ACX_AUD_BCH","https://acx.io:443/api/v2/tickers/bchaud.json"},
		{"ACX_AUD_XRP", "https://acx.io:443/api/v2/tickers/ltcaud.json"},
		{"ACX_AUD_LTC","https://acx.io:443/api/v2/tickers/xrpaud.json"}}
	requestToExchange(responseObjectACX, urlList4, ch)

	var responseObjectCoinjar structs.Coinjar
	urlList5 := []Pair{
		{"Coinjar_AUD_BTC", "https://data.exchange.coinjar.com/products/BTCAUD/ticker"},
		{"Coinjar_AUD_ETH", "https://data.exchange.coinjar.com/products/ETHAUD/ticker"},
		{"Coinjar_AUD_XRP","https://data.exchange.coinjar.com/products/XRPAUD/ticker"},
		{"Coinjar_AUD_LTC", "https://data.exchange.coinjar.com/products/LTCAUD/ticker"}}
	requestToExchange(responseObjectCoinjar, urlList5, ch)

	// check for errors
	/*for _, v := range groupList {
		if v.error != nil {
			log.Println(v.name, v.error.Error())
		}
	}*/

	for range urlList {
		// Use the response (<-ch).body
		fmt.Println(<-ch)
	}

	for range urlList2 {
		// Use the response (<-ch).body
		fmt.Println(<-ch)
	}

	for range urlList2a {
		// Use the response (<-ch).body
		fmt.Println(<-ch)
	}

	for range urlList2b {
		// Use the response (<-ch).body
		fmt.Println(<-ch)
	}

	for range urlList3 {
		// Use the response (<-ch).body
		fmt.Println(<-ch)
	}

	for range urlList4 {
		// Use the response (<-ch).body
		fmt.Println(<-ch)
	}

	for range urlList5 {
		// Use the response (<-ch).body
		fmt.Println(<-ch)
	}
	/*if DEBUG {
		log.Println("got values from Crypto exchange")
		log.Println(groupList)
	}*/
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	// 14.40s and 8s elapsed before async



	// send email if arbitage found
}