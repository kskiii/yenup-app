package registry

import (
	"yenup/internal/config"
	"yenup/internal/handler"
	rateHandler "yenup/internal/handler/rate"
	notifierRepo "yenup/internal/infrastructure/repository/notifier"
	rateRepo "yenup/internal/infrastructure/repository/rate"
	usecase "yenup/internal/usecase"
)

type Registry struct {
	config     *config.Config
	AppHandler *handler.Handler
}

func NewRegistry(cfg *config.Config) (*Registry, error) {
	rateFetcher := rateRepo.NewFetcher(cfg.ExchangeRateAPIKey, cfg.ExchangeRateAPIURL)
	slackNotifier := notifierRepo.NewSlackNotifier(cfg.SlackWebhookURL)

	// usecase
	rateUsecase := usecase.NewRateChecker(rateFetcher, slackNotifier)

	// handler
	rateHandler := rateHandler.NewRateHandler(rateUsecase)

	// app handler
	appHandler := handler.NewHandler(rateHandler)

	return &Registry{
		config:     cfg,
		AppHandler: appHandler,
	}, nil
}
