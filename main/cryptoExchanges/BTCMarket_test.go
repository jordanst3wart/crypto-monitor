package CryptoExchanges

import (
	"testing"
	"time"
)

// TODO use interface
func TestCryptoExchange(t *testing.T) {
	ch := make(chan CryptoDTO)
	var responseObjectBTC BTCMarket
	go responseObjectBTC.RequestUpdate("BTCMarket_AUD_BTC", "https://api.btcmarkets.net/market/BTC/AUD/tick", ch, "AUD", "BTC")
	select {
	case result := <-ch:
		if result.Error != nil {
			t.Errorf("Didn't expect an error but got one")
		}
		// check payload...
		/*if !tc.expectedError && result.Err != nil {
			t.Errorf("Did not expect an error but got one: %v", result.Err)
		}*/
		// ... additional assertions ...
	case <-time.After(10 * time.Second):
		t.Errorf("Test timed out")
	}
}
