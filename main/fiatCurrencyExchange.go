package main

import (
	"crypto-monitor/structs"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ExchangeClient interface {
	getRates() (resp *http.Response, err error)
}

type MockErrorExchangeClient struct{}

//type MockExchangeClient struct{}

type RealExchangeClient struct{}

func (e MockErrorExchangeClient) getRates() (resp *http.Response, err error) {
	// return json scaffold
	return &http.Response{Body: nil}, errors.New("some error")
}

type ExchangeRates struct {
	rates map[string]float64
	err   error
}

/*func (e MockExchangeClient) getRates() (resp *http.Response, err error) {
	// return json scaffold
	return &http.Response{Body: nil}, nil
}*/

/*
	// 1000 Monthly Requests which updates hourly
	// free access key 8c0b46946137202ce9e9152e255f7b52
	// can't use USD, and can't use https
	// http://api.exchangeratesapi.io/latest?&base=EUR&access_key=8c0b46946137202ce9e9152e255f7b52
*/
func (e RealExchangeClient) getRates() (resp *http.Response, err error) {
	return http.Get("http://api.exchangeratesapi.io/latest?&base=EUR&access_key=8c0b46946137202ce9e9152e255f7b52")
}

/*
	Gets exchange rate, every 5 minutes
*/

func fiatCurrencyExchangeRates(ch chan ExchangeRates, updateFrequency time.Duration, client ExchangeClient) {
	exchangeMap := make(map[string]float64)
	var responseObject structs.CurrencyExchangeAPI
	// TODO need to cache
	// TODO does the 'return' statement return out of the for loop?
	for {
		resp, err := client.getRates()
		if err != nil {
			log.Println("Failed to send http request to Exchange client")
			// should wrap error
			ch <- ExchangeRates{nil, err}
			time.Sleep(1 * time.Minute)
			continue
		}
		// https://fixer.io/product

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Failed to read message from Exchange client")
			ch <- ExchangeRates{nil, err}
			time.Sleep(1 * time.Minute)
			continue
		}
		err = json.Unmarshal(responseData, &responseObject)
		if err != nil {
			log.Println("Failed to parse message from Exchange client")
			ch <- ExchangeRates{nil, err}
			time.Sleep(1 * time.Minute)
			continue
		} else {
			exchangeMap["USD2AUD"] = responseObject.Rates.AUD / responseObject.Rates.USD
			exchangeMap["GBP2AUD"] = responseObject.Rates.AUD / responseObject.Rates.GBP
			ch <- ExchangeRates{exchangeMap, err}
		}
		time.Sleep(updateFrequency)
	}
}
