package handler

import (
	"yenup/internal/handler/rate"
)

type Handler struct {
	RateHandler *rate.RateHandler
}

func NewHandler(rateHandler *rate.RateHandler) *Handler {
	return &Handler{
		RateHandler: rateHandler,
	}
}
