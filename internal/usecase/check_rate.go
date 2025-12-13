package usecase

import (
	"fmt"
	"time"
	"yenup/internal/domain/rate"
	domain "yenup/internal/domain/rate"
)

// rateCheckUsecase is the interface for the rate check usecase
type rateCheckUsecase interface {
	CheckRates(date time.Time, base string, target string) (domain.Rate, error)
}

// RateChecker is the usecase for checking the rate
type RateChecker struct {
	Repo     rate.Repository
	Notifier rate.Notifier
}

func NewRateChecker(repo rate.Repository, notifier rate.Notifier) *RateChecker {
	return &RateChecker{
		Repo:     repo,
		Notifier: notifier,
	}
}

func (r *RateChecker) CheckRates(base, target string) error {
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	// Get rates from repository
	todayRate, err := r.Repo.FetchRate(today, base, target)
	if err != nil {
		return fmt.Errorf("failed to fetch today's rate: %w", err)
	}
	yesterdayRate, err := r.Repo.FetchRate(yesterday, base, target)
	if err != nil {
		return fmt.Errorf("failed to fetch yesterday's rate: %w", err)
	}

	if todayRate.Value < yesterdayRate.Value {
		msg := fmt.Sprintf(
			"CAD/JPY 上昇！昨日: %.4f -> 今日： %.4f",
			yesterdayRate.Value,
			todayRate.Value,
		)
		return r.Notifier.Notify(msg)
	}
	return nil
}
