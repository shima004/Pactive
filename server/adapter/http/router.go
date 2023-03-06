package http

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

func InitRouter() {
	// ルーティングの設定
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())

}
