package service

import (
	"errors"
	"fmt"

	"github.com/frangar97/testapi/internal/entities"
	"github.com/frangar97/testapi/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(string, string) error
}

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func (u userServiceImpl) CreateUser(username string, password string) error {
	user, err := u.userRepository.GetUserByUsername(username)

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	if user != nil {
		return fmt.Errorf("user %s already exist ", username)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	newUser := &entities.User{UserName: username, Password: string(hashPassword)}

	err = u.userRepository.CreateUser(newUser)
	if err != nil {
		return err
	}

	return nil
}
