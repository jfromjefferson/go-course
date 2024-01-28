package database

import (
	"github.com/jfromjefferson/gi-course-9/internal/entity"
	pkgentity "github.com/jfromjefferson/gi-course-9/pkg/entity"
	"gorm.io/gorm"
)

type ProductDB struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *ProductDB {
	return &ProductDB{
		DB: db,
	}
}

func (productDB *ProductDB) Create(product *entity.Product) error {
	return productDB.DB.Create(&product).Error
}

func (productDB *ProductDB) GetByID(ID pkgentity.ID) (*entity.Product, error) {
	var product entity.Product
	err := productDB.DB.Where("id = ?", ID).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (productDB *ProductDB) Update(ID pkgentity.ID, product *entity.Product) error {
	_, err := productDB.GetByID(ID)
	if err != nil {
		return err
	}

	return productDB.DB.Save(product).Error
}

func (productDB *ProductDB) Delete(ID pkgentity.ID, product *entity.Product) error {
	_, err := productDB.GetByID(ID)
	if err != nil {
		return err
	}
	return productDB.DB.Delete(product).Error
}
