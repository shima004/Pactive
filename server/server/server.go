package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/shima004/pactive/adapter/http"
	"github.com/shima004/pactive/config"
	"github.com/shima004/pactive/domain/service"
	"github.com/shima004/pactive/infra/postgres"
	"github.com/shima004/pactive/usecase"
)

func Run() {
	e := echo.New()

	postgresql := postgres.InitDB()
	db, _ := postgresql.DB()
	defer db.Close()

	userRepository := postgres.NewUserRepository(postgresql)
	userService := service.NewUserService(userRepository)
	userUsecase := usecase.NewUserUsecase(userService)
	userHandler := http.NewUserHandler(userUsecase)

	http.InitRouter(e, userHandler)

	serverInfo := config.GetServerInfo()
	if serverInfo.Protocol == "https" {
		log.Println("Server is running on https mode")
		e.Logger.Fatal(e.StartTLS(":443", serverInfo.CertFile, serverInfo.KeyFile))
	} else {
		log.Println("Server is running on http mode")
		e.Logger.Fatal(e.Start(":80"))
	}
}
