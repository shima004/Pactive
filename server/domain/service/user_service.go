package service

import (
	"context"

	"github.com/shima004/pactive/domain/model"
	"github.com/shima004/pactive/domain/repository"
)

type IUserService interface {
	AddUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, id int) (*model.User, error)
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

func (s *userService) GetUser(ctx context.Context, id int) (*model.User, error) {
	return s.UserRepository.GetUser(ctx, id)
}
