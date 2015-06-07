package dynamodb_storage

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/nickschuch/karma/storage"
)

func TestInit(t *testing.T) {
	keys := storage.List()
	assert.Contains(t, keys, "dynamodb", "The dynamodb storage is registered.")
}

func TestStorage(t *testing.T) {
	// Nothing to see here.
}
