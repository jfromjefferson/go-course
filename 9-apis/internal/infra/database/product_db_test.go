package database

import (
	"github.com/jfromjefferson/gi-course-9/internal/entity"
	pkgentity "github.com/jfromjefferson/gi-course-9/pkg/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestProductDB_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Bala de goma", "BLADEGOMA", 1500)
	assert.Nil(t, err)
	assert.NotNil(t, product)

	productDB := NewProduct(db)
	assert.NotNil(t, productDB)

	err = productDB.Create(product)
	assert.Nil(t, err)

	product, err = entity.NewProduct("Bala de goma", "BLADEGOMA", -1000)
	assert.Nil(t, product)
	assert.NotNil(t, err)
}

func TestProductDB_GetByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Bala de goma", "BLADEGOMA", 1500)
	assert.Nil(t, err)
	assert.NotNil(t, product)

	productDB := NewProduct(db)
	assert.NotNil(t, productDB)

	err = productDB.Create(product)
	assert.Nil(t, err)

	productFound, err := productDB.GetByID(product.ID)
	assert.Nil(t, err)
	assert.NotNil(t, productFound)
	assert.Equal(t, productFound.Code, "BLADEGOMA")

	productFound, err = productDB.GetByID(pkgentity.NewID())
	assert.Nil(t, productFound)
	assert.NotNil(t, err)
}

func TestProductDB_Update(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Bala de goma", "BLADEGOMA", 1500)
	assert.Nil(t, err)
	assert.NotNil(t, product)

	productDB := NewProduct(db)
	assert.NotNil(t, productDB)

	err = productDB.Create(product)
	assert.Nil(t, err)

	productFound, err := productDB.GetByID(product.ID)
	assert.Nil(t, err)
	assert.NotNil(t, productFound)

	productFound.Name = "Caneca azul"
	productFound.Code = "CAZUL"

	err = productDB.Update(productFound.ID, productFound)
	assert.Nil(t, err)

	productID := pkgentity.NewID()
	productFound, err = productDB.GetByID(productID)
	assert.Nil(t, productFound)
	assert.NotNil(t, err)

	err = productDB.Update(productID, productFound)
	assert.NotNil(t, err)
}

func TestProductDB_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Bala de goma", "BLADEGOMA", 1500)
	assert.Nil(t, err)
	assert.NotNil(t, product)

	productDB := NewProduct(db)
	assert.NotNil(t, productDB)

	err = productDB.Create(product)
	assert.Nil(t, err)

	productFound, err := productDB.GetByID(product.ID)
	assert.Nil(t, err)
	assert.NotNil(t, productFound)

	err = productDB.Delete(productFound.ID, productFound)
	assert.Nil(t, err)

	productID := pkgentity.NewID()
	productFound, err = productDB.GetByID(productID)
	assert.NotNil(t, err)
	assert.Nil(t, productFound)

	err = productDB.Delete(productID, productFound)
	assert.NotNil(t, err)

}
