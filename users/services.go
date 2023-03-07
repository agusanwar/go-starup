package users

import "golang.org/x/crypto/bcrypt"

// interface
type Services interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

// struct tidak dapat diakses dari package lain
type services struct {
	// mapping struct input ke struct user
	repository Repository
}

// untuk membuat struct service membutuhkan sebuah function
func NewServices(repository Repository) *services {
	return &services{repository}
}

// mapping struct input to struct user
func (s *services) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Passion = input.Passion
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(password)
	user.Roles = "User"

	// save struct user to repository
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, err
}
