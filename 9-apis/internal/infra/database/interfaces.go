package database

import (
	"github.com/jfromjefferson/gi-course-9/internal/entity"
	pkgentity "github.com/jfromjefferson/gi-course-9/pkg/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
	Delete(user *entity.User) error
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindByID(ID pkgentity.ID) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(product *entity.Product) error
}
