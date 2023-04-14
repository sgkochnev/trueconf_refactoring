package usecase

var _ UserUseCase = (*manager)(nil)

type manager struct {
	UserUseCase
}

func NewManager(repo Repository) *manager {
	return &manager{
		UserUseCase: NewUserUsecase(repo),
	}
}
