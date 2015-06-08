package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	name = "nickschuch"
)

func TestGetUser(t *testing.T) {
	var user string

	// Small increase.
	user, _ = getUser(name + "++")
	assert.Equal(t, name, user, "Found user '"+name+"'.")

	// Large increase.
	user, _ = getUser(name + "+=10")
	assert.Equal(t, name, user, "Found user '"+name+"'.")

	// Small decrease.
	user, _ = getUser(name + "--")
	assert.Equal(t, name, user, "Found user '"+name+"'.")

	// Large decrease.
	user, _ = getUser(name + "-=10")
	assert.Equal(t, name, user, "Found user '"+name+"'.")
}

func TestIncreaseAmount(t *testing.T) {
	amount := increaseAmount(name + "++")
	assert.Equal(t, 1, amount, "Found small increase.")
	amount = increaseAmount(name + "+=10")
	assert.Equal(t, 10, amount, "Found large increase.")
}

func TestDecreaseAmount(t *testing.T) {
	amount := decreaseAmount(name + "--")
	assert.Equal(t, 1, amount, "Found small decrease.")
	amount = decreaseAmount(name + "-=10")
	assert.Equal(t, 10, amount, "Found large decrease.")
}
