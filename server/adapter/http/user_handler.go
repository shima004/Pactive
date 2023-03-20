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
		jsonString := "{ \"@context\": [ \"https://www.w3.org/ns/activitystreams\", \"https://w3id.org/security/v1\", { \"manuallyApprovesFollowers\": \"as:manuallyApprovesFollowers\", \"sensitive\": \"as:sensitive\", \"Hashtag\": \"as:Hashtag\", \"quoteUrl\": \"as:quoteUrl\", \"toot\": \"http://joinmastodon.org/ns#\", \"Emoji\": \"toot:Emoji\", \"featured\": \"toot:featured\", \"discoverable\": \"toot:discoverable\", \"schema\": \"http://schema.org#\", \"PropertyValue\": \"schema:PropertyValue\", \"value\": \"schema:value\", \"misskey\": \"https://misskey-hub.net/ns#\", \"_misskey_content\": \"misskey:_misskey_content\", \"_misskey_quote\": \"misskey:_misskey_quote\", \"_misskey_reaction\": \"misskey:_misskey_reaction\", \"_misskey_votes\": \"misskey:_misskey_votes\", \"isCat\": \"misskey:isCat\", \"vcard\": \"http://www.w3.org/2006/vcard/ns#\" } ], \"type\": \"Person\", \"id\": \"https://misskey.io/users/9c1sd6g6p0\", \"inbox\": \"https://misskey.io/users/9c1sd6g6p0/inbox\", \"outbox\": \"https://misskey.io/users/9c1sd6g6p0/outbox\", \"followers\": \"https://misskey.io/users/9c1sd6g6p0/followers\", \"following\": \"https://misskey.io/users/9c1sd6g6p0/following\", \"featured\": \"https://misskey.io/users/9c1sd6g6p0/collections/featured\", \"sharedInbox\": \"https://misskey.io/inbox\", \"endpoints\": { \"sharedInbox\": \"https://misskey.io/inbox\" }, \"url\": \"https://misskey.io/@ShimaPaca\", \"preferredUsername\": \"ShimaPaca\", \"name\": null, \"summary\": null, \"icon\": null, \"image\": null, \"tag\": [], \"manuallyApprovesFollowers\": false, \"discoverable\": true, \"publicKey\": { \"id\": \"https://misskey.io/users/9c1sd6g6p0#main-key\", \"type\": \"Key\", \"owner\": \"https://misskey.io/users/9c1sd6g6p0\", \"publicKeyPem\": \"-----BEGIN PUBLIC KEY-----\nMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAvJvbxyjy96SegUkhSIW5\nre4OhFa2ucFw6u4G2n54LO36h5C0gJl11MYDsMkKWcQrAYAbOXJfYB51e3GSvhxa\n6lmJ3N4F2iMAd9F/leO2Ugrk+DfLQzSAQEtJL4QiKVu/ftuy+rqQGheX1a71g5Ka\nZAirLHqk2iDhbnv6PRzqoTBZpGS98bil6dpMCeGZGdw8EPyWWDc6HmdzQodyUuCl\nChQh2XkY2Af6gBrwo1w1OMON7o+wAcu5UwmqLCgJ7YZsBS2RlzqlpcXMnoWCnMok\nLifms1JlSXUJ14wTSkUay0vIqnwGXQm7feobYBpfWtFc7IxYnprrufzgsaDQd/9c\nuwRPdjaUx7sxE3IA1VJTK+D1FYdS/U+3LCRwN1Ys2ZJulqcH2W+YX9iYwiq6EfiV\nUgA6UtHt1ma+vy/mBjjKie8uPzFj/mYBd3jtEyt9lx/ga3UPJ3jeusHdF1rO7WD0\n3jbsxcUo1y1nAwLIFWH999a8LDhuOzCFb9ZU2GLzgagkYwXm3ZkNUmq1aOklQeyz\nUI3mG62yYbiNqhJodCeX8GVen7GlTSRPJYHaIt9tBkLG/848UcZEASiBACdPSUH8\ngpOu2W5MpTeMUDi/JTz5a+u28v0777yAj+FWCVU4s8LyNUtBDUSvvtIAOJeVLW2g\nbp+QgYtY1KjOUQkjdUsuavMCAwEAAQ==\n-----END PUBLIC KEY-----\n\" }, \"isCat\": false }"
		var jsons map[string]interface{}
		err := json.Unmarshal([]byte(jsonString), &jsons)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(400, "invalid json")
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
