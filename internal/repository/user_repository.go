package repository

import (
	"errors"

	"github.com/frangar97/testapi/internal/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(*entities.User) error
	GetUserByUsername(string) (*entities.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func (u userRepositoryImpl) CreateUser(user *entities.User) error {
	result := u.db.Create(user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user was not created")
	}

	return nil
}

func (u userRepositoryImpl) GetUserByUsername(username string) (*entities.User, error) {
	user := &entities.User{}

	result := u.db.First(user, "user_name=?", username)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
