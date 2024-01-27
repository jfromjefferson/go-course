package database

import (
	"github.com/jfromjefferson/gi-course-9/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestUser_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)

	db.AutoMigrate(&entity.User{})

	user, err := entity.NewUser("Jefferson", "Silva", "email@email.com", "12345")
	assert.Nil(t, err)
	assert.NotNil(t, user)

	userDB := NewUser(db)
	assert.NotNil(t, userDB)

	err = userDB.Create(user)

	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail("email@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.NotEmpty(t, userFound.FirstName)
	assert.NotEmpty(t, userFound.LastName)
	assert.NotEmpty(t, userFound.Email)

	userFound, err = userDB.FindByEmail("email@email.com.br")
	assert.NotNil(t, err)
	assert.Nil(t, userFound)

}
