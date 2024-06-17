package repository

import (
	"fmt"

	"github.com/vibinnair/user-management-in-mem/model"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUser(id int) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id string) error
}

type UserRepositoryImpl struct {
	users map[int]*model.User
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{
		users: make(map[int]*model.User),
	}
}

func (userRepository *UserRepositoryImpl) CreateUser(user *model.User) error {
	userRepository.users[user.ID] = user
	return nil
}

func (userRepository *UserRepositoryImpl) GetUser(id int) (*model.User, error) {
	user, exists := userRepository.users[id]

	if !exists {
		return nil, fmt.Errorf("User Not Found")
	}
	return user, nil
}

func (userRepository *UserRepositoryImpl) GetAllUsers() ([]*model.User, error) {
	return nil, nil
}

func (userRepository *UserRepositoryImpl) UpdateUser(user *model.User) error {
	return nil
}

func (userRepository *UserRepositoryImpl) DeleteUser(id string) error {
	return nil
}
