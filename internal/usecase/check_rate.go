package usecase

import (
	"fmt"
	"time"
	"yenup/internal/domain/rate"
)

// RateCheckUsecase is the interface for the rate check usecase
type RateCheckUsecase interface {
	CheckRates(base, target string) (*CheckRateResult, error)
}

// CheckRateResult is the result of checking the rate
type CheckRateResult struct {
	TodayRate     float64
	YesterdayRate float64
	IsNotified    bool
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

func (r *RateChecker) CheckRates(base, target string) (*CheckRateResult, error) {
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	// Get rates from repository
	todayRate, err := r.Repo.FetchRate(today, base, target)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch today's rate: %w", err)
	}
	yesterdayRate, err := r.Repo.FetchRate(yesterday, base, target)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch yesterday's rate: %w", err)
	}

	if todayRate.Value < yesterdayRate.Value {
		msg := fmt.Sprintf(
			"JPY Stronger Alert! %s/%s: Yesterday %.4f -> Today %.4f",
			base,
			target,
			yesterdayRate.Value,
			todayRate.Value,
		)
		err := r.Notifier.Notify(msg)
		if err != nil {
			return nil, fmt.Errorf("failed to notify: %w", err)
		}
		return &CheckRateResult{
			TodayRate:     todayRate.Value,
			YesterdayRate: yesterdayRate.Value,
			IsNotified:    true,
		}, nil
	}
	return &CheckRateResult{
		TodayRate:     todayRate.Value,
		YesterdayRate: yesterdayRate.Value,
		IsNotified:    false,
	}, nil
}
