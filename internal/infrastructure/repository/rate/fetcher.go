package rate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	domain "yenup/internal/domain/rate"
)

// The response structure for time series data
type APIResponse struct {
	Success    bool                          `json:"success"`
	Timeseries bool                          `json:"timeseries"`
	Rates      map[string]map[string]float64 `json:"rates"`
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

// FetchRates fetches yesterday's and today's JPY rate from CAD
func (f *Fetcher) FetchRate(date time.Time, base string, target string) (domain.Rate, error) {
	dateStr := date.Format("2006-01-02")

	fullUrl := fmt.Sprintf(
		"%s%s?base=%s&symbols=%s&access_key=%s",
		f.URL,
		dateStr,
		base,
		target,
		f.APIKey,
	)

	resp, err := http.Get(fullUrl)
	if err != nil {
		// HTTP request failed
		return domain.Rate{}, fmt.Errorf("failed to fetch rate: %w", err)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return domain.Rate{}, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	var data APIResponse
	if err := json.Unmarshal(body, &data); err != nil {
		// Failed to parse JSON
		return domain.Rate{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	rateValue := data.Rates[date.Format("2006-01-02")][target]

	Rate := domain.Rate{
		Base:   base,
		Target: target,
		Value:  rateValue,
		Date:   date,
	}

	return Rate, nil
}
