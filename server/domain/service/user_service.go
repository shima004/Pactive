package service

import (
	"context"

	"github.com/go-fed/activity/streams/vocab"
	"github.com/shima004/pactive/domain/model"
	"github.com/shima004/pactive/domain/repository"
)

type IUserService interface {
	AddUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, resource string) (vocab.ActivityStreamsPerson, error)
	GetWebFinger(ctx context.Context, resource string) (*model.WebFinger, error)
}

type userService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &userService{
		UserRepository: userRepository,
	}
}

func (s *userService) AddUser(ctx context.Context, user *model.User) error {
	return s.UserRepository.AddUser(ctx, user)
}

func (s *userService) GetUser(ctx context.Context, resource string) (vocab.ActivityStreamsPerson, error) {
	return s.UserRepository.GetUser(ctx, resource)
}

func (s *userService) GetWebFinger(ctx context.Context, resource string) (*model.WebFinger, error) {
	return s.UserRepository.GetWebFinger(ctx, resource)
}
