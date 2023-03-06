package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shima004/pactive/db"
	"github.com/shima004/pactive/model"
)

func InitRouting(e *echo.Echo) {
	e.GET("/.well-known/webfinger", func(c echo.Context) error {
		if c.QueryParam("resource") != "" {

		}
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/post", func(c echo.Context) error {
		db, err := db.GetDB()
		if err != nil {
			return err
		}
		db.Create(&model.Test{
			Name: "test",
		})
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/get", func(c echo.Context) error {
		db, err := db.GetDB()
		if err != nil {
			return err
		}
		var test model.Test
		db.First(&test)
		return c.String(http.StatusOK, test.Name)
	})
}
