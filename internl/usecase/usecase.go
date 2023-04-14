package usecase

import (
	"refactoring/internl/dto"
	"refactoring/internl/entity"
)

type Repository interface {
	GetUsers() (*entity.UserStore, error)
	CreateUser(*entity.User) (string, error)
	GetUser(string) (*entity.User, error)
	UpdateUser(string, *entity.User) error
	DeleteUser(string) error
}

type UserUseCase interface {
	SearchUsers() (*entity.UserStore, error)
	CreateUser(*dto.CreateUserRequest) (string, error)
	GetUser(string) (*entity.User, error)
	UpdateUser(string, *dto.UpdateUserRequest) error
	DeleteUser(string) error
}
