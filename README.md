# 💴 YenUp

YenUp is a currency exchange rate monitoring tool built with Go. It checks the JPY exchange rate against a base currency (e.g., CAD) and sends Slack notifications when the Japanese Yen becomes stronger compared to the previous day.

This project demonstrates the implementation of **Clean Architecture** and **Dependency Injection** in Go.

## 🚀 Features

- **Rate Monitoring**: Fetches daily exchange rates using external APIs.
- **Trend Alert**: Sends a Slack notification automatically when JPY strengthens (i.e., the base currency/JPY rate drops).
- **Smart Calculation**: Implements cross-rate calculation (via EUR) to support free-tier limitations of exchange rate APIs.
- **REST API**: Provides a RESTful endpoint to trigger checks manually and retrieve detailed rate data.

## 🛠 Tech Stack

- **Language**: Go 1.23+
- **Framework**: Gin (HTTP Web Framework)
- **Architecture**: Clean Architecture (Handlers, Usecases, Domains, Repositories)
- **Dependency Injection**: Registry pattern
- **External API**: exchangeratesapi.io / Frankfurter
- **Notification**: Slack Incoming Webhook

## 🏗 Architecture

The project follows the **Clean Architecture** principles to ensure separation of concerns and testability.

```text
cmd/
  └── yenup/        # Entry point (main.go)
internal/
  ├── config/       # Configuration management
  ├── domain/       # Domain models and repository interfaces
  ├── usecase/      # Business logic (rate comparison)
  ├── handler/      # HTTP handlers (Gin)
  ├── infrastructure/
  │   └── repository/ # External API & Slack implementation
  └── registory/    # Dependency Injection container
```

## 🏁 Getting Started

### Prerequisites

- Go 1.23 or higher
- Slack Webhook URL (for notifications)
- Exchange Rates API Key (optional if using free tier logic)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/yenup.git
   cd yenup
   ```

2. Create a `.env` file from the example:
   ```bash
   cp .env.example .env
   ```

3. Configure environment variables in `.env`:
   ```env
   APP_PORT=8080
   BASE_CURRENCY=CAD
   TARGET_CURRENCY=JPY
   EXCHANGE_RATE_API_URL=http://api.exchangeratesapi.io/v1/
   EXCHANGE_RATE_API_KEY=your_api_key_here
   SLACK_WEBHOOK_URL=https://hooks.slack.com/services/xxx/yyy/zzz
   ```

### Running the Application

```bash
go run cmd/yenup/main.go
```

### Usage

Trigger a rate check via HTTP request:

```bash
curl "http://localhost:8080/check-rate?base=CAD&target=JPY"
```

Response example:
```json
{
  "status": "success",
  "message": "Rate check executed successfully",
  "data": {
    "base": "CAD",
    "target": "JPY",
    "today_rate": 112.7396,
    "yesterday_rate": 113.2207,
    "change": "down (JPY stronger)",
    "is_notified": true
  }
}
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
