package postgres

import (
	"github.com/shima004/pactive/domain/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) AddUser(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUser(id int) (*model.User, error) {
	var user model.User
	err := r.DB.First(&user, id).Error
	return &user, err
}
