package fiatCurrencyExchange

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

// MockExchangeClient is a mock implementation of the ExchangeClient interface
type MockExchangeClient struct {
	Response *http.Response
	Err      error
}

func (m *MockExchangeClient) getRates() (*http.Response, error) {
	return m.Response, m.Err
}

func TestFiatCurrencyExchangeRates(t *testing.T) {
	// Define different test scenarios
	tests := []struct {
		name          string
		client        ExchangeClient
		expectedError bool
	}{
		{
			name: "Successful response",
			client: &MockExchangeClient{
				Response: &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
						"Rates": {
							"AUD": 1.5,
							"USD": 1.2,
							"GDP": 0.8
						}
					}`))),
				},
				Err: nil,
			},
			expectedError: false,
		},
		{
			name: "Exchange client returns error",
			client: &MockExchangeClient{
				Response: nil,
				Err:      errors.New("failed to get rates"),
			},
			expectedError: true,
		},
		// ... other test scenarios ...
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ch := make(chan ExchangeRates)
			go FiatCurrencyExchangeRates(ch, time.Second, tc.client)

			select {
			case result := <-ch:
				if tc.expectedError && result.Err == nil {
					t.Errorf("Expected an error but didn't get one")
				}
				if !tc.expectedError && result.Err != nil {
					t.Errorf("Did not expect an error but got one: %v", result.Err)
				}
				// ... additional assertions ...
			case <-time.After(2 * time.Second):
				t.Errorf("Test timed out")
			}
		})
	}
}

func TestFiatCurrencyExchangeRatesDataPayload(t *testing.T) {
	mockResponse := `{
		"Rates": {
			"AUD": 1.5,
			"USD": 1.2,
			"GBP": 0.8
		}
	}`
	client := &MockExchangeClient{
		Response: &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mockResponse))),
		},
		Err: nil,
	}

	ch := make(chan ExchangeRates)
	go FiatCurrencyExchangeRates(ch, time.Second, client)

	select {
	case result := <-ch:
		if result.Err != nil {
			t.Errorf("Did not expect an error but got one: %v", result.Err)
		}

		// Check if the data payload is correct
		expectedUSD2AUD := 1.5 / 1.2
		expectedGBP2AUD := 1.5 / 0.8
		if result.Rates["USD2AUD"] != expectedUSD2AUD {
			t.Errorf("Expected USD2AUD rate to be %v, got %v", expectedUSD2AUD, result.Rates["USD2AUD"])
		}
		if result.Rates["GBP2AUD"] != expectedGBP2AUD {
			t.Errorf("Expected GBP2AUD rate to be %v, got %v", expectedGBP2AUD, result.Rates["GBP2AUD"])
		}
	case <-time.After(2 * time.Second):
		t.Errorf("Test timed out")
	}
}
