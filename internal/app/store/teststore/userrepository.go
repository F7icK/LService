package teststore

import (
	"github.com/F7icK/LService/internal/app/model"
	"github.com/F7icK/LService/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
	use []*model.User
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil{
		return err
	}

	r.users[u.Telephone] = u
	u.ID = len(r.users)

	return nil
}

// FindByTelephone ...
func (r *UserRepository) FindByTelephone(telephone string) (*model.User, error) {
	u, ok := r.users[telephone]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}


func (r *UserRepository) AllSelect() ([]*model.User, error) {

	return r.use, nil
}

func (r *UserRepository) DeleteFromID(id int) (*model.User, error) {

	for _, us := range r.users {
		if us.ID == id {
			return us, nil
		}
	}

	return nil, store.ErrRecordNotFound
}