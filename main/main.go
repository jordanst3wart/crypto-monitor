package main

import (
	"crypto-monitor/structs"
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
	DEBUG := false
	// log setup
	log.SetOutput(os.Stdout)

	log.Println("Starting log...")
	// log setup finished
	var ARB_RATIO float64
	ARB_RATIO = 1.02
	fiatRates := make(chan ExchangeRates)

	// get exchange rates to start
	go fiatCurrencyExchangeRates(fiatRates)
	msg := <-fiatRates
	if msg.err != nil {
		log.Println("Fiat exchange error: ", msg.err)
	} else if DEBUG {
		log.Println("Fiat rates: ", msg.rates)
	}

	for {
		start := time.Now()
		if DEBUG {
			log.Println("Starting iteration...")
		}

		//DEBUG := true
		//fiatRates := make(chan ExchangeRates)
		select {
		case msg := <-fiatRates:
			if msg.err != nil {
				log.Println("Fiat exchange error: ", msg.err)
			} else if DEBUG {
				log.Println("Fiat received message", msg)
			}
		default:
			if DEBUG {
				log.Println("No fiat message received.")
			}
		}

		ch := make(chan structs.CryptoDTO)

		startData := getStartData()
		for _, elem := range startData {
			calculate(elem, ch)
		}
		if DEBUG {
			log.Println("elapsed time: ", time.Since(start).Seconds())
		}

		listThing := []structs.CryptoDTO{}

		for _, app := range startData {
			for range app.list {
				val := <-ch
				if val.Error != nil {
					log.Println("Name:", val.Name, "Error", val.Error, "Coin", val.Coin)
				} else {
					//log.Println(val)
					tmpVal := ConvertCurrency(val, msg)
					listThing = append(listThing, tmpVal)
					if DEBUG {
						log.Println(tmpVal)
					}
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
		mapCrypto := map[string][]structs.CryptoDTO{}
		for i := range uniqueCryptos {
			listCrypto := []structs.CryptoDTO{}
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
					arb := CheckArbitage(itemInner, itemOuter)
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

		//log.Println(listArb)

		/*if DEBUG {
			log.Println("got values from Crypto exchange")
			log.Println(groupList)
		}*/
		if DEBUG {
			log.Println("elapsed time: ", time.Since(start).Seconds())
		}
		// 14.40s and 8s elapsed before async
		// 1.6s after async

		// send email if arbitage found
		// wait five minutes for next iteration
		// biggest limit seen is 1 call per second
		time.Sleep(time.Second*100)
	}
}
