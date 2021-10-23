package model

import (
	"testing"
)

// TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		ID: 1,
		Name: "Ильян",
		Surname: "Ахметов",
		Telephone: "+79991316119",
	}
}