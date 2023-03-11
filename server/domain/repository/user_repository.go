package repository

import (
	"context"

	"github.com/go-fed/activity/streams/vocab"
	"github.com/shima004/pactive/domain/model"
)

type IUserRepository interface {
	AddUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, resource string) (vocab.ActivityStreamsPerson, error)
	GetWebFinger(ctx context.Context, resource string) (*model.WebFinger, error)
}
