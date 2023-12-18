package CryptoExchanges

import (
	"fmt"
	"io"
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

func processResponse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK HTTP status received: %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

func exponentialBackoffRequest(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("request creation failed: %w", err)
	}
	req.Header.Add("Accept", "application/json")

	retryCount := 0
	maxRetries := 3
	backoff := 1 * time.Second
	maxBackoff := 1 * time.Minute

	for {
		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("request failed: %w", err)
		}

		responseData, err := processResponse(resp)
		if err == nil {
			return responseData, nil
		}

		if resp.StatusCode >= 500 && resp.StatusCode <= 599 && retryCount < maxRetries {
			time.Sleep(backoff)
			backoff *= 2
			if backoff > maxBackoff {
				backoff = maxBackoff
			}
			retryCount++
		} else {
			return nil, err
		}
	}
}

func requestWrapper(url string) ([]byte, error) {
	return exponentialBackoffRequest(url)
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
