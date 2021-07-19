package main

import (
	"g2/domain"
	"g2/interfaces"
	"g2/interfaces/handler"
	"g2/room"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net"
)

func main() {
	e := echo.New()

	handler.ConnectionPool = make(map[domain.RoomId][]net.Conn)
	room.InitRooms()

	e.Use(middleware.Logger())

	allHandler := handler.NewAllHandler()
	r := interfaces.NewRouter(allHandler)
	r.Set(e)

	e.Logger.Debug(e.Start(":8090"))
}
