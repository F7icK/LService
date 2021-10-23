package teststore

import (
	"github.com/F7icK/LService/internal/app/model"
	"github.com/F7icK/LService/internal/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
		use: make([]*model.User, 0),
	}

	return s.userRepository
}
