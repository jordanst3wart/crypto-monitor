package fiatCurrencyExchange

import (
	"crypto-monitor/main/cryptoExchanges"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type ExchangeClient interface {
	getRates() (resp *http.Response, err error)
}

type RealExchangeClient struct{}

type ExchangeRates struct {
	Rates map[string]float64
	Err   error
}

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

func FiatCurrencyExchangeRates(ch chan ExchangeRates, updateFrequency time.Duration, client ExchangeClient) {
	exchangeMap := make(map[string]float64)
	var responseObject CryptoExchanges.CurrencyExchangeAPI
	for {
		resp, err := client.getRates()
		if err != nil {
			log.Println("Failed to send http request to Exchange client")
			// should wrap error
			ch <- ExchangeRates{nil, err}
			time.Sleep(1 * time.Minute)
			continue
		}

		responseData, err := io.ReadAll(resp.Body)
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
