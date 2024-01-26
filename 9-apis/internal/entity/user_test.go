package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John", "Doe", "email@email.com", "12345")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Contains(t, user.Email, "@")
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John", "Doe", "email@email.com", "12345")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("12345"))
	assert.False(t, user.ValidatePassword("123456"))
	assert.NotEqual(t, user.Password, "12345")
}
