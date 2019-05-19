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

func currencyExchangeRates(ch chan ExchangeRates) {
	exchangeMap := make(map[string]float64)

	resp, err := http.Get("https://api.exchangeratesapi.io/latest?base=USD")
	if err != nil {
		// should wrap error
		ch <- ExchangeRates{nil, err}
		return
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- ExchangeRates{nil, err}
		return
	}
	var responseObject structs.CurrencyExchangeAPI
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		ch <- ExchangeRates{nil, err}
		return
	} else {
		exchangeMap["USD2AUD"] = responseObject.Rates.AUD / responseObject.Rates.USD
		exchangeMap["GBP2AUD"] = responseObject.Rates.AUD / responseObject.Rates.GDP
		ch <-ExchangeRates{exchangeMap, err}
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
		fmt.Printf("Unknown currency")
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
		fmt.Println("Requesting data from CoinfloorTickerAndBitstamp")
		var resseObjectCoinfloorAndBitstamp structs.CoinfloorTickerAndBitstamp
		requestToExchange(resseObjectCoinfloorAndBitstamp, data.list, ch)
	case "IndepentReserve":
		fmt.Println("Requesting data from IndepentReserve")
		var responseObjectIndependentReserve structs.IndepentReserve
		requestToExchange(responseObjectIndependentReserve, data.list, ch)
	case "GeminiTickerBTC":
		fmt.Println("Requesting data from IndepentReserve")
		var responseObjectGeminiBTC structs.GeminiTickerBTC
		requestToExchange(responseObjectGeminiBTC, data.list, ch)
	case "GeminiTickerETH":
		fmt.Println("Requesting data from GeminiTickerETH")
		var responseObjectGeminiETH structs.GeminiTickerETH
		requestToExchange(responseObjectGeminiETH, data.list, ch)
	case "BTCMarket":
		fmt.Println("Requesting data from BTCMarket")
		var responseObjectBTC structs.BTCMarket
		requestToExchange(responseObjectBTC, data.list, ch)
	case "ACXTicker":
		fmt.Println("Requesting data from ACX")
		var responseObjectACX structs.ACXTicker
		requestToExchange(responseObjectACX, data.list, ch)
	case "Coinjar":
		fmt.Println("Requesting data from Coinjar")
		var responseObjectCoinjar structs.Coinjar
		requestToExchange(responseObjectCoinjar, data.list, ch)
	default:
		fmt.Printf("Invalid key in startData")
	}
}

func foo2(){

}

func main() {
	start := time.Now()
	// log setup
	f, err := os.OpenFile("server.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f) // if not local
	log.Println("This is a test log entry")
	// log setup finished

	var ARB_RATIO float64
	ARB_RATIO = 1.04
	//DEBUG := true
	chRates := make(chan ExchangeRates)
	go currencyExchangeRates(chRates)
	ch := make(chan structs.CryptoDTO)

	startData := getStartData()
	for _, elem := range startData {
		calculate(elem, ch)
	}

	fmt.Println("%.2fs elapsed\n", time.Since(start).Seconds())
	val1:=<-chRates
	if val1.err != nil {
		log.Println(val1.err)
	} else {
		log.Println(val1)
	}

	listThing := []structs.CryptoDTO{}

	for _, app := range startData {
		for range app.list {
			val:=<-ch
			if val.Error != nil {
				log.Println("Name:", val.Name, "Error", val.Error)
			} else {
				//log.Println(val)
				tmpVal:=ConvertCurrency(val,val1)
				listThing = append(listThing,tmpVal)
				log.Println(tmpVal)
				// TODO check for arbitage
				// for each val go other each other val
				//CheckArbitage()
				// if greater than some margin send email
				// standardise logging
			}
		}
	}

	// sort cryptos by crypto-currency
	set := []string{}
	for _, item := range listThing {
		set = append(set,item.Crypto)
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
		arb float64
	}

	listArb := []arbStruct{}
	for _, cryptoList := range mapCrypto {
		for _, itemOuter := range cryptoList {
			for _, itemInner := range cryptoList {
				arb := CheckArbitage(itemInner, itemOuter)
				listArb = append(listArb, arbStruct{"bid: " + itemInner.Name + ", ask:" + itemOuter.Name,itemOuter.Crypto,arb})
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
	fmt.Println("%.2fs elapsed\n", time.Since(start).Seconds())
	// 14.40s and 8s elapsed before async
	// 1.6s after async


	// send email if arbitage found
}




func getStartData() []startData {
	return []startData{{"CoinfloorTickerAndBitstamp", []Four{
		{"Coinfloor_BTC", "https://webapi.coinfloor.co.uk:8090/bist/XBT/GBP/ticker/", "GBP", "BTC"},
		{"Coinfloor_ETH","https://webapi.coinfloor.co.uk:8090/bist/ETH/GBP/ticker/", "GBP", "ETH"},
		{"Coinfloor_BCH","https://webapi.coinfloor.co.uk:8090/bist/BCH/GBP/ticker/", "GBP", "BCH"},
		{"Bitstamp_BTC","https://www.bitstamp.net/api/v2/ticker/btcusd/", "USD", "BTC"},
		{"Bitstamp_XRP","https://www.bitstamp.net/api/v2/ticker/xrpusd/", "USD", "XRP"},
		{"Bitstamp_LTC","https://www.bitstamp.net/api/v2/ticker/ltcusd/", "USD", "LTC"},
		{"Bitstamp_ETH","https://www.bitstamp.net/api/v2/ticker/ethusd/", "USD", "ETH"},
		{"Bitstamp_BCH","https://www.bitstamp.net/api/v2/ticker/bchusd/", "USD", "BCH"}}},
		startData{"IndepentReserve",[]Four{
			{"IndependentReserve_BTC", "https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xbt&secondaryCurrencyCode=aud", "AUD","BTC"},
			{"IndependentReserve_ETH","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=eth&secondaryCurrencyCode=aud", "AUD", "ETH"},
			{"IndependentReserve_BCH","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=bch&secondaryCurrencyCode=aud", "AUD", "BCH"},
			{"IndependentReserve_XRP","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=xrp&secondaryCurrencyCode=aud", "AUD", "XRP"},
			{"IndependentReserve_LTC","https://api.independentreserve.com/Public/GetMarketSummary?primaryCurrencyCode=ltc&secondaryCurrencyCode=aud", "AUD", "LTC"}}},
		startData{"GeminiTickerBTC", []Four{
			{"GEMINI_BTC", "https://api.gemini.com/v1/pubticker/btcusd", "USD", "BTC"}}},
		startData{"GeminiTickerETH",[]Four{
			{"GEMINI_ETH", "https://api.gemini.com/v1/pubticker/ethusd", "USD", "ETH"}}},
		startData{"BTCMarket", []Four{
			{"BTCMarket_AUD_BTC", "https://api.btcmarkets.net/market/BTC/AUD/tick", "AUD", "BTC"},
			{"BTCMarket_AUD_ETH", "https://api.btcmarkets.net/market/ETH/AUD/tick", "AUD", "ETH"},
			{"BTCMarket_AUD_BCH","https://api.btcmarkets.net/market/BCHABC/AUD/tick", "AUD", "BCH"},
			{"BTCMarket_AUD_XRP", "https://api.btcmarkets.net/market/XRP/AUD/tick", "AUD", "XRP"},
			{"BTCMarket_AUD_LTC","https://api.btcmarkets.net/market/LTC/AUD/tick", "AUD", "LTC"}}},
		startData{"ACXTicker",[]Four{
			{"ACX_AUD_BTC", "https://acx.io:443/api/v2/tickers/btcaud.json", "AUD", "BTC"},
			{"ACX_AUD_ETH", "https://acx.io:443/api/v2/tickers/ethaud.json", "AUD", "ETH"},
			{"ACX_AUD_BCH","https://acx.io:443/api/v2/tickers/bchaud.json", "AUD", "BCH"},
			{"ACX_AUD_LTC", "https://acx.io:443/api/v2/tickers/ltcaud.json", "AUD", "LTC"},
			{"ACX_AUD_XRP","https://acx.io:443/api/v2/tickers/xrpaud.json","AUD", "XRP"}}},
		startData{"Coinjar",[]Four{
			{"Coinjar_AUD_BTC", "https://data.exchange.coinjar.com/products/BTCAUD/ticker", "AUD", "BTC"},
			{"Coinjar_AUD_ETH", "https://data.exchange.coinjar.com/products/ETHAUD/ticker", "AUD", "ETH"},
			{"Coinjar_AUD_XRP","https://data.exchange.coinjar.com/products/XRPAUD/ticker","AUD", "XRP"},
			{"Coinjar_AUD_LTC", "https://data.exchange.coinjar.com/products/LTCAUD/ticker", "AUD", "LTC"}}}}
}