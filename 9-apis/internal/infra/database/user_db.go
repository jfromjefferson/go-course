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

func (userDB *UserDB) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := userDB.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userDB *UserDB) Create(user *entity.User) error {
	return userDB.DB.Create(&user).Error
}

func (userDB *UserDB) Update(user *entity.User) error {
	return userDB.DB.Save(&user).Error
}

func (userDB *UserDB) Delete(user *entity.User) error {
	return userDB.DB.Delete(&user).Error
}
