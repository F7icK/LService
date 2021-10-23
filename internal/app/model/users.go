package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// User ...
type User struct {
	ID int `json:"id"`
	Name string	`json:"name"`
	Surname string `json:"surname"`
	Telephone string `json:"telephone"`
}


func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Name, validation.Required, is.UTFLetter, validation.Length(2,255)),
		validation.Field(&u.Surname, validation.Required, is.UTFLetter, validation.Length(2,255)),
		validation.Field(&u.Telephone, validation.Required, is.E164, validation.Length(12,12)),
	)
}