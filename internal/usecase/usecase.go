package usecase

//go:generate mockgen -source=usecase.go -destination=mock_usecase/usecase.go -package=mock_usecase

import (
	"context"

	"boilerplate-go/internal/model"
	"boilerplate-go/internal/store"
)

type IUseCase interface {
	GetFruits(ctx context.Context) ([]model.Fruit, error)
}

type UseCase struct {
	store store.IStore
}

func NewUseCase(store store.IStore) *UseCase {
	return &UseCase{
		store: store,
	}
}

func (u *UseCase) GetFruits(context.Context) ([]model.Fruit, error) {
	fruits, err := u.store.ReadFruitFile(nil)
	if err != nil {
		return nil, err
	}

	return fruits, nil
}
