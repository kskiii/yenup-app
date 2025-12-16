package rate

import (
	"encoding/json"
	"fmt"
	"time"
	domain "yenup/internal/domain/rate"
)

// ExchangeRatesResponse is the response structure for ExchangeRates API
type ExchangeRatesResponse struct {
	Success bool               `json:"success"`
	Base    string             `json:"base"`
	Date    string             `json:"date"`
	Rates   map[string]float64 `json:"rates"`
}

// ExchangeRatesFetcher fetches rates from the ExchangeRates API (requires API key)
type ExchangeRatesFetcher struct {
	APIKey string
	URL    string
}

// NewExchangeRatesFetcher creates a new ExchangeRatesFetcher
func NewExchangeRatesFetcher(apiKey, url string) *ExchangeRatesFetcher {
	return &ExchangeRatesFetcher{
		APIKey: apiKey,
		URL:    url,
	}
}

// FetchRate fetches the exchange rate for base/target by using EUR as intermediate
// Since free plan only supports EUR as base, we calculate:
// base/target = EUR/target ÷ EUR/base
func (f *ExchangeRatesFetcher) FetchRate(date time.Time, base string, target string) (domain.Rate, error) {
	dateStr := date.Format("2006-01-02")

	url := fmt.Sprintf(
		"%s%s?base=EUR&symbols=%s,%s&access_key=%s",
		f.URL,
		dateStr,
		base,
		target,
		f.APIKey,
	)

	body, err := doGet(url)
	if err != nil {
		return domain.Rate{}, fmt.Errorf("failed to fetch rate: %w", err)
	}

	var data ExchangeRatesResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return domain.Rate{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Get EUR/base and EUR/target rates
	eurToBase := data.Rates[base]
	eurToTarget := data.Rates[target]

	if eurToBase == 0 {
		return domain.Rate{}, fmt.Errorf("rate for %s not found in response", base)
	}
	if eurToTarget == 0 {
		return domain.Rate{}, fmt.Errorf("rate for %s not found in response", target)
	}

	// Calculate cross rate: base/target = EUR/target ÷ EUR/base
	rateValue := eurToTarget / eurToBase

	return domain.Rate{
		Base:   base,
		Target: target,
		Value:  rateValue,
		Date:   date,
	}, nil
}
