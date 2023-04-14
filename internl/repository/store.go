package repository

import (
	"refactoring/config"
	jrepo "refactoring/internl/repository/json_repo"
)

var _ Repository = (*store)(nil)

type store struct {
	Repository
}

func NewStore(cfg *config.Config) (*store, error) {
	stroe, err := jrepo.NewJsonStore(cfg.JsonStore.Name)
	if err != nil {
		return nil, err
	}
	return &store{
		Repository: stroe,
	}, nil
}
