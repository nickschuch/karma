package memory_storage

import (
	storage "github.com/nickschuch/karma/storage"
)

type User struct {
	Name  string
	Value int
}

type MemoryStorage struct {
	Users map[string]User
}

func init() {
	users := make(map[string]User)
	storage.Register("memory", &MemoryStorage{Users: users})
}

func (s *MemoryStorage) Get(n string) int {
	user := exist(s.Users, n)
	s.Users[n] = user
	return user.Value
}

func (s *MemoryStorage) Set(n string, v int) {
	user := User{
		Name:  n,
		Value: v,
	}
	s.Users[n] = user
}

func (s *MemoryStorage) Increase(n string, v int) {
	user := exist(s.Users, n)
	user.Value = user.Value + v
	s.Users[n] = user
}

func (s *MemoryStorage) Decrease(n string, v int) {
	user := exist(s.Users, n)
	user.Value = user.Value - v
	s.Users[n] = user
}

func exist(u map[string]User, n string) User {
	// Check if the user exists in the map.
	if val, ok := u[n]; ok {
		return val
	}

	// Else we create a new user and use that instead.
	user := User{
		Name:  n,
		Value: 0,
	}
	return user
}
