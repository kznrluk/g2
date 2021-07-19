package handler

import (
	"github.com/labstack/echo/v4"
)

type AllHandler struct {
	HealthCheckHandler echo.HandlerFunc
}

func NewAllHandler() AllHandler {
	return AllHandler{
		HealthCheckHandler: ProvideHealthCheckHandler(),
	}
}

func ProvideHealthCheckHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(200, "OK")
	}
}
