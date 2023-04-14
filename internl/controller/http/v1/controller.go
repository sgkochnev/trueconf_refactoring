package v1

import (
	"refactoring/internl/dto"
	"refactoring/internl/entity"
)

type UserUseCase interface {
	SearchUsers() (*entity.UserStore, error)
	CreateUser(*dto.CreateUserRequest) (string, error)
	GetUser(string) (*entity.User, error)
	UpdateUser(string, *dto.UpdateUserRequest) error
	DeleteUser(string) error
}

type UseCase interface {
	UserUseCase
}
