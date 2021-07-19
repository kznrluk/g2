package handler

import (
	"github.com/gobwas/ws"
	"github.com/labstack/echo/v4"
)

type AllHandler struct {
	HealthCheckHandler      echo.HandlerFunc
	WebSocketUpgradeHandler echo.HandlerFunc
}

func NewAllHandler() AllHandler {
	return AllHandler{
		HealthCheckHandler:      ProvideHealthCheckHandler(),
		WebSocketUpgradeHandler: ProvideWebSocketUpgradeHandler(),
	}
}

func ProvideHealthCheckHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(200, "OK")
	}
}

func ProvideWebSocketUpgradeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn, _, _, err := ws.UpgradeHTTP(c.Request(), c.Response())
		if err != nil {
			return c.String(500, err.Error())
		}

		go HandleWebSocket(conn)

		return nil
	}
}
