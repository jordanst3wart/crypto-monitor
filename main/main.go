package main

import (
	"crypto-monitor/main/cryptoExchanges"
	"crypto-monitor/main/fiatCurrencyExchange"
	"log"
	"os"
	"strconv"
	"time"
)

// TODO Use orders to determine the volume of available & log amount of volume is available to buy/sell
// TODO logrus to have better debug logic
// TODO write tests
// TODO use config file

func main() {
	// TODO setup reading from config file
	ARB_RATIO := 1.01

	// log setup
	log.SetOutput(os.Stdout)
	log.Println("Starting log...")

	fiatRates := make(chan fiatCurrencyExchange.ExchangeRates)

	// get exchange rates to start
	go fiatCurrencyExchange.FiatCurrencyExchangeRates(fiatRates, 60*time.Minute, fiatCurrencyExchange.RealExchangeClient{})
	msg := <-fiatRates
	if msg.Err != nil {
		log.Println("Fiat exchange error: ", msg.Err)
	}

	for {
		select {
		case msg := <-fiatRates:
			if msg.Err != nil {
				log.Println("Fiat exchange error: ", msg.Err)
			}
		default:
			log.Println("No fiat message received.")
		}

		ch := make(chan CryptoExchanges.CryptoDTO)

		startData := getStartData()
		for _, elem := range startData {
			calculate(elem, ch)
		}

		listThing := []CryptoExchanges.CryptoDTO{}

		for _, app := range startData {
			for range app.list {
				val := <-ch
				if val.Error != nil {
					log.Println("Name:", val.Name, "Error", val.Error, "Coin", val.Coin)
				} else {
					//log.Println(val)
					tmpVal := ConvertCurrency(val, msg)
					listThing = append(listThing, tmpVal)
					// if greater than some margin send email
					// standardise logging
				}
			}
		}

		// sort cryptos by crypto-currency
		set := []string{}
		for _, item := range listThing {
			set = append(set, item.Crypto)
		}
		uniqueCryptos := UniqueStrings(set)
		mapCrypto := map[string][]CryptoExchanges.CryptoDTO{}
		for i := range uniqueCryptos {
			listCrypto := []CryptoExchanges.CryptoDTO{}
			for _, item := range listThing {
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

		listArb := []arbStruct{}
		for _, cryptoList := range mapCrypto {
			for _, itemOuter := range cryptoList {
				for _, itemInner := range cryptoList {
					arb := CheckArbitrage(itemInner, itemOuter)
					listArb = append(listArb, arbStruct{"bid: " + itemInner.Name + ", ask:" + itemOuter.Name, itemOuter.Crypto, arb})
				}
			}
		}

		for _, item := range listArb {
			if item.arb > ARB_RATIO {
				//val, _ := strconv.ParseFloat(item.arb, 64)
				val := strconv.FormatFloat(item.arb, 'f', -1, 64)
				log.Println("ARBITRAGE!!! on " + item.name + " at " + val)
			}
		}

		// biggest limit seen is 1 call per second
		time.Sleep(time.Second * 100)
	}
}
