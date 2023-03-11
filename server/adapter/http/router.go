package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(e *echo.Echo, userHandler *UserHandler) {
	e.Use(middleware.Logger(), middleware.Recover())
	e.GET("/.well-known/webfinger", userHandler.GetWebFinger())
	e.POST("/users", userHandler.AddUser())
	e.GET("/users/:resource ", userHandler.GetUser())
	e.POST("/users/:resource/inbox", userHandler.PostInbox())
}
