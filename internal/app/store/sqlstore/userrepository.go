package sqlstore

import (
	"database/sql"

	"github.com/F7icK/LService/internal/app/model"
	"github.com/F7icK/LService/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil{
		return  err
	}

	return r.store.db.QueryRow(
			"INSERT INTO users (name, surname, telephone) VALUES ($1, $2, $3) RETURNING id",
			u.Name,
			u.Surname,
			u.Telephone,
		).Scan(&u.ID)
}

// FindByTelephone ...
func (r *UserRepository) FindByTelephone(telephone string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT * FROM users WHERE telephone = $1",
		telephone,
		).Scan(
			&u.ID,
			&u.Name,
			&u.Surname,
			&u.Telephone,
			); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

// AllSelect ...
func (r *UserRepository) AllSelect() ([]*model.User, error) {
	rows, err := r.store.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	us := make([]*model.User, 0)

	for rows.Next() {
		u := new(model.User)
		err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.Telephone)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return us, nil
}

// DeleteFromID ...
func (r *UserRepository) DeleteFromID(id int) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"DELETE FROM users WHERE id  = $1 RETURNING id, name, surname, telephone",
		id,
	).Scan(
		&u.ID,
		&u.Name,
		&u.Surname,
		&u.Telephone,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}