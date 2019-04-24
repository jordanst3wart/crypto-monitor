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

func geminiBTC() (structs.GeminiTickerBTC, error) {
	var responseObject structs.GeminiTickerBTC
	resp, err := http.Get("https://api.gemini.com/v1/pubticker/btcusd")
	// also the same for https://api.gemini.com/v1/pubticker/ethusd
	if err != nil {
		return responseObject, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return responseObject, err
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	return responseObject, nil
}

func geminiETH() (structs.GeminiTickerETH, error) {
	var responseObject structs.GeminiTickerETH
	resp, err := http.Get("https://api.gemini.com/v1/pubticker/ethusd")
	// also the same for https://api.gemini.com/v1/pubticker/ethusd
	if err != nil {
		return responseObject, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return responseObject, err
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	return responseObject, nil
}

func coinfloor(url string) (structs.CoinfloorTicker, error) {
	var responseObject structs.CoinfloorTicker
	resp, err := http.Get(url)
	if err != nil {
		return responseObject, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return responseObject, err
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	return responseObject, nil
}

func coinfloorAndBitstamp(url string) (structs.CoinfloorTickerAndBitstamp, error) {
	var responseObject structs.CoinfloorTickerAndBitstamp
	resp, err := http.Get(url)
	if err != nil {
		return responseObject, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
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

func main() {
	//val1, err1 := cryptoCurrencies()
	val, err := currencyExchangeRates()
	// log error at main level
	if err != nil {
		// possibly send email
		log.Fatal(err.Error())
	}
	log.Println(val)
	//btc, err := geminiBTC()
	//eth, err := geminiETH()
	btc, err := coinfloorAndBitstamp("https://webapi.coinfloor.co.uk:8090/bist/XBT/GBP/ticker/")
	eth, err := coinfloorAndBitstamp("https://webapi.coinfloor.co.uk:8090/bist/ETH/GBP/ticker/")
	bch, err := coinfloorAndBitstamp("https://webapi.coinfloor.co.uk:8090/bist/BCH/GBP/ticker/")
	bch2, err := coinfloorAndBitstamp("https://www.bitstamp.net/api/v2/ticker/btcusd/")


	log.Println(btc)
	log.Println(eth)
	log.Println(bch)


	// send email if arbitage found
}