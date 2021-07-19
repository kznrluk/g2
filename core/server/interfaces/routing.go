package interfaces

import (
	"g2/interfaces/handler"
	"github.com/labstack/echo/v4"
)

type Router interface {
	Set(e *echo.Echo)
}

type router struct {
	handlers handler.AllHandler
}

func NewRouter(handlers handler.AllHandler) Router {
	return &router{handlers: handlers}
}

func (r router) Set(e *echo.Echo) {
	e.GET("/check", r.handlers.HealthCheckHandler)
	e.GET("/ws", r.handlers.WebSocketUpgradeHandler)
}
