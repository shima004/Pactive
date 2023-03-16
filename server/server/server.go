package server

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	logger "github.com/labstack/gommon/log"
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
	// static image
	e.Static("/images", "assets/images")

	serverInfo := config.GetServerInfo()
	mode := os.Getenv("SERVER_MODE")

	switch mode {
	case "production":
		log.Println("Server is running on production mode")
		serverInfo.Protocol = "https"
		e.Logger.SetLevel(logger.INFO)
		e.Logger.Fatal(e.StartTLS(":443", serverInfo.CertFile, serverInfo.KeyFile))
	case "development":
		log.Println("Server is running on development mode")
		serverInfo.Protocol = "https"
		e.Logger.SetLevel(logger.DEBUG)
		e.Logger.Fatal(e.StartTLS(":443", serverInfo.CertFile, serverInfo.KeyFile))
	case "local":
		log.Println("Server is running on test mode")
		serverInfo.Protocol = "http"
		serverInfo.Host = "localhost"
		e.Logger.SetLevel(logger.DEBUG)
		e.Logger.Fatal(e.Start(":8080"))
	}
}
