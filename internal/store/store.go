package store

//go:generate mockgen -source=store.go -destination=mock_store/store.go -package=mock_store

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log/slog"
	"os"

	"boilerplate-go/internal/config"
	"boilerplate-go/internal/model"
)

type IStore interface {
	ReadFruitFile(ctx context.Context) ([]model.Fruit, error)
}

type Store struct {
	config *config.Config
}

func NewStore(config *config.Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) ReadFruitFile(ctx context.Context) ([]model.Fruit, error) {
	file, err := os.Open("./data/fruit.json")
	if err != nil {
		slog.ErrorContext(ctx, "Failed to Open a file", slog.Any("err", err))
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var fruits []model.Fruit
	err = json.Unmarshal(byteValue, &fruits)
	if err != nil {
		return nil, err
	}

	return fruits, nil
}
