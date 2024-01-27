package database

import (
	"github.com/jfromjefferson/gi-course-9/internal/entity"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *UserDB {
	return &UserDB{
		DB: db,
	}
}

func (userDB *UserDB) Create(user *entity.User) error {
	return userDB.DB.Create(user).Error
}

func (userDB *UserDB) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := userDB.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
