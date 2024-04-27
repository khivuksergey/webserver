package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Health checks the health of the service.
//
// @Tags Health
// @Summary Check service health
// @Description Checks if the service is working properly
// @ID health
// @Produce json
// @Success 200 {string} string "OK"
// @Router /health [get]
func health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
