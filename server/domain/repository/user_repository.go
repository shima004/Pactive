package repository

import (
	"context"

	"github.com/shima004/pactive/domain/model"
)

type IUserRepository interface {
	AddUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, id int) (*model.User, error)
}
