package server

import (
	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()
	InitRouting(e)
	e.Logger.Fatal(e.Start(":8080"))
}
