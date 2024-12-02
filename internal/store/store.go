package store

//go:generate mockgen -source=store.go -destination=mock_store/store.go -package=mock_store

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"boilerplate-go/internal/model"
)

type IStore interface {
	ReadFruitFile(ctx context.Context) ([]model.Fruit, error)
}

type Store struct {
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) ReadFruitFile(context.Context) ([]model.Fruit, error) {
	file, err := os.Open("./data/fruit.json")
	if err != nil {
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
