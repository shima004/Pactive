package usecase

import (
	"context"
	"log"
	"strings"

	"github.com/go-fed/activity/streams/vocab"
	"github.com/shima004/pactive/domain/model"
	"github.com/shima004/pactive/domain/service"
)

type IUserUsecase interface {
	AddUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, resource string) (vocab.ActivityStreamsPerson, error)
	GetWebFinger(ctx context.Context, resource string) (*model.WebFinger, error)
}

type userUsecase struct {
	UserService service.IUserService
}

func NewUserUsecase(userService service.IUserService) IUserUsecase {
	return &userUsecase{
		UserService: userService,
	}
}

func (u *userUsecase) AddUser(ctx context.Context, user *model.User) error {
	return u.UserService.AddUser(ctx, user)
}

func (u *userUsecase) GetUser(ctx context.Context, resource string) (vocab.ActivityStreamsPerson, error) {
	return u.UserService.GetUser(ctx, resource)
}

func (u *userUsecase) GetWebFinger(ctx context.Context, resource string) (*model.WebFinger, error) {
	log.Println("GetWebFinger", resource)
	name := strings.Split(resource, "@")[0]
	name = strings.Split(name, ":")[1]
	return u.UserService.GetWebFinger(ctx, name)
}
