package http

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/shima004/pactive/config"
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

type AddUserRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	PublicKey string `json:"public_key"`
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
		if c.Request().Header.Get("Accept") != "application/activity+json" {
			return c.JSON(400, "invalid accept header")
		}
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
		resource := c.QueryParam("resource")
		if resource == "" {
			c.Logger().Error("resource is empty")
			return c.JSON(400, "invalid resource")
		}
		webfinger, err := h.usecase.GetWebFinger(c.Request().Context(), resource)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(400, "invalid resource")
		}
		return c.JSON(200, webfinger)
	}
}

func (h *UserHandler) GetHostMeta() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := os.Open("assets/host-meta.xml")
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(400, "invalid resource")
		}
		defer file.Close()
		decoder := xml.NewDecoder(file)
		hostMeta := &model.HostMeta{}
		if err := decoder.Decode(hostMeta); err != nil {
			c.Logger().Error(err)
			return c.JSON(400, "invalid resource")
		}
		serverInfo := config.GetServerInfo()
		hostMeta.Link.Template = serverInfo.Protocol + "://" + serverInfo.Host + "/.well-known/webfinger?resource={uri}"
		return c.XML(200, hostMeta)
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
