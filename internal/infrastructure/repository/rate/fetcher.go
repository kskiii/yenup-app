package rate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	domain "yenup/internal/domain/rate"
)

// APIResponse is the response structure for single date endpoint
type APIResponse struct {
	Success bool               `json:"success"`
	Base    string             `json:"base"`
	Date    string             `json:"date"`
	Rates   map[string]float64 `json:"rates"`
}

type Fetcher struct {
	APIKey string
	URL    string
}

func NewFetcher(apiKey, url string) *Fetcher {
	return &Fetcher{
		APIKey: apiKey,
		URL:    url,
	}
}

// FetchRate fetches the exchange rate for base/target by using EUR as intermediate
// Since free plan only supports EUR as base, we calculate:
// base/target = EUR/target ÷ EUR/base
// Example: CAD/JPY = EUR/JPY ÷ EUR/CAD = 160 ÷ 1.5 = 106.67
func (f *Fetcher) FetchRate(date time.Time, base string, target string) (domain.Rate, error) {
	dateStr := date.Format("2006-01-02")

	// Request both base and target currencies with EUR as base
	// Example: symbols=CAD,JPY
	fullUrl := fmt.Sprintf(
		"%s%s?base=EUR&symbols=%s,%s&access_key=%s",
		f.URL,
		dateStr,
		base,
		target,
		f.APIKey,
	)

	resp, err := http.Get(fullUrl)
	if err != nil {
		return domain.Rate{}, fmt.Errorf("failed to fetch rate: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Rate{}, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)

	var data APIResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return domain.Rate{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Get EUR/base and EUR/target rates
	eurToBase := data.Rates[base]
	eurToTarget := data.Rates[target]

	// Validate rates exist
	if eurToBase == 0 {
		return domain.Rate{}, fmt.Errorf("rate for %s not found in response", base)
	}
	if eurToTarget == 0 {
		return domain.Rate{}, fmt.Errorf("rate for %s not found in response", target)
	}

	// Calculate base/target rate
	// base/target = EUR/target ÷ EUR/base
	rateValue := eurToTarget / eurToBase

	// If the rate is zero, return an error
	if eurToBase == 0 {
		return domain.Rate{}, fmt.Errorf("rate for %s is zero, cannot calculate cross rate", base)
	}

	return domain.Rate{
		Base:   base,
		Target: target,
		Value:  rateValue,
		Date:   date,
	}, nil
}
