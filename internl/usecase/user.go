package usecase

import (
	"refactoring/internl/dto"
	"refactoring/internl/entity"
)

type userUseCase struct {
	repo Repository
}

func NewUserUsecase(repo Repository) *userUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (uc *userUseCase) SearchUsers() (*entity.UserStore, error) {
	return uc.repo.GetUsers()
}

func (uc *userUseCase) CreateUser(userDTO *dto.CreateUserRequest) (string, error) {
	user := entity.User{
		DisplayName: userDTO.DisplayName,
		Email:       userDTO.Email,
	}

	return uc.repo.CreateUser(&user)
}

func (us *userUseCase) GetUser(id string) (*entity.User, error) {
	return us.repo.GetUser(id)
}

func (uc *userUseCase) UpdateUser(id string, userDTO *dto.UpdateUserRequest) error {
	user := entity.User{
		DisplayName: userDTO.DisplayName,
	}

	return uc.repo.UpdateUser(id, &user)
}

func (uc *userUseCase) DeleteUser(id string) error {
	return uc.repo.DeleteUser(id)
}
