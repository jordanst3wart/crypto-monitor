package main


// just a scratch
// go run pokemon.go

import (
	"./structs"
	"encoding/json"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"log"
	"net/http"
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

func geminiBTC(url string) (structs.GenericCryptoResponse, error) {
	var responseObject structs.GeminiTickerBTC
	var object structs.GenericCryptoResponse
	// also the same for https://api.gemini.com/v1/pubticker/ethusd
	responseData, err := requestWrapper(url)
	if err != nil {
		return object, err
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return object, err
	}

	object.Volume = responseObject.Volume
	object.Ask = responseObject.Ask
	object.Bid = responseObject.Bid
	object.High = responseObject.High
	object.Low = responseObject.Low
	object.Last = responseObject.Last

	return object, nil
}

func geminiETH(url string) (structs.GenericCryptoResponse, error) {
	var responseObject structs.GeminiTickerETH
	var object structs.GenericCryptoResponse
	// also the same for https://api.gemini.com/v1/pubticker/ethusd
	responseData, err := requestWrapper(url)
	if err != nil {
		return object, err
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return object, err
	}

	object.Volume = responseObject.Volume
	object.Ask = responseObject.Ask
	object.Bid = responseObject.Bid
	object.High = responseObject.High
	object.Low = responseObject.Low
	object.Last = responseObject.Last

	return object, nil
}

func IndepentReserve(url string) (structs.GenericCryptoResponse, error) {
	var responseObject structs.IndepentReserve
	var object structs.GenericCryptoResponse
	responseData, err := requestWrapper(url)
	if err != nil {
		return object, err
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return object, err
	}

	object.Volume = responseObject.Volume
	object.Ask = responseObject.Ask
	object.Bid = responseObject.Bid
	object.High = responseObject.High
	object.Low = responseObject.Low
	object.Last = responseObject.Last

	return object, nil
}

func coinfloorAndBitstamp(url string) (structs.GenericCryptoResponse, error) {
	var responseObject structs.CoinfloorTickerAndBitstamp
	var object structs.GenericCryptoResponse
	responseData, err := requestWrapper(url)
	if err != nil {
		return object, err
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return object, err
	}

	object.Volume = responseObject.Volume
	object.Ask = responseObject.Ask
	object.Bid = responseObject.Bid
	object.High = responseObject.High
	object.Low = responseObject.Low
	object.Last = responseObject.Last

	return object, nil
}

func BTCMarket(url string) (structs.GenericCryptoResponse, error) {
	var responseObject structs.BTCMarket
	var object structs.GenericCryptoResponse
	responseData, err := requestWrapper(url)
	if err != nil {
		return object, err
	}
	//var A structs.CryptoExchange


	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return object, err
	}

	object.Volume = responseObject.Volume
	object.Ask = responseObject.Ask
	object.Bid = responseObject.Bid
	object.High = responseObject.High
	object.Low = responseObject.Low
	object.Last = responseObject.Last

	return object, nil
}

func genericCryptoExchange(url string) (structs.GenericCryptoResponse, error) {
	var responseObject structs.BTCMarket
	var object structs.GenericCryptoResponse
	responseData, err := requestWrapper(url)
	if err != nil {
		return object, err
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return object, err
	}

	object.Volume = responseObject.Volume
	object.Ask = responseObject.Ask
	object.Bid = responseObject.Bid
	object.High = responseObject.High
	object.Low = responseObject.Low
	object.Last = responseObject.Last

	return object, nil
}


/*func cryptoCurrencies() (map[string]int, error) {
	exchangeMap := make(map[string]map[string]int) // map[string]int or error ??
	gemini()
}*/

func currencyExchangeRates() (map[string]decimal.Decimal, error) {
	exchangeMap := make(map[string]decimal.Decimal)

	resp, err := http.Get("https://api.exchangeratesapi.io/latest?base=USD")
	// same for https://api.gemini.com/v1/pubticker/ethusd
	// should use panic... and defer recover to deal with it; in the future...
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

// https://api.exchangeratesapi.io/latest?base=USD

type Pair struct {
	name, url string
}

type CryptoDTO struct {
	name string
	coin structs.GenericCryptoResponse
	error error
}

func main() {
	DEBUG := false
	//val1, err1 := cryptoCurrencies()
	val, err := currencyExchangeRates()
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
	groupListBitstampAndCoinfloor := []CryptoDTO{}
	for _, v := range urlList {
		val , err := coinfloorAndBitstamp(v.url)
		groupListBitstampAndCoinfloor = append(groupListBitstampAndCoinfloor, CryptoDTO{v.name,val, err})
	}

	urlList2 := []Pair{
		{"IndependentReserve_AUS_BTC", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xbt&secondaryCurrencyCode=aud"},
		{"IndependentReserve_AUS_ETH","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=eth&secondaryCurrencyCode=aud"},
		{"IndependentReserve_AUS_BCH","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=bch&secondaryCurrencyCode=aud"},
		{"IndependentReserve_AUS_XRP","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xrp&secondaryCurrencyCode=aud"},
		{"IndependentReserve_AUS_LTC","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=ltc&secondaryCurrencyCode=aud"}}
	groupListIndependentReserve := []CryptoDTO{}
	for _, v := range urlList2 {
		val , err := IndepentReserve(v.url)
		groupListIndependentReserve = append(groupListIndependentReserve, CryptoDTO{v.name,val, err})
	}

	btc, err1 := geminiBTC("https://api.gemini.com/v1/pubticker/btcusd")
	eth, err2 := geminiETH("https://api.gemini.com/v1/pubticker/ethusd")
	groupListGemini := []CryptoDTO{}
	groupListGemini = append(groupListGemini, CryptoDTO{"GEMINI_USD_BTC", btc,err1})
	groupListGemini = append(groupListGemini, CryptoDTO{"GEMINI_USD_ETH", eth,err2})

	urlList3 := []Pair{
		{"BTCMarket_AUS_BTC", "https://api.btcmarkets.net/market/BTC/AUD/tick"},
		{"BTCMarket_AUS_ETH", "https://api.btcmarkets.net/market/ETH/AUD/tick"},
		{"BTCMarket_AUS_BCH","https://api.btcmarkets.net/market/BCHABC/AUD/tick"},
		{"BTCMarket_AUS_XRP", "https://api.btcmarkets.net/market/XRP/AUD/tick"},
		{"BTCMarket_AUS_LTC","https://api.btcmarkets.net/market/LTC/AUD/tick"}}
	groupListBTCMarket := []CryptoDTO{}
	for _, v := range urlList3 {
		val , err := BTCMarket(v.url)
		groupListBTCMarket = append(groupListBTCMarket, CryptoDTO{v.name,val, err})
	}

	// create a list

	/*btc, err := coinfloorAndBitstamp("https://webapi.coinfloor.co.uk:8090/bist/XBT/GBP/ticker/")
	eth, err := coinfloorAndBitstamp("https://webapi.coinfloor.co.uk:8090/bist/ETH/GBP/ticker/")
	bch, err := coinfloorAndBitstamp("https://webapi.coinfloor.co.uk:8090/bist/BCH/GBP/ticker/")
	bch2, err := coinfloorAndBitstamp("https://www.bitstamp.net/api/v2/ticker/btcusd/")
*/

	//log.Println(btc)
	//log.Println(eth)
	if DEBUG {
		log.Println("got values from Crypto exchange")
		log.Println(groupListBitstampAndCoinfloor)
	}



	// send email if arbitage found
}