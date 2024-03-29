package http

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
	"strings"

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
		request := &AddUserRequest{}
		if err := c.Bind(request); err != nil {
			return c.JSON(400, "invalid json body")
		}
		user := &model.User{
			Name:      request.Name,
			Email:     request.Email,
			Password:  request.Password,
			PublicKey: request.PublicKey,
		}
		if err := h.usecase.AddUser(ctx, user); err != nil {
			return c.JSON(400, "invalid user")
		}
		return c.JSON(200, "success")
	}
}

func (h *UserHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		// ctx := c.Request().Context()
		accepts := strings.Split(c.Request().Header.Get("Accept"), ",")
		allowAccepts := []string{"application/activity+json", "application/ld+json"}
		accept := ""
		for _, a := range accepts {
			for _, aa := range allowAccepts {
				if a == aa {
					accept = a
					break
				}
			}
		}
		if accept == "" {
			c.Logger().Info("invalid accept: " + c.Request().Header.Get("Accept"))
			return c.JSON(400, "invalid accept")
		}

		c.Response().Header().Set("Content-Type", accept)

		file, err := os.Open("assets/misskey.json")
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(400, "invalid id")
		}
		defer file.Close()
		decoder := json.NewDecoder(file)
		var jsons map[string]interface{}
		if err := decoder.Decode(&jsons); err != nil {
			c.Logger().Error(err)
			return c.JSON(400, "invalid id")
		}
		// id := c.Param("id")
		// user, err := h.usecase.GetUser(ctx, id)
		// if err != nil {
		// 	c.Logger().Info(err.Error() + " id: " + id)
		// 	return c.JSON(400, "invalid id")
		// }
		// json, err := user.Serialize()
		// if err != nil {
		// 	c.Logger().Error(err.Error() + " id: " + id)
		// 	return c.JSON(400, "invalid id")
		// }
		return c.JSON(200, jsons)
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
