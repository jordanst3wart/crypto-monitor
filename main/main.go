package main


// go run pokemon.go

import (
	"crypto-monitor/structs"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)


// TODO change interface to use float64 for numbers rather than strings
// TODO simplify
// TODO logrus ???


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
