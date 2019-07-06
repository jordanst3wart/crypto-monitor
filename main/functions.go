package main



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
