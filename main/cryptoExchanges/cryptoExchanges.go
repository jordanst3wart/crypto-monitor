package CryptoExchanges

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type CryptoData struct {
	Name     string
	Coin     CryptoExchange
	Error    error
	Currency string
	Crypto   string
}

func requestWrapper(url string) ([]byte, error) {
	var responseData []byte

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("req failed for %s error is %v", url, err)
	}
	// TODO do a retry if a 500 code in x number of seconds, and maybe do exponential backoff maxing out at 5 minutes

	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return responseData, fmt.Errorf("get request failed for URL: %s, error: %w ", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return responseData, fmt.Errorf("non-OK HTTP status received for URL %s: %d", url, resp.StatusCode)
	}

	responseData, err = io.ReadAll(resp.Body)
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
