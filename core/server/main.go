package main

import (
	"g2/interfaces"
	"g2/interfaces/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	allHandler := handler.NewAllHandler()
	r := interfaces.NewRouter(allHandler)
	r.Set(e)

	e.Logger.Debug(e.Start(":8090"))
}
