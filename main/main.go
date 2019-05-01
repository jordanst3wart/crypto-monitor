package main


// go run pokemon.go

import (
	"./structs"
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

func requestWrapper(url string) ([]byte, error) {
	var responseData []byte
	resp, err := http.Get(url)
	if err != nil {
		return responseData, err
	}

	return ioutil.ReadAll(resp.Body)
}

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

type CryptoDTO struct {
	name string
	coin structs.CryptoExchange
	error error
}

func requestToExchange(exchange structs.CryptoExchange, urlList []Pair, groupList []CryptoDTO) ([]CryptoDTO){
	for _, v := range urlList {
		val, err := exchange.RequestUpdate(v.url)
		groupList = append(groupList, CryptoDTO{v.name,val, err})
	}
	return groupList
}


func main() {
	start := time.Now()

	DEBUG := true
	val, err := currencyExchangeRates()
	groupList := []CryptoDTO{}
	// ch := make(chan CryptoDTO)
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
	groupList = requestToExchange(resseObjectCoinfloorAndBitstamp, urlList, groupList)

	var responseObjectIndependentReserve structs.IndepentReserve
	urlList2 := []Pair{
		{"IndependentReserve_AUD_BTC", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xbt&secondaryCurrencyCode=aud"},
		{"IndependentReserve_AUD_ETH","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=eth&secondaryCurrencyCode=aud"},
		{"IndependentReserve_AUD_BCH","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=bch&secondaryCurrencyCode=aud"},
		{"IndependentReserve_AUD_XRP","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xrp&secondaryCurrencyCode=aud"},
		{"IndependentReserve_AUD_LTC","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=ltc&secondaryCurrencyCode=aud"}}
	groupList = requestToExchange(responseObjectIndependentReserve, urlList2, groupList)


	var responseObjectGeminiBTC structs.GeminiTickerBTC
	btc, err1 := responseObjectGeminiBTC.RequestUpdate("https://api.gemini.com/v1/pubticker/btcusd")
	var responseObjectGeminiETH structs.GeminiTickerETH
	eth, err2 := responseObjectGeminiETH.RequestUpdate("https://api.gemini.com/v1/pubticker/ethusd")
	groupList = append(groupList, CryptoDTO{"GEMINI_USD_BTC", btc,err1})
	groupList = append(groupList, CryptoDTO{"GEMINI_USD_ETH", eth,err2})


	var responseObjectBTC structs.BTCMarket
	urlList3 := []Pair{
		{"BTCMarket_AUD_BTC", "https://api.btcmarkets.net/market/BTC/AUD/tick"},
		{"BTCMarket_AUD_ETH", "https://api.btcmarkets.net/market/ETH/AUD/tick"},
		{"BTCMarket_AUD_BCH","https://api.btcmarkets.net/market/BCHABC/AUD/tick"},
		{"BTCMarket_AUD_XRP", "https://api.btcmarkets.net/market/XRP/AUD/tick"},
		{"BTCMarket_AUD_LTC","https://api.btcmarkets.net/market/LTC/AUD/tick"}}
	groupList = requestToExchange(responseObjectBTC, urlList3, groupList)

	var responseObjectACX structs.ACXTicker
	urlList4 := []Pair{
		{"ACX_AUD_BTC", "https://acx.io:443/api/v2/tickers/btcaud.json"},
		{"ACX_AUD_ETH", "https://acx.io:443/api/v2/tickers/ethaud.json"},
		{"ACX_AUD_BCH","https://acx.io:443/api/v2/tickers/bchaud.json"},
		{"ACX_AUD_XRP", "https://acx.io:443/api/v2/tickers/ltcaud.json"},
		{"ACX_AUD_LTC","https://acx.io:443/api/v2/tickers/xrpaud.json"}}
	groupList = requestToExchange(responseObjectACX, urlList4, groupList)

	var responseObjectCoinjar structs.Coinjar
	urlList5 := []Pair{
		{"Coinjar_AUD_BTC", "https://data.exchange.coinjar.com/products/BTCAUD/ticker"},
		{"Coinjar_AUD_ETH", "https://data.exchange.coinjar.com/products/ETHAUD/ticker"},
		{"Coinjar_AUD_XRP","https://data.exchange.coinjar.com/products/XRPAUD/ticker"},
		{"Coinjar_AUD_LTC", "https://data.exchange.coinjar.com/products/LTCAUD/ticker"}}
	groupList = requestToExchange(responseObjectCoinjar, urlList5, groupList)

	// check for errors
	for _, v := range groupList {
		if v.error != nil {
			log.Println(v.name, v.error.Error())
		}
	}

	if DEBUG {
		log.Println("got values from Crypto exchange")
		log.Println(groupList)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	// 14.40s and 8s elapsed before async



	// send email if arbitage found
}