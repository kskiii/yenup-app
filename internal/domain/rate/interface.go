package rate

import "time"

type Repository interface {
	// Get the rate for a given base and target currency
	FetchRate(date time.Time, base string, target string) (Rate, error)
}

type Notifier interface {
	// Notify the user via　lack
	Notify(message string) error
}
