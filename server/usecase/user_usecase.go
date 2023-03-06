package usecase

import (
	"context"

	"github.com/shima004/pactive/domain/model"
	"github.com/shima004/pactive/domain/service"
)

type IUserUsecase interface {
	AddUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, id int) (*model.User, error)
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
	// TODO: ユーザー登録時のバリデーション処理を記述する
	return u.UserService.AddUser(ctx, user)
}

func (u *userUsecase) GetUser(ctx context.Context, id int) (*model.User, error) {
	// TODO: ユーザー取得時のバリデーション処理を記述する
	return u.UserService.GetUser(ctx, id)
}
