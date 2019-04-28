package main


// just a scratch
// go run pokemon.go

// implement unmarshal
// https://stackoverflow.com/questions/44380095/polymorphic-json-unmarshalling-of-embedded-structs

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

func geminiBTC(url string) (structs.GeminiTickerBTC, error) {
	var responseObject structs.GeminiTickerBTC
	// also the same for https://api.gemini.com/v1/pubticker/ethusd
	responseData, err := requestWrapper(url)
	if err != nil {
		return responseObject, err
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	return responseObject, nil
}

func geminiETH(url string) (structs.GeminiTickerETH, error) {
	var responseObject structs.GeminiTickerETH
	// also the same for https://api.gemini.com/v1/pubticker/ethusd
	responseData, err := requestWrapper(url)
	if err != nil {
		return responseObject, err
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	return responseObject, nil
}

//type foo struct{
//}

/*func bar(baz interface{}) {
	f, ok := baz.(*foo)
	if !ok {
		// baz was not of type *foo. The assertion failed
	}

	// f is of type *foo
}*/

func IndepentReserve(url string) (structs.IndepentReserve, error) {
	var responseObject structs.IndepentReserve
	responseData, err := requestWrapper(url)
	if err != nil {
		return responseObject, err
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	//f := &structs.IndepentReserve{}

	//bar(f)
	//foo2 := &structs.IndepentReserve{}
	//IndepentReserve2("foo", foo2)

	return responseObject, nil
}

/*func UnmarshalToStruct(url string, exchange structs.CryptoExchange) error {
	//var coin structs.CryptoExchange
	responseData, err := requestWrapper(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(responseData, &exchange)
	if err != nil {
		return err
	}
	log.Println(exchange)

	return nil
}*/

func coinfloorAndBitstamp(url string) (structs.CoinfloorTickerAndBitstamp, error) {
	var responseObject structs.CoinfloorTickerAndBitstamp
	responseData, err := requestWrapper(url)
	if err != nil {
		return responseObject, err
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	return responseObject, nil
}

func BTCMarket(url string) (structs.BTCMarket, error) {
	var responseObject structs.BTCMarket
	responseData, err := requestWrapper(url)
	if err != nil {
		return responseObject, err
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	return responseObject, nil
}

func ACX(url string) (structs.ACXTicker, error) {
	var responseObject structs.ACXTicker
	responseData, err := requestWrapper(url)
	if err != nil {
		return responseObject, err
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	return responseObject, nil
}

func Coinjar(url string) (structs.Coinjar, error) {
	var responseObject structs.Coinjar
	responseData, err := requestWrapper(url)
	if err != nil {
		return responseObject, err
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	return responseObject, nil
}

/*func cryptoCurrencies() (map[string]int, error) {
	exchangeMap := make(map[string]map[string]int) // map[string]int or error ??
	gemini()
}*/

func currencyExchangeRates() (map[string]decimal.Decimal, error) {
	exchangeMap := make(map[string]decimal.Decimal)

	resp, err := http.Get("https://api.exchangeratesapi.io/latest?base=USD")
	// same for https://api.gemini.com/v1/pubticker/ethusd
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


func main() {
	start := time.Now()
	var responseObject2 structs.IndepentReserve

	blah, err1 := responseObject2.RequestUpdate("https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xbt&secondaryCurrencyCode=aud")
	if err1 != nil {
		log.Println("rarr")
	}

	log.Println(blah)

	//CryptoExchange
	//err := UnmarshalToStruct("https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xbt&secondaryCurrencyCode=aud", responseObject2) // (interface{}, error) {
	//var foo := CryptoDTO{"hi",*foo2, err}

	//foo4, err3 := something2("https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xbt&secondaryCurrencyCode=aud", responseObject2)
	/*if err != nil {
		log.Fatal(err.Error())
	}*/
	log.Println(responseObject2)

	DEBUG := false
	//val1, err1 := cryptoCurrencies()
	val, err := currencyExchangeRates()
	groupList := []CryptoDTO{}
	ch := make(chan CryptoDTO)
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


	for _, v := range urlList {
		val , err := coinfloorAndBitstamp(v.url)
		ch <- CryptoDTO{v.name,val, err}
		groupList = append(groupList, CryptoDTO{v.name,val, err})
	}

	urlList2 := []Pair{
		{"IndependentReserve_AUD_BTC", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xbt&secondaryCurrencyCode=aud"},
		{"IndependentReserve_AUD_ETH","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=eth&secondaryCurrencyCode=aud"},
		{"IndependentReserve_AUD_BCH","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=bch&secondaryCurrencyCode=aud"},
		{"IndependentReserve_AUD_XRP","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xrp&secondaryCurrencyCode=aud"},
		{"IndependentReserve_AUD_LTC","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=ltc&secondaryCurrencyCode=aud"}}
	for _, v := range urlList2 {
		val, err := IndepentReserve(v.url)
		groupList = append(groupList, CryptoDTO{v.name,val, err})
	}

	btc, err1 := geminiBTC("https://api.gemini.com/v1/pubticker/btcusd")
	eth, err2 := geminiETH("https://api.gemini.com/v1/pubticker/ethusd")
	groupList = append(groupList, CryptoDTO{"GEMINI_USD_BTC", btc,err1})
	groupList = append(groupList, CryptoDTO{"GEMINI_USD_ETH", eth,err2})

	urlList3 := []Pair{
		{"BTCMarket_AUD_BTC", "https://api.btcmarkets.net/market/BTC/AUD/tick"},
		{"BTCMarket_AUD_ETH", "https://api.btcmarkets.net/market/ETH/AUD/tick"},
		{"BTCMarket_AUD_BCH","https://api.btcmarkets.net/market/BCHABC/AUD/tick"},
		{"BTCMarket_AUD_XRP", "https://api.btcmarkets.net/market/XRP/AUD/tick"},
		{"BTCMarket_AUD_LTC","https://api.btcmarkets.net/market/LTC/AUD/tick"}}
	for _, v := range urlList3 {
		// go routine goes here
		val , err := BTCMarket(v.url)
		groupList = append(groupList, CryptoDTO{v.name,val, err})
	}


	urlList4 := []Pair{
		{"ACX_AUD_BTC", "https://acx.io:443/api/v2/tickers/btcaud.json"},
		{"ACX_AUD_ETH", "https://acx.io:443/api/v2/tickers/ethaud.json"},
		{"ACX_AUD_BCH","https://acx.io:443/api/v2/tickers/bchaud.json"},
		{"ACX_AUD_XRP", "https://acx.io:443/api/v2/tickers/ltcaud.json"},
		{"ACX_AUD_LTC","https://acx.io:443/api/v2/tickers/xrpaud.json"}}
	for _, v := range urlList4 {
		val , err := ACX(v.url)
		groupList = append(groupList, CryptoDTO{v.name,val, err})
	}

	urlList5 := []Pair{
		{"Coinjar_AUD_BTC", "https://data.exchange.coinjar.com/products/BTCAUD/ticker"},
		{"Coinjar_AUD_ETH", "https://data.exchange.coinjar.com/products/ETHAUD/ticker"},
		{"Coinjar_AUD_XRP","https://data.exchange.coinjar.com/products/XRPAUD/ticker"},
		{"Coinjar_AUD_LTC", "https://data.exchange.coinjar.com/products/LTCAUD/ticker"}}
	for _, v := range urlList5 {
		val , err := Coinjar(v.url)
		groupList = append(groupList, CryptoDTO{v.name,val, err})
	}

	// check for errors
	for _, v := range groupList {
		if v.error != nil {
			log.Println(v.error.Error())
		}
	}
	//log.Println(btc)
	//log.Println(eth)
	if DEBUG {
		log.Println("got values from Crypto exchange")
		log.Println(groupList)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	// 14.40s elapsed before async



	// send email if arbitage found
}