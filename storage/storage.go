package storage

import (
	"errors"
	"fmt"
)

type Storage interface {
	Get(string) int
	Set(string, int)
  Increase(string, int)
  Decrease(string, int)
}

var (
	storages    map[string]Storage
	ErrNotFound = errors.New("Could not find the storage.")
)

func init() {
	storages = make(map[string]Storage)
}

func Register(name string, storage Storage) error {
	if _, exists := storages[name]; exists {
		return fmt.Errorf("Scheme already registered %s", name)
	}
	storages[name] = storage

	return nil
}

func New(name string) (Storage, error) {
	if p, exists := storages[name]; exists {
		return p, nil
	}

	return nil, ErrNotFound
}
