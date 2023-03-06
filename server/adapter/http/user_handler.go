package http

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/shima004/pactive/domain/model"
	"github.com/shima004/pactive/usecase"
)

type UserHandler struct {
	usecase usecase.IUserUsecase
}

func NewUserHandler(usecase usecase.IUserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (h *UserHandler) AddUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		name := c.QueryParam("name")
		email := c.QueryParam("email")
		password := c.QueryParam("password")
		user := &model.User{
			Name:     name,
			Email:    email,
			Password: password,
		}
		if err := h.usecase.AddUser(ctx, user); err != nil {
			return c.JSON(400, "invalid user")
		}
		return c.JSON(200, "success")
	}
}

func (h *UserHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.QueryParam("id")
		// string型のidをint型に変換する
		i, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(400, "invalid id")
		}
		user, err := h.usecase.GetUser(ctx, i)
		if err != nil {
			return c.JSON(400, "invalid id")
		}
		return c.JSON(200, user)
	}
}
