package dynamodb_storage

import (
  storage "github.com/nickschuch/karma/storage"
)

type DynamoDBStorage struct{}

func init() {
	storage.Register("dynamodb", &DynamoDBStorage{})
}

func (p *DynamoDBStorage) Get(n string) int {
	return 0
}

func (p *DynamoDBStorage) Set(n string, v int) {
	// Nothing for now.
}

func (p *DynamoDBStorage) Increase(n string, v int) {
	// Nothing for now.
}

func (p *DynamoDBStorage) Decrease(n string, v int) {
	// Nothing for now.
}
