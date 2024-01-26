package entity

import (
	"github.com/jfromjefferson/gi-course-9/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        entity.ID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
}

func NewUser(firstName, lastName, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		ID:        entity.NewID(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  string(hash),
	}

	return &user, nil
}

func (user *User) ValidatePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}
