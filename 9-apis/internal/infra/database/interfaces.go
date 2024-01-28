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
	GetByID(ID pkgentity.ID) (*entity.Product, error)
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	Update(ID pkgentity.ID, product *entity.Product) error
	Delete(ID pkgentity.ID, product *entity.Product) error
}
