package usecase

//go:generate mockgen -source=usecase.go -destination=mock_usecase/usecase.go -package=mock_usecase

import (
	"context"

	"boilerplate-go/internal/config"
	"boilerplate-go/internal/model"
	"boilerplate-go/internal/store"
)

type IUseCase interface {
	GetFruits(ctx context.Context) ([]model.Fruit, error)
	CheckFruits(ctx context.Context, fruits []model.Fruit) error
}

type UseCase struct {
	config *config.Config
	store  store.IStore
}

func NewUseCase(config *config.Config, store store.IStore) *UseCase {
	return &UseCase{
		config: config,
		store:  store,
	}
}
