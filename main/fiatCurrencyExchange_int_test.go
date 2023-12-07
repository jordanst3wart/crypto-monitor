package main

import (
	"testing"
	"time"
)

func TestFiatCurrencyExchangeRatesDataPayloadInt(t *testing.T) {
	client := &RealExchangeClient{}

	ch := make(chan ExchangeRates)
	go fiatCurrencyExchangeRates(ch, 10*time.Second, client)

	select {
	case result := <-ch:
		if result.err != nil {
			t.Errorf("Did not expect an error but got one, server could be down: %v", result.err)
		}

		// Check if the data payload is coming through
		if result.rates["USD2AUD"] <= 0.1 && result.rates["USD2AUD"] >= 3 {
			t.Errorf("Expected USD2AUD rate outside of range could be an error got %v", result.rates["USD2AUD"])
		}
		if result.rates["GBP2AUD"] <= 0.1 && result.rates["GBP2AUD"] >= 3 {
			t.Errorf("Expected GBP2AUD rate outside of range could be an error got %v", result.rates["GBP2AUD"])
		}
		t.Logf("Got USD2AUD rate %v and GBP2AUD rate %v", result.rates["USD2AUD"], result.rates["GBP2AUD"])
	case <-time.After(20 * time.Second):
		t.Errorf("Test timed out")
	}
}
