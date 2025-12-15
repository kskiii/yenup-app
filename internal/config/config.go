package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort            string
	BaseCurrency       string
	TargetCurrency     string
	ExchangeRateAPIKey string
	ExchangeRateAPIURL string
	SlackWebhookURL    string
}

func Load() (*Config, error) {
	// load the config from the environment variables
	_ = godotenv.Load()

	cfg := &Config{
		AppPort:            getEnv("APP_PORT", "8080"),
		BaseCurrency:       getEnv("BASE_CURRENCY", "CAD"),
		TargetCurrency:     getEnv("TARGET_CURRENCY", "JPY"),
		ExchangeRateAPIKey: getEnv("EXCHANGE_RATE_API_KEY", ""),
		ExchangeRateAPIURL: getEnv("EXCHANGE_RATE_API_URL", ""),
		SlackWebhookURL:    getEnv("SLACK_WEBHOOK_URL", ""),
	}
	return cfg, nil
}

func getEnv(key string, fallback string) string {
	// return the value of the environment variable if it exists, otherwise return an empty string
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
