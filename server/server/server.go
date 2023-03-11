package server

import (
	"github.com/labstack/echo/v4"
	"github.com/shima004/pactive/adapter/http"
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

	e.Logger.Fatal(e.Start(":8080"))

}
