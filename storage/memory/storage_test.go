package memory_storage

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/nickschuch/karma/storage"
)

var (
	name = "nickschuch"
)

func TestInit(t *testing.T) {
	keys := storage.List()
	assert.Contains(t, keys, "memory", "The memory storage is registered.")
}

func TestStorage(t *testing.T) {
	s, _ := storage.New("memory")

	// Get a new user.
	amount := s.Get(name)

	fmt.Println(amount)

	assert.Equal(t, 0, amount, "Get a new amount.")

	// Set a value for the user.
	s.Set(name, 10)
	amount = s.Get(name)
	assert.Equal(t, 10, amount, "Set the amount.")

	// Increase the users karma.
	s.Increase(name, 1)
	amount = s.Get(name)
	assert.Equal(t, 11, amount, "Increase the amount.")

	// Decrease the users karma.
	s.Decrease(name, 1)
	amount = s.Get(name)
	assert.Equal(t, 10, amount, "Decrease the amount.")
}
