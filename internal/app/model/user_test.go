package model_test

import (
	"strings"
	"testing"

	"github.com/F7icK/LService/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct{
		name string
		u func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty telephone",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Telephone = ""

				return u
			},
			isValid: false,
		},
		{
			name: "short telephone",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Telephone = "+7123456789"

				return u
			},
			isValid: false,
		},
		{
			name: "long telephone",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Telephone = "+700123456789"

				return u
			},
			isValid: false,
		},
		{
			name: "invalid characters in the telephone",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Telephone = "+7s123456789"

				return u
			},
			isValid: false,
		},
		{
			name: "short name",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Name = "A"

				return u
			},
			isValid: false,
		},
		{
			name: "short surname",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Surname = "A"

				return u
			},
			isValid: false,
		},
		{
			name: "long name",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Name = strings.Repeat("a", 256)
				return u
			},
			isValid: false,
		},
		{
			name: "long surname",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Surname = strings.Repeat("a", 256)

				return u
			},
			isValid: false,
		},
		{
			name: "invalid characters in the name",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Name = "Ильян1"

				return u
			},
			isValid: false,
		},
		{
			name: "invalid characters in the surname",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Surname = "1Ахметов1"

				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
