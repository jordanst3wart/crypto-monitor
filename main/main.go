package main


// go run pokemon.go

import (
	"crypto-monitor/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)


// TODO change interface to use float64 for numbers rather than strings
// TODO simplify
// TODO logrus ???


/*
	Gets exchange rate, every 5 minutes
 */
func currencyExchangeRates(ch chan ExchangeRates) {
	exchangeMap := make(map[string]float64)
	var responseObject structs.CurrencyExchangeAPI
	// TODO need to cache
	for {
		resp, err := http.Get("https://api.exchangeratesapi.io/latest?base=USD")
		if err != nil {
			// should wrap error
			ch <- ExchangeRates{nil, err}
			return
		}
		// https://fixer.io/product

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			ch <- ExchangeRates{nil, err}
			return
		}
		err = json.Unmarshal(responseData, &responseObject)
		if err != nil {
			ch <- ExchangeRates{nil, err}
			return
		} else {
			exchangeMap["USD2AUD"] = responseObject.Rates.AUD / responseObject.Rates.USD
			exchangeMap["GBP2AUD"] = responseObject.Rates.AUD / responseObject.Rates.GDP
			ch <- ExchangeRates{exchangeMap, err}
		}
		time.Sleep(5 * time.Minute)
	}
}

type Four struct {
	name, url, currency, crypto string
}

func requestToExchange(exchange structs.CryptoExchange, urlList []Four, ch chan structs.CryptoDTO){
	for _, v := range urlList {
		go exchange.RequestUpdate(v.name, v.url, ch, v.currency, v.crypto)
		//ch<-CryptoDTO{v.name,val, err}
		//groupList = append(groupList, CryptoDTO{v.name,val, err})
	}
	//return groupList
}


type ExchangeRates struct {
	rates map[string]float64
	err error
}

type startData struct {
	exchange string
	list []Four
}

func convertHelper(conversion float64, dto structs.CryptoDTO) (structs.CryptoDTO){

	last, _   :=dto.Coin.LastFloat()
	high, _   := dto.Coin.HighFloat()
	low, _    := dto.Coin.LowFloat()
	volume, _ := dto.Coin.VolumeFloat()
	ask, _    := dto.Coin.AskFloat()
	bid, _    := dto.Coin.BidFloat()

	tmpCoin := structs.BTCMarket{
		last * conversion,
		high * conversion,
		low * conversion,
		volume * conversion,
		bid * conversion,
		ask * conversion}

	return structs.CryptoDTO{
		dto.Name,
		tmpCoin,
		dto.Error,
		"AUD",
		dto.Crypto,
	}
}

func CheckArbitage(exchange1 structs.CryptoDTO, exchange2 structs.CryptoDTO) float64 {
	bid, _ := exchange1.Coin.BidFloat()
	ask, _ := exchange2.Coin.AskFloat()
	return bid / ask
}

func ConvertCurrency(crypto structs.CryptoDTO, exchangeRate ExchangeRates) structs.CryptoDTO {
	switch crypto.Currency {
	case "USD":
		return convertHelper(exchangeRate.rates["USD2AUD"],crypto)
	case "GBP":
		return convertHelper(exchangeRate.rates["GBP2AUD"],crypto)
	case "AUD":
		return crypto
	default:
		log.Println("Unknown currency trying to be converted")
		return structs.CryptoDTO{}
	}
}

func UniqueStrings(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}


func calculate(data startData, ch chan structs.CryptoDTO) {
	switch data.exchange {
	case "CoinfloorTickerAndBitstamp":
		//log.Println("Requesting data from CoinfloorTickerAndBitstamp")
		var resseObjectCoinfloorAndBitstamp structs.CoinfloorTickerAndBitstamp
		requestToExchange(resseObjectCoinfloorAndBitstamp, data.list, ch)
	case "IndepentReserve":
		//log.Println("Requesting data from IndepentReserve")
		var responseObjectIndependentReserve structs.IndepentReserve
		requestToExchange(responseObjectIndependentReserve, data.list, ch)
	case "GeminiTickerBTC":
		//log.Println("Requesting data from IndepentReserve")
		var responseObjectGeminiBTC structs.GeminiTickerBTC
		requestToExchange(responseObjectGeminiBTC, data.list, ch)
	case "GeminiTickerETH":
		//log.Println("Requesting data from GeminiTickerETH")
		var responseObjectGeminiETH structs.GeminiTickerETH
		requestToExchange(responseObjectGeminiETH, data.list, ch)
	case "BTCMarket":
		//log.Println("Requesting data from BTCMarket")
		var responseObjectBTC structs.BTCMarket
		requestToExchange(responseObjectBTC, data.list, ch)
	case "ACXTicker":
		//log.Println("Requesting data from ACX")
		var responseObjectACX structs.ACXTicker
		requestToExchange(responseObjectACX, data.list, ch)
	case "Coinjar":
		//log.Println("Requesting data from Coinjar")
		var responseObjectCoinjar structs.Coinjar
		requestToExchange(responseObjectCoinjar, data.list, ch)
	default:
		log.Println("Invalid key in startData")
	}
}

func main() {
	DEBUG := true
	// log setup
	if DEBUG {
		log.SetOutput(os.Stdout)
	} else {
		f, err := os.OpenFile("server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("error opening file: %v", err)
		}
		defer f.Close()
		log.SetOutput(f) // if not local
	}

	log.Println("Starting log")
	// log setup finished
	var ARB_RATIO float64
	ARB_RATIO = 1.02
	chRates := make(chan ExchangeRates)

	// get exchange rates to start
	go currencyExchangeRates(chRates)
	msg := <-chRates
	if msg.err != nil {
		log.Println("Error: ", msg.err)
	} else if DEBUG {
		log.Println(msg.rates)
	}

	for {
		log.Println("msg", msg.rates)
		start := time.Now()
		log.Println("Starting iteration...")

		//DEBUG := true
		//chRates := make(chan ExchangeRates)
		select {
		case msg := <-chRates:
			fmt.Println("received message", msg)
			if msg.err != nil {
				log.Println("Error: ", msg.err)
			} else if DEBUG {
				log.Println(msg.rates)
			}
		default:
			fmt.Println("no message received")
		}

		ch := make(chan structs.CryptoDTO)

		startData := getStartData()
		for _, elem := range startData {
			calculate(elem, ch)
		}
		if DEBUG {
			log.Println("%.2fs elapsed\n", time.Since(start).Seconds())
		}

		listThing := []structs.CryptoDTO{}

		for _, app := range startData {
			for range app.list {
				val := <-ch
				if val.Error != nil {
					log.Println("Name:", val.Name, "Error", val.Error)
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
				log.Println("ARBITAGE!!! on " + item.name + " at " + val)
			}
		}

		//log.Println(listArb)

		/*if DEBUG {
			log.Println("got values from Crypto exchange")
			log.Println(groupList)
		}*/
		if DEBUG {
			log.Println("%.2fs elapsed\n", time.Since(start).Seconds())
		}
		// 14.40s and 8s elapsed before async
		// 1.6s after async

		// send email if arbitage found
		// wait five minutes for next iteration
		// biggest limit seen is 1 call per second
		time.Sleep(time.Second*100)
	}
}
