package rate

import (
	"net/http"
	"yenup/internal/usecase"

	"github.com/gin-gonic/gin"
)

// RateHandler is the handler for the rate route
type RateHandler struct {
	Usecase *usecase.RateChecker
}

// NewRateHandler creates a new RateHandler
func NewRateHandler(u *usecase.RateChecker) *RateHandler {
	return &RateHandler{
		Usecase: u,
	}
}

// Response is the response for the rate route
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// CheckRate checks the rate of the base and target currencies
func (h *RateHandler) CheckRate(c *gin.Context) {
	base := c.Query("base")
	target := c.Query("target")

	if base == "" || target == "" {
		c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "base and target are required",
			Data:    nil,
		})
		return
	}

	// use usecase to check the rate
	err := h.Usecase.CheckRates(base, target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "Rate check executed successfully",
		Data:    nil,
	})
}
