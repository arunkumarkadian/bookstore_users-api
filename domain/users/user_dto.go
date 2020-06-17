package users

import (
	"github.com/arunkumarkadian/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate()  *errors.RestErr {
	// Trimming spaces in email id field and converting email id to lowercase
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}

	return nil
}