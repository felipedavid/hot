package types

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var emailRx = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

type UpdateUserParams struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	Password  *string `json:"password"`

	UpdateUserParamsErrors
}

type UpdateUserParamsErrors struct {
	FirstNameError string `json:"first_name,omitempty"`
	LastNameError  string `json:"last_name,omitempty"`
	EmailError     string `json:"email,omitempty"`
	PasswordError  string `json:"password,omitempty"`
}

func (p *UpdateUserParams) Validate() (valid bool) {
	valid = true

	if p.Password != nil && len(*p.Password) < 8 {
		p.PasswordError = "email too short"
		valid = false
	}

	if p.FirstName != nil && len(*p.FirstName) < 2 {
		p.FirstNameError = "first name too short"
		valid = false
	}

	if p.LastName != nil && len(*p.LastName) < 2 {
		p.LastNameError = "last name too short"
		valid = false
	}

	if p.Email != nil && !emailRx.MatchString(*p.Email) {
		p.EmailError = "invalid email address"
		valid = false
	}

	return valid
}

type CreateUserParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`

	CreateUserParamsErrors
}

type CreateUserParamsErrors struct {
	FirstNameError string `json:"first_name,omitempty"`
	LastNameError  string `json:"last_name,omitempty"`
	EmailError     string `json:"email,omitempty"`
	PasswordError  string `json:"password,omitempty"`
}

func (p *CreateUserParams) Validate() (valid bool) {
	if len(p.Password) < 8 {
		p.PasswordError = "email too short"
	}

	if len(p.FirstName) < 2 {
		p.FirstNameError = "first name too short"
	}

	if len(p.LastName) < 2 {
		p.LastNameError = "last name too short"
	}

	if !emailRx.MatchString(p.Email) {
		p.EmailError = "invalid email address"
	}

	return p.PasswordError == "" && p.FirstNameError == "" && p.LastNameError == "" && p.EmailError == ""
}

type User struct {
	ID             string `bson:"_id,omitempty" json:"id"`
	FirstName      string `bson:"first_name" json:"first_name"`
	LastName       string `bson:"last_name" json:"last_name"`
	Email          string `bson:"email" json:"email"`
	HashedPassword string `bson:"hashed_password" json:"-"`
}

func NewUser(params *CreateUserParams) (*User, error) {
	hPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:      params.FirstName,
		LastName:       params.LastName,
		Email:          params.Email,
		HashedPassword: string(hPassword),
	}, nil
}
