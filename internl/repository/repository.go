package repository

import "refactoring/internl/entity"

type Repository interface {
	GetUsers() (*entity.UserStore, error)
	CreateUser(*entity.User) (string, error)
	GetUser(string) (*entity.User, error)
	UpdateUser(string, *entity.User) error
	DeleteUser(string) error
}
