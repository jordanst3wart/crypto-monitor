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

// TODO Use orders to determine the volume of available & log amount of volume is available to buy/sell
// TODO logrus to have better debug logic

func main() {
	minimumArbitrageRatio := 1.01
	exchangeDataList := ExchangeDataList()

	// log setup
	log.SetOutput(os.Stdout)
	log.Println("Starting log...")

	fiatRatesChannel := make(chan fiatCurrencyExchange.ExchangeRates)

	// get exchange rates to start
	go fiatCurrencyExchange.FiatCurrencyExchangeRates(fiatRatesChannel, 60*time.Minute, fiatCurrencyExchange.RealExchangeClient{})
	fiatMsg := <-fiatRatesChannel
	if fiatMsg.Err != nil {
		log.Printf("Fiat exchange error: %v \n", fiatMsg.Err)
	}

	for {
		select {
		case msg := <-fiatRatesChannel:
			if msg.Err != nil {
				log.Printf("Fiat exchange error: %v \n", msg.Err)
			}
		default:
			log.Println("No fiat exchange rate message received...")
		}

		// fiatMsg = <-fiatRatesChannel

		cryptoExchangeChannel := make(chan CryptoExchanges.CryptoData)

		for _, elem := range exchangeDataList {
			exchangeMutex(elem, cryptoExchangeChannel)
		}

		var cryptoExchangeData []CryptoExchanges.CryptoData

		for _, app := range exchangeDataList {
			for range app.list {
				val := <-cryptoExchangeChannel
				if val.Error != nil {
					log.Printf("Name: %s, Error: %v, Coin: %v\n", val.Name, val.Error, val.Coin)
				} else {
					tmpVal := ConvertCurrency(val, fiatMsg) // TODO this is wrong
					cryptoExchangeData = append(cryptoExchangeData, tmpVal)
				}
			}
		}

		// sort cryptos by crypto-currency
		var set []string
		for _, item := range cryptoExchangeData {
			set = append(set, item.Crypto)
		}
		uniqueCryptos := DeduplicateStrings(set)
		mapCrypto := map[string][]CryptoExchanges.CryptoData{}
		for i := range uniqueCryptos {
			var listCrypto []CryptoExchanges.CryptoData
			for _, item := range cryptoExchangeData {
				if item.Crypto == uniqueCryptos[i] {
					listCrypto = append(listCrypto, item)
				}
			}
			mapCrypto[uniqueCryptos[i]] = listCrypto
		}

		type arbStruct struct {
			name, crypto string
			arb          float64
		}

		var listArb []arbStruct
		for _, cryptoList := range mapCrypto {
			for _, itemOuter := range cryptoList {
				for _, itemInner := range cryptoList {
					arb := CheckArbitrage(itemInner, itemOuter)
					listArb = append(listArb, arbStruct{fmt.Sprintf("bid: %s, ask: %s", itemInner.Name, itemOuter.Name), itemOuter.Crypto, arb})
				}
			}
		}

		for _, item := range listArb {
			if item.arb > minimumArbitrageRatio {
				val := strconv.FormatFloat(item.arb, 'f', -1, 64)
				log.Printf("ARBITRAGE on %s at %v\n", item.name, val)
			}
		}

		// biggest limit seen is 1 call per second
		time.Sleep(time.Second * 100)
	}
}
