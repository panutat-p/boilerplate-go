package usecase

import (
	"context"
	"errors"
	"log/slog"

	"boilerplate-go/internal/model"
)

func (u *UseCase) GetFruits(ctx context.Context) ([]model.Fruit, error) {
	fruits, err := u.store.ReadFruitFile(ctx)
	if err != nil {
		slog.Error("Failed to ReadFruitFile", slog.Any("err", err))
		return nil, err
	}

	err = u.CheckFruits(ctx, fruits)
	if err != nil {
		slog.Error("Failed to CheckFruits", slog.Any("err", err))
		return nil, err
	}

	return fruits, nil
}

func (u *UseCase) CheckFruits(ctx context.Context, fruits []model.Fruit) error {
	if len(fruits) == 0 {
		slog.ErrorContext(ctx, "empty fruits")
		return errors.New("empty fruits")
	}

	return nil
}
