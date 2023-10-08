package service

import "github.com/frangar97/testapi/internal/repository"

type Service struct {
	UserService UserService
}

func NewService(repositories repository.Repository) Service {
	return Service{
		UserService: userServiceImpl{userRepository: repositories.UserRepository},
	}
}