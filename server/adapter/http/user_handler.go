package http

import (
	"encoding/json"
	"io"
	"log"

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
		resource := c.Param("resource")
		user, err := h.usecase.GetUser(ctx, resource)
		if err != nil {
			log.Println(err)
			return c.JSON(400, "invalid id")
		}
		json, err := user.Serialize()
		if err != nil {
			log.Println(err)
			return c.JSON(400, "invalid id")
		}
		return c.JSON(200, json)
	}
}

func (h *UserHandler) GetWebFinger() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		resource := c.QueryParam("resource")
		if resource == "" {
			return c.JSON(200, "<XRD><Link rel=\"lrdd\" type=\"application/xrd+xml\" template=\"https://localhost:8080/.well-known/webfinger?resource={uri}\"/></XRD>")
		}
		webFinger, err := h.usecase.GetWebFinger(ctx, resource)
		if err != nil {
			return c.JSON(400, "invalid resource")
		}
		return c.JSON(200, webFinger)
	}
}

func (h *UserHandler) PostInbox() echo.HandlerFunc {
	return func(c echo.Context) error {
		body := c.Request().Body
		body_bytes, err := io.ReadAll(body)
		if err != nil {
			return c.JSON(400, "invalid json body")
		}
		inboxCallback := &InboxCallbackFuncs{}
		jsonResolver, err := inboxCallback.GetJsonResolver()
		if err != nil {
			return c.JSON(400, "invalid jsonResolver")
		}

		ctx := c.Request().Context()
		var jsonMap map[string]interface{}
		if err := json.Unmarshal(body_bytes, &jsonMap); err != nil {
			return c.JSON(400, "failed to unmarshal json")
		}
		if err := jsonResolver.Resolve(ctx, jsonMap); err != nil {
			return c.JSON(400, "failed to resolve json")
		}
		return c.JSON(200, "success")
	}
}
