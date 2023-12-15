package main

import (
	CryptoExchanges "crypto-monitor/main/cryptoExchanges"
	"testing"
	"time"
)

func TestIntCryptoExchanger(t *testing.T) {
	exchangeDataList := ExchangeDataList()
	cryptoExchangeChannel := make(chan CryptoExchanges.CryptoData)
	go func() {
		for _, elem := range exchangeDataList {
			exchangeMutex(elem, cryptoExchangeChannel)
		}
		time.Sleep(10 * time.Second)
		close(cryptoExchangeChannel)
	}()

	for cryptoData := range cryptoExchangeChannel {
		if cryptoData.Error != nil {
			t.Errorf("Didn't expect an error but got one: %v", cryptoData.Error)
		}

		bid, _ := cryptoData.Coin.BidFloat()
		ask, _ := cryptoData.Coin.AskFloat()
		if bid == float64(0) {
			t.Errorf("Bid not defined for %s", cryptoData.Name)
		}
		if ask == float64(0) {
			t.Errorf("Ask not defined for %s", cryptoData.Name)
		}
	}
}
