package fiatCurrencyExchange

import (
	"testing"
	"time"
)

func TestFiatCurrencyExchangeRatesDataPayloadInt(t *testing.T) {
	client := &RealExchangeClient{}

	ch := make(chan ExchangeRates)
	go FiatCurrencyExchangeRates(ch, 10*time.Second, client)

	select {
	case result := <-ch:
		if result.Err != nil {
			t.Errorf("Did not expect an error but got one, server could be down: %v", result.Err)
		}

		// Check if the data payload is coming through
		if result.Rates["USD2AUD"] <= 0.1 && result.Rates["USD2AUD"] >= 3 {
			t.Errorf("Expected USD2AUD rate outside of range could be an error got %v", result.Rates["USD2AUD"])
		}
		if result.Rates["GBP2AUD"] <= 0.1 && result.Rates["GBP2AUD"] >= 3 {
			t.Errorf("Expected GBP2AUD rate outside of range could be an error got %v", result.Rates["GBP2AUD"])
		}
		t.Logf("Got USD2AUD rate %v and GBP2AUD rate %v", result.Rates["USD2AUD"], result.Rates["GBP2AUD"])
	case <-time.After(20 * time.Second):
		t.Errorf("Test timed out")
	}
}
