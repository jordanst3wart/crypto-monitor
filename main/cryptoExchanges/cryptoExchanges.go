package CryptoExchanges

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	bid_aud
	offer_aud
*/

// the `json:""` are called struct tags
// by default the json is parsed input the matching key
// currency struct
type CurrencyExchangeAPI struct {
	Base  string           `json:"base"`
	Rates CurrencyExchange `json:"rates"`
	Date  string           `json:"date"`
}

type CurrencyExchange struct {
	GBP float64 `json:"GBP"`
	AUD float64 `json:"AUD"`
	EUR float64 `json:"EUR"`
	USD float64 `json:"USD"`
}

type CryptoData struct {
	Name     string
	Coin     CryptoExchange
	Error    error
	Currency string
	Crypto   string
}

// TODO put logic in here for RequestUpdate to reduce repeating code
func requestWrapper(url string) ([]byte, error) {
	var responseData []byte

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("req failed for %s error is %v", url, err)
	}

	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return responseData, fmt.Errorf("get request failed for URL: %s, error: %w ", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return responseData, fmt.Errorf("non-OK HTTP status received for URL %s: %d", url, resp.StatusCode)
	}

	responseData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return responseData, fmt.Errorf("error reading response data for URL: %s, error: %d", url, err)
	}

	return responseData, nil
}

type CryptoExchange interface {
	VolumeFloat() (float64, error)
	AskFloat() (float64, error)
	BidFloat() (float64, error)
	HighFloat() (float64, error)
	LowFloat() (float64, error)
	LastFloat() (float64, error)
	RequestUpdate(name string, url string, ch chan CryptoData, currency string, crypto string)
	// RequestUpdate(v.name, v.url, ch, v.currency, v.crypto)
}
