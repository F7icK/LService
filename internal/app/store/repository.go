package store

import (
	"github.com/F7icK/LService/internal/app/model"
)

// UserRepository ...
type UserRepository interface {
	Create(user *model.User) error
	FindByTelephone(string) (*model.User, error)
	AllSelect() ([]*model.User, error)
	DeleteFromID(id int) (*model.User, error)
}
