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

func (productDB *ProductDB) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = productDB.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = productDB.DB.Order("created_at " + sort).Find(&products).Error
	}

	return products, err
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
