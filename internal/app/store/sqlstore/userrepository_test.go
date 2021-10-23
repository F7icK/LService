package sqlstore_test

import (
	"testing"

	"github.com/F7icK/LService/internal/app/model"
	"github.com/F7icK/LService/internal/app/store"
	"github.com/F7icK/LService/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByTelephone(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	telephone := "+79991111111"
	_, err := s.User().FindByTelephone(telephone)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Telephone = telephone
	s.User().Create(u)
	u, err = s.User().FindByTelephone(telephone)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_AllSelect(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	s.User().Create(u)
	us, err := s.User().AllSelect()
	assert.NoError(t, err)
	assert.NotNil(t, us)
}

func TestUserRepository_DeleteFromID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	uid1 := model.TestUser(t)

	s.User().Create(uid1)
	u, err := s.User().DeleteFromID(uid1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}