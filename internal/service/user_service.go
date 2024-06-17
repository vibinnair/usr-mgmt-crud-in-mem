package service

import (
	"github.com/vibinnair/user-management-in-mem/internal/repository"
	"github.com/vibinnair/user-management-in-mem/model"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (service *UserService) CreateUser(user *model.User) error {
	return service.userRepository.CreateUser(user)
}

func (service *UserService) GetUser(id int) (*model.User, error) {
	return service.userRepository.GetUser(id)
}

func (service *UserService) UpdateUser(user *model.User) error {
	return service.userRepository.UpdateUser(user)
}

func (service *UserService) DeleteUser(id string) error {
	return service.DeleteUser(id)
}
