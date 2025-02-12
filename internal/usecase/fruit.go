package usecase

import (
	"context"
	"errors"
	"log/slog"
	"runtime"

	"golang.org/x/sync/errgroup"

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

func (u *UseCase) WriteFruits(ctx context.Context, fruits []model.Fruit) error {
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(runtime.NumCPU())
	for _, fruit := range fruits {
		fruit := fruit
		g.Go(func() error {
			err := u.store.WriteFruitFile(ctx, []model.Fruit{fruit})
			if err != nil {
				slog.ErrorContext(
					ctx,
					"Failed to WriteFruitFile for fruit",
					slog.Any("fruit", fruit),
					slog.Any("err", err),
				)
				return err
			}
			return nil
		})
	}
	err := g.Wait()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to WriteFruits", slog.Any("err", err))
		return err
	}

	slog.InfoContext(ctx, "Succeeded to WriteFruits")
	return nil
}
