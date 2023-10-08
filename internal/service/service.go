package service

import "github.com/frangar97/testapi/internal/repository"

type Service struct {
	UserService   UserService
	DeviceService DeviceService
}

func NewService(repositories repository.Repository, secret string) Service {
	return Service{
		UserService:   userServiceImpl{userRepository: repositories.UserRepository, secret: secret},
		DeviceService: deviceServiceImpl{deviceRepository: repositories.DeviceRepository},
	}
}
