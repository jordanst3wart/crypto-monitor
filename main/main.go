package main

import (
	"crypto-monitor/main/cryptoExchanges"
	"crypto-monitor/main/fiatCurrencyExchange"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type ArbitrageData struct {
	name, crypto string
	arb          float64
}

// TODO Use orders to determine the volume of available & log amount of volume is available to buy/sell
// TODO logrus to have better debug logic

func main() {
	minimumArbitrageRatio := 1.017
	exchangeDataList := ExchangeDataList()

	// log setup
	os.Setenv("TZ", "Australia/Sydney")
	log.SetOutput(os.Stdout)
	log.Println("Starting log...")

	fiatRatesChannel := make(chan fiatCurrencyExchange.ExchangeRates)

	// get exchange rates to start
	go fiatCurrencyExchange.FiatExchangeRatesRoutine(fiatRatesChannel, 60*time.Minute, fiatCurrencyExchange.RealExchangeClient{})
	fiatMsg := <-fiatRatesChannel
	if fiatMsg.Err != nil {
		log.Fatalf("Fiat exchange error on start up: %v \n", fiatMsg.Err)
	}

	cryptoExchangeChannel := make(chan CryptoExchanges.CryptoData)

	for {
		go updateFiatRates(fiatRatesChannel, fiatMsg)

		for _, elem := range exchangeDataList {
			exchangeMutex(elem, cryptoExchangeChannel)
		}

		var aggregatedCryptoData []CryptoExchanges.CryptoData
		for _, exchangeApp := range exchangeDataList {
			for range exchangeApp.cryptoList {
				cryptoData := <-cryptoExchangeChannel
				if cryptoData.Error != nil {
					log.Printf("Exchange App: %s, Error: %v, Coin: %v\n", cryptoData.Name, cryptoData.Error, cryptoData.Coin)
				} else {
					aggregatedCryptoData = append(aggregatedCryptoData, ConvertCurrency(cryptoData, fiatMsg))
				}
			}
		}

		// sort cryptos by crypto-currency
		var set []string
		for _, item := range aggregatedCryptoData {
			set = append(set, item.Crypto)
		}
		uniqueCryptos := DeduplicateStrings(set)
		mapCrypto := map[string][]CryptoExchanges.CryptoData{}
		for i := range uniqueCryptos {
			var listCrypto []CryptoExchanges.CryptoData
			for _, item := range aggregatedCryptoData {
				if item.Crypto == uniqueCryptos[i] {
					listCrypto = append(listCrypto, item)
				}
			}
			mapCrypto[uniqueCryptos[i]] = listCrypto
		}

		var arbitrageList []ArbitrageData
		for _, cryptoList := range mapCrypto {
			for _, itemOuter := range cryptoList {
				for _, itemInner := range cryptoList {
					arb := CheckArbitrage(itemInner, itemOuter)
					bid, _ := itemInner.Coin.BidFloat()
					ask, _ := itemOuter.Coin.AskFloat()
					arbitrageList = append(arbitrageList, ArbitrageData{fmt.Sprintf("buy: %s at %v, sell: %s at %v", itemOuter.Name, ask, itemInner.Name, bid), itemOuter.Crypto, arb})
				}
			}
		}

		for _, item := range arbitrageList {
			if item.arb > minimumArbitrageRatio {
				val := strconv.FormatFloat(item.arb, 'f', -1, 64)
				log.Printf("ARBITRAGE on %s at %v\n", item.name, val)
			}
		}

		// biggest limit seen is 1 call per second
		time.Sleep(time.Second * 25)
	}
}

func updateFiatRates(fiatRatesChannel chan fiatCurrencyExchange.ExchangeRates, fiatMsg fiatCurrencyExchange.ExchangeRates) {
	func() {
		select {
		case fiatRateMsg := <-fiatRatesChannel:
			if fiatRateMsg.Err != nil {
				log.Printf("Fiat exchange error: %v \n", fiatRateMsg.Err)
			} else {
				fiatMsg = fiatRateMsg
			}
		}
	}()
}
